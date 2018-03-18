package common

import "encoding/json"

type JsonRpcRequest struct {
	JsonRpc string          `json:"jsonrpc"`
	Qid     string          `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type JsonRpcResponse struct {
	Method    string      `json:"method"`
	Qid       string      `json:"id"`
	ErrorCode int         `json:"error_code"`
	ErrorInfo string      `json:"error_info"`
	Result    interface{} `json:"result"`
}

