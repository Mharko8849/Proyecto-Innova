# PRD — Product Requirements Document

> **Proyecto:** ÍndicePartidas 
> **Asignatura:** Innovación y Emprendimiento 2-2026 — USACH  
> **Equipo:** Grupo 1

---

## 1. Identificación del producto

|Campo|Descripción|
|---|---|
|Nombre del producto|**ÍndicePartidas**|
|Versión del documento|`v0.1`|
|Fecha|`2026-10-22`|
|Responsable de producto|Grupo 1|
|Responsable técnico|Equipo de desarrollo — Grupo 1|
|Estado|`En revisión`|

---

## 2. Resumen ejecutivo

**ÍndicePartidas** es una plataforma web orientada a pequeñas y medianas empresas, contratistas y trabajadores independientes que necesitan elaborar cotizaciones y controlar posteriormente el costo real de sus trabajos.

La plataforma permitirá crear presupuestos mediante partidas, registrar la ejecución real de los trabajos y comparar automáticamente lo cotizado con lo realmente gastado. Inicialmente contará con los rubros de **Construcción, Electricidad y Vidriería**, además de una opción **“Otro”**, mediante la cual cada usuario podrá configurar su propia estructura de cotización.

A partir del historial generado, el sistema permitirá conocer desviaciones, rendimientos y rentabilidad, además de alimentar de manera anonimizada el **ÍndicePartidas**, cuyo objetivo es generar referencias de costos y rendimientos reales basadas en trabajos efectivamente ejecutados.

---

## 3. Contexto y problema

### Problema actual

- Pequeñas empresas y contratistas dedican una cantidad importante de tiempo a preparar cotizaciones de forma manual mediante cuadernos, planillas Excel o presupuestos anteriores.
    
- Es frecuente reutilizar cotizaciones antiguas que contienen precios desactualizados o estructuras que no corresponden exactamente al nuevo trabajo.
    
- Muchos contratistas no realizan una comparación formal entre lo presupuestado y lo realmente gastado al finalizar un proyecto.
    
- La falta de seguimiento impide conocer con precisión el margen real obtenido en cada trabajo.
    
- Errores de cubicación, variaciones en precios de materiales, mayor uso de mano de obra e imprevistos pueden provocar desviaciones importantes respecto del presupuesto inicial.
    
- La información histórica de costos y rendimientos suele quedar dispersa o no registrarse.
    
- No existe una fuente accesible que reúna sistemáticamente costos y rendimientos reales de partidas ejecutadas por pequeños y medianos contratistas.
    

### Evidencia o señales

Las entrevistas realizadas mostraron que:

- Algunos contratistas tardan entre **30 minutos y 2 horas** en realizar una cotización.
    
- Otros destinan aproximadamente **2 a 3 horas por sesión** a preparar presupuestos.
    
- En empresas pequeñas de construcción, una cubicación y presupuesto puede tomar entre **1 y 2 días**.
    
- Algunos entrevistados no realizan ninguna comparación formal entre lo cotizado y lo ejecutado.
    
- Se detectaron desviaciones informadas de aproximadamente **2 % a 10 %** en construcción y de hasta **20 % a 30 %** en trabajos más pequeños o especializados.
    
- Algunos contratistas reconocen no conocer su margen real de ganancia.
    
- Existe disposición a registrar información si el proceso es rápido y sencillo.
    
- Existe disposición a compartir información anonimizada siempre que no se exponga la identidad, clientes o márgenes comerciales.
    
- Un potencial segundo cliente manifestó interés en consultar costos de mercado y disponer adicionalmente de contactos de prestadores por especialidad.
    

### Consecuencia de no resolverlo

- Cotizaciones realizadas por debajo del costo real.
    
- Pérdida de rentabilidad.
    
- Riesgo de problemas de liquidez.
    
- Repetición de errores de presupuestación.
    
- Dependencia de estimaciones subjetivas o información antigua.
    
- Pérdida del conocimiento generado en proyectos anteriores.
    
- Dificultad para conocer si un precio se encuentra dentro de valores razonables de mercado.
    
- Riesgo de pérdidas económicas significativas para pequeñas empresas y contratistas.
    

