package handler

import (
	"DNA/core/ledger"
	"DNA/core/transaction"
	"DNA/net/httpjsonrpc/common"
	"DNA/net/httpjsonrpc/serialize"
	"bytes"
	"encoding/json"
	log "github.com/alecthomas/log4go"
)

type SendRawTransactionReq struct {
	Tx string
}

type SendRawTransactionResp struct {
	TxHash string
}

func SendRawTransaction(req *common.JsonRpcRequest, resp *common.JsonRpcResponse) {
	txReq := &SendRawTransactionReq{}
	err := json.Unmarshal(req.Params, txReq)
	if err != nil {
		log.Info("SendRawTransaction qid:%v json.Unmarshal SendRawTransactionReq:%s error:%s", req.Qid, req.Params, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	data, err := serialize.StringToByteArray(txReq.Tx)
	if err != nil {
		log.Info("SendRawTransaction qid:%v StringToByteArray error:%s", req.Qid, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	tx := &transaction.Transaction{}
	err = tx.Deserialize(bytes.NewReader(data))
	if err != nil {
		log.Info("SendRawTransaction qid:%v Transaction Deserialize error:%s", req.Qid, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	errCode := common.VerifyAndSendTx(tx)
	if errCode != common.RPCERR_OK {
		resp.ErrorCode = int(errCode)
		return
	}

	resp.Result = &SendRawTransactionResp{
		TxHash: serialize.Uint256ToString(tx.Hash()),
	}
}

type GetRawTransactionReq struct {
	TxHash string
}

type GetRawTransactionResp struct {
	Tx *serialize.TransactionInfo
}

func GetRawTransaction(req *common.JsonRpcRequest, resp *common.JsonRpcResponse) {
	txReq := &GetRawTransactionReq{}
	err := json.Unmarshal(req.Params, txReq)
	if err != nil {
		log.Info("GetRawTransaction qid:%v json.Unmarshal GetRawTransactionReq:%s error:%s", req.Qid, req.Params, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	txHash, err := serialize.StringToUint256(txReq.TxHash)
	if err != nil {
		log.Info("GetRawTransaction qid:%v Uint256ToString:%s error:%s", req.Qid, txReq.TxHash, err)
		resp.ErrorCode = common.RPCERR_INVALID_PARAMS
		return
	}
	tx, err := ledger.DefaultLedger.Store.GetTransaction(txHash)
	if err != nil {
		log.Info("GetRawTransaction qid:%v Store.GetTransaction txHash:%x error:%s", req.Qid, txHash, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_IO_ERR
		return
	}
	txInfo, err := serialize.TransactionToTransactionInfo(tx)
	if err != nil {
		log.Info("GetRawTransaction qid:%s txHash:%x TransactionToTransactionInfo error:%s", req.Qid, txHash, err)
		resp.ErrorCode = common.RPCERR_INTERNAL_ERR
		return
	}
	resp.Result = &GetRawTransactionResp{
		Tx: txInfo,
	}
}
