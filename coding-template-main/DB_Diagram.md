# DB Diagram — Modelo de datos: ÍndicePartidas

> Este documento define el modelo lógico y físico de base de datos para la plataforma **ÍndicePartidas**, según los requerimientos funcionales y no funcionales especificados en [PRD.md](PRD.md).

---

## 1. Propósito del modelo

Este documento describe la estructura de datos relacional para soportar el ciclo de vida completo del negocio: **Gestión de Empresas/Usuarios $\rightarrow$ Clientes y Proyectos $\rightarrow$ Catálogo y APU $\rightarrow$ Cotizaciones Versionadas $\rightarrow$ Órdenes de Trabajo y Terreno $\rightarrow$ Comparación Analítica $\rightarrow$ Índice Estadístico Anonimizado**.

| Campo | Definición |
| --- | --- |
| Motor de base de datos | **PostgreSQL 15+** |
| ORM / Query layer | **pgx / sqlx / GORM** (Stack Go) |
| Estrategia de migraciones | **golang-migrate** (migraciones SQL numeradas `000001_init.up.sql` / `.down.sql`) |
| Estado del documento | **Objetivo (Diseño inicial MVP)** |

---

## 2. Convenciones

- **Nombres de tablas:** `snake_case`, plural (ej. `companies`, `quotation_versions`, `execution_logs`).
- **Nombres de columnas:** `snake_case` (ej. `company_id`, `created_at`, `unit_cost_snapshot`).
- **Claves primarias (PK):** `id UUID` generado con `gen_random_uuid()` para entidades transaccionales de negocio. Llaves de catálogo fijo usan `INT` (ej. `roles`, `trades`).
- **Claves foráneas (FK):** `[entidad_singular]_id`, con constraint referencial física obligatoria en motor.
- **Campos de auditoría estándar:** `created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()`, `updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()`.
- **Soft delete:** No destructivo mediante columna `deleted_at TIMESTAMPTZ NULL` o bandera `is_active BOOLEAN` en catálogos.
- **Zona horaria:** Siempre `TIMESTAMPTZ` en almacenamiento (estándar UTC).
- **Valores monetarios:** `NUMERIC(14,2)` para garantizar precisión decimal exacta sin errores de redondeo de punto flotante.
- **Cantidades y rendimientos:** `NUMERIC(12,4)` para soportar decimales finos en cubicaciones y consumos de APU.
- **Estructuras extensibles:** `JSONB` para almacenar fórmulas, variables y configuraciones dinámicas del rubro **“Otro”** (RF22).

---

## 3. Glosario de entidades

