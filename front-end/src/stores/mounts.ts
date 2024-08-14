import type { MountCategory, MountFilters, MountInformation } from '~/types';

export const useMountsStore = defineStore('mounts', {
  state: () => ({
    _mounts: null as MountCategory[] | null,
    _allMounts: null as MountInformation[] | null,
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
      const sortMounts = (
        a: MountInformation,
        b: MountInformation,
        sortType: string
      ) => {
        switch (sortType) {
          case 'Name Ascending':
            return a.name.localeCompare(b.name);
          case 'Name Descending':
            return b.name.localeCompare(a.name);
          case 'Collected':
            if (a.isCollected && !b.isCollected) return -1;
            if (!a.isCollected && b.isCollected) return 1;
            return 0;
          case 'Not Collected':
            if (!a.isCollected && b.isCollected) return -1;
            if (a.isCollected && !b.isCollected) return 1;
            return 0;
          default:
            return 0;
        }
      };

      const traverseSort = (category: MountCategory, sortType: string) => {
        if (category.categories) {
          category.categories = category.categories.map((cat) =>
            traverseSort(cat, sortType)
          );
          if (sortType === 'Default' && category.order) {
            category.categories.sort((a, b) => b.order - a.order);
          }
        }

        if (category.mounts) {
          category.mounts.sort((a, b) => sortMounts(a, b, sortType));
        }

        return category;
      };

      result = result
        .sort((a, b) =>
          state.filters.sort === 'Default' ? b.order - a.order : 0
        )
        .map((category) => traverseSort(category, state.filters.sort));

      // Root Categories
      if (state.filters.rootCategories.length)
        result = result.filter((category) =>
          state.filters.rootCategories.includes(category.name)
        );

      // Sub Categories
      const subCategoryFilter = (category: MountCategory) => {
        if (category.categories) {
          category.categories = category.categories.filter((subCategory) =>
            state.filters.subCategories.includes(subCategory.name)
          );
          category.categories = category.categories.map(subCategoryFilter);
        }
        return category;
      };
      if (state.filters.subCategories.length)
        result = result.map(subCategoryFilter);

      // Misc Filters
      return result;
    },
    allMounts(state) {
      return state._allMounts || [];
    },
    rootCategoryNames(state) {
      return getRootCategoryNames(state._mounts || []).sort((a, b) =>
        a.localeCompare(b)
      );
    },
    subCategoryNames(state) {
      return getSubCategoryNames(state._mounts || []).sort((a, b) =>
        a.localeCompare(b)
      );
    },
  },
  actions: {
    setMounts(newMounts: MountCategory[]) {
      this._mounts = newMounts;
    },
    setAllMounts(newMounts: MountInformation[]) {
      this._allMounts = newMounts;
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
    paths: ['_mounts'],
    storage: persistedState.sessionStorage,
  },
});
