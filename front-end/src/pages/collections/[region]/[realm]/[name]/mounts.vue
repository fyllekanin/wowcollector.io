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
  middleware: 'collection-mounts',
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
  ogTitle: page.value.og.title,
  ogDescription: page.value.og.description,
  ogImage: page.value.og.image,
});

const toast = useToast();

const characterStore = useCharacterStore();
const { character } = storeToRefs(characterStore);
const mountsStore = useMountsStore();
const { mounts } = storeToRefs(mountsStore);

const viewId = ref('');
const loading = ref(false);

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

async function loadView() {
  try {
    loading.value = true;

    if (!viewId.value)
      return navigateTo(
        `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/mounts`
      );

    const failed = await navigateTo(
      `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/mounts?viewId=${viewId.value}`
    );
    // @ts-expect-error - NavigationFailure is not defined in the type definitions
    if (failed?.message?.includes('redundant navigation')) return;
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
          icon="heroicons-outline:eye"
          placeholder="Load view by ID"
        />
        <UTooltip text="Load view">
          <UButton
            @click="loadView"
            icon="material-symbols:send-rounded"
            color="gray"
            :loading="loading"
          />
        </UTooltip>
      </UButtonGroup>
    </div>
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
