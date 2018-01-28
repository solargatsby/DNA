package httpjsonrpc

import (
	. "DNA/common"
	"DNA/core/asset"
	. "DNA/core/transaction"
	"DNA/core/transaction/payload"
)

type PayloadInfo interface{}

//implement PayloadInfo define BookKeepingInfo
type BookKeepingInfo struct {
	Nonce  uint64
	Issuer IssuerInfo
}

//implement PayloadInfo define IssueAssetInfo
type IssueAssetInfo struct {
}

type IssuerInfo struct {
	X, Y string
}

//implement PayloadInfo define RegisterAssetInfo
type RegisterAssetInfo struct {
	Asset      *asset.Asset
	Amount     Fixed64
	Issuer     IssuerInfo
	Controller string
}

//implement PayloadInfo define TransferAssetInfo
type TransferAssetInfo struct {
}

type RecordInfo struct {
	RecordType string
	RecordData string
}

type BookkeeperInfo struct {
	PubKey     string
	Action     string
	Issuer     IssuerInfo
	Controller string
}

type DataFileInfo struct {
	IPFSPath string
	Filename string
	Note     string
	Issuer   IssuerInfo
}

func TransPayloadToHex(p Payload) PayloadInfo {
	switch object := p.(type) {
	case *payload.BookKeeping:
		obj := new(BookKeepingInfo)
		obj.Nonce = object.Nonce
		return obj
	case *payload.BookKeeper:
		obj := new(BookkeeperInfo)
		encodedPubKey, _ := object.PubKey.EncodePoint(true)
		obj.PubKey = ToHexString(encodedPubKey)
		if object.Action == payload.BookKeeperAction_ADD {
			obj.Action = "add"
		} else if object.Action == payload.BookKeeperAction_SUB {
			obj.Action = "sub"
		} else {
			obj.Action = "nil"
		}
		obj.Issuer.X = object.Issuer.X.String()
		obj.Issuer.Y = object.Issuer.Y.String()

		return obj
	case *payload.IssueAsset:
	case *payload.TransferAsset:
	case *payload.RegisterAsset:
		obj := new(RegisterAssetInfo)
		obj.Asset = object.Asset
		obj.Amount = object.Amount
		obj.Issuer.X = object.Issuer.X.String()
		obj.Issuer.Y = object.Issuer.Y.String()
		obj.Controller = ToHexString(object.Controller.ToArray())
		return obj
	case *payload.Record:
		obj := new(RecordInfo)
		obj.RecordType = object.RecordType
		obj.RecordData = ToHexString(object.RecordData)
		return obj
	case *payload.DataFile:
		obj := new(DataFileInfo)
		obj.IPFSPath = object.IPFSPath
		obj.Filename = object.Filename
		obj.Note = object.Note
		obj.Issuer.X = object.Issuer.X.String()
		obj.Issuer.Y = object.Issuer.Y.String()
		return obj
	}
	return nil
}