---

## 4. Objetivos

### Objetivo principal

Desarrollar una plataforma que permita gestionar el ciclo completo entre **cotización, ejecución y análisis del costo real**, utilizando la información histórica generada para mejorar futuras estimaciones y construir progresivamente un índice anonimizado de costos y rendimientos reales.

### Objetivos secundarios

- Reducir el tiempo y esfuerzo requerido para elaborar nuevas cotizaciones.
    
- Centralizar las partidas, materiales, mano de obra, rendimientos y precios utilizados por cada empresa.
    
- Mantener un historial de versiones de las cotizaciones.
    
- Evitar el uso de precios desactualizados.
    
- Transformar cotizaciones aceptadas directamente en órdenes de trabajo.
    
- Facilitar el registro de costos y cantidades reales durante la ejecución.
    
- Comparar automáticamente lo cotizado con lo realmente ejecutado.
    
- Permitir conocer desviaciones y rentabilidad real por proyecto y partida.
    
- Facilitar el aprendizaje a partir de proyectos anteriores.
    
- Permitir la configuración de rubros personalizados además de Construcción, Electricidad y Vidriería.
    
- Generar información agregada y anonimizada que permita construir el ÍndicePartidas.
    
- Proteger la información comercial sensible de los usuarios.
    

### No objetivos

En la primera versión del producto no se busca:

- Reemplazar sistemas ERP existentes.
    
- Desarrollar un sistema completo de contabilidad o facturación electrónica.
    
- Administrar remuneraciones o contratos laborales.
    
- Automatizar completamente la cubicación desde planos.
    
- Publicar información identificable sobre costos, clientes o márgenes de otras empresas.
    
- Convertirse inicialmente en un marketplace completo de contratación de servicios.
    
- Garantizar que el valor mostrado por el índice corresponda al precio que un contratista debe cobrar.
    
- Reemplazar el criterio técnico o profesional del usuario al elaborar una cotización.
    

---

## 5. Usuarios objetivo

|Usuario / actor|Necesidad|Nivel de acceso|Frecuencia de uso|
|---|---|---|---|
|Dueño de Pyme / Contratista|Cotizar, controlar costos, conocer rentabilidad y revisar indicadores|Alto|Diaria / semanal|
|Administrador de Obra|Gestionar partidas, presupuestos, ejecución y órdenes de trabajo|Alto|Diaria|
|Jefe de Terreno|Registrar cantidades, avances, costos e incidencias|Medio|Diaria|
|Maestro / trabajador|Registrar principalmente fotografías, avances o incidencias simples|Limitado|Diaria / eventual|
|Contratista independiente|Cotizar, ejecutar y controlar sus propios trabajos|Alto|Diaria / semanal|
|Tasador|Consultar referencias de costos y rendimientos|Consulta|Eventual|
|Aseguradora|Consultar información agregada de costos|Consulta|Eventual|
|Mandante / Inmobiliaria|Consultar referencias y costos estimados de mercado|Consulta|Eventual|
|Administrador del sistema|Gestionar información, validar datos y controlar funcionamiento general|Administrador|Diaria / eventual|

---

## 6. Ficha técnica del producto

|Categoría|Definición|
|---|---|
|Tipo de producto|Plataforma web tipo SaaS|
|Dominio|Cotizaciones, costos, construcción y servicios|
|Plataforma|Web responsiva|
|Dispositivos prioritarios|Computador, tablet y teléfono móvil|
|Backend|Go|
|Frontend|React + JavaScript|
|Base de datos|PostgreSQL|
|Comunicación|API REST mediante JSON|
|Integraciones iniciales|No obligatorias para el MVP|
|Integraciones futuras posibles|Email, proveedores, servicios de almacenamiento, ERP|
|Datos sensibles|Sí: clientes, costos internos, márgenes, datos de empresa y proyectos|
|Seguridad|Autenticación, autorización por roles, anonimización y HTTPS en producción|
|Requisitos regulatorios|Protección de información personal y comercial conforme a normativa aplicable|

---

# 7. Alcance funcional

## Módulos del producto

