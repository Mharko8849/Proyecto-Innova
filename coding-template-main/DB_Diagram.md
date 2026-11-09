# DB Diagram — Modelo de datos: ÍndicePartidas

> Este documento define el modelo lógico y físico de base de datos para la plataforma **ÍndicePartidas**, de acuerdo con los requerimientos funcionales y no funcionales definidos en el PRD y con la arquitectura tecnológica del proyecto.

---

## 1. Propósito del modelo

Este documento describe la estructura de datos relacional necesaria para soportar el ciclo completo de ÍndicePartidas:

**Empresas y usuarios → Clientes y proyectos → Rubros → Catálogo y APU → Cotizaciones versionadas → Órdenes de Trabajo → Registro de ejecución real → Comparación Cotizado vs. Ejecutado → Normalización de partidas → Validación de aportes → ÍndicePartidas agregado y anonimizado.**

|Campo|Definición|
|---|---|
|Motor de base de datos|**PostgreSQL**|
|Backend|**Go**|
|Frontend|**React + JavaScript**|
|Comunicación|**API REST mediante JSON**|
|ORM / Query layer|**pgx / sqlx / GORM**|
|Estrategia de migraciones|**golang-migrate**|
|Estado del documento|**Modelo objetivo actualizado**|

---

## 2. Convenciones

- **Nombres de tablas:** `snake_case`, plural y en español.
- **Nombres de columnas:** `snake_case`, en español y sin tildes.
- **Claves primarias:** `id UUID` generado mediante `gen_random_uuid()` para entidades de negocio.
- **Roles:** se mantiene `INT` incremental por tratarse de un catálogo fijo del sistema.
- **Claves foráneas:** formato `[entidad]_id`.
- **Fechas y auditoría:** `TIMESTAMPTZ` para eventos del sistema y `DATE` para fechas de negocio.
- **Soft delete:** mediante `eliminado_en TIMESTAMPTZ NULL` o `esta_activo BOOLEAN`.
- **Valores monetarios:** `NUMERIC(14,2)`.
- **Cantidades y rendimientos:** `NUMERIC(12,4)`.
- **Porcentajes:** `NUMERIC(5,2)` o `NUMERIC(7,2)` según rango requerido.
- **Configuraciones extensibles:** `JSONB` para almacenar variables, fórmulas, conceptos y parámetros de rubros personalizados.
- **Snapshots históricos:** las cotizaciones guardan nombres, costos y precios utilizados en ese momento para impedir que cambios posteriores modifiquen el histórico.

---

## 3. Glosario de entidades

|Entidad|Descripción|Módulo / propósito|
|---|---|---|
|`empresas`|Empresa, contratista o pyme que utiliza la plataforma.|Usuarios / Configuración|
|`usuarios`|Usuarios asociados a una empresa.|Usuarios / Configuración|
|`roles`|Catálogo de roles del sistema.|Control de acceso|
|`usuarios_roles`|Relación entre usuarios y roles.|Control de acceso|
|`clientes`|Clientes finales o mandantes de la empresa.|Clientes|
|`proyectos`|Trabajos u obras asociados a clientes.|Proyectos|
|`rubros`|Rubros predefinidos o personalizados.|Configuración|
|`categorias`|Agrupación de partidas dentro de un rubro.|Catálogo|
|`recursos`|Materiales, mano de obra, equipos, subcontratos u otros recursos.|Catálogo / APU|
|`historial_precios_recursos`|Histórico de precios por recurso.|Catálogo / Precios|
|`partidas`|Partidas maestras utilizadas para cotizar.|Catálogo|
|`recursos_apu_partida`|Composición APU de cada partida.|APU|
|`cotizaciones`|Cabecera comercial de una cotización.|Cotizaciones|
|`versiones_cotizacion`|Versiones inmutables de una cotización.|Cotizaciones|
|`partidas_cotizacion`|Partidas incluidas en cada versión.|Cotizaciones|
|`recursos_partida_cotizacion`|Snapshot de recursos y costos utilizados al cotizar.|Cotizaciones|
|`ordenes_trabajo`|Orden creada desde una cotización aceptada.|Órdenes de Trabajo|
|`partidas_orden_trabajo`|Partidas activas a ejecutar y controlar.|Ejecución|
|`registros_ejecucion`|Registro periódico de cantidad ejecutada y horas.|Terreno|
|`recursos_ejecucion`|Recursos realmente utilizados durante la ejecución.|Terreno / Costos reales|
|`costos_adicionales_ejecucion`|Imprevistos y costos adicionales.|Terreno / Costos reales|
|`fotos_ejecucion`|Evidencias fotográficas e incidencias.|Terreno|
|`consentimientos_datos`|Estado actual del consentimiento de una empresa.|Índice|
|`historial_consentimientos`|Historial de otorgamiento y revocación del consentimiento.|Índice|
|`partidas_estandar_indice`|Catálogo estándar para agrupar partidas equivalentes.|Índice|
|`equivalencias_partidas_indice`|Relación entre partidas privadas y partidas estándar.|Índice|
|`aportes_indice_partidas`|Datos validados que pueden alimentar el índice.|Índice|
|`registros_indice_partidas`|Estadísticas agregadas y anonimizadas.|Índice|
|`perfiles_directorio_contratistas`|Perfiles públicos voluntarios por rubro.|Directorio|

