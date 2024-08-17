import type { MountCategory, MountInformation } from '~/types';

export const useViewBuilderStore = defineStore('view-builder', {
  state: () => ({
    _allMounts: [] as MountInformation[], // Used as a fallback to reset the state
    _mounts: [] as MountInformation[],
    _mountCategories: [
      {
        id: newId(10),
        name: 'New Category',
        categories: [],
        mounts: [],
        order: 0,
      },
    ] as MountCategory[],
    _cloneableCategory: [
      {
        id: newId(10),
        name: 'New Category',
        categories: [],
        mounts: [],
        order: 0,
      },
    ] as MountCategory[],
    _dragState: {
      state: false,
      type: '' as 'mount' | 'category' | '',
    },
    _searchFilter: '',
    _settings: {
      showBorders: true,
      showMountTooltips: true,
    },
  }),
  getters: {
    flatMounts(state): MountInformation[] {
      return state._mounts || [];
    },
    dragState(state) {
      return state._dragState;
    },
  },
  actions: {
    setMounts(newMounts: MountInformation[]) {
      this._mounts = newMounts;
    },
    setMountCategories(newCategories: MountCategory[]) {
      this._mountCategories = newCategories;
    },
    addMounts(mounts: MountInformation[]) {
      this._mounts.push(...mounts);
    },
    addRootCategory(category: MountCategory) {
      this._mountCategories.push(category);
    },
    setNewIdForCloneableCategory() {
      this._cloneableCategory = [
        {
          id: newId(10),
          name: 'New Category',
          categories: [],
          mounts: [],
          order: 0,
        },
      ];
    },
    addSubCategory(category: MountCategory, parentId: string) {
      const parentCategory = this._mountCategories.find(
        (category) => category.id === parentId
      );

      if (parentCategory) {
        parentCategory.categories.push(category);
      }
    },
    removeRootCategory(categoryId: string) {
      this._mountCategories = this._mountCategories.filter(
        (category) => category.id !== categoryId
      );
    },
    removeSubCategory(categoryId: string, parentId: string) {
      const parentCategory = this._mountCategories.find(
        (category) => category.id === parentId
      );

      if (parentCategory) {
        parentCategory.categories = parentCategory.categories.filter(
          (category) => category.id !== categoryId
        );
      }
    },
    resetMountCategories() {
      this._mountCategories = [
        {
          id: newId(10),
          name: 'New Category',
          categories: [],
          mounts: [],
          order: 0,
        },
      ];
    },
    setDragState(state: boolean, type: 'mount' | 'category') {
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
  // persist: {
  //   paths: ['_mounts', '_mountCategories'],
  //   storage: persistedState.localStorage,
  // },
});
