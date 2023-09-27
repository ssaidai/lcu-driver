package model

type BlockPlayer struct {
	GameName   string `json:"gameName"`
	GameTag    string `json:"gameTag"`
	Icon       int    `json:"icon"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Pid        string `json:"pid"`
	Puuid      string `json:"puuid"`
	SummonerId int    `json:"summonerId"`
}

type ChatConfig struct {
	ChatDomain ChatDomain `json:"ChatDomain"`
	LcuSocial  LcuSocial  `json:"LcuSocial"`
}

type ChatDomain struct {
	ChampSelectDomainName string `json:"ChampSelectDomainName"`
	ClubDomainName        string `json:"ClubDomainName"`
	CustomGameDomainName  string `json:"CustomGameDomainName"`
	P2PDomainName         string `json:"P2PDomainName"`
	PostGameDomainName    string `json:"PostGameDomainName"`
}

type PlatformToRegionMap struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}

type LcuSocial struct {
	AggressiveScanning     bool                `json:"AggressiveScanning"`
	ForceChatFilter        bool                `json:"ForceChatFilter"`
	QueueJobGraceSeconds   int                 `json:"QueueJobGraceSeconds"`
	ReplaceRichMessages    bool                `json:"ReplaceRichMessages"`
	SilenceChatWhileInGame bool                `json:"SilenceChatWhileInGame"`
	AllowGroupByGame       bool                `json:"allowGroupByGame"`
	GameNameTaglineEnabled bool                `json:"gameNameTaglineEnabled"`
	PlatformToRegionMap    PlatformToRegionMap `json:"platformToRegionMap"`
}
