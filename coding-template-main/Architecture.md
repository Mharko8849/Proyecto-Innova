# Architecture — Arquitectura técnica del sistema

> **Proyecto:** ÍndicePartidas
> **Versión del documento:** `v0.1`
> **Fecha:** `2026-10-27`
> **Estado:** `En definición`

---

## 1. Resumen arquitectónico

```text
ÍndicePartidas usa una arquitectura de Monolito Modular (Layered Architecture) con un
frontend React + JavaScript desacoplado, un backend Go con API REST, y PostgreSQL como
base de datos relacional. La prioridad técnica es entregar el MVP con simplicidad operativa,
manteniendo reglas de negocio testeables y aisladas, con una estructura que permita escalar
hacia Fase 2 (ML, índice avanzado) sin refactorizar el núcleo.
```

---

## 2. Objetivos técnicos

- **Entregar un MVP funcional** sin sobre-ingeniería, cubriendo el ciclo completo: cotización -> ejecución -> comparación.
- **Aislar las reglas de negocio** en la capa de servicios/casos de uso, haciendo el sistema testeable con independencia del framework.
- **Mantener el frontend desacoplado** del backend para permitir modificaciones independientes en cada capa.
- **Soportar registro móvil de bajo rozamiento:** el sistema debe responder en <= 3 segundos y permitir registros de ejecución en <= 30 segundos (RNF02, RNF05).
- **Garantizar confidencialidad comercial** desde el diseño: márgenes, clientes y precios privados jamás deben cruzar a otros contextos de empresa.
- **Permitir extensibilidad de rubros** sin modificar código fuente: los rubros personalizados ("Otro") se almacenan como configuración de datos (RNF25).
- **Preparar la base para el ÍndicePartidas anonimizado**, dejando los datos de ejecución estructurados para su agregación futura sin identificar empresas.

---

## 3. Restricciones

| Restricción | Impacto |
| --- | --- |
| Backend en **Go** | Condiciona el framework HTTP (Gin/Fiber/Chi), el ORM (GORM o raw SQL) y la estructura de paquetes (`cmd/`, `internal/`). |
| Frontend en **React + JavaScript** | Sin TypeScript en el MVP; tipado y migración a TS evaluable en Fase 2. |
| Base de datos: **PostgreSQL** | Integridad referencial obligatoria; queries e índices optimizados para búsquedas de partidas y ejecuciones. |
| Equipo académico (semestre) | Alcance técnico acotado al MVP; microservicios y ML diferidos a Fase 2 o postentrega. |
| Entorno de despliegue inicial | VPS o contenedor Docker; sin Kubernetes en MVP. |
| Privacidad y confidencialidad comercial | Márgenes, clientes y costos privados nunca deben ser accesibles entre empresas; el índice solo puede exponer datos agregados y anonimizados. |
| API documentada con **OpenAPI/Swagger** | Todos los endpoints públicos deben contar con documentación generada automáticamente (RNF30). |

---

## 4. Stack tecnológico

| Capa | Tecnología | Motivo |
| --- | --- | --- |
| Frontend | React + JavaScript | Definido en PRD (RNF20); ecosistema amplio, componentes reutilizables, soporte responsivo. |
| Backend | Go | Definido en PRD (RNF19); rendimiento, concurrencia nativa, compilación estática y despliegue simple. |
| HTTP framework | Gin | Liviano, maduro, amplia documentación y compatible con middleware estándar de Go. |
| Base de datos | PostgreSQL | Integridad referencial, transaccionalidad ACID, índices avanzados, escalabilidad relacional (RNF14). |
| ORM / SQL | GORM | Abstracción de queries con soporte a migraciones; raw SQL cuando se requiera optimizar (RNF07). |
| Auth | JWT (JSON Web Tokens) | Stateless, compatible con RBAC, simple de implementar en Go y validar en React (RNF08, RNF09). |
| Almacenamiento de archivos | Sistema de archivos local (MVP) / S3-compatible (Fase 2) | Fotografías y documentos adjuntos a partidas e incidencias (RF13). |
| Testing (backend) | `go test` + `testify` | Testing nativo de Go; `testify` para assertions y mocks. |
| Testing (frontend) | Vitest + React Testing Library | Testing unitario y de integración de componentes React. |
| Documentación API | Swagger / OpenAPI (`swaggo/swag`) | Generación automática desde anotaciones Go (RNF30). |
| Bundler frontend | Vite | Rápido, moderno, compatible con React + JavaScript sin configuración compleja. |
| Control de versiones | Git + GitHub | Definido en PRD (RNF29). |
| Deploy | Docker + Docker Compose | Contenedores para backend, frontend y base de datos; reproducibilidad entre ambientes. |

