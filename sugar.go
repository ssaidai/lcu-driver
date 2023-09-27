package lcuapi

import (
	"encoding/json"
	"fmt"
	"lcu/model"
)

var (
	sugarConn = &Connection{
		IInquirer: &NilInquirer{},
		IWatcher:  &NilWatcher{},
	}
)

func Init() (err error) {
	sugarConn, err = NewConnection()
	if err != nil {
		return
	}
	err = sugarConn.Init()
	return
}

func GetCurrentSummoner() (data model.Summoner, err error) {
	resp, err := sugarConn.Get("/lol-summoner/v1/current-summoner")
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &data)
	return
}

func GetSummonerByName(name string) (data model.Summoner, err error) {
	resp, err := sugarConn.Get("/lol-summoner/v1/summoners?name=" + name)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &data)
	return
}

func GetMatchHistoryByPuuid(puuid string, begI, endI int) (data model.MatchHistory, err error) {
	url := fmt.Sprintf("/lol-match-history/v1/products/lol/%s/matches?begIndex=%d&endIndex=%d", puuid, begI, endI)
	resp, err := sugarConn.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &data)
	return
}

func GetMatchDetailsByGameId(gameId int) (data model.MatchData, err error) {
	url := fmt.Sprintf("/lol-match-history/v1/games/%d", gameId)
	resp, err := sugarConn.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &data)
	return
}

func GetRankedStatsByPuuid(puuid string) (data model.RankedStats, err error) {
	url := fmt.Sprintf("/lol-ranked/v1/ranked-stats/%s", puuid)
	resp, err := sugarConn.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &data)
	return
}
