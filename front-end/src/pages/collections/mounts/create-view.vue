<script lang="ts" setup>
import draggable from 'vuedraggable';

definePageMeta({
  layout: 'empty',
});

const { data: page } = await useAsyncData('mounts', () =>
  queryContent('/collections/mounts/create-view').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No mounts create view page found in the content.',
  });
}

const mountsStore = useMountsStore();

const search = ref('');
const sidebarMounts = ref(flatMapMounts(mountsStore.mounts));

const filteredMounts = computed(() => {
  if (!search.value) return sidebarMounts.value;
  return sidebarMounts.value.filter((mount) =>
    mount.name.toLowerCase().includes(search.value.toLowerCase())
  );
});

const items1 = ref([{ text: 'hello' }]);
const items2 = ref([{ text: 'world' }]);

function onMove(event: any) {
  console.log('onMove', event);
}
</script>

<template>
  <CreateViewContainer>
    <template #sidebar-content>
      <div class="flex flex-col mt-6 gap-10">
        <UCard
          draggable
          class="cursor-grab"
          :ui="{
            ring: 'hover:ring-primary dark:hover:ring-primary transition ease-in-out',
          }"
        >
          New category
        </UCard>
        <UInput
          v-model="search"
          placeholder="Search for a mount"
          icon="heroicons-outline:search"
        />
        <div class="flex grow flex-wrap w-full gap-4">
          <MountIcon
            class="select-none"
            draggable
            v-for="mount in filteredMounts"
            :key="mount?.id"
            :mount="mount"
            :clickable="false"
            build-mode
          />
        </div>
      </div>
    </template>
    <template #main-content>
      <div class="flex flex-col w-full p-10 gap-2">
        <draggable
          class="h-fit w-full p-2 border-2 border-dashed"
          v-model="items1"
          group="1"
          item-key="text"
          @end="onMove"
        >
          <template #item="{ element }">
            <div
              class="w-20 h-20 bg-red-500 rounded text-center content-center cursor-grab"
            >
              {{ element.text }}
            </div>
          </template>
        </draggable>
        <draggable
          class="h-fit w-full p-2 border-2 border-dashed"
          v-model="items2"
          group="2"
          item-key="text"
          @end="onMove"
          :data-id="1"
        >
          <template #item="{ element }">
            <div
              class="w-20 h-20 bg-red-500 rounded text-center content-center cursor-grab"
            >
              {{ element.text }}
            </div>
          </template>
        </draggable>
      </div>
    </template>
  </CreateViewContainer>
</template>
