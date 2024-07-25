package response

type MountCollectionMount struct {
	Name            string
	Description     string
	Id              int
	IsCollected     bool
	CreatureDisplay string
	Icon            string
}

type MountCollectionCategory struct {
	Name       string
	Mounts     []MountCollectionMount
	Categories []MountCollectionCategory
	Order      int
}
