import type { PetCategory, PetFilters, PetInformation } from '~/types';

export const usePetsStore = defineStore({
  id: 'pets',
  state: () => ({
    _pets: null as PetCategory[] | null,
    filters: {
      search: '',
      sort: 'Default',
      rootCategories: [],
      subCategories: [],
      miscFilters: [],
    } as PetFilters,
  }),
  getters: {
    pets(state) {
      let result = state._pets || [];

      const filterBySearch = (toys: PetInformation[] | null) => {
        if (!toys) return null;
        return toys?.filter(
          (mount) =>
            mount.name
              .toLowerCase()
              .includes(state.filters.search.toLowerCase()) ||
            mount.id.toString().includes(state.filters.search)
        );
      };
      const mapSubCategories = (category: PetCategory) => {
        return {
          ...category,
          toys: filterBySearch(category.pets),
        };
      };
      const mapRootCategories = (category: PetCategory) => {
        return {
          ...category,
          pets: filterBySearch(category.pets),
          categories: category.categories
            ?.map(mapSubCategories)
            .filter((subCategory) => subCategory.pets?.length),
        };
      };

      // Search
      result = result
        ?.map(mapRootCategories)
        .filter(
          (category) => category.pets?.length || category.categories?.length
        );

      // Sort
      const sortPets = (
        a: PetInformation,
        b: PetInformation,
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

      const traverseSort = (category: PetCategory, sortType: string) => {
        if (category.categories) {
          category.categories = category.categories.map((cat) =>
            traverseSort(cat, sortType)
          );
          if (sortType === 'Default' && category.order) {
            category.categories.sort((a, b) => b.order - a.order);
          }
        }

        if (category.pets) {
          category.pets.sort((a, b) => sortPets(a, b, sortType));
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
      const subCategoryFilter = (category: PetCategory) => {
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
      return getRootCategoryNames(state._pets || []).sort((a, b) =>
        a.localeCompare(b)
      );
    },
    subCategoryNames(state) {
      return getSubCategoryNames(state._pets || []).sort((a, b) =>
        a.localeCompare(b)
      );
    },
  },
  actions: {
    setPets(newPets: PetCategory[]) {
      this._pets = newPets;
    },
    clearPets() {
      this._pets = null;
    },
    setPetFilters(newFilters: Partial<PetFilters>) {
      for (const key in newFilters) {
        // @ts-expect-error
        this.filters[key] = newFilters[key as keyof PetFilters];
      }
    },
    clearPetFilters() {
      this.filters = {
        search: '',
        sort: 'Default',
        rootCategories: [],
        subCategories: [],
        miscFilters: [],
      };
    },
    clearPetFilter(filter: keyof PetFilters) {
      delete this.filters[filter];
    },
  },
  persist: {
    pick: ['_pets'],
    storage: persistedState.sessionStorage,
  },
});
