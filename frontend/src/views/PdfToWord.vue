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
        <Loading :visible="loading" message="Converting PDF to Word" />
        <div class="flex flex-col items-center gap-3 p-4 text-center">
            <h1 class="text-charcoal text-4xl font-black leading-tight tracking-[-0.033em] md:text-5xl">
                Pdf to Word Converter
            </h1>
            <p class="text-subtle-text-light text-base font-normal leading-normal max-w-xl">
                Convert your PDF files into editable Word documents quickly and easily.
            </p>
        </div>
        <div class="mt-8 flex flex-col gap-8 p-4">
            <DropZone v-if="fileItems.length === 0" title="Drag & drop your PDF files here"
                subtitle="Supports PDF files. Max 10MB." :icon="Upload" display-name="PDFs (*.pdf)" pattern="*.pdf"
                @files-uploaded="handleFilesUploaded" />
            <FilesList v-if="fileItems.length > 0" :fileItems="fileItems" @remove-file="removeFile" />
            <div class="justify-center pt-4 flex flex-col gap-3 items-center">
                <div class="flex w-full flex-col gap-3 sm:flex-row sm:max-w-md">
                    <button type="button" @click="processFiles"
                        class="flex flex-1 cursor-pointer items-center justify-center overflow-hidden rounded-lg h-14 px-5 bg-primary text-charcoal text-base font-bold leading-normal tracking-[0.015em]">
                        <span class="material-symbols-outlined mr-2">
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" class="w-6 h-6">
                                <path d="M4 4a2 2 0 012-2h6l6 6v12a2 2 0 01-2 2H6a2 2 0 01-2-2V4z" stroke="currentColor"
                                    stroke-width="1.5" fill="none" />
                                <path d="M12 2v6h6" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"
                                    stroke-linejoin="round" />
                                <path d="M7 13h4M7 16h2" stroke="currentColor" stroke-width="1.5"
                                    stroke-linecap="round" />
                                <path d="M15 14l3 3m0 0l-3 3m3-3h-5" stroke="currentColor" stroke-width="1.5"
                                    stroke-linecap="round" stroke-linejoin="round" />
                                <rect x="18" y="13" width="4" height="3" rx="0.5" stroke="currentColor" stroke-width="1"
                                    fill="none" />
                            </svg>
                        </span>
                        <span class="truncate">
                            Convert to Word
                        </span>
                    </button>
                    <button type="button" @click="selectFolder"
                        class="flex flex-1 cursor-pointer items-center justify-center overflow-hidden rounded-lg h-14 px-5 bg-gray-200 text-charcoal text-base font-bold leading-normal tracking-[0.015em]">
                        <span class="material-symbols-outlined mr-2">
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" class="w-6 h-6">
                                <path d="M6 2a2 2 0 00-2 2v16a2 2 0 002 2h12a2 2 0 002-2V8l-6-6H6z"
                                    stroke="currentColor" stroke-width="1.5" fill="none" />
                                <path d="M14 2v6h6" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"
                                    stroke-linejoin="round" />
                                <path d="M9 13h6M9 16h4" stroke="currentColor" stroke-width="1.5"
                                    stroke-linecap="round" />
                            </svg>
                        </span>
                        <span class="truncate"> Select Folder </span>
                    </button>
                </div>
            </div>
        </div>
    </section>
</template>
<script setup lang="ts">
import { ref } from "vue";
import DropZone from "@/components/DropZone.vue";
import FilesList from "@/components/FilesList.vue";
import { Upload } from "@/components/icons";
import Loading from "@/components/Loading.vue";
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
    console.log("File removed. Remaining files:", fileItems.value);
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
        const result = await ProcessFiles("pdfToWord", "", 0);
        console.log("Files processed:", result);
        fileItems.value = [];
    } catch (error) {
        console.error("Error processing files:", error);
    } finally {
        loading.value = false;
    }
};
</script>
