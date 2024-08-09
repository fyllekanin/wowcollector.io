interface Changelog {
  version: string;
  timestamp: number;
  changes: Change[];
}

interface Change {
  type: number;
  content: string;
}

export default defineEventHandler({
  handler: () =>
    [
      {
        version: '0.0.5',
        changes: [
          {
            type: 3,
            content: 'Fixed sticking tooltip',
          },
        ],
      },
      {
        version: '0.0.4',
        changes: [
          {
            type: 3,
            content: 'Fixed bug with tooltip',
          },
          {
            type: 3,
            content: 'Fixed mounts middleware not triggering redirect',
          },
        ],
      },
    ] as Changelog[],
});
