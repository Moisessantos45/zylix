import { ref } from "vue";
import { defineStore } from "pinia";
import { GetFolder, RemoveFile, ClearALl } from "../../wailsjs/go/main/App";
import type { FileItem } from "../types";

const useFileStore = defineStore("file", () => {
  const fileItems = ref<FileItem[]>([]);
  const isFirstUpload = ref<boolean>(true);

  const handleFilesUploaded = (files: FileItem[]) => {
    fileItems.value = [...files];
    isFirstUpload.value=false;
    console.log("Files uploaded:", fileItems.value);
  };

  const removeFile = async (id: string) => {
    const file = fileItems.value.find((item) => item.id === id);
    if (!file?.id) {
      return;
    }

    fileItems.value = fileItems.value.filter((file) => file.id !== id);
    console.log("File removed. Remaining files:", fileItems.value);
    try {
      await RemoveFile(file.path);
    } catch (error) {
      console.warn("ocurrrio un error en la eliminacion", error);
    }
  };

  const selectFolder = async () => {
    try {
      const folderPath = await GetFolder();
      console.log("Selected folder:", folderPath);
    } catch (error) {
      console.error("Error selecting folder:", error);
    }
  };

  const clearFiles = async () => {
    fileItems.value = [];
    isFirstUpload.value=true
    try {
      await ClearALl();
    } catch (error) {
      console.error("Error clearing files:", error);
    }
  };

  return {
    fileItems,
    isFirstUpload,
    handleFilesUploaded,
    removeFile,
    selectFolder,
    clearFiles,
  };
});

export default useFileStore;
