import type { Character } from '~/types';

export const useCharacterStore = defineStore({
  id: 'character',
  state: () => ({
    character: null as Partial<Character> | null,
  }),
  actions: {
    setCharacter(newCharacter: Partial<Character>) {
      this.character = newCharacter;
    },
    clearCharacter() {
      this.character = null;
    },
  },
  persist: {
    pick: ['character'],
    storage: persistedState.sessionStorage,
  },
});
