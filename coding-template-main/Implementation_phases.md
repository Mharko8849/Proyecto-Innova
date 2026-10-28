# Implementation Phases — Plan de entrega de ÍndicePartidas

> **Versión:** 0.4 — 28-10-2026. **Estado:** plan propuesto, revisado contra el árbol local del commit `a8dea5d`, incluidas la arquitectura y la base de datos incorporadas en `865c55d` y `50aa50d`. Ninguna fase de software se declara implementada por la existencia de este documento.
>
> **Flujo:** Cotizar → aceptar una base → crear trabajo → registrar ejecución → comparar → reutilizar lo aprendido.

## 1. Objetivo y límites

El objetivo no es terminar todos los mantenedores antes de mostrar valor: es entregar un recorrido completo con pocos datos y ampliarlo después. La referencia técnica actual es **Go + React/JavaScript + PostgreSQL**, coincidente en PRD, RNF y [Architecture.md](Architecture.md). La arquitectura propone monolito modular por capas, Gin, GORM, JWT, Vite, OpenAPI con `swaggo/swag` y Docker Compose; el modelo de datos propone PostgreSQL 15+, precisión decimal y migraciones SQL con `golang-migrate`. Son definiciones documentadas, no dependencias instaladas ni servicios implementados. Confirmar versiones y las alternativas que aún dejan abiertas sus checklists antes de configurar el proyecto.

**Interpretación de entregas:** P1 es un incremento interno del recorrido de problema-solución, no el MVP completo del PRD. Dashboard básico no equivale a dashboard completo; vigencia simple no equivale a actualización automática de precios; versiones básicas no equivalen a gestión avanzada. Este plan distribuye esas capacidades sin quitar RF del compromiso vigente. Confirmar con producto qué incluye el prototipo inicial del índice y el alcance completo de RF22; las reglas y pruebas propuestas aquí no modifican ni aprueban documentos ajenos.

Este plan desarrolla los tres problemas de [problema-solucion.md](../Trabajos-Clases/problema-solucion.md): presupuestar, capturar costo real y comparar. Identidad, catálogo y aceptación son soportes de ese recorrido, no tres productos adicionales.

### Hitos académicos y técnicos

| Hito | Qué se busca | Qué no se debe confundir |
|---|---|---|
| Hasta 03-11 | Cerrar el experimento según su criterio previo y registrar aprendizaje | Esta planificación no modifica el criterio ni inventa resultados |
| PEP 1, 10 y 12-11 | Evidencia, propuesta, negocio, prototipo y defensa | No exige SaaS funcionando |
| Desde 17-11 | Construcción y pruebas del MVP | No obliga a anticipar una implementación antes de validar |
| Primera entrega técnica, fecha por acordar | Un usuario recorre cotizar → registrar → comparar | No equivale a todo el producto del semestre |
| Entrega semestral | Cobertura acordada del catálogo, software desplegado y evidencia de uso | ML e índice requieren acuerdo de alcance y datos, no resultados simulados presentados como reales |

Las semanas siguientes son una secuencia relativa propuesta, no fechas ni asignaciones ya aprobadas.

## 2. Coordinación con el trabajo de los compañeros

| Artefacto | Qué debe aportar su responsable | Dependencia de este plan |
|---|---|---|
| [Architecture.md](Architecture.md) | Monolito modular Go/React, JWT, capas, errores, transacciones, pruebas y despliegue | Diseño técnico disponible; P0 concilia sus pendientes y P1 implementa solo el soporte del recorrido |
| [DB_Diagram.md](DB_Diagram.md) | Diseño inicial con empresas, APU, versiones, OT, ejecución, consentimiento y agregados | Modelo objetivo disponible, no esquema migrado; cada migración se revisa con su responsable |
| [Design.md](Design.md) | Recorrido móvil/escritorio, estados, permisos visibles y captura mínima | Cada pantalla antes de integrarse |
| [PRD.md](PRD.md) | Alcance, reglas de negocio y aceptación | Fuente funcional de cada fase |
| [Rules.md](Rules.md) | Revisión, cambios, evidencias y coordinación | Forma de entregar todas las fases |

