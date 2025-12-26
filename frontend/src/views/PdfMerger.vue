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
        <Loading :visible="loading" message="Merging PDFs" />
        <div class="flex flex-col items-center gap-3 p-4 text-center">
            <h1 class="text-charcoal text-4xl font-black leading-tight tracking-[-0.033em] md:text-5xl">
                Merge Your PDFs Instantly
            </h1>
            <p class="text-subtle-text-light text-base font-normal leading-normal max-w-xl">
                Combine multiple PDF files into a single document. Simply drag, drop, reorder, and merge.
            </p>
        </div>
        <div class="mt-8 flex flex-col gap-8 p-4 w-full max-w-[960px] flex-1">
            <DropZone v-if="fileItems.length === 0" title="Drag & drop your PDF files here"
                subtitle="Supports PDF files. Max 10MB." :icon="Upload" display-name="PDFs (*.pdf)" pattern="*.pdf" />
            <FilesList v-if="fileItems.length > 0" :fileItems="fileItems" @remove-file="useFile.removeFile"
                @clear-all="useFile.clearFiles" />
            <button v-if="fileItems.length > 0" type="button" @click="useFile.handleFileInput"
                class="flex w-full items-center justify-center gap-2 rounded-xl h-12 px-6 bg-primary/10 text-primary border-2 border-primary border-dashed text-base font-medium hover:bg-primary/20 transition-colors">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                    <path d="M12 5v14m-7-7h14" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                        stroke-linejoin="round" />
                </svg>
                <span>Add More PDFs</span>
            </button>
            <div class="justify-center pt-4 flex flex-col gap-3 items-center">
                <div class="flex flex-col gap-2 w-full max-w-md">
                    <label class="text-sm font-medium text-charcoal" for="name">Output File Name</label>
                    <input v-model="namePdf" class="w-full h-11 px-4 bg-white border border-gray-200 rounded-lg text-gray-900 
                               placeholder:text-gray-400 focus:border-primary focus:ring-2 focus:ring-primary/20 
                               focus:outline-none transition-opacity" id="name" type="text"
                        placeholder="merged-document" />
                </div>
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
                        <span class="truncate">Merge PDFs</span>
                    </button>
                    <button type="button" @click="useFile.selectFolder"
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
import { ref, onMounted, onBeforeUnmount } from "vue";
import { storeToRefs } from "pinia";
import DropZone from "@/components/DropZone.vue";
import FilesList from "@/components/FilesList.vue";
import { Upload } from "@/components/icons";
import Loading from "@/components/Loading.vue";
import { ProcessFiles } from "../../wailsjs/go/main/App";
import useFileStore from "@/stores/file";

const useFile = useFileStore();
const { fileItems, currentTool } = storeToRefs(useFile)

const namePdf = ref<string>("");
const loading = ref<boolean>(false);

const processFiles = async () => {
    try {
        if (namePdf.value.trim() === "") {
            console.warn("No files to process.");
            return;
        }
        loading.value = true;
        const result = await ProcessFiles("pdfMerge", namePdf.value, 0);
        console.info("Files processed:", result);
        fileItems.value = [];
        namePdf.value = "";
    } catch (error) {
        console.error("Error processing files:", error);
    } finally {
        loading.value = false;
    }
};

onMounted(() => {
    currentTool.value = "pdf"
})

onBeforeUnmount(() => {
    useFile.clearAll();
});
</script>
