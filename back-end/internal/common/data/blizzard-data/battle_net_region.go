package blizzarddata

type BattleNetRegion string

const (
	REGION_EU BattleNetRegion = "eu"
	REGION_US BattleNetRegion = "us"
)

func FromString(s string) BattleNetRegion {
	switch s {
	case "eu":
		return REGION_EU
	case "us":
		return REGION_US
	default:
		return ""
	}
}
