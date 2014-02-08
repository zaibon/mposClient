package mposClient

type PoolStatus struct {
	CurrentNetWorkBlock uint64  `json:"currentnetworkblock"`
	Efficency           float32 `json:"efficiency"`
	EstShare            float64 `json:"estshares"`
	EstTime             float64 `json:"esttime"`
	HashRate            uint64  `json:"hashrate"`
	LastBlock           uint    `json:"lastblock"`
	NetHashRate         uint64  `json:"nethashrate"`
	NetDiff             float64 `json:"networkdiff"`
	NextNetWorkBlock    uint64  `json:"nextnetworkblock"`
	PoolName            string  `json:"pool_name"`
	TimeSinceLast       uint    `json:"timesincelast"`
	WorkersNbr          uint    `json:"workers"`
}
