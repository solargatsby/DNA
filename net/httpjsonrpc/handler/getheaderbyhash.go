package handler

import (
	"DNA/net/httpjsonrpc/serialize"
	"DNA/net/httpjsonrpc/common"
	log "github.com/alecthomas/log4go"
	"encoding/json"
	"DNA/core/ledger"
)

type GetHeaderByHashReq struct {
	BlockHash string
}

type GetHeaderByHeightReq struct {
	Height uint32
}

type GetHeaderResp struct {
	Header *serialize.BlockHeaderInfo
}

func GetHeaderByHash(req *common.JsonRpcRequest, resp *common.JsonRpcResponse){
	headerReq := &GetHeaderByHashReq{}
	err := json.Unmarshal(req.Params, headerReq)
	if err != nil {
		log.Info("GetHeadByHash qid:%v json.Unmarshal GetHeadByHashReq:%s error:%s", req.Qid, req.Params, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}

	hash, err := serialize.StringToUint256(headerReq.BlockHash)
	if err != nil {
		log.Info("GetHeadByHash qid:%v StringToUint256:%s error:%s", req.Qid, headerReq.BlockHash, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}

	header, err := ledger.DefaultLedger.Store.GetHeader(hash)
	if err != nil {
		log.Info("GetHeadByHash qid:%v Store.GetHeader hash:%x error:%s", req.Qid, hash, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_IO_ERR
		return
	}

	resp.Result = &GetHeaderResp{
		Header:serialize.HeaderToHeaderInfo(header),
	}
}

func GetHeaderByHeight(req *common.JsonRpcRequest, resp *common.JsonRpcResponse){
	headerReq := &GetHeaderByHeightReq{}
	err := json.Unmarshal(req.Params, headerReq)
	if err != nil {
		log.Info("GetHeaderByHeight qid:%v json.Unmarshal GetHeaderByHeightReq:%s error:%s", req.Qid, req.Params, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}

	hash, err := ledger.DefaultLedger.Store.GetBlockHash(headerReq.Height)
	if err != nil {
		log.Info("GetHeaderByHeight qid:%v Store.GetBlockHash height:%v error:%s", req.Qid, headerReq.Height, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_IO_ERR
		return
	}

	header, err := ledger.DefaultLedger.Store.GetHeader(hash)
	if err != nil {
		log.Info("GetHeadByHash qid:%v Store.GetHeader hash:%x error:%s", req.Qid, hash, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_IO_ERR
		return
	}

	resp.Result = &GetHeaderResp{
		Header:serialize.HeaderToHeaderInfo(header),
	}
}


