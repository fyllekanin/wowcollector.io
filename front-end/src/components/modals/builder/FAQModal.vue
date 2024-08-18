<script lang="ts" setup>
import type { AccordionItem } from '#ui/types';

defineEmits(['close']);

const props = defineProps({
  kind: {
    type: String,
    required: true,
    enum: ['mounts', 'pets', 'achievements', 'toys'],
  },
});

const kindWithoutS = props.kind.slice(0, -1);

const items = [
  {
    label: 'What is the View Builder?',
    content: `
      The View Builder is a tool that allows you to create custom views for our collections. You can add ${props.kind} and categories to your view however you like, and customize the layout to suit your needs.
    `,
    defaultOpen: true,
  },
  {
    label: `How do I add a new ${kindWithoutS} to my view?`,
    content: `
      To add a new ${kindWithoutS} to your view, grab a ${kindWithoutS}-icon in the sidebar and drag it to a category in the main view area. This will add the ${kindWithoutS} to your view.
    `,
  },
  {
    label: `How do I remove ${props.kind} from my view?`,
    content: `
      To remove a ${kindWithoutS} from your view, grab the ${kindWithoutS}-icon and drag it back to the sidebar. This will remove the ${kindWithoutS} from your view.
    `,
  },
  {
    label: 'How do I change the layout of my view?',
    content: `
      To change the layout of your view, either grab an already added category and drag it to a new position, or grab a new category and drag it to the main view area.
    `,
  },
  {
    label: 'How do I save my view?',
    content: `
      To save your view, click the "Create view" button in the topbar. You will be prompted to enter a name for your view, and then you can save it. This will in return give you a link that you can use to load your view later.
    `,
  },
  {
    label: 'How do I load a saved view?',
    content: `
      To load a saved view, use the link that was generated when you saved the view. Paste the link in your browser and the view will be loaded.
    `,
  },
  {
    label: 'How do I delete a saved view?',
    content: `
      At the moment, you cannot delete a saved view.
    `,
  },
  {
    label: 'How do I edit a saved view?',
    content: `
      At the moment, you cannot edit a saved view. Once you have saved a view, it is final.
    `,
  },
  {
    label: 'How do I share my view with others?',
    content: `
      To share your view with others, use the link that was generated when you saved the view. Share the link with others and they will be able to view your view.
    `,
  },
] as AccordionItem[];
</script>

<template>
  <UModal fullscreen>
    <UCard
      :ui="{
        ring: '',
        divide: 'divide-y divide-gray-100 dark:divide-gray-800',
        body: {
          base: 'flex justify-center',
        },
      }"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <h2 class="text-lg font-semibold">View Builder - FAQ</h2>
          <UButton
            variant="ghost"
            color="gray"
            size="xl"
            @click="$emit('close')"
          >
            <UIcon class="scale-125" name="mdi:close" />
          </UButton>
        </div>
      </template>

      <div class="flex w-[80%] justify-center">
        <ULandingFAQ :items="items" multiple />
      </div>

      <template #footer>
        <div class="flex justify-end space-x-4 p-4">
          <UButton variant="solid" @click="$emit('close')">Close</UButton>
        </div>
      </template>
    </UCard>
  </UModal>
</template>
