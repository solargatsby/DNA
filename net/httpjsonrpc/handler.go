package httpjsonrpc

import "DNA/net/httpjsonrpc/handler"

func init(){
	DefJsonRpcSvr.RegHandler("getversion", handler.GetVersion)
	DefJsonRpcSvr.RegHandler("getcurrentblockheight", handler.GetCurrentBlockHeight)
	DefJsonRpcSvr.RegHandler("getcurrentblockhash", handler.GetCurrentBlockHash)
	DefJsonRpcSvr.RegHandler("getblockhash", handler.GetBlockHash)
	DefJsonRpcSvr.RegHandler("getheaderbyheight", handler.GetHeaderByHeight)
	DefJsonRpcSvr.RegHandler("getheaderbyhash", handler.GetHeaderByHash)
	DefJsonRpcSvr.RegHandler("getblockbyheight", handler.GetBlockByHeight)
	DefJsonRpcSvr.RegHandler("getblockbyhash", handler.GetBlockByHash)
	DefJsonRpcSvr.RegHandler("sendrawtransaction", handler.SendRawTransaction)
	DefJsonRpcSvr.RegHandler("getrawtransaction", handler.GetRawTransaction)
}
