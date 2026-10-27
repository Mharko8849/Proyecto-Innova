# Problema y solución — Grupo 1 — Desafío: ÍndicePartidas

## El problema

**Enunciado:**

> Los contratistas independientes y pequeñas empresas de especialidades necesitan preparar cotizaciones que cubran correctamente sus costos y les permitan obtener una utilidad, pero actualmente muchos trabajan con cuadernos, planillas o presupuestos anteriores y no comparan sistemáticamente lo cotizado con lo realmente ejecutado. Esta situación aparece cada vez que preparan y posteriormente ejecutan un trabajo, pudiendo requerir desde 30 minutos hasta varias horas para cotizar y generando desviaciones informadas de entre 2 % y 30 % según el tipo de proyecto. Como resultado, algunos contratistas descubren los errores cuando ya compraron materiales o terminaron el trabajo, sin conocer con certeza su margen real. [E-01, E-02, E-04, E-05]

- **Quién lo tiene:** contratistas independientes, maestros contratistas y pequeñas empresas que realizan trabajos por especialidad, particularmente en construcción, electricidad, vidriería y servicios similares. [E-01, E-02, E-05]
    
- **Qué quiere lograr, y qué pasa hoy en vez de eso:** quieren preparar presupuestos que reflejen correctamente sus costos y conocer si el trabajo fue realmente rentable, pero hoy muchos estiman antes de comenzar y no realizan una comparación estructurada al finalizar. [E-01, E-02]
    
- **Cuándo le pasa, y cuántas veces:** cada vez que deben preparar una nueva cotización y posteriormente ejecutar el trabajo. E-01 indicó que cada cotización puede tomar entre 30 minutos y 2 horas, mientras E-02 dedica sesiones de 2 a 3 horas a presupuestar. [E-01, E-02]
    
- **Cuánto le cuesta:** errores de cubicación, variaciones de materiales y mayor mano de obra pueden generar desviaciones. E-01 reportó desviaciones de 20 % a 30 % en trabajos complejos, mientras E-05 indicó desviaciones habituales de 2 % a 10 % en proyectos de construcción. [E-01, E-05]
    
- **Evidencia:** [E-01, E-02, E-04, E-05]
    

### Reencuadre

Una alternativa considerada fue definir el problema únicamente como **“los contratistas tardan demasiado en realizar cotizaciones”**. Sin embargo, las entrevistas mostraron que la velocidad no es el problema más crítico: E-04 indicó que una mala estimación puede afectar gravemente la sostenibilidad financiera de una pyme y E-01 y E-02 mostraron que el problema continúa después de cotizar, porque no existe una comparación sistemática con el costo real.

Por eso el equipo decidió enfocarse en un problema más amplio: **la falta de trazabilidad entre lo que se presupuesta y lo que realmente cuesta ejecutar el trabajo**, ya que esta falta de información provoca que los mismos errores puedan repetirse en futuras cotizaciones. [E-01, E-02, E-04]

---

## Alternativas de hoy

Actualmente los usuarios resuelven el problema mediante:

- Cuadernos y registros manuales para guardar precios y preparar presupuestos. [E-02]
    
- Microsoft Excel y planillas propias para APU, cubicaciones y presupuestos. [E-05]
    
- Reutilización de presupuestos y datos históricos de trabajos anteriores. [E-04, E-05]
    
- Comparación manual de presupuestos solicitando precios a diferentes maestros o especialistas. [E-03]
    
- Estimación basada en experiencia personal, sin realizar posteriormente un análisis formal de rentabilidad. [E-01, E-02]
    
- Software especializado de construcción, principalmente disponible para empresas con procesos más estructurados. E-02 señaló haber evaluado software, pero desistió por su costo. [E-02]
    

---

## Hipótesis de solución

Antes de elegir la solución se consideraron tres alternativas:

1. **Un cotizador digital simple**, que permita crear presupuestos mediante partidas, materiales, mano de obra y márgenes, guardando plantillas para reutilizarlas. 
    