---

## 4. Estado del esquema

|Módulo|Estado|Observación|
|---|---|---|
|Empresas, usuarios y roles|Diseñado|Soporta múltiples empresas y control de acceso.|
|Clientes y proyectos|Diseñado|Relaciona proyectos con sus clientes.|
|Rubros personalizados|Diseñado|Permite Construcción, Electricidad, Vidriería y rubros creados desde “Otro”.|
|Catálogo y APU|Diseñado|Soporta partidas, recursos, rendimientos e historial de precios.|
|Cotizaciones y versiones|Diseñado|Mantiene trazabilidad comercial e histórico por snapshots.|
|Órdenes de Trabajo|Diseñado|Se generan desde una versión aceptada.|
|Registro de ejecución|Diseñado|Registra cantidades, horas, recursos reales, adicionales y fotografías.|
|Cotizado vs. Ejecutado|Diseñado|Puede calcularse desde partidas de OT y registros reales.|
|Consentimiento|Diseñado|Incluye estado actual e historial de cambios.|
|Normalización del índice|Diseñado|Permite agrupar nombres diferentes bajo una partida estándar.|
|Aportes y validación|Diseñado|Separa datos privados de los registros agregados del índice.|
|ÍndicePartidas agregado|Diseñado|Considera muestra, costos, rendimiento, desviación, región y período.|
|Directorio voluntario|Diseñado|Publicación opcional por rubro y zona.|

---

## 5. Diagrama DBML

