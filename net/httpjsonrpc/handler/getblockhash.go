package handler

import (
	"DNA/net/httpjsonrpc/common"
	"encoding/json"
	log "github.com/alecthomas/log4go"
	"DNA/core/ledger"
	"DNA/net/httpjsonrpc/serialize"
)

type GetBlockHashReq struct {
	Height uint32	`json:"height"`
}

type GetBlockHashResp struct {
	BlockHash string 	`json:"block_hash"`
}

func GetBlockHash(req *common.JsonRpcRequest, resp *common.JsonRpcResponse){
	blockReq := &GetBlockHashReq{}
	err := json.Unmarshal(req.Params, blockReq)
	if err != nil {
		log.Info("GetBlockHash qid:%v json.Unmarshal GetBlockHashReq:%s error:%s",req.Qid,  req.Params, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	blockHash, err := ledger.DefaultLedger.Store.GetBlockHash(blockReq.Height)
	if err != nil {
		log.Error("GetBlockHash Qid:%v Height:%d error %s", req.Qid, blockReq.Height, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_IO_ERR
		return
	}
	resp.Result = &GetBlockHashResp{
		BlockHash:serialize.Uint256ToString(blockHash),
	}
}
