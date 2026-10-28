# Design — UI, UX y sistema visual

> Documento de referencia para la interfaz de **ÍndicePartidas**. Define la dirección visual, los tokens, los componentes y las reglas de interacción del frontend. **Estado:** `v0.1 — En revisión` · **Última actualización:** `2026-10-28` · **Base de esta versión:** vista de **creación y edición de cotización** diseñada en Google Stitch, más las decisiones de producto del `PRD.md` y los requerimientos no funcionales. Las vistas de **registro de ejecución** y **comparación Cotizado vs. Ejecutado** se incorporarán a medida que se diseñen; sus patrones quedan definidos aquí como especificación objetivo.

## 1. Identidad visual del producto

| Campo | Definición |
| --- | --- |
| Nombre del producto | **ÍndicePartidas** |
| Personalidad visual | Técnica, profesional, sobria y accesible. Una “planilla técnica moderna”: precisa para trabajar con cifras, pero lo bastante simple para un contratista que hoy usa cuaderno o Excel. |
| Audiencia principal | Dueños y contratistas independientes, administradores de obra, jefes de terreno y maestros. Usuarios con baja tolerancia a la complejidad y alta necesidad de claridad financiera. |
| Sensación deseada | Confianza, control, orden, rapidez y transparencia. La interfaz debe transmitir que los números son confiables y que el estado del trabajo siempre está claro. |
| Referencias visuales | Vista “Technical Chapter Quotation” generada en Google Stitch (editor de cotización), herramientas de estimación de obra como Presto y Procore, y planillas técnicas de construcción. El diseño prioriza densidad de datos legible, jerarquía por capítulos y acentos verdes para dinero y acciones. |

**Rasgos de la identidad**

- **Grafito estructural:** la navegación, las barras de herramientas y el resumen económico usan superficies oscuras que ordenan la pantalla y separan el “marco” del “documento”.
- **Verde financiero:** el verde identifica acciones principales, selección, confirmaciones y cifras positivas. No se usa como decoración.
- **Documento blanco:** las tablas, fichas y vistas previas viven sobre superficies blancas con bordes finos, como una hoja de cálculo impresa.
- **Densidad controlada:** la información se agrupa por capítulos y se revela en niveles; no se muestra todo al mismo tiempo.
- **Cifras como protagonistas:** los importes se alinean a la derecha, usan cifras tabulares y reciben pesos tipográficos distintos según su importancia.

## 2. Principios de diseño

- **Claridad financiera primero:** en cada pantalla debe ser evidente qué se está viendo, si es costo o precio de venta, y cuál es la acción siguiente. Nunca mezclar ambos conceptos sin etiqueta.
- **Separación entre información interna y comercial:** costos, márgenes y utilidad son información confidencial. La vista para el cliente muestra solo lo que corresponde a la propuesta.
- **Captura rápida en terreno:** el registro de ejecución debe poder completarse en **30 segundos o menos** para una operación habitual. Menos campos, valores predeterminados y controles táctiles amplios.
- **Divulgación progresiva:** primero lo esencial —proyecto, partidas, total y guardar—; el detalle de APU, materiales, fórmulas y configuración se abre solo cuando se solicita.
- **Estados explícitos:** toda vista conectada a datos debe cubrir carga, vacío, error, éxito, permisos insuficientes y, cuando corresponda, falta de conexión.
- **Accesibilidad por defecto:** contraste, foco visible, navegación por teclado, etiquetas y objetivos táctiles adecuados desde el primer componente.
- **Consistencia sobre novedad:** no se crea un patrón nuevo si ya existe uno equivalente. Las variantes se resuelven con `variant` y `size`, no con props booleanas acumulativas.
- **Componentes reutilizables:** se compone antes de duplicar. Un componente específico solo se crea cuando tiene comportamiento propio.
- **Confianza y trazabilidad:** toda acción que modifique dinero o estado debe dejar claro qué cambió, cuándo y quién puede verlo.

## 3. Design tokens

> **Fuente canónica:** los valores de esta sección. La primera vista exportada desde Stitch usa aproximaciones en algunos casos (por ejemplo, radios de 4 px); al implementar en Vue se deben usar los tokens definidos aquí.

### Color

