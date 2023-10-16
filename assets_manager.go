package lcuapi

import (
	"encoding/json"
	"fmt"
	"github.com/Vikyanite/lcu-driver/model"
	"strconv"
	"strings"
)

type AssetsManager struct {
	champions    map[int]model.Champion
	queues       map[int]model.Queue
	perks        map[int]model.Perk
	items        map[int]model.Item
	spells       map[int]model.Spell
	perkStyles   map[int]model.PerkStyle
	tiers        map[string]model.Tier
	profileIcons []model.ProfileIcon
}

var assetsManagerInstance *AssetsManager

func AssetsManagerInstance() *AssetsManager {
	if assetsManagerInstance != nil {
		return assetsManagerInstance
	}
	assetsManagerInstance = &AssetsManager{
		champions:  map[int]model.Champion{},
		queues:     map[int]model.Queue{},
		perks:      map[int]model.Perk{},
		items:      map[int]model.Item{},
		spells:     map[int]model.Spell{},
		perkStyles: map[int]model.PerkStyle{},
		tiers: map[string]model.Tier{
			"": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/unranked.png",
				Name:     "暂无段位",
			},
			"IRON": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/iron.png",
				Name:     "黑铁",
			},
			"BRONZE": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/bronze.png",
				Name:     "英勇青铜",
			},
			"SILVER": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/silver.png",
				Name:     "不屈白银",
			},
			"GOLD": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/goldpng",
				Name:     "荣耀黄金",
			},
			"PLATINUM": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/platinum.png",
				Name:     "华贵铂金",
			},
			"EMERALD": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/emerald.png",
				Name:     "流光翡翠",
			},
			"DIAMOND": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/diamond.png",
				Name:     "璀璨钻石",
			},
			"MASTER": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/master.png",
				Name:     "超凡大师",
			},
			"GRANDMASTER": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/grandmaster.png",
				Name:     "傲世宗师",
			},
			"CHALLENGER": {
				IconPath: "https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-shared-components/global/default/challenger.png",
				Name:     "最强王者",
			},
		},
		profileIcons: []model.ProfileIcon{},
	}
	return assetsManagerInstance
}

func (a *AssetsManager) GetChampionById(id int) model.Champion {
	return a.champions[id]
}

func (a *AssetsManager) GetQueueById(id int) model.Queue {
	return a.queues[id]
}

func (a *AssetsManager) GetPerkById(id int) model.Perk {
	return a.perks[id]
}

func (a *AssetsManager) GetItemById(id int) model.Item {
	return a.items[id]
}

func (a *AssetsManager) GetSpellById(id int) model.Spell {
	return a.spells[id]
}

func (a *AssetsManager) GetPerkStyleById(id int) model.PerkStyle {
	return a.perkStyles[id]
}

func (a *AssetsManager) GetProfileIconById(id int) model.ProfileIcon {
	return a.profileIcons[id]
}

func (a *AssetsManager) GetTierByName(name string) model.Tier {
	return a.tiers[name]
}

func (a *AssetsManager) Init() (err error) {
	if err = a.getChampions(); err != nil {
		return
	}
	if err = a.getQueues(); err != nil {
		return
	}
	if err = a.getPerks(); err != nil {
		return
	}
	if err = a.getItems(); err != nil {
		return
	}
	if err = a.getSpells(); err != nil {
		return
	}
	if err = a.getPerkStyles(); err != nil {
		return
	}
	//if err = a.getTiers(); err != nil {
	//	return
	//}
	if err = a.getProfileIcons(); err != nil {
		return
	}

	return
}

func (a *AssetsManager) Fill(data interface{}) (err error) {
	switch ob := data.(type) {
	case *model.Summoner:
		err = a.fillSummoner(ob)
	case *model.MatchData:
		err = a.fillMatchData(ob)
	case *model.RankedStats:
		err = a.fillRankedStats(ob)
	case *model.MatchHistory:
		err = a.fillMatchHistory(ob)
	default:
		err = fmt.Errorf("unknown type %T", ob)
	}
	return
}

func (a *AssetsManager) fillSummoner(data *model.Summoner) (err error) {
	data.ProfileIconOb = a.GetProfileIconById(data.ProfileIconId)
	return
}

func (a *AssetsManager) fillMatchHistory(data *model.MatchHistory) (err error) {
	return
}

