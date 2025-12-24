import type { Tool } from "../types";

const tools: Tool[] = [
  {
    id: "img-to-pdf",
    title: "Image to PDF",
    subtitle: "Convert Images to PDF",
    description: "Convert your images into high-quality PDF documents easily.",
    icon: "ImageToPdf",
    route: "ImageToPdf",
  },
  {
    id: "image-optimizer",
    title: "Image Optimizer",
    subtitle: "Optimize Image Files",
    description:
      "Easily optimize images to reduce file size while maintaining quality.",
    icon: "ImageOptimizer",
    route: "ImageOptimizer",
  },
  {
    id: "pdf-to-image",
    title: "PDF to Image",
    subtitle: "Convert PDF Pages to Images",
    description:
      "Convert each page of your PDF documents into high-resolution images.",
    icon: "PdfToImage",
    route: "PdfToImage",
  },
  {
    id: "pdf-optimizer",
    title: "PDF Optimizer",
    subtitle: "Optimize PDF Files",
    description:
      "Reduce the size of your PDF files while preserving quality and content.",
    icon: "PdfOptimizer",
    route: "PdfOptimizer",
  },
  {
    id: "pdf-merger",
    title: "PDF Merger",
    subtitle: "Merge Multiple PDFs",
    description:
      "Combine multiple PDF documents into a single, organized file.",
    icon: "PdfMerger",
    route: "PdfMerger",
  },
  {
    id: "extract-pages-pdf",
    title: "Extract Pages from PDF",
    subtitle: "Extract specific pages from your PDF",
    description: "Extract each page from your PDF as a separate PDF file.",
    icon: "Scan",
    route: "ExtractPagesPdf",
  },
];

export { tools };
