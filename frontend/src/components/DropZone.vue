<template>
  <div
    class="flex flex-col items-center justify-center gap-6 rounded-xl border-2 border-dashed border-slate-300 p-8 sm:p-14 text-center bg-white">
    <div class="flex items-center justify-center size-16 bg-primary/20 rounded-full">
      <component :is="icon" class="w-8 h-8 text-primary" />
    </div>
    <div class="flex flex-col items-center gap-2">
      <p class="text-lg font-bold text-slate-900">{{ title }}</p>
      <p class="text-sm text-slate-500">{{ subtitle }}</p>
    </div>
    <button type="button" @click="handleFileInput"
      class="flex items-center justify-center rounded-lg h-10 px-4 bg-slate-100 text-slate-900 text-sm font-bold hover:bg-slate-200 transition-colors">
      <span class="truncate">Browse Files</span>
    </button>
  </div>
</template>

<script setup lang="ts">
import { type Component } from 'vue';
import { storeToRefs } from 'pinia';
import type { FileItem } from '../types';
import { GetFiles } from "../../wailsjs/go/main/App"
import useFileStore from '@/stores/file';

const useFile = useFileStore();
const { isFirstUpload } = storeToRefs(useFile)

const props = defineProps<{
  title: string
  subtitle: string
  icon: Component
  displayName: string
  pattern: string
}>()


const emit = defineEmits<{
  'files-uploaded': [files: FileItem[]]
  'file-removed': [index: number]
}>();

const handleFileInput = async () => {
  try {
    const files = await GetFiles(props.displayName, props.pattern, isFirstUpload.value);
    console.log('Selected files:', files);
    processFiles(files);
  } catch (error) {
    console.warn('File selection was cancelled or failed.', error);
  }
};


const processFiles = (rawFiles: string[]) => {
  const fileItems: FileItem[] = rawFiles.map((file, index) => ({
    id: `${Date.now()}-${index}`,
    path: file,
    name: file.split('/').pop() || 'unknown',
    extension: file.split('.').pop() || '',
  }));
  emit('files-uploaded', fileItems);
};


</script>

<style scoped>
.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}

.list-enter-from {
  opacity: 0;
  transform: translateX(-20px);
}

.list-leave-to {
  opacity: 0;
  transform: translateX(20px);
}
</style>