| Token | Uso | Valor |
| --- | --- | --- |
| `color.primary` | Acción principal, selección, navegación activa, enlaces destacados | `#15803D` |
| `color.primaryHover` | Hover de acciones principales | `#166534` |
| `color.primaryActive` | Estado presionado de acciones principales | `#14532D` |
| `color.primarySoft` | Fondos suaves, etiquetas y estados positivos secundarios | `#DCFCE7` |
| `color.onPrimarySoft` | Texto sobre fondos verdes suaves | `#14532D` |
| `color.secondary` | Superficies estructurales y acciones secundarias oscuras | `#18181B` |
| `color.secondaryRaised` | Controles sobre superficies oscuras | `#27272A` |
| `color.success` | Confirmaciones y estados validados | `#047857` |
| `color.successContainer` | Fondo de confirmaciones y badges aprobados | `#ECFDF5` |
| `color.successBorder` | Borde de badges aprobados | `#A7F3D0` |
| `color.warning` | Alertas preventivas, vigencias y pendientes | `#B45309` |
| `color.warningContainer` | Fondo de advertencias | `#FFFBEB` |
| `color.warningBorder` | Borde de advertencias | `#FDE68A` |
| `color.danger` | Errores, sobrecostos y acciones destructivas | `#B91C1C` |
| `color.dangerContainer` | Fondo de errores y desviaciones negativas | `#FEF2F2` |
| `color.dangerBorder` | Borde de errores y desviaciones | `#FECACA` |
| `color.info` | Información contextual y ayuda | `#1D4ED8` |
| `color.infoContainer` | Fondo de mensajes informativos | `#EFF6FF` |
| `color.background` | Fondo general de la aplicación | `#F4F4F5` |
| `color.surface` | Tarjetas, tablas, modales y documentos | `#FFFFFF` |
| `color.surfaceMuted` | Encabezados de tabla, pies y filas alternas | `#FAFAFA` |
| `color.surfaceInverse` | Sidebar, toolbar y barra de resumen | `#18181B` |
| `color.text` | Texto principal | `#18181B` |
| `color.textMuted` | Texto secundario, ayudas y metadatos | `#52525B` |
| `color.textSubtle` | Texto terciario, placeholders y marcas de agua | `#6D6D76` |
| `color.textInverse` | Texto sobre superficies oscuras | `#F4F4F5` |
| `color.textInverseMuted` | Texto secundario sobre superficies oscuras | `#A1A1AA` |
| `color.border` | Bordes decorativos de tablas, tarjetas y separadores | `#E4E4E7` |
| `color.borderStrong` | Bordes de énfasis y separadores de capítulo | `#D4D4D8` |
| `color.borderInput` | Bordes de campos y controles interactivos | `#8C8C95` |
| `color.focus` | Anillo de foco | `#15803D` |
| `color.overlay` | Fondo de modales y paneles | `rgba(24, 24, 27, 0.50)` |

**Reglas de color**

- El verde de acción (`primary`) y el verde de confirmación (`success`) pertenecen a la misma familia; por eso **el estado nunca se comunica solo con color**: siempre va acompañado de texto e icono.
- El rojo no se usa para “restar” en una tabla normal; se reserva para errores, sobrecostos y acciones destructivas.
- El grafito no es un color de texto cualquiera: se usa para estructurar la navegación y las barras persistentes.
- Las superficies oscuras no llevan sombras profundas; se separan por contraste tonal y bordes sutiles.

### Tipografía

| Token | Uso | Valor |
| --- | --- | --- |
| `font.family.base` | Texto general | `Inter, system-ui, sans-serif` |
| `font.feature.numeric` | Cifras, cantidades e importes | `'tnum' 1, 'zero' 1` (`font-variant-numeric: tabular-nums`) |
| `font.size.2xs` | Badges, atajos de teclado, microtexto | `10px / 14px` |
| `font.size.xs` | Etiquetas de columna y ayudas | `11px / 14px` |
| `font.size.sm` | Texto compacto y tablas densas | `12px / 16px` |
| `font.size.md` | Texto base de tablas y controles | `14px / 20px` |
| `font.size.lg` | Texto de lectura y formularios | `16px / 24px` |
| `font.size.xl` | Subtítulos y encabezados de sección | `20px / 28px` |
| `font.size.2xl` | Títulos de página | `24px / 32px` |
| `font.size.3xl` | Título principal / portada | `32px / 40px` |

| Rol | Peso | Notas |
| --- | --- | --- |
| `display` | 700 | 32 px, tracking `-0.02em`. Solo títulos de página y documentos. |
| `headline.lg` | 600 | 24 px, tracking `-0.015em`. Encabezados principales de vista. |
| `headline.md` | 600 | 20 px, tracking `-0.01em`. Secciones y modales. |
| `headline.sm` | 600 | 16–17 px. Encabezados de tarjeta y capítulo. |
| `body.lg` | 400 | 16 px. Texto de lectura, formularios y móvil. |
| `body.md` | 400 | 14 px. Texto base de escritorio y tablas. |
| `body.sm` | 400 | 12 px. Metadatos, ayudas y celdas secundarias. |
| `label` | 600 | 11 px, mayúsculas, tracking `0.06em`. Encabezados de columna y campos. |
| `numeric.sm` | 500 | 11 px, tabular. Totales secundarios y porcentajes. |
| `numeric.md` | 600 | 14 px, tabular. Importes y cantidades. |
| `numeric.lg` | 700 | 18 px, tabular. Total de venta y cifras protagonistas. |

**Reglas tipográficas**

- Todo número que se compare verticalmente —cantidades, precios, subtotales, márgenes, porcentajes— usa `tabular-nums`.
- Los importes se alinean a la derecha; las descripciones a la izquierda; las unidades y códigos al centro.
- En móvil, el texto de formularios y lectura no baja de **16 px** para evitar zoom automático y mejorar legibilidad.
- No se usan más de tres tamaños tipográficos por bloque visual.
- Los títulos no se escriben en mayúsculas sostenidas salvo las etiquetas de columna y badges.

