# Pricing — Grupo 1 — ÍndicePartidas

# 1. Métrica de valor

La métrica propuesta corresponde a la cantidad de cotizaciones y trabajos gestionados mensualmente.

El valor del software aumenta a medida que el usuario administra más cotizaciones, proyectos y registros de ejecución.

Esta métrica es comprensible para los clientes porque conocen aproximadamente su volumen de trabajo.

**Estado:** hipótesis comercial pendiente de validación [H].

---

# 2. Cuánto vale el problema

Para realizar una estimación inicial se considera un contratista que genera diez cotizaciones mensuales.

E-01 indicó que una cotización puede requerir entre 30 minutos y 2 horas.

Se utiliza como referencia un promedio estimado de 1,25 horas.

El valor de una hora del contratista se estima en $10.000.

|Concepto|Valor|
|---|---|
|Cotizaciones mensuales|10|
|Tiempo promedio|1,25 horas|
|Valor de la hora|$10.000|
|Valor mensual estimado del tiempo|$125.000|

Los valores de cotizaciones mensuales y costo por hora son supuestos de análisis.

La estimación no incluye las posibles pérdidas económicas por errores de presupuesto.

---

# 3. Paquetes y precios

Se proponen tres planes de suscripción.

|Característica|Independiente|Pyme|Empresa|
|---|---|---|---|
|Cliente objetivo|Trabajador independiente|Pequeña empresa|Empresa con varios proyectos|
|Cotizaciones mensuales|Hasta 10 [H]|Hasta 40 [H]|Sin límite inicial definido [H]|
|Usuarios|1 [H]|Hasta 5 [H]|Hasta 15 [H]|
|Cotizaciones|Sí|Sí|Sí|
|Catálogo de partidas|Sí|Sí|Sí|
|Control Cotizado vs. Ejecutado|Sí|Sí|Sí|
|Historial avanzado|Básico|Completo|Completo|
|Roles y permisos|Básico|Sí|Avanzado|
|Dashboard|Básico|Sí|Avanzado|
|Rubro personalizado|Por validar|Sí|Sí|
|Precio mensual|$9.990 + IVA|$24.990 + IVA|$59.990 + IVA|

Los límites y diferencias funcionales son propuestas comerciales pendientes de validación.

---

# 4. Modalidades de pago

E-01 y E-02 manifestaron preferencia por disponer de opciones mensuales, semestrales y anuales.

También consideran importante que los períodos más largos incluyan descuentos.

## Plan Independiente

|Modalidad|Precio base|Ahorro|
|---|---|---|
|Mensual|$9.990|—|
|Semestral|$49.990|$9.950|
|Anual|$99.900|$19.980|

Los valores se presentan antes de IVA, manteniendo el criterio del modelo de precios original.

### Cálculo semestral

Seis mensualidades:

6 × $9.990 = $59.940.

Precio semestral:

$49.990.

Ahorro:

$59.940 - $49.990 = $9.950.

Descuento aproximado: **16,6 %**.

### Cálculo anual

Doce mensualidades:

12 × $9.990 = $119.880.

Precio anual:

$99.900.

Ahorro:

$119.880 - $99.900 = $19.980.

Descuento aproximado: **16,7 %**.

Los descuentos semestrales y anuales son prácticamente equivalentes.

Se recomienda evaluar una diferencia mayor para incentivar la contratación anual.

Los precios semestrales de los planes Pyme y Empresa permanecen pendientes de definición.

---

# 5. Prueba gratuita

Las entrevistas E-01 y E-02 muestran preferencia por un período de prueba gratuito de 30 días.

El objetivo de este período será permitir que los usuarios:

1. Configuren su empresa.
    
2. Carguen componentes y precios.
    
3. Generen cotizaciones.
    
4. Registren costos reales.
    
5. Comparen resultados.
    
6. Evalúen si el software se adapta a su trabajo.
    

Al comenzar la prueba se deberá informar claramente el precio posterior.

La conversión al finalizar los 30 días debe medirse mediante un experimento real.

---

# 6. Justificación del precio del paquete Pyme

Se propone un precio mensual de $24.990 + IVA.

Considerando el escenario estimado de $125.000 mensuales asociados al tiempo de cotización:

$24.990 / $125.000 ≈ 20 %.

El precio representaría aproximadamente un 20 % del valor mensual estimado del problema.

Esta comparación es orientativa, porque los supuestos económicos todavía necesitan validación.

---

# 7. Validación de precios

## Evidencia obtenida

E-01 y E-02 manifestaron conformidad con los precios presentados.

También prefirieron:

- Prueba gratuita de un mes.
    
- Pagos mensuales.
    
- Pagos semestrales.
    
- Pagos anuales.
    
- Descuentos por mayor período de contratación.
    

## Interpretación

