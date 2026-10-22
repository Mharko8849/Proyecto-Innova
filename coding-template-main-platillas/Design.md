# Design — UI, UX y sistema visual

> Plantilla reutilizable para definir dirección visual, experiencia de usuario, componentes y reglas frontend. Está pensada para trabajar con Atomic Design sin obligar a construir todo el catálogo antes del producto.

## 1. Identidad visual del producto

| Campo | Definición |
| --- | --- |
| Nombre del producto | `[Nombre]` |
| Personalidad visual | `[Formal / cálida / técnica / institucional / lúdica / premium / minimalista]` |
| Audiencia principal | `[Usuario objetivo]` |
| Sensación deseada | `[Confianza, rapidez, claridad, cercanía, eficiencia, etc.]` |
| Referencias visuales | `[Links, productos comparables, screenshots, moodboard]` |

## 2. Principios de diseño

- **Claridad primero:** la interfaz debe explicar el estado actual y la acción siguiente.
- **Accesibilidad por defecto:** contraste, foco visible, navegación por teclado y objetivos táctiles adecuados.
- **Componentes reutilizables:** preferir composición antes que crear variantes específicas por pantalla.
- **Estados explícitos:** cada vista debe cubrir carga, vacío, error, éxito y permisos insuficientes.
- **Consistencia sobre novedad:** no crear un patrón nuevo si ya existe uno equivalente en el sistema.

## 3. Design tokens

### Color

| Token | Uso | Valor |
| --- | --- | --- |
| `color.primary` | Acción principal, links destacados | `[#HEX]` |
| `color.secondary` | Acciones secundarias | `[#HEX]` |
| `color.success` | Confirmaciones | `[#HEX]` |
| `color.warning` | Alertas preventivas | `[#HEX]` |
| `color.danger` | Errores o acciones destructivas | `[#HEX]` |
| `color.surface` | Fondos de tarjetas | `[#HEX]` |
| `color.text` | Texto principal | `[#HEX]` |
| `color.textMuted` | Texto secundario | `[#HEX]` |

### Tipografía

| Token | Uso | Valor |
| --- | --- | --- |
| `font.family.base` | Texto general | `[Inter / system-ui / etc.]` |
| `font.size.xs` | Ayudas, captions | `[12px]` |
| `font.size.sm` | Texto compacto | `[14px]` |
| `font.size.md` | Texto base | `[16px]` |
| `font.size.lg` | Subtítulos | `[18px]` |
| `font.size.xl` | Títulos | `[24px+]` |

### Espaciado, radios y sombras

| Token | Uso | Valor |
| --- | --- | --- |
| `space.1` | Separaciones mínimas | `[4px]` |
| `space.2` | Separaciones pequeñas | `[8px]` |
| `space.4` | Separación estándar | `[16px]` |
| `space.6` | Bloques medianos | `[24px]` |
| `radius.sm` | Inputs, badges | `[6px]` |
| `radius.md` | Cards, modales | `[12px]` |
| `shadow.card` | Tarjetas elevadas | `[valor]` |

## 4. Accesibilidad

- Cumplir WCAG 2.2 AA cuando aplique.
- Todo input debe tener `label` visible o accesible.
- Todo control interactivo debe tener estado `focus-visible`.
- No depender sólo del color para comunicar estado.
- Objetivos táctiles mínimos: `44px x 44px`.
- Modales y menús flotantes deben atrapar/liberar foco correctamente.

## 5. Arquitectura de componentes con Atomic Design

### Átomos

Componentes base indivisibles.

| Componente | Responsabilidad | Estados mínimos | Recomendación |
| --- | --- | --- | --- |
| `Button` | Ejecutar acciones | `default`, `hover`, `focus`, `disabled`, `loading` | Propio |
| `Input` | Capturar texto | `default`, `focus`, `error`, `disabled` | Propio |
| `Label` | Nombrar campos | `default`, `required`, `disabled` | Propio |
| `Icon` | Mostrar pictogramas | `size`, `color` | Librería SVG |
| `Badge` | Comunicar estado breve | `success`, `warning`, `danger`, `neutral` | Propio |
| `Spinner` | Carga breve | `sm`, `md`, `lg` | Propio |

