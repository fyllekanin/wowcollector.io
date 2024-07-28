export default defineNuxtRouteMiddleware(async (to) => {
  const { region, realm, name } = to.params as Record<string, string>;

  if (!region || !realm || !name) {
    return abortNavigation();
  }

  const characterStore = useCharacterStore();
  const mountsStore = useMountsStore();
  const { character } = characterStore;
  const { mounts } = mountsStore;

  if (!character?.name || character.name !== name) {
    characterStore.setCharacter({ region, realm, name });
  }

  if (!mounts?.length) {
    const { data: mountData } = await useFetch(
      `/api/character/${region}/${realm}/${name}/mounts`
    );
    if (!mountData.value) {
      return abortNavigation();
    }

    mountsStore.setMounts(mountData.value);
  }
});
