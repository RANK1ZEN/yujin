package db

import "gorm.io/gorm"

type Region string

const (
	RegionNA Region = "na"
	RegionEUW Region = "euw"
	RegionKR Region = "kr"
)

type Summoner struct {
	gorm.Model
	Puuid         string `json:"puuid"`
	AccountId     string `json:"accountId"`
	SummonerId    string `json:"summonerId"`
	Level         int    `json:"summonerLevel"`
	ProfileIconId int    `json:"profileIconId" `
	Name          string `json:"name"`
}