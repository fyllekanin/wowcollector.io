export interface PetInformation {
  name: string;
  id: number | string;
  isCollected: boolean;
  assets: {
    largeIcon: string;
  };
  order: number;
}

export interface PetCategory {
  name: string;
  pets: PetInformation[] | null;
  categories: PetCategory[];
  order: number;
  id: string;
}

export interface PetFilters {
  search: string;
  rootCategories: string[];
  subCategories: string[];
  miscFilters: string[];
  sort: PetSort;
}

export type PetSort =
  | 'Default'
  | 'Not Collected'
  | 'Collected'
  | 'Name Ascending'
  | 'Name Descending';
