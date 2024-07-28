export interface MountInformation {
  name: string;
  description: string;
  id: number;
  isCollected: boolean;
  creatureDisplay: string;
  icon: string;
}

export interface MountCategory {
  category: string;
  name: string;
  mounts: MountInformation[] | null;
  categories: MountCategory[];
  order: number;
}

export interface WoWHeadMountInformation {
  completion_category: string;
  icon: string;
  name: string;
  quality: number;
  spells: unknown[];
  tooltip: string;
}
