import type { MountCategory } from '~/types';

export const useMountsStore = defineStore('mounts', {
  state: () => ({
    mounts: null as MountCategory[] | null,
    mountFilters: {} as Record<string, string[] | string>,
  }),
  actions: {
    setMounts(newMounts: MountCategory[]) {
      this.mounts = newMounts;
    },
    setMountFilters(newFilters: Record<string, string[] | string>) {
      for (const key in newFilters) {
        this.mountFilters[key] = newFilters[key];
      }
    },
  },
  persist: {
    paths: ['mounts', 'mountFilters'],
    storage: persistedState.localStorage,
  },
});