|Módulo|Propósito|Actor principal|Prioridad|
|---|---|---|---|
|Catálogo de Partidas y APU|Gestionar partidas, materiales, mano de obra, rendimientos y precios|Contratista / Administrador|MVP|
|Cotizaciones y Presupuestos|Crear, modificar, versionar y controlar cotizaciones|Contratista|MVP|
|Flujo Comercial y Órdenes de Trabajo|Controlar aceptación y transformar cotizaciones en trabajos ejecutables|Contratista / Administrador|MVP|
|Captura de Ejecución en Terreno|Registrar cantidades, costos, horas, fotografías e incidencias reales|Jefe de Terreno / Contratista|MVP|
|Comparación y Analítica|Comparar presupuesto con ejecución y calcular indicadores|Dueño / Contratista|MVP|
|ÍndicePartidas|Generar referencias agregadas y anonimizadas de costos y rendimientos|Contratista / Segundo cliente|MVP / Fase 2|
|Gestión de Usuarios y Configuración|Gestionar empresas, usuarios, clientes, proyectos y rubros|Dueño / Administrador|MVP|
|Red de Prestadores|Facilitar la búsqueda voluntaria de contratistas por especialidad|Mandante / Contratista|Futuro|
|Predicción de costos|Estimar desviaciones futuras mediante información histórica|Contratista|Fase 2|

---

## Requerimientos funcionales

|ID|Módulo|Requerimiento|Prioridad|Estado|
|---|---|---|---|---|
|RF01|Catálogo|Visualización del catálogo de partidas|Must|Pendiente|
|RF02|Catálogo|CRUD de partidas y APU|Must|Pendiente|
|RF03|Catálogo|CRUD de categorías de partidas|Must|Pendiente|
|RF04|Catálogo|Actualización de indicadores y precios de insumos|Must|Pendiente|
|RF05|Cotizaciones|CRUD de cotizaciones por partidas|Must|Pendiente|
|RF05.1|Cotizaciones|Configuración de márgenes y gastos generales|Must|Pendiente|
|RF06|Cotizaciones|Control de versiones|Must|Pendiente|
|RF07|Cotizaciones|Partidas extraordinarias|Should|Pendiente|
|RF08|Cotizaciones|Vigencia de precios y cotizaciones|Must|Pendiente|
|RF09|Flujo comercial|Envío y registro de aceptación|Must|Pendiente|
|RF10|Órdenes de trabajo|Conversión automática a Orden de Trabajo|Must|Pendiente|
|RF11|Órdenes de trabajo|Cierre de Orden de Trabajo|Must|Pendiente|
|RF12|Ejecución|Registro simplificado de costos y cantidades|Must|Pendiente|
|RF13|Ejecución|Reporte fotográfico de avances e incidentes|Should|Pendiente|
|RF14|Analítica|Comparación Cotizado vs. Ejecutado|Must|Pendiente|
|RF15|Analítica|Estimación predictiva de desviación|Should|Pendiente|
|RF16|Analítica|Dashboard de rentabilidad|Must|Pendiente|
|RF17|Índice|Consulta del Índice Agregado Anonimizado|Must|Pendiente|
|RF18|Índice|Consentimiento para uso anonimizado|Must|Pendiente|
|RF19|Índice|Validación y depuración de datos|Must|Pendiente|
|RF20|Configuración|Gestión de usuarios, empresas y roles|Must|Pendiente|
|RF21|Configuración|Gestión de clientes y proyectos|Must|Pendiente|
|RF22|Configuración|Configuración de rubros y cotización personalizada|Must|Pendiente|
|RF23|Red de prestadores|Directorio voluntario de contratistas|Could|Pendiente|

---

# Detalle por requerimiento

## RF01 — Visualización del catálogo de partidas

**Descripción:**  
El sistema permitirá consultar las partidas disponibles según el rubro seleccionado, mostrando materiales, mano de obra, rendimientos, unidades, precios e historial asociado.

**Actores:**  
Contratista, Administrador de Obra.

**Precondiciones:**

- El usuario debe estar autenticado.
    
- Debe existir al menos un rubro configurado.
    

**Flujo principal:**

1. El usuario accede al catálogo.
    
