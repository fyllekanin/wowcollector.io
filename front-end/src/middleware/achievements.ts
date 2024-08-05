import type { AchievementCategory, AchievementCategoryResponse } from '~/types';

export default defineNuxtRouteMiddleware(async (to) => {
  const { region, realm, name } = to.params as Record<string, string>;

  if (!region || !realm || !name) {
    return abortNavigation();
  }

  const characterStore = useCharacterStore();
  const achievementsStore = useAchievementsStore();
  const { character } = storeToRefs(characterStore);
  const { achievements } = storeToRefs(achievementsStore);

  characterStore.setCharacter({ region, realm, name });

  if (!achievements.value?.length || character.value?.name !== name) {
    const { data: rootCategories } = await useFetch(
      '/api/battle-net/achievement-root-categories'
    );
    if (!rootCategories.value) {
      return abortNavigation();
    }

    achievementsStore.setAchievements(
      rootCategories.value as AchievementCategory[]
    );

    const { data } = useFetch<AchievementCategoryResponse>(
      `/api/character/${region}/${realm}/${name}/achievements`
    );

    if (!data.value) {
      return abortNavigation();
    }

    achievementsStore.setAchievements(data.value.categories);
    achievementsStore.setTotal(data.value.total);
    achievementsStore.setCompleted(data.value.completed);
  }
});
