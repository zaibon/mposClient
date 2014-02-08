package mposClient

type Shares struct {
	DonatePercent float64 `json:"donate_percent"`
	Id            uint64  `json:"id"`
	Invalid       uint    `json:"invalid"`
	IsAnonymous   int     `json:"is_anonymous"`
	Username      string  `json:"username"`
	Valid         uint64  `json:"valid"`
}
