export interface Auth {
  id: string;
  displayName: string;
  tokens: {
    accessToken: string;
    refreshToken: string;
  };
  connections: {
    battleTag: string;
    discordId: string;
  };
}