2. **Una herramienta de control posterior al trabajo**, donde el usuario registre los materiales, cantidades, horas y costos realmente utilizados para conocer automáticamente la rentabilidad y desviación del proyecto. 
    
3. **Una plataforma integrada de cotización y aprendizaje histórico**, que conecte la creación del presupuesto con la ejecución real, compare ambos resultados y reutilice los datos históricos para mejorar las siguientes cotizaciones. 
    

### La que elegimos, y por qué

Se eligió la tercera alternativa porque responde al problema completo identificado en las entrevistas.

Un cotizador por sí solo permite ahorrar tiempo, pero no evita que el usuario vuelva a cometer los mismos errores si después no sabe cuánto costó realmente el trabajo. Del mismo modo, registrar solamente el costo final no soluciona el problema de preparar mejores presupuestos futuros.

ÍndicePartidas conectará ambos procesos:

**Cotizar → Ejecutar → Registrar → Comparar → Aprender → Volver a cotizar.**

**Creemos que** construyendo una plataforma simple que conecte la cotización con el registro del costo real **para** contratistas independientes y pequeñas empresas de especialidades **lograremos que** puedan conocer sus desviaciones y utilizar sus propios trabajos anteriores como referencia para preparar nuevas cotizaciones. **Lo sabremos cuando** al menos 3 de 5 usuarios del experimento registren por cuenta propia información real o reconstruida de al menos 3 trabajos y la mayoría de los registros pueda realizarse en aproximadamente 30 segundos o menos. [H]

Este criterio coincide con el supuesto principal definido en `experimento.md`: el sistema solo generará valor si los usuarios efectivamente registran la ejecución real.

---

## Alcance del MVP

|Problema (ordenado)|Funcionalidad que responde|Cajón|
|---|---|---|
|1. El contratista prepara presupuestos utilizando información dispersa o histórica y puede equivocarse en cantidades, costos o márgenes. [E-01, E-02, E-04]|Crear cotizaciones mediante partidas, cantidades, costos, mano de obra y margen, guardando la información para reutilizarla.|Imprescindible|
|2. Después de ejecutar un trabajo, muchos contratistas no registran sistemáticamente cuánto gastaron realmente. [E-01, E-02]|Registro simplificado del costo real ejecutado por partida: cantidad, materiales, horas y costos adicionales.|Imprescindible|
|3. El contratista no conoce claramente cuánto se desvió ni cuál fue su rentabilidad real. [E-01, E-02, E-05]|Comparación automática Cotizado vs. Ejecutado, mostrando diferencia en monto y porcentaje.|Imprescindible|

### Después

Funcionalidades que agregan valor pero no son necesarias para demostrar la propuesta principal del MVP:

- Dashboard completo de rentabilidad e indicadores históricos.
    
- Control avanzado de versiones de cotizaciones.
    
- Partidas extraordinarias.
    
- Reporte fotográfico de avances.
    
- Vigencia automática de precios.
    
- Índice agregado anonimizado.
    
- Validación y depuración automática de datos para el índice.
    
- Estimación predictiva mediante aprendizaje automático.
    
- Directorio voluntario de contratistas.
    
- Automatización avanzada de precios.
    

### Nunca en este MVP

Estas funcionalidades pueden formar parte de una evolución futura, pero no se construirán durante el MVP de cuatro semanas:

- Marketplace completo para contratar prestadores.
    
- Sistema de reputación y calificaciones.
    
- Facturación electrónica.
    
- Contabilidad completa.
    
- Gestión de remuneraciones.
    
- Lectura automática de planos y cubicación mediante inteligencia artificial.
    
- Integración completa con ERP.
    
- Aplicación móvil nativa independiente.
    
- Modelo avanzado de machine learning entrenado con grandes volúmenes de información.
    

---

### Lo que tiene que ser verdad técnicamente

- La aplicación debe permitir almacenar cotizaciones, partidas, cantidades y costos históricos en una base de datos relacional. — **Sí**.
    
