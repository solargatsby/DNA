package handler

import (
	"DNA/net/httpjsonrpc/common"
	"DNA/common/config"
)

type GetVersionResp struct {
	Version string
}

func GetVersion(req *common.JsonRpcRequest, resp *common.JsonRpcResponse){
	resp.Result = &GetVersionResp{
		Version:config.Version,
	}
}