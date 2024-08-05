export interface PetInformation {
  name: string;
  id: number;
  isCollected: boolean;
  assets: {
    largeIcon: string;
  };
}

export interface PetCategory {
  name: string;
  pets: PetInformation[] | null;
  categories: PetCategory[];
  order: number;
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
