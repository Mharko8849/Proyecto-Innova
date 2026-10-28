# Rules — Reglas de trabajo, desarrollo y mantenimiento del proyecto

> **Proyecto:** ÍndicePartidas  
> **Asignatura:** Innovación y Emprendimiento 2-2026 — USACH  
> **Equipo:** Grupo 1  
> **Integrantes:** Marco Ortiz, Angel Benavides, Hector Berguño, Vicente Carrillo, Martin Fuentes, Sarumi Mora, Luis Peluk.

---

## 1. Propósito

Este documento establece las reglas operativas, estándares de ingeniería, convenciones de código, flujo de trabajo en Git y criterios de calidad para el desarrollo y mantenimiento de **ÍndicePartidas**.

Su objetivo es asegurar que todos los miembros del Grupo 1 y cualquier asistente de desarrollo (IA o colaboradores externos) implementen funcionalidades y corrijan incidencias con consistencia arquitectónica, preservando siempre la confidencialidad de los datos comerciales, la inmutabilidad de presupuestos y la trazabilidad del ciclo **presupuestar → ejecutar → comparar → aprender**.

---

## 2. Fuente de verdad por tema

Antes de modificar código o documentación, se debe consultar el documento rector correspondiente:

| Pregunta / Ámbito | Documento Rector | Ubicación |
| :--- | :--- | :--- |
| ¿Qué producto se construye, para quién y qué requerimientos funcionales cubre (RF01–RF23)? | [PRD.md](PRD.md) | `coding-template-main/PRD.md` y [requerimientos-funcionales.md](../requerimientos-funcionales.md) |
| ¿Cuáles son los criterios de rendimiento, seguridad, usabilidad y calidad (RNF01–RNF30)? | [requerimientos-no-funcionales.md](../requerimientos-no-funcionales.md) | Raíz del proyecto |
| ¿Cómo debe verse, estructurarse y responder la UI/UX web y móvil? | [Design.md](Design.md) | `coding-template-main/Design.md` |
| ¿Cuál es el modelo de datos relacional, entidades, tipos y constraints en PostgreSQL? | [DB_Diagram.md](DB_Diagram.md) | `coding-template-main/DB_Diagram.md` |
| ¿Cómo se organiza el sistema técnico, capas, stack y flujos de control? | [Architecture.md](Architecture.md) | `coding-template-main/Architecture.md` |
| ¿En qué secuencia vertical se implementan los módulos del sistema? | [Implementation_phases.md](Implementation_phases.md) | `coding-template-main/Implementation_phases.md` |
| ¿Cuáles son las evidencias de entrevistas, validación de clientes y propuesta de valor? | Trabajos y Bitácoras de Clase | `Trabajos-Clases/` |
| ¿Cuáles son las reglas de desarrollo, commits, testing y mantenimiento? | [Rules.md](Rules.md) | `coding-template-main/Rules.md` |

---

## 3. Regla de actualización documental

Ningún cambio sustancial de software se considera terminado si no se actualiza la documentación rectora correspondiente:

| Tipo de cambio | Documento obligatorio a actualizar | Puntos clave a revisar |
| :--- | :--- | :--- |
| **Nueva funcionalidad o flujo de negocio** | `PRD.md`, `Implementation_phases.md` | RF asociado, criterios de aceptación y prioridad de fase. |
| **Nuevo requerimiento técnico o umbral de calidad** | `requerimientos-no-funcionales.md`, `Architecture.md` | RNF impactado (latencia, seguridad, compatibilidad). |
| **Cambio de UI, componente o interacción en terreno** | `Design.md` | Flujo móvil (< 30s en terreno), estados de carga, tokens. |
| **Nueva tabla, relación, columna o índice** | `DB_Diagram.md` | Constraints relacionales, soft delete, inmutabilidad de versiones. |
| **Modificación de stack, middlewares o capas** | `Architecture.md` | Mapeo de responsabilidades (Router → Service → Repo). |
| **Cambio en el proceso del equipo o normas de PR** | `Rules.md` | Convenciones, checklists y estándares de validación. |