```
// ==========================================================
// INDICEPARTIDAS
// MODELO DE BASE DE DATOS
// PostgreSQL
// Backend: Go
// Frontend: React + JavaScript
// ==========================================================


// ==========================================================
// 1. EMPRESAS, USUARIOS Y CONTROL DE ACCESO
// ==========================================================

Table empresas {
  id uuid [pk, default: `gen_random_uuid()`]
  nombre varchar(255) [not null]
  rut varchar(50) [not null, unique]
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]
  eliminado_en timestamptz [null]
}

Table usuarios {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  correo varchar(255) [not null, unique]
  contrasena_hash varchar(255) [not null]
  nombre_completo varchar(255) [not null]
  telefono varchar(50) [null]
  estado varchar(20) [not null, default: 'ACTIVO']
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]
  eliminado_en timestamptz [null]
}

Table roles {
  id int [pk, increment]
  nombre varchar(50) [not null, unique]
  descripcion varchar(255) [null]
}

Table usuarios_roles {
  usuario_id uuid [not null]
  rol_id int [not null]
  asignado_en timestamptz [not null, default: `now()`]

  indexes {
    (usuario_id, rol_id) [pk]
  }
}


// ==========================================================
// 2. CLIENTES Y PROYECTOS
// ==========================================================

Table clientes {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  nombre varchar(255) [not null]
  rut varchar(50) [null]
  nombre_contacto varchar(255) [null]
  correo_contacto varchar(255) [null]
  telefono_contacto varchar(50) [null]
  direccion text [null]
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]
  eliminado_en timestamptz [null]
}

Table proyectos {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  cliente_id uuid [not null]
  nombre varchar(255) [not null]
  descripcion text [null]
  direccion text [null]
  estado varchar(30) [not null, default: 'PLANIFICACION']
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]
  eliminado_en timestamptz [null]
}


// ==========================================================
// 3. RUBROS
// ==========================================================

Table rubros {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [null]
  nombre varchar(150) [not null]
  descripcion text [null]
  es_personalizado boolean [not null, default: false]
  configuracion_json jsonb [null]
  esta_activo boolean [not null, default: true]
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]

  indexes {
    (empresa_id, nombre) [unique]
  }
}


// ==========================================================
// 4. CATALOGO, RECURSOS Y APU
// ==========================================================

Table categorias {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  rubro_id uuid [not null]
  nombre varchar(150) [not null]
  descripcion text [null]
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]
  eliminado_en timestamptz [null]
}

Table recursos {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  nombre varchar(255) [not null]
  tipo_recurso varchar(30) [not null]
  unidad varchar(30) [not null]
  precio_unitario_actual numeric(14,2) [not null, default: 0]
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]
  eliminado_en timestamptz [null]
}

Table historial_precios_recursos {
  id uuid [pk, default: `gen_random_uuid()`]
  recurso_id uuid [not null]
  precio_unitario numeric(14,2) [not null]
  fecha_desde date [not null]
  fecha_hasta date [null]
  fuente varchar(255) [null]
  creado_en timestamptz [not null, default: `now()`]
}

Table partidas {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  categoria_id uuid [not null]
  codigo varchar(50) [null]
  nombre varchar(255) [not null]
  descripcion text [null]
  unidad varchar(30) [not null]
  esta_activa boolean [not null, default: true]
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]
  eliminado_en timestamptz [null]
}

Table recursos_apu_partida {
  id uuid [pk, default: `gen_random_uuid()`]
  partida_id uuid [not null]
  recurso_id uuid [not null]
  cantidad_rendimiento numeric(12,4) [not null]
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]

  indexes {
    (partida_id, recurso_id) [unique]
  }
}


// ==========================================================
// 5. COTIZACIONES
// ==========================================================

Table cotizaciones {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  proyecto_id uuid [not null]
  codigo varchar(50) [not null]
  version_actual int [not null, default: 1]
  estado varchar(30) [not null, default: 'BORRADOR']
  creado_por uuid [not null]
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]
  eliminado_en timestamptz [null]

  indexes {
    (empresa_id, codigo) [unique]
  }
}

Table versiones_cotizacion {
  id uuid [pk, default: `gen_random_uuid()`]
  cotizacion_id uuid [not null]
  numero_version int [not null]
  vigente_hasta date [not null]
  porcentaje_gastos_generales numeric(5,2) [not null, default: 0]
  porcentaje_margen_ganancia numeric(5,2) [not null, default: 0]
  subtotal_costo_directo numeric(14,2) [not null, default: 0]
  precio_final numeric(14,2) [not null, default: 0]
  notas text [null]
  estado varchar(30) [not null, default: 'BORRADOR']
  enviada_en timestamptz [null]
  aceptada_en timestamptz [null]
  rechazada_en timestamptz [null]
  motivo_rechazo text [null]
  medio_envio varchar(30) [null]
  creado_por uuid [not null]
  creado_en timestamptz [not null, default: `now()`]

  indexes {
    (cotizacion_id, numero_version) [unique]
  }
}

Table partidas_cotizacion {
  id uuid [pk, default: `gen_random_uuid()`]
  version_cotizacion_id uuid [not null]
  partida_id uuid [null]
  nombre_partida_respaldo varchar(255) [not null]
  unidad varchar(30) [not null]
  cantidad numeric(12,4) [not null]
  costo_unitario_directo_respaldo numeric(14,2) [not null]
  subtotal_costo_directo numeric(14,2) [not null]
  es_extraordinaria boolean [not null, default: false]
  creado_en timestamptz [not null, default: `now()`]
}

Table recursos_partida_cotizacion {
  id uuid [pk, default: `gen_random_uuid()`]
  partida_cotizacion_id uuid [not null]
  recurso_id uuid [null]
  nombre_recurso_respaldo varchar(255) [not null]
  tipo_recurso varchar(30) [not null]
  unidad varchar(30) [not null]
  precio_unitario_respaldo numeric(14,2) [not null]
  cantidad_rendimiento numeric(12,4) [not null]
  subtotal numeric(14,2) [not null]
}


// ==========================================================
// 6. ORDENES DE TRABAJO
// ==========================================================

Table ordenes_trabajo {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  proyecto_id uuid [not null]
  version_cotizacion_id uuid [not null]
  codigo varchar(50) [not null]
  estado varchar(30) [not null, default: 'EN_EJECUCION']
  fecha_inicio date [not null]
  cerrado_en timestamptz [null]
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]

  indexes {
    (empresa_id, codigo) [unique]
  }
}

Table partidas_orden_trabajo {
  id uuid [pk, default: `gen_random_uuid()`]
  orden_trabajo_id uuid [not null]
  partida_cotizacion_id uuid [not null]
  nombre_partida_respaldo varchar(255) [not null]
  unidad varchar(30) [not null]
  cantidad_presupuestada numeric(12,4) [not null]
  costo_unitario_presupuestado numeric(14,2) [not null]
  cantidad_ejecutada_acumulada numeric(12,4) [not null, default: 0]
  costo_real_acumulado numeric(14,2) [not null, default: 0]
  estado varchar(30) [not null, default: 'PENDIENTE']
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]
}


// ==========================================================
// 7. REGISTRO DE EJECUCION
// ==========================================================

Table registros_ejecucion {
  id uuid [pk, default: `gen_random_uuid()`]
  partida_orden_trabajo_id uuid [not null]
  usuario_id uuid [not null]
  fecha_registro date [not null]
  cantidad_ejecutada numeric(12,4) [not null, default: 0]
  horas_trabajadas numeric(6,2) [not null, default: 0]
  notas text [null]
  creado_en timestamptz [not null, default: `now()`]
}

Table recursos_ejecucion {
  id uuid [pk, default: `gen_random_uuid()`]
  registro_ejecucion_id uuid [not null]
  recurso_id uuid [null]
  nombre_recurso_respaldo varchar(255) [not null]
  tipo_recurso varchar(30) [not null]
  unidad varchar(30) [not null]
  cantidad_real numeric(12,4) [not null]
  precio_unitario_real numeric(14,2) [not null]
  subtotal_real numeric(14,2) [not null]
  creado_en timestamptz [not null, default: `now()`]
}

Table costos_adicionales_ejecucion {
  id uuid [pk, default: `gen_random_uuid()`]
  registro_ejecucion_id uuid [not null]
  concepto varchar(255) [not null]
  monto numeric(14,2) [not null]
  descripcion text [null]
  creado_en timestamptz [not null, default: `now()`]
}

Table fotos_ejecucion {
  id uuid [pk, default: `gen_random_uuid()`]
  registro_ejecucion_id uuid [not null]
  url_foto varchar(1024) [not null]
  descripcion text [null]
  es_incidencia boolean [not null, default: false]
  creado_en timestamptz [not null, default: `now()`]
}


// ==========================================================
// 8. CONSENTIMIENTO
// ==========================================================

Table consentimientos_datos {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null, unique]
  consentimiento_otorgado boolean [not null, default: false]
  otorgado_en timestamptz [null]
  revocado_en timestamptz [null]
  actualizado_en timestamptz [not null, default: `now()`]
}

Table historial_consentimientos {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  accion varchar(20) [not null]
  realizado_en timestamptz [not null, default: `now()`]
}


// ==========================================================
// 9. NORMALIZACION DEL INDICE
// ==========================================================

Table partidas_estandar_indice {
  id uuid [pk, default: `gen_random_uuid()`]
  rubro_id uuid [not null]
  nombre varchar(255) [not null]
  descripcion text [null]
  unidad varchar(30) [not null]
  esta_activa boolean [not null, default: true]
  creado_en timestamptz [not null, default: `now()`]
}

Table equivalencias_partidas_indice {
  id uuid [pk, default: `gen_random_uuid()`]
  partida_id uuid [not null]
  partida_estandar_id uuid [not null]
  validada boolean [not null, default: false]
  creado_en timestamptz [not null, default: `now()`]

  indexes {
    (partida_id, partida_estandar_id) [unique]
  }
}


// ==========================================================
// 10. APORTES AL INDICE
// ==========================================================

Table aportes_indice_partidas {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  partida_orden_trabajo_id uuid [not null]
  partida_estandar_id uuid [not null]
  unidad varchar(30) [not null]
  cantidad_ejecutada numeric(12,4) [not null]
  costo_real_total numeric(14,2) [not null]
  costo_unitario_real numeric(14,2) [not null]
  horas_trabajadas numeric(10,2) [not null, default: 0]
  rendimiento_real numeric(12,4) [null]
  desviacion_porcentaje numeric(7,2) [null]
  fecha_ejecucion date [not null]
  region varchar(100) [null]
  comuna varchar(100) [null]
  estado_validacion varchar(30) [not null, default: 'PENDIENTE']
  motivo_validacion text [null]
  revisado_por uuid [null]
  revisado_en timestamptz [null]
  creado_en timestamptz [not null, default: `now()`]
}


// ==========================================================
// 11. INDICEPARTIDAS AGREGADO
// ==========================================================

Table registros_indice_partidas {
  id uuid [pk, default: `gen_random_uuid()`]
  partida_estandar_id uuid [not null]
  region varchar(100) [null]
  comuna varchar(100) [null]
  periodo_desde date [null]
  periodo_hasta date [null]
  cantidad_muestras int [not null, default: 0]
  costo_unitario_promedio numeric(14,2) [not null]
  costo_unitario_mediana numeric(14,2) [not null]
  costo_unitario_minimo numeric(14,2) [not null]
  costo_unitario_maximo numeric(14,2) [not null]
  rendimiento_promedio numeric(12,4) [null]
  desviacion_promedio_porcentaje numeric(7,2) [null]
  calculado_en timestamptz [not null, default: `now()`]
}


// ==========================================================
// 12. DIRECTORIO
// ==========================================================

Table perfiles_directorio_contratistas {
  id uuid [pk, default: `gen_random_uuid()`]
  empresa_id uuid [not null]
  rubro_id uuid [not null]
  nombre_comercial_publico varchar(255) [not null]
  telefono_contacto varchar(50) [null]
  correo_contacto varchar(255) [null]
  zona_cobertura varchar(255) [not null]
  descripcion text [null]
  esta_publicado boolean [not null, default: false]
  creado_en timestamptz [not null, default: `now()`]
  actualizado_en timestamptz [not null, default: `now()`]

  indexes {
    (empresa_id, rubro_id) [unique]
  }
}


// ==========================================================
// 13. REFERENCIAS
// ==========================================================

Ref: usuarios.empresa_id > empresas.id [delete: cascade]
Ref: usuarios_roles.usuario_id > usuarios.id [delete: cascade]
Ref: usuarios_roles.rol_id > roles.id [delete: restrict]

Ref: clientes.empresa_id > empresas.id [delete: cascade]
Ref: proyectos.empresa_id > empresas.id [delete: cascade]
Ref: proyectos.cliente_id > clientes.id [delete: restrict]

Ref: rubros.empresa_id > empresas.id [delete: cascade]

Ref: categorias.empresa_id > empresas.id [delete: cascade]
Ref: categorias.rubro_id > rubros.id [delete: restrict]
Ref: recursos.empresa_id > empresas.id [delete: cascade]
Ref: historial_precios_recursos.recurso_id > recursos.id [delete: cascade]
Ref: partidas.empresa_id > empresas.id [delete: cascade]
Ref: partidas.categoria_id > categorias.id [delete: restrict]
Ref: recursos_apu_partida.partida_id > partidas.id [delete: cascade]
Ref: recursos_apu_partida.recurso_id > recursos.id [delete: restrict]

Ref: cotizaciones.empresa_id > empresas.id [delete: cascade]
Ref: cotizaciones.proyecto_id > proyectos.id [delete: restrict]
Ref: cotizaciones.creado_por > usuarios.id [delete: restrict]
Ref: versiones_cotizacion.cotizacion_id > cotizaciones.id [delete: cascade]
Ref: versiones_cotizacion.creado_por > usuarios.id [delete: restrict]
Ref: partidas_cotizacion.version_cotizacion_id > versiones_cotizacion.id [delete: cascade]
Ref: partidas_cotizacion.partida_id > partidas.id [delete: set null]
Ref: recursos_partida_cotizacion.partida_cotizacion_id > partidas_cotizacion.id [delete: cascade]
Ref: recursos_partida_cotizacion.recurso_id > recursos.id [delete: set null]

Ref: ordenes_trabajo.empresa_id > empresas.id [delete: cascade]
Ref: ordenes_trabajo.proyecto_id > proyectos.id [delete: restrict]
Ref: ordenes_trabajo.version_cotizacion_id > versiones_cotizacion.id [delete: restrict]
Ref: partidas_orden_trabajo.orden_trabajo_id > ordenes_trabajo.id [delete: cascade]
Ref: partidas_orden_trabajo.partida_cotizacion_id > partidas_cotizacion.id [delete: restrict]

Ref: registros_ejecucion.partida_orden_trabajo_id > partidas_orden_trabajo.id [delete: cascade]
Ref: registros_ejecucion.usuario_id > usuarios.id [delete: restrict]
Ref: recursos_ejecucion.registro_ejecucion_id > registros_ejecucion.id [delete: cascade]
Ref: recursos_ejecucion.recurso_id > recursos.id [delete: set null]
Ref: costos_adicionales_ejecucion.registro_ejecucion_id > registros_ejecucion.id [delete: cascade]
Ref: fotos_ejecucion.registro_ejecucion_id > registros_ejecucion.id [delete: cascade]

Ref: consentimientos_datos.empresa_id > empresas.id [delete: cascade]
Ref: historial_consentimientos.empresa_id > empresas.id [delete: cascade]

Ref: partidas_estandar_indice.rubro_id > rubros.id [delete: restrict]
Ref: equivalencias_partidas_indice.partida_id > partidas.id [delete: cascade]
Ref: equivalencias_partidas_indice.partida_estandar_id > partidas_estandar_indice.id [delete: restrict]

Ref: aportes_indice_partidas.empresa_id > empresas.id [delete: cascade]
Ref: aportes_indice_partidas.partida_orden_trabajo_id > partidas_orden_trabajo.id [delete: restrict]
Ref: aportes_indice_partidas.partida_estandar_id > partidas_estandar_indice.id [delete: restrict]
Ref: aportes_indice_partidas.revisado_por > usuarios.id [delete: set null]

Ref: registros_indice_partidas.partida_estandar_id > partidas_estandar_indice.id [delete: restrict]

Ref: perfiles_directorio_contratistas.empresa_id > empresas.id [delete: cascade]
Ref: perfiles_directorio_contratistas.rubro_id > rubros.id [delete: restrict]
```

