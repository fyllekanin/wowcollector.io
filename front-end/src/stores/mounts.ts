import type { MountCategory } from '~/types';

export const useMountsStore = defineStore('mounts', {
  state: () => ({
    mounts: null as MountCategory[] | null,
    mountFilters: {
      renderType: 'grid-compact',
    } as Record<string, string[] | string>,
  }),
  actions: {
    setMounts(newMounts: MountCategory[]) {
      this.mounts = newMounts;
    },
    clearMounts() {
      this.mounts = null;
    },
    setMountFilters(newFilters: Record<string, string[] | string>) {
      for (const key in newFilters) {
        this.mountFilters[key] = newFilters[key];
      }
    },
    clearMountFilters() {
      this.mountFilters = {};
    },
    clearMountFilter(filter: string) {
      delete this.mountFilters[filter];
    },
  },
  persist: {
    paths: ['mounts', 'mountFilters'],
    storage: persistedState.sessionStorage,
  },
});
