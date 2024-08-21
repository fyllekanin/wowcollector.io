export default defineNuxtRouteMiddleware(async () => {
  const petViewBuilderStore = usePetViewBuilderStore();

  const { data: pets } = await useFetch('/api/battle-net/pets');

  if (!pets.value) {
    return abortNavigation();
  }

  petViewBuilderStore.setPets(pets.value);
});
