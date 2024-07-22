<script lang="ts" setup>
import { object, string, type InferType } from 'yup';
import type { FormSubmitEvent } from '#ui/types';

const { data: page } = await useAsyncData('search', () =>
  queryContent('/search').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No search page found in the content.',
  });
}

useSeoMeta({
  title: page.value.title,
  description: page.value.description,
});

const { isMobile } = useScreenSize();

const schema = object({
  name: string().required(),
  realm: string().required(),
  region: string().required(),
});

type Schema = InferType<typeof schema>;
const state = reactive({
  name: undefined,
  realm: undefined,
  region: undefined,
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  // Do something with event.data
  console.log(event.data);
}
</script>

<template>
  <UContainer class="w-full flex items-center justify-center">
    <UCard
      class="w-96 mt-32"
      :ui="{
        ring: isMobile ? '' : 'ring-1 ring-gray-200 dark:ring-gray-800',
      }"
    >
      <template #header>Search for a character</template>

      <UForm
        :schema="schema"
        :state="state"
        class="space-y-4"
        @submit="onSubmit"
      >
        <UFormGroup label="Character Name" name="name">
          <UInput v-model="state.name" />
        </UFormGroup>

        <UFormGroup label="Realm" name="realm">
          <UInput v-model="state.realm" />
        </UFormGroup>

        <UFormGroup label="Region" name="region">
          <USelect
            v-model="state.region"
            :options="['EU', 'US', 'KR', 'TW', 'CN']"
          />
        </UFormGroup>
      </UForm>

      <template #footer>
        <div class="w-full flex justify-end">
          <UButton :disabled="!state.name || !state.realm || !state.region">
            Search
          </UButton>
        </div>
      </template>
    </UCard>
  </UContainer>
</template>
