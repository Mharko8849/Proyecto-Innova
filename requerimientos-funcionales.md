# Requerimientos Funcionales (RF) — ÍndicePartidas (Desafío 4)

> **Proyecto-Innova** — Grupo 1, Innovación y Emprendimiento 2-2026, USACH.

---

### Módulo 1 — Catálogo de Partidas y APU

| N° | Nombre de Requerimiento | Descripción (Historia de Usuario) | Fuente / Justificación |
| :---: | :--- | :--- | :--- |
| 1 | Visualización del catálogo de partidas | Como usuario del sistema quiero poder ver el catálogo de partidas constructivas con sus materiales, rendimientos, precios unitarios e historial de costos. | Catálogo D4 (Punto 1) + E-05 (Resp. 3, 4, 17) |
| 2 | C.R.U.D de Partidas y Análisis de Precio Unitario (APU) | Como usuario "Administrador de Obra" o "Contratista" quiero poder crear, observar, actualizar y/o eliminar partidas con su respectivo Análisis de Precio Unitario (materiales, mano de obra y rendimiento). | Catálogo D4 (Punto 1) + E-05 (Resp. 3 y 17: uso de APU) |
| 3 | C.R.U.D de Categorías de Partidas | Como usuario "Administrador de Obra" o "Contratista" quiero poder crear, actualizar y/o eliminar categorías de partidas (ej. Obra Gruesa, Terminaciones, Instalaciones) asociadas al catálogo. | E-05 (Resp. 1 y 8: distinción entre Terminaciones y Obra Gruesa) |
| 4 | Actualización de indicadores y precios de insumos | Como usuario "Administrador de Obra" o "Contratista" quiero poder fijar y actualizar los precios de materiales y parámetros base (valor UF, dólar referencial) para mantener vigentes los costos del catálogo. | Catálogo D4 (Punto 1) + E-05 (Resp. 4: fijación de UF y dólar entre $900–$1000) |

### Módulo 2 — Cotizaciones y Presupuestos

| N° | Nombre de Requerimiento | Descripción (Historia de Usuario) | Fuente / Justificación |
| :---: | :--- | :--- | :--- |
| 5 | C.R.U.D de Cotizaciones por partidas | Como usuario "Contratista" o "Administrador de Obra" quiero poder armar, editar o eliminar un presupuesto seleccionando partidas del catálogo e ingresando las cubicaciones (cantidades) del proyecto. | Catálogo D4 (Punto 2) + E-05 (Resp. 1: cubicación y presupuesto) |
| 5.1 | Configuración de márgenes y gastos generales | Como usuario "Contratista" o "Dueño" quiero poder configurar de forma confidencial los porcentajes de gastos generales, costos indirectos, sobrecostos de seguridad y margen de ganancia en cada cotización. | Catálogo D4 (Punto 2) + E-05 (Resp. 4, 5 y 16: suma alzada y temor a exponer márgenes) |
| 6 | Control de versiones de cotización | Como usuario "Contratista" o "Administrador de Obra" quiero poder generar nuevas versiones de un presupuesto manteniendo el registro exacto de qué partidas, cantidades o precios cambiaron entre una versión y otra. | Catálogo D4 (Punto 3: versionado con registro de cambios) |
| 7 | Adición de partidas extraordinarias (Modificaciones) | Como usuario "Administrador de Obra" quiero poder agregar nuevas partidas extraordinarias a un proyecto ya presupuestado sin alterar el presupuesto base original (modalidad suma alzada). | E-05 (Resp. 4 y 5: en suma alzada no se modifica el presupuesto base, se crea partida nueva) |

### Módulo 3 — Flujo Comercial y Órdenes de Trabajo

| N° | Nombre de Requerimiento | Descripción (Historia de Usuario) | Fuente / Justificación |
| :---: | :--- | :--- | :--- |
| 8 | Envío de cotización y registro de aceptación | Como usuario "Contratista" o "Administrador de Obra" quiero poder exportar/enviar el presupuesto al mandante y registrar formalmente su estado (En revisión, Aceptado, Rechazado). | Catálogo D4 (Punto 4) + E-05 (Resp. 2, 4 y 5) |
| 9 | Conversión automática a Orden de Trabajo | Como usuario "Administrador de Obra" quiero poder convertir una cotización aceptada en una Orden de Trabajo activa en un clic, sin necesidad de reescribir las partidas ni sus cubicaciones. | Catálogo D4 (Punto 5: conversión sin reescribir) |

### Módulo 4 — Captura de Ejecución en Terreno

| N° | Nombre de Requerimiento | Descripción (Historia de Usuario) | Fuente / Justificación |
| :---: | :--- | :--- | :--- |
| 10 | Registro simplificado de costos y cantidades en terreno | Como usuario "Administrador de Obra" o "Jefe de Terreno" quiero poder registrar desde un dispositivo móvil o escritorio las cantidades reales utilizadas y los costos ejecutados por partida. | Catálogo D4 (Punto 6) + E-05 (Resp. 11 y 13: lo hace el Admin. de Obra / Jefe de Terreno) |
| 11 | Reporte fotográfico de avances e incidentes | Como usuario "Jefe de Terreno" o "Maestro" quiero poder subir fotografías y notas rápidas desde el celular asociadas a una partida para evidenciar errores de plano o imprevistos en obra gruesa/terminaciones. | E-05 (Resp. 8, 9 y 12: maestros envían fotos desde el celular cuando detectan errores) |

### Módulo 5 — Comparación, Analítica e Índice

| N° | Nombre de Requerimiento | Descripción (Historia de Usuario) | Fuente / Justificación |
| :---: | :--- | :--- | :--- |
| 12 | Comparación Cotizado vs. Ejecutado por partida | Como usuario "Contratista" o "Administrador de Obra" quiero poder visualizar un cuadro comparativo por partida entre lo presupuestado y el costo real ejecutado, mostrando la desviación en monto y porcentaje. | Catálogo D4 (Punto 7: corazón del producto) + E-05 (Resp. 6 y 7: desviación del 2 % al 10 %) |
| 13 | Estimación predictiva de desviación de costos (ML) | Como usuario "Contratista" quiero poder ver una alerta o estimación del costo real esperado al armar una partida, calculada a partir del histórico de desviaciones propias y agregadas. | Catálogo D4 (Sección Dato y aprendizaje automático) |
| 14 | Dashboard de rentabilidad e indicadores propios | Como usuario "Dueño" o "Contratista" quiero poder ver un tablero con la tasa de adjudicación de cotizaciones, desviación promedio (ej. 2 % a 10 %) y margen real obtenido tras finalizar cada obra. | E-05 (Resp. 2, 7 y 19: a la constructora le interesa la información interna de su empresa) |
| 15 | Consulta del Índice Agregado Anonimizado | Como usuario "Tasador / Aseguradora / Mandante" (Segundo Cliente) quiero poder consultar y exportar el índice agregado de precios y rendimientos reales de partidas del mercado sin visualizar datos comerciales sensibles ni márgenes individuales. | Catálogo D4 (Segundo cliente) + E-05 (Resp. 15, 16 y 19) |
