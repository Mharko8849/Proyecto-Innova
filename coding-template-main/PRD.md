# PRD — Product Requirements Document

> Plantilla reutilizable para describir qué producto se está construyendo, para quién, por qué existe, qué alcance tiene y cómo se valida que cumple su objetivo.

## 1. Identificación del producto

| Campo | Descripción |
| --- | --- |
| Nombre del producto | `[Nombre del producto]` |
| Versión del documento | `v0.1` |
| Fecha | `[YYYY-MM-DD]` |
| Responsable de producto | `[Nombre / rol]` |
| Responsable técnico | `[Nombre / rol]` |
| Estado | `Borrador / En revisión / Aprobado / Obsoleto` |

## 2. Resumen ejecutivo

Describe el producto en 3 a 5 líneas.

```text
[Producto] permite a [usuario objetivo] resolver [problema principal] mediante [capacidades principales], reduciendo/mejorando [impacto esperado].
```

## 3. Contexto y problema

### Problema actual

- `[Dolor 1]`
- `[Dolor 2]`
- `[Dolor 3]`

### Evidencia o señales

- `[Datos, observaciones, entrevistas, procesos manuales, tickets, errores frecuentes]`

### Consecuencia de no resolverlo

- `[Costo operativo, pérdida de información, mala experiencia, riesgo legal, etc.]`

## 4. Objetivos

### Objetivo principal

- `[Resultado principal medible]`

### Objetivos secundarios

- `[Objetivo secundario 1]`
- `[Objetivo secundario 2]`
- `[Objetivo secundario 3]`

### No objetivos

Todo lo que explícitamente queda fuera para evitar scope creep.

- `[Fuera de alcance 1]`
- `[Fuera de alcance 2]`

## 5. Usuarios objetivo

| Usuario / actor | Necesidad | Nivel de acceso | Frecuencia de uso |
| --- | --- | --- | --- |
| `[Actor 1]` | `[Qué necesita lograr]` | `[Rol/permisos]` | `[Diaria/semanal/eventual]` |
| `[Actor 2]` | `[Qué necesita lograr]` | `[Rol/permisos]` | `[Diaria/semanal/eventual]` |

## 6. Ficha técnica del producto

| Categoría | Definición |
| --- | --- |
| Tipo de producto | `[SaaS / backoffice / marketplace / sistema interno / app móvil / landing / API / etc.]` |
| Dominio | `[Eventos, educación, salud, finanzas, logística, etc.]` |
| Plataforma | `[Web / móvil / escritorio / API / híbrido]` |
| Dispositivos prioritarios | `[Desktop / tablet / móvil / lector QR / kiosko / etc.]` |
| Integraciones | `[Email, pagos, SSO, ERP, CRM, mapas, calendarios, etc.]` |
| Datos sensibles | `[Sí/No + cuáles]` |
| Requisitos regulatorios | `[Privacidad, auditoría, retención, accesibilidad, etc.]` |

## 7. Alcance funcional

### Módulos del producto

| Módulo | Propósito | Actor principal | Prioridad |
| --- | --- | --- | --- |
| `[Módulo 1]` | `[Qué resuelve]` | `[Actor]` | `MVP / Fase 2 / Futuro` |
| `[Módulo 2]` | `[Qué resuelve]` | `[Actor]` | `MVP / Fase 2 / Futuro` |

### Requerimientos funcionales

Usa IDs estables. No reutilices IDs eliminados.

| ID | Módulo | Requerimiento | Prioridad | Estado |
| --- | --- | --- | --- | --- |
| RF01 | `[Módulo]` | El sistema debe `[acción verificable]`. | `Must / Should / Could` | `Pendiente` |
| RF02 | `[Módulo]` | El sistema debe `[acción verificable]`. | `Must / Should / Could` | `Pendiente` |

### Detalle por requerimiento

#### RF01 — `[Nombre del requerimiento]`

**Descripción:**  
`[Explica qué debe ocurrir y por qué.]`

**Actores:**  
`[Actor primario, actores secundarios]`

**Precondiciones:**

- `[Condición previa 1]`
- `[Condición previa 2]`

**Flujo principal:**

1. `[Paso 1]`
2. `[Paso 2]`
3. `[Paso 3]`

**Flujos alternativos / errores:**

- `[Caso alternativo 1]`
- `[Caso alternativo 2]`

**Criterios de aceptación BDD:**

- **Dado** `[contexto]`, **cuando** `[acción]`, **entonces** `[resultado observable]`.
- **Dado** `[contexto de error]`, **cuando** `[acción inválida]`, **entonces** `[rechazo esperado]`.

## 8. Requerimientos no funcionales

| ID | Categoría | Requerimiento verificable | Métrica / criterio |
| --- | --- | --- | --- |
| RNF01 | Seguridad | `[Ej: sesiones seguras, cifrado, hash, RBAC]` | `[Cómo se valida]` |
| RNF02 | Performance | `[Ej: tiempo máximo de respuesta]` | `[p95 < X ms]` |
| RNF03 | Disponibilidad | `[Ej: tolerancia a reinicios]` | `[Criterio]` |
| RNF04 | Accesibilidad | `[Ej: WCAG 2.2 AA]` | `[Checklist / auditoría]` |
| RNF05 | Auditoría | `[Ej: acciones sensibles auditadas]` | `[Eventos mínimos]` |

## 9. Priorización del MVP

### Incluido en MVP

- `[Función imprescindible 1]`
- `[Función imprescindible 2]`
- `[Función imprescindible 3]`

### Diferido

- `[Función no crítica 1]`
- `[Función no crítica 2]`

### Descartado

- `[Función descartada 1]`

## 10. Métricas de éxito

| Métrica | Línea base | Meta | Cómo se mide |
| --- | --- | --- | --- |
| `[Métrica 1]` | `[Actual]` | `[Meta]` | `[Fuente]` |
| `[Métrica 2]` | `[Actual]` | `[Meta]` | `[Fuente]` |

## 11. Riesgos y decisiones

| Riesgo / decisión | Impacto | Mitigación / resolución | Estado |
| --- | --- | --- | --- |
| `[Riesgo 1]` | `Bajo / Medio / Alto` | `[Plan]` | `Abierto` |
| `[Decisión 1]` | `[Impacto]` | `[Motivo]` | `Aprobada` |

## 12. Dependencias

- `[Dependencia técnica]`
- `[Dependencia de negocio]`
- `[Dependencia externa]`

## 13. Checklist de completitud

- [ ] El problema está claramente definido.
- [ ] Los usuarios objetivo están identificados.
- [ ] El MVP está separado de fases futuras.
- [ ] Cada RF tiene criterio de aceptación verificable.
- [ ] Cada RNF tiene forma de medición.
- [ ] Las decisiones relevantes están registradas.
- [ ] Los riesgos principales tienen mitigación.
