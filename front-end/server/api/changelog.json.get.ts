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
[
  {
    "version": "0.1.0",
    "timestamp": "1723275848",
    "changes": [
      {
        "type": "2",
        "content": "Add release functionality to the repository"
      }
    ]
  }
] as Changelog[];

export default defineEventHandler({
  handler: () => data,
});