# Design — UI, UX y sistema visual

> Documento de referencia para la interfaz de **ÍndicePartidas**. Define la dirección visual, los tokens, los componentes y las reglas de interacción del frontend. **Estado:** `v0.2 — En revisión` · **Última actualización:** `2026-11-09` · **Base de esta versión:** especificación textual del dashboard **Inicio** de [Figma](https://www.figma.com/design/JoXshj0r0VakO5EXiBxW9i/Sin-t%C3%ADtulo?node-id=2-3), página **Page 1**, nodo `2:3`, capa de código de **1512 × 982 px**, proporcionada por el equipo. La referencia describe el lienzo de escritorio de **Cristales Santiago Demo**, rubro **Vidriería**; colores, tamaños y espaciados son estimaciones visuales, sin extracción de tokens ni inspección directa del archivo en esta revisión. Las demás pantallas, estados ocultos, interacciones, implementación y comportamiento en ejecución siguen sin verificar; sus definiciones se mantienen como propuestas basadas en el `PRD.md` y los requerimientos no funcionales.

## 1. Identidad visual del producto

| Campo | Definición |
| --- | --- |
| Nombre del producto | **ÍndicePartidas** |
| Personalidad visual | Administrativa, profesional, serena y accesible. Un entorno claro para gestionar cotizaciones, trabajos y costos, con densidad moderada y separación legible entre indicadores, catálogo, acciones y actividad. |
| Audiencia principal | Dueños y contratistas independientes, administradores de obra, jefes de terreno y maestros. Usuarios con baja tolerancia a la complejidad y alta necesidad de claridad financiera. |
| Sensación deseada | Confianza, control, orden y claridad financiera. La navegación blanca, el espacio de trabajo azul grisáceo muy claro y la jerarquía azul marino permiten identificar el estado del negocio y la siguiente acción. |
| Referencias visuales | Dashboard Inicio del mockup de Figma descrito por el equipo: cuenta ficticia **Cristales Santiago Demo**, rubro **Vidriería**, textos en español y marca **ÍndicePartidas — COSTOS CON EVIDENCIA**. Esta vista establece la nueva referencia visual; las otras pantallas requieren validación con sus propios frames. |

**Rasgos de la identidad**

- **Marco claro:** sidebar y cabecera blancas sobre un espacio de trabajo azul grisáceo muy claro; bordes finos separan la navegación del contenido.
- **Jerarquía azul:** el azul marino identifica marca, encabezados, valores destacados y acción principal; el azul de acento identifica enlaces y selección. El verde y el ámbar comunican estados con texto.
- **Tarjetas suaves:** superficies blancas, esquinas redondeadas, bordes pálidos y sombras discretas agrupan la información sin elevación pesada.
- **Densidad controlada:** el dashboard separa saludo, cuatro métricas, catálogo del rubro, acciones rápidas y actividad reciente. El detalle técnico de otras vistas se revela progresivamente.
- **Cifras como protagonistas:** importes y valores se destacan en azul marino. En tablas se propone alineación derecha y cifras tabulares; los ejemplos conservan el formato español con puntos de miles, como `$355.000`.

## 2. Principios de diseño

- **Claridad financiera primero:** en cada pantalla debe ser evidente qué se está viendo, si es costo o precio de venta, y cuál es la acción siguiente. Nunca mezclar ambos conceptos sin etiqueta.
- **Separación entre información interna y comercial:** costos, márgenes y utilidad son información confidencial. La vista para el cliente muestra solo lo que corresponde a la propuesta.
- **Captura rápida en terreno:** el registro de ejecución debe poder completarse en **30 segundos o menos** para una operación habitual. Menos campos, valores predeterminados y controles táctiles amplios.
- **Divulgación progresiva:** Inicio prioriza resumen del negocio y acciones frecuentes; cada flujo muestra primero lo esencial y abre el detalle de APU, materiales, fórmulas y configuración bajo demanda.
- **Estados explícitos:** toda vista conectada a datos debe cubrir carga, vacío, error, éxito, permisos insuficientes y, cuando corresponda, falta de conexión.
- **Accesibilidad por defecto:** contraste, foco visible, navegación por teclado, etiquetas y objetivos táctiles adecuados desde el primer componente.
- **Consistencia sobre novedad:** tarjetas, filas de acción, badges e iconos comparten lenguaje visual. Las variantes se resuelven con `variant` y `size`, no con props booleanas acumulativas; hover, foco, presionado y deshabilitado requieren definición.
- **Componentes reutilizables:** se compone antes de duplicar. Un componente específico solo se crea cuando tiene comportamiento propio.
- **Confianza y trazabilidad:** toda acción que modifique dinero o estado debe dejar claro qué cambió, cuándo y quién puede verlo.

## 3. Design tokens

> **Referencia provisional:** los valores de esta sección trasladan las estimaciones visuales del dashboard de Figma; no son tokens extraídos del archivo. Hover, presionado, errores, bordes auxiliares, alturas de línea, sombras y medidas sin referencia explícita son propuestas de implementación. Confirmar valores, tipografía y contraste antes de aprobarlos como tokens de marca para React + JavaScript.

### Color

| Token | Uso | Valor |
| --- | --- | --- |
| `color.primary` | Azul marino estimado: botón principal, marca y panel promocional | `#153F6A` |
| `color.primaryHover` | Propuesta de hover; no visible en la referencia | `#12365A` |
| `color.primaryActive` | Propuesta de estado presionado; no visible en la referencia | `#102F50` |
| `color.primarySoft` | Azul claro estimado: navegación seleccionada y contenedores de iconos | `#EAF3FB` |
| `color.onPrimarySoft` | Azul marino para texto de navegación seleccionada | `#153F6A` |
| `color.secondary` | Azul secundario estimado: fondo de ilustración del catálogo | `#1D527F` |
| `color.secondaryRaised` | Azul de acento estimado: enlaces, etiquetas e indicador activo | `#287ACA` |
| `color.success` | Verde estimado: textos de “Prueba activa” y “Finalizado”; contraste por ajustar | `#23966E` |
| `color.successContainer` | Fondo verde pálido estimado para prueba activa y finalización | `#ECFAF3` |
| `color.successBorder` | Borde de éxito propuesto; no medido en la referencia | `#A7F3D0` |
| `color.warning` | Ámbar estimado: texto de “En ejecución”; contraste por ajustar | `#CF7B16` |
| `color.warningContainer` | Fondo ámbar pálido estimado para trabajo en ejecución | `#FFF5E7` |
| `color.warningBorder` | Borde de advertencia propuesto; no medido en la referencia | `#FDE2BC` |
| `color.danger` | Propuesta para errores, sobrecostos y acciones destructivas | `#B91C1C` |
| `color.dangerContainer` | Propuesta de fondo de error; estado no visible en el dashboard | `#FEF2F2` |
| `color.dangerBorder` | Propuesta de borde de error | `#FECACA` |
| `color.info` | Azul de acento estimado: enlaces, etiquetas superiores y ayuda | `#287ACA` |
| `color.infoContainer` | Azul claro para información contextual | `#EAF3FB` |
| `color.background` | Azul grisáceo muy claro estimado del espacio de trabajo | `#F4F7FA` |
| `color.surface` | Tarjetas, tablas, modales y documentos | `#FFFFFF` |
| `color.surfaceMuted` | Propuesta para encabezados de tabla y filas alternas | `#F4F7FA` |
| `color.surfaceInverse` | Panel promocional marino; sidebar y cabecera permanecen blancas | `#153F6A` |
| `color.text` | Azul marino estimado: encabezados y valores destacados | `#173A5E` |
| `color.textMuted` | Azul grisáceo estimado: etiquetas, descripciones y metadatos | `#6B84A0` |
| `color.textSubtle` | Reutilización provisional del tono secundario para texto terciario | `#6B84A0` |
| `color.textInverse` | Blanco para botón principal y texto del panel marino | `#FFFFFF` |
| `color.textInverseMuted` | Propuesta de texto secundario sobre superficies marinas | `#DCE7F0` |
| `color.border` | Borde pálido estimado de tarjetas y divisores | `#DCE7F0` |
| `color.borderStrong` | Propuesta para separadores de mayor énfasis | `#C7D8E8` |
| `color.borderInput` | Propuesta de borde perceptible para campos y controles | `#6B84A0` |
| `color.focus` | Propuesta de anillo de foco azul | `#287ACA` |
| `color.overlay` | Propuesta de fondo de modales y paneles | `rgba(23, 58, 94, 0.50)` |

**Reglas de color**

- El azul marino identifica la acción principal; el azul de acento identifica enlaces y selección. Verde y ámbar expresan estados, siempre acompañados de texto y, cuando corresponda, icono.
- El rojo no se usa para “restar” en una tabla normal; se reserva para errores, sobrecostos y acciones destructivas.
- Sidebar, cabecera y tarjetas usan blanco; el azul marino se concentra en marca, títulos, cifras y promoción del catálogo.
- Los tonos estimados son una base de fidelidad visual; los pares con contraste insuficiente deben ajustarse antes de aprobar la implementación. Evitar sombras profundas y elevación pesada.

### Tipografía

| Token | Uso | Valor |
| --- | --- | --- |
| `font.family.base` | Sans-serif; Inter se conserva como propuesta, familia del mockup por confirmar | `Inter, system-ui, sans-serif` |
| `font.feature.numeric` | Cifras, cantidades e importes | `'tnum' 1, 'zero' 1` (`font-variant-numeric: tabular-nums`) |
| `font.size.2xs` | Etiquetas superiores en mayúsculas; tamaño estimado, interlineado propuesto | `10px / 14px` |
| `font.size.xs` | Metadatos pequeños y badges; tamaño estimado, interlineado propuesto | `11px / 16px` |
| `font.size.sm` | Metadatos y badges; tamaño estimado, interlineado propuesto | `12px / 16px` |
| `font.size.md` | Navegación, botones, etiquetas y cuerpo; tamaño estimado | `14px / 20px` |
| `font.size.lg` | Propuesta para lectura y formularios móviles | `16px / 24px` |
| `font.size.xl` | Encabezados de sección y métricas; tamaño estimado | `20px / 28px` |
| `font.size.2xl` | Encabezado promocional y valores destacados; tamaño estimado | `24px / 32px` |
| `font.size.3xl` | Saludo de Inicio, estimado entre 28–30 px; interlineado propuesto | `30px / 38px` |

| Rol | Peso | Notas |
| --- | --- | --- |
| `display` | 700 propuesto | Saludo de Inicio de aproximadamente 28–30 px, negrita; tracking por confirmar. |
| `headline.lg` | 700 propuesto | Encabezado promocional de aproximadamente 24 px, negrita. |
| `headline.md` | 600–700 por confirmar | Encabezados de sección de aproximadamente 20–24 px. |
| `headline.sm` | 600 propuesto | Encabezados secundarios de 16–18 px como propuesta para vistas pendientes. |
| `body.lg` | 400 | 16 px. Texto de lectura, formularios y móvil. |
| `body.md` | 400 propuesto | Aproximadamente 14 px para navegación, botones, etiquetas y cuerpo de escritorio. |
| `body.sm` | 400 propuesto | Aproximadamente 11–12 px para metadatos y badges. |
| `label` | 600 propuesto | Aproximadamente 10–11 px, mayúsculas y tracking ampliado; `0.06em` provisional. |
| `numeric.sm` | 500 | 11 px, tabular. Totales secundarios y porcentajes. |
| `numeric.md` | 600 | 14 px, tabular. Importes y cantidades. |
| `numeric.lg` | 600–700 por confirmar | Aproximadamente 20–24 px para los valores de las cuatro métricas; cifras tabulares propuestas. |

**Reglas tipográficas**

- Todo número que se compare verticalmente —cantidades, precios, subtotales, márgenes, porcentajes— usa `tabular-nums`.
- En tablas, los importes se alinean a la derecha, las descripciones a la izquierda y las unidades y códigos al centro. Las tarjetas de Inicio mantienen su lectura a la izquierda y los valores destacados.
- En móvil, el texto de formularios y lectura no baja de **16 px** para evitar zoom automático y mejorar legibilidad.
- No se usan más de tres tamaños tipográficos por bloque visual.
- Los títulos usan azul marino y las descripciones un tono secundario; las mayúsculas se reservan para etiquetas como “VIDRIERÍA”, “RESUMEN DEL NEGOCIO”, “TU RUBRO” y “COSTOS CON EVIDENCIA”.

### Espaciado, radios y sombras

| Token | Uso | Valor |
| --- | --- | --- |
| `space.1` | Incremento menor propuesto para relaciones entre textos | `4px` |
| `space.2` | Base propuesta del ritmo de espaciado | `8px` |
| `space.3` | Separación compacta | `12px` |
| `space.4` | Separación estándar | `16px` |
| `space.5` | Padding interior estimado de las tarjetas | `24px` |
| `space.6` | Bloques grandes | `32px` |
| `space.7` | Separaciones de página | `48px` |
| `radius.xs` | Propuesta para detalles pequeños no medidos | `4px` |
| `radius.sm` | Propuesta para controles compactos | `8px` |
| `radius.md` | Botones y filas de acción, estimados entre 10–12 px | `12px` |
| `radius.lg` | Tarjetas principales, estimadas entre 16–18 px | `16px` |
| `radius.xl` | Variante de tarjeta redondeada; uso en modales por validar | `18px` |
| `radius.full` | Avatares, búsqueda circular, iconos circulares y badges tipo píldora | `9999px` |
| `shadow.card` | Propuesta de sombra discreta para tarjetas | `0 2px 6px rgba(23, 58, 94, 0.04)` |
| `shadow.raised` | Propuesta para popovers, menús y autocompletado | `0 4px 12px rgba(23, 58, 94, 0.08)` |
| `shadow.modal` | Propuesta para modales y paneles flotantes | `0 16px 40px rgba(23, 58, 94, 0.18)` |
| `shadow.sticky` | Propuesta para barras fijas de vistas pendientes | `0 -2px 8px rgba(23, 58, 94, 0.06)` |

**Reglas de forma y profundidad**

- La interfaz se organiza en **capas claras**: fondo azul grisáceo muy pálido, navegación blanca y tarjetas blancas; la promoción del catálogo aporta el contraste marino.
- Las tarjetas usan radios estimados de 16–18 px y padding de 24 px; botones y filas de acción, radios de 10–12 px. Avatares e indicadores de estado usan círculos o píldoras.
- Bordes de aproximadamente 1 px y sombras discretas delimitan las superficies. La ilustración del catálogo conserva su tratamiento geométrico plano, sin sustituirse por fotografía.

**Layout y breakpoints**

| Token | Uso | Valor |
| --- | --- | --- |
| `layout.sidebar` | Ancho estimado de la sidebar blanca de altura completa | `270px` |
| `layout.header` | Altura estimada de la cabecera blanca con divisor inferior | `76px` |
| `layout.contentMax` | Propuesta de ancho máximo; el lienzo de referencia mide 1512 × 982 px | `1440px` |
| `layout.marginMobile` | Propuesta de margen lateral en móvil; sin frame disponible | `16px` |
| `layout.marginTablet` | Propuesta de margen lateral en tablet; sin frame disponible | `24px` |
| `layout.marginDesktop` | Padding exterior estimado del espacio de trabajo | `30px` |
| `layout.gutter` | Propuesta para separaciones compactas entre bloques | `16px` |
| `layout.gutterDesktop` | Separación estimada de 18–22 px entre bloques del dashboard | `20px` |

| Breakpoint | Rango | Comportamiento general |
| --- | --- | --- |
| `mobile` | `< 768px` propuesto | Una columna y navegación desplegable por validar. Conservar el orden saludo → métricas → catálogo → acciones rápidas → actividad; adaptar tablas a tarjetas o listas. |
| `tablet` | `768px – 1199px` propuesto | Sidebar colapsable y métricas en dos columnas como propuesta; catálogo y acciones se apilan cuando falte espacio. Sin frame inspeccionado. |
| `desktop` | `≥ 1200px` propuesto | Referencia disponible a 1512 × 982 px: sidebar visible, cuatro métricas iguales en fila, catálogo de aproximadamente dos tercios y acciones de un tercio; actividad a ancho completo. |
| `wide` | `≥ 1440px` propuesto | Mantener proporciones del dashboard y limitar el estiramiento como propuesta; el frame disponible no demuestra comportamiento en otros anchos. |

## 4. Accesibilidad

El objetivo es cumplir **WCAG 2.2 AA** en las vistas principales. Es un requisito de implementación; la especificación visual proporcionada no acredita accesibilidad, estados de interacción ni comportamiento de lectores de pantalla.

**Contraste y color**

- Texto normal sobre fondo: contraste mínimo **4.5:1**. Texto grande y elementos gráficos: **3:1**.
- Ningún estado se comunica solo por color: los badges llevan texto; las desviaciones llevan signo, cifra e icono; los errores llevan mensaje.
- El foco visible usa un anillo de **2 px** en `color.focus` con separación de 2 px respecto del control.
- Revisar especialmente etiquetas pequeñas azul grisáceo, enlaces azules y badges verdes y ámbar: los valores estimados incluyen pares que no alcanzan el contraste de texto normal.

**Teclado y navegación**

- Toda acción debe ser accesible por teclado, incluida la navegación lateral, los menús, las celdas editables y los modales.
- Orden de tabulación lógico: navegación → cabecera → saludo y acción principal → métricas → catálogo → acciones rápidas → actividad. En otras vistas, continuar hacia sus acciones y resumen económico.
- En la tabla de cotización: `Tab` avanza de celda, `Enter` confirma la edición, `Escape` cancela y devuelve el foco a la celda original.
- Los atajos de teclado deben ser adicionales, nunca la única forma de ejecutar una acción.
- Se incluye un enlace “Saltar al contenido” visible al recibir foco; la cuenta truncada conserva un nombre accesible completo y los controles de búsqueda y salida tienen etiqueta.

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

Cálculo de luminancia relativa según **WCAG 2.2**, realizado el `2026-11-09` sobre los valores hexadecimales provisionales de este documento, con ratios redondeados a dos decimales. Evalúa estos pares de color, no los valores reales de Figma; no reemplaza la revisión de tamaños, pesos, opacidad y estados en navegador.

| Combinación | Ratio | Criterio aplicable | Resultado |
| --- | --- | --- | --- |
| Texto principal sobre fondo | 10,83:1 | Texto normal 4,5:1 | AAA |
| Texto secundario sobre superficie | 3,87:1 | Texto normal 4,5:1 | No cumple AA para texto normal; pendiente de ajuste |
| Texto terciario sobre fondo | 3,59:1 | Texto normal 4,5:1 | No cumple AA para texto normal; pendiente de ajuste |
| Blanco sobre `primary` | 10,76:1 | Texto normal 4,5:1 | AAA |
| Blanco sobre `success` | 3,71:1 | Texto normal 4,5:1 | No cumple AA para texto normal; pendiente de ajuste |
| Blanco sobre `danger` | 6,47:1 | Texto normal 4,5:1 | AA |
| `success` sobre `successContainer` | 3,45:1 | Texto normal 4,5:1 | No cumple AA para badges pequeños; pendiente de ajuste |
| `warning` sobre `warningContainer` | 3,00:1 | Texto normal 4,5:1 | No cumple AA para badges pequeños; pendiente de ajuste |
| `danger` sobre `dangerContainer` | 5,91:1 | Texto normal 4,5:1 | AA |
| `info` sobre `infoContainer` | 3,96:1 | Texto normal 4,5:1 | No cumple AA para texto normal; pendiente de ajuste |
| `onPrimarySoft` sobre `primarySoft` | 9,59:1 | Texto normal 4,5:1 | AAA |
| `textInverseMuted` sobre `surfaceInverse` | 8,58:1 | Texto normal 4,5:1 | AAA |
| `focus` sobre superficie | 4,44:1 | Componentes gráficos 3:1 | Cumple para este par |
| `borderInput` sobre superficie | 3,87:1 | Componentes gráficos 3:1 | Cumple para este par |

**Hallazgos y correcciones**

- Los tonos secundarios `#6B84A0`, verde `#23966E`, ámbar `#CF7B16` y azul de enlace `#287ACA` requieren revisar su uso en texto pequeño. Se conservan como estimaciones de la referencia, con ajustes accesibles pendientes de aprobación.
- `color.borderInput` (`#6B84A0`) es una propuesta para controles, distinta del borde decorativo pálido `#DCE7F0`. Los bordes claros **no deben ser el único indicio de un control interactivo**.
- Los ratios anteriores corresponden a la nueva paleta provisional. Hover, presionado, deshabilitado, foco sobre otros fondos, overlays y los colores exactos del archivo requieren validación antes de declarar conformidad.

## 5. Arquitectura de componentes con Atomic Design

La implementación será en **React + JavaScript**. Los componentes de UI puros viven separados de los que consumen datos o estado global. El inventario es conceptual: los patrones repetidos del dashboard sugieren límites reutilizables, pero no se verificaron componentes nativos de Figma ni su implementación; las piezas de otras pantallas siguen siendo propuestas.

**Base técnica propuesta**

| Tema | Propuesta | Motivo | Estado |
| --- | --- | --- | --- |
| Base headless | Librería accesible compatible con React, por seleccionar | Resolver foco y teclado en diálogos, menús, tabs, popovers y tooltips; el render no identifica una librería. | Pendiente de selección y aprobación |
| Iconos | SVG de contorno mediante una librería coherente; Iconify para React como opción | Mantener glifos pequeños, trazo uniforme y tonos azul grisáceo o azul de acento; familia exacta por confirmar. | Propuesta, pendiente de aprobación |
| Estilos | Tailwind CSS o CSS con custom properties | Conservar tokens reutilizables y el sistema claro de superficies; la tecnología del mockup no fue verificada. | Por definir |
| Formato de datos | Formateador compartido para CLP, porcentajes y fechas en español | Conservar puntos de miles, como `$355.000` y `$231.500`, y etiquetas del ejemplo como `65% del presupuesto`. | Propuesta de implementación basada en el contenido visible |

> La especificación disponible describe el render de Figma. Framework, librerías, fuente exacta, rutas de navegación y comportamiento en ejecución no fueron inspeccionados; las elecciones técnicas de esta tabla se confirmarán para el frontend React del proyecto.

### Átomos

Componentes base indivisibles.

| Componente | Responsabilidad | Estados mínimos | Recomendación |
| --- | --- | --- | --- |
| `Button` | Ejecutar acciones | `default`, `hover`, `focus`, `active`, `disabled`, `loading` | Propio; principal marino, secundario blanco o con borde, enlace azul y peligro. Interacciones propuestas. |
| `IconButton` | Búsqueda circular, salida de cuenta y acciones compactas | `default`, `hover`, `focus`, `disabled`, `loading` | Nombre accesible obligatorio y tooltip; búsqueda sobre círculo azul pálido. Estados distintos del predeterminado por definir. |
| `Input` | Capturar texto | `default`, `focus`, `error`, `disabled`, `readonly` | Propio; borde `color.borderInput` |
| `MoneyInput` | Capturar importes CLP | `default`, `focus`, `error`, `disabled` | Propio; alineación derecha, cifras tabulares y borde `color.borderInput` |
| `NumericInput` | Capturar cantidades y porcentajes | `default`, `focus`, `error`, `disabled` | Propio; teclado numérico en móvil y borde `color.borderInput` |
| `Select` | Elegir una opción | `default`, `focus`, `error`, `disabled` | Propio o headless; borde `color.borderInput` |
| `Label` | Nombrar campos | `default`, `required`, `disabled` | Propio |
| `Icon` | Mostrar pictogramas de contorno | `size`, `color` | Trazo consistente; composición con contenedor azul pálido, cuadrado redondeado o circular según el patrón. |
| `Badge` | Comunicar estado breve | `success`, `warning`, `danger`, `neutral`, `info` | Píldora compacta con texto: “Prueba activa” y “Finalizado” verdes; “En ejecución” ámbar. Otros estados propuestos. |
| `Spinner` | Carga breve | `sm`, `md`, `lg` | Propio |
| `Kbd` | Mostrar atajos | `default` | Propio |
| `Divider` | Separar bloques | `horizontal`, `vertical` | Propio |

### Moléculas

Combinaciones pequeñas de átomos.

| Componente | Responsabilidad | Composición sugerida |
| --- | --- | --- |
| `FormField` | Campo con etiqueta, ayuda y error | `Label + Input/NumericInput + HelpText + ErrorText` |
| `SearchBar` | Acceso a búsqueda; campo completo por definir | En la cabecera se describe solo un `IconButton` circular; apertura, campo, filtros y resultados pendientes. |
| `Toast` | Notificación temporal | Cola + icono + mensaje + acción opcional |
| `Tabs` | Navegación entre secciones | Headless accesible + estilo propio |
| `DropdownMenu` | Acciones contextuales | Headless accesible |
| `CompanySwitcher` | Resumen de empresa; cambio de empresa por confirmar | Avatar azul pálido “CR” + nombre truncado “Cristales Santiago Demo” con etiqueta accesible completa + subtítulo “Vidriería”. |
| `UserMenu` | Acciones de cuenta | Área inferior de sidebar con divisor y control de salida; menú, rol y otras acciones no verificados. |
| `QuotationStatus` / `StatusBadge` | Mostrar estados comerciales, de trabajo y de prueba | `Badge` con texto; variantes visibles de prueba activa, trabajo en ejecución y finalizado, y variantes propuestas para cotizaciones. |
| `MoneyValue` | Mostrar un importe | Cifra tabular + unidad/moneda + énfasis |
| `DeltaValue` | Mostrar una desviación | Signo + monto + porcentaje + color/icono |
| `ChapterHeading` / `SectionHeading` | Encabezar capítulos o secciones | Título + texto secundario + enlace opcional, como “Actividad reciente” y “Ver historial”; código y subtotal solo en el patrón propuesto de cotización. |
| `EmptyState` | Explicar una ausencia de datos | Icono + título + descripción + acción |
| `ConfirmDialog` | Confirmar acciones destructivas | Título + consecuencia + acciones |

### Organismos

Bloques funcionales de alto nivel. Marca, navegación, métricas, promoción de rubro, filas de acción y actividad son composiciones inferidas de la referencia; sus nombres no implican que existan como componentes en Figma o en el código.

| Componente | Responsabilidad | Decisión |
| --- | --- | --- |
| `AppSidebar` | Marca, navegación y cuenta de empresa | Blanca, aproximadamente 270 px y altura completa. Marca: cuadrado marino redondeado con “i” minúscula + “ÍndicePartidas” + “COSTOS CON EVIDENCIA”. Orden: Inicio → Cotizaciones → Trabajos / Órdenes → Catálogo de partidas → Componentes y materiales → Clientes → Historial → ÍndicePartidas → Directorio de prestadores → Configuración. Filas con icono de contorno y texto; Inicio seleccionado con fondo azul pálido, texto azul oscuro y punto azul a la derecha. Cuenta separada por divisor al pie. |
| `AppHeader` | Rubro, sección, prueba y acceso a búsqueda | Blanca, aproximadamente 76 px y divisor inferior. Izquierda: “VIDRIERÍA” pequeño en mayúsculas sobre “Inicio”. Derecha: píldora verde pálido “Prueba activa” y búsqueda circular azul pálido. Posición sticky no verificada. |
| `QuotationHeaderStrip` | Ficha compacta de la cotización | Propuesta: cliente, obra, plazo, moneda y versión sobre superficies claras; pendiente del frame de cotización. |
| `QuotationToolbar` | Acciones de edición de la hoja | Propuesta: insertar capítulo, agregar APU y configurar columnas; importación y comportamiento pendientes de confirmación. |
| `QuotationTable` | Matriz de partidas por capítulo | Propuesta: edición en celda y agrupación por capítulos, con encabezados marinos y bordes suaves. |
| `EconomicSummaryBar` | Resumen de costo, utilidad, neto, IVA y total | Propuesta para cotización: datos por rol, superficie clara y cifras marinas; fijación y expansión móvil por validar. |
| `ClientPreviewModal` | Documento comercial sin costos internos | Propuesta: imprimible y accesible, con identidad azul y sin márgenes ni utilidad. |
| `ExecutionForm` | Registro rápido de ejecución en terreno | Propuesta: móvil primero, campos mínimos y captura habitual en ≤ 30 segundos. |
| `ComparisonTable` | Comparación Cotizado vs. Ejecutado | Propuesta: cifras y desviaciones por partida, separando costos de precios de venta. |
| `Modal` | Confirmaciones y formularios flotantes | Headless accesible |
| `Card` | Contenedor composable de métricas, promoción y acciones rápidas | Blanco, borde pálido, radio estimado 16–18 px, padding de 24 px y sombra discreta. Componer tarjeta de métrica, promoción de rubro y filas de acción con icono y chevron; la ilustración de ventana conserva su geometría plana. |
| `FilterPanel` | Agrupar filtros de búsqueda | Propio |
| `DataTable` / `ActivityList` | Listados generales y actividad reciente | Actividad en filas con icono circular, título y metadatos a la izquierda, badge a la derecha y divisores finos; tabla avanzada como propuesta para otros listados. |

### Templates / layouts

| Layout | Uso | Reglas |
| --- | --- | --- |
| `AppLayout` | Marco autenticado y composición de Inicio | Sidebar y header blancos + área azul grisácea con padding estimado de 30 px y gaps de 18–22 px. Inicio: saludo → cuatro métricas iguales → catálogo ancho y acciones estrechas → actividad a ancho completo. Colapsar columnas y navegación conservando ese orden es una propuesta para otros tamaños. |
| `QuotationEditorLayout` | Crear y editar cotización | Propuesta: ficha compacta + toolbar + tabla por capítulos + resumen económico; adaptar al sistema visual claro y validar con su frame. |
| `ExecutionCaptureLayout` | Registrar ejecución en terreno | Propuesta: contexto de OT + formulario breve + acción principal fija; prioridad móvil por RNF. |
| `ComparisonLayout` | Análisis Cotizado vs. Ejecutado | Propuesta: resumen de indicadores + tabla comparativa + detalle progresivo. |
| `ClientDocumentLayout` | Vista previa y exportación comercial | Propuesta: documento blanco, membrete, partidas y totales; sin controles ni datos internos. |
| `AuthLayout` | Login y recuperación | Propuesta: formulario centrado o dos columnas, sin navegación global; sin referencia visual inspeccionada. |
| `CrudLayout` | Mantenedores | Propuesta: título + acciones + filtros + tabla sobre tarjetas claras. |
| `MasterDetailLayout` | Lista + detalle | Propuesta: paneles con scroll independiente; comportamiento pendiente del diseño correspondiente. |

## 6. Reglas de composición

- **No crear componentes específicos sin comportamiento propio.** Si una tarjeta solo cambia textos o datos, se compone con `Card` y sus subcomponentes.
- **Evitar props booleanas acumulativas** como `isBig`, `isRed`, `hasBorder`, `compact` o `adminMode`. Preferir `variant`, `size`, `tone` y composición.
- **Separar UI pura de UI conectada a datos.** Los componentes de presentación reciben props y emiten eventos; los contenedores resuelven estado, permisos y llamadas a la API.
- **Los componentes que manejan foco, teclado, portales o colisiones de pantalla usan una base headless accesible** y reciben solo estilos propios.
- **El dinero se formatea en un único lugar.** Los componentes no concatenan símbolos ni separadores manualmente: reciben valores ya formateados o usan un formateador compartido (CLP, sin decimales por defecto).
- **Los números comparables usan cifras tabulares.** Las tablas alinean importes a la derecha, descripciones a la izquierda y unidades al centro; las métricas de Inicio conservan su alineación visual izquierda.
- **La información interna se encapsula.** Costos, márgenes y utilidad se renderizan solo en componentes internos; la vista cliente usa componentes distintos o un modo explícito que no recibe esos datos.
- **La edición en tabla no debe sorprender.** Toda celda editable se distingue visualmente y confirma o cancela con `Enter` y `Escape`.
- **En móvil, la tarea principal no depende de scroll horizontal.** Las tablas se transforman en tarjetas, listas o formularios apilados.
- **La acción principal tiene una jerarquía clara.** Inicio destaca “Nueva cotización” en marino junto al saludo y repite el acceso como fila secundaria en “Acciones rápidas”; catálogo e historial mantienen menor énfasis.
- **Los formularios no pierden datos.** Antes de abandonar con cambios pendientes se muestra confirmación; el guardado no depende del navegador.
- **Los componentes de presentación no contienen reglas de negocio.** Márgenes, IVA y totales oficiales se validan en backend; las ayudas de presentación se encapsulan en funciones o hooks testeables de React.

## 7. Patrones de pantalla

### Listado CRUD

Debe incluir, como patrón propuesto para las pantallas de listado aún no documentadas en Figma:

- Título y descripción breve.
- Acción primaria visible.
- Filtros principales y búsqueda.
- Tabla o lista con estados de carga, vacío y error.
- Paginación o scroll controlado.
- Confirmación para acciones destructivas.
- En móvil: tarjetas o filas apiladas; filtros en panel desplegable.

### Formulario

Debe incluir, como patrón propuesto para las pantallas de captura aún no documentadas en Figma:

- Validación visible por campo, con mensaje orientado a la corrección.
- Mensaje general de error si falla el envío.
- Acción primaria y secundaria.
- Prevención de doble envío y estado `guardando`.
- Confirmación si se perderán cambios.
- Etiquetas visibles; ayudas solo cuando aportan.

### Detalle

Debe incluir estado, metadatos, acciones por rol, trazabilidad y protección de información confidencial. La referencia visual disponible para el resumen del negocio es **Inicio**, con la cuenta ficticia **Cristales Santiago Demo** y rubro **Vidriería**; el dashboard de escritorio se compone así:

- Introducción: etiqueta azul pequeña en mayúsculas “RESUMEN DEL NEGOCIO”, saludo grande y en negrita “Hola, Cristales Santiago Demo.” y texto secundario “Aquí tienes una vista actualizada de tu operación.” A la derecha, botón principal marino con icono más y texto “Nueva cotización”.
- Métricas: cuatro tarjetas blancas de igual ancho en una fila, con etiqueta atenuada, valor destacado, contexto e icono sobre cuadrado azul pálido redondeado en la esquina superior derecha. Contenido de demostración: “Cotizaciones” → `3` → “2 aceptadas”; “Trabajos activos” → `1` → “Requieren seguimiento”; “Costo presupuestado” → `$355.000` → “En trabajos activos”; “Costos registrados” → `$231.500` → “65% del presupuesto”. Son ejemplos del render, no resultados calculados ni datos reales.
- Catálogo del rubro: tarjeta izquierda de aproximadamente dos tercios del ancho de la fila intermedia. Etiqueta “TU RUBRO”, título “Vidriería” y píldora superior derecha “3 partidas”. Panel interior con izquierda marina y derecha azul secundario: “Catálogo adaptado a tu actividad”, “Configura ventanas a medida” y “Ajusta dimensiones, vidrio, perfiles, accesorios y mano de obra en un solo flujo.” Botón blanco redondeado “Ir a mi catálogo” con chevron derecho. Ilustración frontal plana de ventana de dos paneles, vidrio azul, marco blanco y manillas triangulares simples.
- Acciones rápidas: tarjeta derecha de aproximadamente un tercio del ancho, titulada “Acciones rápidas”. Cuatro filas con borde, alturas y radios consistentes, separación vertical uniforme, icono sobre fondo azul pálido a la izquierda y chevron a la derecha: “Nueva cotización”, “Nuevo cliente”, “Registrar ejecución” y “Consultar historial”. Destinos y comportamiento pendientes de verificar.
- Actividad: tarjeta blanca a ancho completo bajo la fila intermedia, con “Actividad reciente”, subtítulo “Operaciones guardadas en tu cuenta.” y enlace “Ver historial” a la derecha. Filas con icono circular pálido, título y metadatos a la izquierda y estado a la derecha: “Mampara de vidrio templado” — “OT-0018 · Comercial Los Alerces (ficticio)” — badge ámbar “En ejecución”; “Espejos para recepción” — “OT-0017 · Comercial Los Alerces (ficticio)” — badge verde “Finalizado”. Divisores finos entre filas; la lista continúa fuera del borde inferior visible de 982 px y el recorte no representa su final.

**Editor de cotización**

Es un patrón funcional propuesto, pendiente de contrastar con su propio frame de Figma. La referencia visual vigente es el dashboard Inicio descrito arriba; el editor hereda superficies claras, encabezados marinos, acentos azules, bordes suaves y jerarquía financiera. Las adaptaciones de escritorio, tablet y móvil siguientes son propuestas, no vistas verificadas del mockup.

**Escritorio**

- Ficha compacta blanca con código, estado, cliente, obra, plazo, moneda, versión y fecha de vigencia.
- Toolbar clara con acciones de edición de la hoja e iconos de contorno.
- Tabla sobre tarjeta blanca, con bordes pálidos y organización propuesta por capítulos desplegables.
- Columnas principales: partida, unidad, cantidad, precio unitario, subtotal de venta y acciones.
- Columnas internas —costo y margen— visualmente diferenciadas y protegidas.
- Resumen económico con costo, utilidad, neto, IVA y total cliente; superficie clara, cifras marinas y fijación por validar.
- Acción principal marina “Guardar cotización” y acción secundaria “Vista cliente”.

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

**Prioridad:** móvil por el requisito de registro habitual en ≤ 30 segundos. El dashboard solo muestra el acceso “Registrar ejecución”; el formulario y sus interacciones siguientes son propuestas pendientes del frame correspondiente.

- Contexto visible: proyecto, Orden de Trabajo y partida seleccionada.
- Fecha con valor predeterminado.
- Acciones rápidas: “Material”, “Mano de obra” y “Otro costo”.
- Solo los campos esenciales; el resto opcional.
- Cantidad ejecutada como acción diferenciada del registro de materiales u horas para evitar doble conteo.
- Nota y fotografía opcionales, con acceso directo a cámara.
- Acción principal fija “Guardar registro”.
- Confirmación tras guardado exitoso; deshacer solo cuando exista una operación de corrección compatible con la trazabilidad.
- Estado sin conexión: conservar lo ingresado y ofrecer reintento. Solo informar guardado local o sincronización automática si esas capacidades se implementan y verifican.

**Comparación Cotizado vs. Ejecutado**

**Escritorio**

- Propuesta pendiente de frame: resumen superior en tarjetas claras con costo estimado, costo real, precio de venta, utilidad y margen según permisos.
- Comparación propuesta por partida con desviación en monto y porcentaje, etiquetas marinas e importes legibles.
- Identificación propuesta de sobrecostos con signo, texto y color, sin depender solo del tono.
- Gráfico simple propuesto de costo estimado vs. real; no está documentado por el dashboard.
- Filtros por capítulo y detalle de partida como propuesta funcional, con iconos de contorno y bordes suaves.

**Móvil**

- Tarjetas de indicadores principales.
- Lista de partidas con desviación, ordenable por impacto.
- El gráfico se simplifica o se reemplaza por una comparación textual.
- El detalle de una partida se abre en panel completo.

**Vista previa comercial del cliente**

- Documento blanco propuesto con membrete de la empresa y marca marina; frame e interacción de impresión pendientes de verificar.
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
| Búsqueda / notificaciones | “Buscar” para el control circular de Inicio; “Notificaciones, {n} sin leer” para el control propuesto de otras vistas. |
| Ayuda | “Abrir ayuda” |
| Cerrar sesión | “Cerrar sesión” |
| Cerrar vista previa | “Cerrar vista previa de la cotización” |
| Imprimir | “Imprimir cotización {código}” |

**Reglas de nombres accesibles**

- El `aria-label` incluye el contexto: no basta con “Eliminar” o “Ver”.
- El texto visible del tooltip puede ser más breve, pero debe coincidir en intención con el nombre accesible.
- Los controles que solo muestran un icono siempre tienen nombre accesible y tooltip; el nombre completo de la empresa sigue disponible aunque se trunque visualmente.
- Los estados de los controles se comunican con `aria-expanded`, `aria-pressed` o `aria-disabled` según corresponda, no solo con un cambio de color.

**Comportamiento responsivo por vista**

| Vista | Escritorio | Tablet | Móvil |
| --- | --- | --- | --- |
| Cotización | Propuesta: tabla completa y resumen económico | Propuesta: tabla con scroll y columnas internas ocultas | Propuesta: tarjetas por capítulo y edición en panel |
| Ejecución | Propuesta: formulario con contexto lateral | Propuesta: formulario de una columna | Propuesta: formulario breve con acción fija |
| Comparación | Propuesta: indicadores + tabla + gráfico | Propuesta: indicadores + tabla resumida | Propuesta: indicadores + tarjetas de partidas |
| Vista cliente | Propuesta: modal amplio | Propuesta: modal adaptado | Propuesta: documento vertical |

## 8. Estados obligatorios

Cada pantalla o componente conectado a datos debe especificar los estados siguientes como requisitos de implementación; no fueron inspeccionados en Figma. En Inicio se propone carga independiente de métricas y actividad, vacío con acción pertinente y error con reintento, sin sustituir datos desconocidos por ceros ni por los ejemplos de demostración. Los badges descritos en el render son “Prueba activa” en verde pálido, “En ejecución” en ámbar y “Finalizado” en verde; no acreditan transiciones de estado implementadas.

| Estado | Significado | Tratamiento obligatorio |
| --- | --- | --- |
| `loading` | Los datos aún no están disponibles | Skeletons o spinner; en Inicio conservar el espacio de las tarjetas y evitar cifras desconocidas como cero. |
| `empty` | No existen datos | Explicación breve y acción para crear el primer registro; en Inicio permitir crear una cotización o acceder al catálogo según el bloque. |
| `error` | Falló la carga o la operación | Mensaje claro, causa cuando se conozca y acción de reintento |
| `success` | La operación se completó | Confirmación breve; no bloquear la pantalla |
| `forbidden` | El rol no tiene acceso | Explicar la restricción sin revelar información sensible |
| `offline` | No hay conexión | Conservar entradas y ofrecer reintento; solo declarar persistencia local si está implementada y confirmada. |
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
| Pendiente de sincronizar | Guardado local, si existe persistencia verificada | Mostrar icono y acción “Reintentar”; sin persistencia local, indicar “Sin enviar” y conservar el formulario. |
| Sincronizado | Guardado en servidor | Confirmación discreta con hora |
| Error de sincronización | No se pudo guardar | Mantener los datos ingresados y explicar qué hacer; la recuperación tras recargar depende de persistencia local verificada. |
| Validación | Datos incompletos o fuera de rango | Señalar el campo y no perder lo ingresado |

**Estados por pantalla**

**Editor de cotización**

| Estado | Texto y comportamiento |
| --- | --- |
| Cargando | Skeletons en ficha y tabla; totales ocultos. “Cargando cotización…” |
| Vacío | “Esta cotización aún no tiene partidas.” Acción: “Agregar partida”; “Importar Excel” solo si se confirma e implementa esa capacidad. |
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
| Sin conexión | “Sin conexión. El registro no se ha enviado.” Acción: “Reintentar”; indicar guardado en el dispositivo y envío automático únicamente si ambas capacidades están verificadas. |
| Sin permiso de costos | El rol Maestro solo registra fotografías y notas; los campos de costo se ocultan. |

**Comparación Cotizado vs. Ejecutado**

| Estado | Texto y comportamiento |
| --- | --- |
| Cargando | Skeletons en indicadores y tabla; no mostrar cifras en $0. |
| Vacío | “La comparación estará disponible cuando existan costos registrados.” Acción: “Ir a ejecución”. |
| Error | “No pudimos calcular la comparación.” Acción: “Reintentar”. |
| Sin permiso | “No tienes permiso para ver costos internos.” Sin datos parciales que revelen márgenes. |
| Sin conexión | Si existen datos previamente cargados, mostrarlos con “Actualizado al {fecha}” y banner de conexión; en otro caso, indicar que no están disponibles. |

**Vista previa comercial**

| Estado | Texto y comportamiento |
| --- | --- |
| Cargando | Spinner dentro del modal. “Generando vista previa…” |
| Vacío | “La cotización no tiene partidas para mostrar.” Acción: “Cerrar”. |
| Error | “No pudimos generar la vista previa.” Acción: “Reintentar”. |
| Sin permiso | La acción “Vista cliente” se oculta o deshabilita según el rol. |
| Sin conexión | Si los datos ya están cargados, permitir imprimir; si no, mostrar error y opción de reintentar. |

**Reglas generales de estado**

- Una cifra desconocida no se muestra como `$0`; se muestra como “Sin datos” o se oculta con explicación. Los valores, nombres y códigos de Inicio son contenido ficticio de referencia y no rellenan estados vacíos reales.
- Los errores de red no eliminan el formulario ni el trabajo en curso.
- Las confirmaciones no tapan la acción principal ni exigen un clic adicional innecesario.
- Los estados críticos se anuncian a tecnologías de asistencia mediante regiones `aria-live`.

## 9. Checklist de diseño antes de implementar

- [] Los tokens provisionales de color, tipografía, espaciado y radios se confirmaron con Figma y se ajustaron los pares de contraste pendientes.
- [] El layout de Inicio se contrastó con el frame de referencia y las propuestas de otras pantallas se validaron con sus diseños.
- [] Cada flujo crítico tiene estado vacío, carga y error definidos y verificados; no se confunden con estados observados en el mockup.
- [] Las acciones destructivas tienen confirmación.
- [] Los formularios tienen validación y mensajes claros.
- [] Los componentes complejos usan una base accesible compatible con React cuando corresponde.
- [] La navegación por teclado, los nombres accesibles y los estados de foco, hover, presionado y deshabilitado están verificados.
- [] El diseño móvil y tablet se validó; las adaptaciones propuestas conservan el orden saludo, métricas, catálogo, acciones rápidas y actividad del escritorio.
