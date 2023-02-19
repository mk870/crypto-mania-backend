package models

type Coin struct {
	MyModel
	Id                 int    `json:"id" gorm:"primary_key"`
	UserId             int    `json:"userId"`
	Name               string `json:"name"`
	MktCap             string `json:"mktcap"`
	Rank               string `json:"rank"`
	AllTimeHighPrice   string `json:"alltimehighprice"`
	CoinsInCirculation string `json:"coinsincirculation"`
}