### Moléculas

Combinaciones pequeñas de átomos.

| Componente | Responsabilidad | Composición sugerida |
| --- | --- | --- |
| `FormField` | Campo con label, ayuda y error | `Label + Input + HelpText + ErrorText` |
| `SearchBar` | Búsqueda textual | `Input + Icon + optional Button` |
| `Toast` | Notificación temporal | Librería de cola + estilo propio |
| `Tabs` | Navegación entre secciones | Headless accesible + estilo propio |
| `DropdownMenu` | Acciones contextuales | Headless accesible |

### Organismos

Bloques funcionales de alto nivel.

| Componente | Responsabilidad | Decisión |
| --- | --- | --- |
| `DataTable` | Listar, ordenar y paginar datos | Usar librería si requiere filtros/orden avanzado |
| `Modal` | Confirmaciones y formularios flotantes | Headless accesible |
| `Sidebar` | Navegación principal | Propio |
| `Header` | Contexto, usuario y acciones globales | Propio |
| `FilterPanel` | Agrupar filtros de búsqueda | Propio |
| `Card` | Contenedor composable | Propio con subcomponentes |

### Templates / layouts

| Layout | Uso | Reglas |
| --- | --- | --- |
| `BaseLayout` | Sitios públicos | Header + contenido + footer |
| `DashboardLayout` | Backoffice | Sidebar + topbar + área scrollable |
| `AuthLayout` | Login/registro | Formulario centrado o dos columnas |
| `CrudLayout` | Mantenedores | Título + acciones + filtros + tabla |
| `MasterDetailLayout` | Lista + detalle | Paneles con scroll independiente |
| `FullscreenLayout` | Flujos inmersivos | Sin navegación global |

## 6. Reglas de composición

- `CardProduct`, `CardUser`, `CardEvent`, etc. sólo se crean si tienen comportamiento propio; si no, usar `<Card>` composable.
- Evitar props booleanas acumulativas tipo `isBig`, `isRed`, `hasBorder`, `compact`, `adminMode`. Preferir `variant`, `size` y composición.
- Si un componente maneja foco, teclado, portales o colisiones de pantalla, preferir librería headless.
- Separar componentes UI puros de componentes conectados a datos.

## 7. Patrones de pantalla

### Listado CRUD

Debe incluir:

- Título y descripción breve.
- Acción primaria visible.
- Filtros principales.
- Tabla/lista con estados de carga, vacío y error.
- Paginación o scroll controlado.
- Confirmación para acciones destructivas.

### Formulario

Debe incluir:

- Validación visible por campo.
- Mensaje general de error si falla el submit.
- Acción primaria y secundaria.
- Prevención de doble envío.
- Confirmación si se perderán cambios.

### Detalle

Debe incluir:

- Estado actual de la entidad.
- Metadatos importantes.
- Acciones permitidas por rol.
- Historial o auditoría si aplica.

## 8. Estados obligatorios

Cada pantalla o componente conectado a datos debe especificar:

- `loading`
- `empty`
- `error`
- `success`
- `forbidden`
- `offline` si aplica

## 9. Checklist de diseño antes de implementar

- [ ] Hay tokens definidos para color, tipografía, espaciado y radios.
- [ ] Los layouts principales están definidos.
- [ ] Cada flujo crítico tiene estado vacío, carga y error.
- [ ] Las acciones destructivas tienen confirmación.
- [ ] Los formularios tienen validación y mensajes claros.
- [ ] Los componentes complejos usan librerías headless cuando corresponde.
- [ ] La navegación por teclado fue considerada.
- [ ] El diseño móvil/escritorio está especificado según prioridad del producto.
