package handler

import (
	"DNA/net/httpjsonrpc/common"
	. "DNA/net/httpjsonrpc/serialize"
	"encoding/json"
	log "github.com/alecthomas/log4go"
	"DNA/core/ledger"
)

type GetBlockByHashReq struct {
	BlackHash string `json:"block_hash"`
}

type GetBlockByHeightReq struct {
	Height uint32
}

type GetBlockResp struct {
	Block *BlockInfo `json:"block"`
}

func GetBlockByHash(req *common.JsonRpcRequest, resp *common.JsonRpcResponse) {
	blockReq := &GetBlockByHashReq{}
	err := json.Unmarshal(req.Params, blockReq)
	if err != nil {
		log.Info("GetBlockByHash qid:%v json.Unmarshal GetBlockByHashReq:%s error:%s", req.Qid, req.Params, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	blockHash, err := StringToUint256(blockReq.BlackHash)
	if err != nil {
		log.Info("GetBlockByHash qid:%v StringToUint256:%s error:%s", req.Qid, blockReq.BlackHash, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	block, err := ledger.DefaultLedger.Store.GetBlock(blockHash)
	if err != nil {
		log.Info("GetBlockByHash qid:%v Store.GetBlock:%x error:%s", req.Qid, blockHash, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_IO_ERR
		return
	}
	blockInfo, err := BlockToBlockInfo(block)
	if err != nil {
		log.Info("GetBlockByHash qid:%v block hash:%x BlockToBlockInfo error:%s", req.Qid, blockHash, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_ERR
		return
	}
	resp.Result = &GetBlockResp{
		Block:blockInfo,
	}
}


func GetBlockByHeight(req *common.JsonRpcRequest, resp *common.JsonRpcResponse){
	blockReq := &GetBlockByHeightReq{}
	err := json.Unmarshal(req.Params, blockReq)
	if err != nil {
		log.Info("GetBlockByHeight qid:%v json.Unmarshal GetBlockByHeightReq:%s error:%s", req.Qid, req.Params, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	blockHash, err := ledger.DefaultLedger.Store.GetBlockHash(blockReq.Height)
	if err != nil {
		log.Info("GetBlockByHeight qid:%v Store.GetBlockHash:%v error:%s", req.Qid, blockReq.Height, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_IO_ERR
		return
	}
	block, err := ledger.DefaultLedger.Store.GetBlock(blockHash)
	if err != nil {
		log.Info("GetBlockByHeight qid:%v Store.GetBlock:%x error:%s", req.Qid, blockHash, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_IO_ERR
		return
	}
	blockInfo, err := BlockToBlockInfo(block)
	if err != nil {
		log.Info("GetBlockByHeight qid:%v block hash:%x BlockToBlockInfo error:%s", req.Qid, blockHash, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_ERR
		return
	}
	resp.Result = &GetBlockResp{
		Block:blockInfo,
	}
}
