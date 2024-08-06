<script lang="ts" setup>
import type { PropType } from 'vue';
import type { MountInformation } from '~/types';

defineProps({
  mount: {
    type: Object as PropType<MountInformation>,
    required: true,
  },
});
</script>

<template>
  <a
    :href="`https://www.wowhead.com/mount=${mount.id}`"
    target="_blank"
    :data-wowhead="`mount=${mount.id}`"
  >
    <img
      :src="mount.assets.largeIcon"
      :class="[
        !mount.isCollected
          ? 'brightness-50 grayscale blur-[1px] transition ease-in-out hover:grayscale-0 hover:blur-[0px] hover:brightness-100 hover:ring-1 hover:ring-primary'
          : '',
      ]"
      width="38"
      @error="
        mount.assets.largeIcon =
          'https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg'
      "
    />
  </a>
</template>
