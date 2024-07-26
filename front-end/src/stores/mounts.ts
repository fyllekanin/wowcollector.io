import type { MountCategory } from '~/types';

export const useMountsStore = defineStore(
  'mounts',
  () => {
    const _mounts = ref<MountCategory[] | null>(null);
    const mounts = computed(() => _mounts.value);

    const _mountFilters = ref<Record<string, string[] | string>>({});
    const mountFilters = computed(() => _mountFilters.value);

    const setMounts = (newMounts: MountCategory[]) => {
      _mounts.value = newMounts;
    };
    const setMountFilters = (newFilters: Record<string, string[] | string>) => {
      for (const key in newFilters) {
        _mountFilters.value[key] = newFilters[key];
      }
    };

    return {
      mounts,
      setMounts,

      mountFilters,
      setMountFilters,
    };
  },
  {
    persist: {
      paths: ['mounts', 'mountFilters'],
    },
  }
);
