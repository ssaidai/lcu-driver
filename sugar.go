package lcuapi

import (
	"encoding/json"
	"fmt"
	"lcu/model"
)

var (
	sugarConn = NewDriver()
)

func Run() (err error) {
	err = sugarConn.Run()
	return
}

func GET(uri string) (resp []byte, err error) {
	if !sugarConn.IsRunning() {
		err = fmt.Errorf("lcu-driver not running")
		return
	}
	r, err := sugarConn.Get(uri)
	if err != nil {
		return
	}
	resp = r.Body()
	return
}

func POST(uri string, body interface{}) (resp []byte, err error) {
	if !sugarConn.IsRunning() {
		err = fmt.Errorf("lcu-driver not running")
		return
	}
	r, err := sugarConn.Post(uri, body)
	if err != nil {
		return
	}
	resp = r.Body()
	return
}

func GetCurrentSummoner() (data model.Summoner, err error) {
	binData, err := GET("/lol-summoner/v1/current-summoner")
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	return
}

func GetSummonerByName(name string) (data model.Summoner, err error) {
	binData, err := GET("/lol-summoner/v1/summoners?name=" + name)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	return
}

func GetMatchHistoryByPuuid(puuid string, begI, endI int) (data model.MatchHistory, err error) {
	url := fmt.Sprintf("/lol-match-history/v1/products/lol/%s/matches?begIndex=%d&endIndex=%d", puuid, begI, endI)
	binData, err := GET(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	return
}

func GetMatchDetailsByGameId(gameId int) (data model.MatchData, err error) {
	url := fmt.Sprintf("/lol-match-history/v1/games/%d", gameId)
	binData, err := GET(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	return
}

func GetRankedStatsByPuuid(puuid string) (data model.RankedStats, err error) {
	url := fmt.Sprintf("/lol-ranked/v1/ranked-stats/%s", puuid)
	binData, err := GET(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	return
}