### Diferencias documentales a conciliar, sin modificar trabajo ajeno

| Tema | Estado observado en `a8dea5d` | Efecto sobre este plan |
|---|---|---|
| Stack | PRD, RNF y Architecture coinciden en React/JavaScript; Rules §6.1 y ADR-001, y problema-solución, aún mencionan Vue | Mantener React/JavaScript y solicitar actualización de menciones residuales; no reabrir el stack desde este plan |
| Arquitectura, datos e interfaz | Architecture y DB_Diagram ya están desarrollados; Design sigue siendo plantilla | Usar los diseños existentes; confirmar alternativas ORM/framework/migraciones y acordar pantallas mínimas sin sustituir al responsable de diseño |
| Prioridad de entrega | PRD incluye capacidades básicas/prototipo del índice en MVP; problema-solución difiere versiones avanzadas y dashboard completo, pero su mapa también dice versionado/índice después | Distinguir incremento interno, MVP básico y ampliaciones; aprobar la interpretación antes de declarar cobertura, sin asumir que toda diferencia de nivel es una contradicción |
| “Otro” | RF22 completo incluye fórmulas y parámetros, no solo nombre de rubro | Configuración básica en P1 es un subconjunto propuesto; fórmulas libres requieren una ampliación aprobada, no se dan por eliminadas |
| Versiones y entidades | DB separa `quotations`, `quotation_versions` y snapshots; Architecture §9 describe nuevas filas de `cotizaciones` y otros nombres | P1 usa desde el inicio una versión base y sus snapshots del modelo objetivo; P3 completa revisiones/diff. Conciliar los nombres del contrato con sus responsables, sin crear un modelo paralelo |
| Migraciones | DB §11 define `backend/migrations/` y `golang-migrate`; Architecture muestra `internal/migrations/` y deja AutoMigrate/SQL por resolver | Acordar una sola ruta y herramienta antes de crear scripts; no mantener dos fuentes de esquema ni generar las tablas de fases futuras por anticipado |
| Identidad y consentimiento | DB usa `company_id`, `user_roles` y `data_sharing_consents` por empresa; Architecture usa claims `empresa_id`/`rol` y menciona consentimiento en usuarios | Acordar mapeo JWT/API/BD y usar el consentimiento empresarial del modelo como referencia a conciliar; no crear un booleano adicional en usuarios ni mezclar roles sin traducción explícita |
| Costo ejecutado | DB §9.3 multiplica `accumulated_quantity` por `accumulated_real_cost`, mientras los registros tienen `actual_cost` | Definir si el costo es unitario o total antes de implementar. Propuesta del caso 3.1: `actual_cost` es importe incremental y el acumulado es suma de importes, no cantidad multiplicada por un costo ya total |
| Predicción y directorio | Rules §10 propone promedios ponderados y >500 OT, además de directorio básico; PRD no define ese umbral y sitúa red de prestadores en futuro | Revisar con producto/validación; 500 OT no garantiza comparabilidad ni ML útil y el directorio no se incorpora automáticamente al primer incremento |
| Privacidad | Rules añade aislamiento por empresa y exclusión de identificadores/márgenes en exportación | Incorporar esas verificaciones desde P1; para índice, añadir protección de cohortes pequeñas, porque quitar campos no prueba anonimización |

La jerarquía documental y las convenciones de integración están en Rules; una contradicción no se resuelve eligiendo silenciosamente la opción más conveniente. Estas observaciones no acreditan aprobación de ADRs ni existencia de ramas/PR, esquema, autenticación o código.

Los responsables nominales siguen **por asignar por el equipo**. No se asumen a partir de commits ni se sobrescriben sus documentos. Si falta una definición, abrir una decisión pendiente en la tarea; se puede trabajar con un contrato propuesto o datos sintéticos explícitos, pero no integrarlo como definitivo sin revisión conjunta.

### Acuerdos mínimos antes de la primera tarea de código

