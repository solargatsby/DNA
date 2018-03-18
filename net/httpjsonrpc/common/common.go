package common

import (
	tx "DNA/core/transaction"
	. "DNA/errors"
	. "DNA/net/protocol"
	log "github.com/alecthomas/log4go"
)

var Node Noder

func VerifyAndSendTx(txn *tx.Transaction) ErrCode {
	// if transaction is verified unsucessfully then will not put it into transaction pool
	if errCode := Node.AppendTxnPool(txn); errCode != ErrNoError {
		log.Warn("Can NOT add the transaction to TxnPool")
		log.Info("[httpjsonrpc] VerifyTransaction failed when AppendTxnPool.")
		return errCode
	}
	if err := Node.Xmit(txn); err != nil {
		log.Error("Xmit Tx Error:Xmit transaction failed.", err)
		return ErrXmitFail
	}
	return ErrNoError
}

