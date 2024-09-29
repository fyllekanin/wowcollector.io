import equal from 'fast-deep-equal';

import { Images } from '~/constants';

import type { PetCategory, PetInformation } from '~/types';

export const usePetViewBuilderStore = defineStore({
  id: 'pet-view-builder',
  state: () => ({
    _allPets: [] as PetInformation[], // Used as a fallback to reset the state
    _pets: [] as PetInformation[],
    _petCategories: [
      {
        id: newId(10),
        name: 'New Category',
        categories: [],
        pets: [],
        order: 0,
      },
    ] as PetCategory[],
    _cloneableCategory: [
      {
        id: newId(10),
        name: 'New Category',
        categories: [],
        pets: [],
        order: 0,
      },
    ] as PetCategory[],
    _dragState: {
      state: false,
      type: '' as 'pet' | 'category' | '',
    },
    _searchFilter: '',
    _settings: {
      showBorders: true,
      showPetTooltips: true,
      showFaction: {
        label: 'Both',
        value: 'both',
        avatar: {
          src: Images.BOTH,
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
    pets(state): PetInformation[] {
      let pets = state._pets || [];

      if (state._settings.showFaction.value === 'both') {
        return pets;
      }

      if (state._settings.showFaction.value === 'horde') {
        pets = pets.filter((pet) => pet.faction === 'HORDE');
      }

      if (state._settings.showFaction.value === 'alliance') {
        pets = pets.filter((pet) => pet.faction === 'ALLIANCE');
      }
      console.log(pets.map((pet) => pet.faction));
      return pets;
    },
    flatPets(state): PetInformation[] {
      return state._pets || [];
    },
    getFinalCategories(state): PetCategory[] {
      const setOrder = (category: PetCategory, order: number) => {
        category.order = order;
        category.categories.forEach((subCategory, index) => {
          setOrder(subCategory, index);
        });
        if (category.pets)
          category.pets.forEach((pet, index) => {
            pet.order = index;
          });
      };

      const categories = [...state._petCategories];
      return categories.map((category, index) => {
        setOrder(category, index);
        return category;
      });
    },
    dragState(state) {
      return state._dragState;
    },
    hasChanges(state) {
      const cloned = [...state._petCategories];
      if (cloned.length !== 1) return true;

      if (cloned.length === 1) {
        const [{ id, ...category }] = cloned;

        return !equal(category, {
          name: 'New Category',
          categories: [],
          pets: [],
          order: 0,
        });
      }

      return false;
    },
    isValid(state) {
      return state._petCategories.length;
    },
    highlightPetDropzones(state) {
      return state._dragState.state && state._dragState.type === 'pet';
    },
    highlightCategoryDropzones(state) {
      return state._dragState.state && state._dragState.type === 'category';
    },
  },
  actions: {
    setPets(newPets: PetInformation[]) {
      this._pets = newPets;
    },
    setPetCategories(newCategories: PetCategory[]) {
      this._petCategories = newCategories;
    },
    addPets(pets: PetInformation[]) {
      this._pets.push(...pets);
    },
    addRootCategory(category: PetCategory) {
      this._petCategories.push(category);
    },
    setNewIdForCloneableCategory() {
      this._cloneableCategory = [
        {
          id: newId(10),
          name: 'New Category',
          categories: [],
          pets: [],
          order: 0,
        },
      ];
    },
    addSubCategory(category: PetCategory, parentId: string) {
      const parentCategory = this._petCategories.find(
        (category) => category.id === parentId
      );

      if (parentCategory) {
        parentCategory.categories.push(category);
      }
    },
    removeRootCategory(categoryId: string) {
      this._petCategories = this._petCategories.filter(
        (category) => category.id !== categoryId
      );
    },
    removeSubCategory(categoryId: string, parentId: string) {
      const parentCategory = this._petCategories.find(
        (category) => category.id === parentId
      );

      if (parentCategory) {
        parentCategory.categories = parentCategory.categories.filter(
          (category) => category.id !== categoryId
        );
      }
    },
    resetStore() {
      this._pets = this._allPets;
      this._petCategories = [
        {
          id: newId(10),
          name: 'New Category',
          categories: [],
          pets: [],
          order: 0,
        },
      ];
    },
    setDragState(state: boolean, type: 'pet' | 'category') {
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