| Entidad | Descripción | Dueño funcional / Módulo |
| --- | --- | --- |
| `companies` | Empresa, contratista o persona jurídica que utiliza la plataforma (Multi-tenant). | Módulo Configuración (RF20) |
| `users` | Usuarios con acceso a la plataforma vinculados a una empresa. | Módulo Configuración (RF20) |
| `roles` | Catálogo maestro de roles del sistema (`DUENO`, `ADMIN_OBRA`, `JEFE_TERRENO`, etc.). | Módulo Configuración (RF20) |
| `user_roles` | Tabla pivote de asignación de roles a usuarios. | Módulo Configuración (RF20) |
| `clients` | Clientes finales o mandantes a quienes se les cotizan los trabajos. | Módulo Clientes (RF21) |
| `projects` | Obras o proyectos contratados que agrupan cotizaciones y trabajos. | Módulo Proyectos (RF21) |
| `trades` | Rubros base del sistema (*Construcción, Electricidad, Vidriería, Otro*). | Módulo Catálogo (RF22) |
| `trade_custom_configs` | Estructuras personalizadas de variables, unidades y fórmulas para "Otro". | Módulo Configuración (RF22) |
| `categories` | Categorías temáticas para organizar partidas dentro de un rubro. | Módulo Catálogo (RF03) |
| `resources` | Insumos elementales (materiales, mano de obra, maquinarias, subcontratos). | Módulo Catálogo (RF04) |
| `items` | Partidas maestras del catálogo base de la empresa con su unidad de medida. | Módulo Catálogo (RF01, RF02) |
| `item_apu_resources` | Composición del Análisis de Precios Unitarios (rendimiento de insumos por partida). | Módulo Catálogo (RF02) |
| `quotations` | Cabecera comercial de la cotización asociada a un proyecto y cliente. | Módulo Cotizaciones (RF05, RF09) |
| `quotation_versions` | Versión inmutable de una cotización con vigencia, gastos y márgenes. | Módulo Cotizaciones (RF05.1, RF06, RF08) |
| `quotation_items` | Partidas presupuestadas en una versión específica (congelamiento de costo y cubicación). | Módulo Cotizaciones (RF05, RF07) |
| `quotation_item_resources` | Snapshot inmutable de insumos y rendimientos calculados para esa cotización. | Módulo Cotizaciones (RF04, RF06) |
| `work_orders` | Orden de Trabajo formal creada a partir de una cotización aprobada. | Flujo Comercial (RF10, RF11) |
| `work_order_items` | Partidas activas a ejecutar dentro de la obra con metas de cubicación. | Módulo Ejecución (RF10, RF12) |
| `execution_logs` | Registro diario/periódico en terreno de avances, costos reales y horas hombre. | Módulo Terreno (RF12) |
| `execution_photos` | Evidencias fotográficas y notas asociadas a avances o incidentes. | Módulo Terreno (RF13) |
| `data_sharing_consents` | Registro explícito de autorización de la empresa para anonimizar datos. | Módulo Índice (RF18) |
| `benchmark_index_records` | Estadísticas agregadas, depuradas y anonimizadas de costos y rendimientos. | Módulo Índice (RF17, RF19) |
| `contractor_directory_profiles`| Perfil público voluntario de prestadores por especialidad y zona. | Directorio / Red (RF23) |

---

## 4. Estado del esquema

| Entidad / módulo | Estado | Notas |
| --- | --- | --- |
| `1. Tenancy y Usuarios` | Diseñado | Soporta multi-empresa y asignación granular de roles. |
| `2. Clientes y Proyectos` | Diseñado | Vincula clientes con proyectos/obras. |
| `3. Catálogo y APU` | Diseñado | Permite rubros fijos y campo JSONB para rubro “Otro”. |
| `4. Cotizaciones y Versiones` | Diseñado | Cumple inmutabilidad histórica mediante snapshots de precios. |
| `5. Órdenes de Trabajo y Terreno` | Diseñado | Optimizado para registro rápido en terreno ($\le$ 30 s) y fotos. |
| `6. Índice Anonimizado y Consentimiento` | Diseñado | Respeta opt-in legal y aísla métricas sin identificar empresas. |

---

## 5. Diagrama DBML

