<template>
  <header
    class="flex items-center justify-between whitespace-nowrap border-b border-solid border-border-light px-4 sm:px-6 md:px-10 py-3 w-full bg-background-light text-text-dark font-display group/design-root">
    <div class="flex items-center gap-4 text-text-dark">
      <div class="size-6 text-primary">
        <svg fill="none" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
          <path d="M6 6H42L36 24L42 42H6L12 24L6 6Z" fill="currentColor"></path>
        </svg>
      </div>
      <h2 class="text-lg font-bold leading-tight tracking-[-0.015em]">Zylix</h2>
    </div>
    <div class="hidden md:flex flex-1 justify-end gap-8">
      <div class="flex gap-2">
        <button
          class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-primary text-text-dark text-sm font-bold leading-normal tracking-[0.015em] hover:opacity-90 transition-opacity">
          <span class="truncate">Sign Up</span>
        </button>
        <button
          class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-text-light text-text-dark text-sm font-bold leading-normal tracking-[0.015em] hover:bg-gray-200 transition-colors">
          <span class="truncate">Log In</span>
        </button>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
</script>
