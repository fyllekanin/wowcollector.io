import type { AchievementCategory } from '~/types';

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

    console.log(rootCategories.value);

    achievementsStore.setAchievements(
      rootCategories.value as AchievementCategory[]
    );

    const responses = await Promise.all([
      useFetch<AchievementCategory[]>(
        `/api/character/${region}/${realm}/${name}/achievements`,
        {
          query: {
            rootCategoryId: achievements.value[0]?.id,
          },
        }
      ),
      useFetch<AchievementCategory[]>(
        `/api/character/${region}/${realm}/${name}/achievements`,
        {
          query: {
            rootCategoryId: achievements.value[1]?.id,
          },
        }
      ),
    ]);

    if (responses.some((response) => !response.data.value)) {
      return abortNavigation();
    }

    const mergedAchievements = responses.reduce((acc, { data }, index) => {
      // const [category] = data.value;
      const merged = {
        ...achievements.value[index],
        categories: data.value,
        achievements: null,
      } as AchievementCategory;
      acc.push(merged);
      return acc;
    }, [] as AchievementCategory[]);

    mergedAchievements.forEach((achievementCategory) => {
      // console.log(achievementCategory);
      achievementsStore.setAchievement(achievementCategory);
    });
  }
});