---

## 4. Flujo estándar para implementar una funcionalidad

Para garantizar la entrega por flujo vertical utilizable, cada funcionalidad (ej. Catálogo de Partidas, Versionado de Cotizaciones, Cierre de OT, Anonimización de Índice) debe seguir esta secuencia:

1. **Revisión del requerimiento:** Leer el RF en [PRD.md](PRD.md) o [requerimientos-funcionales.md](../requerimientos-funcionales.md) y los RNF aplicables (ej. RNF-02 de registro rápido en terreno, RNF-11 de confidencialidad comercial).
2. **Confirmación del modelo de datos:** Verificar entidades y campos en [DB_Diagram.md](DB_Diagram.md). Si requiere migración en PostgreSQL, definir constraints e índices necesarios.
3. **Definición del contrato de API:** Especificar endpoints REST (método HTTP, ruta, payload JSON, headers de autenticación y respuestas de error estandarizadas).
4. **Implementación de Repositorio (Capa de Acceso a Datos):** Construir queries parametrizadas con soporte de transacciones ACID para operaciones críticas (ej. creación de cotización con partidas y APU).
5. **Implementación de Servicio / Caso de Uso (Lógica de Dominio):** Codificar reglas de negocio puras (cálculo de APU, aplicación de márgenes, bloqueo de modificación en cotizaciones aceptadas, anonimización).
6. **Implementación de Handler / Controlador:** Parsear request, validar sintaxis, extraer identidad del token JWT y delegar al caso de uso.
7. **Implementación de Frontend (React responsivo):** Construir vistas y componentes considerando visualización en escritorio y dispositivos móviles para terreno.
8. **Testing:** Añadir tests unitarios para las reglas de cálculo y tests de integración para endpoints clave.
9. **Verificación manual y de seguridad:** Probar caso exitoso, caso con datos erróneos y verificar que usuarios de una empresa nunca puedan acceder a datos de otra.
10. **Actualización documental y registro:** Actualizar la documentación y marcar el estado en la fase correspondiente.

---

## 5. Flujo estándar para corregir bugs

### 5.1 Diagnóstico
Antes de tocar cualquier línea de código:
- Reproducir el bug en ambiente local con datos controlados.
- Identificar el comportamiento esperado según el PRD vs. el comportamiento observado.
- Localizar la capa responsable (Frontend, Handler, Servicio, Query o Base de Datos).
- Determinar si el bug compromete datos sensibles (ej. exposición de márgenes confidenciales o cálculo erróneo de APU).

### 5.2 Corrección
- Atacar la causa raíz y no aplicar parches superficiales que oculten inconsistencias.
- Escribir un test automatizado que reproduzca la falla (debe fallar antes del fix y pasar después).
- No introducir refactorizaciones fuera del alcance del bugfix.
- Si el bug se debió a ambigüedad funcional, actualizar [PRD.md](PRD.md).
- Si el bug involucra inconsistencia en datos, generar script de migración correctiva y actualizar [DB_Diagram.md](DB_Diagram.md).

### 5.3 Cierre y registro de incidencias
Todo bug resuelto debe registrarse en el commit o PR con el siguiente formato:
```text
Bug: [Descripción clara del error y módulo afectado]
Causa raíz: [Por qué ocurría el problema a nivel técnico]
Solución: [Qué se modificó para corregirlo definitivamente]
Validación: [Pruebas ejecutadas y tests agregados]
Documentación afectada: [Documentos actualizados si aplicó]
```

---

## 6. Responsabilidades por capa

El sistema respeta una separación estricta de responsabilidades (Clean / Layered Architecture):

