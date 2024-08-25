import type { ToyCategory } from '../collections';

export interface ToyViewPayload {
  name: string;
  categories: ToyCategory[];
  isUnknownIncluded: boolean;
}
