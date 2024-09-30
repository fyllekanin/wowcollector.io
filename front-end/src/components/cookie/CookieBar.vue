<script lang="ts" setup>
import CookieModal from '~/components/cookie/CookieModal.vue';

const { isConsentGiven, setConsent } = useConsent();
const modal = useModal();

const showCookieBar = computed(() => !isConsentGiven.value);

const giveConsent = () => {
  setConsent({ essential: true, analytics: true });
};

const rejectConsent = () => {
  setConsent({ essential: true, analytics: false });
};
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
    class="min-h-20 w-full px-4 py-2 gap-10 flex flex-col md:flex-row grow items-end md:items-center justify-between border-t-[1px] border-gray-700 dark:border-gray-700 fixed bottom-0 bg-white dark:bg-slate-900"
  >
    <div class="flex gap-1 self-start md:self-center">
      <p class="text-md">
        We use cookies to ensure you get the best experience on our website.
        <UButton
          variant="link"
          size="sm"
          @click="
            modal.open(CookieModal, {
              onClose: () => {
                modal.close();
              },
            })
          "
          :ui="{
            padding: {
              sm: 'px-0',
            },
          }"
          >Show Purposes</UButton
        >
      </p>
    </div>
    <div class="flex grow gap-2 justify-end">
      <UButton variant="ghost" color="gray" size="lg" @click="rejectConsent"
        >Reject</UButton
      >
      <UButton variant="outline" size="lg" @click="giveConsent"
        >Accept All</UButton
      >
    </div>
  </div>
</template>

<style scoped></style>
