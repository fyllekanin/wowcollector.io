export interface ToyInformation {
  name: string;
  id: number;
  isCollected: boolean;
  itemId: number;
  assets: {
    largeIcon: string;
  };
}

export interface ToyCategory {
  name: string;
  toys: ToyInformation[] | null;
  categories: ToyCategory[];
  order: number;
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
