<script lang="ts" setup>
const { data: page } = await useAsyncData('feedback', () =>
  queryContent('/footer/feedback').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No feedback page found in the content.',
  });
}

const feedBackState = reactive({
  feedback: '',
  attachments: [],
  currentUserExperience: 0,
  email: '',
  battleTag: '',
});
</script>

<template>
  <UContainer class="flex max-w-[1200px] mx-auto mb-10">
    <UTabs class="w-full" :items="page?.tabs" orientation="horizontal">
      <template #item="{ item }">
        <UCard @submit.prevent="">
          <template #header>
            <div class="flex flex-col gap-2">
              <p class="text-xl font-semibold leading-6">
                {{ item.label }}
              </p>
              <p class="mt-1 text-sm">
                {{ item.description }}
              </p>
              <p class="mt-1 text-xs italic">
                {{ item.additionalNotes }}
              </p>
            </div>
          </template>

          <div v-if="item.feedbackType === 'feedback'" class="space-y-6">
            <UFormGroup label="Description" name="feedback" required>
              <UTextarea
                v-model="feedBackState.feedback"
                resize
                autoresize
                :rows="10"
                placeholder="Enter your suggestions here, be as detailed as possible."
              />
            </UFormGroup>

            <UFormGroup label="Attachments" name="attachments">
              <UButtonGroup>
                <UInput
                  type="file"
                  size="md"
                  icon="i-heroicons-folder"
                  accept="image/*,video/*"
                  multiple
                />
                <UButton
                  icon="i-heroicons-trash"
                  color="red"
                  :disabled="!feedBackState.attachments.length"
                />
              </UButtonGroup>
            </UFormGroup>

            <UFormGroup
              class="flex flex-col gap-4"
              label="How would you rate the current user experience so far?"
              name="currentUserExperience"
            >
              <div class="flex items-center gap-8">
                <p>Poor</p>
                <URadio
                  class="flex flex-col-reverse gap-2 items-center"
                  :ui="{
                    inner: 'ms-0 flex flex-col gap-2',
                  }"
                  v-for="i in 5"
                  :key="i"
                >
                  <template #label>
                    {{ i }}
                  </template>
                </URadio>
                <p>Excellent</p>
              </div>
            </UFormGroup>

            <UFormGroup label="Email Address" name="email">
              <UInput
                v-model="feedBackState.email"
                type="email"
                placeholder="Email Address"
              />
            </UFormGroup>
            <UFormGroup label="Battle tag" name="battleTag">
              <UInput
                v-model="feedBackState.battleTag"
                type="text"
                placeholder="Battle tag"
              />
            </UFormGroup>
          </div>

          <div v-else-if="item.feedbackType === 'bug'" class="space-y-6"></div>

          <template #footer>
            <div class="flex w-full justify-end">
              <UButton type="submit" color="primary">
                Submit {{ item.label }}</UButton
              >
            </div>
          </template>
        </UCard>
      </template>
    </UTabs>
  </UContainer>
</template>