2. Selecciona un rubro o categoría.
    
3. El sistema muestra las partidas disponibles y sus antecedentes.
    

**Flujos alternativos / errores:**

- Si no existen partidas, se informará al usuario y se ofrecerá crear una nueva.
    

**Criterio de aceptación BDD:**

- **Dado** un usuario autenticado y un rubro con partidas registradas, **cuando** acceda al catálogo, **entonces** deberá visualizar las partidas y sus principales antecedentes.
    

---

## RF02 — CRUD de Partidas y APU

**Descripción:**  
El usuario podrá crear, consultar, modificar y eliminar partidas, incorporando materiales, mano de obra, unidades, precios y rendimientos.

**Actores:**  
Contratista, Administrador de Obra.

**Precondiciones:**

- Usuario autenticado con permisos de edición.
    

**Flujo principal:**

1. El usuario crea o selecciona una partida.
    
2. Registra sus componentes.
    
3. El sistema calcula y almacena el APU.
    

**Flujos alternativos / errores:**

- No se permitirá guardar una partida sin nombre o unidad obligatoria.
    

**Criterio de aceptación BDD:**

- **Dado** un usuario autorizado, **cuando** complete correctamente los datos de una partida, **entonces** el sistema deberá almacenarla y dejarla disponible para futuras cotizaciones.
    

---

## RF03 — CRUD de Categorías de Partidas

**Descripción:**  
Permitirá organizar las partidas mediante categorías configurables.

**Actores:**  
Contratista, Administrador de Obra.

**Criterio de aceptación BDD:**

- **Dado** un usuario autorizado, **cuando** cree una nueva categoría, **entonces** ésta deberá quedar disponible para asociarle partidas.
    

---

## RF04 — Actualización de precios de insumos

**Descripción:**  
Permitirá actualizar materiales y parámetros utilizados en el cálculo de costos.

**Actores:**  
Contratista, Administrador de Obra.

**Criterio de aceptación BDD:**

- **Dado** un insumo existente, **cuando** el usuario actualice su precio, **entonces** el nuevo valor deberá quedar disponible para futuras cotizaciones sin alterar las versiones históricas ya cerradas.
    

---

## RF05 — CRUD de Cotizaciones por partidas

**Descripción:**  
Permitirá crear presupuestos seleccionando partidas e indicando sus cantidades.

**Actores:**  
Contratista, Administrador de Obra.

**Criterio de aceptación BDD:**

- **Dado** un proyecto y partidas disponibles, **cuando** el usuario ingrese las cantidades, **entonces** el sistema deberá calcular el costo correspondiente y generar la cotización.
    

---

## RF05.1 — Configuración de márgenes y gastos generales

**Descripción:**  
Permitirá incorporar gastos generales, costos indirectos y margen comercial manteniendo esta información como confidencial.

**Actores:**  
Dueño, Contratista.

**Criterio de aceptación BDD:**

- **Dado** un presupuesto calculado, **cuando** un usuario autorizado configure gastos y margen, **entonces** el sistema deberá calcular el precio final sin exponer esta información a roles no autorizados.
    

---

## RF06 — Control de versiones

**Descripción:**  
El sistema conservará las distintas versiones de una cotización y permitirá identificar los cambios realizados.

**Actores:**  
Contratista, Administrador de Obra.

**Criterio de aceptación BDD:**

- **Dado** un presupuesto previamente guardado, **cuando** el usuario genere una nueva versión, **entonces** deberán conservarse la versión anterior y las diferencias entre ambas.
    

---

## RF07 — Partidas extraordinarias

**Descripción:**  
Permitirá incorporar trabajos adicionales sin modificar el presupuesto base aceptado.

**Actores:**  
Administrador de Obra, Contratista.

**Criterio de aceptación BDD:**

- **Dado** un proyecto con presupuesto aceptado, **cuando** se incorpore una modificación, **entonces** deberá registrarse como partida extraordinaria independiente.
    

---

## RF08 — Vigencia de precios y cotizaciones

**Descripción:**  
Cada cotización podrá disponer de una fecha límite de validez.

**Actores:**  
Contratista.

**Criterio de aceptación BDD:**

