Revisa este archivo Vue y elimina o reemplaza las clases CSS que causan lag en el scroll:

CLASES PROBLEMÁTICAS (eliminar o reemplazar):
- shadow-sm, shadow-md, shadow-lg, shadow-xl (especialmente en listas con v-for)
- shadow-[color]/opacity (ej: shadow-primary/20)
- transition-shadow
- transition-colors (en muchos elementos)
- hover:shadow-* 
- backdrop-blur-*
- filter, blur-*

REEMPLAZOS RECOMENDADOS:
- shadow-* → eliminar o usar solo en hover de elementos individuales
- transition-shadow → transition-opacity o transition-transform
- hover:shadow-lg → hover:opacity-90 o hover:bg-gray-50
- transition-colors → transition-opacity (cuando sea posible)

AGREGAR PARA OPTIMIZAR SCROLL:
- will-change-scroll en contenedores con overflow-y-auto
- will-change-transform en elementos animados

REGLAS:
1. Las transiciones de opacity y transform son aceleradas por GPU
2. Las sombras con colores personalizados son muy costosas
3. Evitar sombras en elementos dentro de v-for con muchos items
4. backdrop-blur causa repaint constante

Aplica estos cambios manteniendo un diseño visual atractivo.