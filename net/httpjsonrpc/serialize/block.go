package serialize

import (
	"DNA/core/ledger"
	"fmt"
)

func HeaderToHeaderInfo(header *ledger.Header)(*BlockHeaderInfo){
	blockHeaderInfo := &BlockHeaderInfo{
		Version:header.Blockdata.Version,
		PrevBlockHash:Uint256ToString(header.Blockdata.PrevBlockHash),
		TransactionsRoot:Uint256ToString(header.Blockdata.TransactionsRoot),
		Timestamp:header.Blockdata.Timestamp,
		Height:header.Blockdata.Height,
		ConsensusData:header.Blockdata.ConsensusData,
		NextBookKeeper:Uint160ToString(header.Blockdata.NextBookKeeper),
		Program:&ProgramInfo{
			Code:ByteArrayToString(header.Blockdata.Program.Code),
			Parameter:ByteArrayToString(header.Blockdata.Program.Parameter),
		},
	}
	return blockHeaderInfo
}

func BlockToBlockInfo(block *ledger.Block)(*BlockInfo, error){
	header := &ledger.Header{
		Blockdata:block.Blockdata,
	}
	blockHeaderInfo := HeaderToHeaderInfo(header)

	txInfos := make([]*TransactionInfo, 0, len(block.Transactions))
	for _, tx := range block.Transactions{
		txInfo, err := TransactionToTransactionInfo(tx)
		if err != nil {
			return nil, fmt.Errorf("TransactionToTransactionInfo TxHash%x error:%s", tx.Hash(), err)
		}
		txInfos = append(txInfos, txInfo)
	}

	blocInfo := &BlockInfo{
		BlockData:blockHeaderInfo,
		Transactions:txInfos,
	}
	return blocInfo, nil
}

type BlockHeaderInfo struct {
	Version          uint32
	PrevBlockHash    string
	TransactionsRoot string
	Timestamp        uint32
	Height           uint32
	ConsensusData    uint64
	NextBookKeeper   string
	Program          *ProgramInfo
}

type BlockInfo struct {
	Hash         string
	BlockData    *BlockHeaderInfo
	Transactions []*TransactionInfo
}