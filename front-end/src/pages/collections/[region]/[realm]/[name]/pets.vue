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
  middleware: 'collection-pets',
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

const toast = useToast();

const characterStore = useCharacterStore();
const { character } = storeToRefs(characterStore);
const petsStore = usePetsStore();
const { pets } = storeToRefs(petsStore);

const viewId = ref('');
const loading = ref(false);

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

async function loadView() {
  try {
    loading.value = true;

    if (!viewId.value)
      return navigateTo(
        `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/pets`
      );

    const failed = await navigateTo(
      `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/pets?viewId=${viewId.value}`
    );
    if (failed) toast.add({ title: 'View not found', color: 'red' });
  } catch (error) {
    console.error(error);
    toast.add({
      title: 'Failed to load view',
      description: 'An error occurred while loading the view.',
      color: 'red',
    });
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <UContainer class="flex flex-col gap-4 pb-6">
    <div class="flex grow flex-wrap gap-2 items-center justify-between">
      <UBreadcrumb :links="mapContentNavigation(page?.breadcrumbs)" />
      <UButtonGroup>
        <UInput
          v-model="viewId"
          readonly
          icon="heroicons-outline:link"
          placeholder="Load view by ID"
        />
        <UTooltip text="Load view">
          <UButton
            @click="loadView"
            icon="material-symbols:send-rounded"
            color="gray"
          />
        </UTooltip>
      </UButtonGroup>
    </div>
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
