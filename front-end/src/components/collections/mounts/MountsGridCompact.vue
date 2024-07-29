<script lang="ts" setup>
const mountsStore = useMountsStore();

const { mounts, mountFilters } = storeToRefs(mountsStore);
console.log(mounts.value);
</script>

<template>
  <UContainer
    v-for="(category, i) in mounts"
    :key="i"
    class="flex flex-col w-full px-0 lg:px-0 sm:px-0 mx-0 gap-2"
  >
    <h2 class="text-lg">{{ category.name }}</h2>
    <UContainer
      v-if="category.mounts?.length"
      class="flex flex-wrap w-full self-start px-0 lg:px-0 sm:px-0 mx-0 gap-4"
    >
      <div v-for="(mount, j) in category.mounts" :key="j">
        <Lazy>
          <a
            :href="`https://www.wowhead.com/mount/${mount.id}`"
            target="_blank"
            :data-wowhead="`mount=${mount.id}`"
          >
            <img
              :src="mount.icon"
              :class="[!mount.isCollected ? 'not-collected' : '', 'mount-icon']"
              width="32"
              :on-error="(e: any) => (e.target.src = mount.creatureDisplay)"
              @error="
                mount.icon =
                  'https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg'
              "
            />
          </a>
        </Lazy>
      </div>
    </UContainer>
    <UContainer
      class="flex flex-wrap w-full self-start px-0 lg:px-0 sm:px-0 mx-0 gap-4"
    >
      <UContainer
        v-for="(subCategory, j) in category.categories?.filter(
          (c) => c.mounts?.length
        )"
        :key="j"
        class="flex grow flex-wrap px-0 lg:px-0 sm:px-0 mx-0"
      >
        <div class="flex flex-col gap-4">
          <h3 class="text-xs">{{ subCategory.name }}</h3>
          <UContainer
            class="flex grow flex-wrap gap-2 px-0 lg:px-0 sm:px-0 mx-0"
          >
            <div v-for="(mount, k) in subCategory.mounts" :key="k">
              <Lazy>
                <a
                  :href="`https://www.wowhead.com/mount/${mount.id}`"
                  target="_blank"
                  :data-wowhead="`mount=${mount.id}`"
                >
                  <img
                    :src="mount.icon"
                    :class="[
                      !mount.isCollected ? 'not-collected' : '',
                      'mount-icon',
                    ]"
                    width="32"
                    :on-error="(e: any) => (e.target.src = mount.creatureDisplay)"
                    @error="
                      mount.icon =
                        'https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg'
                    "
                  />
                </a>
              </Lazy>
            </div>
          </UContainer>
        </div>
      </UContainer>
    </UContainer>
  </UContainer>
</template>
