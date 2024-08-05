import type { MountCategory } from './mounts';
import type { AchievementCategory } from './achievements';
import type { ToyCategory } from './toys';

export * from './mounts';
export * from './achievements';
export * from './toys';

export type CollectionInformationResponse<T> =
  | MountCategory[]
  | AchievementCategory[]
  | ToyCategory[];