1. Conciliar el primer recorrido de este plan con el alcance amplio de la sección 9 del PRD y aprobar qué corresponde a primera entrega, ampliación semestral y posterior.
2. Asignar responsable y revisor de arquitectura, datos, interfaz y flujo funcional.
3. Acordar nombres de conceptos, unidades, estados y representación de dinero.
4. Acordar contrato de autenticación, empresa activa y permisos en backend.
5. Acordar contratos HTTP: campos, errores, paginación, revisión concurrente y reintentos.
6. Seleccionar las partidas y participantes del primer piloto, sin inventar un rubro ganador.
7. Confirmar los ADRs y el contrato JWT descritos por Rules con arquitectura, incluidas pertenencia a empresa y permisos efectivos; no tratar claims sin validar como autorización suficiente.
8. Conciliar el alcance MVP del PRD con este plan: RF08, RF16–19 y RF22 completo necesitan entrega identificada o cambio de prioridad aprobado. Mientras no exista ese acuerdo, conservarlos como pendientes del compromiso vigente.
9. Revisar con datos los contratos de la sección 2.1 antes de escribir migraciones/cálculos. Usar la versión inicial del modelo, no reproducir automáticamente todas sus tablas o asumir que sus checklists prueban integridad implementada.

### 2.1. Contratos del modelo objetivo que deben cerrarse por incremento

No se edita aquí el DBML ni se imponen migraciones al compañero. Estas son condiciones de aceptación para convertir su diseño inicial en software:

| Contrato | Referencia actual | Acuerdo y prueba necesarios |
|---|---|---|
| Dinero y cantidades | DB §2: `NUMERIC(14,2)` y `NUMERIC(12,4)` | Mantener decimales exactos entre JSON, Go y PostgreSQL; definir moneda, redondeo y límites. En `performance_quantity`, fijar orientación del rendimiento (recurso/unidad o inversa) |
| Costo y acumulados | DB `execution_logs` / `work_order_items` y §9.3 | Fijar semántica de `actual_cost`; evitar doble multiplicación y doble suma. Si se persisten acumulados, actualizarlos atómicamente y comprobar que coinciden con registros vigentes |
| Datos completos y correcciones | DB muestra cantidades/costos con default cero y logs sin campos de corrección/completitud | Acordar cómo distinguir cero confirmado de pendiente y cómo auditar correcciones sin perder historia. No cerrar una OT solo porque todos los campos tienen cero por defecto |
| OT única y reintentos | `work_orders` tiene código único por empresa, pero no unicidad declarada de `quotation_version_id`; logs sin clave de operación | Acordar restricciones/idempotencia que soporten peticiones simultáneas y reintentos. Comprobar existencia antes de insertar, por sí solo, no garantiza unicidad concurrente |
| Propiedad empresarial | DB distribuye `company_id` entre cabeceras; hijos heredan contexto por FK | Verificar cadena de propiedad y referencias entre empresas, también al escribir relaciones. Acordar restricciones compuestas o estrategia equivalente con datos; no consultar hijos sin comprobar su empresa |
| Snapshot e historial | Versiones y recursos snapshot, FK con borrados en cascada | Acordar protección de versiones aceptadas y borrado lógico; probar que borrar/archivar una cabecera no destruye históricos. Los snapshots no impiden por sí mismos UPDATE o DELETE |
| Resultado total e indirectos | Versión con porcentajes/precio; ejecución con costo total por registro | Definir cómo registrar gastos generales/indirectos reales y evitar doble conteo antes de mostrar resultado neto; desglose detallado requiere contrato compatible con el modelo |
| Índice y consentimiento | `data_sharing_consents`, `benchmark_index_records`, claves por rubro/nombre/unidad | Acordar selección de origen, cohorte, período, calidad, retiro y supresión de muestras pequeñas; ausencia de `company_id` en la salida no acredita anonimización |

No todos estos contratos bloquean P1: dinero, versión base, propiedad y reintento son previos al recorrido; completitud/corrección se consolidan en P2; índice/consentimiento bloquean únicamente su línea de publicación en P5. Un flujo no resuelto no se habilita al usuario como si estuviera terminado.

## 3. Backlog por fases

### P0 — Validación y contratos previos