### Espaciado, radios y sombras

| Token | Uso | Valor |
| --- | --- | --- |
| `space.1` | Separaciones mínimas | `4px` |
| `space.2` | Separaciones pequeñas | `8px` |
| `space.3` | Separación compacta | `12px` |
| `space.4` | Separación estándar | `16px` |
| `space.5` | Bloques medianos | `24px` |
| `space.6` | Bloques grandes | `32px` |
| `space.7` | Separaciones de página | `48px` |
| `radius.xs` | Detalles mínimos | `2px` |
| `radius.sm` | Inputs, botones y badges compactos | `4px` |
| `radius.md` | Controles estándar y celdas editables | `6px` |
| `radius.lg` | Tarjetas, tablas y paneles | `8px` |
| `radius.xl` | Modales y documentos | `12px` |
| `radius.full` | Avatares y píldoras excepcionales | `9999px` |
| `shadow.card` | Tarjetas elevadas | `0 1px 2px rgba(24, 24, 27, 0.06)` |
| `shadow.raised` | Popovers, menús y autocompletado | `0 4px 12px rgba(15, 23, 42, 0.08)` |
| `shadow.modal` | Modales y paneles flotantes | `0 20px 50px rgba(15, 23, 42, 0.25)` |
| `shadow.sticky` | Barra de resumen fija | `0 -2px 8px rgba(0, 0, 0, 0.08)` |

**Reglas de forma y profundidad**

- La interfaz se organiza en **capas planas**: fondo gris, documento blanco y marco grafito. Las sombras son sutiles y no decorativas.
- Los radios se mantienen contenidos; se evita el aspecto “burbuja” en una herramienta financiera.
- Un panel no lleva sombra si ya está delimitado por borde y contraste de superficie.

**Layout y breakpoints**

| Token | Uso | Valor |
| --- | --- | --- |
| `layout.sidebar` | Ancho de navegación lateral en escritorio | `288px` |
| `layout.header` | Altura de cabecera | `64px` |
| `layout.contentMax` | Ancho máximo del área de trabajo | `1440px` |
| `layout.marginMobile` | Margen lateral en móvil | `16px` |
| `layout.marginTablet` | Margen lateral en tablet | `24px` |
| `layout.marginDesktop` | Margen lateral en escritorio | `32px` |
| `layout.gutter` | Separación entre columnas | `16px` |
| `layout.gutterDesktop` | Separación entre columnas en escritorio | `24px` |

| Breakpoint | Rango | Comportamiento general |
| --- | --- | --- |
| `mobile` | `< 768px` | Una columna, navegación en panel desplegable, acciones prioritarias fijas, tablas transformadas en tarjetas o listas. |
| `tablet` | `768px – 1199px` | Navegación colapsable, tablas con scroll horizontal controlado y columnas secundarias ocultas. |
| `desktop` | `≥ 1200px` | Sidebar visible, tabla completa y barra de resumen persistente. |
| `wide` | `≥ 1440px` | Contenido centrado con ancho máximo; no se estira indefinidamente. |

## 4. Accesibilidad

El objetivo es cumplir **WCAG 2.2 AA** en las vistas principales.

**Contraste y color**

- Texto normal sobre fondo: contraste mínimo **4.5:1**. Texto grande y elementos gráficos: **3:1**.
- Ningún estado se comunica solo por color: los badges llevan texto; las desviaciones llevan signo, cifra e icono; los errores llevan mensaje.
- El foco visible usa un anillo de **2 px** en `color.focus` con separación de 2 px respecto del control.
- No se usan fondos verdes con texto blanco en tamaños pequeños sin verificar contraste.

**Teclado y navegación**

- Toda acción debe ser accesible por teclado, incluida la navegación lateral, los menús, las celdas editables y los modales.
- Orden de tabulación lógico: navegación → cabecera → contenido → resumen económico.
- En la tabla de cotización: `Tab` avanza de celda, `Enter` confirma la edición, `Escape` cancela y devuelve el foco a la celda original.
- Los atajos de teclado deben ser adicionales, nunca la única forma de ejecutar una acción.
- Se incluye un enlace “Saltar al contenido” visible al recibir foco.

**Formularios**

- Todo input tiene `label` visible. El placeholder no reemplaza la etiqueta.
- Los campos obligatorios se marcan con texto o icono, no solo con asterisco sin explicación.
- Los errores se muestran junto al campo, describen cómo corregir y se anuncian a lectores de pantalla.
- No se valida solo en el cliente: la interfaz muestra el error que devuelve el backend.
- Los formularios evitan el doble envío y muestran el estado `guardando` en la acción principal.

**Tablas**

