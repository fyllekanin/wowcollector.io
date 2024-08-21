export default defineNuxtRouteMiddleware(async () => {
  const mountViewBuilderStore = useMountViewBuilderStore();

  const { data: mounts } = await useFetch('/api/battle-net/mounts');

  if (!mounts.value) {
    return abortNavigation();
  }

  mountViewBuilderStore.setMounts(mounts.value);
});
