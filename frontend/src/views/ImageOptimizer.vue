<template>
    <main class="w-full px-4 py-5 items-center flex flex-col will-change-scroll">
        <div class="w-full p-4 flex">
            <RouterLink to="/"
                class="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-slate-200 text-slate-700 font-medium hover:bg-slate-300 transition-colors">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                    <path d="M15 18l-6-6 6-6" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                        stroke-linejoin="round" />
                </svg>
                Back to Home
            </RouterLink>
        </div>

        <Loading :visible="loading" message="Optimizing images" />

        <div class="layout-content-container flex flex-col w-full max-w-4xl gap-4">
            <div class="flex flex-col items-center gap-2 text-center">
                <h1 class="text-4xl sm:text-5xl font-black tracking-tighter text-slate-900">Image Optimizer</h1>
                <p class="text-base text-slate-500 max-w-lg">
                    Effortlessly reduce image file sizes without losing quality. Perfect for web, email, and storage.
                </p>
            </div>

            <DropZone v-if="fileItems.length === 0" title="Drag & drop your image here"
                subtitle="Supports JPG, PNG, WebP, GIF. Max 10MB." :icon="Upload" display-name="Images (*.png;*.jpg)"
                pattern="*.png;*.jpg" @files-uploaded="useFile.handleFilesUploaded" />

            <FilesList v-if="fileItems.length > 0" :fileItems="fileItems" @remove-file="useFile.removeFile"
                @clear-all="useFile.clearFiles" />
            <button v-if="fileItems.length > 0" type="button"
                class="flex w-full items-center justify-center gap-2 rounded-xl h-12 px-6 bg-primary/10 text-primary border-2 border-primary border-dashed text-base font-medium hover:bg-primary/20 transition-colors">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                    <path d="M12 5v14m-7-7h14" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                        stroke-linejoin="round" />
                </svg>
                <span>Add More Images</span>
            </button>


            <div class="flex flex-col gap-6 rounded-xl border border-slate-200 bg-white p-6 sm:p-8">
                <h3 class="text-lg font-bold text-slate-900">Choose your optimization settings</h3>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
                    <div class="flex flex-col gap-3">
                        <h4 class="text-base font-bold text-slate-900">Compression Level</h4>
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
                    <!--
                    <div class="flex flex-col gap-3">
                        <h4 class="text-base font-bold text-slate-900">Output Format</h4>
                        <div class="relative">
                            <DropDown
                                :list="[
                                    {
                                        name: 'JPEG',
                                        value: 'jpeg',
                                    },
                                    {
                                        name: 'PNG',
                                        value: 'png',
                                    },
                                    {
                                        name: 'WebP',
                                        value: 'webp',
                                    },
                                    {
                                        name: 'Auto',
                                        value: 'auto',
                                    },
                                ]"
                            />
                        </div>
                    </div> -->
                </div>

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
            </div>
            <button type="button" @click="useFile.selectFolder"
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
                class="flex w-full items-center justify-center gap-2 rounded-xl h-12 px-6 bg-primary text-slate-900 text-base font-bold hover:opacity-90 transition-opacity disabled:opacity-50 disabled:cursor-not-allowed">
                <span class="material-symbols-outlined">bolt</span>
                <span>Optimize Image</span>
            </button>
        </div>
    </main>
</template>
<script setup lang="ts">
import { ref, onBeforeUnmount } from "vue";
import { storeToRefs } from "pinia";
import DropZone from "@/components/DropZone.vue";
import FilesList from "@/components/FilesList.vue";
import Loading from "@/components/Loading.vue";
import { Upload } from "@/components/icons";
import { ProcessFiles } from "../../wailsjs/go/main/App";
import useFileStore from "@/stores/file";

const useFile = useFileStore();
const { fileItems,isFirstUpload } = storeToRefs(useFile)

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

const processFiles = async () => {
    try {
        loading.value = true;
        if (quality.value <= 0 || quality.value > 100) {
            console.warn("No files to process.");
            return;
        }
        const result = await ProcessFiles("imgOptimizer", "", quality.value);
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
    isFirstUpload.value=true;
});
</script>
