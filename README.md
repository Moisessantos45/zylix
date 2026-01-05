# Zylix

<div align="center">

**Una suite completa de herramientas para manipulación de PDFs e imágenes**

[![Wails](https://img.shields.io/badge/Wails-v2-red?style=flat-square)](https://wails.io)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![Vue.js](https://img.shields.io/badge/Vue.js-3.x-4FC08D?style=flat-square&logo=vue.js)](https://vuejs.org)

</div>

## 📋 Descripción

**Zylix** es una aplicación de escritorio multiplataforma que proporciona un conjunto de herramientas profesionales para trabajar con archivos PDF e imágenes. Diseñada con una interfaz moderna e intuitiva, permite realizar operaciones complejas de manera simple y rápida.

## 📥 Descarga

### Versión Estable

| Versión | Plataforma | Descarga | Fecha |
|---------|-----------|----------|-------|
| v2.7.1 | Windows | [⬇️ Descargar](https://rmovevnbyamzdvslzqaq.supabase.co/storage/v1/object/public/apps/Programas/Zylix-v2_7_1.exe) | 2026-01-04 |
| v2.7 | Windows | [⬇️ Descargar](https://rmovevnbyamzdvslzqaq.supabase.co/storage/v1/object/public/apps/Programas/Zylix-v2_7.exe) | 2026-01-03 |
| v2.5 | Windows | [⬇️ Descargar](https://rmovevnbyamzdvslzqaq.supabase.co/storage/v1/object/public/apps/Programas/Zylix-v1_5.exe) | 2025-12-19 |

> **Nota**: Para otras plataformas o versiones anteriores, consulta la sección [Releases](https://github.com/Moisessantos45/zylix/releases).

### Instalación Rápida (Windows)

1. Descarga el archivo `.exe` desde el link anterior
2. Ejecuta el instalador
3. Sigue las instrucciones del asistente de instalación
4. ¡Listo! Ya puedes usar Zylix

## ✨ Características

### 📄 Herramientas PDF

- **Optimizar PDF**: Reduce el tamaño de archivos PDF manteniendo la calidad
- **Unir PDFs**: Combina múltiples archivos PDF en un solo documento
- **Separar PDF**: Extrae páginas individuales de un documento PDF
- **PDF a Imágenes**: Convierte páginas de PDF a imágenes de alta calidad
- **Imágenes a PDF**: Crea documentos PDF a partir de múltiples imágenes

### 🖼️ Herramientas de Imagen

- **Optimizar Imágenes**: Reduce el tamaño de las imágenes manteniendo la calidad visual
- **Conversión de Formatos**: Soporte para JPG, JPEG, PNG, BMP, TIFF y WEBP

## 🛠️ Stack Tecnológico

### Backend
- **[Go](https://golang.org)**: Lenguaje de programación principal
- **[Wails v2](https://wails.io)**: Framework para aplicaciones de escritorio
- **[pdfcpu](https://github.com/pdfcpu/pdfcpu)**: Biblioteca para manipulación de PDFs
- **[imaging](https://github.com/disintegration/imaging)**: Procesamiento de imágenes

### Frontend
- **[Vue.js 3](https://vuejs.org)**: Framework JavaScript progresivo
- **[TypeScript](https://www.typescriptlang.org)**: Tipado estático para JavaScript
- **[Vite](https://vitejs.dev)**: Build tool y dev server ultrarrápido

## 📦 Requisitos

- **Go** 1.21 o superior
- **Node.js** 16 o superior (o **Bun** como alternativa)
- **Wails CLI** v2

## 🚀 Instalación

### 1. Instalar Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 2. Clonar el repositorio

```bash
git clone https://github.com/Moisessantos45/zylix.git
cd zylix
```

### 3. Instalar dependencias

```bash
# Instalar dependencias de Go
go mod download

# Instalar dependencias del frontend
cd frontend
bun install  # o npm install
cd ..
```

## 🏃 Desarrollo

### Modo de desarrollo

Ejecuta la aplicación en modo de desarrollo con hot-reload:

```bash
wails dev
```

Esto iniciará:
- Un servidor de desarrollo Vite para el frontend
- La aplicación de escritorio con hot-reload
- Un servidor de desarrollo en http://localhost:34115 para debugging

### Características del modo desarrollo

- ✅ Hot reload automático del frontend
- ✅ DevTools del navegador disponibles
- ✅ Logs en tiempo real
- ✅ Acceso a métodos Go desde la consola del navegador

## 🔨 Compilación

### Compilar para producción

```bash
wails build
```

### Compilación optimizada

```bash
# Windows
wails build -clean -upx

# macOS
wails build -clean -upx -platform darwin/universal

# Linux
wails build -clean -upx -platform linux/amd64
```

Los archivos compilados se generarán en el directorio `build/bin/`.

## 📖 Uso

1. **Selecciona los archivos**: Haz clic en el botón correspondiente para seleccionar archivos PDF o imágenes
2. **Elige la herramienta**: Selecciona la operación que deseas realizar
3. **Configura opciones**: Ajusta parámetros como la calidad de compresión (si aplica)
4. **Selecciona carpeta de salida**: Indica dónde guardar los archivos procesados
5. **Procesa**: ¡Y listo! Tus archivos estarán listos en segundos

## 📁 Estructura del Proyecto

```
zylix/
├── frontend/          # Código del frontend (Vue.js)
│   ├── src/          # Código fuente
│   └── dist/         # Build de producción
├── build/            # Archivos de configuración de build
├── app.go            # Lógica principal de la aplicación
├── tool.go           # Implementación de herramientas
├── main.go           # Punto de entrada
├── wails.json        # Configuración de Wails
└── go.mod            # Dependencias de Go
```

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 👨‍💻 Autor

**Moisessantos45**
- Email: santosxphdz34@gmail.com
- GitHub: [@Moisessantos45](https://github.com/Moisessantos45)

## 🙏 Agradecimientos

- [Wails](https://wails.io) - Por el increíble framework
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) - Por la potente biblioteca de PDF
- [Vue.js](https://vuejs.org) - Por el framework frontend

---

<div align="center">
Hecho por Moisessantos45 usando Wails, Go y Vue.js
</div>
