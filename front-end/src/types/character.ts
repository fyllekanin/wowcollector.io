export interface Character {
  name: string;
  realm: string;
  region: string;
  faction: string;
  level: number;
  assets: {
    avatar: string;
    inset: string;
    mainRaw: string;
  };
}