- **Dado** un presupuesto cuya vigencia haya terminado, **cuando** el usuario lo consulte, **entonces** el sistema deberá advertir que los precios pueden requerir actualización.
    

---

## RF09 — Envío y registro de aceptación

**Descripción:**  
Permitirá exportar o enviar una cotización y registrar su estado comercial.

**Estados mínimos:**

- Borrador.
    
- En revisión.
    
- Aceptada.
    
- Rechazada.
    

**Criterio de aceptación BDD:**

- **Dado** una cotización enviada, **cuando** se registre la aceptación del cliente, **entonces** su estado deberá cambiar a “Aceptada”.
    

---

## RF10 — Conversión automática a Orden de Trabajo

**Descripción:**  
Una cotización aceptada podrá convertirse directamente en una Orden de Trabajo.

**Criterio de aceptación BDD:**

- **Dado** una cotización aceptada, **cuando** el usuario seleccione “Crear Orden de Trabajo”, **entonces** deberán copiarse automáticamente las partidas y cantidades aprobadas.
    

---

## RF11 — Cierre de Orden de Trabajo

**Descripción:**  
Permitirá finalizar formalmente un trabajo y consolidar sus costos reales.

**Criterio de aceptación BDD:**

- **Dado** una Orden de Trabajo en ejecución, **cuando** sea marcada como finalizada, **entonces** el sistema deberá consolidar los datos ejecutados y habilitar el análisis final.
    

---

## RF12 — Registro simplificado de costos y cantidades

**Descripción:**  
Permitirá registrar desde terreno cantidades ejecutadas, materiales, horas de trabajo y costos adicionales.

**Actores:**  
Contratista, Administrador de Obra, Jefe de Terreno.

**Criterio de aceptación BDD:**

- **Dado** una Orden de Trabajo activa, **cuando** el usuario registre la ejecución de una partida, **entonces** los datos deberán quedar asociados a dicha partida y proyecto.
    

---

## RF13 — Reporte fotográfico

**Descripción:**  
Permitirá adjuntar fotografías y notas a partidas o incidencias.

**Actores:**  
Jefe de Terreno, Maestro.

**Criterio de aceptación BDD:**

- **Dado** una partida en ejecución, **cuando** el usuario adjunte una fotografía y comentario, **entonces** éstos deberán quedar vinculados al proyecto y partida correspondiente.
    

---

## RF14 — Comparación Cotizado vs. Ejecutado

**Descripción:**  
Permitirá visualizar costos presupuestados, costos reales y desviaciones.

**Criterio de aceptación BDD:**

- **Dado** una partida con datos presupuestados y ejecutados, **cuando** el usuario consulte su análisis, **entonces** el sistema deberá mostrar ambos valores y su diferencia monetaria y porcentual.
    

---

## RF15 — Estimación predictiva de desviación

**Descripción:**  
El sistema utilizará información histórica propia y agregada para estimar el costo real esperado de una partida.

**Criterio de aceptación BDD:**

- **Dado** suficiente historial de datos comparables, **cuando** el usuario prepare una nueva partida, **entonces** el sistema podrá mostrar una estimación de costo esperado o riesgo de desviación.
    

---

## RF16 — Dashboard de rentabilidad

**Descripción:**  
Permitirá visualizar indicadores como tasa de adjudicación, desviación promedio y margen real.

**Criterio de aceptación BDD:**

- **Dado** un usuario con proyectos históricos, **cuando** acceda al dashboard, **entonces** deberá visualizar indicadores calculados a partir de sus propios datos.
    

---

## RF17 — Índice Agregado Anonimizado

**Descripción:**  
Permitirá consultar referencias agregadas de costos y rendimientos provenientes de trabajos ejecutados.

**Criterio de aceptación BDD:**

- **Dado** que existan suficientes registros válidos, **cuando** un usuario consulte una partida del índice, **entonces** deberá visualizar estadísticas agregadas sin identificar empresas individuales.
    

---

## RF18 — Consentimiento para uso anonimizado

**Descripción:**  
Los contratistas podrán aceptar o rechazar la utilización anonimizada de sus datos para construir el índice.

