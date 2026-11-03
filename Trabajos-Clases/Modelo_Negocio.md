# Modelo de negocio — Grupo 1 — Desafío: ÍndicePartidas

## Tarjeta de aprendizaje (cierre del experimento)

- **Creíamos que:** si el registro de ejecución requería menos de 30 segundos y solicitaba solamente información esencial, los contratistas estarían dispuestos a registrar datos reales de sus trabajos de manera recurrente. [H]
    
- **Para verificarlo hicimos:** un experimento mediante un formulario móvil simplificado que simula el futuro registro de ejecución de ÍndicePartidas. Se solicitó registrar información básica de trabajos o partidas: cantidad ejecutada, materiales o costo, horas de trabajo y observaciones.
    
- **El criterio era:** al menos **3 de 5 participantes** debían realizar por cuenta propia un mínimo de **3 registros reales o reconstruidos de trabajos distintos**, y al menos **80 % de los registros** debían completarse en **30 segundos o menos**, antes del 3 de noviembre de 2026. [H definido antes del experimento]
    
- **Observamos:** `[Completar con el resultado real del experimento: cantidad de participantes, cantidad de registros, tiempos obtenidos y personas que no respondieron]`.
    
- **De eso aprendimos que:** `[Completar después de revisar los resultados reales]`.
    
- **Lo que todavía no sabemos:** si el comportamiento observado durante el experimento se mantendrá durante varias semanas de uso real y si los contratistas seguirán registrando información una vez terminada la etapa inicial de prueba.
    

---

## Decisión

- **Perseverar:** el criterio se cumplió o casi, y lo aprendido no cambia la pieza central.
    
- **Redefinir:** cambia una pieza del modelo.
    
- **Concluir que no hay negocio:** el problema existe, pero la evidencia muestra que el modelo no funciona.
    

**Por qué:**

`[Completar una vez cerrado el experimento.]`

Si el criterio del experimento se cumple o queda muy cerca, la decisión recomendada sería **perseverar**, ya que las entrevistas ya validaron el problema de no conocer el costo real y la principal incertidumbre era la disposición a registrar la ejecución.

Si el registro resulta demasiado lento o los usuarios no lo realizan, corresponde **redefinir el mecanismo de captura de datos**, no necesariamente abandonar el problema identificado.

---

# Business Model Canvas

|Bloque|Contenido|Fuente|
|---|---|---|
|**Segmentos de clientes**|**Cliente inicial:** contratistas independientes, dueños de pequeñas empresas y pymes que elaboran cotizaciones y ejecutan trabajos por especialidad. Los rubros iniciales serán Construcción, Electricidad y Vidriería. **Segundo cliente futuro:** contratistas generales, tasadores, aseguradoras, inmobiliarias y mandantes interesados en información agregada de costos y rendimientos.|E-01, E-02, E-03, E-05|
|**Propuesta de valor**|Para contratistas independientes y pequeñas empresas que necesitan preparar presupuestos rentables y controlar sus trabajos, **ÍndicePartidas** es una plataforma de cotización y control de costos que conecta lo presupuestado con lo realmente ejecutado para mostrar desviaciones, rentabilidad y utilizar la experiencia anterior en futuras cotizaciones. A diferencia de trabajar únicamente con cuadernos, planillas o presupuestos antiguos, cierra el ciclo **cotizar → ejecutar → comparar → aprender**.|E-01, E-02, E-04, E-05|
|**Canales**|En la etapa inicial, el principal canal será la **venta directa**, mostrando el prototipo y posteriormente el producto a contratistas y pequeñas empresas. También se consideran **referidos entre contratistas**, presencia digital mediante sitio web y contenido, y en una etapa posterior gremios o asociaciones de especialidades. [H]|Venta directa: [H], Referidos/gremios: [H]|
|**Relaciones con clientes**|Modelo principalmente de **autoservicio asistido**. El usuario podrá registrarse, configurar su empresa y crear cotizaciones por sí mismo, pero durante la incorporación inicial se contempla apoyo para cargar partidas y configurar el rubro. El soporte deberá ser sencillo y rápido, especialmente durante los primeros usos. Para mantener al cliente se buscará que obtenga valor temprano mediante su primera cotización y su primera comparación Cotizado vs. Ejecutado.|E-02, E-05 + [H]|
|**Fuentes de ingreso**|**Primera línea:** suscripción SaaS mensual por tramos según cantidad de cotizaciones/trabajos gestionados. Paquete Independiente: **$9.990 + IVA/mes [H]**; Pyme: **$24.990 + IVA/mes [H]**; Empresa: **$59.990 + IVA/mes [H]**. **Segunda línea futura:** suscripción profesional por acceso al ÍndicePartidas agregado, reportes y análisis de mercado. Precio todavía no validado. [H]|E-03 + pricing.md + [H]|
|**Recursos clave**|Plataforma web, base de datos PostgreSQL, backend en Go, frontend en Vue, infraestructura de alojamiento, catálogo inicial de partidas, información histórica de cotizaciones y ejecuciones, datos anonimizados aportados por usuarios con consentimiento y equipo de desarrollo. El recurso estratégico a largo plazo es la base acumulada de costos y rendimientos reales.|D4 + RF/RNF + E-01, E-02, E-05|
|**Actividades clave**|Desarrollar y mantener la plataforma; facilitar la creación de cotizaciones; mantener y mejorar estructuras de partidas; permitir el registro de ejecución; comparar Cotizado vs. Ejecutado; validar y depurar datos; anonimizar la información destinada al índice; dar soporte a usuarios; captar nuevos contratistas y mantener activa la participación de los existentes.|problema-solucion.md + RF|
|**Socios clave**|Proveedor de infraestructura/cloud [H]; proveedor de dominio y servicios web [H]; eventualmente una pasarela de pago para suscripciones [H]; gremios o asociaciones de contratistas como canal de llegada [H]; en el futuro, fuentes externas de precios o proveedores de materiales si se decide automatizar actualización de insumos [H]. No se requiere socio financiero para el modelo inicial.|[H]|
|**Estructura de costos**|**Fijos:** desarrollo y mantenimiento del software, alojamiento base, dominio, herramientas del equipo y administración. **Variables:** almacenamiento adicional, tráfico, soporte por cliente, procesamiento de archivos/fotografías y comisiones de pasarela de pago. **Adquisición de clientes:** tiempo destinado a contacto directo, demostraciones, incorporación inicial y soporte. Los montos se calcularán posteriormente en `unit-economics.md`.|[H]|

