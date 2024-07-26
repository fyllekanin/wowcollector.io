export interface MountInformation {
  name: string;
  description: string;
  id: number;
  isCollected: boolean;
  creatureDisplay: string;
  icon: string;
}

export interface MountCategory {
  category: Category;
  name: string;
  mounts: MountInformation[] | null;
  subCategories: MountCategory[];
  order: number;
}

export type Category = string;
// // Root
// | 'Unknown'
// | 'Classic'
// | 'Race'
// | 'The burning crusade'
// | 'Wrath of the lich king'
// | 'Cataclysm'
// | 'Mists of pandaria'
// | 'Warlords of draenor'
// | 'Legion'
// | 'Battle for azeroth'
// | 'Shadowlands'
// | 'Dragonflight'
// | 'Miscellaneous'
// | 'Professions'
// | 'PVP'
// | 'Promotion'
// | 'World events'
// | 'Class mounts'
// // Sub
// | 'Reputation'
// | 'Dungeon drop'
// | 'Raid drop'
// | 'Blood elves'
// | 'Dracthyr'
// | 'Goblin'

export interface WoWHeadMountInformation {
  completion_category: string;
  icon: string;
  name: string;
  quality: number;
  spells: unknown[];
  tooltip: string;
}
