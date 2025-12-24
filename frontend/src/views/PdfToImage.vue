<template>
    <section class="w-full px-4 py-5 items-center flex flex-col will-change-scroll">
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
        <Loading :visible="loading" message="Converting PDF to Images" />
        <div class="flex flex-wrap justify-between gap-3 p-4">
            <div class="flex min-w-72 flex-col gap-3 text-center w-full items-center">
                <p class="text-gray-900 text-4xl font-black leading-tight tracking-[-0.033em]">
                    Convert Your PDF to High-Quality Images
                </p>
                <p class="text-gray-600 text-base font-normal leading-normal max-w-2xl">
                    Instantly turn any PDF document into JPG or PNG files online.
                </p>
            </div>
        </div>
        <div v-if="fileItems.length === 0" class="flex flex-col p-4 w-full max-w-[960px] flex-1">
            <DropZone title="Drag & drop your PDF here" subtitle="Supports PDF files. Max 10MB." :icon="Upload"
                display-name="PDFs (*.pdf)" pattern="*.pdf" @files-uploaded="useFile.handleFilesUploaded" />
        </div>

        <FilesList v-if="fileItems.length > 0" :fileItems="fileItems" @remove-file="useFile.removeFile"
            @clear-all="useFile.clearFiles" />
        <button v-if="fileItems.length > 0" type="button"
            class="flex w-full items-center justify-center gap-2 rounded-xl h-12 px-6 bg-primary/10 text-primary border-2 border-primary border-dashed text-base font-medium hover:bg-primary/20 transition-colors mx-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                <path d="M12 5v14m-7-7h14" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                    stroke-linejoin="round" />
            </svg>
            <span>Add More PDFs</span>
        </button>

        <div class="flex flex-col gap-4 p-4 w-full max-w-[960px] flex-1">
            <h3 class="text-gray-900 text-lg font-bold leading-tight tracking-[-0.015em] px-0 pb-2 pt-4">
                Configuration
            </h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="flex flex-col gap-3 p-6 rounded-xl bg-white border border-gray-200">
                    <label class="text-sm font-medium text-gray-700">Output Format</label>
                    <div class="grid grid-cols-2 gap-2 p-1 bg-gray-100 rounded-lg">
                        <button type="button" @click="handleFormatChange('JPG')" :class="[
                            'px-3 py-2 text-sm font-semibold rounded-md text-gray-900',
                            formatSelected === 'JPG' ? 'bg-white text-gray-900' : 'text-gray-600',
                        ]">
                            JPG
                        </button>
                        <button type="button" @click="handleFormatChange('PNG')" :class="[
                            'px-3 py-2 text-sm font-semibold rounded-md text-gray-900',
                            formatSelected === 'PNG' ? 'bg-white text-gray-900' : 'text-gray-600',
                        ]">
                            PNG
                        </button>
                    </div>
                </div>
                <!-- <div class="flex flex-col gap-3 p-6 rounded-xl bg-white border border-gray-200">
                    <div class="flex justify-between items-center">
                        <label class="text-sm font-medium text-gray-700">Image Quality</label>
                        <span class="text-sm font-semibold text-gray-900">High</span>
                    </div>
                    <div class="flex items-center gap-4">
                        <input class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer accent-primary"
                            max="100" min="0" type="range" value="80" />
                    </div>
                </div> -->
            </div>
        </div>
        <div class="p-4 flex flex-col gap-3 w-full max-w-[960px] flex-1">
            <button type="button" @click="useFile.selectFolder"
                class="flex w-full items-center justify-center gap-2 rounded-xl h-12 px-6 bg-slate-200 text-slate-700 text-base font-medium hover:opacity-90 transition-opacity">
                <span class="material-symbols-outlined">
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                        <path
                            d="M10 4H4C2.89543 4 2 4.89543 2 6V18C2 19.1046 2.89543 20 4 20H20C21.1046 20 22 19.1046 22 18V8C22 6.89543 21.1046 6 20 6H12L10 4Z"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
                    </svg></span>
                <span>Select Output Folder</span>
            </button>
            <button type="button" @click="processFiles"
                class="w-full flex min-w-[84px] max-w-full cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-6 bg-primary text-gray-900 text-base font-bold leading-normal tracking-[0.015em] hover:opacity-90 transition-opacity">
                <span class="truncate">Convert to Image</span>
            </button>
        </div>
    </section>
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
const { fileItems, isFirstUpload } = storeToRefs(useFile)

const formatSelected = ref<string>("JPG");
const loading = ref<boolean>(false);

const handleFormatChange = (format: string) => {
    formatSelected.value = format;
};

const processFiles = async () => {
    try {
        loading.value = true;
        const result = await ProcessFiles("pdfToImg", formatSelected.value, 0);
        console.log("Files processed:", result);
        fileItems.value = [];
    } catch (error) {
        console.error("Error processing files:", error);
    } finally {
        loading.value = false;
    }
};

onBeforeUnmount(() => {
    fileItems.value = [];
    isFirstUpload.value = true;
});
</script>
