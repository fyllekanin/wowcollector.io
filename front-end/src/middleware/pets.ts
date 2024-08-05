import type { PetCategory } from '~/types';

export default defineNuxtRouteMiddleware(async (to) => {
  const { region, realm, name } = to.params as Record<string, string>;

  if (!region || !realm || !name) {
    return abortNavigation();
  }

  const characterStore = useCharacterStore();
  const petsStore = usePetsStore();
  const { character } = storeToRefs(characterStore);
  const { pets } = storeToRefs(petsStore);

  characterStore.setCharacter({ region, realm, name });

  if (!pets.value?.length || character.value?.name !== name) {
    const { data: petData } = await useFetch<PetCategory[]>(
      `/api/character/${region}/${realm}/${name}/pets`
    );
    if (!petData.value) {
      return abortNavigation();
    }
    petsStore.setPets(petData.value);
  }
});
