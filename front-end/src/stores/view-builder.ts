import type { BuilderMountCategory, BuilderMountInformation } from '~/types';

export const useViewBuilderStore = defineStore('view-builder', {
  state: () => ({
    _mounts: [] as BuilderMountInformation[],
    _mountCategories: [
      {
        id: newId(10),
        name: 'New Category',
        categories: [],
        mounts: [],
        order: 0,
      },
    ] as BuilderMountCategory[],
  }),
  getters: {
    mounts(state) {
      // return state._mounts || [];
      // // should map the mounts into a category structure
      const mapMountsToCategories = (
        mounts: BuilderMountInformation[],
        categories: BuilderMountCategory[]
      ): BuilderMountCategory[] => {
        return categories.map((category) => {
          return {
            ...category,
            mounts: mounts.filter((mount) => mount.category === category.id),
            categories: mapMountsToCategories(mounts, category.categories),
          };
        });
      };

      return mapMountsToCategories(state._mounts || [], state._mountCategories);
    },
    flatMounts(state) {
      return state._mounts || [];
    },
  },
  actions: {
    setMounts(newMounts: BuilderMountInformation[]) {
      this._mounts = newMounts;
    },
    setMountCategories(newCategories: BuilderMountCategory[]) {
      this._mountCategories = newCategories;
    },
    addRootCategory(category: BuilderMountCategory) {
      this._mountCategories.push(category);
    },
    updateRootOrder(idOrder: Array<string>): void {

    },
    updateOrder(category: BuilderMountCategory, idOrder: Array<string>): void {

    },
    addSubCategory(category: BuilderMountCategory, parentId: string) {
      const parentCategory = this._mountCategories.find(
        (category) => category.id === parentId
      );

      if (parentCategory) {
        parentCategory.categories.push(category);
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
  },
  // persist: {
  //   paths: ['_mounts', '_mountCategories'],
  //   storage: persistedState.localStorage,
  // },
});
