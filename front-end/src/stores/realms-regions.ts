import type { Realm, Region } from '~/types';

export const useRealmsRegionsStore = defineStore({
  id: 'realms-regions',
  state: () => ({
    realms: null as Realm[] | null,
    regions: null as Region[] | null,
  }),
  actions: {
    setRealms(newRealms: Realm[]) {
      this.realms = newRealms;
    },
    setRegions(newRegions: Region[]) {
      this.regions = newRegions;
    },
  },
  persist: {
    pick: ['realms', 'regions'],
    storage: persistedState.localStorage,
  },
});