- Se implementan con semántica de tabla (`table`, `thead`, `th`, `tbody`) o con roles ARIA equivalentes cuando se usen componentes virtualizados.
- Cada columna numérica declara su unidad o formato en el encabezado.
- Las filas editables mantienen una relación clara entre celda, cabecera y fila.
- En móvil no se fuerza el desplazamiento horizontal para completar tareas principales: la información se reorganiza en tarjetas.

**Modales y paneles**

- Atrapan el foco mientras están abiertos y lo devuelven al control que los abrió al cerrarse.
- Se cierran con `Escape` salvo que exista una operación en curso.
- El título del modal se asocia con `aria-labelledby`.
- Si el modal contiene un documento largo, el foco inicial va al encabezado y no al final del contenido.

**Objetivos táctiles y movimiento**

- En dispositivos táctiles, los controles tienen un área mínima de **44 × 44 px**, aunque el icono sea más pequeño.
- En escritorio se permiten controles compactos de 32–36 px, siempre con un área de interacción suficiente.
- Se respeta `prefers-reduced-motion`; las animaciones son breves y no bloquean la lectura.
- La interfaz debe funcionar con zoom al **200 %** sin pérdida de contenido ni acciones.

**Verificación de contraste de tokens**

Cálculo de luminancia relativa según **WCAG 2.2**, realizado el `2026-10-28` sobre los pares definidos en este documento. No reemplaza la revisión visual en navegador ni la verificación de estados con opacidad.

| Combinación | Ratio | Criterio aplicable | Resultado |
| --- | --- | --- | --- |
| Texto principal sobre fondo | 16,12:1 | Texto normal 4,5:1 | AAA |
| Texto secundario sobre superficie | 7,73:1 | Texto normal 4,5:1 | AAA |
| Texto terciario sobre fondo | 4,66:1 | Texto normal 4,5:1 | AA |
| Blanco sobre `primary` | 5,02:1 | Texto normal 4,5:1 | AA |
| Blanco sobre `success` | 5,48:1 | Texto normal 4,5:1 | AA |
| Blanco sobre `danger` | 6,47:1 | Texto normal 4,5:1 | AA |
| `success` sobre `successContainer` | 5,21:1 | Texto normal 4,5:1 | AA |
| `warning` sobre `warningContainer` | 4,84:1 | Texto normal 4,5:1 | AA |
| `danger` sobre `dangerContainer` | 5,91:1 | Texto normal 4,5:1 | AA |
| `info` sobre `infoContainer` | 6,16:1 | Texto normal 4,5:1 | AA |
| `onPrimarySoft` sobre `primarySoft` | 8,30:1 | Texto normal 4,5:1 | AAA |
| `textInverseMuted` sobre `surfaceInverse` | 6,91:1 | Texto normal 4,5:1 | AA |
| `focus` sobre superficie | 5,02:1 | Componentes gráficos 3:1 | Cumple |
| `borderInput` sobre superficie | 3,33:1 | Componentes gráficos 3:1 | Cumple |

**Hallazgos y correcciones**

- `color.textSubtle` se ajustó de `#71717A` a `#6D6D76`: sobre `color.background` pasaba de **4,40:1** a **4,66:1**.
- Se incorporó `color.borderInput` (`#8C8C95`) porque `color.border` y `color.borderStrong` no alcanzan 3:1. Los bordes decorativos pueden seguir siendo claros, pero **no deben ser el único indicio de un control interactivo**.
- Los pares restantes cumplen AA. La verificación de estados `hover`, `disabled`, `focus` y overlays queda pendiente en la implementación.

## 5. Arquitectura de componentes con Atomic Design

La implementación será en **Vue**. Los componentes de UI puros viven separados de los que consumen datos o estado global.

**Base técnica propuesta**

| Tema | Propuesta | Motivo | Estado |
| --- | --- | --- | --- |
| Base headless | **Reka UI** | Componentes accesibles para diálogos, menús, tabs, popovers y tooltips; compatible con Vue 3 y con foco/teclado resueltos. | Propuesta, pendiente de aprobación |
| Iconos | **Material Symbols** mediante Iconify (`@iconify/vue`) | Continuidad con la maqueta y carga por icono en lugar de la fuente completa. | Propuesta, pendiente de aprobación |
| Estilos | Tailwind CSS o CSS con custom properties | La maqueta usa Tailwind; los tokens deben sobrevivir a cualquiera de las dos opciones. | Por definir |
| Formato de datos | Formateador compartido para CLP, porcentajes y fechas | Evita duplicar reglas y permite testearlas. | Propuesta |

> La maqueta actual usa Tailwind vía CDN y Material Symbols como fuente. Ninguna de esas decisiones se considera definitiva para producción hasta cerrar esta tabla.

### Átomos

Componentes base indivisibles.