---

## 6. Relaciones explícitas

|Origen|Destino|Cardinalidad|Borrado|Motivo|
|---|---|---|---|---|
|`usuarios.empresa_id`|`empresas.id`|N:1|CASCADE|Cada usuario pertenece a una empresa.|
|`usuarios_roles.usuario_id`|`usuarios.id`|N:1|CASCADE|Elimina asignaciones al eliminar el usuario.|
|`usuarios_roles.rol_id`|`roles.id`|N:1|RESTRICT|No eliminar roles que estén asignados.|
|`clientes.empresa_id`|`empresas.id`|N:1|CASCADE|Los clientes pertenecen a una empresa.|
|`proyectos.cliente_id`|`clientes.id`|N:1|RESTRICT|Mantiene la trazabilidad histórica.|
|`rubros.empresa_id`|`empresas.id`|N:1 opcional|CASCADE|`NULL` indica rubro global; con empresa indica personalizado.|
|`categorias.rubro_id`|`rubros.id`|N:1|RESTRICT|Una categoría pertenece a un rubro.|
|`historial_precios_recursos.recurso_id`|`recursos.id`|N:1|CASCADE|Historial dependiente del recurso.|
|`partidas.categoria_id`|`categorias.id`|N:1|RESTRICT|Evita partidas huérfanas.|
|`recursos_apu_partida.partida_id`|`partidas.id`|N:1|CASCADE|El APU depende de la partida.|
|`recursos_apu_partida.recurso_id`|`recursos.id`|N:1|RESTRICT|No eliminar insumos utilizados en APU.|
|`versiones_cotizacion.cotizacion_id`|`cotizaciones.id`|N:1|CASCADE|Las versiones pertenecen a la cotización.|
|`partidas_cotizacion.version_cotizacion_id`|`versiones_cotizacion.id`|N:1|CASCADE|Las partidas pertenecen a una versión.|
|`partidas_cotizacion.partida_id`|`partidas.id`|N:1 opcional|SET NULL|El snapshot histórico permanece aunque se borre la partida base.|
|`ordenes_trabajo.version_cotizacion_id`|`versiones_cotizacion.id`|N:1|RESTRICT|La versión aceptada debe conservarse.|
|`partidas_orden_trabajo.orden_trabajo_id`|`ordenes_trabajo.id`|N:1|CASCADE|Las partidas pertenecen a una OT.|
|`registros_ejecucion.partida_orden_trabajo_id`|`partidas_orden_trabajo.id`|N:1|CASCADE|Registro real de cada partida de OT.|
|`recursos_ejecucion.registro_ejecucion_id`|`registros_ejecucion.id`|N:1|CASCADE|Recursos reales utilizados en el registro.|
|`costos_adicionales_ejecucion.registro_ejecucion_id`|`registros_ejecucion.id`|N:1|CASCADE|Imprevistos asociados al registro.|
|`fotos_ejecucion.registro_ejecucion_id`|`registros_ejecucion.id`|N:1|CASCADE|Evidencia del avance o incidencia.|
|`consentimientos_datos.empresa_id`|`empresas.id`|1:1|CASCADE|Estado actual del consentimiento.|
|`historial_consentimientos.empresa_id`|`empresas.id`|N:1|CASCADE|Mantiene trazabilidad de cambios.|
|`partidas_estandar_indice.rubro_id`|`rubros.id`|N:1|RESTRICT|Clasificación de partidas estándar.|
|`equivalencias_partidas_indice.partida_id`|`partidas.id`|N:1|CASCADE|Relaciona partida privada con estándar.|
|`aportes_indice_partidas.partida_estandar_id`|`partidas_estandar_indice.id`|N:1|RESTRICT|Cada aporte debe corresponder a una partida comparable.|
|`registros_indice_partidas.partida_estandar_id`|`partidas_estandar_indice.id`|N:1|RESTRICT|El índice agregado se genera por partida estándar.|
|`perfiles_directorio_contratistas.rubro_id`|`rubros.id`|N:1|RESTRICT|Perfil público por especialidad.|

