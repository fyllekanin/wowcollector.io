import type { MountCategory } from './mounts';
import type { AchievementCategory } from './achievements';

export * from './mounts';
export * from './achievements';

export type CollectionInformationResponse<T> =
  | MountCategory[]
  | AchievementCategory[];
