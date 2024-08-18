<script lang="ts" setup></script>

<template>
  <div class="w-screen h-full">
    <div class="flex flex-col gap-4 h-full w-full items-center mt-20">
      <img src="../../assets/img/sad_orc.webp" />
      <h1 class="text-4xl font-semibold">Oh no!</h1>
      <p class="text-md text-center">
        It seems like your screen is too small to use the view builder. Please
        resize your screen to at least 1024px width.
      </p>
      <UButton to="/">Go back home</UButton>
    </div>
  </div>
</template>