- Debe existir una relación entre una cotización aceptada y los datos reales registrados posteriormente durante la ejecución. — **Sí**.
    
- El sistema debe poder calcular automáticamente diferencias monetarias y porcentuales entre presupuesto y ejecución. — **Sí**.
    
- El registro desde terreno debe funcionar correctamente desde un teléfono mediante una interfaz web responsiva. — **No comprobado todavía / [H]**.
    
- La arquitectura debe permitir guardar diferentes estructuras de cotización para los rubros predefinidos y para la opción “Otro”. — **No comprobado todavía / [H]**.
    
- Los datos que eventualmente alimenten el ÍndicePartidas deben poder anonimizarse antes de utilizarse de forma agregada. — **No comprobado todavía / [H]**.
    
- El backend será desarrollado en Go, el frontend en Vue y la persistencia principal en PostgreSQL. — **Definición técnica del proyecto**.
    

---

## Mapa de historias

|Paso 1|Paso 2|Paso 3|Paso 4|
|---|---|---|---|
|**Preparar cotización**|**Aceptar y ejecutar trabajo**|**Registrar costo real**|**Revisar resultado**|
|Crear cliente/proyecto — **primera entrega**|Marcar cotización como aceptada — **primera entrega**|Registrar cantidad, materiales, horas y costo — **primera entrega**|Mostrar Cotizado vs. Ejecutado — **primera entrega**|
|Seleccionar/agregar partidas — **primera entrega**|Convertir en trabajo activo — **primera entrega**|Guardar ejecución por partida — **primera entrega**|Mostrar desviación en $ y % — **primera entrega**|
|Calcular costo y precio final — **primera entrega**|Gestión avanzada de OT — después|Fotografías e incidencias — después|Dashboard, histórico e índice — después|
|Versionado y vigencia — después|Partidas extraordinarias — después|Registro avanzado por usuarios/roles — después|Predicción mediante ML — después|

La primera entrega debe recorrer el ciclo completo:

**Crear cotización → marcarla como aceptada → registrar ejecución → visualizar comparación.**

---

## Prototipo de baja fidelidad

- **Qué se muestra:** las pantallas mínimas correspondientes al camino principal del MVP:
    
    1. Inicio / selección de proyecto.
        
    2. Creación de cotización.
        
    3. Detalle de partidas.
        
    4. Cotización aceptada / trabajo activo.
        
    5. Registro simplificado de costo real.
        
    6. Comparación Cotizado vs. Ejecutado.
        
- **Quién lo hace y para cuándo:** Grupo 1 — `[completar responsable específico si el equipo lo requiere]`.
    
- **Dónde queda:** carpeta del repositorio del Grupo 1 destinada a prototipos / `[completar ruta definitiva]`.
    

---

## Entrevistas de solución

|Fecha|Rol y sector (E-NN si ya hablaron)|Qué queremos ver que hace|Siguiente paso que le vamos a pedir|
|---|---|---|---|
|`[Fecha]`|Contratista independiente / vidriería — E-01|Que intente crear una cotización y luego registrar el costo ejecutado sin recibir instrucciones paso a paso|Que realice un registro de prueba por cuenta propia|
|`[Fecha]`|Maestro contratista de obras menores — E-02|Observar si entiende el flujo Cotizar → Ejecutar → Comparar y en qué parte se pierde|Que pruebe registrar información de un trabajo real o pasado|
|`[Fecha]`|Constructor / Administrador de Obra — E-05|Observar si la información solicitada en terreno corresponde realmente a lo que registraría un responsable de obra|Que indique qué campos eliminaría o mantendría antes de desarrollar|

---

## Después de las entrevistas de solución

> Esta sección debe completarse después de mostrar el prototipo a los usuarios. No corresponde inventar los resultados antes de realizar las pruebas.

- **Qué cambió en el alcance:** `[Completar después de las entrevistas]`
    
- **Qué cambió en el prototipo:** `[Completar después de las entrevistas]`
    
- **Qué sigue siendo [H]:** `[Completar después de las entrevistas]`