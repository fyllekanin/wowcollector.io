export default defineNuxtRouteMiddleware(async (to) => {
  const { region, realm, name } = to.params as Record<string, string>;

  if (!region || !realm || !name) {
    return abortNavigation();
  }

  const characterStore = useCharacterStore();
  const toysStore = useToysStore();
  const { character } = storeToRefs(characterStore);

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

  const { data: toyData } = await useFetch(
    `/api/character/${region}/${realm}/${name}/toys`
  );

  if (!toyData.value) {
    return abortNavigation();
  }

  toysStore.setToys(toyData.value);
});