```dbml
// ==========================================
// 1. TENANCY, USUARIOS Y CONTROL DE ACCESO
// ==========================================

Table companies {
  id uuid [pk, default: `gen_random_uuid()`]
  name varchar(255) [not null]
  rut_tax_id varchar(50) [not null, unique]
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
  deleted_at timestamptz [null]
}

Table users {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null]
  email varchar(255) [not null, unique]
  password_hash varchar(255) [not null]
  full_name varchar(255) [not null]
  phone varchar(50) [null]
  status varchar(20) [not null, default: 'ACTIVO'] // ACTIVO, INACTIVO, BLOQUEADO
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
  deleted_at timestamptz [null]
}

Table roles {
  id int [pk, increment]
  name varchar(50) [not null, unique] // DUENO, ADMIN_OBRA, JEFE_TERRENO, MAESTRO, CONSULTA, ADMIN_SISTEMA
  description varchar(255) [null]
}

Table user_roles {
  user_id uuid [not null]
  role_id int [not null]
  assigned_at timestamptz [not null, default: `now()`]

  indexes {
    (user_id, role_id) [pk]
  }
}

// ==========================================
// 2. CLIENTES Y PROYECTOS
// ==========================================

Table clients {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null]
  name varchar(255) [not null]
  rut_tax_id varchar(50) [null]
  contact_name varchar(255) [null]
  contact_email varchar(255) [null]
  contact_phone varchar(50) [null]
  address text [null]
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
  deleted_at timestamptz [null]
}

Table projects {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null]
  client_id uuid [not null]
  name varchar(255) [not null]
  description text [null]
  address text [null]
  status varchar(30) [not null, default: 'PLANIFICACION'] // PLANIFICACION, COTIZADO, EN_EJECUCION, FINALIZADO, CANCELADO
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
  deleted_at timestamptz [null]
}

// ==========================================
// 3. RUBROS, CATÁLOGO Y COMPOSICIÓN APU
// ==========================================

Table trades {
  id int [pk, increment]
  name varchar(100) [not null, unique] // Construcción, Electricidad, Vidriería, Otro
  is_custom boolean [not null, default: false]
}

Table trade_custom_configs {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null]
  trade_id int [not null]
  name varchar(150) [not null]
  config_json jsonb [not null] // Variables, formulas, conceptos y parametros
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
}

Table categories {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null]
  trade_id int [not null]
  name varchar(150) [not null]
  description text [null]
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
  deleted_at timestamptz [null]
}

Table resources {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null]
  name varchar(255) [not null]
  resource_type varchar(30) [not null] // MATERIAL, MANO_DE_OBRA, EQUIPO, SUBCONTRATO
  unit varchar(30) [not null] // m2, m3, kg, ltr, hh, dia, un
  current_unit_price numeric(14,2) [not null, default: 0]
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
  deleted_at timestamptz [null]
}

Table items {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null]
  category_id uuid [not null]
  code varchar(50) [null]
  name varchar(255) [not null]
  description text [null]
  unit varchar(30) [not null]
  is_active boolean [not null, default: true]
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
  deleted_at timestamptz [null]
}

Table item_apu_resources {
  id uuid [pk, default: `gen_random_uuid()`]
  item_id uuid [not null]
  resource_id uuid [not null]
  performance_quantity numeric(12,4) [not null] // Cantidad o rendimiento necesario por unidad de partida
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]

  indexes {
    (item_id, resource_id) [unique]
  }
}

// ==========================================
// 4. COTIZACIONES Y CONTROL DE VERSIONES
// ==========================================

Table quotations {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null]
  project_id uuid [not null]
  code varchar(50) [not null]
  current_version int [not null, default: 1]
  status varchar(30) [not null, default: 'BORRADOR'] // BORRADOR, EN_REVISION, ACEPTADA, RECHAZADA
  created_by uuid [not null]
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
  deleted_at timestamptz [null]

  indexes {
    (company_id, code) [unique]
  }
}

Table quotation_versions {
  id uuid [pk, default: `gen_random_uuid()`]
  quotation_id uuid [not null]
  version_number int [not null]
  valid_until date [not null]
  general_expenses_pct numeric(5,2) [not null, default: 0] // Gastos generales %
  profit_margin_pct numeric(5,2) [not null, default: 0]    // Margen comercial %
  direct_cost_subtotal numeric(14,2) [not null, default: 0]
  final_price numeric(14,2) [not null, default: 0]
  notes text [null]
  status varchar(30) [not null, default: 'BORRADOR'] // BORRADOR, ENVIADA, VENCIDA, ACEPTADA, RECHAZADA, SUPERADA
  created_by uuid [not null]
  created_at timestamptz [not null, default: `now()`]

  indexes {
    (quotation_id, version_number) [unique]
  }
}

Table quotation_items {
  id uuid [pk, default: `gen_random_uuid()`]
  quotation_version_id uuid [not null]
  item_id uuid [null] // Nullable si fue creada ad-hoc
  item_name_snapshot varchar(255) [not null]
  unit varchar(30) [not null]
  quantity numeric(12,4) [not null]
  unit_direct_cost_snapshot numeric(14,2) [not null] // Costo unitario calculado en ese momento
  subtotal_direct_cost numeric(14,2) [not null]      // quantity * unit_direct_cost_snapshot
  is_extraordinary boolean [not null, default: false] // RF07: Partidas extraordinarias
  created_at timestamptz [not null, default: `now()`]
}

Table quotation_item_resources {
  id uuid [pk, default: `gen_random_uuid()`]
  quotation_item_id uuid [not null]
  resource_id uuid [null]
  resource_name_snapshot varchar(255) [not null]
  resource_type varchar(30) [not null]
  unit varchar(30) [not null]
  unit_price_snapshot numeric(14,2) [not null]
  performance_quantity numeric(12,4) [not null]
  subtotal numeric(14,2) [not null] // performance_quantity * unit_price_snapshot
}

// ==========================================
// 5. ÓRDENES DE TRABAJO Y REGISTRO EN TERRENO
// ==========================================

Table work_orders {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null]
  project_id uuid [not null]
  quotation_version_id uuid [not null]
  code varchar(50) [not null]
  status varchar(30) [not null, default: 'EN_EJECUCION'] // EN_EJECUCION, PAUSADA, CERRADA
  start_date date [not null]
  closed_at timestamptz [null]
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]

  indexes {
    (company_id, code) [unique]
  }
}

Table work_order_items {
  id uuid [pk, default: `gen_random_uuid()`]
  work_order_id uuid [not null]
  quotation_item_id uuid [not null]
  item_name_snapshot varchar(255) [not null]
  unit varchar(30) [not null]
  budgeted_quantity numeric(12,4) [not null]
  budgeted_unit_cost numeric(14,2) [not null]
  accumulated_quantity numeric(12,4) [not null, default: 0]
  accumulated_real_cost numeric(14,2) [not null, default: 0]
  status varchar(30) [not null, default: 'PENDIENTE'] // PENDIENTE, EN_PROGRESO, COMPLETADA
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
}

Table execution_logs {
  id uuid [pk, default: `gen_random_uuid()`]
  work_order_item_id uuid [not null]
  user_id uuid [not null]
  log_date date [not null]
  executed_quantity numeric(12,4) [not null, default: 0]
  worked_hours numeric(6,2) [not null, default: 0]
  actual_cost numeric(14,2) [not null, default: 0]
  notes text [null]
  created_at timestamptz [not null, default: `now()`]
}

Table execution_photos {
  id uuid [pk, default: `gen_random_uuid()`]
  execution_log_id uuid [not null]
  photo_url varchar(1024) [not null]
  caption text [null]
  is_incident boolean [not null, default: false]
  created_at timestamptz [not null, default: `now()`]
}

// ==========================================
// 6. CONSENTIMIENTO, ÍNDICE Y DIRECTORIO
// ==========================================

Table data_sharing_consents {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null, unique]
  consent_granted boolean [not null, default: false]
  granted_at timestamptz [null]
  revoked_at timestamptz [null]
  updated_at timestamptz [not null, default: `now()`]
}

Table benchmark_index_records {
  id uuid [pk, default: `gen_random_uuid()`]
  trade_id int [not null]
  standard_item_name varchar(255) [not null]
  unit varchar(30) [not null]
  sample_size int [not null, default: 0]
  avg_unit_cost numeric(14,2) [not null]
  median_unit_cost numeric(14,2) [not null]
  min_unit_cost numeric(14,2) [not null]
  max_unit_cost numeric(14,2) [not null]
  avg_deviation_pct numeric(5,2) [not null]
  last_calculated_at timestamptz [not null, default: `now()`]

  indexes {
    (trade_id, standard_item_name, unit) [unique]
  }
}

Table contractor_directory_profiles {
  id uuid [pk, default: `gen_random_uuid()`]
  company_id uuid [not null, unique]
  trade_id int [not null]
  public_business_name varchar(255) [not null]
  contact_phone varchar(50) [null]
  contact_email varchar(255) [null]
  coverage_zone varchar(255) [not null]
  description text [null]
  is_published boolean [not null, default: false]
  created_at timestamptz [not null, default: `now()`]
  updated_at timestamptz [not null, default: `now()`]
}

// ==========================================
// DEFINICIÓN DE REFERENCIAS / FOREIGN KEYS
// ==========================================

Ref: users.company_id > companies.id [delete: cascade]
Ref: user_roles.user_id > users.id [delete: cascade]
Ref: user_roles.role_id > roles.id [delete: restrict]

Ref: clients.company_id > companies.id [delete: cascade]
Ref: projects.company_id > companies.id [delete: cascade]
Ref: projects.client_id > clients.id [delete: restrict]

Ref: trade_custom_configs.company_id > companies.id [delete: cascade]
Ref: trade_custom_configs.trade_id > trades.id [delete: restrict]

Ref: categories.company_id > companies.id [delete: cascade]
Ref: categories.trade_id > trades.id [delete: restrict]

Ref: resources.company_id > companies.id [delete: cascade]
Ref: items.company_id > companies.id [delete: cascade]
Ref: items.category_id > categories.id [delete: restrict]

Ref: item_apu_resources.item_id > items.id [delete: cascade]
Ref: item_apu_resources.resource_id > resources.id [delete: restrict]

Ref: quotations.company_id > companies.id [delete: cascade]
Ref: quotations.project_id > projects.id [delete: restrict]
Ref: quotations.created_by > users.id [delete: restrict]

Ref: quotation_versions.quotation_id > quotations.id [delete: cascade]
Ref: quotation_versions.created_by > users.id [delete: restrict]

Ref: quotation_items.quotation_version_id > quotation_versions.id [delete: cascade]
Ref: quotation_items.item_id > items.id [delete: set null]

Ref: quotation_item_resources.quotation_item_id > quotation_items.id [delete: cascade]
Ref: quotation_item_resources.resource_id > resources.id [delete: set null]

Ref: work_orders.company_id > companies.id [delete: cascade]
Ref: work_orders.project_id > projects.id [delete: restrict]
Ref: work_orders.quotation_version_id > quotation_versions.id [delete: restrict]

Ref: work_order_items.work_order_id > work_orders.id [delete: cascade]
Ref: work_order_items.quotation_item_id > quotation_items.id [delete: restrict]

Ref: execution_logs.work_order_item_id > work_order_items.id [delete: cascade]
Ref: execution_logs.user_id > users.id [delete: restrict]

Ref: execution_photos.execution_log_id > execution_logs.id [delete: cascade]

Ref: data_sharing_consents.company_id > companies.id [delete: cascade]
Ref: benchmark_index_records.trade_id > trades.id [delete: restrict]
Ref: contractor_directory_profiles.company_id > companies.id [delete: cascade]
Ref: contractor_directory_profiles.trade_id > trades.id [delete: restrict]
```