---

## ¿Cuenta una sola historia?

Sí, de manera preliminar.

El modelo comienza con un segmento que ya fue identificado en las entrevistas: **contratistas pequeños que cotizan y no siempre conocen el resultado real de sus trabajos**. [E-01, E-02]

La propuesta de valor responde directamente a ese problema mediante una plataforma que conecta cotización y ejecución.

El cliente inicial pagaría mediante una **suscripción mensual**, coherente con el modelo SaaS y con la preferencia expresada en E-03 por una suscripción mensual con reportes.

La principal incertidumbre está en los canales y en la adquisición del cliente: el precio del paquete principal es de **$24.990 + IVA mensuales [H]**, por lo que una venta directa demasiado costosa no sería sostenible si requiere muchas horas por cliente.

Por esta razón, la venta directa se utilizaría inicialmente para validar el modelo, pero a largo plazo deberán probarse canales más escalables como referidos, asociaciones y presencia digital. [H]

Otra incertidumbre importante es la generación del ÍndicePartidas. El valor futuro del dato depende de que suficientes usuarios registren costos reales de manera constante y autoricen su utilización anonimizada.

---

# Las tres hipótesis que faltan

|Hipótesis (bloque)|Cómo la van a probar|Cuándo|
|---|---|---|
|**1. Relaciones / Actividades clave:** los contratistas registrarán de manera recurrente sus costos reales si el proceso es suficientemente simple y rápido. [H]|Experimento de registro simplificado: medir cantidad de registros, recurrencia y tiempo utilizado por cada participante.|Experimento actual / inicio de Fase 3|
|**2. Fuentes de ingreso:** pequeños contratistas y pymes estarán dispuestos a asumir una suscripción cercana a **$24.990 + IVA mensuales [H]** por el paquete principal.|Presentar una propuesta con precio concreto y medir una acción real: firma de intención, solicitud de piloto o pago.|Antes de la PEP 1 y continuar en Fase 3|
|**3. Canales:** es posible captar clientes a un costo compatible con una suscripción mensual relativamente baja. [H]|Comparar venta directa, referidos y contacto mediante asociaciones; medir cuántos contactos se necesitan para conseguir una demostración y posteriormente un cliente.|Fase 3|

---

## Lo que cambió respecto del desafío del catálogo

Inicialmente el desafío planteaba principalmente un software de cotización como mecanismo para obtener datos de costos reales y construir un índice sectorial.

Las entrevistas mostraron que el problema inmediato para los contratistas pequeños no es solamente crear una cotización, sino **no saber con precisión cuánto ganaron o perdieron después de ejecutar el trabajo**, además de la dificultad de registrar información adicional en terreno. [E-01, E-02, E-05]

Por ello, el proyecto mantiene la tesis original del **ÍndicePartidas**, pero prioriza primero entregar valor al contratista mediante el ciclo **Cotizado vs. Ejecutado**. La generación del índice ocurre como consecuencia de acumular esos datos de forma estructurada, validada y anonimizada, en lugar de exigir que el usuario registre información únicamente para alimentar una base de datos externa.