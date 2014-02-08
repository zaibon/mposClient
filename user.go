package mposClient

type User struct {
	HashRate  int    `json:"hashrate"`
	ShareRate string `sjson:"shareharerate"`
	UserName  string `json:"username"`
	Share     Shares `json:"shares"`
}
