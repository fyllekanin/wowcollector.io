import type { MountCategory } from '../collections';

export interface MountViewPayload {
  name: string;
  categories: MountCategory[];
  isUnknownIncluded: boolean;
}
