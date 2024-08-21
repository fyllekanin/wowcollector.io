<script lang="ts" setup>
const { data: page } = await useAsyncData('mounts', () =>
  queryContent('/collections/toys').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No toys page found in the content.',
  });
}

definePageMeta({
  middleware: 'collection-toys',
});
useHead({
  title: 'WoW Collector - Toys',
  meta: [
    {
      name: 'description',
      content: 'Toys collection page',
    },
  ],
});
useSeoMeta({
  title: page.value.title,
  description: page.value.description,
});

const toysStore = useToysStore();
const { toys } = storeToRefs(toysStore);

const total = computed(() => {
  if (!toys.value) return 0;
  return flatMapToys(toys.value).length;
});
const collected = computed(() => {
  if (!toys.value) return 0;
  return flatMapToys(toys.value).filter((toy) => toy.isCollected).length;
});
const percentageToysCollected = computed(() => {
  if (!toys.value) return 0;
  return Math.round((collected.value / total.value) * 100);
});
</script>

<template>
  <UContainer class="flex flex-col gap-4 pb-6">
    <UBreadcrumb :links="mapContentNavigation(page?.breadcrumbs)" />
    <CollectionHeader
      :progress="percentageToysCollected"
      :collected="collected"
      :total="total"
      collection="toys"
    >
      <ToyFilters />
    </CollectionHeader>

    <ToysGridCompact />
  </UContainer>
  <ScrollToTop />
</template>
