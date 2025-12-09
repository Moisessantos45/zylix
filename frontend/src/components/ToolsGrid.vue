<template>
  <section class="flex flex-col w-10/12 mx-auto">
    <h2 class="text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5 text-text-dark">
      Explore Our Tools</h2>
    <div class="grid grid-cols-[repeat(auto-fit,minmax(200px,1fr))] gap-4 p-4">
      <ToolCard v-for="tool in tools" :key="tool.title" :tool="tool" />
    </div>
  </section>
</template>

<script setup lang="ts">
import ToolCard from './ToolCard.vue';
import { tools } from '@/composables/tools';
</script>
