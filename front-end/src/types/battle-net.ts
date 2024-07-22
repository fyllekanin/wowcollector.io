interface Realm {
  Name: string;
  Region: string;
  Slug: string;
}

interface Region {
  Name: string;
  Value: string;
}

export interface RealmsRegions {
  Realms: Realm[];
  Regions: Region[];
}
