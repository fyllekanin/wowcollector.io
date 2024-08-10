export enum ChangelogType {
  BREAKING = '1',
  FEATURE = '2',
  BUG = '3',
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

const data: Changelog[] = [
  {
    version: '1.0.0',
    timestamp: '1723129817',
    changes: [
      {
        type: ChangelogType.FEATURE,
        content: 'Initial release',
      },
    ],
  },
  {
    version: '1.0.1',
    timestamp: '1723216217',
    changes: [
      {
        type: ChangelogType.FEATURE,
        content: 'Added new feature',
      },
      {
        type: ChangelogType.BUG,
        content: 'Fixed bug',
      },
    ],
  },
] as Changelog[];

export default defineEventHandler({
  handler: () => data,
});
