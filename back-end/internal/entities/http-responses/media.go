package httpresponses

type BattleNetAsset struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BattleNetMedia struct {
	Assets []*BattleNetAsset `json:"assets"`
	Id     int               `json:"id"`
}

func (m *BattleNetMedia) GetIconAsset() string {
	return m.GetAssetByKey("icon")
}

func (m *BattleNetMedia) GetAssetByKey(key string) string {
	for _, element := range m.Assets {
		if element.Key == key {
			return element.Value
		}
	}
	return ""
}
