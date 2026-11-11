# Experimento — Grupo 1 — ÍndicePartidas

# 1. Nuestros supuestos

|Supuesto|Tipo|Importancia|Evidencia|
|---|---|---|---|
|Los contratistas necesitan comparar costos cotizados y reales|Deseable|Alta|E-01, E-02, E-05|
|Utilizarán una herramienta si el registro es sencillo|Deseable|Alta|E-01, E-02, E-05|
|Es posible registrar costos desde teléfonos|Factible|Alta|Hipótesis técnica|
|Aportarán datos anonimizados|Viable|Alta|Interés condicionado de entrevistados|
|Existirá un segundo cliente para el índice|Viable|Media|E-03|
|El software podrá configurarse para otros rubros|Factible|Media|Hipótesis de diseño|
|Los usuarios valoran cotizaciones por componentes|Deseable|Alta|E-01, E-02|
|Los usuarios contratarán planes después de una prueba gratuita|Viable|Alta|Preferencia declarada, pago pendiente|
|Los mecanismos de validación harán confiable el índice|Factible|Alta|Riesgo identificado por E-03|

---

# 2. Supuesto más riesgoso

Los contratistas registrarán regularmente sus costos reales si el proceso es suficientemente rápido y sencillo.

Este supuesto es fundamental porque el valor principal del software depende de comparar los costos presupuestados con los ejecutados.

Si los usuarios solamente crean cotizaciones, el sistema no podrá construir correctamente el historial de costos reales.

---

# 3. Tarjeta de prueba

## Creemos que

Un contratista independiente o responsable de una pequeña empresa puede completar el ciclo básico en la plataforma sin fricción excesiva si cuenta con una interfaz simplificada y un catálogo con partidas predefinidas. Específicamente creemos que:

1. El registro de cuenta y configuración inicial tomará **menos de 1 minuto**.
2. La elaboración de una cotización para un trabajo estándar (usando partidas preconfiguradas con materiales y mano de obra) tomará **menos de 10 minutos**.
3. La captura de un costo real ejecutado en terreno tomará **30 segundos o menos**.

## Para verificarlo

Se construyó un prototipo funcional en **Figma** que simula el flujo de la aplicación:

- Pantalla de registro de usuario y selección de rubro inicial.
- Módulo de cotización con **partidas base ya predefinidas** (con sus componentes de materiales y mano de obra).
- Generación y aprobación de cotización.
- Formulario móvil simplificado para registrar costos y cantidades reales ejecutadas en terreno.
- Vista de comparación automática entre costo presupuestado y costo real.

## Mediremos

1. Tiempo cronometrado en cada etapa: registro de cuenta, creación de cotización y registro de costos.
2. Errores, desorientaciones y dudas durante la navegación.
3. Tasa de tareas completadas de forma autónoma versus con asistencia del facilitador.
4. Solicitudes de cambio o campos ignorados por el usuario.

## Tenemos razón si

- El registro de cuenta se completa en $\le$ 1 minuto.
- La cotización se elabora en $\le$ 10 minutos (reduciendo drásticamente las 2 a 3 horas habituales de los entrevistados).
- Los registros de costos habituales se completan en $\le$ 30 segundos.
- La mayoría de los evaluados logra completar el recorrido sin bloqueos insalvables y comprende el resultado de la comparación de costos.

---

# 4. Plan

|Actividad|Responsable|Fecha|
|---|---|---|
|Diseñar formulario|Grupo 1|30 de septiembre|
|Realizar prueba interna|Grupo 1|30 de septiembre|
|Contactar participantes|Grupo 1|30 de septiembre|
|Ejecutar piloto|Grupo 1|30 de octubre|
|Revisar registros|Grupo 1|29 oct. al 2 nov.|
|Cerrar experimento|Grupo 1|3 de noviembre|
|Completar aprendizaje|Grupo 1|3 de noviembre|

---

# 5. Resultados de la prueba observada con el prototipo

La prueba se aplicó a dos personas del segmento objetivo de vidriería:

- **P-01:** Instalador independiente de vidriería y ventanería (E-01).
- **P-02:** Dueño de taller/pyme de vidriería y aluminio (E-02).

## Tiempos cronometrados por tarea