**Criterio de aceptación BDD:**

- **Dado** un usuario propietario de sus datos, **cuando** configure su consentimiento, **entonces** el sistema deberá respetar esa selección al generar información agregada.
    

---

## RF19 — Validación y depuración de datos

**Descripción:**  
Los registros inconsistentes, anómalos o incompletos deberán ser detectados antes de incorporarse al índice.

**Criterio de aceptación BDD:**

- **Dado** un registro que incumpla los criterios definidos de validación, **cuando** sea procesado para el índice, **entonces** deberá excluirse o quedar marcado para revisión.
    

---

## RF20 — Gestión de usuarios, empresas y roles

**Descripción:**  
Permitirá crear usuarios y definir permisos según su responsabilidad dentro de la organización.

**Criterio de aceptación BDD:**

- **Dado** un usuario con rol limitado, **cuando** intente acceder a información restringida, **entonces** el sistema deberá impedir el acceso.
    

---

## RF21 — Gestión de clientes y proyectos

**Descripción:**  
Permitirá almacenar clientes y proyectos asociados a cotizaciones y órdenes de trabajo.

**Criterio de aceptación BDD:**

- **Dado** un cliente registrado, **cuando** se cree una cotización, **entonces** ésta podrá quedar vinculada al cliente y proyecto correspondiente.
    

---

## RF22 — Configuración de rubros y estructuras personalizadas

**Descripción:**  
El usuario podrá utilizar Construcción, Electricidad o Vidriería, o crear un rubro mediante la opción “Otro”.

En un rubro personalizado podrá definir:

- Nombre del rubro.
    
- Categorías.
    
- Conceptos utilizados en la cotización.
    
- Unidades.
    
- Variables.
    
- Precios.
    
- Fórmulas.
    
- Parámetros de cálculo.
    

La configuración deberá almacenarse y reutilizarse posteriormente.

**Criterio de aceptación BDD:**

- **Dado** un usuario que seleccione “Otro”, **cuando** configure y guarde su estructura de cotización, **entonces** ésta deberá quedar disponible para crear futuras cotizaciones sin modificar el código fuente del sistema.
    

---

## RF23 — Directorio voluntario de contratistas

**Descripción:**  
Permitirá que contratistas que lo autoricen aparezcan en un directorio por especialidad y zona.

**Criterio de aceptación BDD:**

- **Dado** un contratista que haya autorizado la publicación de su perfil, **cuando** un usuario consulte prestadores de esa especialidad, **entonces** podrá visualizar sus datos públicos de contacto.
    

---

# 8. Requerimientos no funcionales

