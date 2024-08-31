<script lang="ts" setup>
const consent = useCookie<undefined | null | '0' | '1'>('_wc_consent_optin');
const isMounted = ref(false);
const showCookieBar = computed(
  () =>
    // (consent.value === undefined || consent.value === null) && isMounted.value
    false
);

const giveConsent = () => (consent.value = '1');
const rejectConsent = () => (consent.value = '0');

onMounted(() => {
  isMounted.value = true;
});
</script>

<template>
  <div
    v-if="showCookieBar"
    v-motion
    :initial="{
      opacity: 0,
      y: 100,
    }"
    :enter="{
      opacity: 1,
      y: 0,
      transition: {
        type: 'keyframes',
        stiffness: '100',
        delay: 0,
        duration: 500,
      },
    }"
    :leave="{
      opacity: 0,
      y: 100,
      transition: {
        type: 'keyframes',
        stiffness: '100',
        delay: 0,
        duration: 500,
      },
    }"
    class="min-h-20 w-full px-4 flex grow items-center justify-between border-t-[1px] border-gray-700 dark:border-gray-700 fixed bottom-0 bg-white dark:bg-slate-900"
  >
    <p class="text-md">
      We only use cookies for analytics to improve our service for your benefit.
    </p>
    <div class="flex gap-2">
      <UButton variant="outline" size="lg" @click="giveConsent"
        >Accept All</UButton
      >
      <UButton variant="ghost" color="gray" size="lg" @click="rejectConsent"
        >Reject</UButton
      >
    </div>
  </div>
</template>

<style scoped></style>
