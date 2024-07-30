export interface Filters {
  search: string;
  rootCategories: string[];
  subCategories: string[];
  miscFilters: string[];
}

export interface MountFilters {
  search: string;
  rootCategories: string[];
  subCategories: string[];
  miscFilters: string[];
  sort: Sort;
  viewStyle: ViewStyle;
}

export type ViewStyle = 'grid' | 'list' | 'grid-compact';

export type Sort =
  | 'Default'
  | 'Not Collected'
  | 'Collected'
  | 'Name Ascending'
  | 'Name Descending';
