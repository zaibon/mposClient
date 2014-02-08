package mposClient

type PoolPublicInfo struct {
	HashRate      int    `json:"hashrate"`
	LastBlock     int    `json:"last_block"`
	NetHashRate   int    `json:"network_hashrate"`
	PoolName      string `json:"pool_name"`
	ShareCurRound int    `json:"shares_this_round"`
	WorkersNbr    int    `json:"workers"`
}