Las entrevistas respaldan cualitativamente la aceptación del esquema comercial [E-01][E-02].

Sin embargo, no demuestran disposición efectiva a pagar.

Esta deberá comprobarse mediante contratos, pagos, compromisos verificables o conversión después del piloto [H].

## Comparación con alternativas del mercado

| Alternativa | Precio | Modalidad | Cierra el ciclo Cotizar → Registrar → Comparar |
|---|---|---|---|
| Microsoft Excel | $179.990 | Compra única | No — el usuario construye y mantiene las planillas a mano [E-01] |
| APU Chile | Desde $54.990 | Compra única | No — solo presupuesto y base de datos, sin registro de ejecución |
| Presto | Cotizar (licencia empresa) | Licencia de alto costo | Sí, pero orientado a grandes constructoras y obra civil [competencia.md] |
| Procore / Buildertrend | Cotizar (suscripción) | Suscripción empresarial | Sí, pero excede las necesidades y presupuesto de pymes pequeñas [competencia.md] |
| **ÍndicePartidas** | **$9.990 – $24.990 + IVA/mes** | **Suscripción mensual accesible** | **Sí — diseñado específicamente para independientes y pymes de especialidad** [H] |

La propuesta de ÍndicePartidas es la única solución con suscripción mensual de bajo costo que cubre el recorrido continuo (presupuestar, registrar costo real en terreno y comparar desviaciones), resolviendo la brecha que hoy dejan las planillas manuales y el software corporativo [E-02][competencia.md].

## Resultado del experimento de precio (compromisos verificables)

**Objetivo comercial:** Obtener **dos compromisos verificables con pymes del sector**, presentando la oferta del plan Pyme ($24.990 + IVA/mes) tras los 30 días de prueba gratuita.

| Contacto | Fecha | Perfil | Oferta presentada | Resultado |
|---|---|---|---|---|
| Pyme-01 | _pendiente_ | _pendiente_ | Plan Pyme $24.990 + 30 días gratis | _pendiente_ |
| Pyme-02 | _pendiente_ | _pendiente_ | Plan Pyme $24.990 + 30 días gratis | _pendiente_ |

- **Compromisos obtenidos:** _pendiente_ / 2
- **Rechazos y razones:** _pendiente_

> **Nota metodológica:** Un compromiso verificable corresponde a una confirmación concreta (por escrito, correo o mensajería) en la que el decisor manifiesta su aceptación a participar en el piloto con el precio informado. Las declaraciones verbales de simpatía hacia el producto registradas en entrevistas exploratorias previas no constituyen prueba de compromiso de pago.

---

# 8. Segunda línea de negocio

El ÍndicePartidas agregado podría comercializarse posteriormente como una fuente de información para:

- Contratistas generales.
    
- Tasadores.
    
- Aseguradoras.
    
- Inmobiliarias.
    
- Mandantes.
    
- Profesionales que revisan presupuestos.
    

E-03 manifestó interés en disponer de:

1. Precios actualizados.
    
2. Cantidad de registros.
    
3. Valores mínimos y máximos.
    
4. Información segmentada.
    
5. Directorio de prestadores.
    

El precio de esta segunda línea todavía no se encuentra validado.

---

# 9. Experimento de precio

## Creemos que

Los contratistas encontrarán valor suficiente en ÍndicePartidas para contratar una suscripción después de conocer sus funciones y utilizar el producto.

## Para verificarlo

Se ofrecerá una prueba de 30 días con el precio posterior informado.

## Mediremos

- Cantidad de usuarios que activan el producto.
    
- Cantidad de cotizaciones generadas.
    
- Cantidad de trabajos registrados.
    
- Frecuencia de uso.
    
- Solicitudes de contratación.
    
- Conversiones a planes pagados.
    

## Criterio original de precio

El experimento inicial estableció como criterio obtener al menos dos compromisos verificables entre cinco potenciales clientes.

Las entrevistas verbales no son suficientes para considerar cumplido dicho criterio.

---

# 10. Hipótesis pendientes

- Disposición efectiva a pagar.
    
- Conversión al finalizar el período gratuito.
    
- Límites de cotizaciones.
    
- Cantidad adecuada de usuarios por plan.
    
- Sostenibilidad de los descuentos.
    
- Precio adecuado del plan Pyme.
    
- Precio de la segunda línea de negocio.
    
- Retención de clientes.
    
- Costo de adquisición.
    

---

# 11. Conclusión

Las entrevistas respaldan cualitativamente la estructura de precios propuesta y la conveniencia de incorporar una prueba gratuita de 30 días.

Se recomienda mantener los planes Independiente, Pyme y Empresa, agregando modalidades mensuales, semestrales y anuales.

La siguiente etapa deberá comprobar si la aceptación declarada se transforma efectivamente en contratación y uso recurrente.