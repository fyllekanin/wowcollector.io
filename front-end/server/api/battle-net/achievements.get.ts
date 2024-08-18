import BattlenetService from '~~/server/services/BattlenetService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const achievements = await BattlenetService.getAchievements();
    if (achievements) return achievements;
  },
});
