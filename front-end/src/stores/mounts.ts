import type { MountCategory, MountFilters } from '~/types';

export const useMountsStore = defineStore('mounts', {
  state: () => ({
    mounts: null as MountCategory[] | null,
    mountFilters: {
      search: '',
      sort: 'Default',
      viewStyle: 'grid-compact',
      rootCategories: [],
      subCategories: [],
      miscFilters: [],
    } as MountFilters,
  }),
  actions: {
    setMounts(newMounts: MountCategory[]) {
      this.mounts = newMounts;
    },
    clearMounts() {
      this.mounts = null;
    },
    setMountFilters(newFilters: Partial<MountFilters>) {
      for (const key in newFilters) {
        // @ts-expect-error
        this.mountFilters[key] = newFilters[key as keyof MountFilters];
      }
    },
    clearMountFilters() {
      this.mountFilters = {
        search: '',
        sort: 'Default',
        viewStyle: 'grid-compact',
        rootCategories: [],
        subCategories: [],
        miscFilters: [],
      };
    },
    clearMountFilter(filter: keyof MountFilters) {
      delete this.mountFilters[filter];
    },
  },
  persist: {
    paths: ['mounts', 'mountFilters'],
    storage: persistedState.sessionStorage,
  },
});
