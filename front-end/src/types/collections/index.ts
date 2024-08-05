import type { MountCategory } from './mounts';
import type { AchievementCategory } from './achievements';
import type { ToyCategory } from './toys';
import type { PetCategory } from './pets';

export * from './mounts';
export * from './achievements';
export * from './toys';
export * from './pets';

export type CollectionInformationResponse<T> =
  | MountCategory[]
  | AchievementCategory[]
  | ToyCategory[]
  | PetCategory[];
