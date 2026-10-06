# Architecture — Arquitectura técnica del sistema

> Plantilla reutilizable para describir stack, estructura, capas, decisiones técnicas y scaffolding inicial. Debe responder cómo se construye el sistema, no qué funcionalidades tiene.

## 1. Resumen arquitectónico

```text
[Producto] usa [estilo arquitectónico] con [frontend], [backend], [base de datos] e integraciones [externas]. La prioridad técnica es [simplicidad / escalabilidad / auditabilidad / rapidez MVP / modularidad].
```

## 2. Objetivos técnicos

- `[Objetivo técnico 1: ej. entregar MVP rápido sin sobre-ingeniería]`
- `[Objetivo técnico 2: ej. mantener reglas de negocio testeables]`
- `[Objetivo técnico 3: ej. permitir despliegue simple]`

## 3. Restricciones

| Restricción | Impacto |
| --- | --- |
| `[Lenguaje/stack obligatorio]` | `[Cómo condiciona la arquitectura]` |
| `[Tiempo/equipo]` | `[Cómo condiciona alcance técnico]` |
| `[Infraestructura]` | `[Cloud, on-premise, Docker, etc.]` |
| `[Compliance]` | `[Privacidad, auditoría, trazabilidad]` |

## 4. Stack tecnológico

| Capa | Tecnología | Motivo |
| --- | --- | --- |
| Frontend | `[React / Vue / Angular / Next / etc.]` | `[Motivo]` |
| Backend | `[Go / Node / Python / Java / etc.]` | `[Motivo]` |
| HTTP framework | `[Fiber / Gin / Express / FastAPI / etc.]` | `[Motivo]` |
| Base de datos | `[PostgreSQL / MySQL / etc.]` | `[Motivo]` |
| ORM / SQL | `[GORM / Prisma / SQLAlchemy / raw SQL]` | `[Motivo]` |
| Auth | `[JWT / sessions / OAuth / SSO]` | `[Motivo]` |
| Jobs | `[Postgres queue / Redis / SQS / BullMQ / Celery]` | `[Motivo]` |
| Testing | `[Vitest / Jest / Go test / Pytest / etc.]` | `[Motivo]` |
| Deploy | `[Docker / VPS / Kubernetes / serverless]` | `[Motivo]` |

## 5. Estilo arquitectónico

### Patrón principal

`[Layered Architecture / Clean Architecture / Hexagonal / Modular Monolith / Microservices / etc.]`

### Flujo recomendado

```text
Request
  -> Router / Controller / Handler
  -> Service / Use Case
  -> Repository / Gateway
  -> Database / External Provider
```

### Responsabilidades por capa

#### Router / Handler / Controller

- Parsear request.
- Validar formato básico.
- Extraer contexto de autenticación.
- Llamar al caso de uso.
- Traducir errores a HTTP.
- No contener reglas de negocio.

#### Service / Use Case

- Implementar reglas de negocio.
- Validar permisos específicos del recurso.
- Orquestar transacciones.
- Aplicar idempotencia.
- Emitir eventos, jobs o auditoría.

#### Repository / Gateway

- Encapsular queries y persistencia.
- No decidir reglas de negocio.
- Exponer métodos alineados al dominio.

#### Jobs / Workers

- Ejecutar tareas diferidas.
- Reintentar operaciones externas.
- Registrar estado e intentos.
- No bloquear respuestas HTTP.

## 6. Scaffolding inicial

Adapta esta estructura al stack real del proyecto.

```text
/
  documentation/
    PRD.md
    Design.md
    DB_Diagram.md
    Architecture.md
    Backend_Implementation_Phases.md
    Rules.md
  backend/
    cmd/
      server/
        main.*
    internal/
      config/
      domain/
      handlers/
      services/
      repositories/
      middlewares/
      jobs/
      migrations/
      tests/
  frontend/
    src/
      app/
      features/
      components/
      layouts/
      hooks/
      lib/
      styles/
  infra/
    docker/
    scripts/
```

## 7. Configuración y ambientes

| Ambiente | Uso | Base de datos | Integraciones |
| --- | --- | --- | --- |
| Local | Desarrollo individual | `[local container]` | `[mocks/sandbox]` |
| Test | CI | `[efímera]` | `[mocks]` |
| Staging | Validación funcional | `[staging]` | `[sandbox]` |
| Producción | Usuarios reales | `[prod]` | `[prod]` |

### Variables de entorno

| Variable | Requerida | Descripción | Ejemplo |
| --- | --- | --- | --- |
| `APP_ENV` | Sí | Ambiente de ejecución | `local` |
| `DATABASE_URL` | Sí | Conexión a BD | `postgres://...` |
| `JWT_SECRET` | Sí | Firma de tokens | `***` |
| `[VAR]` | `[Sí/No]` | `[Descripción]` | `[Ejemplo]` |

## 8. Seguridad

- Autenticación: `[estrategia]`.
- Autorización: `[RBAC / ABAC / scopes / ownership]`.
- Secretos: nunca commitear; usar variables de entorno o secret manager.
- Contraseñas/tokens: guardar hash, no texto plano.
- Inputs: validar en backend aunque el frontend valide.
- Auditoría: registrar acciones sensibles.

## 9. Transacciones e idempotencia

Toda operación que modifique más de una tabla o recurso externo debe definir:

- Límite transaccional.
- Qué ocurre si falla a mitad.
- Si se puede reintentar.
- Clave de idempotencia si aplica.
- Eventos/auditoría emitidos.

## 10. Manejo de errores

Formato recomendado:

```json
{
  "error": {
    "code": "RESOURCE_NOT_FOUND",
    "message": "Mensaje legible para el cliente",
    "details": {},
    "request_id": "..."
  }
}
```

Los servicios devuelven errores de dominio; la capa HTTP los traduce a status codes.

## 11. Observabilidad

- Logs estructurados con `request_id`.
- Métricas mínimas: latencia, errores, throughput.
- Healthcheck de API y base de datos.
- Trazas si el sistema lo justifica.

## 12. Decisiones arquitectónicas

| ID | Decisión | Contexto | Consecuencia | Estado |
| --- | --- | --- | --- | --- |
| ADR-001 | `[Decisión]` | `[Por qué se tomó]` | `[Trade-off]` | `Aprobada` |

## 13. Checklist de arquitectura

- [ ] El stack está definido y justificado.
- [ ] Las capas tienen responsabilidades claras.
- [ ] Existe scaffolding inicial.
- [ ] La estrategia de configuración por ambiente está clara.
- [ ] Seguridad, errores y transacciones tienen reglas explícitas.
- [ ] Hay una sección para decisiones técnicas.
- [ ] La arquitectura no contradice el PRD ni el modelo de datos.