- **Ventana:** antes de construir; compatible con preparación de PEP 1.
- **Responsables:** producto/validación y referentes de arquitectura, datos e interfaz; nombres pendientes.
- **Entrada:** entrevistas, experimento y documentos actuales.
- **Trabajo:** cerrar aprendizajes reales; probar prototipo; priorizar partidas del piloto; revisar la arquitectura y el DBML existentes; acordar sus contratos pendientes, ruta de migraciones y mapeo JWT/empresa; registrar dudas y decisiones. No rehacer los documentos de los compañeros.
- **Salida verificable:** alcance revisado, recorrido dibujado, ejemplos numéricos y tareas con responsable/revisor. Resultados experimentales enlazados cuando existan.
- **No incluye:** implementar una aplicación para sustituir evidencia de validación.
- **Bloqueos:** no convertir intenciones de uso en adopción probada; no cambiar retrospectivamente el umbral del experimento.

### P1 — Primer recorrido completo, mínimo y privado

- **Ventana propuesta:** primera semana de construcción.
- **RF:** RF01–03, RF05, RF05.1, RF09–12, RF14, RF20–22 en su alcance mínimo.
- **Entrada:** P0 y contratos iniciales revisados.
- **Recorrido demostrable:** entrar como dueño de una empresa piloto, registrar cliente/proyecto, crear una partida y cotización de una línea, aceptar la base, crear trabajo, cargar un costo real y consultar desviación.
- **Implementación vertical:** contrato y modelo/migración acordados → persistencia y caso de uso Go → API documentada → pantalla React/JavaScript → pruebas y revisión según Rules. Cada incremento debe conectar esas piezas, no terminar solo una capa.
- **Modelo objetivo a usar:** `companies`, `users`, `roles`/`user_roles`, `clients`, `projects`, `trades`, `categories`, `resources`, `items`/`item_apu_resources`, `quotations`, una `quotation_version` inicial con `quotation_items`/`quotation_item_resources`, `work_orders`/`work_order_items` y `execution_logs`. Crear solo las migraciones necesarias al recorrido y revisar nombres/tipos con datos; no duplicar estas entidades con nombres del ejemplo anterior de Architecture.
- **Soporte técnico del incremento:** scaffold mínimo del monolito Go y React/Vite, configuración sin secretos, PostgreSQL de prueba, JWT/RBAC y errores JSON con `request_id`. Tomar Gin/GORM y las herramientas de prueba de Architecture como referencia sujeta a sus pendientes, no instalar ni desplegar por la sola aprobación de este documento.
- **Interfaz:** captura simple; estados vacío, carga, error, guardado y acceso denegado. No requiere un sistema visual completo.
- **Reglas:** base aceptada inmutable; propiedad por empresa; costos previstos distintos de precios de venta; dinero exacto; registros parciales explícitos.
- **Permisos del recorrido:** Architecture §8 reserva márgenes/gastos generales a dueño/contratista; los roles de captura no reciben esos campos. Confirmar equivalencia de esos roles con las semillas de DB y verificar recursos encadenados, no solo el `company_id` de la cabecera.
- **Revisiones comerciales:** P1 sin historial completo no permite sobrescribir ofertas enviadas/aceptadas. Si el piloto necesita revisar una oferta, incorporar RF06 mínimo antes de habilitar ese uso, coordinando con ADR-002 de Rules.
- **Pruebas de salida:** caso numérico de la sección 3.1 de este plan; usuario anónimo rechazado; dos empresas aisladas; editar catálogo no cambia la base; una cotización no aceptada no crea trabajo.
- **Fuera:** publicación de índice, ML, archivos, motor libre de fórmulas y administración avanzada de roles.
- **Criterio de paso:** una demostración reproducible recorre el flujo completo con persistencia y pruebas reales. Si no se logra, reducir variedad, no eliminar la comparación final.

### P2 — Captura útil, repetible y corregible