| Componente | Responsabilidad | Estados mínimos | Recomendación |
| --- | --- | --- | --- |
| `Button` | Ejecutar acciones | `default`, `hover`, `focus`, `active`, `disabled`, `loading` | Propio; variantes `primary`, `secondary`, `ghost`, `danger` |
| `IconButton` | Acción compacta con icono | `default`, `hover`, `focus`, `disabled`, `loading` | Propio; `aria-label` obligatorio y tooltip |
| `Input` | Capturar texto | `default`, `focus`, `error`, `disabled`, `readonly` | Propio; borde `color.borderInput` |
| `MoneyInput` | Capturar importes CLP | `default`, `focus`, `error`, `disabled` | Propio; alineación derecha, cifras tabulares y borde `color.borderInput` |
| `NumericInput` | Capturar cantidades y porcentajes | `default`, `focus`, `error`, `disabled` | Propio; teclado numérico en móvil y borde `color.borderInput` |
| `Select` | Elegir una opción | `default`, `focus`, `error`, `disabled` | Propio o headless; borde `color.borderInput` |
| `Label` | Nombrar campos | `default`, `required`, `disabled` | Propio |
| `Icon` | Mostrar pictogramas | `size`, `color` | Librería SVG coherente (Material Symbols o equivalente) |
| `Badge` | Comunicar estado breve | `success`, `warning`, `danger`, `neutral`, `info` | Propio; siempre con texto |
| `Spinner` | Carga breve | `sm`, `md`, `lg` | Propio |
| `Kbd` | Mostrar atajos | `default` | Propio |
| `Divider` | Separar bloques | `horizontal`, `vertical` | Propio |

### Moléculas

Combinaciones pequeñas de átomos.

| Componente | Responsabilidad | Composición sugerida |
| --- | --- | --- |
| `FormField` | Campo con etiqueta, ayuda y error | `Label + Input/NumericInput + HelpText + ErrorText` |
| `SearchBar` | Búsqueda textual | `Input + Icon + atajo opcional` |
| `Toast` | Notificación temporal | Cola + icono + mensaje + acción opcional |
| `Tabs` | Navegación entre secciones | Headless accesible + estilo propio |
| `DropdownMenu` | Acciones contextuales | Headless accesible |
| `CompanySwitcher` | Cambiar empresa activa | Avatar + nombre + RUT + menú |
| `UserMenu` | Contexto del usuario | Avatar + nombre + rol + acciones |
| `QuotationStatus` | Mostrar estado de cotización | `Badge` + texto + icono |
| `MoneyValue` | Mostrar un importe | Cifra tabular + unidad/moneda + énfasis |
| `DeltaValue` | Mostrar una desviación | Signo + monto + porcentaje + color/icono |
| `ChapterHeading` | Encabezar un capítulo | Código + título + contador + subtotal |
| `EmptyState` | Explicar una ausencia de datos | Icono + título + descripción + acción |
| `ConfirmDialog` | Confirmar acciones destructivas | Título + consecuencia + acciones |

### Organismos

Bloques funcionales de alto nivel.

| Componente | Responsabilidad | Decisión |
| --- | --- | --- |
| `AppSidebar` | Navegación principal y contexto de empresa | Propio; colapsable en tablet/móvil |
| `AppHeader` | Contexto, búsqueda, notificaciones y usuario | Propio; sticky |
| `QuotationHeaderStrip` | Ficha compacta de la cotización | Propio; cliente, obra, plazo, moneda y versión |
| `QuotationToolbar` | Acciones de edición de la hoja | Propio; insertar capítulo, agregar APU, importar y configurar columnas |
| `QuotationTable` | Matriz de partidas por capítulo | Propio; edición en celda y agrupación por capítulos |
| `EconomicSummaryBar` | Resumen fijo de costo, utilidad, neto, IVA y total | Propio; persistente en escritorio y expandible en móvil |
| `ClientPreviewModal` | Documento comercial sin costos internos | Propio; imprimible y accesible |
| `ExecutionForm` | Registro rápido de ejecución en terreno | Propio; móvil primero |
| `ComparisonTable` | Comparación Cotizado vs. Ejecutado | Propio; cifras y desviaciones por partida |
| `Modal` | Confirmaciones y formularios flotantes | Headless accesible |
| `Card` | Contenedor composable | Propio con subcomponentes |
| `FilterPanel` | Agrupar filtros de búsqueda | Propio |
| `DataTable` | Listados generales | Usar librería si requiere filtros u orden avanzado |

### Templates / layouts

| Layout | Uso | Reglas |
| --- | --- | --- |
| `AppLayout` | Marco general autenticado | Sidebar + header + contenido; sidebar colapsable en tablet/móvil |
| `QuotationEditorLayout` | Crear y editar cotización | Ficha compacta + toolbar + tabla por capítulos + resumen económico fijo |
| `ExecutionCaptureLayout` | Registrar ejecución en terreno | Contexto de OT + formulario breve + acción principal fija; prioridad móvil |
| `ComparisonLayout` | Análisis Cotizado vs. Ejecutado | Resumen de indicadores + tabla comparativa + detalle progresivo |
| `ClientDocumentLayout` | Vista previa y exportación comercial | Documento blanco, membrete, partidas y totales; sin controles internos |
| `AuthLayout` | Login y recuperación | Formulario centrado o dos columnas; sin navegación global |
| `CrudLayout` | Mantenedores | Título + acciones + filtros + tabla |
| `MasterDetailLayout` | Lista + detalle | Paneles con scroll independiente |

