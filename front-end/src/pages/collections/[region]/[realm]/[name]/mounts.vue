<script lang="ts" setup>
definePageMeta({
  middleware: 'mounts',
});
useHead({
  title: 'WoW Collector - Mounts',
  meta: [
    {
      name: 'description',
      content: 'Mounts collection page',
    },
  ],
});
useSeoMeta({
  title: 'WoW Collector - Mounts',
  description: 'Mounts collection page',
});

const mountsStore = useMountsStore();
const { mounts } = storeToRefs(mountsStore);

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
  return Math.round((collectedMounts.length / flattenedMounts.length) * 100);
});
</script>

<template>
  <UContainer class="flex flex-col gap-4">
    <CollectionHeader
      :progress="percentageMountsCollected"
      :collected="collectedMounts"
      :available="availableMounts"
      collection="mounts"
    >
      <MountFilters />
    </CollectionHeader>

    <MountsGridCompact />
  </UContainer>
  <ScrollToTop />
</template>
