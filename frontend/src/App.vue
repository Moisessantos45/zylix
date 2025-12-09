<script lang="ts" setup>
import Header from './components/Header.vue';
</script>

<template>
  <Header />
  <main class="min-h-screen bg-background-light text-[#111813] font-display group/design-root">
    <RouterView />
  </main>
</template>