## 6. Reglas de composición

- **No crear componentes específicos sin comportamiento propio.** Si una tarjeta solo cambia textos o datos, se compone con `Card` y sus subcomponentes.
- **Evitar props booleanas acumulativas** como `isBig`, `isRed`, `hasBorder`, `compact` o `adminMode`. Preferir `variant`, `size`, `tone` y composición.
- **Separar UI pura de UI conectada a datos.** Los componentes de presentación reciben props y emiten eventos; los contenedores resuelven estado, permisos y llamadas a la API.
- **Los componentes que manejan foco, teclado, portales o colisiones de pantalla usan una base headless accesible** y reciben solo estilos propios.
- **El dinero se formatea en un único lugar.** Los componentes no concatenan símbolos ni separadores manualmente: reciben valores ya formateados o usan un formateador compartido (CLP, sin decimales por defecto).
- **Los números siempre usan cifras tabulares y alineación derecha.** Las descripciones se alinean a la izquierda y las unidades al centro.
- **La información interna se encapsula.** Costos, márgenes y utilidad se renderizan solo en componentes internos; la vista cliente usa componentes distintos o un modo explícito que no recibe esos datos.
- **La edición en tabla no debe sorprender.** Toda celda editable se distingue visualmente y confirma o cancela con `Enter` y `Escape`.
- **En móvil, la tarea principal no depende de scroll horizontal.** Las tablas se transforman en tarjetas, listas o formularios apilados.
- **La acción principal es única por pantalla.** Las acciones secundarias se agrupan en menús o botones de menor jerarquía.
- **Los formularios no pierden datos.** Antes de abandonar con cambios pendientes se muestra confirmación; el guardado no depende del navegador.
- **Los componentes de dominio no contienen reglas de negocio.** El cálculo de márgenes, IVA y totales vive en el backend o en un composable de dominio testeable, no en la plantilla.

## 7. Patrones de pantalla

### Listado CRUD

Debe incluir:

- Título y descripción breve.
- Acción primaria visible.
- Filtros principales y búsqueda.
- Tabla o lista con estados de carga, vacío y error.
- Paginación o scroll controlado.
- Confirmación para acciones destructivas.
- En móvil: tarjetas o filas apiladas; filtros en panel desplegable.

### Formulario

Debe incluir:

- Validación visible por campo, con mensaje orientado a la corrección.
- Mensaje general de error si falla el envío.
- Acción primaria y secundaria.
- Prevención de doble envío y estado `guardando`.
- Confirmación si se perderán cambios.
- Etiquetas visibles; ayudas solo cuando aportan.

### Detalle

Debe incluir:

- Estado actual de la entidad.
- Metadatos importantes y fecha de actualización.
- Acciones permitidas por rol.
- Historial o auditoría cuando corresponda.
- Información confidencial protegida y claramente marcada.

**Editor de cotización**

Es la pantalla de referencia del sistema.

**Escritorio**

- Ficha compacta con código, estado, cliente, obra, plazo, moneda, versión y fecha de vigencia.
- Toolbar con acciones de edición de la hoja.
- Tabla organizada por capítulos desplegables.
- Columnas principales: partida, unidad, cantidad, precio unitario, subtotal de venta y acciones.
- Columnas internas —costo y margen— visualmente diferenciadas y protegidas.
- Resumen económico fijo con costo, utilidad, neto, IVA y total cliente.
- Acción principal “Guardar cotización” y acción secundaria “Vista cliente”.

**Tablet**

- Sidebar colapsable.
- La tabla conserva scroll horizontal controlado y fija las columnas de descripción y total.
- Las columnas internas se ocultan por defecto y se consultan en el detalle de la partida.
- El resumen económico se mantiene visible, con menor altura.

**Móvil**

- Navegación en panel desplegable.
- Ficha de cotización resumida con acción “Editar datos”.
- Partidas en tarjetas agrupadas por capítulo; cada tarjeta muestra descripción, cantidad, subtotal y estado.
- La edición se realiza en un formulario o panel de pantalla completa, no en una tabla horizontal.
- El detalle de APU, costo y margen se abre bajo demanda.
- El resumen económico se reduce a total y acción de guardar; el desglose se expande.
- La vista cliente se abre como documento vertical imprimible.

**Captura de ejecución en terreno**

**Prioridad:** móvil.

- Contexto visible: proyecto, Orden de Trabajo y partida seleccionada.
- Fecha con valor predeterminado.
- Acciones rápidas: “Material”, “Mano de obra” y “Otro costo”.
- Solo los campos esenciales; el resto opcional.
- Cantidad ejecutada como acción diferenciada del registro de materiales u horas para evitar doble conteo.
- Nota y fotografía opcionales, con acceso directo a cámara.
- Acción principal fija “Guardar registro”.
- Confirmación inmediata con opción de deshacer.
- Estado sin conexión: indicar claramente si el registro quedó guardado localmente y ofrecer reintento.