- **Ventana propuesta:** segunda semana; incluye pruebas con usuarios.
- **RF:** RF04, RF11, RF12 y RF14; completar RF02, RF05, RF20–22 básicos.
- **Entrada:** recorrido P1 operativo.
- **Trabajo:** APU con componentes; varias partidas; precios para nuevas cotizaciones; rubros y unidades separados; registros incrementales; correcciones auditadas; estado de completitud por partida; cierre controlado; reutilización privada de partidas. Conciliar campos/estrategia de corrección, datos pendientes e indirectos antes de ampliar `execution_logs`, sin cambiar el modelo por cuenta propia.
- **Dashboard básico RF16:** a partir de datos privados, mostrar tasa de aceptación con período/denominador definido, desviación por trabajos completos y resultado según costos incluidos. Con poca historia mostrar vacío/datos insuficientes; dashboard avanzado e indicadores adicionales pueden esperar. RF14 por trabajo no reemplaza RF16 completo.
- **Captura propuesta:** seleccionar partida, indicar cantidad del período y costos incrementales; horas y su valoración cuando corresponda. Fecha y moneda se heredan cuando sea seguro. Precargar no significa inventar ejecución.
- **Errores:** impedir unidades incompatibles, costos negativos sin flujo de corrección, doble envío y edición concurrente silenciosa. Con red caída, no mostrar “guardado” hasta confirmación; sincronización offline completa no se promete en esta entrega.
- **Pruebas de salida:** dos registros suman una sola vez; el total real no se multiplica otra vez por la cantidad; reintento no duplica; acumulados coinciden con logs vigentes; corrección conserva trazabilidad; otra empresa no accede; incompletos no generan margen final; cierre y edición concurrentes no dejan un resultado inconsistente. Dashboard no incluye márgenes ajenos ni mezcla períodos/unidades incompatibles.
- **Piloto:** cronometrar registros, registrar ayuda necesaria y completitud; conservar por separado resultados sintéticos, retrospectivos y prospectivos. No sustituye ni reinterpreta el experimento previo.
- **Criterio de paso:** cantidades/costos se consolidan sin duplicidad y usuarios pueden intentar el flujo sin instrucciones paso a paso.

### P3 — Ampliación semestral: versiones y flujo comercial

- **Ventana propuesta:** tercera semana, una vez estabilizado el recorrido.
- **RF:** RF06, RF08, RF09 y RF10; RF07 según alcance aprobado. RF08 permanece Must en el PRD: planificar su versión simple aquí salvo cambio de prioridad aprobado.
- **Entrada:** P2 y política de versiones acordada con datos/arquitectura.
- **Trabajo obligatorio propuesto para cubrir catálogo:** versiones inmutables con cambios registrados; salida comercial compartible sin costos/márgenes internos; aceptación de versión específica; conversión sin reescribir; OT única por versión aceptada.
- **Diferencia con P1:** P1 ya usa `quotation_versions` con una base inicial y snapshots; P3 incorpora crear nuevas versiones comerciales, comparación de cambios, vigencia y envío. No reemplaza el modelo ni migra desde una cotización sin versión; una nueva versión nunca reescribe una OT existente.
- **Extraordinarias:** si se habilitan, aprobación y presupuesto separado; si se difieren, explicar que el MVP no administra cambios de alcance y no recalcularlos como sobrecosto de la base.
- **Vigencia:** fecha y aviso simple de vencimiento para cubrir RF08; confirmar regla con producto. Actualización automática de precios es posterior y nunca recalcula la base aceptada.
- **Pruebas de salida:** tres versiones conservan valores; diff identifica cantidad/precio; exportación no filtra margen; rechazo no habilita OT; reintentos y peticiones paralelas no crean dos OT; actualización de catálogo no altera versiones.
- **Criterio de paso:** los siete puntos del catálogo tienen evidencia funcional o un recorte explícitamente acordado con la contraparte docente, y RF08 tiene prueba de vigencia. Esa cobertura no acredita por sí sola todos los Must adicionales del PRD. No eliminar una obligación por decisión unilateral del agente.

### P4 — Endurecimiento, despliegue y aceptación del MVP

