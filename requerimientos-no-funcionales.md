# Requerimientos No Funcionales (RNF) — ÍndicePartidas (Desafío 4)

> **Proyecto-Innova** — Grupo 1, Innovación y Emprendimiento 2-2026, USACH.

### 1. Usabilidad

|N°|Requerimiento no funcional|Descripción / Criterio|
|---|---|---|
|**RNF-01**|Interfaz intuitiva|El sistema deberá permitir ejecutar las funciones principales sin requerir conocimientos técnicos especializados.|
|**RNF-02**|Registro rápido en terreno|El registro de ejecución de una partida deberá requerir la menor cantidad posible de campos y permitir completar una operación habitual en menos de 30 segundos.|
|**RNF-03**|Diseño responsivo|La interfaz deberá adaptarse correctamente a computadores, tablets y teléfonos móviles.|
|**RNF-04**|Consistencia visual|Formularios, botones, tablas y navegación deberán mantener un diseño y comportamiento uniforme en todos los módulos.|

El RNF-02 es especialmente importante porque as entrevistas mostraron que el gran riesgo es que los usuarios **no registren los costos por ser demasiado lento o complejo**.

### 2. Rendimiento

|N°|Requerimiento no funcional|Descripción / Criterio|
|---|---|---|
|**RNF-05**|Tiempo de respuesta|Las operaciones habituales, como consultar partidas, cotizaciones o proyectos, deberán responder en un máximo aproximado de 3 segundos bajo condiciones normales.|
|**RNF-06**|Generación de cotización|El cálculo y visualización de una cotización deberá completarse en un máximo aproximado de 5 segundos.|
|**RNF-07**|Consultas optimizadas|Las búsquedas y filtros deberán utilizar mecanismos eficientes de consulta e índices en PostgreSQL cuando corresponda.|

### 3. Seguridad

|N°|Requerimiento no funcional|Descripción / Criterio|
|---|---|---|
|**RNF-08**|Autenticación|Todo usuario deberá autenticarse antes de acceder a información privada de la plataforma.|
|**RNF-09**|Autorización por roles|El acceso a funciones e información deberá restringirse según el rol del usuario.|
|**RNF-10**|Protección de contraseñas|Las contraseñas deberán almacenarse mediante algoritmos seguros de hash y nunca en texto plano.|
|**RNF-11**|Confidencialidad comercial|Márgenes, clientes, costos privados e información identificable de una empresa no deberán ser visibles para otros usuarios.|
|**RNF-12**|Anonimización|Los datos utilizados en ÍndicePartidas deberán ser procesados de forma que el índice agregado no permita identificar directamente a la empresa que los aportó.|
|**RNF-13**|Comunicación segura|En producción, la comunicación entre cliente y servidor deberá utilizar HTTPS.|

Spring Security 

### 4. Base de datos e integridad

|N°|Requerimiento no funcional|Descripción / Criterio|
|---|---|---|
|**RNF-14**|Motor de base de datos|La persistencia principal deberá implementarse mediante PostgreSQL.|
|**RNF-15**|Integridad referencial|La base de datos deberá utilizar claves primarias, claves foráneas y restricciones para evitar registros inconsistentes.|
|**RNF-16**|Persistencia histórica|Las cotizaciones aceptadas, versiones y costos ejecutados relevantes no deberán perderse al actualizar información posterior.|
|**RNF-17**|Transacciones|Las operaciones críticas deberán utilizar transacciones para evitar estados incompletos ante errores.|
|**RNF-18**|Respaldo|La base de datos deberá permitir la generación periódica de copias de seguridad y su recuperación.|

### 5. Compatibilidad y tecnología

|N°|Requerimiento no funcional|Descripción / Criterio|
|---|---|---|
|**RNF-19**|Backend|El backend será desarrollado en Go, utilizando una arquitectura modular para la implementación de la lógica de negocio y servicios del sistema.|
|**RNF-20**|Frontend|La interfaz web será desarrollada utilizando React y JavaScript para construir una interfaz web responsiva e interactiva|
|**RNF-21**|Arquitectura de comunicación|Frontend y backend deberán comunicarse mediante una API REST utilizando JSON.|
|**RNF-22**|Compatibilidad web|El sistema deberá funcionar correctamente en versiones modernas de Chrome, Edge y Firefox.|
|**RNF-23**|Independencia cliente-servidor|El frontend y backend deberán mantenerse desacoplados para permitir modificaciones independientes.|

### 6. Escalabilidad y extensibilidad

|N°|Requerimiento no funcional|Descripción / Criterio|
|---|---|---|
|**RNF-24**|Extensibilidad de rubros|La arquitectura deberá permitir incorporar nuevos rubros y estructuras de cotización sin rediseñar el núcleo del sistema.|
|**RNF-25**|Configuración dinámica|Las estructuras creadas mediante “Otro” deberán almacenarse como configuraciones y no mediante cambios directos al código fuente.|
|**RNF-26**|Escalabilidad de datos|La base de datos deberá permitir el crecimiento progresivo del historial de partidas, ejecuciones y costos utilizados para generar el índice.|
### 7. Mantenibilidad

|N°|Requerimiento no funcional|Descripción / Criterio|
|---|---|---|
|**RNF-27**|Arquitectura modular|El sistema deberá separar responsabilidades entre frontend, backend y persistencia de datos.|
|**RNF-28**|Código mantenible|El código deberá organizarse siguiendo convenciones consistentes, nombres descriptivos y separación por capas o módulos.|
|**RNF-29**|Control de versiones|El código fuente deberá mantenerse mediante un sistema de control de versiones, preferentemente Git.|
|**RNF-30**|Documentación de API|Los endpoints principales de la API deberán estar documentados para facilitar mantenimiento e integración.|

Para el RNF-30 --> usar después **Swagger / OpenAPI**.
