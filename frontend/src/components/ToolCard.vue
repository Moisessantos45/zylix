<template>
  <RouterLink :to="{ name: tool.route }"
    class="flex flex-1 cursor-pointer flex-col gap-3 rounded-xl border border-border-light bg-white p-4 hover:bg-gray-50 transition-colors">
    <div class="text-primary w-10 h-10">
      <component :is="iconComponent" class="w-full h-full" />
    </div>
    <div class="flex flex-col gap-1">
      <h2 class="text-base font-bold leading-tight text-text-dark">
        {{ tool.subtitle }}
      </h2>
      <p class="text-sm font-normal leading-normal text-subtle-text-light">
        {{ tool.description }}
      </p>
    </div>
  </RouterLink>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { Tool } from '../types';
import * as Icons from './icons';

const props = defineProps<{
  tool: Tool
}>();

const iconComponent = computed(() => {
  const iconName = props.tool.icon as keyof typeof Icons;
  return Icons[iconName] || Icons.FileIcon;
});
</script>
