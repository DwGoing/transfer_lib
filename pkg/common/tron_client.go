package common

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"time"

	"github.com/ahmetb/go-linq"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type TronClient struct {
	ChainClient
	apiKeys []string
}

func NewTronClient(nodes map[string]int, apiKeys []string) *TronClient {
	return &TronClient{
		ChainClient: *newChainClient(Chain_TRON, nodes),
		apiKeys:     apiKeys,
	}
}

// @title	获取Tron客户端
// @param	Self	*TronClient			模块实例
// @return	_		*client.GrpcClient	客户端
// @return	_		error				异常信息
func (Self *TronClient) getTronRpcClient() (*client.GrpcClient, error) {
	sum := 0
	for _, v := range Self.nodes {
		sum += v
	}
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(sum)
	var node string
	for k, v := range Self.nodes {
		if v >= i {
			node = k
			break
		}
		i = i - v
	}
	client := client.NewGrpcClient(node)
	apiKey := Self.apiKeys[rand.Int()%len(Self.apiKeys)]
	err := client.SetAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	err = client.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (Self *TronClient) GetBalance(address string, addressType AddressType, currency Currency) (float64, error) {
	switch addressType {
	case AddressType_TRON:
		client, err := Self.getTronRpcClient()
		if err != nil {
			return 0, err
		}
		switch currency {
		case Currency_TRX:
			account, err := client.GetAccount(address)
			if err != nil {
				return 0, err
			}
			balance, _ := new(big.Float).Quo(new(big.Float).SetInt64(account.Balance), big.NewFloat(1e6)).Float64()
			return balance, nil
		case Currency_USDT:
			balanceBigInt, err := client.TRC20ContractBalance(address, "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
			if err != nil {
				return 0, err
			}
			balance, _ := new(big.Float).Quo(new(big.Float).SetInt(balanceBigInt), big.NewFloat(1e6)).Float64()
			return balance, nil
		default:
			return 0, errors.New("unsupported currency")
		}
	default:
		return 0, errors.New("unsupported address type")
	}
}

type Tron struct{}

// @title	发送Tron交易
// @param	Self		*Tron					模块实例
// @param	client		*client.GrpcClient		客户端
// @param	privateKey	*ecdsa.PrivateKey		私钥
// @param	tx			*core.Transaction		交易
// @param	waitReceipt	*client.GrpcClient		是否等待结果
// @return	_			*core.TransactionInfo	交易信息
// @return	_			error					异常信息
func (Self *Tron) SendTronTransaction(client *client.GrpcClient, privateKey *ecdsa.PrivateKey, tx *core.Transaction, waitReceipt bool) (*core.TransactionInfo, error) {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}
	tx.Signature = append(tx.Signature, signature)
	result, err := client.Broadcast(tx)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, fmt.Errorf("bad transaction: %v", string(result.GetMessage()))
	}
	var transaction *core.TransactionInfo
	start := 0
	for {
		if start++; start > 10 {
			return nil, errors.New("transaction info not found")
		}
		transaction, err = client.GetTransactionInfoByID(common.BytesToHexString(hash))
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		if transaction.Result != 0 {
			return nil, errors.New(string(transaction.ResMessage))
		}
		break
	}
	return transaction, err
}

// @title	解析交易
// @param	Self		*Tron				模块实例
// @param	txHash		string				交易Hash
// @return	_			*Transaction	交易信息
// @return	_			error				异常信息
func (Self *Tron) DecodeTronTransaction(client *client.GrpcClient, txHash string) (*Transaction, error) {
	result := Transaction{
		Chain: Chain_TRON,
		Hash:  txHash,
	}
	tx, err := client.GetTransactionInfoByID(txHash)
	if err != nil {
		return nil, err
	}
	result.Height = tx.BlockNumber
	result.TimeStamp = tx.GetBlockTimeStamp()
	coreTx, err := client.GetTransactionByID(txHash)
	if err != nil {
		return nil, err
	}
	contracts := coreTx.RawData.GetContract()
	if len(contracts) < 1 {
		return nil, errors.New("not transfer transaction")
	}
	if contracts[0].Type != core.Transaction_Contract_TransferContract &&
		contracts[0].Type != core.Transaction_Contract_TriggerSmartContract {
		return nil, errors.New("not transfer transaction")
	}
	if tx.ContractAddress == nil {
		var contract core.TransferContract
		err = contracts[0].GetParameter().UnmarshalTo(&contract)
		if err != nil {
			return nil, err
		}
		result.From = common.EncodeCheck(contract.GetOwnerAddress())
		result.To = common.EncodeCheck(contract.GetToAddress())
		result.Amount, _ = new(big.Float).Quo(new(big.Float).SetInt64(contract.GetAmount()), big.NewFloat(1e6)).Float64()
	} else {
		contractAddress := common.EncodeCheck(tx.ContractAddress)
		result.Contract = &contractAddress
		var contract core.TriggerSmartContract
		err = contracts[0].GetParameter().UnmarshalTo(&contract)
		if err != nil {
			return nil, err
		}
		logs := tx.GetLog()
		if len(logs) < 1 {
			return nil, errors.New("not transfer transaction")
		}
		log := logs[0]
		topics := log.GetTopics()
		if len(topics) < 3 {
			return nil, errors.New("not transfer transaction")
		}
		// 签名校验
		if common.BytesToHexString(topics[0]) != common.BytesToHexString(common.Keccak256([]byte("Transfer(address,address,uint256)"))) {
			return nil, errors.New("not transfer transaction")
		}
		result.From = common.EncodeCheck(contract.GetOwnerAddress())
		result.To = common.EncodeCheck(append([]byte{0x41}, topics[2][12:]...))
		decimalsBigInt, err := client.TRC20GetDecimals(contractAddress)
		if err != nil {
			return nil, err
		}
		result.Amount, _ = new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).SetBytes(log.Data)), big.NewFloat(math.Pow10(int(decimalsBigInt.Int64())))).Float64()
	}
	receiptResult := tx.GetReceipt().GetResult()
	result.Result = receiptResult == core.Transaction_Result_SUCCESS
	return &result, nil
}

