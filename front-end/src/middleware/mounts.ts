export default defineNuxtRouteMiddleware(async (to) => {
  const { region, realm, name } = to.params as Record<string, string>;

  if (!region || !realm || !name) {
    return abortNavigation();
  }

  const characterStore = useCharacterStore();
  const mountsStore = useMountsStore();
  const { character } = storeToRefs(characterStore);
  const { mounts } = mountsStore;

  characterStore.setCharacter({ region, realm, name });

  if (!mounts?.length || character.value?.name !== name) {
    const { data: mountData } = await useFetch(
      `/api/character/${region}/${realm}/${name}/mounts`
    );
    if (!mountData.value) {
      return abortNavigation();
    }
    mountsStore.setMounts(mountData.value);
  }
});
