package blockchain

type Config struct {
	RpcURL         string `yaml:"rpc_url"`
	AccountPrivKey string `yaml:"account_priv_key"`
	NetworkID      int64  `yaml:"network_id"`
	Init           bool   `yaml:"init"`

	SportAddr             string `yaml:"sport_addr"`
	PersonAddr            string `yaml:"person_addr"`
	TeamAddr              string `yaml:"team_addr"`
	TeamPersonAddr        string `yaml:"team__person_addr"`
	CompetitionAddr       string `yaml:"competition_addr"`
	MatchAddr             string `yaml:"match_addr"`
	MatchParticipantsAddr string `yaml:"match_participants_addr"`
	PrizeAddr             string `yaml:"prize_addr"`
	CompetitionTeamAddr   string `yaml:"competition_team_addr"`
	TeamAchievementsAddr  string `yaml:"team_achievements_addr"`
}
