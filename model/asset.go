package model

type ProfileIcon struct {
	Id       int    `json:"id"`
	IconPath string `json:"iconPath"`
}

type PerkStyles struct {
	SchemaVersion int         `json:"schemaVersion"`
	Styles        []PerkStyle `json:"styles"`
}

type Perk struct {
	Id                       int      `json:"id"`
	Name                     string   `json:"name"`
	MajorChangePatchVersion  string   `json:"majorChangePatchVersion"`
	Tooltip                  string   `json:"tooltip"`
	ShortDesc                string   `json:"shortDesc"`
	LongDesc                 string   `json:"longDesc"`
	RecommendationDescriptor string   `json:"recommendationDescriptor"`
	IconPath                 string   `json:"iconPath"`
	EndOfGameStatDescs       []string `json:"endOfGameStatDescs"`
	RecommendationAttributes struct{} `json:"recommendationDescriptorAttributes"`
}

type Spell struct {
	Id            int      `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	SummonerLevel int      `json:"summonerLevel"`
	Cooldown      int      `json:"cooldown"`
	GameModes     []string `json:"gameModes"`
	IconPath      string   `json:"iconPath"`
}

type Queue struct {
	Name                string `json:"name"`
	ShortName           string `json:"shortName"`
	Description         string `json:"description"`
	Id                  int    `json:"id"`
	DetailedDescription string `json:"detailedDescription"`
}

type Champion struct {
	Id                 int      `json:"id"`
	Name               string   `json:"name"`
	Alias              string   `json:"alias"`
	SquarePortraitPath string   `json:"squarePortraitPath"`
	Roles              []string `json:"roles"`
}

type Item struct {
	Id                       int      `json:"id"`
	Name                     string   `json:"name"`
	Description              string   `json:"description"`
	Active                   bool     `json:"active"`
	InStore                  bool     `json:"inStore"`
	From                     []int    `json:"from"`
	To                       []int    `json:"to"`
	Categories               []string `json:"categories"`
	MaxStacks                int      `json:"maxStacks"`
	RequiredChampion         string   `json:"requiredChampion"`
	RequiredAlly             string   `json:"requiredAlly"`
	RequiredBuffCurrencyName string   `json:"requiredBuffCurrencyName"`
	RequiredBuffCurrencyCost int      `json:"requiredBuffCurrencyCost"`
	SpecialRecipe            int      `json:"specialRecipe"`
	IsEnchantment            bool     `json:"isEnchantment"`
	Price                    int      `json:"price"`
	PriceTotal               int      `json:"priceTotal"`
	IconPath                 string   `json:"iconPath"`
}

type Tier struct {
	IconPath string `json:"iconPath"`
	Name     string `json:"name"`
	Division string `json:"division"`
}

type PerkStyle struct {
	Id                         int                          `json:"id"`
	Name                       string                       `json:"name"`
	Tooltip                    string                       `json:"tooltip"`
	IconPath                   string                       `json:"iconPath"`
	AssetMap                   map[string]string            `json:"assetMap"`
	IsAdvanced                 bool                         `json:"isAdvanced"`
	AllowedSubStyles           []int                        `json:"allowedSubStyles"`
	SubStyleBonus              []SubStyleBonus              `json:"subStyleBonus"`
	Slots                      []Slots                      `json:"slots"`
	DefaultPageName            string                       `json:"defaultPageName"`
	DefaultSubStyle            int                          `json:"defaultSubStyle"`
	DefaultPerks               []int                        `json:"defaultPerks"`
	DefaultPerksWhenSplashed   []int                        `json:"defaultPerksWhenSplashed"`
	DefaultStatModsPerSubStyle []DefaultStatModsPerSubStyle `json:"defaultStatModsPerSubStyle"`
}

type SubStyleBonus struct {
	StyleId int `json:"styleId"`
	PerkId  int `json:"perkId"`
}

type Slots struct {
	Type      string `json:"type"`
	SlotLabel string `json:"slotLabel"`
	Perks     []int  `json:"perks"`
}

type DefaultStatModsPerSubStyle struct {
	Id    string `json:"id"`
	Perks []int  `json:"perks"`
}