---

## 5. Estilo arquitectónico

### Patrón principal

**Monolito Modular con Layered Architecture** — un único proceso de backend organizado en capas bien separadas por responsabilidad. Los módulos funcionales (Catálogo, Cotizaciones, Ejecución, Índice, etc.) son carpetas dentro de `internal/`, no servicios independientes.

> Se elige el monolito modular sobre microservicios por la restricción de equipo y tiempo. La separación clara de capas permite evolucionar hacia microservicios en Fase 2 si fuera necesario, sin reescribir la lógica de negocio.

### Flujo de una solicitud HTTP

```text
Cliente (React) -- HTTPS
  -> Router (Gin)
    -> Middleware (Auth JWT / RBAC / Logger)
      -> Handler / Controller
        -> Service / Use Case   (reglas de negocio, validaciones de dominio)
          -> Repository          (queries PostgreSQL via GORM)
            -> Base de datos (PostgreSQL)
```

### Responsabilidades por capa

#### Router / Handler

- Parsear y deserializar el request (JSON -> struct).
- Validar formato y presencia de campos requeridos.
- Extraer identidad y rol del token JWT.
- Invocar el caso de uso correspondiente.
- Serializar la respuesta o traducir errores de dominio a HTTP status codes.
- **No** contiene lógica de negocio ni queries directas.

#### Service / Use Case

- Implementar las reglas de negocio del dominio (p. ej.: solo se puede crear OT desde cotización Aceptada; los márgenes solo son visibles por roles autorizados).
- Validar permisos específicos del recurso (ownership por empresa).
- Orquestar transacciones que afectan múltiples entidades.
- Aplicar idempotencia en operaciones críticas.
- Emitir auditoría de acciones sensibles.

#### Repository / Gateway

- Encapsular toda la interacción con PostgreSQL.
- Exponer métodos orientados al dominio (p. ej.: `FindPartidaByRubroAndCategoria`), no queries crudas.
- No tomar decisiones de negocio.

#### Middleware

- **Auth:** Validar JWT y rechazar requests no autenticadas (RNF08).
- **RBAC:** Verificar que el rol del usuario tenga permiso para el recurso/acción (RNF09).
- **Logger:** Registrar `request_id`, método, ruta, duración y status code.
- **CORS:** Permitir solo orígenes configurados en `ALLOWED_ORIGINS`.

#### Jobs / Workers (Fase 2)

- Procesamiento del ÍndicePartidas: agregación y anonimización periódica de datos de ejecuciones consentidas.
- Notificaciones de vigencia vencida de cotizaciones (RF08).
- Detección de anomalías en datos del índice (RF19).

---

## 6. Scaffolding inicial

