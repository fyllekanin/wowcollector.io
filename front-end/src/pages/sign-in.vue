<script lang="ts" setup>
import { Icons } from '~/constants';

definePageMeta({
  layout: 'signin',
});

const router = useRouter();
</script>

<template>
  <div class="flex flex-col items-center justify-center h-full p-4">
    <div class="flex flex-col items-center w-full max-w-lg gap-4">
      <LogoFull class="cursor-pointer" @click="router.push('/')" />
      <p class="text-md font-semibold">Connect using your favorite provider</p>
      <div class="flex flex-col w-full gap-4">
        <UButton
          block
          :icon="Icons.BATTLENET"
          color="gray"
          size="lg"
          :ui="{
            icon: {
              base: 'bg-blue-500',
            },
          }"
        >
          Battle.net
        </UButton>
        <UButton block :icon="Icons.DISCORD_COLOR" color="gray" size="lg">
          Discord
        </UButton>
      </div>
      <UDivider />
      <p class="text-sm dark:text-gray-400 text-white">
        By signing in, you agree to our
        <UButton
          variant="link"
          size="xs"
          :ui="{
            padding: { xs: 'px-0 py-0' },
          }"
          >Terms of Service</UButton
        >
        and
        <UButton
          variant="link"
          size="xs"
          :ui="{
            padding: { xs: 'px-0 py-0' },
          }"
          >Privacy Policy</UButton
        >
      </p>
    </div>
  </div>
</template>