**Comparación Cotizado vs. Ejecutado**

**Escritorio**

- Resumen superior: costo estimado, costo real, precio de venta, utilidad y margen.
- Comparación por partida con desviación en monto y porcentaje.
- Identificación visual de sobrecostos.
- Gráfico simple de costo estimado vs. real.
- Filtros por capítulo y acceso al detalle de cada partida.

**Móvil**

- Tarjetas de indicadores principales.
- Lista de partidas con desviación, ordenable por impacto.
- El gráfico se simplifica o se reemplaza por una comparación textual.
- El detalle de una partida se abre en panel completo.

**Vista previa comercial del cliente**

- Documento blanco con membrete de la empresa.
- Datos del cliente y proyecto.
- Partidas y totales **sin costos internos, márgenes ni utilidad**.
- Totales: subtotal neto, IVA y total.
- Acciones “Imprimir” y “Cerrar”; sin controles de edición.
- En móvil se presenta como documento vertical desplazable.

**Acciones destructivas**

| Acción | Cuándo se usa | Texto de confirmación | Acción principal | Acción secundaria |
| --- | --- | --- | --- | --- |
| Eliminar partida | Al quitar una partida de la cotización | “¿Eliminar la partida «{descripción}»? Se quitará de esta cotización y el total se recalculará.” | Eliminar partida | Cancelar |
| Eliminar capítulo | Al quitar un capítulo con partidas | “¿Eliminar el capítulo «{nombre}» y sus {n} partidas? Esta acción no se puede deshacer.” | Eliminar capítulo | Cancelar |
| Eliminar cotización | Solo en estado Borrador | “¿Eliminar la cotización {código}? Se perderán sus partidas y versiones.” | Eliminar cotización | Cancelar |
| Descartar cambios | Al salir con cambios sin guardar | “¿Salir sin guardar? Los cambios de esta cotización se perderán.” | Salir sin guardar | Seguir editando |
| Eliminar registro de ejecución | Al corregir un registro en terreno | “¿Eliminar este registro? El avance y los costos volverán a los valores anteriores.” | Eliminar registro | Cancelar |

**Reglas de confirmación**

- El foco inicial va en **Cancelar**, nunca en la acción destructiva.
- La acción destructiva usa `variant="danger"` y un verbo específico; no se usa “Aceptar” ni “OK”.
- El texto explica la consecuencia concreta, no solo pregunta “¿Estás seguro?”.
- Si la acción se puede deshacer, se ofrece “Deshacer” en lugar de una confirmación previa.
- No se confirma dos veces la misma acción.

**Nombres accesibles de controles de icono**

| Control | Nombre accesible |
| --- | --- |
| Expandir o contraer capítulo | “Expandir capítulo {nombre}” / “Contraer capítulo {nombre}” |
| Ver desglose de APU | “Ver desglose de APU de {partida}” |
| Eliminar partida | “Eliminar partida {descripción}” |
| Agregar partida al capítulo | “Agregar partida al capítulo {nombre}” |
| Menú de empresa | “Cambiar empresa activa. Empresa actual: {nombre}” |
| Notificaciones | “Notificaciones, {n} sin leer” |
| Ayuda | “Abrir ayuda” |
| Cerrar sesión | “Cerrar sesión” |
| Cerrar vista previa | “Cerrar vista previa de la cotización” |
| Imprimir | “Imprimir cotización {código}” |

**Reglas de nombres accesibles**

- El `aria-label` incluye el contexto: no basta con “Eliminar” o “Ver”.
- El texto visible del tooltip puede ser más breve, pero debe coincidir en intención con el nombre accesible.
- Los controles que solo muestran un icono siempre tienen nombre accesible y tooltip.
- Los estados de los controles se comunican con `aria-expanded`, `aria-pressed` o `aria-disabled` según corresponda, no solo con un cambio de color.

**Comportamiento responsivo por vista**

| Vista | Escritorio | Tablet | Móvil |
| --- | --- | --- | --- |
| Cotización | Tabla completa y resumen fijo | Tabla con scroll y columnas internas ocultas | Tarjetas por capítulo y edición en panel |
| Ejecución | Formulario con contexto lateral | Formulario de una columna | Formulario breve con acción fija |
| Comparación | Indicadores + tabla + gráfico | Indicadores + tabla resumida | Indicadores + tarjetas de partidas |
| Vista cliente | Modal amplio | Modal adaptado | Documento vertical |

## 8. Estados obligatorios

Cada pantalla o componente conectado a datos debe especificar:

| Estado | Significado | Tratamiento obligatorio |
| --- | --- | --- |
| `loading` | Los datos aún no están disponibles | Skeletons o spinner; no mostrar cifras vacías como si fueran cero |
| `empty` | No existen datos | Explicación breve y acción para crear el primer registro |
| `error` | Falló la carga o la operación | Mensaje claro, causa cuando se conozca y acción de reintento |
| `success` | La operación se completó | Confirmación breve; no bloquear la pantalla |
| `forbidden` | El rol no tiene acceso | Explicar la restricción sin revelar información sensible |
| `offline` | No hay conexión | Estado honesto de guardado local y opción de reintentar |
| `dirty` | Hay cambios sin guardar | Advertir antes de salir o cambiar de vista |

