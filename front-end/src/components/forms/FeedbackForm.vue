<script lang="ts" setup>
import { number, object, string, type InferType } from 'yup';
import type { FormSubmitEvent } from '#ui/types';

// collections , character search etc
const SITE_AREA_REGEXP = /^(collections|character_search)$/;
const schema = object({
  siteArea: string().matches(SITE_AREA_REGEXP).required(),
  feedback: string().required(),
  attachment: string().nullable(),
  currentUserExperience: number(),
});

type Schema = InferType<typeof schema>;

const state = reactive({
  siteArea: 'collections',
  feedback: '',
  attachment: null as string | null,
  currentUserExperience: 0,
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  // Do something with event.data
  console.log(event.data);
}
</script>

<template>
  <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit">
    <UFormGroup label="Site Area" name="siteArea">
      <URadioGroup
        v-model="state.siteArea"
        color="primary"
        :options="[
          {
            value: 'collections',
            label: 'Collections (Mounts, Achievements..)',
          },
          { value: 'character_search', label: 'Character Search' },
        ]"
      />
    </UFormGroup>

    <UFormGroup label="Feedback" name="feedback">
      <UTextarea
        v-model="state.feedback"
        resize
        autoresize
        :rows="10"
        placeholder="Enter your suggestions here, be as detailed as possible."
      />
    </UFormGroup>

    <UButton @click="$emit('submit', state)" color="primary">Submit</UButton>
  </UForm>
</template>