- **Ventana propuesta:** cuarta semana; seguridad y pruebas comienzan en P1, no se posponen hasta aquí.
- **RF:** todos los incluidos; **RNF:** seguridad, integridad, usabilidad, rendimiento, respaldo y API.
- **Entrada:** recorrido y ampliaciones elegidas estables.
- **Trabajo:** instalación reproducible; migraciones probadas; API actualizada; configuración sin secretos; HTTPS en entorno desplegado; logs sin datos sensibles; respaldo/restauración; pruebas de aislamiento; medición de rendimiento; accesibilidad del camino crítico.
- **Verificación técnica propuesta según Architecture:** en entorno de prueba autorizado, compilación Go, `go test ./...`, `go test -race ./...`, `go vet ./...`, revisión de formato, Vitest/React Testing Library y construcción frontend con scripts previamente revisados. Validar migraciones SQL de subida y recuperación en BD desechable, no sobre datos reales; comprobar `GET /health`, errores OpenAPI y arranque Docker Compose. No se ejecutan por figurar aquí ni se inventa cobertura.
- **Evidencia de salida:** commit probado, entorno, comandos, resultados y limitaciones; URL si se despliega con autorización; restauración en base de prueba; reporte del piloto sin identificadores.
- **Control de alcance antes de declarar MVP completo:** revisar uno por uno los Must del PRD. RF16 necesita dashboard básico si conserva esa prioridad; RF17–19 necesitan la entrega de datos acordada en P5 antes del cierre, o un recorte explícitamente aprobado. RF22 básico no equivale al motor de fórmulas completo. P4 puede cerrar el recorrido básico, pero no dar por cumplidos esos pendientes.
- **Criterio de paso:** persona distinta del autor reproduce el recorrido y puede explicar datos, cálculos y permisos. No declarar “listo” por capturas de pantalla solamente.

### P5 — Índice y aprendizaje: línea separada, no bloqueo de primera entrega

- **Calendario:** línea de trabajo coordinada con P2–P4, no necesariamente posterior a despliegue. Iniciar definición de datos/permisos en P0; usar registros reales cuando estén disponibles. Si RF17–19 conservan prioridad Must/MVP, su entrega acordada precede a la aceptación del MVP completo. No se declara eliminado del semestre.
- **RF:** RF15, RF17–19; RF16 básico se construye en P2 y sus ampliaciones históricas se evalúan aparte. RF23 sigue posterior mientras producto no apruebe otro alcance.
- **Entrada:** registros comparables, procedencia y permisos verificables. No basta con quitar nombres.
- **Trabajo inicial:** contrato de comparabilidad; control de calidad; exclusión de datos sintéticos; política de cohortes y reidentificación; evaluación del comprador; línea base descriptiva privada.
- **Base del modelo:** `data_sharing_consents` registra permiso empresarial y `benchmark_index_records` almacena únicamente resultados agregados. Conciliar con Architecture la selección y depuración privada previa, retiro del aporte y cálculo sin exponer identificadores; no crear consentimiento adicional en `users`. Definir filtros/períodos antes de ampliar la clave agregada del modelo.
- **ML:** definir objetivo (costo real por unidad terminada), variables disponibles al cotizar, cortes temporales y separación por obra/empresa cuando corresponda; medir MAE monetario por unidad frente al APU presupuestado y una referencia histórica simple. No usar el costo final como entrada.
- **Línea base de Rules:** los promedios ponderados propuestos pueden evaluarse como referencia estadística, con pesos, cohortes y errores documentados. No presentarlos como modelo supervisado validado ni asumir que superar 500 OT habilita ML: importan diversidad, comparabilidad y evaluación temporal por categoría.
- **Publicación:** solo con umbrales de diversidad, precisión y privacidad aprobados; si no se cumplen, mostrar “datos insuficientes”. No inventar un mínimo universal de registros.
- **Pruebas de salida:** abstención con muestra insuficiente; unidad/rubro incompatible excluido; empresa sin permiso no contribuye; evaluación reproducible contra línea base sin fuga de datos.
- **Decisión:** si no hay datos adecuados, entregar diagnóstico y plan verificable, y confirmar con el profesor el impacto sobre la entrega. No presentar datos simulados como validación real ni promedios como ML entrenado.

