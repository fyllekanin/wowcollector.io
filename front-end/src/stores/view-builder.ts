import type { MountCategory, MountInformation } from '~/types';

function id(length: number) {
  const characters =
    'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  let result = '';
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * characters.length));
  }
  return result;
}

export const useViewBuilderStore = defineStore('view-builder', {
  state: () => ({
    _mounts: [] as MountInformation[],
    _mountCategories: [
      {
        id: id(10),
        name: 'New Category',
        categories: [],
        mounts: [],
        order: 0,
      },
    ] as MountCategory[],
  }),
  getters: {
    mounts(state) {
      return state._mounts || [];
      // // should map the mounts into a category structure
      // const mapMountsToCategories = (
      //   mounts: MountInformation[],
      //   categories: MountCategory[]
      // ): MountCategory[] => {
      //   return categories.map((category) => {
      //     return {
      //       ...category,
      //       mounts: mounts.filter((mount) => mount.category === category.id),
      //       categories: mapMountsToCategories(mounts, category.categories),
      //     };
      //   });
      // };

      // return mapMountsToCategories(state._mounts || [], state.mountCategories);
    },
    flatMounts(state) {
      return state._mounts || [];
    },
  },
  actions: {
    setMounts(newMounts: MountInformation[]) {
      this._mounts = newMounts;
    },
    setMountCategories(newCategories: MountCategory[]) {
      this._mountCategories = newCategories;
    },
    resetMountCategories() {
      this._mountCategories = [
        {
          id: id(10),
          name: 'New Category',
          categories: [],
          mounts: [],
          order: 0,
        },
      ];
    },
  },
  persist: {
    paths: ['_mounts', '_mountCategories'],
    storage: persistedState.localStorage,
  },
});
