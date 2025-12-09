<template>
    <div class="w-full">
        <div class="w-full p-2 flex justify-between items-center border-b border-gray-200 mb-2">
            <div>
                <span class="text-gray-600 text-sm font-medium">Total Files: {{ fileItems.length }}</span>
            </div>
            <button @click="$emit('clear-all')"
                class="flex w-24 items-center justify-center gap-2 rounded-lg h-10 px-4 bg-gray-100 text-gray-700 text-sm font-medium hover:bg-gray-200 transition-opacity">
                <span>Clear All</span>
            </button>
        </div>        
        <div class="max-h-52 overflow-y-auto will-change-scroll">
            <div v-for="(file) in fileItems" :key="file.id"
                class="flex items-center gap-4 mt-1 bg-white px-4 min-h-14 justify-between rounded-xl border border-gray-200">
                <div class="flex items-center gap-4 overflow-hidden">
                    <div class="text-primary flex items-center justify-center rounded-lg bg-gray-100 shrink-0 size-10">
                        <FileIcon class="w-6 h-6" />
                    </div>
                    <p class="text-gray-900 text-base font-normal leading-normal flex-1 truncate">
                        {{ file.name }}{{ file.extension }}</p>
                </div>
                <div class="shrink-0">
                    <button @click="$emit('remove-file', file.id)"
                        class="text-gray-600 hover:text-gray-900 flex size-7 items-center justify-center">
                        <X class="w-5 h-5" />
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { FileIcon, X } from './icons';
import type { FileItem } from '../types';

defineProps<{
    fileItems: FileItem[],
}>()

defineEmits<{
    (e: 'remove-file', id: string): void,
    (e: 'clear-all'): void,
}>()

</script>