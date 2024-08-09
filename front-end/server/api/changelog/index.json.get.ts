export enum ChangelogType {
  BREAKING = '1',
  FEATURE = '2',
  BUG = '3'
}
  
interface Changelog {
  version: string;
  timestamp: string;
  changes: Change[];
}

interface Change {
  type: ChangelogType;
  content: string;
}

const data: Changelog[] =
[] as Changelog[];

export default defineEventHandler({
  handler: () => data,
});