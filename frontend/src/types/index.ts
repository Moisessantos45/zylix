export interface Tool {
  id: string;
  title: string;
  subtitle: string;
  description: string;
  icon: string;
  route: string;
}

export interface FileItem {
  id: string;
  path: string;
  name: string;
  extension: string;
}

export interface ToolOptions {
  format?: string;
  orientation?: string;
  outputFormat?: string;
  compressionLevel?: number;
  language?: string;
  [key: string]: any;
}
