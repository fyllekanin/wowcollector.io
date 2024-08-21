export default defineNuxtRouteMiddleware(async () => {
  const toyViewBuilderStore = useToyViewBuilderStore();

  const { data: toys } = await useFetch('/api/battle-net/toys');

  if (!toys.value) {
    return abortNavigation();
  }

  toyViewBuilderStore.setToys(toys.value);
});
