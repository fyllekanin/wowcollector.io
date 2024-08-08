import CharacterService from '~~/server/services/CharacterService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const { region, realm, name } = getRouterParams(event);
    const achievements = await CharacterService.getCharacterAchievements(
      name.toLowerCase(),
      realm.toLowerCase(),
      region.toLowerCase()
    );
    if (achievements) return achievements;
  },
});
