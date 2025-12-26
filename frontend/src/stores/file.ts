import { ref } from "vue";
import { defineStore } from "pinia";
import { GetFolder, GetFiles, RemoveFile, ClearALl } from "../../wailsjs/go/main/App";
import type { FileItem } from "../types";

const useFileStore = defineStore("file", () => {
  const fileItems = ref<FileItem[]>([]);
  const isFirstUpload = ref<boolean>(true);
  const currentTool = ref<string>("")
  const params = ref<{ [key: string]: any }>({
    "pdf": {
      displayName: "PDFs (*.pdf)",
      pattern: ".pdf"
    },
    "image": {
      displayName: "Images (*.png;*.jpg)",
      pattern: "*.png;*.jpg"
    }
  })

  const handleFilesUploaded = (files: FileItem[]) => {
    fileItems.value = [...fileItems.value, ...files];
    isFirstUpload.value = false;
    console.log("Files uploaded:", fileItems.value);
  };

  const processFiles = (rawFiles: string[]) => {
    const newFileItems: FileItem[] = rawFiles.map((file, index) => ({
      id: `${Date.now()}-${index}`,
      path: file,
      name: file.split('/').pop() || 'unknown',
      extension: file.split('.').pop() || '',
    })).filter((item) => !fileItems.value.some(file => file.name == item.name));

    handleFilesUploaded(newFileItems);
  };


  const handleFileInput = async () => {
    try {
      if (!currentTool.value.trim()) return;
      const files = await GetFiles(params.value[currentTool.value].displayName, params.value[currentTool.value].pattern, isFirstUpload.value);
      console.info('Selected files:', files);
      processFiles(files);
    } catch (error) {
      console.warn('File selection was cancelled or failed.', error);
    }
  };


  const removeFile = async (id: string) => {
    const file = fileItems.value.find((item) => item.id === id);
    if (!file?.id) {
      return;
    }

    fileItems.value = fileItems.value.filter((file) => file.id !== id);
    console.info("File removed. Remaining files:", fileItems.value);
    try {
      await RemoveFile(file.path);
    } catch (error) {
      console.warn("ocurrrio un error en la eliminacion", error);
    }
  };

  const selectFolder = async () => {
    try {
      const folderPath = await GetFolder();
      console.info("Selected folder:", folderPath);
    } catch (error) {
      console.error("Error selecting folder:", error);
    }
  };

  const clearFiles = async () => {
    fileItems.value = [];
    isFirstUpload.value = true
    try {
      await ClearALl();
    } catch (error) {
      console.error("Error clearing files:", error);
    }
  };

  const clearAll = () => {
    fileItems.value = [];
    isFirstUpload.value = true;
    currentTool.value = ""
  }

  return {
    fileItems,
    isFirstUpload,
    currentTool,
    handleFilesUploaded,
    handleFileInput,
    removeFile,
    selectFolder,
    clearFiles,
    clearAll
  };
});

export default useFileStore;