```text
proyecto-innova/
+-- documentation/
|   +-- PRD.md
|   +-- Design.md
|   +-- DB_Diagram.md
|   +-- Architecture.md          <- este archivo
|   +-- Implementation_phases.md
|   +-- Rules.md
|
+-- backend/
|   +-- cmd/
|   |   +-- server/
|   |       +-- main.go           <- punto de entrada, wiring de dependencias
|   +-- internal/
|   |   +-- config/               <- lectura de variables de entorno
|   |   +-- domain/               <- structs y errores de dominio (entidades puras)
|   |   |   +-- partida.go
|   |   |   +-- cotizacion.go
|   |   |   +-- orden_trabajo.go
|   |   |   +-- usuario.go
|   |   |   +-- empresa.go
|   |   +-- handlers/             <- controllers HTTP (Gin)
|   |   |   +-- catalogo_handler.go
|   |   |   +-- cotizacion_handler.go
|   |   |   +-- ejecucion_handler.go
|   |   |   +-- indice_handler.go
|   |   |   +-- auth_handler.go
|   |   +-- services/             <- casos de uso y reglas de negocio
|   |   |   +-- catalogo_service.go
|   |   |   +-- cotizacion_service.go
|   |   |   +-- ejecucion_service.go
|   |   |   +-- indice_service.go
|   |   |   +-- auth_service.go
|   |   +-- repositories/         <- acceso a datos (GORM / raw SQL)
|   |   |   +-- partida_repository.go
|   |   |   +-- cotizacion_repository.go
|   |   |   +-- ejecucion_repository.go
|   |   |   +-- usuario_repository.go
|   |   +-- middlewares/
|   |   |   +-- auth_middleware.go
|   |   |   +-- rbac_middleware.go
|   |   |   +-- logger_middleware.go
|   |   +-- migrations/           <- migraciones SQL ordenadas (001_init.sql, etc.)
|   |   +-- tests/                <- tests de integracion y unitarios
|   +-- docs/                     <- Swagger generado por swaggo/swag
|   +-- go.mod
|   +-- go.sum
|
+-- frontend/
|   +-- src/
|   |   +-- app/                  <- configuracion de rutas y providers globales
|   |   +-- features/             <- modulos funcionales por dominio
|   |   |   +-- catalogo/
|   |   |   +-- cotizaciones/
|   |   |   +-- ejecucion/
|   |   |   +-- indice/
|   |   |   +-- auth/
|   |   +-- components/           <- componentes reutilizables (Atomic Design)
|   |   +-- layouts/              <- estructuras de pagina (sidebar, topbar)
|   |   +-- hooks/                <- custom hooks (useAuth, usePartidas, etc.)
|   |   +-- lib/                  <- cliente HTTP, helpers, constantes
|   |   +-- styles/               <- CSS global y design tokens
|   +-- public/
|   +-- index.html
|   +-- vite.config.js
|   +-- package.json
|
+-- infra/
    +-- docker/
    |   +-- Dockerfile.backend
    |   +-- Dockerfile.frontend
    |   +-- docker-compose.yml    <- orquesta backend + frontend + postgres
    +-- scripts/
        +-- seed.sql              <- datos semilla: rubros, categorias y partidas base
        +-- migrate.sh
```

---

## 7. Configuración y ambientes

| Ambiente | Uso | Base de datos | Integraciones |
| --- | --- | --- | --- |
| **Local** | Desarrollo individual | PostgreSQL local (Docker Compose) | Sin externos / mocks |
| **Test** | CI / tests automáticos | PostgreSQL efímera en contenedor | Mocks |
| **Staging** | Validación funcional pre-entrega | PostgreSQL staging | Sandbox |
| **Producción** | Usuarios reales (piloto académico) | PostgreSQL prod | Prod |

### Variables de entorno

| Variable | Requerida | Descripción | Ejemplo |
| --- | --- | --- | --- |
| `APP_ENV` | Sí | Ambiente de ejecución | `local` / `production` |
| `DATABASE_URL` | Sí | Cadena de conexión PostgreSQL | `postgres://user:pass@localhost:5432/indice_partidas` |
| `JWT_SECRET` | Sí | Clave de firma de tokens JWT | `***` (mínimo 32 caracteres) |
| `JWT_EXPIRY_HOURS` | No | Duración de los tokens | `24` |
| `PORT` | No | Puerto del servidor HTTP | `8080` |
| `ALLOWED_ORIGINS` | Sí | Orígenes permitidos para CORS | `http://localhost:5173` |
| `UPLOAD_PATH` | No | Ruta para archivos subidos (fotos) | `./uploads` |
| `LOG_LEVEL` | No | Nivel de logging | `info` |

> Nunca commitear valores de producción. Usar `.env` local excluido por `.gitignore`. En CI/CD usar secrets del repositorio.

---

## 8. Seguridad