---

## 7. Índices y restricciones recomendadas

Además de los `UNIQUE` ya representados en DBML, el esquema físico deberá agregar índices para consultas frecuentes.

|Tabla|Índice / restricción|Columnas|Tipo|Motivo|
|---|---|---|---|---|
|`empresas`|`uq_empresas_rut`|`rut`|UNIQUE|Evitar empresas duplicadas.|
|`usuarios`|`uq_usuarios_correo`|`correo`|UNIQUE|Correo único para autenticación.|
|`usuarios`|`idx_usuarios_empresa`|`empresa_id`|BTREE|Listar usuarios por empresa.|
|`clientes`|`idx_clientes_empresa`|`empresa_id`|BTREE|Listado de clientes.|
|`proyectos`|`idx_proyectos_empresa_estado`|`empresa_id, estado`|BTREE|Dashboard de proyectos.|
|`rubros`|`uq_rubros_empresa_nombre`|`empresa_id, nombre`|UNIQUE|Evitar duplicación de rubros personalizados.|
|`categorias`|`idx_categorias_empresa_rubro`|`empresa_id, rubro_id`|BTREE|Catálogo por rubro.|
|`recursos`|`idx_recursos_empresa`|`empresa_id`|BTREE|Búsqueda de recursos.|
|`historial_precios_recursos`|`idx_historial_recurso_fecha`|`recurso_id, fecha_desde`|BTREE|Consultar precios históricos.|
|`partidas`|`idx_partidas_empresa_categoria`|`empresa_id, categoria_id`|BTREE|Cargar catálogo.|
|`recursos_apu_partida`|`uq_recurso_apu_partida`|`partida_id, recurso_id`|UNIQUE|Evitar recurso duplicado en un APU.|
|`cotizaciones`|`uq_cotizaciones_empresa_codigo`|`empresa_id, codigo`|UNIQUE|Código de cotización único.|
|`versiones_cotizacion`|`uq_version_cotizacion`|`cotizacion_id, numero_version`|UNIQUE|Número de versión único.|
|`versiones_cotizacion`|`idx_versiones_vigencia_estado`|`vigente_hasta, estado`|BTREE|Detectar cotizaciones vencidas.|
|`ordenes_trabajo`|`uq_ot_empresa_codigo`|`empresa_id, codigo`|UNIQUE|Código de OT único.|
|`ordenes_trabajo`|`idx_ot_empresa_estado`|`empresa_id, estado`|BTREE|Listar trabajos activos.|
|`registros_ejecucion`|`idx_registro_partida_fecha`|`partida_orden_trabajo_id, fecha_registro`|BTREE|Acumulados y reportes.|
|`recursos_ejecucion`|`idx_recursos_ejecucion_registro`|`registro_ejecucion_id`|BTREE|Cargar recursos reales del avance.|
|`partidas_estandar_indice`|`idx_partida_estandar_rubro`|`rubro_id`|BTREE|Búsqueda por rubro.|
|`equivalencias_partidas_indice`|`uq_equivalencia_partida`|`partida_id, partida_estandar_id`|UNIQUE|Evitar equivalencias duplicadas.|
|`aportes_indice_partidas`|`idx_aportes_estado_partida`|`estado_validacion, partida_estandar_id`|BTREE|Procesar solo aportes válidos.|
|`aportes_indice_partidas`|`idx_aportes_fecha_zona`|`fecha_ejecucion, region, comuna`|BTREE|Estadísticas por zona y período.|
|`registros_indice_partidas`|`idx_indice_partida_zona_periodo`|`partida_estandar_id, region, comuna, periodo_desde, periodo_hasta`|BTREE|Consultar índice segmentado.|
|`perfiles_directorio_contratistas`|`uq_perfil_empresa_rubro`|`empresa_id, rubro_id`|UNIQUE|Un perfil por empresa y rubro.|

