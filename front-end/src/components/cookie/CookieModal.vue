<script setup lang="ts">
import { Icons } from '~/constants';

const emit = defineEmits(['close']);

const { consent, setConsent } = useConsent();
const toast = useToast();

const localConsent = ref({ ...consent.value });

const tabs = [
  {
    label: 'Essential Cookies',
    icon: 'i-heroicons-check-circle',
    content: '',
  },
  {
    label: 'Analytics Cookies',
    icon: 'i-heroicons-chart-pie',
    content: '',
  },
];

const savePreferences = () => {
  setConsent(localConsent.value);
  toast.add({
    title: 'Cookie preferences saved',
    color: 'green',
    timeout: 2000,
  });
  emit('close');
};

const closeModal = () => emit('close');
</script>

<template>
  <UModal
    @close="closeModal"
    title="Manage Cookie Preferences"
    :ui="{
      width: 'w-full sm:max-w-2xl',
    }"
  >
    <UCard>
      <template #header>
        <div class="flex justify-between items-center">
          <h2 class="text-xl font-semibold">Cookie Preferences</h2>
          <UButton
            @click="closeModal"
            variant="ghost"
            color="gray"
            :icon="Icons.CLOSE"
          >
          </UButton>
        </div>
      </template>
      <div class="flex grow justify-between gap-10">
        <UTabs :items="tabs" class="w-1/3" orientation="vertical">
          <template #default="{ item, selected }">
            <span
              class="truncate"
              :class="[selected && 'text-primary-500 dark:text-primary-400']"
              >{{ item.label }}</span
            >
          </template>
        </UTabs>
        <div class="flex flex-col gap-4 justify-start">
          <div class="flex gap-10 items-center w-full justify-between">
            <span class="text-lg">Analytics Cookies</span>
            <UToggle v-model="localConsent.analytics" />
          </div>
          <span class="text-sm">
            We use cookies to improve your experience on our website. By
            allowing analytics cookies, you agree to the storing of cookies on
            your device to enhance site navigation, analyze site usage, and
            assist in the feuture development of the website.
          </span>
        </div>
      </div>
      <template #footer>
        <div class="flex justify-between">
          <UButton @click="savePreferences">Save Preferences</UButton>
          <UButton @click="setConsent({ essential: true, analytics: true })"
            >Allow All</UButton
          >
        </div>
      </template>
    </UCard>
  </UModal>
</template>

<style scoped>
.p-4 {
  padding: 1rem;
}
</style>