### 6.1 Frontend (Vue / Web Responsiva)
- **Debe:**
  - Proveer interfaces rápidas y limpias, priorizando el registro ágil en terreno (< 30 segundos, RNF-02).
  - Validar tipos, rangos y campos obligatorios en el cliente para brindar retroalimentación inmediata.
  - Consumir la API REST mediante JSON y manejar adecuadamente los estados de UI: *Cargando (Loading)*, *Vacío (Empty)*, *Error* y *Éxito*.
  - Diseñar bajo enfoque responsivo (móvil, tablet y escritorio).
- **No debe:**
  - Ser la única barrera de autorización o seguridad (las validaciones deben repetirse en el backend).
  - Calcular márgenes finales de venta o precios oficiales en el cliente sin validación del servidor.
  - Almacenar tokens o información sensible en lugares inseguros sin control de expiración.

### 6.2 Backend Handlers / Controllers (API REST en Go)
- **Debe:**
  - Parsear el cuerpo de la petición (JSON) y validar sintaxis de entrada.
  - Extraer y validar el contexto de autenticación (Claims del JWT: `user_id`, `company_id`, `role`).
  - Traducir parámetros HTTP a estructuras de dominio y llamar a los Servicios.
  - Retornar códigos de estado HTTP semánticos (200, 201, 400, 401, 403, 404, 409, 500) con formato de error uniforme.
- **No debe:**
  - Contener reglas de negocio (cálculo de costos, fórmulas de APU o porcentajes de desviación).
  - Ejecutar queries SQL directamente contra la base de datos.
  - Decidir permisos complejos sin consultar al servicio o middleware de autorización.

### 6.3 Services / Casos de Uso (Capa de Dominio)
- **Debe:**
  - Implementar todas las reglas de negocio de ÍndicePartidas:
    - Desglose y cálculo de APU (materiales + mano de obra + rendimiento).
    - Aplicación confidencial de gastos generales y márgenes comerciales.
    - Mecanismo de versionado de cotizaciones (inmutabilidad de versiones aceptadas).
    - Conversión en un clic de Cotización aceptada a Orden de Trabajo (OT).
    - Cálculo de Desviación Cotizado vs. Ejecutado (`desviación = costo_real - costo_presupuestado`).
    - Algoritmo de anonimización y depuración de outliers para el ÍndicePartidas (RF18, RF19).
  - Garantizar la confidencialidad multitenant: un usuario sólo puede operar sobre recursos de su propia empresa.
  - Orquestar transacciones atómicas que afecten múltiples entidades.
- **No debe:**
  - Depender de detalles HTTP (no usar `http.Request`, `http.ResponseWriter` ni frameworks web dentro del servicio).

### 6.4 Repositories / Acceso a Datos (PostgreSQL)
- **Debe:**
  - Encapsular todas las sentencias SQL y operaciones de persistencia.
  - Utilizar consultas parametrizadas para prevenir inyecciones SQL.
  - Respetar la integridad referencial y las llaves foráneas.
  - Asegurar la inmutabilidad histórica requerida por RNF-16 (no borrar cotizaciones históricas ni registros de costos ejecutados).
  - Filtrar obligatoriamente por `company_id` en todas las consultas de entidades privadas.
- **No debe:**
  - Aplicar reglas de negocio o cálculos de rentabilidad dentro del repositorio.

---

## 7. Reglas de Git, Ramas y Commits

Para coordinar eficazmente el trabajo de los 7 integrantes del Grupo 1, se establecen las siguientes convenciones:

### 7.1 Estructura de Ramas
- `main`: Rama de producción y entregas académicas evaluadas. Siempre estable y funcional.
- `develop`: Rama de integración continua del equipo.
- `feature/<modulo>-<nombre>`: Nuevas funcionalidades (ej. `feature/catalogo-apu`, `feature/orden-trabajo`).
- `fix/<modulo>-<nombre>`: Corrección de errores (ej. `fix/calculo-margen`, `fix/login-jwt`).
- `docs/<tema>`: Cambios exclusivamente documentales (ej. `docs/actualizacion-rules`).