---

## 8. Reglas de integridad y negocio

1. **Aislamiento por empresa:** ningún usuario debe consultar o modificar datos privados pertenecientes a una empresa distinta de la suya.
    
2. **Rubros globales y personalizados:** Construcción, Electricidad y Vidriería se crean como rubros globales (`empresa_id = NULL`). Cuando un usuario selecciona “Otro”, se crea un rubro personalizado vinculado a su empresa.
    
3. **Historial de precios:** cada modificación importante de `precio_unitario_actual` debe generar un registro en `historial_precios_recursos`.
    
4. **Inmutabilidad de versiones:** una `versiones_cotizacion` en estado `ACEPTADA`, `RECHAZADA`, `VENCIDA` o `SUPERADA` no debe modificarse. Los cambios posteriores generan una nueva versión.
    
5. **Snapshots:** las partidas y recursos de una cotización almacenan nombres y precios utilizados en ese instante. Los cambios posteriores al catálogo no alteran cotizaciones históricas.
    
6. **Partidas extraordinarias:** los cambios posteriores a una aceptación se registran mediante nuevas partidas marcadas con `es_extraordinaria = true`, sin reescribir la versión original aceptada.
    
7. **Creación de Orden de Trabajo:** una OT solo puede generarse desde una versión de cotización aceptada.
    
