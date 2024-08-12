<script lang="ts" setup>
import type { PropType } from 'vue';
import type { ToyInformation } from '~/types';

defineProps({
  toy: {
    type: Object as PropType<ToyInformation>,
    required: true,
  },
});
</script>

<template>
  <a
    :href="`https://www.wowhead.com/item/${toy.itemId}`"
    target="_blank"
    :data-wowhead="`item=${toy.itemId}`"
  >
    <img
      :src="toy.assets.largeIcon"
      :class="[
        !toy.isCollected
          ? 'brightness-50 grayscale blur-[1px] transition ease-in-out hover:grayscale-0 hover:blur-[0px] hover:brightness-100 hover:ring-1 hover:ring-primary'
          : '',
      ]"
      width="38"
      @error="
        toy.assets.largeIcon =
          'https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg'
      "
    />
  </a>
</template>
