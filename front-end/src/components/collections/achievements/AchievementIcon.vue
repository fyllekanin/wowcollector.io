<script lang="ts" setup>
import { useIntersectionObserver } from '@vueuse/core';

import type { PropType } from 'vue';
import type { AchievementInformation } from '~/types';

const props = defineProps({
  achievement: {
    type: Object as PropType<AchievementInformation>,
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
    default: false,
  },
});

const achievementViewBuilderStore = useAchievementViewBuilderStore();
const { _settings } = storeToRefs(achievementViewBuilderStore);

const showTooltip = computed(() => {
  return _settings.value.showAchievementTooltips;
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
    :href="
      clickable ? `https://www.wowhead.com/achievement/${achievement.id}` : '#'
    "
    target="_blank"
    :data-wowhead="showTooltip ? `achievement=${achievement.id}` : ''"
    @click="!clickable && $event.preventDefault()"
  >
    <nuxt-img
      v-if="targetIsVisible"
      :src="achievement.icon"
      :class="[
        !achievement.isCompleted && !buildMode
          ? 'brightness-50 grayscale blur-[1px] transition ease-in-out hover:grayscale-0 hover:blur-[0px] hover:brightness-100 hover:ring-1 hover:ring-primary'
          : 'hover:ring-1 hover:ring-primary transition ease-in-out',
      ]"
      width="38"
      loading="lazy"
      @error="
        achievement.icon =
          'https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg'
      "
    />
  </a>
</template>
