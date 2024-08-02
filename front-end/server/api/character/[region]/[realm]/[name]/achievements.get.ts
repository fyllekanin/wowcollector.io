import { flatMapAchievements } from '~/utils';
import CharacterService from '~~/server/services/CharacterService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const { region, realm, name } = getRouterParams(event);
    const achievements = await CharacterService.getCharacterAchievements(
      name,
      realm,
      region
    );

    if (achievements) {
      const flattenedAchievements = flatMapAchievements(achievements);
      console.log(
        flattenedAchievements.filter((achievement) => achievement.id === 8891)
      );
    }

    if (achievements) return achievements;
  },
});
