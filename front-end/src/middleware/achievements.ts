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
    const { data } = await useLazyFetch<AchievementCategory[]>(
      `/api/character/${region}/${realm}/${name}/achievements`
    );

    if (!data.value) {
      return abortNavigation();
    }

    const minDisplayOrder =
      Math.min(...data.value?.map((c) => c.displayOrder || 0)) || 0;
    const accordionCategories = data.value.map((category) => ({
      ...category,
      defaultOpen: category.displayOrder === minDisplayOrder,
    }));

    achievementsStore.setAchievements(accordionCategories);
  }
});
