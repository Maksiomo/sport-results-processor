package blockchain

import (
	"context"
	"math/big"
	"sport-results-pocessor/internal/common/adapter/logger"
	"sport-results-pocessor/pkg/gen/blockchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type BlockchainRegistryClient struct {
	EthClient        *ethclient.Client
	Sport            *blockchain.SportRegistry
	Person           *blockchain.PersonRegistry
	Team             *blockchain.TeamRegistry
	TeamPerson       *blockchain.TeamPersonRegistry
	Competition      *blockchain.CompetitionRegistry
	Match            *blockchain.MatchRegistry
	MatchParticipant *blockchain.MatchParticipantRegistry
	Prize            *blockchain.PrizeRegistry
	CompetitionTeams *blockchain.CompetitionTeamsRegistry
	TeamAchievements *blockchain.TeamAchievementsRegistry
}

func NewBlockchainRegistryClient(ctx context.Context, cfg *Config, logg logger.Logger) (*BlockchainRegistryClient, error) {
	log := logg.WithComponent(ctx, "blockchain-client")

	ec, err := ethclient.DialContext(ctx, cfg.RpcURL)
	if err != nil {
		return nil, err
	}
	c := &BlockchainRegistryClient{EthClient: ec}

	// приватный ключ первого аккаунта Ganache (из логов)
	key, err := crypto.HexToECDSA(cfg.AccountPrivKey)
	if err != nil {
		log.Error("can not parse priv key", zap.Error(err))
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(cfg.NetworkID))
	if err != nil {
		log.Error("can not make auth", zap.Error(err))
		return nil, err
	}

	// Initialize each registry

	if cfg.Init {
		// deploy contracts from scratch
		_, _, c.Sport, err = blockchain.DeploySportRegistry(auth, c.EthClient)
		if err != nil {
			log.Error("can not make sport registry", zap.Error(err))
			return nil, err
		}
		_, _, c.Person, err = blockchain.DeployPersonRegistry(auth, c.EthClient)
		if err != nil {
			log.Error("can not make person registry", zap.Error(err))
			return nil, err
		}
		_, _, c.Team, err = blockchain.DeployTeamRegistry(auth, c.EthClient)
		if err != nil {
			log.Error("can not make team registry", zap.Error(err))
			return nil, err
		}
		_, _, c.TeamPerson, err = blockchain.DeployTeamPersonRegistry(auth, c.EthClient)
		if err != nil {
			log.Error("can not make team person registry", zap.Error(err))
			return nil, err
		}
		_, _, c.Competition, err = blockchain.DeployCompetitionRegistry(auth, c.EthClient)
		if err != nil {
			log.Error("can not make competition registry", zap.Error(err))
			return nil, err
		}
		_, _, c.Match, err = blockchain.DeployMatchRegistry(auth, c.EthClient)
		if err != nil {
			log.Error("can not make match registry", zap.Error(err))
			return nil, err
		}
		_, _, c.MatchParticipant, err = blockchain.DeployMatchParticipantRegistry(auth, c.EthClient)
		if err != nil {
			log.Error("can not make match participants registry", zap.Error(err))
			return nil, err
		}
		_, _, c.Prize, err = blockchain.DeployPrizeRegistry(auth, c.EthClient)
		if err != nil {
			log.Error("can not make prize registry", zap.Error(err))
			return nil, err
		}
		_, _, c.CompetitionTeams, err = blockchain.DeployCompetitionTeamsRegistry(auth, c.EthClient)
		if err != nil {
			log.Error("can not make competition teams registry", zap.Error(err))
			return nil, err
		}
		_, _, c.TeamAchievements, err = blockchain.DeployTeamAchievementsRegistry(auth, c.EthClient)
		if err != nil {
			log.Error("can not make team achievements registry", zap.Error(err))
			return nil, err
		}
	} else {
		// connect to existing regisrties
		if c.Sport, err = blockchain.NewSportRegistry(common.HexToAddress(cfg.SportAddr), ec); err != nil {
			log.Error("can not connect to sport registry", zap.Error(err))
			return nil, err
		}
		if c.Person, err = blockchain.NewPersonRegistry(common.HexToAddress(cfg.PersonAddr), ec); err != nil {
			log.Error("can not connect to person registry", zap.Error(err))
			return nil, err
		}
		if c.Team, err = blockchain.NewTeamRegistry(common.HexToAddress(cfg.TeamAddr), ec); err != nil {
			log.Error("can not connect to team registry", zap.Error(err))
			return nil, err
		}
		if c.TeamPerson, err = blockchain.NewTeamPersonRegistry(common.HexToAddress(cfg.TeamPersonAddr), ec); err != nil {
			log.Error("can not connect to team person registry", zap.Error(err))
			return nil, err
		}
		if c.Competition, err = blockchain.NewCompetitionRegistry(common.HexToAddress(cfg.CompetitionAddr), ec); err != nil {
			log.Error("can not connect to competition registry", zap.Error(err))
			return nil, err
		}
		if c.Match, err = blockchain.NewMatchRegistry(common.HexToAddress(cfg.MatchAddr), ec); err != nil {
			log.Error("can not connect to match registry", zap.Error(err))
			return nil, err
		}
		if c.MatchParticipant, err = blockchain.NewMatchParticipantRegistry(common.HexToAddress(cfg.MatchParticipantsAddr), ec); err != nil {
			log.Error("can not connect to match participant registry", zap.Error(err))
			return nil, err
		}
		if c.Prize, err = blockchain.NewPrizeRegistry(common.HexToAddress(cfg.PrizeAddr), ec); err != nil {
			log.Error("can not connect to prize registry", zap.Error(err))
			return nil, err
		}
		if c.CompetitionTeams, err = blockchain.NewCompetitionTeamsRegistry(common.HexToAddress(cfg.CompetitionTeamAddr), ec); err != nil {
			log.Error("can not connect to competition teams registry", zap.Error(err))
			return nil, err
		}
		if c.TeamAchievements, err = blockchain.NewTeamAchievementsRegistry(common.HexToAddress(cfg.TeamAchievementsAddr), ec); err != nil {
			log.Error("can not connect to team achievements registry", zap.Error(err))
			return nil, err
		}
	}

	return c, nil
}
