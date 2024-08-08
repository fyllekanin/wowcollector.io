export default defineNuxtRouteMiddleware(async (to) => {
  const { region, realm, name } = to.params as Record<string, string>;

  if (!region || !realm || !name) {
    return abortNavigation();
  }

  const characterStore = useCharacterStore();
  const mountsStore = useMountsStore();
  const { character } = storeToRefs(characterStore);
  const { mounts } = storeToRefs(mountsStore);

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

    characterStore.setCharacter({
      ...characterData.value,
      region,
    });
  }

  if (!mounts.value?.length || character.value?.name !== name) {
    const { data: mountData } = await useFetch(
      `/api/character/${region}/${realm}/${name}/mounts`
    );

    if (!mountData.value) {
      return abortNavigation();
    }

    mountsStore.setMounts(mountData.value);
  }
});