| Etapa evaluada | Meta esperada | Participante 1 (P-01) | Participante 2 (P-02) | ¿Cumple meta? | Observación de asistencia |
|---|---|---|---|---|---|
| **1. Registro de cuenta (Onboarding)** | $\le$ 1:00 min | 1 min 15 s | 0 min 52 s | Parcial | P-01 dudó brevemente al seleccionar el rubro; P-02 autónomo. |
| **2. Elaboración de cotización (Partidas base)** | $\le$ 10:00 min | 4 min 20 s | 3 min 10 s | **Sí (amplio)** | Al utilizar partidas predefinidas, ambos cotizaron en menos de la mitad del tiempo meta. |
| **3. Registro de costo (Intento 1: exploratorio)** | $\le$ 30 s | 41 s | 36 s | No (primera vez) | Se detuvieron a entender el formulario y campos opcionales. |
| **3. Registro de costo (Intento 2: compra insumos)** | $\le$ 30 s | **22 s** | **19 s** | **Sí** | Registro directo de compra de perfiles/vidrio. Rápido y sin ayuda. |
| **3. Registro de costo (Intento 3: mano de obra)** | $\le$ 30 s | **24 s** | **21 s** | **Sí** | Registro rutinario sin asistencia. |

## Fricciones y hallazgos observados

### Dónde se pierden

1. Tras aceptar la cotización, no sabían claramente dónde registrar los costos reales.
2. Al momento de ingresar los costos se perdían con el formato del sistema.

### Qué ignoran

Ambos omitieron los campos de "Horas de trabajo" y "Notas descriptivas"; se concentraron exclusivamente en el monto en dinero y la cantidad.

### Qué ayuda necesitaron

1. Consultaron si el valor a ingresar correspondía a monto neto o con IVA.
2. Necesitaron ayuda para navegar entre las distintas pantallas.

## Cambios aplicados al prototipo tras el feedback

1. **Acceso directo:** Incorporar un botón visible para registrar costos reales de forma inmediata tras aprobar la cotización, evitando desorientación.
2. **Navegación guiada:** Mejorar los enlaces y botones de transición entre pantallas para facilitar el flujo continuo.
3. **Formato simplificado:** Rediseñar y clarificar el formato de ingreso de costos para evitar confusiones de estructura.
4. **Claridad en montos:** Agregar etiqueta explícita indicando si el valor ingresado es Neto o Total con IVA.
5. **Reducción de campos:** Ocultar o colapsar campos de horas y observaciones para no recargar la vista en terreno.

---

# 6. Tarjeta de aprendizaje (Tarjeta de salida)

## Creíamos que

Un contratista completaría el registro en menos de 1 minuto, cotizaría en menos de 10 minutos con partidas preconfiguradas y registraría costos reales en menos de 30 segundos, permitiendo su uso continuado.

## Observamos

- En la prueba de usabilidad, los tiempos de cotización (3 a 4 minutos) y de registro habitual de costos (19 a 24 segundos) cumplieron con holgura los umbrales de velocidad.
- Sin embargo, respecto al experimento de campo prolongado original (5 participantes registrando autónomamente durante varios días), este quedó parcial / inconcluso debido a la falta de seguimiento.

## Aprendimos que

- La estructura por partidas base agiliza significativamente la cotización, pero el flujo entre cotizar y registrar debe ser mucho más directo y guiado.
- La velocidad de captura es técnicamente viable ($\le$ 30 s), pero la interfaz debe evitar formatos confusos y reducir pasos intermedios de navegación entre pantallas.

## Decisión

**Perseverar con ajustes de producto**. Incorporar las mejoras de usabilidad, formato y navegación identificadas en el prototipo.

---

# 7. Conclusión

Las pruebas observadas con el prototipo en Figma permitieron comprobar que la elaboración de cotizaciones mediante partidas predefinidas y el registro ágil de costos son técnicamente viables en tiempos reducidos, cumpliendo con los umbrales de velocidad planteados.

Sin embargo, la experiencia de los participantes demostró que la usabilidad, la claridad en el formato de ingreso y la navegación directa entre pantallas son factores decisivos para evitar desorientación y asegurar la adopción en terreno.

Por ello, el equipo resuelve perseverar en el desarrollo de la solución, priorizando la implementación de las mejoras de interfaz identificadas en el prototipo para consolidar un flujo simple e intuitivo.
Se recomienda continuar con el prototipo, medir tiempos reales, comprobar el registro recurrente de costos y realizar un piloto comercial antes de considerar validado el modelo de negocio.