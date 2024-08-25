<script lang="ts" setup>
defineEmits(['cancel', 'confirm']);

const loading = ref(false);

const state = reactive({
  name: '',
  isUnknownIncluded: false,
});
</script>

<template>
  <UModal>
    <UCard
      :ui="{
        ring: '',
        divide: 'divide-y divide-gray-100 dark:divide-gray-800',
      }"
    >
      <template #header>
        <h2 class="text-lg font-semibold">Create View</h2>
      </template>

      <div class="flex flex-col gap-10">
        <div class="flex flex-col gap-2">
          <p class="text-sm text-gray-600 dark:text-gray-400">
            Provide a name for your new view
          </p>
          <UInput
            v-model="state.name"
            required
            label="View Name"
            placeholder="My Awesome View"
            size="lg"
          />
        </div>

        <div class="flex gap-2">
          <UToggle v-model="state.isUnknownIncluded" />
          <p class="text-sm text-gray-600 dark:text-gray-400">
            Include any unselected mounts inside an "Unknown" category
          </p>
        </div>
      </div>

      <template #footer>
        <div class="flex justify-end space-x-4 p-4">
          <UButton variant="outline" @click="$emit('cancel')">Cancel</UButton>
          <UButton
            variant="solid"
            :loading="loading"
            :disabled="!state.name"
            @click="
              {
                loading = true;
                $emit('confirm', state);
              }
            "
            >Submit</UButton
          >
        </div>
      </template>
    </UCard>
  </UModal>
</template>

<style lang="scss" scoped></style>
