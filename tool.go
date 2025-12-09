package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func (a *App) ConvertImagesToPDF(name string) error {

	var images []string

	for _, img := range a.Files {
		img = strings.TrimSpace(img)
		if esImagenValida(filepath.Base(img)) {
			images = append(images, img)
		}
	}

	if len(images) == 0 {
		return fmt.Errorf("no se encontraron imágenes válidas para convertir")
	}

	fmt.Printf("Encontradas %d imagen(es)\n", len(images))

	pdfPath := filepath.Join(a.Folder, name)

	os.MkdirAll(a.Folder, 0755)

	err := api.ImportImagesFile(images, pdfPath, nil, nil)
	if err != nil {
		return fmt.Errorf("error convirtiendo a PDF: %v", err)
	}

	fmt.Printf("✅ PDF creado: %s\n", name)
	fmt.Println("🎉 Conversión completada!")

	return nil
}

func esImagenValida(nombre string) bool {
	ext := strings.ToLower(filepath.Ext(nombre))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".bmp" || ext == ".tiff" || ext == ".webp"
}

func esPDFValido(nombre string) bool {
	ext := strings.ToLower(filepath.Ext(nombre))
	return ext == ".pdf"
}

func nombreSinExtension(ruta string) string {
	return strings.TrimSuffix(filepath.Base(ruta), filepath.Ext(ruta))
}

func fileInfo(path string) float64 {
	info, _ := os.Stat(path)
	return float64(info.Size()) / 1024 / 1024
}

func (a *App) OptimizeImages(quality int) error {
	if quality == 0 {
		quality = 85
	}

	for _, imgPath := range a.Files {
		nameWithoutExt := nombreSinExtension(imgPath)
		outputPath := filepath.Join(a.Folder, nameWithoutExt+"_opt.jpg")

		srcImg, err := imaging.Open(imgPath)
		if err != nil {
			fmt.Printf("⚠️ Error %s: %v\n", filepath.Base(imgPath), err)
			continue
		}

		dstImg := imaging.Resize(srcImg, 0, 2048, imaging.Lanczos)

		f, _ := os.Create(outputPath)
		jpeg.Encode(f, dstImg, &jpeg.Options{Quality: quality})
		f.Close()

		origSize := fileInfo(imgPath)
		newSize := fileInfo(outputPath)
		fmt.Printf("✅ %s → %.1fMB (antes %.1fMB, -%.0f%%)\n",
			filepath.Base(imgPath), newSize, origSize, (1-newSize/origSize)*100)
	}
	return nil
}

func (a *App) ConvertPDFToImages(name string) error {
	var cmd *exec.Cmd
	var file string = a.MergeFiles(a.Files)

	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error getting executable path: %v", err)
	}
	execDir := filepath.Dir(execPath)

	cmd = exec.Command(filepath.Join(execDir, "dist", "pdf_to_img"), file, a.Folder, name)

	output, err := cmd.CombinedOutput()
	if err != nil {

		log.Printf("Error processing file %s: %v", file, err)
		return err
	}

	log.Printf("Output for file %s: %s", file, string(output))

	return nil
}

func (a *App) ExtractPDFPages() error {
	for _, file := range a.Files {
		if !esPDFValido(filepath.Base(file)) {
			continue
		}

		nameWithoutExt := nombreSinExtension(file)
		pathSubFolder := filepath.Join(a.Folder, nameWithoutExt)
		os.MkdirAll(pathSubFolder, 0755)

		err := api.ExtractPagesFile(file, pathSubFolder, nil, nil)
		if err != nil {
			fmt.Printf("⚠️ Error %s: %v\n", filepath.Base(file), err)
			continue
		}

		fmt.Printf("✅ %s → SVGs en %s/\n", filepath.Base(file), pathSubFolder)
	}
	return nil
}

func (a *App) OptimizePDF(name string) error {
	if len(a.Files) == 0 {
		return fmt.Errorf("no hay PDFs para optimizar")
	}

	for _, file := range a.Files {
		if !esPDFValido(filepath.Base(file)) {
			continue
		}

		nameWithoutExt := nombreSinExtension(file)
		optimizedPath := filepath.Join(a.Folder, nameWithoutExt+"_optimizado.pdf")

		fmt.Printf("Optimizando %s → %s\n", filepath.Base(file), optimizedPath)

		err := api.OptimizeFile(file, optimizedPath, nil)
		if err != nil {
			fmt.Printf("⚠️ Error %s: %v\n", filepath.Base(file), err)
			continue
		}

		origSize := fileInfo(file)
		newSize := fileInfo(optimizedPath)
		fmt.Printf("✅ %s → %.1fMB (antes %.1fMB, -%.1f%%)\n",
			filepath.Base(file), newSize, origSize,
			(1-newSize/origSize)*100)
	}
	return nil
}

func (a *App) MergePDFs(outputName string) error {

	var validPDFs []string
	for _, file := range a.Files {
		if esPDFValido(filepath.Base(file)) {
			validPDFs = append(validPDFs, file)
		}
	}

	outputPath := filepath.Join(a.Folder, outputName)
	os.MkdirAll(a.Folder, 0755)

	fmt.Printf("Uniendo %d PDFs → %s\n", len(validPDFs), outputPath)

	err := api.MergeCreateFile(validPDFs, outputPath, false, nil)
	if err != nil {
		return fmt.Errorf("error uniendo: %v", err)
	}

	fmt.Printf("✅ %d PDFs → %s\n", len(validPDFs), outputPath)
	return nil
}