- **Autenticación (RNF08):** JWT firmado con `JWT_SECRET`. El token incluye `user_id`, `empresa_id` y `rol`. Expiración configurable vía `JWT_EXPIRY_HOURS`.
- **Autorización (RNF09):** RBAC (Role-Based Access Control). Roles: `dueno`, `contratista`, `admin_obra`, `jefe_terreno`, `maestro`, `consultor`, `admin_sistema`.
- **Aislamiento por empresa:** Todo recurso (cotización, partida, proyecto, OT) pertenece a una `empresa_id`. Los servicios verifican que el `empresa_id` del token coincida con el del recurso. Ningún usuario puede acceder a datos de otra empresa.
- **Confidencialidad de márgenes (RNF11):** El margen/gastos generales de una cotización solo es visible para roles `dueno` y `contratista` de la misma empresa.
- **Anonimización del índice (RNF12):** Al agregar datos al ÍndicePartidas se eliminan `empresa_id`, datos de clientes, márgenes y cualquier campo identificable. Solo se retienen métricas de costo y rendimiento por partida y rubro.
- **Contraseñas (RNF10):** Almacenadas con `bcrypt` (cost >= 12). Nunca en texto plano.
- **Inputs:** Validados en backend con independencia de la validación del frontend.
- **HTTPS (RNF13):** Obligatorio en producción. En local se usa HTTP.
- **Secretos:** Nunca commitear en el repositorio. Variables de entorno o secret manager.
- **Auditoría:** Las acciones sensibles (cambio de estado de cotización, conversión a OT, cierre de OT, modificación de consentimiento del índice) se registran con `user_id`, `empresa_id`, `timestamp` y acción.

---

## 9. Transacciones e idempotencia

Las siguientes operaciones modifican más de una entidad y deben ejecutarse dentro de una transacción PostgreSQL (RNF17):

| Operación | Tablas afectadas | Si falla a mitad | Idempotencia |
| --- | --- | --- | --- |
| Crear cotización con partidas | `cotizaciones` + `cotizacion_partidas` | Rollback: ni cotización ni líneas persisten | Clave: `(proyecto_id, version_numero)` |
| Generar nueva versión de cotización | `cotizaciones` (nueva fila) + copia de `cotizacion_partidas` | Rollback: versión nueva no existe | — |
| Conversión de cotización a OT (RF10) | `cotizaciones.estado` + `ordenes_trabajo` + `ot_partidas` | Rollback: cotización permanece en estado anterior | Verificar que OT no exista antes de crear |
| Cierre de OT (RF11) | `ordenes_trabajo.estado = Cerrada` + desviaciones calculadas | Rollback: OT permanece abierta | Solo se permite cerrar una vez |
| Consentimiento del índice (RF18) | `usuarios.consentimiento_indice` + cola de anonimización | Reversible: el usuario puede cambiar su consentimiento | — |

---

## 10. Manejo de errores

Todos los endpoints de la API devuelven errores en el siguiente formato JSON:

```json
{
  "error": {
    "code": "COTIZACION_NO_ENCONTRADA",
    "message": "La cotización solicitada no existe o no pertenece a tu empresa.",
    "details": {},
    "request_id": "req_a1b2c3d4"
  }
}
```

### Convenciones de códigos de error

| Código de dominio | HTTP Status | Descripción |
| --- | --- | --- |
| `UNAUTHORIZED` | 401 | Token ausente, inválido o expirado |
| `FORBIDDEN` | 403 | El usuario no tiene permiso para el recurso |
| `NOT_FOUND` | 404 | Recurso no existe o no pertenece a la empresa |
| `VALIDATION_ERROR` | 422 | Campos requeridos faltantes o formato incorrecto |
| `CONFLICT` | 409 | Operación no permitida por estado actual (p. ej.: OT ya cerrada) |
| `INTERNAL_ERROR` | 500 | Error no controlado; registrado en logs con `request_id` |

- Los servicios devuelven errores de dominio tipados (structs en Go).
- Los handlers los traducen a HTTP status code + cuerpo JSON.
- Los errores 5xx incluyen `request_id` para correlación en logs pero **no** exponen stack traces al cliente.

---

## 11. Observabilidad

