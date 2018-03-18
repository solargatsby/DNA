package serialize

import (
	. "DNA/core/transaction"
	"encoding/json"
	"DNA/core/transaction"
	"fmt"
)

func TransactionToTransactionInfo(tx *transaction.Transaction)(*TransactionInfo, error){
	txInfo := &TransactionInfo{}
	txInfo.TxType = byte(tx.TxType)
	txInfo.PayloadVersion = tx.PayloadVersion

	pload, err := MarshalPayload(tx.Payload)
	if err != nil {
		return nil, fmt.Errorf("MarshalPayload error:%s", err)
	}
	txInfo.Payload = pload

	attrInfos := make([]*TxAttributeInfo, 0, len(tx.Attributes))
	for _, attr := range tx.Attributes{
		attrInfo := &TxAttributeInfo{
			Usage:byte(attr.Usage),
			Data:ByteArrayToString(attr.Data),
		}
		attrInfos = append(attrInfos, attrInfo)
	}
	txInfo.Attributes = attrInfos

	utxoInfos := make([]*UTXOTxInputInfo, 0, len(tx.UTXOInputs))
	for _, utxo := range tx.UTXOInputs{
		utxoInfo := &UTXOTxInputInfo{
			ReferTxID : Uint256ToString(utxo.ReferTxID),
			ReferTxOutputIndex:utxo.ReferTxOutputIndex,
		}
		utxoInfos = append(utxoInfos, utxoInfo)
	}
	txInfo.UTXOInputs = utxoInfos

	outputInfos := make([]*TxOutputInfo, 0, len(tx.Outputs))
	for _, output := range tx.Outputs{
		outputInfo := &TxOutputInfo{
			AssetID:Uint256ToString(output.AssetID),
			Value:int64(output.Value),
			ProgramHash:Uint160ToString(output.ProgramHash),
		}
		outputInfos = append(outputInfos, outputInfo)
	}
	txInfo.Outputs = outputInfos

	programInfos := make([]*ProgramInfo, 0, len(tx.Programs))
	for _, program := range tx.Programs{
		programInfo := &ProgramInfo{
			Code:ByteArrayToString(program.Code),
			Parameter:ByteArrayToString(program.Parameter),
		}
		programInfos = append(programInfos, programInfo)
	}
	txInfo.Programs = programInfos

	return txInfo, nil
}

type TxAttributeInfo struct {
	Usage byte
	Data  string
}

type UTXOTxInputInfo struct {
	ReferTxID          string
	ReferTxOutputIndex uint16
}

type TxOutputInfo struct {
	AssetID     string
	Value       int64
	ProgramHash string
}

type ProgramInfo struct {
	Code      string
	Parameter string
}

type TransactionInfo struct {
	TxType            byte
	PayloadVersion    byte
	Payload           json.RawMessage
	Attributes        []*TxAttributeInfo
	UTXOInputs        []*UTXOTxInputInfo
	Outputs           []*TxOutputInfo
	Programs          []*ProgramInfo
}
