import type { Character } from '~/types';

export const useCharacterStore = defineStore('character', {
  state: () => ({
    character: null as Character | null,
  }),
  actions: {
    setCharacter(newCharacter: Character) {
      this.character = newCharacter;
    },
  },
  persist: {
    paths: ['character'],
    storage: persistedState.localStorage,
  },
});
