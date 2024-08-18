import equal from 'fast-deep-equal';

import type { AchievementCategory, AchievementInformation } from '~/types';

export const useAchievementViewBuilderStore = defineStore(
  'achievement-view-builder',
  {
    state: () => ({
      _allAchievements: [] as AchievementInformation[], // Used as a fallback to reset the state
      _achievements: [] as AchievementInformation[],
      _achievementCategories: [
        {
          id: newId(10),
          name: 'New Category',
          categories: [],
          achievements: [],
          displayOrder: 0,
        },
      ] as AchievementCategory[],
      _cloneableCategory: [
        {
          id: newId(10),
          name: 'New Category',
          categories: [],
          achievements: [],
          displayOrder: 0,
        },
      ] as AchievementCategory[],
      _dragState: {
        state: false,
        type: '' as 'achievement' | 'category' | '',
      },
      _searchFilter: '',
      _settings: {
        showBorders: true,
        showMountTooltips: true,
      },
    }),
    getters: {
      flatAchievements(state): AchievementInformation[] {
        return state._achievements || [];
      },
      dragState(state) {
        return state._dragState;
      },
      hasChanges(state) {
        const cloned = [...state._achievementCategories];
        if (cloned.length !== 1) return true;

        if (cloned.length === 1) {
          const [{ id, ...category }] = cloned;

          return !equal(category, {
            name: 'New Category',
            categories: [],
            achievements: [],
            displayOrder: 0,
          });
        }

        return false;
      },
      isValid(state) {
        return state._achievementCategories.length;
      },
    },
    actions: {
      setAchievements(newAchievements: AchievementInformation[]) {
        this._achievements = newAchievements;
      },
      setAchievementCategories(newCategories: AchievementCategory[]) {
        this._achievementCategories = newCategories;
      },
      addAchievements(achievements: AchievementInformation[]) {
        this._achievements.push(...achievements);
      },
      addRootCategory(category: AchievementCategory) {
        this._achievementCategories.push(category);
      },
      setNewIdForCloneableCategory() {
        this._cloneableCategory = [
          {
            id: newId(10),
            name: 'New Category',
            categories: [],
            achievements: [],
            displayOrder: 0,
          },
        ];
      },
      addSubCategory(category: AchievementCategory, parentId: string) {
        const parentCategory = this._achievementCategories.find(
          (category) => category.id === parentId
        );

        if (parentCategory) {
          parentCategory.categories.push(category);
        }
      },
      removeRootCategory(categoryId: string) {
        this._achievementCategories = this._achievementCategories.filter(
          (category) => category.id !== categoryId
        );
      },
      removeSubCategory(categoryId: string, parentId: string) {
        const parentCategory = this._achievementCategories.find(
          (category) => category.id === parentId
        );

        if (parentCategory) {
          parentCategory.categories = parentCategory.categories.filter(
            (category) => category.id !== categoryId
          );
        }
      },
      resetStore() {
        this._achievements = this._allAchievements;
        this._achievementCategories = [
          {
            id: newId(10),
            name: 'New Category',
            categories: [],
            achievements: [],
            displayOrder: 0,
          },
        ];
      },
      setDragState(state: boolean, type: 'achievement' | 'category') {
        this._dragState = {
          state,
          type,
        };
      },
      clearDragState() {
        this._dragState = {
          state: false,
          type: '',
        };
      },
      setSearchFilter(filter: string) {
        this._searchFilter = filter;
      },
      setSettings(settings: Partial<typeof this._settings>) {
        Object.assign(this._settings, settings);
      },
    },
  }
);
