import equal from 'fast-deep-equal';

import type { ToyCategory, ToyInformation } from '~/types';

export const useToyViewBuilderStore = defineStore('toy-view-builder', {
  state: () => ({
    _allToys: [] as ToyInformation[], // Used as a fallback to reset the state
    _toys: [] as ToyInformation[],
    _toyCategories: [
      {
        id: newId(10),
        name: 'New Category',
        categories: [],
        toys: [],
        order: 0,
      },
    ] as ToyCategory[],
    _cloneableCategory: [
      {
        id: newId(10),
        name: 'New Category',
        categories: [],
        toys: [],
        order: 0,
      },
    ] as ToyCategory[],
    _dragState: {
      state: false,
      type: '' as 'toy' | 'category' | '',
    },
    _searchFilter: '',
    _settings: {
      showBorders: true,
      showToyTooltips: true,
    },
  }),
  getters: {
    flatToys(state): ToyInformation[] {
      return state._toys || [];
    },
    getFinalCategories(state): ToyCategory[] {
      const setOrder = (category: ToyCategory, order: number) => {
        category.order = order;
        category.categories.forEach((subCategory, index) => {
          setOrder(subCategory, index);
        });
        if (category.toys)
          category.toys.forEach((toy, index) => {
            toy.order = index;
          });
      };

      const categories = [...state._toyCategories];
      return categories.map((category, index) => {
        setOrder(category, index);
        return category;
      });
    },
    dragState(state) {
      return state._dragState;
    },
    hasChanges(state) {
      const cloned = [...state._toyCategories];
      if (cloned.length !== 1) return true;

      if (cloned.length === 1) {
        const [{ id, ...category }] = cloned;

        return !equal(category, {
          name: 'New Category',
          categories: [],
          toys: [],
          order: 0,
        });
      }

      return false;
    },
    isValid(state) {
      return state._toyCategories.length;
    },
  },
  actions: {
    setToys(newToys: ToyInformation[]) {
      this._toys = newToys;
    },
    setToyCategories(newCategories: ToyCategory[]) {
      this._toyCategories = newCategories;
    },
    addToys(toys: ToyInformation[]) {
      this._toys.push(...toys);
    },
    addRootCategory(category: ToyCategory) {
      this._toyCategories.push(category);
    },
    setNewIdForCloneableCategory() {
      this._cloneableCategory = [
        {
          id: newId(10),
          name: 'New Category',
          categories: [],
          toys: [],
          order: 0,
        },
      ];
    },
    addSubCategory(category: ToyCategory, parentId: string) {
      const parentCategory = this._toyCategories.find(
        (category) => category.id === parentId
      );

      if (parentCategory) {
        parentCategory.categories.push(category);
      }
    },
    removeRootCategory(categoryId: string) {
      this._toyCategories = this._toyCategories.filter(
        (category) => category.id !== categoryId
      );
    },
    removeSubCategory(categoryId: string, parentId: string) {
      const parentCategory = this._toyCategories.find(
        (category) => category.id === parentId
      );

      if (parentCategory) {
        parentCategory.categories = parentCategory.categories.filter(
          (category) => category.id !== categoryId
        );
      }
    },
    resetStore() {
      this._toys = this._allToys;
      this._toyCategories = [
        {
          id: newId(10),
          name: 'New Category',
          categories: [],
          toys: [],
          order: 0,
        },
      ];
    },
    setDragState(state: boolean, type: 'toy' | 'category') {
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
});
