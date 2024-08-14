export default defineNuxtRouteMiddleware(async () => {
  const viewBuilderStore = useViewBuilderStore();

  const { data: mounts } = await useFetch('/api/battle-net/mounts');

  if (!mounts.value) {
    return abortNavigation();
  }

  viewBuilderStore.setMounts(
    mounts.value.map((mount) => ({
      ...mount,
      category: null,
      level: null,
    }))
  );
});
