import type {
  AchievementCategory,
  AchievementFilters,
  AchievementInformation,
} from '~/types';

export const useAchievementsStore = defineStore('achievements', {
  state: () => ({
    _achievements: null as AchievementCategory[] | null,
    filters: {
      search: '',
      sort: 'Default',
      rootCategories: [],
      subCategories: [],
      miscFilters: [],
    } as AchievementFilters,
  }),
  getters: {
    achievements(state) {
      let result = state._achievements || [];

      const filterBySearch = (
        achievements: AchievementInformation[] | null
      ) => {
        if (!achievements) return null;
        return achievements?.filter(
          (achievement) =>
            achievement.name
              .toLowerCase()
              .includes(state.filters.search.toLowerCase()) ||
            achievement.id.toString().includes(state.filters.search)
        );
      };
      const mapSubCategories = (category: AchievementCategory) => {
        if (!category.hasOwnProperty('achievements')) return category;
        return {
          ...category,
          achievements: filterBySearch(category?.achievements || []),
        };
      };
      const mapRootCategories = (category: AchievementCategory) => {
        if (!category.hasOwnProperty('achievements')) return category;
        return {
          ...category,
          achievements: filterBySearch(category?.achievements || []),
          categories: category.categories
            ?.map(mapSubCategories)
            .filter((subCategory) => subCategory.achievements?.length),
        };
      };

      // Search
      result = result?.map(mapRootCategories).filter((category) => {
        if (!('achievements' in category)) return true;
        return category.achievements?.length || category.categories?.length;
      });

      // Sort
      const sortAchievements = (
        a: AchievementInformation,
        b: AchievementInformation,
        sortType: string
      ) => {
        switch (sortType) {
          case 'Name Ascending':
            return a.name.localeCompare(b.name);
          case 'Name Descending':
            return b.name.localeCompare(a.name);
          case 'Collected':
            if (a.isCompleted && !b.isCompleted) return -1;
            if (!a.isCompleted && b.isCompleted) return 1;
            return 0;
          case 'Not Collected':
            if (!a.isCompleted && b.isCompleted) return -1;
            if (a.isCompleted && !b.isCompleted) return 1;
            return 0;
          default:
            return 0;
        }
      };

      const traverseSort = (
        category: AchievementCategory,
        sortType: string
      ) => {
        if (category.categories) {
          category.categories = category.categories.map((cat) =>
            traverseSort(cat, sortType)
          );
          if (sortType === 'Default' && category.displayOrder) {
            category.categories.sort((a, b) => b.displayOrder - a.displayOrder);
          }
        }

        if (category.achievements) {
          category.achievements.sort((a, b) =>
            sortAchievements(a, b, sortType)
          );
        }

        return category;
      };

      result = result
        .sort((a, b) =>
          state.filters.sort === 'Default' ? b.displayOrder - a.displayOrder : 0
        )
        .map((category) => traverseSort(category, state.filters.sort));

      // Root Categories
      if (state.filters.rootCategories.length)
        result = result.filter((category) =>
          state.filters.rootCategories.includes(category.name)
        );

      // Sub Categories
      const subCategoryFilter = (category: AchievementCategory) => {
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
      return getRootCategoryNames(state._achievements || []).sort((a, b) =>
        a.localeCompare(b)
      );
    },
    subCategoryNames(state) {
      return getSubCategoryNames(state._achievements || []).sort((a, b) =>
        a.localeCompare(b)
      );
    },
  },
  actions: {
    setAchievement(newAchievementCategory: AchievementCategory) {
      const rootCategoryIndex = this._achievements?.findIndex(
        (category) => category.id === newAchievementCategory.id
      );
      console.log(this._achievements);
      if (
        rootCategoryIndex !== undefined &&
        rootCategoryIndex !== -1 &&
        this._achievements
      ) {
        this._achievements[rootCategoryIndex] = newAchievementCategory;
      }
      console.log(this._achievements);
    },
    setAchievements(newAchievements: AchievementCategory[]) {
      this._achievements = newAchievements;
    },
    clearAchievements() {
      this._achievements = null;
    },
    setAchievementFilters(newFilters: Partial<AchievementFilters>) {
      for (const key in newFilters) {
        // @ts-expect-error
        this.filters[key] = newFilters[key as keyof AchievementFilters];
      }
    },
    clearAchievementFilters() {
      this.filters = {
        search: '',
        sort: 'Default',
        rootCategories: [],
        subCategories: [],
        miscFilters: [],
      };
    },
    clearAchievementFilter(filter: keyof AchievementFilters) {
      delete this.filters[filter];
    },
  },
  persist: {
    paths: ['_achievements'],
    storage: persistedState.sessionStorage,
  },
});
