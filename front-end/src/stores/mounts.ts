import type { MountCategory, MountFilters, MountInformation } from '~/types';

export const useMountsStore = defineStore('mounts', {
  state: () => ({
    _mounts: null as MountCategory[] | null,
    filters: {
      search: '',
      sort: 'Default',
      viewStyle: 'grid-compact',
      rootCategories: [],
      subCategories: [],
      miscFilters: [],
    } as MountFilters,
  }),
  getters: {
    mounts(state) {
      let result = state._mounts || [];

      const filterBySearch = (mounts: MountInformation[] | null) => {
        if (!mounts) return null;
        return mounts?.filter(
          (mount) =>
            mount.name
              .toLowerCase()
              .includes(state.filters.search.toLowerCase()) ||
            mount.id.toString().includes(state.filters.search)
        );
      };
      const mapSubCategories = (category: MountCategory) => {
        return {
          ...category,
          mounts: filterBySearch(category.mounts),
        };
      };
      const mapRootCategories = (category: MountCategory) => {
        return {
          ...category,
          mounts: filterBySearch(category.mounts),
          categories: category.categories
            ?.map(mapSubCategories)
            .filter((subCategory) => subCategory.mounts?.length),
        };
      };

      // Search
      result = result
        ?.map(mapRootCategories)
        .filter(
          (category) => category.mounts?.length || category.categories?.length
        );

      // Sort
      // Root Categories
      // Sub Categories
      // Misc Filters
      return result;
    },
  },
  actions: {
    setMounts(newMounts: MountCategory[]) {
      this._mounts = newMounts;
    },
    clearMounts() {
      this._mounts = null;
    },
    setMountFilters(newFilters: Partial<MountFilters>) {
      for (const key in newFilters) {
        // @ts-expect-error
        this.filters[key] = newFilters[key as keyof MountFilters];
      }
    },
    clearMountFilters() {
      this.filters = {
        search: '',
        sort: 'Default',
        viewStyle: 'grid-compact',
        rootCategories: [],
        subCategories: [],
        miscFilters: [],
      };
    },
    clearMountFilter(filter: keyof MountFilters) {
      delete this.filters[filter];
    },
  },
  persist: {
    paths: ['_mounts', 'filters'],
    storage: persistedState.sessionStorage,
  },
});
