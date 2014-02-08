package mposClient

type Block struct {
	AccountId     int64   `json:"account_id"`
	Accounted     int     `json:"accounted"`
	Amount        float64 `json:"amount"`
	BlockHash     string  `json:"blockhash"`
	Confirmations int64   `json:"confirmations"`
	Difficulty    float64 `json:"difficulty"`
	EstShares     float64 `json:"estshares"`
	Finder        string  `json:"finder"`
	Height        uint64  `json:"height"`
	Id            uint64  `json:"id"`
	IsAnonymous   int     `json:"is_anonymous"`
	ShareId       uint64  `json:"share_id"`
	Shares        float64 `json:"shares"`
	Time          int64   `json:"time"`
	WorkerName    string  `json:"worker_name"`
}

func (b *Block) Ratio() float64 {
	return (b.Shares / b.EstShares) * 100.0
}
