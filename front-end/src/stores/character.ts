import type { Character } from '~/types';

export const useCharacterStore = defineStore('character', {
  state: () => ({
    character: null as Partial<Character> | null,
  }),
  actions: {
    setCharacter(newCharacter: Partial<Character>) {
      console.log('newCharacter', newCharacter);
      this.character = newCharacter;
    },
    clearCharacter() {
      this.character = null;
    },
  },
  persist: {
    paths: ['character'],
    storage: persistedState.sessionStorage,
  },
});
