import equal from 'fast-deep-equal';

import type { MountCategory, MountInformation } from '~/types';

export const useMountViewBuilderStore = defineStore('mount-view-builder', {
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
      showFaction: {
        label: 'Both',
        value: 'both',
        avatar: {
          src: 'https://cdn.discordapp.com/attachments/1161263238554599464/1279760555237838899/12d4f5a73e9c1b830c95229ac396a449.png?ex=66d59d65&is=66d44be5&hm=84d99f37f7c08df863528175f78e6853a19b5c6a48f31b1e7e0684708856139a&',
        },
      } as {
        label: string;
        value: string;
        avatar: {
          src: string;
        };
      },
    },
    successfulCreation: false,
  }),
  getters: {
    mounts(state): MountInformation[] {
      let mounts = state._mounts || [];

      if (state._settings.showFaction.value === 'both') {
        return mounts;
      }

      if (state._settings.showFaction.value === 'horde') {
        mounts = mounts.filter((mount) => mount.faction === 'HORDE');
      }

      if (state._settings.showFaction.value === 'alliance') {
        mounts = mounts.filter((mount) => mount.faction === 'ALLIANCE');
      }
      console.log(mounts.map((mount) => mount.faction));
      return mounts;
    },
    flatMounts(state): MountInformation[] {
      return state._mounts || [];
    },
    getFinalCategories(state): MountCategory[] {
      const setOrder = (category: MountCategory, order: number) => {
        category.order = order;
        category.categories.forEach((subCategory, index) => {
          setOrder(subCategory, index);
        });
        if (category.mounts)
          category.mounts.forEach((mount, index) => {
            mount.order = index;
          });
      };

      const categories = [...state._mountCategories];
      return categories.map((category, index) => {
        setOrder(category, index);
        return category;
      });
    },
    dragState(state) {
      return state._dragState;
    },
    hasChanges(state) {
      const cloned = [...state._mountCategories];
      if (cloned.length !== 1) return true;

      if (cloned.length === 1) {
        const [{ id, ...category }] = cloned;

        return !equal(category, {
          name: 'New Category',
          categories: [],
          mounts: [],
          order: 0,
        });
      }

      return false;
    },
    isValid(state) {
      return state._mountCategories.length;
    },
    highlightMountDropzones(state) {
      return state._dragState.state && state._dragState.type === 'mount';
    },
    highlightCategoryDropzones(state) {
      return state._dragState.state && state._dragState.type === 'category';
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
    resetStore() {
      this._mounts = this._allMounts;
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
    setSuccessfulCreation(successful: boolean) {
      this.successfulCreation = successful;
    },
  },
});
