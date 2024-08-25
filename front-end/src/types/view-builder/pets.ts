import type { PetCategory } from '../collections';

export interface PetViewPayload {
  name: string;
  categories: PetCategory[];
  isUnknownIncluded: boolean;
}
