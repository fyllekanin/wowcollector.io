<script lang="ts" setup>
const { data: page } = await useAsyncData('mounts', () =>
  queryContent('/collections/mounts').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No mounts page found in the content.',
  });
}

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
  title: page.value.title,
  description: page.value.description,
});

const mountsStore = useMountsStore();
const { mounts } = storeToRefs(mountsStore);

const total = computed(() => {
  if (!mounts.value) return 0;
  return flatMapMounts(mounts.value).length;
});
const collected = computed(() => {
  if (!mounts.value) return 0;
  return flatMapMounts(mounts.value).filter((mount) => mount.isCollected)
    .length;
});
const percentageMountsCollected = computed(() => {
  if (!mounts.value) return 0;
  return Math.round((collected.value / total.value) * 100);
});
</script>

<template>
  <UContainer class="flex flex-col gap-4">
    <CollectionHeader
      :progress="percentageMountsCollected"
      :collected="collected"
      :total="total"
      collection="mounts"
    >
      <MountFilters />
    </CollectionHeader>

    <MountsGridCompact />
  </UContainer>
  <ScrollToTop />
</template>