|ID|Categoría|Requerimiento verificable|Métrica / criterio|
|---|---|---|---|
|RNF01|Usabilidad|Interfaz comprensible para usuarios sin conocimientos técnicos especializados|Prueba de las funciones principales con usuarios objetivo|
|RNF02|Usabilidad|Registro rápido en terreno|Operación habitual ≤ 30 segundos|
|RNF03|Usabilidad|Diseño responsivo|Visualización funcional en desktop, tablet y móvil|
|RNF04|Usabilidad|Consistencia visual|Mismos patrones de navegación, botones y formularios|
|RNF05|Rendimiento|Tiempo de respuesta en operaciones habituales|≤ 3 segundos bajo condiciones normales|
|RNF06|Rendimiento|Generación y visualización de cotización|≤ 5 segundos|
|RNF07|Rendimiento|Consultas optimizadas|Uso adecuado de consultas e índices PostgreSQL|
|RNF08|Seguridad|Autenticación obligatoria para información privada|Acceso rechazado a usuario no autenticado|
|RNF09|Seguridad|Autorización mediante roles|Pruebas de permisos por rol|
|RNF10|Seguridad|Contraseñas almacenadas mediante hash seguro|Ninguna contraseña almacenada en texto plano|
|RNF11|Seguridad|Confidencialidad comercial|Empresas no pueden visualizar información privada de terceros|
|RNF12|Seguridad|Anonimización de información del índice|No debe identificarse directamente al proveedor del dato|
|RNF13|Seguridad|Comunicación segura en producción|HTTPS|
|RNF14|Datos|PostgreSQL como motor principal|Persistencia implementada en PostgreSQL|
|RNF15|Datos|Integridad referencial|PK, FK y restricciones donde corresponda|
|RNF16|Datos|Persistencia histórica|Versiones y ejecuciones históricas no deben sobrescribirse|
|RNF17|Datos|Uso de transacciones|Operaciones críticas consistentes ante errores|
|RNF18|Datos|Respaldo y recuperación|Mecanismo de copia y recuperación disponible|
|RNF19|Tecnología|Go|Proyecto ejecutable en el stack definido|
|RNF20|Tecnología|Vue|Aplicación web implementada con el stack definido|
|RNF21|Arquitectura|API REST con JSON|Comunicación cliente-servidor mediante endpoints REST|
|RNF22|Compatibilidad|Compatibilidad web|Chrome, Edge y Firefox modernos|
|RNF23|Arquitectura|Separación cliente-servidor|Frontend y backend desplegables/modificables independientemente|
|RNF24|Extensibilidad|Incorporación de nuevos rubros|No debe requerir rediseñar el núcleo|
|RNF25|Extensibilidad|Configuración dinámica de “Otro”|Configuración almacenada como datos y no como código|
|RNF26|Escalabilidad|Crecimiento del historial|Soportar aumento progresivo de partidas y ejecuciones|
|RNF27|Mantenibilidad|Arquitectura modular|Separación de responsabilidades|
|RNF28|Mantenibilidad|Código mantenible|Convenciones, nombres descriptivos y separación por capas|
|RNF29|Mantenibilidad|Control de versiones|Repositorio Git|
|RNF30|Mantenibilidad|Documentación de API|Endpoints principales documentados mediante OpenAPI/Swagger|

---

# 9. Priorización del MVP

## Incluido en MVP

- Gestión de usuarios y roles.
    
- Gestión de empresas.
    
- Gestión de clientes y proyectos.
    
- Rubros predefinidos:
    
    - Construcción.
        
    - Electricidad.
        
    - Vidriería.
        
- Creación de rubros personalizados mediante “Otro”.
    
- Catálogo de partidas.
    
- APU.
    
- Creación de cotizaciones.
    
- Gastos generales y márgenes.
    
- Control de versiones.
    
- Vigencia de cotizaciones.
    
- Registro de aceptación o rechazo.
    
- Conversión de cotización en Orden de Trabajo.
    
- Registro simplificado de ejecución.
    
- Cierre de Orden de Trabajo.
    
- Comparación Cotizado vs. Ejecutado.
    
- Dashboard básico de indicadores.
    
- Consentimiento para uso de información anonimizada.
    
- Prototipo inicial del ÍndicePartidas.
    

## Diferido / Fase 2

- Predicción avanzada del costo real mediante aprendizaje automático.
    
- Modelos de detección automática de anomalías.
    
- Índice con mayor segmentación territorial y temporal.
    
- Automatización avanzada de actualización de precios.
    
- Integraciones con ERP u otros sistemas externos.
    

## Futuro

- Directorio avanzado de prestadores.
    
- Solicitud directa de cotizaciones a contratistas.
    
- Calificación de prestadores.
    
- Integración con proveedores de materiales.
    
- Aplicación móvil nativa.
    
- Integraciones financieras o contables.
    

---

# 10. Métricas de éxito

> Las siguientes metas corresponden inicialmente a **metas propuestas del proyecto** y deberán validarse mediante pruebas con usuarios.

