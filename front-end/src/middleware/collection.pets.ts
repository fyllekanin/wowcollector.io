import type { PetCategory } from '~/types';

export default defineNuxtRouteMiddleware(async (to) => {
  const { region, realm, name } = to.params as Record<string, string>;

  if (!region || !realm || !name) {
    return abortNavigation();
  }

  const characterStore = useCharacterStore();
  const petsStore = usePetsStore();
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
      return `/api/character/${region}/${realm}/${name}/pets?viewId=${viewId}`;
    else return `/api/character/${region}/${realm}/${name}/pets`;
  })();

  const { data: petData } = await useFetch<PetCategory[]>(path);

  if (!petData.value) {
    return abortNavigation();
  }
  petsStore.setPets(petData.value);
});
