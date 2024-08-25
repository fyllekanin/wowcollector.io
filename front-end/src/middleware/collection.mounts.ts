import type { MountCategory } from '~/types';

export default defineNuxtRouteMiddleware(async (to) => {
  const { region, realm, name } = to.params as Record<string, string>;

  if (!region || !realm || !name) {
    return abortNavigation();
  }

  const characterStore = useCharacterStore();
  const mountsStore = useMountsStore();
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

  const { viewId } = to.query as Record<string, string>;

  const path = (() => {
    if (viewId)
      return `/api/character/${region}/${realm}/${name}/mounts?viewId=${viewId}`;
    else return `/api/character/${region}/${realm}/${name}/mounts`;
  })();

  const { data: mountData } = await useFetch<MountCategory[]>(path);

  if (!mountData.value) {
    return abortNavigation();
  }

  mountsStore.setMounts(mountData.value);
});
