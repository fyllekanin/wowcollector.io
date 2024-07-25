export const useCharacterStore = defineStore(
  'character',
  () => {
    const character = ref<{} | null>(null);

    return {
      character,
    };
  },
  {
    persist: {
      paths: ['character'],
    },
  }
);
