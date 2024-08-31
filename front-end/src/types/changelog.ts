export enum ChangelogType {
  BREAKING = '1',
  FEATURE = '2',
  BUG = '3',
  IMPROVEMENT = '4'
}

export interface Changelog {
  version: string;
  timestamp: string;
  changes: Change[];
}

export interface Change {
  type: ChangelogType;
  content: string;
}
