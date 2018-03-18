package serialize

import (
	"DNA/core/asset"
	"DNA/core/transaction"
	"DNA/core/transaction/payload"
	"fmt"
	"reflect"
	"encoding/json"
)

func MarshalPayload(pload transaction.Payload)([]byte, error){
	var object interface{}
	switch p := pload.(type) {
	case *payload.BookKeeping:
		bookKeepingInfo := &BookKeepingInfo{
			Nonce: p.Nonce,
		}
		object = bookKeepingInfo
	case *payload.BookKeeper:
		pukKey, err := p.PubKey.EncodePoint(true)
		if err != nil {
			return nil, fmt.Errorf("PubKey.EncodePoint error:%s", err)
		}
		issuer, err := p.Issuer.EncodePoint(true)
		if err != nil {
			return nil, fmt.Errorf("Issuer.EncodePoint error:%s", err)
		}
		bookKeeperInfo := &BookkeeperInfo{
			PubKey:ByteArrayToString(pukKey),
			Issuer:ByteArrayToString(issuer),
			Action:byte(p.Action),
		}
		object = bookKeeperInfo
	case *payload.IssueAsset:
	case *payload.TransferAsset:
	case *payload.RegisterAsset:
		issuer, err := p.Issuer.EncodePoint(true)
		if err != nil {
			return nil, fmt.Errorf("Issuer.EncodePoint error:%s", err)
		}
		registerAssetInfo := &RegisterAssetInfo{
			Asset:p.Asset,
			Amount:int64(p.Amount),
			Issuer:ByteArrayToString(issuer),
			Controller:Uint160ToString(p.Controller),
		}
		object = registerAssetInfo
	case *payload.Record:
		recordInfo := &RecordInfo{
			RecordType :p.RecordType,
			RecordData:ByteArrayToString(p.RecordData),
		}
		object = recordInfo
	case *payload.DataFile:
		issuer ,err := p.Issuer.EncodePoint(true)
		if err != nil {
			return nil ,fmt.Errorf("Issuer.EncodePoint error:%s", err)
		}
		dataFileInfo := &DataFileInfo{
			IPFSPath:p.IPFSPath,
			Filename:p.Filename,
			Note:p.Note,
			Issuer:ByteArrayToString(issuer),
		}
		object = dataFileInfo
	default:
		return nil, fmt.Errorf("unknow payload:%v", reflect.TypeOf(pload))
	}
	if object == nil {
		return nil, nil
	}
	return json.Marshal(object)
}

type BookKeepingInfo struct {
	Nonce  uint64
	Issuer string
}

type RegisterAssetInfo struct {
	Asset      *asset.Asset
	Amount     int64
	Issuer     string
	Controller string
}

type RecordInfo struct {
	RecordType string
	RecordData string
}

type BookkeeperInfo struct {
	PubKey     string
	Action     byte
	Issuer     string
	Controller string
}

type DataFileInfo struct {
	IPFSPath string
	Filename string
	Note     string
	Issuer   string
}