---

## 6. Relaciones explícitas

| Origen | Destino | Cardinalidad | Regla de borrado | Motivo |
| --- | --- | --- | --- | --- |
| `users.company_id` | `companies.id` | N:1 | `CASCADE` | Un usuario pertenece obligatoriamente a una empresa tenant. |
| `user_roles.user_id` | `users.id` | N:1 | `CASCADE` | Limpieza de roles asignados al eliminar usuario. |
| `user_roles.role_id` | `roles.id` | N:1 | `RESTRICT` | Los roles del sistema no deben borrarse si tienen usuarios. |
| `projects.client_id` | `clients.id` | N:1 | `RESTRICT` | No se puede borrar un cliente con proyectos históricos asociados. |
| `items.category_id` | `categories.id` | N:1 | `RESTRICT` | Evita partidas huérfanas sin categoría asignada. |
| `item_apu_resources.item_id` | `items.id` | N:1 | `CASCADE` | El desglose APU depende de la partida del catálogo. |
| `item_apu_resources.resource_id` | `resources.id` | N:1 | `RESTRICT` | No se debe eliminar un insumo si está en uso en el catálogo. |
| `quotation_versions.quotation_id`| `quotations.id` | N:1 | `CASCADE` | Las versiones pertenecen al ciclo de la cotización madre. |
| `quotation_items.quotation_version_id` | `quotation_versions.id` | N:1 | `CASCADE` | Las partidas presupuestadas son parte indivisible de la versión. |
| `quotation_items.item_id` | `items.id` | N:1 | `SET NULL` | Snapshot: Si se borra la partida del catálogo, la cotización histórica retiene su snapshot. |
| `work_orders.quotation_version_id`| `quotation_versions.id` | N:1 | `RESTRICT` | La versión aceptada que originó la OT no puede eliminarse. |
| `work_order_items.work_order_id` | `work_orders.id` | N:1 | `CASCADE` | Las partidas activas dependen de la Orden de Trabajo. |
| `execution_logs.work_order_item_id` | `work_order_items.id`| N:1 | `CASCADE` | Registros diarios de avance de la partida de la OT. |
| `execution_photos.execution_log_id`| `execution_logs.id` | N:1 | `CASCADE` | Las fotos documentan un registro o incidencia específica. |
| `data_sharing_consents.company_id` | `companies.id` | 1:1 | `CASCADE` | Cada empresa tiene un único registro de consentimiento legal. |

