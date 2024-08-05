<script lang="ts" setup>
const { data: page } = await useAsyncData('pets', () =>
  queryContent('/collections/pets').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No pets page found in the content.',
  });
}

definePageMeta({
  middleware: 'pets',
});
useHead({
  title: 'WoW Collector - Pets',
  meta: [
    {
      name: 'description',
      content: 'Pets collection page',
    },
  ],
});
useSeoMeta({
  title: page.value.title,
  description: page.value.description,
});

const petsStore = usePetsStore();
const { pets } = storeToRefs(petsStore);

const total = computed(() => {
  if (!pets.value) return 0;
  return flatMapPets(pets.value).length;
});
const collected = computed(() => {
  if (!pets.value) return 0;
  return flatMapPets(pets.value).filter((pet) => pet.isCollected).length;
});
const percentagePetsCollected = computed(() => {
  if (!pets.value) return 0;
  return Math.round((collected.value / total.value) * 100);
});
</script>

<template>
  <UContainer class="flex flex-col gap-4 pb-6">
    <UBreadcrumb :links="mapContentNavigation(page?.breadcrumbs)" />
    <CollectionHeader
      :progress="percentagePetsCollected"
      :collected="collected"
      :total="total"
      collection="pets"
    >
      <PetFilters />
    </CollectionHeader>

    <PetsGridCompact />
  </UContainer>
  <ScrollToTop />
</template>
