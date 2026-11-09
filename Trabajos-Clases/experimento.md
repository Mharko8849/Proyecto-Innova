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

Si el registro requiere menos de 30 segundos y solicita solamente información esencial, los contratistas estarán dispuestos a registrar datos durante varios días.

## Para verificarlo

Se diseñó un formulario móvil simplificado para registrar:

- Trabajo.
    
- Partida.
    
- Cantidad ejecutada.
    
- Materiales.
    
- Costo real.
    
- Horas.
    
- Observación opcional.
    

## Mediremos

1. Cantidad de participantes.
    
2. Cantidad de registros.
    
3. Tiempo utilizado.
    
4. Registros completados sin ayuda.
    
5. Recurrencia.
    

## Tenemos razón si

Al menos tres de cinco participantes realizan tres registros y el 80 % de los registros se completa en 30 segundos o menos.

Este criterio fue definido previamente y debe mantenerse.

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

# 5. Tarjeta de aprendizaje

## Creíamos que

Un registro sencillo de menos de 30 segundos sería suficiente para permitir el uso recurrente.

## Observamos

Las entrevistas muestran que los participantes consideran importante la facilidad de uso.

Sin embargo, no contamos con los resultados completos del experimento original para verificar los tiempos y la recurrencia.

## Aprendimos que

La simplicidad debe ser una característica principal del producto.

## Decisión

Mantener pendiente la validación del experimento hasta disponer de las mediciones reales.

---

# 6. Nuevas hipótesis identificadas

Las entrevistas complementarias permiten identificar nuevas funcionalidades y condiciones comerciales.

## Cotizador modular

E-01 y E-02 prefieren configurar ventanas mediante componentes.

## Prueba gratuita

E-01 y E-02 prefieren un período de 30 días.

## Suscripciones

Existe aceptación declarada de los precios presentados y preferencia por pagos mensuales, semestrales y anuales.

## Confiabilidad del índice

E-03 identifica riesgos relacionados con valores incorrectos o manipulados.

## Directorio

E-03 considera útil encontrar prestadores según especialidad y ubicación.

---

# 7. Experimento de cotización modular

## Creemos que

Un usuario podrá preparar una cotización de una ventana personalizada en menos de diez minutos si dispone de un catálogo previamente configurado.

## Para verificarlo

Solicitaremos a usuarios de vidriería configurar una ventana seleccionando componentes.

## Mediremos

- Tiempo de configuración inicial.
    
- Tiempo de cotización.
    
- Errores.
    
- Solicitudes de ayuda.
    
- Comprensión del documento final.
    

## Resultado esperado

Comprobar si el sistema facilita la elaboración de cotizaciones personalizadas.

**Estado:** pendiente de ejecución.

---

# 8. Experimento comercial

## Creemos que

Los usuarios que obtengan valor durante 30 días estarán dispuestos a contratar una suscripción.

## Para verificarlo

Se ofrecerá una prueba gratuita informando previamente el precio posterior.

## Mediremos

1. Usuarios registrados.
    
2. Usuarios activos.
    
3. Cotizaciones generadas.
    
4. Trabajos registrados.
    
5. Conversiones a planes pagados.
    

**Estado:** pendiente de ejecución.

---

# 9. Experimento de adaptabilidad

## Creemos que

Una estructura configurable permitirá utilizar ÍndicePartidas en rubros diferentes de construcción, electricidad y vidriería.

## Para verificarlo

Se solicitará a un negocio real de otro rubro configurar una plantilla mediante la opción «Otro».

## Mediremos

- Configuración de componentes.
    
- Configuración de unidades.
    
- Uso de fórmulas.
    
- Creación de cotizaciones.
    
- Necesidad de asistencia.
    

El ejemplo de florería E-06 es simulado y no valida esta hipótesis.

---

# 10. Experimento de confiabilidad

## Creemos que

Los mecanismos de validación permitirán mejorar la calidad de los datos del índice.

## Para verificarlo

Se utilizarán registros de prueba con:

- Valores normales.
    
- Valores extremos.
    
- Duplicados.
    
- Datos inconsistentes.
    

## Mediremos

- Detección de anomalías.
    
- Identificación de duplicados.
    
- Efecto sobre los resultados agregados.
    
- Calidad de las referencias generadas.
    

**Estado:** pendiente de implementación.

---

# 11. Experimento del directorio

## Creemos que

Los prestadores aceptarán publicar voluntariamente sus perfiles y los contratistas utilizarán esa información para encontrar especialistas.

## Para verificarlo

Se habilitará un directorio de prueba con perfiles autorizados.

## Mediremos

- Prestadores registrados.
    
- Perfiles publicados.
    
- Búsquedas.
    
- Contactos realizados.
    

**Estado:** pendiente de validación.

---

# 12. Conclusión

Las entrevistas permitieron identificar nuevas necesidades, especialmente respecto de cotizaciones configurables, suscripciones y confiabilidad de información.

Sin embargo, los resultados no reemplazan las pruebas prácticas.

Se recomienda continuar con el prototipo, medir tiempos reales, comprobar el registro recurrente de costos y realizar un piloto comercial antes de considerar validado el modelo de negocio.