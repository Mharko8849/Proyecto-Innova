# Backend Implementation Phases — Plan de implementación por fases

> Plantilla reutilizable para construir el backend por módulos verticales. Cada fase debe entregar funcionalidad usable de extremo a extremo: base de datos, dominio, API, tests y documentación.

## 1. Criterio de trabajo

Trabajar por fases evita implementar piezas sueltas sin valor funcional. Cada fase debe cerrar un flujo vertical.

```text
Modelo/Migración -> Repository -> Service/Use Case -> Handler/API -> Tests -> Documentación
```

## 2. Definición de terminado por fase

Una fase está terminada cuando:

- [ ] Las migraciones/modelos existen y corren en local.
- [ ] Las reglas de negocio viven en servicios/casos de uso.
- [ ] Los endpoints o comandos necesarios están implementados.
- [ ] Hay validación de entrada y manejo de errores.
- [ ] Hay tests unitarios o de integración proporcionales al riesgo.
- [ ] La documentación afectada fue actualizada.
- [ ] El flujo puede probarse manualmente con datos mínimos.

## 3. Fase 0 — Base técnica

**Objetivo:** dejar el proyecto listo para desarrollar funcionalidades.

**Incluye:**

- Scaffolding backend.
- Configuración por ambiente.
- Conexión a base de datos.
- Migraciones base.
- Logger.
- Error handling común.
- CORS / security headers si aplica.
- Healthcheck.
- Setup de tests.

**Entregable:** API arranca localmente, conecta a BD y responde healthcheck.

## 4. Fase 1 — Identidad, usuarios y acceso

**Objetivo:** habilitar autenticación, autorización y administración mínima de usuarios.

**Incluye:**

- Modelo de usuario.
- Roles/permisos.
- Login/logout.
- Sesiones o tokens.
- Middleware de autenticación.
- Middleware o guardas de autorización.
- CRUD mínimo de usuarios si aplica.
- Recuperación/cambio de contraseña si aplica.

**Entregable:** rutas protegidas por rol y usuarios operables.

**Checklist específico:**

- [ ] Contraseñas con hash seguro.
- [ ] Tokens/sesiones revocables según necesidad.
- [ ] Rutas sensibles protegidas en backend.
- [ ] Tests de acceso permitido/denegado.

## 5. Fase 2 — Núcleo de dominio

**Objetivo:** implementar las entidades centrales del producto.

**Incluye:**

- CRUD de entidades principales.
- Reglas de negocio centrales.
- Validaciones de consistencia.
- Relaciones principales.
- Búsqueda/listado/paginación.

**Entregable:** el equipo puede crear, consultar y modificar el objeto principal del negocio.

**Requerimientos vinculados:** `[RFxx, RFyy]`

## 6. Fase 3 — Flujos transaccionales

**Objetivo:** implementar operaciones que modifican varias entidades o requieren consistencia fuerte.

**Incluye:**

- Transacciones ACID.
- Idempotencia.
- Control de concurrencia.
- Validaciones cruzadas.
- Auditoría de acciones críticas.

**Entregable:** flujos críticos funcionan sin duplicados ni estados intermedios corruptos.

**Preguntas obligatorias:**

- ¿Qué tablas cambia el flujo?
- ¿Qué pasa si falla el segundo paso?
- ¿Se puede reintentar?
- ¿Cómo se evita duplicidad?
- ¿Qué se audita?

## 7. Fase 4 — Operación, reportes e integraciones

**Objetivo:** agregar capacidades que conectan el sistema con procesos externos o vistas operativas.

**Incluye:**

- Reportes.
- Exportaciones.
- Integraciones externas.
- Webhooks.
- Emails/notificaciones.
- Jobs asíncronos.

**Entregable:** procesos no interactivos y reportes disponibles sin bloquear la API.

## 8. Fase 5 — Endurecimiento

**Objetivo:** preparar el backend para uso real.

**Incluye:**

- Tests de integración críticos.
- Revisión de seguridad.
- Observabilidad.
- Performance básica.
- Manejo de errores consistente.
- Limpieza de deuda técnica.
- Documentación de despliegue.

**Entregable:** backend estable para staging/producción.

## 9. Plantilla por módulo

Usa esta ficha para cada módulo nuevo.

### Módulo `[Nombre]`

| Campo | Definición |
| --- | --- |
| Objetivo | `[Qué resuelve]` |
| Requerimientos | `[RFxx, RFyy, RNFzz]` |
| Actores | `[Roles/usuarios]` |
| Tablas afectadas | `[Tablas]` |
| Endpoints | `[Rutas]` |
| Riesgo | `Bajo / Medio / Alto` |

#### Implementación

1. `[Migración/modelo]`
2. `[Repository/gateway]`
3. `[Service/use case]`
4. `[Handler/controller]`
5. `[Tests]`
6. `[Documentación]`

#### Casos de prueba mínimos

- `[Caso éxito]`
- `[Caso validación]`
- `[Caso autorización]`
- `[Caso conflicto/concurrencia si aplica]`

## 10. Orden recomendado de implementación por requerimiento

1. Leer el requerimiento en [PRD.md](PRD.md).
2. Revisar entidades en [DB_Diagram.md](DB_Diagram.md).
3. Confirmar arquitectura/capas en [Architecture.md](Architecture.md).
4. Definir contrato API.
5. Crear migración/modelo.
6. Implementar repositorio.
7. Implementar servicio/caso de uso.
8. Implementar handler/controlador.
9. Agregar tests.
10. Actualizar documentación.

## 11. Registro de avance

| Fase | Estado | Fecha | Notas |
| --- | --- | --- | --- |
| Fase 0 | `Pendiente / En progreso / Completa` | `[YYYY-MM-DD]` | `[Notas]` |
| Fase 1 | `Pendiente / En progreso / Completa` | `[YYYY-MM-DD]` | `[Notas]` |
| Fase 2 | `Pendiente / En progreso / Completa` | `[YYYY-MM-DD]` | `[Notas]` |

## 12. Checklist anti-caos

- [ ] No se implementa frontend contra endpoints no documentados.
- [ ] No se agregan queries en handlers/controllers.
- [ ] No se mezclan reglas de negocio con formato HTTP.
- [ ] No se crea una tabla sin actualizar el diagrama.
- [ ] No se cierra una fase sin flujo manual verificable.
- [ ] No se marca un requerimiento como completo sin criterio de aceptación probado.
