<script lang="ts" setup>
import { Icons } from '~/constants';

import type { DropdownItem } from '#ui/types';

const characterStore = useCharacterStore();
const { character } = storeToRefs(characterStore);

const items = computed(() => [
  [
    {
      name: character.value?.name,
      realm: character.value?.realm,
      slot: 'account',
      disabled: true,
    },
  ],
  [
    character.value && {
      label: 'Change character',
      icon: 'material-symbols:logout',
      to: '/search',
    },
  ],
]) as ComputedRef<DropdownItem[][]>;
</script>

<template>
  <UDropdown
    mode="hover"
    :close-delay="50"
    :items="items"
    :ui="{ item: { disabled: 'cursor-text select-text' } }"
    :popper="{ placement: 'bottom-start' }"
  >
    <UChip
      size="md"
      position="bottom-right"
      inset
      :ui="{ base: '-mx-2 rounded-none ring-0', background: '' }"
    >
      <UAvatar
        :src="character?.assets?.avatar"
        :icon="
          !character?.assets?.avatar
            ? 'material-symbols:account-circle-full'
            : undefined
        "
        size="md"
        alt="Avatar"
      />

      <!-- When logged in with Bnet -->
      <template #content>
        <UAvatar
          :icon="Icons.BATTLENET"
          alt="Avatar"
          size="2xs"
          :ui="{
            rounded: 'rounded-md',
            icon: {
              base: 'dark:bg-white bg-black scale-150',
            },
          }"
          class="shadow-md"
        />
      </template>
    </UChip>

    <template #account="{ item }">
      <div class="text-left flex flex-col">
        <span class="font-bold">{{ item.name }}</span>
        <span>{{ item.realm }}</span>
      </div>
    </template>

    <template #item="{ item }">
      <UButton
        :icon="item.icon"
        :to="item.to"
        :disabled="item.disabled"
        variant="ghost"
        color="red"
        class="w-full"
        @click="characterStore.clearCharacter"
        >{{ item.label }}</UButton
      >
    </template>
  </UDropdown>
</template>
