package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx    context.Context
	Files  []string
	Folder string
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetFolder() (string, error) {
	folder, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a folder",
	})

	if err != nil {
		return "", err
	}

	log.Println("Selected folder:", folder)
	a.Folder = folder

	return folder, nil
}

func (a *App) GetFiles(displayName string, patter string) ([]string, error) {
	files, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a folder",
		Filters: []runtime.FileFilter{
			{
				DisplayName: displayName,
				Pattern:     patter,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	log.Println("Selected files:", files)

	a.Files = files

	return files, err
}

func (a *App) ProcessFiles(tool string, name string, quality int) error {

	if len(a.Files) == 0 {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "No files selected",
			Message: "Please select files to process.",
		})

		return fmt.Errorf("no files selected")
	}

	if a.Folder == "" {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "No output folder selected",
			Message: "Please select an output folder.",
		})

		return fmt.Errorf("no output folder selected")
	}

	var err error

	switch tool {
	case "imgToPdf":
		err = a.ConvertImagesToPDF(name)
	case "imgOptimizer":
		// cmd = exec.Command(filepath.Join(execDir, "dist", "img_optimizer"), file, strconv.Itoa(quality), a.Folder)
		err = a.OptimizeImages(quality)
	case "pdfCompressor":
		// cmd = exec.Command(filepath.Join(execDir, "dist", "pdf_optimizer"), file, strconv.Itoa(quality), a.Folder)
		err = a.OptimizePDF(name)
	case "pdfToImg":
		// cmd = exec.Command(filepath.Join(execDir, "dist", "pdf_to_img"), file, a.Folder, name)
		err = a.ConvertPDFToImages(name)
	case "pdfMerge":
		// cmd = exec.Command(filepath.Join(execDir, "dist", "pdf_merge"), file, name, a.Folder)
		err = a.MergePDFs(name)
	case "extractPages":
		err = a.ExtractPDFPages()
		// cmd = exec.Command(filepath.Join(execDir, "dist", "pdf_to_docx"), file, a.Folder)
	default:
		return fmt.Errorf("unknown tool: %s", tool)
	}

	if err != nil {
		log.Println("Error processing files:", err)
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Processing error",
			Message: fmt.Sprintf("An error occurred while processing files: %v", err),
		})
		return err
	}

	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Processing complete",
		Message: "Files have been processed successfully.",
	})

	return nil
}

func (a *App) MergeFiles(files []string) string {

	var builder strings.Builder
	for i, file := range files {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(file)
	}

	result := builder.String()

	return result
}

func (a *App) RemoveFile(file string) error {
	var newListFiles []string
	log.Println("Esta entrando a la funcion", file)

	for _, item := range a.Files {
		if item == file {
			log.Println("el file es: ", item)
			continue
		}
		newListFiles = append(newListFiles, item)
	}

	a.Files = newListFiles

	return nil
}

func (a *App) ClearALl() {
	a.Files = nil
	a.Folder = ""
}