**Estados de cotización**

| Estado | Badge | Tratamiento |
| --- | --- | --- |
| Borrador | Neutro | Editable; no enviado al cliente |
| En revisión | Advertencia | Editable con precaución; registrar cambios |
| Aceptada | Éxito | Solo lectura para partidas base; puede generar Orden de Trabajo |
| Rechazada | Peligro | Solo lectura; permitir duplicar o crear nueva versión |
| Vencida | Advertencia | Advertir que los precios requieren actualización |

**Estados de ejecución**

| Estado | Significado | Tratamiento |
| --- | --- | --- |
| Pendiente de sincronizar | Guardado local | Mostrar icono y acción “Reintentar” |
| Sincronizado | Guardado en servidor | Confirmación discreta con hora |
| Error de sincronización | No se pudo guardar | Mantener el borrador local y explicar qué hacer |
| Validación | Datos incompletos o fuera de rango | Señalar el campo y no perder lo ingresado |

**Estados por pantalla**

**Editor de cotización**

| Estado | Texto y comportamiento |
| --- | --- |
| Cargando | Skeletons en ficha y tabla; totales ocultos. “Cargando cotización…” |
| Vacío | “Esta cotización aún no tiene partidas.” Acciones: “Agregar partida” e “Importar Excel”. |
| Error de carga | “No pudimos cargar la cotización.” Acción: “Reintentar”. |
| Error al guardar | Banner persistente: “No se guardaron los cambios. Revisa tu conexión.” Acción: “Reintentar”; los datos permanecen en pantalla. |
| Sin permiso de edición | Badge “Solo lectura”; acciones de edición deshabilitadas con explicación al enfocar. |
| Sin permiso de costos | Columnas internas, utilidad y margen no se renderizan; el resumen muestra solo el precio de venta. |
| Sin conexión | Acción “Guardar” deshabilitada con mensaje “Sin conexión. No se pueden guardar cambios.” |

**Registro de ejecución**

| Estado | Texto y comportamiento |
| --- | --- |
| Cargando | Skeleton del formulario y de los últimos registros. “Cargando registros…” |
| Vacío | “Esta partida aún no tiene registros.” Acción: “Registrar avance”. |
| Error al guardar | “No pudimos guardar el registro. Tus datos siguen aquí.” Acción: “Reintentar”. |
| Validación | “Ingresa una cantidad mayor que 0.” / “El costo debe ser mayor o igual a 0.” |
| Sin conexión | “Guardado en este dispositivo. Se enviará al recuperar conexión.” Acción: “Ver pendientes”. |
| Sin permiso de costos | El rol Maestro solo registra fotografías y notas; los campos de costo se ocultan. |

**Comparación Cotizado vs. Ejecutado**

| Estado | Texto y comportamiento |
| --- | --- |
| Cargando | Skeletons en indicadores y tabla; no mostrar cifras en $0. |
| Vacío | “La comparación estará disponible cuando existan costos registrados.” Acción: “Ir a ejecución”. |
| Error | “No pudimos calcular la comparación.” Acción: “Reintentar”. |
| Sin permiso | “No tienes permiso para ver costos internos.” Sin datos parciales que revelen márgenes. |
| Sin conexión | Mostrar últimos datos calculados con “Actualizado al {fecha}” y banner de conexión. |

**Vista previa comercial**

| Estado | Texto y comportamiento |
| --- | --- |
| Cargando | Spinner dentro del modal. “Generando vista previa…” |
| Vacío | “La cotización no tiene partidas para mostrar.” Acción: “Cerrar”. |
| Error | “No pudimos generar la vista previa.” Acción: “Reintentar”. |
| Sin permiso | La acción “Vista cliente” se oculta o deshabilita según el rol. |
| Sin conexión | Si los datos ya están cargados, permitir imprimir; si no, mostrar error y opción de reintentar. |

**Reglas generales de estado**

- Una cifra desconocida no se muestra como `$0`; se muestra como “Sin datos” o se oculta con explicación.
- Los errores de red no eliminan el formulario ni el trabajo en curso.
- Las confirmaciones no tapan la acción principal ni exigen un clic adicional innecesario.
- Los estados críticos se anuncian a tecnologías de asistencia mediante regiones `aria-live`.

## 9. Checklist de diseño antes de implementar

- [] Hay tokens definidos para color, tipografía, espaciado y radios.
- [] Los layouts principales están definidos.
- [] Cada flujo crítico tiene estado vacío, carga y error.
- [] Las acciones destructivas tienen confirmación.
- [] Los formularios tienen validación y mensajes claros.
- [] Los componentes complejos usan librerías headless cuando corresponde.
- [] La navegación por teclado fue considerada.
- [] El diseño móvil/escritorio está especificado según prioridad del producto.
