import type { ToyCategory, ToyFilters, ToyInformation } from '~/types';

export const useToysStore = defineStore({
  id: 'toys',
  state: () => ({
    _toys: null as ToyCategory[] | null,
    filters: {
      search: '',
      sort: 'Default',
      rootCategories: [],
      subCategories: [],
      miscFilters: [],
    } as ToyFilters,
  }),
  getters: {
    toys(state) {
      let result = state._toys || [];

      const filterBySearch = (toys: ToyInformation[] | null) => {
        if (!toys) return null;
        return toys?.filter(
          (mount) =>
            mount.name
              .toLowerCase()
              .includes(state.filters.search.toLowerCase()) ||
            mount.id.toString().includes(state.filters.search)
        );
      };
      const mapSubCategories = (category: ToyCategory) => {
        return {
          ...category,
          toys: filterBySearch(category.toys),
        };
      };
      const mapRootCategories = (category: ToyCategory) => {
        return {
          ...category,
          toys: filterBySearch(category.toys),
          categories: category.categories
            ?.map(mapSubCategories)
            .filter((subCategory) => subCategory.toys?.length),
        };
      };

      // Search
      result = result
        ?.map(mapRootCategories)
        .filter(
          (category) => category.toys?.length || category.categories?.length
        );

      // Sort
      const sortToys = (
        a: ToyInformation,
        b: ToyInformation,
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

      const traverseSort = (category: ToyCategory, sortType: string) => {
        if (category.categories) {
          category.categories = category.categories.map((cat) =>
            traverseSort(cat, sortType)
          );
          if (sortType === 'Default' && category.order) {
            category.categories.sort((a, b) => b.order - a.order);
          }
        }

        if (category.toys) {
          category.toys.sort((a, b) => sortToys(a, b, sortType));
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
      const subCategoryFilter = (category: ToyCategory) => {
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
    rootCategoryNames(state) {
      return getRootCategoryNames(state._toys || []).sort((a, b) =>
        a.localeCompare(b)
      );
    },
    subCategoryNames(state) {
      return getSubCategoryNames(state._toys || []).sort((a, b) =>
        a.localeCompare(b)
      );
    },
  },
  actions: {
    setToys(newToys: ToyCategory[]) {
      this._toys = newToys;
    },
    clearToys() {
      this._toys = null;
    },
    setToyFilters(newFilters: Partial<ToyFilters>) {
      for (const key in newFilters) {
        // @ts-expect-error
        this.filters[key] = newFilters[key as keyof PetFilters];
      }
    },
    clearToyFilters() {
      this.filters = {
        search: '',
        sort: 'Default',
        rootCategories: [],
        subCategories: [],
        miscFilters: [],
      };
    },
    clearToyFilter(filter: keyof ToyFilters) {
      delete this.filters[filter];
    },
  },
  persist: {
    pick: ['_toys'],
    storage: persistedState.sessionStorage,
  },
});