|Métrica|Línea base|Meta propuesta|Cómo se mide|
|---|---|---|---|
|Tiempo para realizar una cotización|Entre 30 min y varias horas según entrevistados|Reducir el tiempo respecto del proceso manual del mismo usuario|Comparación antes/después en prueba de usuario|
|Registro de ejecución|Actualmente inexistente o manual en varios entrevistados|≥ 80 % de partidas del piloto con información de ejecución registrada|Registros del sistema|
|Tiempo de registro en terreno|No estandarizado|≤ 30 segundos para registro habitual|Prueba cronometrada|
|Comparación Cotizado vs. Ejecutado|En varios entrevistados no se realiza|100 % de OT cerradas con comparación automática|Base de datos|
|Conocimiento del margen real|Algunos usuarios declaran desconocerlo|Mostrar margen real en 100 % de proyectos con datos completos|Dashboard|
|Trazabilidad de cotizaciones|Archivos/cuadernos dispersos|Mantener historial completo de versiones creadas en plataforma|Registro de versiones|
|Datos utilizables para el índice|Actualmente inexistentes en plataforma|Obtener registros válidos y anonimizados de usuarios participantes en el piloto|Base de datos|
|Satisfacción de usabilidad|Sin línea base|Validación positiva del flujo principal por usuarios piloto|Encuesta / prueba de usuario|

---

# 11. Riesgos y decisiones

|Riesgo / decisión|Impacto|Mitigación / resolución|Estado|
|---|---|---|---|
|Usuarios no registran costos reales por falta de tiempo|Alto|Simplificar el registro, reducir campos y diseñarlo para uso móvil|Abierto|
|Información ingresada no es confiable|Alto|Validación de rangos, detección de anomalías y exclusión de datos inconsistentes|Abierto|
|Usuarios temen compartir información comercial|Alto|Consentimiento explícito y anonimización estricta|Abierto|
|El índice contiene pocos datos inicialmente|Alto|Comenzar con datos propios de usuarios piloto y aumentar progresivamente la muestra|Abierto|
|Rubros personalizados generan estructuras demasiado diferentes|Medio|Utilizar un modelo configurable basado en conceptos, unidades, variables y fórmulas|Abierto|
|Precios cambian rápidamente|Alto|Vigencia de cotizaciones y registro histórico de precios|Abierto|
|Usuarios de terreno presentan baja adopción tecnológica|Alto|Interfaz responsiva, simple y similar a acciones cotidianas de una aplicación móvil|Abierto|
|Desarrollo de ML requiere más datos de los disponibles inicialmente|Alto|Utilizar primero estadísticas e históricos y evolucionar hacia ML al disponer de datos suficientes|Abierto|
|Construcción, Electricidad y Vidriería serán rubros iniciales|Medio|Desarrollar catálogos iniciales y mantener arquitectura extensible|Aprobada|
|Se incorpora opción “Otro” configurable|Alto|Guardar la configuración como datos para evitar programación específica por rubro|Aprobada|
|PostgreSQL como base de datos|Técnico|Base relacional adecuada para trazabilidad y relaciones del sistema|Aprobada|

---

# 12. Dependencias

## Dependencias técnicas

- Go.
- React.
- JavaScript
- PostgreSQL.
- API REST.
- JSON.
- Git.
- OpenAPI / Swagger para documentación de API.
    

## Dependencias de negocio

- Participación de contratistas y pequeñas empresas para validar el producto.
    
- Disponibilidad de datos reales o simulados de partidas, cotizaciones y ejecuciones.
    
- Disposición de los usuarios a registrar costos ejecutados.
    
- Consentimiento para utilizar información anonimizada.
    
- Definición de catálogos iniciales para Construcción, Electricidad y Vidriería.
    

## Dependencias externas

- Navegador web moderno.
    
- Conexión a Internet para la versión web desplegada.
    
- Servicio de alojamiento para backend, frontend y base de datos en ambiente productivo.
    
- Eventual almacenamiento externo para fotografías y archivos.
    
- Eventuales fuentes externas de precios o indicadores en futuras versiones.
    

---

# 13. Checklist de completitud

- [x] El problema está claramente definido.
- [x] Los usuarios objetivo están identificados.
- [x] El MVP está separado de fases futuras.
- [x] Cada RF tiene criterio de aceptación verificable.
- [x] Cada RNF tiene forma de medición.
- [x] Las decisiones relevantes están registradas.
- [x] Los riesgos principales tienen mitigación.
- [ ] Validar las metas propuestas de las métricas de éxito con usuarios.
- [ ] Definir definitivamente quién será responsable de producto.
- [ ] Confirmar el stack frontend con todo el equipo.
- [ ] Definir cuáles funcionalidades de aprendizaje automático deberán estar operativas en la entrega final del semestre.
