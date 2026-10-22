# Rules — Reglas de trabajo y mantenimiento del proyecto

> Plantilla reutilizable para definir cómo se trabaja en el repositorio, cómo se resuelven bugs, qué documento manda sobre cada tema y cómo evitar que la documentación se desordene.

## 1. Propósito

Este archivo define las reglas operativas del repositorio. Su objetivo es que cualquier persona o agente pueda implementar, corregir y documentar cambios sin perder contexto.

## 2. Fuente de verdad por tema

| Pregunta | Documento |
| --- | --- |
| ¿Qué producto se construye y para quién? | [PRD.md](PRD.md) |
| ¿Cómo debe verse y comportarse el frontend? | [Design.md](Design.md) |
| ¿Qué datos existen y cómo se relacionan? | [DB_Diagram.md](DB_Diagram.md) |
| ¿Cómo está organizado técnicamente el sistema? | [Architecture.md](Architecture.md) |
| ¿En qué orden se implementa el backend? | [Backend_Implementation_Phases.md](Backend_Implementation_Phases.md) |
| ¿Cómo se trabaja, corrige y documenta? | [Rules.md](Rules.md) |

## 3. Regla de actualización documental

Si un cambio modifica comportamiento, datos, API, UI o arquitectura, también debe actualizar el documento correspondiente.

| Tipo de cambio | Documento a revisar |
| --- | --- |
| Nueva funcionalidad | `PRD.md`, fases, contratos/API si existen |
| Cambio visual o componente | `Design.md` |
| Nueva tabla/campo/relación | `DB_Diagram.md` |
| Cambio de capas, stack o estructura | `Architecture.md` |
| Cambio de orden de trabajo | `Backend_Implementation_Phases.md` |
| Cambio de proceso del equipo | `Rules.md` |

## 4. Flujo estándar para implementar una funcionalidad

1. Leer el requerimiento en `PRD.md`.
2. Confirmar entidades en `DB_Diagram.md`.
3. Confirmar arquitectura y capa responsable en `Architecture.md`.
4. Definir o actualizar contrato API si aplica.
5. Implementar por flujo vertical.
6. Agregar tests proporcionales al riesgo.
7. Probar manualmente el caso feliz y al menos un caso de error.
8. Actualizar documentación.
9. Registrar decisiones relevantes.

## 5. Flujo estándar para corregir bugs

### 5.1 Diagnóstico

Antes de cambiar código:

- Reproducir o entender el bug.
- Identificar comportamiento esperado.
- Identificar comportamiento actual.
- Ubicar capa responsable.
- Revisar si existe test que debería haber fallado.

### 5.2 Corrección

- Corregir la causa, no sólo el síntoma.
- Agregar o ajustar test si el bug es reproducible.
- Evitar cambios de alcance no relacionados.
- Si el bug revela ambigüedad de producto, actualizar `PRD.md`.
- Si el bug revela inconsistencia de datos, actualizar `DB_Diagram.md` o migración.

### 5.3 Cierre

Registrar:

```text
Bug:
Causa:
Solución:
Test/validación:
Documentación actualizada:
```

## 6. Responsabilidades por capa

### Frontend

Debe:

- Mostrar datos y estados de interacción.
- Validar formato para mejorar UX.
- Consumir contratos API.
- Manejar carga, vacío, error y permisos.

No debe:

- Ser la única barrera de autorización.
- Decidir reglas críticas de negocio.
- Saltarse validaciones backend.

### Backend handlers/controllers

Debe:

- Parsear request.
- Validar sintaxis básica.
- Extraer contexto auth.
- Llamar servicios/casos de uso.
- Traducir errores a HTTP.

No debe:

- Contener reglas de negocio complejas.
- Ejecutar queries directas.
- Decidir permisos contextuales profundos sin servicio.

### Services/use cases

Debe:

- Implementar reglas de negocio.
- Orquestar transacciones.
- Validar permisos contextuales.
- Coordinar repositorios, jobs y auditoría.

### Repositories/gateways

Debe:

- Encapsular persistencia o proveedores externos.
- Exponer operaciones entendibles por dominio.
- No decidir reglas de negocio.

## 7. Reglas de commits y ramas

> Ajustar al flujo real del equipo.

- Ramas: `[feature/nombre]`, `[fix/nombre]`, `[chore/nombre]`.
- Commits pequeños y descriptivos.
- No mezclar refactors grandes con features.
- No commitear secretos, dumps ni archivos generados pesados.
- Todo cambio riesgoso debe tener forma de rollback o mitigación.

## 8. Reglas de testing

| Tipo de cambio | Validación mínima |
| --- | --- |
| Regla de negocio | Test unitario del servicio/use case |
| Query compleja | Test de integración o fixture |
| Endpoint | Test de handler/API o prueba manual documentada |
| UI crítica | Prueba de estados principales |
| Migración | Arranque limpio + datos existentes si aplica |
| Bug corregido | Test que falla antes y pasa después, si es viable |

## 9. Manejo de decisiones

Toda decisión que afecte el futuro del proyecto debe registrarse.

Formato:

```text
Decisión:
Contexto:
Opciones consideradas:
Elección:
Consecuencia:
Fecha:
```

## 10. Manejo de deuda técnica

La deuda técnica aceptada debe quedar visible.

| Deuda | Motivo | Impacto | Plan | Fecha límite |
| --- | --- | --- | --- | --- |
| `[Deuda]` | `[Por qué se acepta]` | `[Riesgo]` | `[Cómo se pagará]` | `[Fecha/fase]` |

## 11. Reglas para agentes o asistentes de coding

- Antes de editar, revisar documentos fuente del tema.
- No asumir arquitectura si está definida en `Architecture.md`.
- No crear componentes visuales fuera de `Design.md` sin actualizarlo.
- No crear tablas/campos sin actualizar `DB_Diagram.md`.
- No marcar tareas como completas sin validación.
- Preservar cambios existentes del usuario.
- Explicar supuestos importantes.

## 12. Checklist antes de entregar cambios

- [ ] El código compila o el cambio es sólo documental.
- [ ] Se ejecutaron pruebas relevantes o se explicó por qué no.
- [ ] No hay cambios no relacionados.
- [ ] La documentación afectada fue actualizada.
- [ ] Los bugs corregidos tienen causa y validación.
- [ ] Las decisiones nuevas quedaron registradas.
