export interface AchievementInformation {
  name: string;
  description: string;
  id: number;
  isCompleted: boolean;
  icon: string;
  points: number;
  displayOrder: number;
}

export interface AchievementCategory {
  name: string;
  achievements: AchievementInformation[] | null;
  categories: AchievementCategory[];
  displayOrder: number;
  id: number;
}

export interface AchievementCategoryResponse {
  total: number;
  completed: number;
  categories: Array<AchievementCategory>
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
