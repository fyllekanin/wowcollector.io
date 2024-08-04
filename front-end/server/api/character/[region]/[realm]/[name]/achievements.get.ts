import CharacterService from '~~/server/services/CharacterService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const { region, realm, name } = getRouterParams(event);
    const { rootCategoryId } = getQuery(event);
    const achievements = await CharacterService.getCharacterAchievements(
      name,
      realm,
      region,
      rootCategoryId as number
    );
    if (achievements) return achievements;
  },
});
