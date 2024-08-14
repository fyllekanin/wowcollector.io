export default defineNuxtRouteMiddleware(async () => {
  const mountsStore = useMountsStore();

  const { data: mounts } = await useFetch('/api/battle-net/mounts');

  if (!mounts.value) {
    return abortNavigation();
  }

  mountsStore.setAllMounts(mounts.value);
});
