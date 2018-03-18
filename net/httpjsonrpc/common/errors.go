package common

const (
	RPCERR_OK                  = 0
	RPCERR_HTTP_METHOD_INVALID = 1001
	RPCERR_INVALID_REQUEST     = 1002
	RPCERR_INVALID_PARAMS      = 1003
	RPCERR_UNSUPPORT_METHOD    = 1004
	RPCERR_INVALID_HASH        = 1005
	RPCERR_INVALID_BLOCK       = 1006
	RPCERR_INVALID_TRANSACTION = 1007
	RPCERR_INTERNAL_ERR        = 1008
	RPCERR_INTERNAL_IO_ERR     = 1009
)

var RPCErrorDesc = map[int]string{
	RPCERR_OK:                  "",
	RPCERR_HTTP_METHOD_INVALID: "",
	RPCERR_INVALID_REQUEST:     "",
}

func GetRPCErrorDesc(errorCode int) string {
	desc, ok := RPCErrorDesc[errorCode]
	if !ok {
		return ""
	}
	return desc
}