### 7.2 Convención de Commits (Conventional Commits)
Los mensajes de commit deben ser claros, en español o inglés consistente, indicando el módulo:
```text
feat(cotizaciones): agregar soporte para control de versiones en presupuestos
fix(terreno): corregir validacion de cantidad ejecutada en partidas
docs(arquitectura): documentar flujo de anonimizacion para indicepartidas
test(apu): anadir pruebas unitarias para calculo de rendimiento de mano de obra
refactor(auth): extraer middleware de validacion de roles por empresa
```

### 7.3 Reglas de Integración
- No realizar commits directos a `main` ni a `develop`.
- Todo cambio se integra mediante Pull Request (PR) con descripción clara y referencia al RF o RNF correspondiente.
- Cada PR debe ser revisado y aprobado por al menos un compañero del equipo antes de fusionarse.
- **Prohibido terminantemente:** subir archivos `.env` con contraseñas reales, volcados de base de datos (`.sql`) con datos de empresas reales o archivos binarios pesados.

---

## 8. Reglas de testing y criterios de validación

| Módulo / Tipo de Cambio | Validación Mínima Requerida | Criterio de Éxito |
| :--- | :--- | :--- |
| **Cálculo de APU y Presupuestos** | Test unitario de servicio | Verificación exacta de costos directos, indirectos y margen alzado sin errores de redondeo. |
| **Control de versiones de cotización** | Test de integración de base de datos | Comprobar que modificar una cotización crea una nueva versión y no sobrescribe la anterior. |
| **Conversión a Orden de Trabajo** | Test de servicio y transacción | Verificación de que partidas y cubicaciones se copian íntegramente a la OT en una sola transacción. |
| **Registro de costos en terreno** | Prueba funcional y de rendimiento | Verificación de que el formulario se completa en menos de 30 segundos (RNF-02). |
| **Comparativa Cotizado vs. Ejecutado** | Test unitario con casos borde | Comprobación de cálculo con sobrecostos, ahorros y cantidades en cero. |
| **Anonimización del ÍndicePartidas** | Test de seguridad y datos | Garantizar que el payload exportado no contenga `company_id`, nombres de clientes, márgenes ni precios individuales (RNF-11, RNF-12). |
| **Endpoints REST** | Test de integración de Handler | Retorno correcto de status code (200, 201, 400, 403, 404) y estructura JSON normalizada. |

---

## 9. Manejo de decisiones técnicas (ADRs)

Toda decisión estructural que altere tecnologías, esquemas de seguridad o flujos de cálculo debe quedar registrada como un **Architectural Decision Record (ADR)** en `Architecture.md` o en la sección correspondiente con el formato:

```text
ADR-[Número]: [Título de la decisión]
Fecha: [YYYY-MM-DD]
Estado: [Propuesta / Aprobada / Reemplazada]
Contexto: [Problema o necesidad detectada en las entrevistas o requisitos del proyecto]
Opciones consideradas: [Alternativa A vs Alternativa B]
Elección: [Opción seleccionada por el equipo]
Consecuencias: [Beneficios obtenidos y compromisos/trade-offs asumidos]
```

### ADRs Iniciales del Proyecto:
- **ADR-001 (Stack Core):** Backend en **Go** y Frontend en **Vue**, con **PostgreSQL** para garantizar alto rendimiento, bajo consumo de recursos y facilidad de despliegue modular (RNF-14, RNF-19, RNF-20).
- **ADR-002 (Inmutabilidad de Cotizaciones):** Cada revisión comercial genera una nueva versión; una cotización aceptada se congela definitivamente para servir de base comparativa histórica inalterable (RNF-16).
- **ADR-003 (Doble capa de privacidad):** Los márgenes de ganancia y gastos generales son estrictamente confidenciales y jamás se comparten al ÍndicePartidas, el cual sólo recibe costos unitarios directos y rendimientos agregados (RNF-11, RNF-12).

---

## 10. Manejo de deuda técnica

La deuda técnica aceptada durante el desarrollo del MVP debe estar registrada y contener un plan explícito de resolución:

