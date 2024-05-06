package api

import (
	"bytes"
	"encoding/json"
	"ethparser/internal/util"
	"io"
	"net/http"
)

const (
	ethGatewayEndpoint = "https://cloudflare-eth.com/v1/mainnet"
)

func GetEthBlockNumber() (int64, error) {
	respBytes, err := request("eth_blockNumber", nil)
	if err != nil {
		return 0, err
	}

	var resp EthBlockNumberResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return 0, err
	}

	return util.HexToInt64(resp.Result), nil
}

func GetEthBlockByNumber(num int64) (EthBlockResponse, error) {
	respBytes, err := request("eth_getBlockByNumber", []any{
		util.Int64ToHex(num),
		true})
	if err != nil {
		return EthBlockResponse{}, err
	}

	var resp EthBlockResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return EthBlockResponse{}, err
	}

	return resp, nil
}

func request(method string, params any) ([]byte, error) {
	var req = map[string]any{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      0,
	}
	reqBytes, _ := json.Marshal(req)

	var buf bytes.Buffer
	buf.Write(reqBytes)

	resp, err := http.Post(ethGatewayEndpoint, "application/json", &buf)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}
