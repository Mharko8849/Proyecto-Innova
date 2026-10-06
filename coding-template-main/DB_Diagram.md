# DB Diagram — Modelo de datos

> Plantilla reutilizable para documentar el modelo lógico y físico de base de datos. Puede usarse con DBML, Mermaid ERD o SQL comentado. La fuente de verdad debe ser explícita: este archivo puede ser objetivo de diseño, espejo de migraciones o ambas cosas si se separan bien.

## 1. Propósito del modelo

Este documento describe la estructura de datos necesaria para soportar el producto definido en [PRD.md](PRD.md).

| Campo | Definición |
| --- | --- |
| Motor de base de datos | `[PostgreSQL / MySQL / SQLite / SQL Server / etc.]` |
| ORM / Query layer | `[GORM / Prisma / TypeORM / SQLAlchemy / Diesel / raw SQL / etc.]` |
| Estrategia de migraciones | `[AutoMigrate / migrations SQL / Liquibase / Flyway / etc.]` |
| Estado del documento | `Objetivo / Implementado / Mixto` |

## 2. Convenciones

- Tablas en `[snake_case / plural / singular]`.
- Claves primarias: `[id bigint / uuid / cuid / etc.]`.
- Foreign keys con constraint física obligatoria salvo justificación.
- Campos de auditoría estándar: `[created_at, updated_at, deleted_at, created_by, etc.]`.
- Soft delete: `[Sí/No + campo]`.
- Timestamps en `[UTC / zona local]`.
- Dinero en `[integer minor units / decimal]`.

## 3. Glosario de entidades

| Entidad | Descripción | Dueño funcional |
| --- | --- | --- |
| `[Entidad 1]` | `[Qué representa en negocio]` | `[Módulo/actor]` |
| `[Entidad 2]` | `[Qué representa en negocio]` | `[Módulo/actor]` |

## 4. Estado del esquema

| Entidad / módulo | Estado | Notas |
| --- | --- | --- |
| `[Módulo 1]` | `Diseñado / Migrado / En producción / Obsoleto` | `[Notas]` |
| `[Módulo 2]` | `Diseñado / Migrado / En producción / Obsoleto` | `[Notas]` |

## 5. Diagrama DBML

> Mantén este bloque como modelo objetivo o sincronízalo con migraciones. Si existen diferencias, documentarlas en la sección 8.

```dbml
// Ejemplo mínimo. Reemplazar por el modelo del proyecto.

Table users {
  id uuid [pk]
  email varchar [not null, unique]
  name varchar [not null]
  status varchar [not null]
  created_at timestamp [not null]
  updated_at timestamp [not null]
}

Table roles {
  id uuid [pk]
  name varchar [not null, unique]
}

Table user_roles {
  user_id uuid [pk]
  role_id uuid [pk]
}

Ref: user_roles.user_id > users.id
Ref: user_roles.role_id > roles.id
```

## 6. Relaciones explícitas

Usa esta tabla para revisar integridad referencial sin depender sólo del diagrama visual.

| Origen | Destino | Cardinalidad | Regla de borrado | Motivo |
| --- | --- | --- | --- | --- |
| `user_roles.user_id` | `users.id` | `N:1` | `RESTRICT / CASCADE / SET NULL` | `[Motivo]` |
| `user_roles.role_id` | `roles.id` | `N:1` | `RESTRICT / CASCADE / SET NULL` | `[Motivo]` |

## 7. Índices y restricciones

| Tabla | Índice / constraint | Columnas | Tipo | Motivo |
| --- | --- | --- | --- | --- |
| `users` | `users_email_unique` | `email` | `unique` | Evitar cuentas duplicadas |
| `[tabla]` | `[nombre]` | `[columnas]` | `[btree/gin/unique/check]` | `[Motivo]` |

## 8. Diferencias entre modelo objetivo y esquema físico

| Elemento | Modelo objetivo | Esquema físico actual | Acción |
| --- | --- | --- | --- |
| `[Tabla/campo/relación]` | `[Esperado]` | `[Actual]` | `[Migrar / diferir / eliminar / revisar]` |

## 9. Reglas de integridad

- `[Regla 1: ej. un usuario activo debe tener email verificado]`
- `[Regla 2: ej. no puede existir inscripción sin evento padre]`
- `[Regla 3: ej. fechas de fin deben ser posteriores a fechas de inicio]`

## 10. Datos semilla

| Semilla | Tabla | Ambiente | Motivo |
| --- | --- | --- | --- |
| `[Roles base]` | `[roles]` | `dev/test/prod` | `[Control de acceso]` |
| `[Catálogo]` | `[tabla]` | `dev/test/prod` | `[Motivo]` |

## 11. Migraciones

### Nueva tabla

1. Crear migración.
2. Agregar tabla con PK y timestamps.
3. Agregar constraints e índices.
4. Agregar FKs físicas.
5. Agregar rollback si la estrategia lo exige.
6. Agregar prueba de migración o smoke test.
7. Actualizar este documento.

### Cambio de columna

1. Verificar datos existentes.
2. Hacer migración compatible hacia adelante.
3. Rellenar datos si aplica.
4. Recién después aplicar `NOT NULL` o constraints estrictas.
5. Actualizar modelos, DTOs y contratos.

## 12. Checklist de base de datos

- [ ] Todas las entidades del PRD tienen representación clara o justificación para no tenerla.
- [ ] Las claves foráneas están declaradas explícitamente.
- [ ] Hay índices para búsquedas frecuentes y constraints de unicidad.
- [ ] Las reglas críticas no dependen sólo del frontend.
- [ ] El modelo objetivo y el esquema físico no se contradicen sin explicación.
- [ ] Los datos semilla están documentados.
- [ ] Las migraciones son reversibles o tienen plan de recuperación.