- **Logs estructurados (JSON):** Cada request logea `request_id`, método HTTP, ruta, `empresa_id` (si autenticado), duración en ms y status code.
- **Métricas mínimas:** Latencia por endpoint, tasa de errores 5xx, requests por segundo.
- **Healthcheck:** `GET /health` retorna `200 OK` con estado de la conexión a PostgreSQL.
- **`request_id`:** Generado en el middleware de logger; propagado al response header `X-Request-ID` para correlación cliente-servidor.
- **Trazas:** No requeridas en MVP. Evaluables en Fase 2 si el sistema crece en módulos.

---

## 12. Decisiones arquitectónicas (ADR)

| ID | Decisión | Contexto | Consecuencia | Estado |
| --- | --- | --- | --- | --- |
| ADR-001 | Backend en **Go** | RNF19: stack definido en PRD; rendimiento, simplicidad de despliegue y compilación estática. | El equipo debe manejar Go; menos librerías ORM/framework que en Java o Node. | Aprobada |
| ADR-002 | Frontend en **React + JavaScript** | RNF20: stack definido en PRD; se evita TypeScript en MVP para reducir curva de entrada. | Sin inferencia de tipos en MVP; migración a TS evaluable en Fase 2. | Aprobada |
| ADR-003 | **Monolito Modular** en lugar de microservicios | Equipo acotado (semestre académico); MVP debe entregarse en tiempo limitado. | Un solo proceso desplegable; la separación por capas facilita extracción futura de servicios. | Aprobada |
| ADR-004 | **PostgreSQL** como única fuente de persistencia | RNF14–RNF17: integridad referencial, historial inmutable, transaccionalidad y respaldo. | No se usa caché ni bases NoSQL en MVP; toda consistencia vive en PostgreSQL. | Aprobada |
| ADR-005 | **JWT stateless** para autenticación | Simple de implementar en Go y validar en React; sin necesidad de store de sesiones en MVP. | Revocación de tokens requiere estrategia adicional (blacklist o expiración corta) si se necesita logout inmediato. | Aprobada |
| ADR-006 | Rubros personalizados **como datos, no como código** | RNF25, RF22: la opción "Otro" no debe requerir modificaciones al código fuente. | La estructura configurable se almacena en tablas (`rubro_config`, `concepto_config`). Mayor complejidad de modelo; mayor flexibilidad. | Aprobada |
| ADR-007 | Anonimización del índice **a nivel de servicio**, no de vista | RNF12: la privacidad debe garantizarse en la lógica, no solo en la presentación. | `IndiceService` opera sobre datos ya desidentificados; nunca recibe `empresa_id` real. | Aprobada |
| ADR-008 | Almacenamiento de fotografías **local en MVP** | RF13 requiere subida de fotos; en MVP no se justifica integración con S3. | En Fase 2 migrar a almacenamiento externo (S3-compatible) para producción real. | Aprobada |
| ADR-009 | **Vite** como bundler de frontend | Rápido, moderno y compatible con React + JavaScript sin configuración compleja. Alternativa a Create React App (deprecado). | Convención moderna; fácil de extender. | Aprobada |
| ADR-010 | Documentación API con **swaggo/swag** | RNF30: OpenAPI/Swagger obligatorio. `swaggo/swag` genera la spec desde anotaciones Go sin archivo YAML manual. | El equipo debe mantener anotaciones de Swagger actualizadas junto al código. | Aprobada |

---

## 13. Checklist de arquitectura

- [x] El stack está definido y justificado (Go + React + PostgreSQL).
- [x] Las capas tienen responsabilidades claras y separadas.
- [x] Existe scaffolding inicial con estructura de directorios.
- [x] La estrategia de configuración por ambiente está definida (variables de entorno).
- [x] Seguridad, errores y transacciones tienen reglas explícitas.
- [x] Hay sección de decisiones técnicas (ADRs).
- [x] La arquitectura es consistente con PRD y RNF.
- [ ] Confirmar framework Go (Gin vs. Fiber vs. Chi) con todo el equipo.
- [ ] Definir estrategia de migraciones (GORM AutoMigrate vs. archivos SQL versionados).
- [ ] Confirmar estrategia de revocación de JWT (blacklist en Redis o expiración corta).
- [ ] Definir cómo se ejecutará Docker Compose en el ambiente de CI.
- [ ] Completar DB_Diagram.md con el modelo de datos completo.