### 3.1. Caso numérico propuesto para el primer recorrido

Este caso es sintético y define una prueba a acordar con producto/datos; no representa una empresa entrevistada ni una fórmula ya aprobada en el PRD. Importes netos CLP, sin impuestos de venta en este ejemplo, con margen objetivo sobre venta y no recargo sobre costo. Se coordina con los `NUMERIC` del modelo, sin convertirlos en punto flotante binario:

| Paso | Resultado esperado |
|---|---|
| Presupuestar 10 m² con APU directo de $1.000/m² | Costo directo previsto $10.000 |
| Añadir gastos generales previstos de 10 %, sin otros indirectos | GG $1.000; costo previsto total $11.000 |
| Aplicar margen objetivo de 20 % sobre venta | Precio neto = $11.000 / (1 − 0,20) = $13.750 |
| Registrar ejecución completa: directo $11.000 y GG real $1.000 | Costo real total $12.000, sin duplicar GG entre partidas y OT |
| Comparar costos directos | Desviación +$1.000; +10 % |
| Comparar costos totales | Desviación +$1.000; +9,09 % de visualización |
| Consultar resultado final del trabajo | Resultado $1.750; margen real 12,73 % sobre precio neto |

**Prueba adicional de acumulados:** dos logs con cantidad 2 y 3, y `actual_cost` $3.000 y $2.000 como importes incrementales, producen cantidad total 5 y costo directo total $5.000. El costo unitario observado es $1.000/unidad; no multiplicar 5 × $5.000. Si los responsables eligen almacenar costo unitario en lugar de importe total, deben renombrar/definir el contrato y ajustar pruebas/API conjuntamente antes de implementar.

Calcular con representación monetaria exacta y no redondear porcentajes intermedios. Acordar precisión y redondeo antes de implementar. Si faltan costos aplicables, no mostrar margen definitivo; si la base prevista es cero, mostrar porcentaje no calculable. Actualizar catálogo después de aceptar no cambia estos valores históricos. El PRD deberá especificar estas reglas o una alternativa acordada antes de convertir el caso en aceptación definitiva.

## 4. Cobertura y dependencias

| Grupo RF | Primera entrega | Ampliación | Dependencias principales |
|---|---|---|---|
| RF01–04 | Catálogo mínimo y precios propios en P1–P2 | Históricos/filtros avanzados posteriores | Unidades, empresa y APU |
| RF05, RF05.1 | Cotizar y calcular en P1 | Mejoras de reutilización en P2 | Catálogo, proyecto, reglas monetarias |
| RF06 | Una versión base en el modelo existente, sin UI de revisiones/diff | Nuevas versiones comerciales y diff en P3 | Snapshots y revisión concurrente |
| RF07 | Diferido del primer recorrido | P3 si se aprueba; sigue Should en PRD | Aceptación, cambios contractuales |
| RF08 | Después del primer recorrido | Vigencia simple en P3, por su prioridad Must vigente | Fecha, estado comercial y snapshot |
| RF09–10 | Aceptar y crear trabajo en P1 | Envío/exportación e historial en P3 | Versión exacta, permisos, idempotencia |
| RF11–12, RF14 | Recorrido P1, robustez P2 | Consolidación y pruebas P4 | Partida aceptada, ejecución incremental |
| RF13 | No bloquea el recorrido | Posterior según alcance aprobado | Archivos seguros y permisos |
| RF16 | No bloquea P1 | Dashboard básico en P2 antes de aceptación del MVP; ampliaciones posteriores | Métricas completas, período y datos privados |
| RF15 | No predicción | Evaluación en P5 con alcance académico acordado | Históricos comparables y línea base |
| RF17–19 | No publicación; datos privados | P5 coordinada con P2–P4; pendiente del MVP vigente salvo recorte aprobado | Datos reales, permiso, calidad y privacidad |
| RF20–21 | Acceso, empresa y proyecto en P1 | Administración ampliada posterior | Contrato de identidad |
| RF22 | Rubro/categoría/unidad con cálculo fijo, subconjunto propuesto | Motor completo requiere fase/contrato o recorte aprobado; no dar RF22 por terminado con el básico | Modelo extensible y evaluación segura de fórmulas |
| RF23 | Fuera | Futuro | Validación de demanda y autorización pública |

