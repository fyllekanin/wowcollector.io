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
    const { data: achievementData } = await useFetch(
      `/api/character/${region}/${realm}/${name}/achievements`
    );
    if (!achievementData.value) {
      return abortNavigation();
    }
    achievementsStore.setAchievements(achievementData.value);
  }
});
