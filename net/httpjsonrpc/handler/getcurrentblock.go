package handler

import (
	"DNA/core/ledger"
	"DNA/net/httpjsonrpc/common"
	"DNA/net/httpjsonrpc/serialize"
)

type GetCurrentBlockHeightResp struct {
	Height uint32
}

func GetCurrentBlockHeight(req *common.JsonRpcRequest, resp *common.JsonRpcResponse) {
	resp.Result = &GetCurrentBlockHeightResp{
		Height: ledger.DefaultLedger.Blockchain.BlockHeight,
	}
}

type GetCurrentBlockHashResp struct {
	BlockHash string
}

func GetCurrentBlockHash(req *common.JsonRpcRequest, resp *common.JsonRpcResponse) {
	resp.Result = &GetCurrentBlockHashResp{
		BlockHash: serialize.Uint256ToString(ledger.DefaultLedger.Store.GetCurrentBlockHash()),
	}

}
