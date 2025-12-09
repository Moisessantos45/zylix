!define ARG_WAILS_AMD64_BINARY "Zylix.exe"
!define ARG_WAILS_ARM64_BINARY "Zylix.exe"
Unicode true
!include "wails_tools.nsh"
!include "UMUI.nsh"
!include "nsDialogs.nsh"

# The version information for this two must consist of 4 parts
VIProductVersion "2.1.0.0"
VIFileVersion    "2.2.0.0"

VIAddVersionKey "CompanyName"     "${INFO_COMPANYNAME}"
VIAddVersionKey "FileDescription" "${INFO_PRODUCTNAME} Installer"
VIAddVersionKey "ProductVersion"  "${INFO_PRODUCTVERSION}"
VIAddVersionKey "FileVersion"     "${INFO_PRODUCTVERSION}"
VIAddVersionKey "LegalCopyright"  "${INFO_COPYRIGHT}"
VIAddVersionKey "ProductName"     "${INFO_PRODUCTNAME}"

# Enable HiDPI support
ManifestDPIAware true

# Configuración de colores y tema
!define UMUI_ICON "..\\icon.ico"          # Para el instalador
!define UMUI_UNICON "..\\icon.ico"        # Para el desinstalador
!define MUI_ICON "..\\icon.ico"           # También para compatibilidad MUI

# Configuración de comportamiento
!define MUI_ABORTWARNING
!define MUI_FINISHPAGE_NOAUTOCLOSE
!define UMUI_ABORTWARNING
!define UMUI_FINISHPAGE_NOAUTOCLOSE

# Personalización de la página de bienvenida
!define MUI_WELCOMEPAGE_TITLE "Bienvenido al instalador de Zylix"
!define MUI_WELCOMEPAGE_TEXT "Gracias por instalar Zylix, tu suite de herramientas PDF local y offline.$\r$\n$\r$\n\
Con Zylix tendrás las herramientas fundamentales de gestión de documentos: convertir imágenes a PDF, PDF a imágenes, optimizar archivos y mucho más, todo sin necesidad de internet.$\r$\n$\r$\n\
Haz clic en 'Siguiente' para continuar."

# Personalización de la página final
!define MUI_FINISHPAGE_TITLE "Instalación Completada"
!define MUI_FINISHPAGE_TEXT "Zylix se ha instalado correctamente en tu sistema.$\r$\n$\r$\n\
Desarrollado por Mmabitec para la gestión profesional de tu negocio."
!define MUI_FINISHPAGE_RUN "$INSTDIR\\${PRODUCT_EXECUTABLE}"
!define MUI_FINISHPAGE_RUN_TEXT "Ejecutar Zylix ahora"

#### Páginas del instalador ####
# Usa MUI_PAGE_* para páginas estándar (NO UMUI_PAGE_*)
!insertmacro MUI_PAGE_WELCOME
!insertmacro MUI_PAGE_DIRECTORY
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_PAGE_FINISH

# Páginas del desinstalador
!insertmacro MUI_UNPAGE_CONFIRM
!insertmacro MUI_UNPAGE_INSTFILES

# Idioma
!insertmacro MUI_LANGUAGE "Spanish"

Name "${INFO_PRODUCTNAME}"
OutFile "..\\..\\bin\\${INFO_PROJECTNAME}-${ARCH}-installer.exe"
InstallDir "$PROGRAMFILES64\\${INFO_COMPANYNAME}\\${INFO_PRODUCTNAME}"
ShowInstDetails show

Function .onInit
   !insertmacro wails.checkArchitecture
FunctionEnd

Section "Instalación"
    !insertmacro wails.setShellContext
    !insertmacro wails.webview2runtime

    SetOutPath $INSTDIR
    !insertmacro wails.files
    
    # Incluir binarios adicionales de la carpeta dist
    CreateDirectory "$INSTDIR\dist"
    SetOutPath "$INSTDIR\dist"
    File "..\..\bin\dist\pdf_to_img.exe"
    
    # Volver al directorio principal
    SetOutPath $INSTDIR

    # Crear accesos directos
    CreateShortcut "$SMPROGRAMS\\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\\${PRODUCT_EXECUTABLE}"
    CreateShortcut "$DESKTOP\\${INFO_PRODUCTNAME}.lnk" "$INSTDIR\\${PRODUCT_EXECUTABLE}"

    !insertmacro wails.associateFiles
    !insertmacro wails.associateCustomProtocols
    !insertmacro wails.writeUninstaller
SectionEnd

Section "Uninstall"
    !insertmacro wails.setShellContext

    # Eliminar datos de usuario
    RMDir /r "$AppData\\${INFO_PRODUCTNAME}"

    # Eliminar archivos del programa
    RMDir /r "$INSTDIR"

    # Eliminar accesos directos
    Delete "$SMPROGRAMS\\${INFO_PRODUCTNAME}.lnk"
    Delete "$DESKTOP\\${INFO_PRODUCTNAME}.lnk"

    !insertmacro wails.unassociateFiles
    !insertmacro wails.unassociateCustomProtocols
    !insertmacro wails.deleteUninstaller
SectionEnd
