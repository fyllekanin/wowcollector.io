export interface Realm {
  name: string;
  region: string;
  slug: string;
}

export interface Region {
  name: string;
  value: string;
}

export interface RealmsRegions {
  realms: Realm[];
  regions: Region[];
}
