package lcuapi

import (
	"encoding/json"
	"fmt"
	"github.com/Vikyanite/lcu-driver/model"
)

var (
	sugarConn = NewDriver()
)

func Start(startCbs ...func() error) (chan error, error) {
	startCbs = append(startCbs, AssetsManagerInstance().Init)
	return sugarConn.Start(startCbs...)
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
	if err != nil {
		return
	}
	err = AssetsManagerInstance().Fill(&data)
	return
}

func GetSummonerByName(name string) (data model.Summoner, err error) {
	binData, err := GET("/lol-summoner/v1/summoners?name=" + name)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	if err != nil {
		return
	}
	err = AssetsManagerInstance().Fill(&data)
	return
}

func GetMatchHistoryByPuuid(puuid string, begI, endI int) (data model.MatchHistory, err error) {
	url := fmt.Sprintf("/lol-match-history/v1/products/lol/%s/matches?begIndex=%d&endIndex=%d", puuid, begI, endI)
	binData, err := GET(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	if err != nil {
		return
	}
	err = AssetsManagerInstance().Fill(&data)
	return
}

func GetMatchDetailsByGameId(gameId int) (data model.MatchData, err error) {
	url := fmt.Sprintf("/lol-match-history/v1/games/%d", gameId)
	binData, err := GET(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	if err != nil {
		return
	}
	err = AssetsManagerInstance().Fill(&data)
	return
}

func GetRankedStatsByPuuid(puuid string) (data model.RankedStats, err error) {
	url := fmt.Sprintf("/lol-ranked/v1/ranked-stats/%s", puuid)
	binData, err := GET(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(binData, &data)
	if err != nil {
		return
	}
	err = AssetsManagerInstance().Fill(&data)
	return
}

func GetChampionById(id int) model.Champion {
	return AssetsManagerInstance().GetChampionById(id)
}

func GetQueueById(id int) model.Queue {
	return AssetsManagerInstance().GetQueueById(id)
}

func GetPerkById(id int) model.Perk {
	return AssetsManagerInstance().GetPerkById(id)
}

func GetItemById(id int) model.Item {
	return AssetsManagerInstance().GetItemById(id)
}

func GetSpellById(id int) model.Spell {
	return AssetsManagerInstance().GetSpellById(id)
}

func GetPerkStyleById(id int) model.PerkStyle {
	return AssetsManagerInstance().GetPerkStyleById(id)
}

func GetProfileIconById(id int) model.ProfileIcon {
	return AssetsManagerInstance().GetProfileIconById(id)
}
