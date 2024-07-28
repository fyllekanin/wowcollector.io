<script lang="ts" setup>
definePageMeta({
  middleware: 'mounts',
});

const slideoverOpen = ref<boolean>(false);

const mountsStore = useMountsStore();
const { mounts, mountFilters } = storeToRefs(mountsStore);

const availableMounts = computed(() => {
  if (!mounts.value) return [];
  return flatMapMounts(mounts.value);
});
const collectedMounts = computed(() => {
  if (!mounts.value) return [];
  return flatMapMounts(mounts.value).filter((mount) => mount.isCollected);
});
const percentageMountsCollected = computed(() => {
  if (!mounts.value) return 0;
  const flattenedMounts = flatMapMounts(mounts.value);
  const collectedMounts = flattenedMounts.filter((mount) => mount.isCollected);
  return (collectedMounts.length / flattenedMounts.length) * 100;
});
</script>

<template>
  <UContainer class="flex flex-col gap-8">
    <div class="flex flex-col">
      <h2 class="text-lg font-bold self-center">Total mounts collected</h2>
      <UProgress :value="percentageMountsCollected" indicator :max="[]" />
    </div>
    <div class="flex flex-grow gap-4">
      <MountFilter v-model="slideoverOpen" />
    </div>
  </UContainer>
</template>
