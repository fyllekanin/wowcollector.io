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
    console.time('achievements 1');
    const { data: rootCategories } = await useFetch(
      '/api/battle-net/achievement-root-categories'
    );
    console.timeEnd('achievements 1');
    if (!rootCategories.value) {
      return abortNavigation();
    }

    achievementsStore.setAchievements(
      rootCategories.value as AchievementCategory[]
    );

    console.time('achievements 2');
    const responses = await Promise.all([
      useFetch<AchievementCategoryResponse>(
        `/api/character/${region}/${realm}/${name}/achievements`,
        {
          query: {
            rootCategoryId: achievements.value[0]?.id,
          },
        }
      ),
      useFetch<AchievementCategoryResponse>(
        `/api/character/${region}/${realm}/${name}/achievements`,
        {
          query: {
            rootCategoryId: achievements.value[1]?.id,
          },
        }
      ),
    ]);
    console.timeEnd('achievements 2');

    if (responses.some((response) => !response.data.value)) {
      return abortNavigation();
    }

    console.time('achievements 3');
    const mergedAchievements = responses.reduce((prev, curr, index) => {
      const merged = {
        ...achievements.value[index],
        categories: curr.data.value?.categories,
        achievements: null,
      } as AchievementCategory;
      prev.push(merged);
      if (curr.data.value?.total)
        achievementsStore.setTotal(curr.data.value.total);
      if (curr.data.value?.completed)
        achievementsStore.setCompleted(curr.data.value.completed);
      return prev;
    }, [] as AchievementCategory[]);
    console.timeEnd('achievements 3');

    console.time('achievements 4');
    mergedAchievements.forEach((achievementCategory) => {
      achievementsStore.setAchievement(achievementCategory);
    });
    console.timeEnd('achievements 4');
  }
});