func (a *AssetsManager) fillMatchData(match *model.MatchData) (err error) {
	match.QueueObject = a.GetQueueById(match.QueueId)
	//因为 "排位赛 单双排/灵活组排"这个描述太长了，所以要处理一下,去掉"排位赛"前缀
	match.QueueObject.Description = strings.Replace(match.QueueObject.Description, "排位赛", "", 1)

	for i := range match.ParticipantIdentities {
		match.ParticipantIdentities[i].Player.ProfileIconOb = a.GetProfileIconById(match.ParticipantIdentities[i].Player.ProfileIcon)
	}

	for i := range match.Participants {

		match.Participants[i].Spell1Object = a.GetSpellById(match.Participants[i].Spell1Id)
		match.Participants[i].Spell2Object = a.GetSpellById(match.Participants[i].Spell2Id)

		match.Participants[i].Stats.Perk0Object = a.GetPerkById(match.Participants[i].Stats.Perk0)

		match.Participants[i].Stats.PerkSubStyleObject = a.GetPerkStyleById(match.Participants[i].Stats.PerkSubStyle)

		match.Participants[i].ChampionObject = a.GetChampionById(match.Participants[i].ChampionId)

		match.Participants[i].Stats.Item0Object = a.GetItemById(match.Participants[i].Stats.Item0)
		match.Participants[i].Stats.Item1Object = a.GetItemById(match.Participants[i].Stats.Item1)
		match.Participants[i].Stats.Item2Object = a.GetItemById(match.Participants[i].Stats.Item2)
		match.Participants[i].Stats.Item3Object = a.GetItemById(match.Participants[i].Stats.Item3)
		match.Participants[i].Stats.Item4Object = a.GetItemById(match.Participants[i].Stats.Item4)
		match.Participants[i].Stats.Item5Object = a.GetItemById(match.Participants[i].Stats.Item5)
		match.Participants[i].Stats.Item6Object = a.GetItemById(match.Participants[i].Stats.Item6)
	}
	for i := range match.Teams {
		for j := range match.Teams[i].Bans {
			match.Teams[i].Bans[j].ChampionObject = a.GetChampionById(match.Teams[i].Bans[j].ChampionId)
		}
	}
	return
}

func (a *AssetsManager) fillRankedStats(data *model.RankedStats) (err error) {
	data.QueueMap.RANKEDSOLO5X5.TierOb = a.GetTierByName(data.QueueMap.RANKEDSOLO5X5.Tier)
	data.QueueMap.RANKEDFLEXSR.TierOb = a.GetTierByName(data.QueueMap.RANKEDFLEXSR.Tier)
	return
}

func (a *AssetsManager) getPerkStyles() (err error) {
	binData, err := GET("/lol-game-data/assets/v1/perkstyles.json")
	if err != nil {
		return err
	}
	var array model.PerkStyles
	err = json.Unmarshal(binData, &array)
	if err != nil {
		return err
	}
	for i := range array.Styles {
		a.perkStyles[array.Styles[i].Id] = array.Styles[i]
	}
	return
}

func (a *AssetsManager) getProfileIcons() (err error) {
	binData, err := GET("/lol-game-data/assets/v1/profile-icons.json")
	if err != nil {
		return err
	}
	var array []model.ProfileIcon
	err = json.Unmarshal(binData, &array)
	if err != nil {
		return err
	}
	for i := range array {
		a.profileIcons = append(a.profileIcons, array[i])
	}
	return
}

func (a *AssetsManager) getChampions() (err error) {
	binData, err := GET("/lol-game-data/assets/v1/champion-summary.json")
	if err != nil {
		return err
	}
	var array []model.Champion
	err = json.Unmarshal(binData, &array)
	if err != nil {
		return err
	}
	for i := range array {
		a.champions[array[i].Id] = array[i]
	}
	return
}

func (a *AssetsManager) getQueues() (err error) {
	binData, err := GET("/lol-game-data/assets/v1/queues.json")
	if err != nil {
		return err
	}
	mp := map[string]model.Queue{}
	err = json.Unmarshal(binData, &mp)
	if err != nil {
		return err
	}
	for k, v := range mp {
		id, _ := strconv.Atoi(k)
		v.Id = id
		a.queues[id] = v
	}
	return
}

func (a *AssetsManager) getPerks() (err error) {
	binData, err := GET("/lol-game-data/assets/v1/perks.json")
	if err != nil {
		return err
	}
	var array []model.Perk
	err = json.Unmarshal(binData, &array)
	if err != nil {
		return err
	}
	for i := range array {
		a.perks[array[i].Id] = array[i]
	}
	return
}

func (a *AssetsManager) getItems() (err error) {
	binData, err := GET("/lol-game-data/assets/v1/items.json")
	if err != nil {
		return err
	}
	var array []model.Item
	err = json.Unmarshal(binData, &array)
	if err != nil {
		return err
	}
	for i := range array {
		a.items[array[i].Id] = array[i]
	}
	a.items[0] = model.Item{
		Id: 0,
	}
	return
}

func (a *AssetsManager) getSpells() (err error) {
	binData, err := GET("/lol-game-data/assets/v1/summoner-spells.json")
	if err != nil {
		return err
	}
	var array []model.Spell
	err = json.Unmarshal(binData, &array)
	if err != nil {
		return err
	}
	for i := range array {
		a.spells[array[i].Id] = array[i]
	}
	return
}