---

## 7. Índices y restricciones

| Tabla | Índice / Constraint | Columnas | Tipo | Motivo |
| --- | --- | --- | --- | --- |
| `companies` | `uq_companies_rut` | `rut_tax_id` | `UNIQUE` | Evitar empresas duplicadas por RUT. |
| `users` | `uq_users_email` | `email` | `UNIQUE` | Credencial única de acceso. |
| `users` | `idx_users_company_id` | `company_id` | `BTREE` | Filtrado rápido de usuarios por empresa tenant. |
| `clients` | `idx_clients_company` | `company_id` | `BTREE` | Búsqueda y listado de clientes por empresa. |
| `projects` | `idx_projects_company_status` | `company_id, status` | `BTREE` | Filtrado del dashboard de proyectos activos. |
| `items` | `idx_items_company_category` | `company_id, category_id` | `BTREE` | Carga ágil del catálogo de partidas. |
| `item_apu_resources` | `uq_item_resource` | `item_id, resource_id` | `UNIQUE` | Evitar duplicidad de un insumo dentro de la misma partida. |
| `quotations` | `uq_quotations_company_code` | `company_id, code` | `UNIQUE` | Código de cotización correlativo único por empresa. |
| `quotation_versions` | `uq_quotation_version_num` | `quotation_id, version_number` | `UNIQUE` | Integridad de número de versión correlativo. |
| `quotation_versions` | `idx_quotation_versions_valid` | `valid_until, status` | `BTREE` | Verificación de cotizaciones vencidas (RF08). |
| `work_orders` | `uq_work_orders_company_code` | `company_id, code` | `UNIQUE` | Código de OT único por empresa. |
| `work_orders` | `idx_work_orders_status` | `company_id, status` | `BTREE` | Consultas de obras en ejecución en terreno. |
| `execution_logs` | `idx_execution_logs_item_date` | `work_order_item_id, log_date` | `BTREE` | Cálculo rápido de acumulados y reportes cronológicos. |
| `benchmark_index_records` | `uq_benchmark_item_unit` | `trade_id, standard_item_name, unit` | `UNIQUE` | Unicidad del registro de referencia del índice. |
| `trade_custom_configs` | `idx_trade_configs_company` | `company_id, trade_id` | `BTREE` | Recuperación de configuración dinámica del rubro "Otro". |

