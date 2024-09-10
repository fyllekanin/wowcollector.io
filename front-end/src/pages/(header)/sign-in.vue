<script lang="ts" setup>
import { Icons } from '~/constants';

definePageMeta({
  layout: 'signin',
});

const toast = useToast();
const router = useRouter();
const runtimeConfig = useRuntimeConfig();
const loading = ref(false);

const route = useRoute();
const oauthCode = route.query.code as string | undefined;

if (oauthCode) {
  loading.value = true;

  try {
    const response = await $fetch(
      `/api/auth/battle-net?code=${oauthCode}&redirect_uri=${runtimeConfig.public.BNET_REDIRECT_URI}&scope=${runtimeConfig.public.BNET_SCOPE}`
    );

    console.log(response);
  } catch (error) {
    console.error('OAuth error:', error);
    toast.add({
      title: 'Error',
      description: 'An error occurred while signing in. Please try again.',
      color: 'red',
    });
  } finally {
    loading.value = false;
  }

  // console.log('OAuth code:', oauthCode);
  // setTimeout(() => {
  //   loading.value = false;
  // }, 3000);
}
</script>

<template>
  <div v-if="loading" class="h-full w-full flex items-center justify-center">
    <LogoLoader class="w-32 h-32" />
  </div>
  <div v-else class="flex flex-col items-center justify-center h-full p-4">
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
          :to="`https://oauth.battle.net/authorize?response_type=code&&state=AbCdEfG&scope=${runtimeConfig.public.BNET_SCOPE}&redirect_uri=${runtimeConfig.public.BNET_REDIRECT_URI}&client_id=${runtimeConfig.public.BNET_CLIENT_ID}`"
        >
          Battle.net
        </UButton>
        <UButton
          block
          :icon="Icons.DISCORD_COLOR"
          color="gray"
          size="lg"
          :to="`https://discord.com/oauth2/authorize?client_id=${runtimeConfig.public.DISCORD_CLIENT_ID}&response_type=code&redirect_uri=${runtimeConfig.public.DISCORD_REDIRECT_URI}&scope=${runtimeConfig.public.DISCORD_SCOPE}`"
        >
          Discord
        </UButton>
      </div>
      <UDivider />
      <p class="text-sm dark:text-gray-400 text-white text-center">
        By signing in, you agree to our
        <UButton
          variant="link"
          size="xs"
          :ui="{
            padding: { xs: 'px-0 py-0' },
          }"
          to="/terms"
          >Terms of Service</UButton
        >
        and
        <UButton
          variant="link"
          size="xs"
          :ui="{
            padding: { xs: 'px-0 py-0' },
          }"
          to="/privacy"
          >Privacy Policy</UButton
        >
      </p>
    </div>
  </div>
</template>
