<script setup lang="ts">
import { Icons } from '~/constants';

const emit = defineEmits(['close']);

const { consent, setConsent } = useConsent();
const toast = useToast();
const { width } = useWindowSize();

const localConsent = ref({ ...consent.value });

const tabs = computed(() => [
  {
    label: `Essential${width.value > 640 ? ' Cookies' : ''}`,
    icon: 'i-heroicons-check-circle',
    slot: 'essential',
  },
  {
    label: `Analytics${width.value > 640 ? ' Cookies' : ''}`,
    icon: 'i-heroicons-chart-pie',
    slot: 'analytics',
  },
]);

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
      <UTabs
        :items="tabs"
        :ui="{
          wrapper: 'flex flex-col sm:flex-row grow justify-between gap-10',
        }"
        :orientation="width > 640 ? 'vertical' : 'horizontal'"
      >
        <template #default="{ item, selected }">
          <span
            class="truncate"
            :class="[selected && 'text-primary-500 dark:text-primary-400']"
            >{{ item.label }}</span
          >
        </template>

        <template #essential>
          <div class="flex flex-col gap-4 justify-start">
            <div class="flex gap-10 items-center w-full justify-between">
              <span class="text-lg">Essential Cookies</span>
              <UToggle v-model="localConsent.essential" :disabled="true" />
            </div>
            <span class="text-sm">
              Essential cookies are necessary for the website to function
              properly. This category only includes cookies that ensures basic
              functionalities and security features of the website such as user
              login, account management, and language preference. These cookies
              do not store any personal information.
            </span>
          </div>
        </template>

        <template #analytics>
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
        </template>
      </UTabs>
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
