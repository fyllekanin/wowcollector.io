export enum ChangelogType {
  BREAKING = '1',
  FEATURE = '2',
  BUG = '3',
  IMPROVEMENT = '4'
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
    "version": "1.0.0",
    "timestamp": "1725276477",
    "changes": [
      {
        "type": "3",
        "content": "Add faction property for mount response"
      },
      {
        "type": "4",
        "content": "Fixed phrasing of sort type depending on the collection type"
      },
      {
        "type": "3",
        "content": "Fixed an issue with view-builder dropzone highlighting"
      },
      {
        "type": "4",
        "content": "Disabled auto zoom when focusing inputs"
      },
      {
        "type": "4",
        "content": "Dynamic logo based on theme preference"
      },
      {
        "type": "2",
        "content": "Added faction filters to view-builder"
      }
    ]
  },
  {
    "version": "0.4.0",
    "timestamp": "1725135716",
    "changes": [
      {
        "type": "4",
        "content": "Updated all database repositories to composite the common one"
      },
      {
        "type": "4",
        "content": "Make common code for creating database indexes"
      },
      {
        "type": "4",
        "content": "Add changelog check for pull requests"
      },
      {
        "type": "3",
        "content": "Fixed bug where toys and pets were not rendering"
      },
      {
        "type": "2",
        "content": "Add KR and TW support"
      },
      {
        "type": "4",
        "content": "Minor UI improvements for view builder"
      }
    ]
  },
  {
    "version": "0.3.0",
    "timestamp": "1725091554",
    "changes": [
      {
        "type": "3",
        "content": "fixed wowhead collection icons redirect url"
      },
      {
        "type": "2",
        "content": "Added API for fetching all the scanned mounts"
      },
      {
        "type": "2",
        "content": "Add API to fetch all scanned toys"
      },
      {
        "type": "2",
        "content": "Add API to fetch all scanned pets"
      },
      {
        "type": "2",
        "content": "Add item view creation APIs"
      },
      {
        "type": "2",
        "content": "Add pet and toy view fetch support in collection"
      }
    ]
  },
  {
    "version": "0.2.0",
    "timestamp": "1723486955",
    "changes": [
      {
        "type": "2",
        "content": "Added feedback / bug report API"
      }
    ]
  },
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