---

## 8. Diferencias entre modelo objetivo y esquema físico

| Elemento | Modelo objetivo | Esquema físico actual | Acción |
| --- | --- | --- | --- |
| Tablas completas (23 tablas) | Modelo relacional documentado en este archivo | Base de datos no inicializada | Ejecutar migración inicial SQL mediante `golang-migrate` al iniciar desarrollo backend. |
| Insumos y snapshots | Desacoplamiento entre catálogo e histórico | No implementado | Mantener columnas `_snapshot` para evitar cambios retroactivos al alterar precios de catálogo. |

---

## 9. Reglas de integridad

1. **Inmutabilidad de versiones aprobadas:** Cuando una `quotation_version` pasa a estado `ACEPTADA` o se convierte en `work_orders`, sus registros hijos (`quotation_items`, `quotation_item_resources`) quedan en modo de solo lectura; cualquier cambio posterior exige generar una nueva versión o una partida extraordinaria (`is_extraordinary = true`).
2. **Snapshot de costos de insumos:** Las tablas de cotización guardan el valor monetario exacto (`unit_price_snapshot`) y nombre (`item_name_snapshot`) en el instante de cotizar. Las variaciones futuras en la tabla `resources` no alteran cotizaciones ya emitidas.
3. **Cálculo de desvíos en tiempo real:** Los valores de desviación no se guardan como columnas estáticas en la OT; se calculan mediante consultas dinámicas o vistas SQL comparando `budgeted_quantity * budgeted_unit_cost` frente a `accumulated_quantity * accumulated_real_cost`.
4. **Consentimiento explícito (Opt-in):** Solo se procesan datos para `benchmark_index_records` si `data_sharing_consents.consent_granted = true` y la OT se encuentra en estado `CERRADA`.
5. **Anonimización absoluta del índice:** Las tablas estadísticas del índice (`benchmark_index_records`) jamás almacenan `company_id`, `client_id`, `project_id` ni márgenes comerciales (`profit_margin_pct`), solo promedios, medianas y desviaciones agregadas.

