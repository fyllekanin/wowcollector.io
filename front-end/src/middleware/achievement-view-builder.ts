export default defineNuxtRouteMiddleware(async () => {
  const achievementViewBuilderStore = useAchievementViewBuilderStore();

  const { data: achievements } = await useFetch('/api/battle-net/achievements');

  if (!achievements.value) {
    return abortNavigation();
  }

  achievementViewBuilderStore.setAchievements(achievements.value);
});
