<template>
    <section class="px-4 justify-center py-5 will-change-scroll">
        <div class="w-full p-4 flex">
            <RouterLink to="/" class="text-primary flex items-center gap-1 hover:underline">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                    <path d="M15 18l-6-6 6-6" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                        stroke-linejoin="round" />
                </svg>
                Back to Home
            </RouterLink>
        </div>
        <Loading :visible="loading" message="Optimizing PDFS" />
        <div class="layout-content-container flex flex-col items-center gap-4 w-full max-w-[960px] flex-1">
            <div class="flex flex-col items-center gap-3 text-center w-full max-w-2xl">
                <h1 class="text-gray-900 text-4xl md:text-5xl font-black leading-tight tracking-[-0.033em]">
                    Optimize PDF
                </h1>
                <p class="text-gray-600 text-base md:text-lg font-normal leading-normal">
                    Reduce the file size of your PDF while maintaining the best possible quality.
                </p>
            </div>
            <div v-if="fileItems.length === 0" class="flex flex-col p-4 w-full">
                <DropZone title="Drag & drop your PDF here" subtitle="Supports PDF files. Max 10MB." :icon="Upload"
                    display-name="PDFs (*.pdf)" pattern="*.pdf" @files-uploaded="handleFilesUploaded" />
            </div>

            <FilesList v-if="fileItems.length > 1" :fileItems="fileItems" @remove-file="removeFile" />

            <div class="flex px-4 py-3 w-full max-w-lg">
                <div class="flex flex-col gap-3 w-full">
                    <h4 class="text-base font-bold text-slate-900">Compression Level</h4>
                    <div class="flex flex-col gap-4 pt-2">
                        <div class="flex justify-between items-center">
                            <h4 class="text-base font-bold text-slate-900">Quality</h4>
                            <span class="text-sm font-medium text-slate-900 bg-slate-100 px-2 py-1 rounded-md">{{
                                quality
                                }}</span>
                        </div>
                        <input v-model="quality"
                            class="w-full h-2 bg-slate-200 rounded-full appearance-none cursor-pointer [&amp;::-webkit-slider-thumb]:appearance-none [&amp;::-webkit-slider-thumb]:h-4 [&amp;::-webkit-slider-thumb]:w-4 [&amp;::-webkit-slider-thumb]:rounded-full [&amp;::-webkit-slider-thumb]:bg-primary"
                            max="100" min="0" type="range" />
                    </div>
                    <div class="grid grid-cols-3 gap-2 rounded-lg bg-slate-100 p-1">
                        <button type="button" @click="handleQuanlityChange('low')" :class="[
                            quanlitySelected === 'low'
                                ? 'flex items-center justify-center rounded-md h-9 px-3 text-sm font-bold text-slate-900 bg-white'
                                : 'flex items-center justify-center rounded-md h-9 px-3 text-sm font-medium text-slate-600 hover:bg-white transition-opacity',
                        ]">
                            Low
                        </button>
                        <button type="button" @click="handleQuanlityChange('medium')" :class="[
                            quanlitySelected === 'medium'
                                ? 'flex items-center justify-center rounded-md h-9 px-3 text-sm font-bold text-slate-900 bg-white'
                                : 'flex items-center justify-center rounded-md h-9 px-3 text-sm font-medium text-slate-600 hover:bg-white transition-opacity',
                        ]">
                            Medium
                        </button>
                        <button type="button" @click="handleQuanlityChange('high')" :class="[
                            quanlitySelected === 'high'
                                ? 'flex items-center justify-center rounded-md h-9 px-3 text-sm font-bold text-slate-900 bg-white'
                                : 'flex items-center justify-center rounded-md h-9 px-3 text-sm font-medium text-slate-600 hover:bg-white transition-opacity',
                        ]">
                            High
                        </button>
                    </div>
                </div>
            </div>
            <div class="flex px-4 py-3 justify-center w-full max-w-lg flex-col gap-3">
                <button type="button" @click="selectFolder"
                    class="flex w-full items-center justify-center gap-2 rounded-xl h-12 px-6 bg-slate-200 text-slate-700 text-base font-medium hover:opacity-90 transition-opacity">
                    <span class="material-symbols-outlined">
                        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                            <path
                                d="M10 4H4C2.89543 4 2 4.89543 2 6V18C2 19.1046 2.89543 20 4 20H20C21.1046 20 22 19.1046 22 18V8C22 6.89543 21.1046 6 20 6H12L10 4Z"
                                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
                        </svg>
                    </span>
                    <span>Select Output Folder</span>
                </button>
                <button type="button" @click="processFiles"
                    class="flex w-full cursor-pointer items-center justify-center overflow-hidden rounded-xl h-12 px-5 bg-primary text-black text-base font-bold leading-normal tracking-[0.015em] hover:opacity-90 transition-opacity disabled:bg-gray-300 disabled:text-gray-500 disabled:cursor-not-allowed">
                    <span class="truncate">Optimize PDF</span>
                </button>
            </div>
        </div>
    </section>
</template>
<script setup lang="ts">
import { ref, onBeforeUnmount } from "vue";
import DropZone from "@/components/DropZone.vue";
import FilesList from "@/components/FilesList.vue";
import Loading from "@/components/Loading.vue";
import { Upload } from "@/components/icons";
import type { FileItem } from "../types";
import { ProcessFiles, GetFolder } from "../../wailsjs/go/main/App";

const fileItems = ref<Array<FileItem>>([]);
const quality = ref<number>(80);
const quanlitySelected = ref<string>("medium");
const loading = ref<boolean>(false);

const handleQuanlityChange = (value: string) => {
    quanlitySelected.value = value;
    switch (value) {
        case "low":
            quality.value = 40;
            break;
        case "medium":
            quality.value = 70;
            break;
        case "high":
            quality.value = 100;
            break;
    }
};

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
        if (quality.value <= 0 || quality.value > 100) {
            console.warn("No files to process.");
            return;
        }
        const result = await ProcessFiles("pdfCompressor", "", quality.value);
        console.log("Files processed:", result);
        fileItems.value = [];
        quality.value = 80;
    } catch (error) {
        console.error("Error processing files:", error);
    } finally {
        loading.value = false;
    }
};

onBeforeUnmount(() => {
    fileItems.value = [];
});
</script>