// @title	从块中获取交易
// @param	Self		*Tron					模块实例
// @param	client		*client.GrpcClient		客户端
// @param	start		int64					开始高度
// @param	end			int64					结束高度
// @return	_			[]model.Transaction		交易信息
// @return	_			error					异常信息
func (Self *Tron) GetTronTransactionsFromBlocks(client *client.GrpcClient, start int64, end int64) ([]Transaction, error) {
	blocklist, err := client.GetBlockByLimitNext(start, end)
	if err != nil {
		return nil, err
	}
	blocks := blocklist.GetBlock()
	result := []Transaction{}
	for _, block := range blocks {
		transactions := block.GetTransactions()
		for _, transaction := range transactions {
			tx, err := Self.DecodeTronTransaction(client, common.Bytes2Hex(transaction.GetTxid()))
			if err != nil {
				continue
			}
			result = append(result, *tx)
		}
	}
	return result, nil
}

type GetTronTransactionsByAddressResponse struct {
	Data []GetTronTransactionsByAddressResponse_Trc20Transaction `json:"data"`
}

type GetTronTransactionsByAddressResponse_Trc20Transaction struct {
	TransactionId  string                                                          `json:"transaction_id"`
	BlockTimestamp int64                                                           `json:"block_timestamp"`
	From           string                                                          `json:"from"`
	To             string                                                          `json:"to"`
	Value          string                                                          `json:"value"`
	TokenInfo      GetTronTransactionsByAddressResponse_Trc20Transaction_TokenInfo `json:"token_info"`
}

type GetTronTransactionsByAddressResponse_Trc20Transaction_TokenInfo struct {
	Address  string `json:"address"`
	Decimals int64  `json:"decimals"`
}

// @title	根据地址获取交易
// @param	Self		*Tron					模块实例
// @param	address		string					地址
// @param	token		*string					币种
// @param	endTime		time.Time				结束时间
// @return	_			[]Transaction	交易信息
// @return	_			error					异常信息
func (Self *Tron) GetTronTransactionsByAddress(url string, address string, token *string, endTime time.Time) ([]Transaction, error) {
	var transactions []Transaction
	if token == nil {
		// 未实现
	} else {
		url := fmt.Sprintf("%s/v1/accounts/%s/transactions/trc20?only_confirmed=true&contract_address=%s&min_timestamp=%d",
			url, address, *token, endTime.UnixMilli(),
		)
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		request.Header.Add("accept", "application/json")
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		var res GetTronTransactionsByAddressResponse
		err = json.Unmarshal(body, &res)
		if err != nil {
			return nil, err
		}
		linq.From(res.Data).SelectT(func(item GetTronTransactionsByAddressResponse_Trc20Transaction) Transaction {
			amountBitInt, _ := new(big.Int).SetString(item.Value, 10)
			amount, _ := new(big.Float).Quo(new(big.Float).SetInt(amountBitInt), big.NewFloat(math.Pow10(int(item.TokenInfo.Decimals)))).Float64()
			return Transaction{
				Chain:     Chain_TRON,
				Hash:      item.TransactionId,
				TimeStamp: item.BlockTimestamp,
				Contract:  &item.TokenInfo.Address,
				From:      item.From,
				To:        item.To,
				Amount:    amount,
				Result:    true,
			}
		}).ToSlice(&transactions)
	}
	return transactions, nil
}