<script lang="ts" setup>
definePageMeta({
  middleware: 'mounts',
});

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
  return Math.round((collectedMounts.length / flattenedMounts.length) * 100);
});
</script>

<template>
  <UContainer class="flex flex-col gap-4">
    <div class="flex grow gap-8 items-center">
      <div class="flex flex-col w-full">
        <UProgress :value="percentageMountsCollected" />
        <p
          class="text-center text-xs sm:text-sm text-nowrap text-gray-500 self-end pt-1"
        >
          {{ collectedMounts.length }} out of
          {{ availableMounts.length }} mounts collected ({{
            percentageMountsCollected
          }}%)
        </p>
      </div>
      <MountFilterSlideover />
    </div>
    <UDivider />
    <div class="flex grow flex-wrap items-end gap-5">
      <MountFilterRow class="hidden sm:flex" />
    </div>
  </UContainer>
</template>
