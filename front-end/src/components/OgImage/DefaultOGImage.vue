<template>
  <div class="h-full w-full flex items-start justify-start bg-slate-900">
    <div class="flex items-start justify-start h-full">
      <div class="flex flex-col justify-between w-full h-full">
        <div class="flex flex-col">
          <h1 class="text-[80px] px-20 pt-20 pb-5 font-black text-left">
            <img
              src="../../assets/img/logo.png"
              style="width: 512px; height: auto"
            />
          </h1>
          <p class="text-2xl pb-20 px-20 font-bold mb-0 text-gray-300">
            The only World of Warcraft collections tracker you will ever need.
          </p>
        </div>

        <div class="flex gap-6 items-center px-20 pb-10">
          <div
            class="w-[48px] h-[48px] border-[1px] border-gray-700 rounded-md flex justify-center items-center"
          >
            <UIcon name="mdi:horse-variant" style="scale: 1.5; color: orange" />
          </div>
          <p class="text-lg font-bold text-gray-300">Collectable tracking</p>
        </div>

        <div class="flex gap-6 items-center px-20 pb-10">
          <div
            class="w-[48px] h-[48px] border-[1px] border-gray-700 rounded-md flex justify-center items-center"
          >
            <UIcon name="mdi:tools" style="scale: 1.5; color: orange" />
          </div>
          <p class="text-lg font-bold text-gray-300">Custom view builder</p>
        </div>

        <div class="flex gap-6 items-center px-20 pb-10">
          <p class="text-lg font-bold text-gray-300">
            ... and much more coming!
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
