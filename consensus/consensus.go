package consensus

import (
	"fmt"
	cl "DNA/account"
	"DNA/common/config"
	"DNA/common/log"
	"DNA/consensus/dbft"
	"DNA/consensus/solo"
	"DNA/net"
	"strings"
	"time"
)

type ConsensusService interface {
	Start() error
	Halt() error
}

func Log(message string) {
	logMsg := fmt.Sprintf("[%s] %s", time.Now().Format("02/01/2006 15:04:05"), message)
	fmt.Println(logMsg)
	log.Info(logMsg)
}

const (
	CONSENSUS_DBFT = "dbft"
	CONSENSUS_SOLO  = "solo"
)

var ConsensusMgr = NewConsensuManager()

type ConsensusManager struct {}

func NewConsensuManager() *ConsensusManager {
	return &ConsensusManager{}
}

func (this *ConsensusManager)NewConsensusService(client cl.Client , localNet net.Neter)ConsensusService{
	consensusType := strings.ToLower(config.Parameters.ConsensusType)
	if consensusType == "" {
		consensusType = CONSENSUS_DBFT
	}
	var consensus ConsensusService
	switch consensusType {
	case CONSENSUS_DBFT:
		consensus = dbft.NewDbftService(client, "dbft", localNet)
	case CONSENSUS_SOLO:
		consensus = solo.NewSoloService(client, localNet)
	}
	log.Infof("ConsensusType:%s", consensusType)
	return consensus
}