| Deuda Técnica | Motivo de Aceptación en MVP | Impacto / Riesgo | Plan de Mitigación / Pago | Fecha / Fase Límite |
| :--- | :--- | :--- | :--- | :--- |
| **Predicción de desviación simplificada** (RF15) | Falta de volumen histórico inicial de datos de entrenamiento para modelos complejos de Machine Learning. | La estimación predictiva se basará en promedios estadísticos ponderados y no en redes neuronales. | Reemplazar por modelo predictivo supervisado una vez acumuladas > 500 órdenes de trabajo finalizadas. | Fase 2 |
| **Soporte offline en terreno limitado** | Complejidad de sincronización bidireccional en primera versión de terreno. | El usuario requiere conexión a internet móvil (3G/4G) para registrar avances fotográficos e incidentes. | Implementar Service Worker y almacenamiento local (PWA/IndexedDB) para cola de subida diferida. | Fase 2 |
| **Directorio de prestadores básico** (RF23) | Foco central en el ciclo Cotizar-Ejecutar-Comparar antes que en el marketplace comercial. | La búsqueda de contratistas será por listado voluntario simple sin pasarela de pagos integrada. | Construir módulo de red y perfiles públicos verificados en etapa posterior. | Fase 3 |

---

## 11. Reglas para agentes y asistentes de IA (Pair Programming)

Cualquier agente de inteligencia artificial (Antigravity, GitHub Copilot, etc.) que asista en el desarrollo debe cumplir estrictamente con:

1. **Respeto a los documentos rectores:** Antes de generar código, verificar [PRD.md](PRD.md), [requerimientos-funcionales.md](../requerimientos-funcionales.md), [requerimientos-no-funcionales.md](../requerimientos-no-funcionales.md) y [Architecture.md](Architecture.md).
2. **Prohibido asumir requerimientos:** No inventar módulos ni entidades fuera de las acordadas (ej. no crear módulos contables o de nómina salarial completa, explícitamente descartados en la sección "No objetivos" del PRD).
3. **Preservación de la confidencialidad:** Todo endpoint o query generado para el usuario debe validar el aislamiento de empresa (`company_id`). Nunca exponer márgenes o utilidades en endpoints públicos o del ÍndicePartidas.
4. **Mantenimiento del idioma:** La documentación, comentarios de dominio, commits y mensajes de error orientados al usuario deben redactarse en español.
5. **No romper contratos existentes:** Si un cambio requiere modificar la estructura de un DTO o tabla existente, debe advertirlo y actualizar la documentación correspondiente antes de proceder.

---

## 12. Checklist antes de entregar cambios (Definition of Done)

Cada funcionalidad, corrección o entrega de sprint debe cumplir con la siguiente lista de verificación:

- [ ] **Compilación y ejecución limpia:** El código compila sin errores ni advertencias graves en el entorno local.
- [ ] **Aislamiento Multitenant verificado:** La funcionalidad respeta la separación de datos entre empresas y no expone información confidencial.
- [ ] **Validación de flujo funcional:** Se probó manualmente el flujo principal (caso feliz) y al menos dos casos de error o borde.
- [ ] **Pruebas automatizadas incluidas:** Se agregaron pruebas unitarias para la lógica de cálculo o pruebas de integración para los nuevos endpoints.
- [ ] **Rendimiento y Usabilidad:** La interfaz cumple con el criterio de registro simplificado en terreno (< 30s) y tiempos de respuesta ágiles.
- [ ] **Documentación sincronizada:** Se actualizaron `PRD.md`, `DB_Diagram.md`, `Architecture.md` o `Rules.md` si hubo cambios en modelos, contratos o reglas.
- [ ] **Commits y ramas según convención:** La rama sigue el formato `feature/...` o `fix/...` y los mensajes de commit son descriptivos.
- [ ] **Revisión de código:** El PR cuenta con la revisión de al menos un integrante del equipo.
