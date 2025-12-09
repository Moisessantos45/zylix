<template>
    <section class="w-full px-4 py-5 will-change-scroll">
        <div class="w-full p-4 flex">
            <RouterLink to="/" class="text-primary flex items-center gap-1 hover:underline">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                    <path d="M15 18l-6-6 6-6" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                        stroke-linejoin="round" />
                </svg>
                Back to Home
            </RouterLink>
        </div>
        <Loading :visible="loading" message="Converting PDF to Images" />
        <div class="flex flex-wrap justify-between gap-3 p-4">
            <div class="flex min-w-72 flex-col gap-3 text-center w-full items-center">
                <p class="text-gray-900 text-4xl font-black leading-tight tracking-[-0.033em]">
                    Extract Pages from Your PDF Easily
                </p>
                <p class="text-gray-600 text-base font-normal leading-normal max-w-2xl">
                    Instantly extract specific pages from any PDF document online.
                </p>
            </div>
        </div>
        <div v-if="fileItems.length === 0" class="flex flex-col p-4">
            <DropZone title="Drag & drop your PDF here" subtitle="Supports PDF files. Max 10MB." :icon="Upload"
                display-name="PDFs (*.pdf)" pattern="*.pdf" @files-uploaded="handleFilesUploaded" />
        </div>

        <FilesList v-if="fileItems.length > 0" :fileItems="fileItems" @remove-file="removeFile" />

        <div class="p-4 flex flex-col gap-3">
            <button type="button" @click="selectFolder"
                class="flex w-full items-center justify-center gap-2 rounded-xl h-12 px-6 bg-slate-200 text-slate-700 text-base font-medium hover:opacity-90 transition-opacity">
                <span class="material-symbols-outlined">
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                        <path d="M10 4H4C2.89543 4 2 4.89543 2 6V18C2 19.1046 2.89543 20 4 20H20C21.1046 20 22 19.1046 22 18V8C22 6.89543 21.1046 6 20 6H12L10 4Z"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
                    </svg>
                </span>
                <span>Select Output Folder</span>
            </button>
            <button type="button" @click="processFiles"
                class="w-full flex min-w-[84px] max-w-full cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-6 bg-primary text-gray-900 text-base font-bold leading-normal tracking-[0.015em] hover:opacity-90 transition-opacity">
                <span class="truncate">Extract Pages</span>
            </button>
        </div>
    </section>
</template>

<script setup lang="ts">
import { ref } from "vue";
import DropZone from "@/components/DropZone.vue";
import FilesList from "@/components/FilesList.vue";
import Loading from "@/components/Loading.vue";
import { Upload } from "@/components/icons";
import type { FileItem } from "../types";
import { ProcessFiles, GetFolder } from "../../wailsjs/go/main/App";

const fileItems = ref<Array<FileItem>>([]);
const loading = ref<boolean>(false);

const handleFilesUploaded = (files: FileItem[]) => {
    fileItems.value = [...files];
    console.log("Files uploaded:", fileItems.value);
};

const removeFile = (id: string) => {
    fileItems.value = fileItems.value.filter((file) => file.id !== id);
};

const selectFolder = async () => {
    try {
        const folderPath = await GetFolder();
        console.log("Selected folder:", folderPath);
    } catch (error) {
        console.error("Error selecting folder:", error);
    }
};

const processFiles = async () => {
    try {
        loading.value = true;
        const result = await ProcessFiles("extractPages", "", 0);
        console.log("Files processed:", result);
        fileItems.value = [];
    } catch (error) {
        console.error("Error processing files:", error);
    } finally {
        loading.value = false;
    }
};
</script>
