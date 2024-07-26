import type { Character } from '~/types';

export const useCharacterStore = defineStore(
  'character',
  () => {
    const _character = ref<Character | null>(null);
    const character = computed(() => _character.value);

    const setCharacter = (newCharacter: Character) => {
      _character.value = newCharacter;
    };

    return {
      character,
      setCharacter,
    };
  },
  {
    persist: {
      paths: ['character'],
    },
  }
);
