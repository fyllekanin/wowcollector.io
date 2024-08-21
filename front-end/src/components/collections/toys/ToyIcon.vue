<script lang="ts" setup>
import { useIntersectionObserver } from '@vueuse/core';

import type { PropType } from 'vue';
import type { ToyInformation } from '~/types';

const props = defineProps({
  toy: {
    type: Object as PropType<ToyInformation>,
    required: true,
  },
  clickable: {
    type: Boolean,
    default: true,
  },
  buildMode: {
    type: Boolean,
    default: false,
  },
  showTooltip: {
    type: Boolean,
    default: true,
  },
  useIntersectionObserver: {
    type: Boolean,
    default: true,
  },
});

const toyViewBuilderStore = useToyViewBuilderStore();
const { _settings } = storeToRefs(toyViewBuilderStore);

const showTooltip = computed(() => {
  return _settings.value.showToyTooltips;
});

const target = ref(null);
const targetIsVisible = ref(!props.useIntersectionObserver);

if (props.useIntersectionObserver) {
  const { stop } = useIntersectionObserver(
    target,
    ([{ isIntersecting }], observerElement) => {
      targetIsVisible.value = isIntersecting;
    }
  );
}
</script>

<template>
  <a
    :class="[buildMode ? 'cursor-move' : '']"
    :href="clickable ? `https://www.wowhead.com/item/${toy.itemId}` : '#'"
    target="_blank"
    :data-wowhead="showTooltip ? `item=${toy.itemId}` : ''"
    @click="!clickable && $event.preventDefault()"
  >
    <img
      v-if="targetIsVisible"
      :src="toy.assets.largeIcon"
      :class="[
        !toy.isCollected && !buildMode
          ? 'brightness-50 grayscale blur-[1px] transition ease-in-out hover:grayscale-0 hover:blur-[0px] hover:brightness-100 hover:ring-1 hover:ring-primary'
          : 'hover:ring-1 hover:ring-primary transition ease-in-out',
      ]"
      width="38"
      @error="
        toy.assets.largeIcon =
          'https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg'
      "
    />
  </a>
</template>
