export interface ToyInformation {
  name: string;
  id: number | string;
  isCollected: boolean;
  itemId: number;
  order: number;
  assets: {
    largeIcon: string;
  };
  faction?: 'ALLIANCE' | 'HORDE';
}

export interface ToyCategory {
  name: string;
  toys: ToyInformation[] | null;
  categories: ToyCategory[];
  order: number;
  id: string;
}

export interface ToyFilters {
  search: string;
  rootCategories: string[];
  subCategories: string[];
  miscFilters: string[];
  sort: ToySort;
}

export type ToySort =
  | 'Default'
  | 'Not Collected'
  | 'Collected'
  | 'Name Ascending'
  | 'Name Descending';
