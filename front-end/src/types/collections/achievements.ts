export interface AchievementInformation {
  name: string;
  description: string;
  id: number;
  isCompleted: boolean;
  icon: string;
  points: number;
  displayOrder: number;
  faction?: 'ALLIANCE' | 'HORDE';
}

export interface AchievementCategory {
  name: string;
  achievements: AchievementInformation[] | null;
  categories: AchievementCategory[];
  displayOrder: number;
  id: number | string;
}

export interface AchievementCategoryAccordion extends AchievementCategory {
  deafultOpen?: boolean;
}

export type AchievementSort =
  | 'Default'
  | 'Not Completed'
  | 'Completed'
  | 'Name Ascending'
  | 'Name Descending';

export interface AchievementFilters {
  search: string;
  rootCategories: string[];
  subCategories: string[];
  miscFilters: string[];
  sort: AchievementSort;
}
