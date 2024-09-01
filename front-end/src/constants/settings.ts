import type { MountSort, PetSort, ToySort, AchievementSort } from '~/types';

export const RENDER_TYPES = [
  {
    label: 'Compact Grid',
    value: 'grid-compact',
    icon: 'material-symbols:background-grid-small-sharp',
  },
  {
    label: 'Grid',
    value: 'grid',
    icon: 'material-symbols:grid-view',
  },
  {
    label: 'List',
    value: 'list',
    icon: 'material-symbols:lists',
  },
];

type Sort = MountSort | PetSort | ToySort;
export const SORT_TYPES: Sort[] = [
  'Default',
  'Collected',
  'Not Collected',
  'Name Ascending',
  'Name Descending',
];
export const ACHIEVEMENT_SORT_TYPES: AchievementSort[] = [
  'Default',
  'Completed',
  'Not Completed',
  'Name Ascending',
  'Name Descending',
];
