package structs

type Taro struct {
	ID   string `json:"id"`

	Loc TaroLoc `json:"name"`
	Advice  TaroMeaning `json:"advice"`
	Warning TaroMeaning `json:"warning"`
	Pic TaroPic `json:"pic"`
}
