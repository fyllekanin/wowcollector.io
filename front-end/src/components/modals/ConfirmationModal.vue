<script lang="ts" setup>
defineEmits(['cancel', 'confirm']);
defineProps({
  title: {
    type: String,
    required: true,
  },
  message: {
    type: String,
    required: true,
  },
  additionaInformation: {
    type: String,
    default: '',
  },
  buttonText: {
    type: String,
    default: 'Confirm',
  },
});

const loading = ref(false);
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
        <h2 class="text-lg font-semibold">{{ title }}</h2>
      </template>

      <div class="flex flex-col">
        <p class="p-4">{{ message }}</p>
        <p v-if="additionaInformation" class="p-4">
          {{ additionaInformation }}
        </p>
      </div>

      <template #footer>
        <div class="flex justify-end space-x-4 p-4">
          <UButton variant="outline" @click="$emit('cancel')">Cancel</UButton>
          <UButton
            variant="solid"
            :loading="loading"
            @click="
              {
                loading = true;
                $emit('confirm');
              }
            "
            >{{ buttonText }}</UButton
          >
        </div>
      </template>
    </UCard>
  </UModal>
</template>

<style lang="scss" scoped></style>
