<template>
    <section class="w-full px-4 py-5 will-change-scroll">
        <div class="w-full p-4 flex">
            <RouterLink to="/" class="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-slate-200 text-slate-700 font-medium hover:bg-slate-300 transition-colors">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                    <path d="M15 18l-6-6 6-6" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                        stroke-linejoin="round" />
                </svg>
                Back to Home
            </RouterLink>
        </div>
        <Loading :visible="loading" message="Optimizing images" />
        <div class="text-center mb-8">
            <h1 class="text-charcoal text-4xl font-black leading-tight tracking-[-0.033em]">Convert Images to PDF</h1>
            <p class="text-gray-500 text-base md:text-lg font-normal leading-normal mt-3 max-w-2xl mx-auto">
                Create high-quality PDF documents from your images in seconds. Drag, drop, customize, and convert with
                ease.
            </p>
        </div>
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 items-start w-full mx-auto">
            <div class="lg:col-span-2 space-y-6">
                <DropZone v-if="fileItems.length === 0" title="Drag & drop images here"
                    subtitle="Supports JPG, PNG, WebP, GIF. Max 10MB." :icon="Upload"
                    display-name="Images (*.png;*.jpg)" pattern="*.png;*.jpg"
                    @files-uploaded="useFile.handleFilesUploaded" />
                <FilesList :expanded="true" v-if="fileItems.length > 0" :fileItems="fileItems"
                    @remove-file="useFile.removeFile" @clear-all="useFile.clearFiles" />
                <button v-if="fileItems.length > 0" type="button"
                    class="flex w-full items-center justify-center gap-2 rounded-xl h-12 px-6 bg-primary/10 text-primary border-2 border-primary border-dashed text-base font-medium hover:bg-primary/20 transition-colors">
                    <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                        <path d="M12 5v14m-7-7h14" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                            stroke-linejoin="round" />
                    </svg>
                    <span>Add More Images</span>
                </button>
            </div>
            <div class="lg:col-span-1 bg-white p-6 rounded-xl border border-gray-100 space-y-6">
                <h3
                    class="text-charcoal text-lg font-bold leading-tight tracking-[-0.015em] pb-2 border-b border-light-gray">
                    Conversion Settings
                </h3>
                <div class="space-y-2">
                    <label class="text-sm font-medium text-charcoal">Page Orientation</label>
                    <div class="flex flex-wrap gap-3">
                        <label
                            class="flex-1 text-sm font-medium leading-normal flex items-center justify-center rounded-lg border border-gray-300 px-4 h-11 text-charcoal relative cursor-pointer">
                            <input checked class="invisible absolute" name="orientation" type="radio" />
                            Portrait
                        </label>
                        <label
                            class="flex-1 text-sm font-medium leading-normal flex items-center justify-center rounded-lg border border-gray-300 px-4 h-11 text-charcoal relative cursor-pointer">
                            <input class="invisible absolute" name="orientation" type="radio" />
                            Landscape
                        </label>
                    </div>
                </div>
                <div class="space-y-2">
                    <label class="text-sm font-medium text-charcoal" for="page-size">Page Size</label>
                    <DropDown :list="[
                        {
                            name: 'A4',
                            value: 'a4',
                        },
                        {
                            name: 'Letter',
                            value: 'letter',
                        },
                        {
                            name: 'Legal',
                            value: 'legal',
                        },
                        {
                            name: 'A3',
                            value: 'a3',
                        },
                    ]" :disabled="true" />
                </div>
                <div class="space-y-2">
                    <label class="text-sm font-medium text-charcoal" for="margins">Margins</label>
                    <DropDown :list="[
                        {
                            name: 'None',
                            value: 'none',
                        },
                        {
                            value: 'small',
                            name: 'Small',
                        },
                        {
                            value: 'normal',
                            name: 'Normal',
                        },
                        {
                            value: 'large',
                            name: 'Large',
                        },
                    ]" :disabled="true" />
                </div>
                <div class="flex flex-col gap-2 w-full max-w-md">
                    <label class="text-sm font-medium text-charcoal" for="name">Output File Name</label>
                    <input v-model="namePdf"
                        class="w-full h-11 px-4 bg-white border border-gray-200 rounded-lg text-gray-900 placeholder:text-gray-400 focus:border-primary focus:ring-2 focus:ring-primary/20 focus:outline-none transition-opacity"
                        id="name" type="text" placeholder="merged-document" />
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
                    class="w-full flex min-w-[84px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-4 bg-primary text-charcoal text-base font-bold tracking-[0.015em] hover:opacity-90 transition-opacity disabled:opacity-50 disabled:cursor-not-allowed">
                    <span class="truncate">Convert to PDF</span>
                </button>
            </div>
        </div>
    </section>
</template>
<script setup lang="ts">
import { ref } from "vue";
import { storeToRefs } from "pinia";
import DropZone from "@/components/DropZone.vue";
import FilesList from "@/components/FilesList.vue";
import DropDown from "@/components/DropDown.vue";
import Loading from "@/components/Loading.vue";
import { Upload } from "@/components/icons";
import { ProcessFiles } from "../../wailsjs/go/main/App";
import useFileStore from "@/stores/file";

const useFile = useFileStore();
const { fileItems } = storeToRefs(useFile)

const namePdf = ref<string>("");
const loading = ref<boolean>(false);

const processFiles = async () => {
    try {
        if (namePdf.value.trim() === "") {
            console.warn("No files to process.");
            return;
        }
        loading.value = true;
        const result = await ProcessFiles("imgToPdf", namePdf.value, 0);
        console.log("Files processed:", result);
        fileItems.value = [];
        namePdf.value = "";
    } catch (error) {
        console.error("Error processing files:", error);
    } finally {
        loading.value = false;
    }
};
</script>
