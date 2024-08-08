export default defineNuxtRouteMiddleware(async (to) => {
  const { region, realm, name } = to.params as Record<string, string>;

  if (!region || !realm || !name) {
    return abortNavigation();
  }

  const characterStore = useCharacterStore();
  const achievementsStore = useAchievementsStore();
  const { character } = storeToRefs(characterStore);
  const { achievements } = storeToRefs(achievementsStore);

  if (
    character.value?.name !== name &&
    character.value?.region !== region &&
    character.value?.realm !== realm
  ) {
    const { data: characterData } = await useFetch(
      `/api/character/${region}/${realm}/${name}`
    );

    if (!characterData.value) {
      return abortNavigation();
    }

    characterStore.setCharacter(characterData.value);
  }

  if (!achievements.value?.length || character.value?.name !== name) {
    const { data: achievementData } = await useFetch(
      `/api/character/${region}/${realm}/${name}/achievements`
    );

    if (!achievementData.value) {
      return abortNavigation();
    }

    const minDisplayOrder =
      Math.min(...achievementData.value?.map((c) => c.displayOrder || 0)) || 0;
    const accordionCategories = achievementData.value.map((category) => ({
      ...category,
      defaultOpen: category.displayOrder === minDisplayOrder,
    }));

    achievementsStore.setAchievements(accordionCategories);
  }
});
