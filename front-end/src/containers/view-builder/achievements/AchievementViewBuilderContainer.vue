<script setup lang="ts">
import ConfirmationModal from '~/components/modals/ConfirmationModal.vue';
import FAQModal from '~/components/modals/builder/FAQModal.vue';
import SubmitViewModal from '~/components/modals/builder/SubmitViewModal.vue';

const achievementViewBuilderStore = useAchievementViewBuilderStore();

const modal = useModal();

function openModal() {
  if (!achievementViewBuilderStore.hasChanges) {
    const router = useRouter();
    router.push('/');
    return;
  }

  modal.open(ConfirmationModal, {
    title: 'Exit build mode',
    message:
      'You have made some changes, are you sure you want to exit build mode?',
    additionaInformation: 'All progress will be lost.',
    onConfirm: () => {
      achievementViewBuilderStore.resetStore();
      const router = useRouter();
      router.push('/');
      modal.close();
    },
    onCancel: () => {
      modal.close();
    },
  });
}

function openHelpModal() {
  modal.open(FAQModal, {
    kind: 'achievements',
    onClose: () => {
      modal.close();
    },
  });
}

function openNamePromptModal() {
  modal.open(SubmitViewModal, {
    onConfirm: (name: string) => {
      modal.close();
    },
    onCancel: () => {
      modal.close();
    },
  });
}
</script>

<template>
  <UDashboardLayout>
    <UDashboardPanel
      :width="400"
      :resizable="{ min: 200, max: 500 }"
      collapsible
    >
      <UDashboardPanelContent>
        <div class="flex flex-col w-full h-full">
          <UDashboardNavbar
            class="!border-transparent"
            :ui="{ left: 'flex-1', wrapper: 'pb-4 px-0' }"
          >
            <template #left>
              <div class="flex grow flex-wrap items-center justify-center">
                <LogoFull class="flex cursor-pointer" @click="openModal" />
              </div>
            </template>
          </UDashboardNavbar>

          <UDivider />

          <slot name="sidebar-content" />
        </div>
      </UDashboardPanelContent>
    </UDashboardPanel>

    <div class="flex flex-col w-full h-full">
      <UDashboardNavbar class="" :ui="{ wrapper: 'pb-' }">
        <template #title>
          <div class="flex gap-2 items-center">
            <UIcon class="scale-125" name="mdi:tools" />
            <h2 class="text-lg font-semibold">Achievement View Builder</h2>
          </div>
        </template>

        <template #center>
          <UButton variant="ghost" color="gray" @click="openHelpModal"
            >What is this?</UButton
          >
        </template>

        <template #right>
          <UButton
            :disabled="
              !achievementViewBuilderStore.hasChanges ||
              !achievementViewBuilderStore.isValid
            "
            icon="mdi:plus"
            @click="openNamePromptModal"
            >Create view</UButton
          >
        </template>
      </UDashboardNavbar>
      <slot name="main-content" />
    </div>
  </UDashboardLayout>
</template>