## 5. Definition of Ready — cuándo empezar una tarea

- [ ] Tiene RF, alcance de entrega y criterio observable.
- [ ] Tiene responsable y revisor, asignados por el equipo.
- [ ] Entradas, salida, permisos, errores y ejemplo numérico están definidos.
- [ ] Los compañeros afectados revisaron el contrato; bloqueos están visibles.
- [ ] No contiene una decisión de negocio escondida como detalle técnico.
- [ ] Los datos de prueba son sintéticos o cuentan con autorización y anonimización apropiadas.

## 6. Definition of Done — cuándo cerrar una tarea

- [ ] Recorre las capas que necesita, incluida interfaz si es tarea de usuario.
- [ ] Se verificó éxito, validación fallida, acceso indebido y reintento/concurrencia cuando aplica.
- [ ] Migración, contrato y decisiones están revisados con sus responsables; los nombres/precisión/relaciones se concilian con el DBML, sin tratarlo como migración ya aplicada.
- [ ] Hay pruebas ejecutadas y su salida, o impedimento explícito; lectura no se informa como ejecución.
- [ ] No filtra datos de otras empresas, costos internos ni identificadores de entrevistados.
- [ ] No sobrescribe versiones aceptadas ni duplica costos por reintento.
- [ ] El revisor puede explicar el cambio y reproducir la demostración.
- [ ] El estado refleja lo comprobado: documentado, implementado o validado no son sinónimos.

## 7. Ficha mínima por tarea

Registrar en la herramienta que acuerde el equipo, no crear un backlog paralelo obligatorio:

Aplicar el flujo de ramas, mensajes y revisión por PR definido en Rules §7 y la evidencia de su Definition of Done. Este plan no crea ramas, commits ni PR ni autoriza operaciones Git automáticamente. Al preparar el incremento, coordinar archivos con sus dueños; no sobrescribir PRD, Rules, arquitectura, datos o diseño para hacer coincidir una propuesta sin revisión.

```text
ID / RF / entrega:
Resultado observable:
Responsable y revisor:
Contratos y dependencias acordados:
Caso de éxito, error, permisos y reintento:
Archivos autorizados para esta tarea:
Evidencia: commit, entorno, comandos y resultados:
Decisiones / bloqueos / siguiente paso:
```

## 8. Estado inicial y gestión del avance

| Fase | Estado al redactar | Responsable | Evidencia para actualizar |
|---|---|---|---|
| P0 | En revisión documental; ejecución de validación por comprobar | Por asignar | Alcance aprobado, prototipo, resultados reales |
| P1 | No iniciada según archivos inspeccionados | Por asignar | Primer recorrido y pruebas |
| P2 | No iniciada | Por asignar | Captura incremental y piloto |
| P3 | No iniciada | Por asignar | Versiones, aceptación y OT |
| P4 | No iniciada; aceptación del MVP sujeta a cobertura Must | Por asignar | Despliegue, seguridad, restauración y matriz de aceptación |
| P5 | Alcance/datos pendientes; coordinación con P2–P4 por acordar | Por asignar | Protocolo de datos, dashboard privado, permisos y evaluación |

Actualizar estados solo con evidencia. Si el equipo no alcanza la planificación, registrar causa y negociar alcance antes de añadir trabajo; no recortar autenticación, aislamiento, integridad monetaria o trazabilidad para simular avance.

**Corte de esta revisión (28-10-2026):** arquitectura y modelo de datos ya son diseños desarrollados; interfaz visual sigue como plantilla y no hay código de aplicación, manifiestos ni migraciones localizados en el árbol revisado. Por eso mejora la preparación de P0, no el estado de implementación de P1–P5. Antes de subir este archivo, revisar su diff y enviarlo por el flujo documental/PR acordado en Rules; no incluir por accidente cambios ajenos del árbol de trabajo.
