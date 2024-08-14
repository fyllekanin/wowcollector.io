import type { MountCategory, MountInformation } from '../collections';

export interface BuilderMountInformation extends MountInformation {
  category: string | null;
  level: string | null;
}

export interface BuilderMountCategory extends MountCategory {
  mounts: BuilderMountInformation[];
  categories: BuilderMountCategory[];
}
