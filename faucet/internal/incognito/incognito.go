package incognito

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hoangnguyen-1312/faucet/logger"
	"io/ioutil"
	"net/http"
)

const (
	createAndSendTransaction                   = "createandsendtransaction"
	getBalanceByPrivateKey                     = "getbalancebyprivatekey"
	getTransactionByHash					   = "gettransactionbyhash"
)

type Client struct {
	host  string
	port  string
	ws    string
	mode  string
	https string
}

type JsonRequest struct {
	Jsonrpc string      `json:"Jsonrpc"`
	Method  string      `json:"Method"`
	Params  interface{} `json:"Params"`
	Id      interface{} `json:"Id"`
}

type RPCError struct {
	Code       int    `json:"Code,omitempty"`
	Message    string `json:"Message,omitempty"`
	StackTrace string `json:"StackTrace"`
	err error `json:"Err"`
}

func (e RPCError) Error() string {
	return fmt.Sprintf("%d: %+v %+v", e.Code, e.err, e.StackTrace)
}

type JsonResponse struct {
	Id      *interface{}    `json:"Id"`
	Result  json.RawMessage `json:"Result"`
	Error   *RPCError       `json:"Error"`
	Params  interface{}     `json:"Params"`
	Method  string          `json:"Method"`
	Jsonrpc string          `json:"Jsonrpc"`
}

func NewClient(host string, port string, ws string, mode string, https string) *Client {
	return &Client{host: host, port: port, ws: ws, mode: mode, https: https}
}

type CreateTransactionResult struct {
	Base58CheckData string
	TxID            string
	ShardID         byte
}

type TransactionDetail struct {
    BlockHash string
    BlockHeight	uint64
    CustomTokenData string
    Fee uint64
    Hash string
    Image string
    Index uint64
    Info string
    InputCoinPubKey string
    IsInBlock bool
    IsInMempool bool
    IsPrivacy bool
    LockTime string
    Metadata string
    PrivacyCustomTokenData string
    PrivacyCustomTokenFee uint64
    PrivacyCustomTokenID string
    PrivacyCustomTokenIsPrivacy bool
    PrivacyCustomTokenName string
    PrivacyCustomTokenProofDetail interface{}
    PrivacyCustomTokenSymbol string
    Proof string
    ProofDetail interface{}
    ShardID uint64
    Sig string
    SigPubKey string
    TxSize uint64
    Type string
    Version uint64
}


func makeRPCRequest(client *Client, method string, params ...interface{}) (*JsonResponse, error) {
	request := JsonRequest{
		Jsonrpc: "1.0",
		Method:  method,
		Params:  params,
		Id:      "1",
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	resp := &http.Response{}
	if client.mode == "release" || client.mode == "debug" {
		logger.Log.Info().
			Str("Https", client.https).
			Str("Mode", client.mode).
			Str("Method", method).
			Interface("Params", params).
			Msg("Make RPC Request")
		resp, err = http.Post(client.https, "application/json", bytes.NewBuffer(requestBytes))
	} else {
		logger.Log.Info().
			Str("Host", client.host).
			Str("Port", client.port).
			Str("Mode", client.mode).
			Msg("Make RPC Request")
		resp, err = http.Post(client.host+":"+client.port, "application/json", bytes.NewBuffer(requestBytes))
	}
	if err != nil {
		return nil, err
	}
	body := resp.Body
	defer body.Close()
	responseBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	response := JsonResponse{}
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *Client) CreateAndSendTransaction(params []interface{}) (*CreateTransactionResult, error) {
	result := &CreateTransactionResult{}
	res, rpcError := makeRPCRequest(client, createAndSendTransaction, params...)
	if rpcError != nil {
		return result, rpcError
	}
	err := json.Unmarshal(res.Result, &result)
	if err != nil {
		return result, err
	}
	if res.Error != nil {
		return result, res.Error
	}
	return result, nil
}

func (client *Client) GetBalanceByPrivateKey(params string) (uint64, error) {
	var result interface{}
	res, rpcError := makeRPCRequest(client, getBalanceByPrivateKey, params)
	if rpcError != nil {
		return 0, rpcError
	}
	err := json.Unmarshal(res.Result, &result)
	if err != nil {
		return 0, err
	}
	return uint64(result.(float64)), nil
}

func (client *Client) GetTransactionByHash(params string) (*TransactionDetail, error) {
	result := &TransactionDetail{}
	res, rpcError := makeRPCRequest(client, getTransactionByHash, params)
	if rpcError != nil {
		return result, rpcError
	}
	err := json.Unmarshal(res.Result, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
