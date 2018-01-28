package main

import (
	"DNA/account"
	"DNA/common/config"
	"DNA/common/log"
	"DNA/core/ledger"
	"DNA/core/store/ChainStore"
	"DNA/core/transaction"
	"DNA/crypto"
	"DNA/net"
	"DNA/net/httpjsonrpc"
	"DNA/net/httpnodeinfo"
	"DNA/net/protocol"
	"os"
	"runtime"
	"time"
	"sort"
	"encoding/hex"
	"DNA/consensus"
	"os/signal"
	"syscall"
)

const (
	DefaultMultiCoreNum = 4
)

func init() {
	log.Init(log.Path, log.Stdout)
	var coreNum int
	if config.Parameters.MultiCoreNum > DefaultMultiCoreNum {
		coreNum = int(config.Parameters.MultiCoreNum)
	} else {
		coreNum = DefaultMultiCoreNum
	}
	log.Debug("The Core number is ", coreNum)
	runtime.GOMAXPROCS(coreNum)
}

func main() {
	var acct *account.Account
	var blockChain *ledger.Blockchain
	var err error
	var noder protocol.Noder

	log.Trace("Node version: ", config.Version)

	log.Info("0. Loading the Ledger")
	ledger.DefaultLedger = new(ledger.Ledger)
	ledger.DefaultLedger.Store, err = ChainStore.NewLedgerStore()
	defer ledger.DefaultLedger.Store.Close()
	if err != nil {
		log.Fatal("open LedgerStore err:", err)
		os.Exit(1)
	}
	ledger.DefaultLedger.Store.InitLedgerStore(ledger.DefaultLedger)
	transaction.TxStore = ledger.DefaultLedger.Store
	crypto.SetAlg(config.Parameters.EncryptAlg)

	log.Info("1. Open the account")
	client := account.GetClient()
	if client == nil {
		log.Fatal("Can't get local account.")
		goto ERROR
	}
	acct, err = client.GetDefaultAccount()
	if err != nil {
		log.Fatal(err)
		goto ERROR
	}
	log.Debug("The Node's PublicKey ", acct.PublicKey)
	ledger.StandbyBookKeepers = GetBookKeepers()

	log.Info("3. BlockChain init")
	blockChain, err = ledger.NewBlockchainWithGenesisBlock(ledger.StandbyBookKeepers)
	if err != nil {
		log.Fatal(err, "  BlockChain generate failed")
		goto ERROR
	}
	ledger.DefaultLedger.Blockchain = blockChain

	log.Info("4. Start the P2P networks")
	// Don't need two return value.
	noder = net.StartProtocol(acct.PublicKey)
	httpjsonrpc.RegistRpcNode(noder)

	noder.SyncNodeHeight()
	noder.WaitForPeersStart()
	noder.WaitForSyncBlkFinish()
	if protocol.SERVICENODENAME != config.Parameters.NodeType {
		log.Info("5. Start Consensus Services")
		consensusSrv := consensus.ConsensusMgr.NewConsensusService(client, noder)
		httpjsonrpc.RegistConsensusService(consensusSrv)
		go consensusSrv.Start()
		time.Sleep(5 * time.Second)
	}

	log.Info("--Start the RPC interface")
	go httpjsonrpc.StartRPCServer()
	go httpjsonrpc.StartLocalServer()
	if config.Parameters.HttpInfoStart {
		go httpnodeinfo.StartServer(noder)
	}

	go func() {
		ticker := time.NewTicker(config.DEFAULTGENBLOCKTIME * time.Second)
		for {
			select {
			case <-ticker.C:
				log.Trace("BlockHeight = ", ledger.DefaultLedger.Blockchain.BlockHeight)
				isNeedNewFile := log.CheckIfNeedNewFile()
				if isNeedNewFile == true {
					log.ClosePrintLog()
					log.Init(log.Path, os.Stdout)
				}
			}
		}
	}()

	func() {
		//等待退出信号
		exit := make(chan bool, 0)
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		go func() {
			for sig := range sc {
				log.Infof("Ontology received exit signal:%v.", sig.String())
				close(exit)
				break
			}
		}()
		<-exit
	}()

ERROR:
	os.Exit(1)
}


func GetBookKeepers() []*crypto.PubKey {
	var pubKeys = []*crypto.PubKey{}
	sort.Strings(config.Parameters.BookKeepers)
	for _, key := range config.Parameters.BookKeepers {
		pubKey := []byte(key)
		pubKey, err := hex.DecodeString(key)
		// TODO Convert the key string to byte
		k, err := crypto.DecodePoint(pubKey)
		if err != nil {
			log.Error("Incorrectly book keepers key")
			return nil
		}
		pubKeys = append(pubKeys, k)
	}

	return pubKeys
}