8. **Cierre de Orden de Trabajo:** al cerrar la OT se consolidan cantidades ejecutadas, recursos reales, costos adicionales y horas registradas.
    
9. **Cotizado vs. Ejecutado:** la desviación se obtiene comparando los costos presupuestados de `partidas_orden_trabajo` con los costos reales derivados de `recursos_ejecucion` y `costos_adicionales_ejecucion`.
    
10. **Consentimiento explícito:** solo se generan `aportes_indice_partidas` cuando la empresa tiene consentimiento vigente.
    
11. **Normalización previa al índice:** toda partida debe vincularse a una `partida_estandar_indice` antes de utilizarse para estadísticas agregadas.
    
12. **Validación del aporte:** un aporte puede tener estado `PENDIENTE`, `VALIDO`, `ANOMALO` o `EXCLUIDO`. Solo los registros `VALIDO` participan en el cálculo del índice.
    
13. **Anonimización pública:** `registros_indice_partidas` no almacena identificadores de empresa, cliente, proyecto ni márgenes comerciales.
    
14. **Tamaño de muestra:** la interfaz del índice debe mostrar `cantidad_muestras` para que el usuario pueda evaluar la confiabilidad del dato.
    
15. **Segmentación:** el índice puede calcularse por partida estándar, región, comuna y período.
    
16. **Directorio independiente:** aparecer en el directorio de contratistas es voluntario y no debe revelar qué datos aportó una empresa al índice.
    

---

## 9. Flujo de datos principal

### 9.1 Cotización y ejecución

```
Empresa
  ↓
Cliente
  ↓
Proyecto
  ↓
Rubro
  ↓
Categoría
  ↓
Partida
  ↓
APU
  ↓
Cotización
  ↓
Versión de cotización
  ↓
Partidas cotizadas
  ↓
Cotización aceptada
  ↓
Orden de Trabajo
  ↓
Partidas de OT
  ↓
Registro de ejecución
  ├── Recursos reales
  ├── Horas trabajadas
  ├── Costos adicionales
  └── Fotografías
  ↓
Cotizado vs. Ejecutado
```

### 9.2 Generación de ÍndicePartidas

```
Orden de Trabajo cerrada
        +
Consentimiento vigente
        ↓
Partida ejecutada
        ↓
Equivalencia con partida estándar
        ↓
Aporte al índice
        ↓
Validación
        ↓
Solo datos VALIDOS
        ↓
Agregación estadística
        ↓
ÍndicePartidas
```

---

## 10. Datos semilla

|Semilla|Tabla|Ambiente|Motivo|
|---|---|---|---|
|`DUENO`|`roles`|dev / test / prod|Propietario o responsable principal.|
|`ADMIN_OBRA`|`roles`|dev / test / prod|Gestión de obra y costos.|
|`JEFE_TERRENO`|`roles`|dev / test / prod|Registro y control en terreno.|
|`MAESTRO`|`roles`|dev / test / prod|Registro limitado de avances/evidencias.|
|`CONSULTA`|`roles`|dev / test / prod|Usuario de lectura.|
|`ADMIN_SISTEMA`|`roles`|dev / test / prod|Administración de plataforma.|
|Construcción|`rubros`|dev / test / prod|Rubro predefinido, `empresa_id = NULL`.|
|Electricidad|`rubros`|dev / test / prod|Rubro predefinido, `empresa_id = NULL`.|
|Vidriería|`rubros`|dev / test / prod|Rubro predefinido, `empresa_id = NULL`.|
|Catálogo demo|`categorias`, `recursos`, `partidas`, `recursos_apu_partida`|dev / test|Pruebas funcionales.|
|Partidas estándar demo|`partidas_estandar_indice`|dev / test|Pruebas del proceso de normalización e índice.|

> **“Otro” no se almacena como rubro global definitivo.** En la interfaz funciona como una acción para crear un nuevo rubro personalizado, por ejemplo Florería, Mecánica o Climatización.

---

## 11. Migraciones

Se utilizará `golang-migrate` con archivos SQL versionados dentro de:

```
backend/migrations/
```

### Propuesta de migraciones

