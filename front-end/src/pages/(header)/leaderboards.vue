<script lang="ts" setup>
const columns = [
  {
    key: 'avatar',
    label: 'Avatar',
    sortable: false,
  },
  {
    key: 'name',
    label: 'Name',
    sortable: true,
  },
];

const paginatedCharacters = ref([
  {
    avatar:
      'https://cdn.wowhead.com/img/wow/icons/large/inv_misc_questionmark.jpg',
    name: 'Character 1',
  },
  {
    avatar:
      'https://cdn.wowhead.com/img/wow/icons/large/inv_misc_questionmark.jpg',
    name: 'Character 2',
  },
  {
    avatar:
      'https://cdn.wowhead.com/img/wow/icons/large/inv_misc_questionmark.jpg',
    name: 'Character 3',
  },
  {
    avatar:
      'https://cdn.wowhead.com/img/wow/icons/large/inv_misc_questionmark.jpg',
    name: 'Character 4',
  },
  {
    avatar:
      'https://cdn.wowhead.com/img/wow/icons/large/inv_misc_questionmark.jpg',
    name: 'Character 5',
  },
  {
    avatar:
      'https://cdn.wowhead.com/img/wow/icons/large/inv_misc_questionmark.jpg',
    name: 'Character 6',
  },
]);
</script>

<template>
  <UContainer class="flex flex-col w-full h-full">
    <div class="w-full">
      <h1 class="text-2xl font-bold py-4">Leaderboard</h1>
    </div>
    <UDivider />
    <UTable
      :columns="columns"
      :rows="paginatedCharacters"
      :loading="paginatedCharacters.length === 0"
      sort-asc-icon="i-heroicons-arrow-up"
      sort-desc-icon="i-heroicons-arrow-down"
      class="w-full lg:max-h-[35rem] max-h-[28rem]"
    >
      <template #avatar-data="{ row }">
        <UAvatar :src="row.avatar" />
      </template>
      <template #authority-data="{ row }"> </template>
      <template #actions-data="{ row }">
        <UDropdown>
          <UButton
            color="gray"
            variant="ghost"
            icon="i-heroicons-ellipsis-horizontal-20-solid"
          />
        </UDropdown>
      </template>
    </UTable>
  </UContainer>
</template>
