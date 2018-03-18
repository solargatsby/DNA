package handler

import (
	. "DNA/common"
	"DNA/core/transaction"
	"DNA/net/httpjsonrpc/common"
	"DNA/net/httpjsonrpc/serialize"
	"encoding/json"
	log "github.com/alecthomas/log4go"
)

type GetUnspentOutputReq struct {
	AssetId     string
	ProgramHash string
}

type GetUnspentOutputResp struct {
	UnspentOutput []*serialize.TxOutputInfo
}

func GetUnspentOutput(req *common.JsonRpcRequest, resp *common.JsonRpcResponse) {
	unspentReq := &GetUnspentOutputReq{}
	err := json.Unmarshal(req.Params, unspentReq)
	if err != nil {
		log.Info("GetUnspentOutput qid:%v json.Unmarshal GetUnspentOutputReq:%s error:%s", req.Qid, req.Params, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	assetId, err := serialize.StringToUint256(unspentReq.AssetId)
	if err != nil {
		log.Info("GetUnspentOutput qid:%v StringToUint256:%s error:%s", req.Qid, unspentReq.AssetId, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	programHash, err := serialize.StringToUint160(unspentReq.ProgramHash)
	if err != nil {
		log.Info("GetUnspentOutput qid:%v StringToUint160:%s error:%s", req.Qid, unspentReq.ProgramHash, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	txOutputs, err := getUnspentOutput(assetId, programHash)
	if err != nil {
		log.Info("GetUnspentOutput qid:%v getUnspentOutput assetId:%s programHash:%s error:%s", req.Qid, unspentReq.AssetId, unspentReq.ProgramHash, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_ERR
		return
	}

	txOutputInfos := make([]*serialize.TxOutputInfo, 0, len(txOutputs))
	for _, output := range txOutputs{
		outputInfo := &serialize.TxOutputInfo{
			AssetID:unspentReq.AssetId,
			Value:int64(output.Value),
			ProgramHash:unspentReq.ProgramHash,
		}
		txOutputInfos = append(txOutputInfos, outputInfo)
	}

	resp.Result = txOutputInfos
}

func getUnspentOutput(assetId Uint256, programHash Uint160) ([]*transaction.TxOutput, error) {
return nil, nil
}