```
backend/migrations/
├── 000001_create_empresas_usuarios_roles.up.sql
├── 000001_create_empresas_usuarios_roles.down.sql
│
├── 000002_create_clientes_proyectos_rubros.up.sql
├── 000002_create_clientes_proyectos_rubros.down.sql
│
├── 000003_create_catalogo_apu_precios.up.sql
├── 000003_create_catalogo_apu_precios.down.sql
│
├── 000004_create_cotizaciones_versiones.up.sql
├── 000004_create_cotizaciones_versiones.down.sql
│
├── 000005_create_ordenes_trabajo_ejecucion.up.sql
├── 000005_create_ordenes_trabajo_ejecucion.down.sql
│
├── 000006_create_consentimientos.up.sql
├── 000006_create_consentimientos.down.sql
│
├── 000007_create_normalizacion_indice.up.sql
├── 000007_create_normalizacion_indice.down.sql
│
├── 000008_create_aportes_indice.up.sql
├── 000008_create_aportes_indice.down.sql
│
├── 000009_create_indice_agregado.up.sql
├── 000009_create_indice_agregado.down.sql
│
├── 000010_create_directorio.up.sql
└── 000010_create_directorio.down.sql
```

### Reglas para migraciones

1. Cada migración debe poseer archivo `.up.sql` y `.down.sql`.
2. Todas las FK deben declarar explícitamente `ON DELETE`.
3. No agregar columnas `NOT NULL` sobre tablas con datos sin establecer antes un valor válido.
4. Las migraciones ejecutadas en producción no deben editarse; los cambios posteriores requieren una nueva migración.
5. Los datos semilla deben mantenerse separados o controlados para evitar duplicación.
6. Las extensiones necesarias de PostgreSQL, como `pgcrypto` para `gen_random_uuid()`, deben habilitarse en la migración inicial.

---

## 12. Diferencias respecto del diseño anterior

|Cambio|Diseño anterior|Diseño actualizado|
|---|---|---|
|Idioma de entidades|Inglés|Español|
|Rubro “Otro”|Un único rubro genérico + configuración aparte|“Otro” crea un rubro personalizado real por empresa|
|Precios|Solo precio actual|Precio actual + historial|
|Aceptación|Solo estado|Estado + fechas + medio + motivo de rechazo|
|Ejecución|Cantidad, horas y costo general|Cantidad + horas + recursos reales + costos adicionales|
|Índice|Estadísticas directas por texto|Partidas estándar + equivalencias + aportes + validación|
|Calidad de datos|No representada explícitamente|Estados PENDIENTE / VALIDO / ANOMALO / EXCLUIDO|
|Segmentación del índice|Básica|Región + comuna + período|
|Consentimiento|Estado actual|Estado actual + historial|
|Directorio|Perfil por empresa|Perfil voluntario por empresa y rubro|

---

## 13. Correspondencia con funcionalidades principales

|Funcionalidad|Tablas principales|
|---|---|
|Gestión de usuarios y roles|`empresas`, `usuarios`, `roles`, `usuarios_roles`|
|Clientes y proyectos|`clientes`, `proyectos`|
|Rubros predefinidos y “Otro”|`rubros`|
|Catálogo|`categorias`, `partidas`|
|APU|`recursos`, `recursos_apu_partida`|
|Historial de precios|`historial_precios_recursos`|
|Cotizaciones|`cotizaciones`, `versiones_cotizacion`, `partidas_cotizacion`|
|Snapshot de recursos|`recursos_partida_cotizacion`|
|Orden de Trabajo|`ordenes_trabajo`, `partidas_orden_trabajo`|
|Registro de ejecución|`registros_ejecucion`|
|Materiales/costos reales|`recursos_ejecucion`, `costos_adicionales_ejecucion`|
|Evidencia fotográfica|`fotos_ejecucion`|
|Cotizado vs. Ejecutado|`partidas_orden_trabajo` + tablas de ejecución|
|Consentimiento|`consentimientos_datos`, `historial_consentimientos`|
|Normalización de partidas|`partidas_estandar_indice`, `equivalencias_partidas_indice`|
|Validación de datos|`aportes_indice_partidas`|
|ÍndicePartidas|`registros_indice_partidas`|
|Directorio|`perfiles_directorio_contratistas`|

---

## 14. Checklist de base de datos

- [x] Empresas y usuarios con separación por empresa.
- [x] Roles y permisos representados.
- [x] Clientes y proyectos.
- [x] Construcción, Electricidad y Vidriería como rubros predefinidos.
- [x] Rubros personalizados mediante opción “Otro”.
- [x] Catálogo de partidas.
- [x] Recursos y APU.
- [x] Historial de precios.
- [x] Cotizaciones versionadas.
- [x] Vigencia, envío, aceptación y rechazo de cotizaciones.
- [x] Snapshots históricos de partidas y recursos.
- [x] Conversión de cotización aceptada a Orden de Trabajo.
- [x] Partidas de Orden de Trabajo.
- [x] Registro simplificado de cantidades y horas.
- [x] Registro detallado de recursos realmente utilizados.
- [x] Registro separado de costos adicionales.
- [x] Fotografías e incidencias.
- [x] Estructura necesaria para Cotizado vs. Ejecutado.
- [x] Consentimiento actual e histórico.
- [x] Normalización de partidas para comparar datos entre empresas.
- [x] Validación y depuración de aportes.
- [x] Índice agregado sin identificadores comerciales.
- [x] Segmentación territorial y temporal del índice.
- [x] Directorio voluntario independiente del índice.
- [x] Modelo compatible con PostgreSQL y backend Go.
- [x] Estrategia de migraciones definida.