---

## 10. Datos semilla

| Semilla | Tabla | Ambiente | Motivo |
| --- | --- | --- | --- |
| Roles base (`DUENO`, `ADMIN_OBRA`, `JEFE_TERRENO`, `MAESTRO`, `CONSULTA`, `ADMIN_SISTEMA`) | `roles` | `dev/test/prod` | Control de autorización por roles (RF20). |
| Rubros iniciales (`Construcción`, `Electricidad`, `Vidriería`, `Otro`) | `trades` | `dev/test/prod` | Rubros estándar del MVP y soporte configurable (RF22). |
| Catálogo base de prueba (partidas y materiales frecuentes) | `categories`, `resources`, `items`, `item_apu_resources` | `dev/test` | Facilitar pruebas funcionales de cotización y cubicación. |

---

## 11. Migraciones

La gestión de migraciones se realizará con **`golang-migrate`**, almacenando los scripts en el directorio `backend/migrations/`:

### Estructura de nombres:
```text
backend/migrations/
  ├── 000001_create_tenancy_and_auth.up.sql
  ├── 000001_create_tenancy_and_auth.down.sql
  ├── 000002_create_catalog_and_apu.up.sql
  ├── 000002_create_catalog_and_apu.down.sql
  ├── 000003_create_quotations_and_versions.up.sql
  ├── 000003_create_quotations_and_versions.down.sql
  ├── 000004_create_work_orders_and_execution.up.sql
  ├── 000004_create_work_orders_and_execution.down.sql
  ├── 000005_create_index_and_consents.up.sql
  └── 000005_create_index_and_consents.down.sql
```

### Reglas para nuevas migraciones:
1. Siempre crear tanto el archivo `.up.sql` como su reverso `.down.sql`.
2. Las claves foráneas deben definirse con `ON DELETE RESTRICT` o `CASCADE` de forma explícita.
3. No alterar columnas en caliente sin antes poblar datos default si se agregan restricciones `NOT NULL`.

---

## 12. Checklist de base de datos

- [x] Todas las entidades del PRD (RF01 a RF23) tienen representación clara y justificada.
- [x] Las claves foráneas cuentan con reglas físicas de integridad referencial declaradas.
- [x] Se diseñaron índices compuestos para consultas críticas (búsquedas por empresa, catálogo y fechas).
- [x] Los datos históricos y versiones cerradas son inmutables mediante snapshots.
- [x] Se soporta la configuración flexible del rubro "Otro" mediante `JSONB` sin modificar tablas.
- [x] Se garantiza la anonimización y protección de márgenes en el módulo del índice.
- [x] La estrategia de migraciones está alineada al stack Go y PostgreSQL.
