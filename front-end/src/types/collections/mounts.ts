export interface MountInformation {
  name: string;
  description: string;
  id: number;
  isCollected: boolean;
  creatureDisplay: string;
  assets: {
    display: string;
    smallIcon: string;
    largeIcon: string;
  };
}

export interface MountCategory {
  id?: string;
  name: string;
  mounts: MountInformation[] | null;
  categories: MountCategory[];
  order: number;
}

export interface MountFilters {
  search: string;
  rootCategories: string[];
  subCategories: string[];
  miscFilters: string[];
  sort: MountSort;
  viewStyle: ViewStyle;
}

export type ViewStyle = 'grid' | 'list' | 'grid-compact';

export type MountSort =
  | 'Default'
  | 'Not Collected'
  | 'Collected'
  | 'Name Ascending'
  | 'Name Descending';

export interface WoWHeadMountInformation {
  completion_category: string;
  icon: string;
  name: string;
  quality: number;
  spells: unknown[];
  tooltip: string;
}
