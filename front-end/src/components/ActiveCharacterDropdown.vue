<script lang="ts" setup>
import { Icons } from '~/constants';

import type { DropdownItem } from '#ui/types';

const { atCookie, logout } = useAuth();
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
    atCookie.value && {
      label: 'Change account',
      icon: 'material-symbols:logout',
      slot: 'logout',
    },
    character.value && {
      label: 'Change character',
      icon: 'basil:exchange-outline',
      to: '/search',
      slot: 'change-character',
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

    <template #logout="{ item }">
      <UButton
        :icon="item.icon"
        :disabled="item.disabled"
        variant="ghost"
        color="red"
        class="w-full"
        @click="logout"
        >{{ item.label }}</UButton
      >
    </template>
    <template #change-character="{ item }">
      <UButton
        :icon="item.icon"
        :to="item.to"
        :disabled="item.disabled"
        variant="ghost"
        color="white"
        class="w-full"
        @click="characterStore.clearCharacter"
        >{{ item.label }}</UButton
      >
    </template>
  </UDropdown>
</template>
