<script lang="ts" setup>
import { useIntersectionObserver } from '@vueuse/core';

import type { PropType } from 'vue';
import type { PetInformation } from '~/types';

const props = defineProps({
  pet: {
    type: Object as PropType<PetInformation>,
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
  <div ref="target">
    <a
      :class="[buildMode ? 'cursor-move' : '']"
      :href="clickable ? `https://www.wowhead.com/battle-pet/${pet.id}` : '#'"
      target="_blank"
      :data-wowhead="showTooltip ? `battle-pet=${pet.id}` : ''"
      @click="!clickable && $event.preventDefault()"
    >
      <img
        v-if="targetIsVisible"
        :src="pet.assets.largeIcon"
        :class="[
          !pet.isCollected && !buildMode
            ? 'brightness-50 grayscale blur-[1px] transition ease-in-out hover:grayscale-0 hover:blur-[0px] hover:brightness-100 hover:ring-1 hover:ring-primary'
            : 'hover:ring-1 hover:ring-primary transition ease-in-out',
        ]"
        width="38"
        @error="
          pet.assets.largeIcon =
            'https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg'
        "
      />
      <div v-else class="w-10 h-10 rounded-lg" />
    </a>
  </div>
</template>
