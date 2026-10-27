# Experimento — Grupo 1 — Desafío: ÍndicePartidas

## Nuestros supuestos

|Supuesto|Tipo (deseable / factible / viable)|Importancia (alta / media / baja)|Evidencia que tenemos hoy|
|---|---|---|---|
|Los contratistas tienen un problema real al no comparar sistemáticamente lo cotizado con el costo ejecutado|Deseable|Alta|E-01 no realiza comparación formal y reporta desviaciones de 20 % a 30 %. E-02 declara no conocer su margen real. E-05 confirma que existen desviaciones entre presupuesto y ejecución|
|Los contratistas utilizarán una herramienta para registrar costos reales si el proceso es suficientemente rápido y simple|Deseable|Alta|E-02 utilizaría el sistema si no requiere rellenar demasiados campos. E-05 exige una UX muy intuitiva. E-01 considera poco viable registrar datos si implica trabajo adicional|
|Es técnicamente posible registrar cantidades, materiales, horas y costos desde un dispositivo móvil|Factible|Alta|La solución puede implementarse como plataforma web responsiva y las entrevistas muestran uso habitual del teléfono para fotografías y comunicaciones|
|Los usuarios estarán dispuestos a aportar datos anonimizados para alimentar el índice|Viable|Alta|E-01, E-02 y E-05 muestran disposición condicionada a privacidad y protección de información comercial|
|Existirá valor para un segundo cliente que consulte información agregada de costos y rendimientos|Viable|Media|E-03 manifiesta interés en pagar mediante suscripción, pero E-05 duda del interés de constructoras grandes en información externa|
|Una estructura configurable permitirá adaptar el sistema a otros rubros además de Construcción, Electricidad y Vidriería|Factible|Media|Hipótesis de diseño del equipo; aún debe validarse con usuarios pertenecientes a otros rubros|

## El supuesto más riesgoso

- **Supuesto:** Los contratistas registrarán de forma recurrente los costos y cantidades realmente ejecutados si el proceso de ingreso es suficientemente rápido y simple.
    
- **Por qué es el más riesgoso:** El valor principal de ÍndicePartidas depende de obtener datos reales de ejecución. Si los contratistas utilizan el sistema solamente para crear cotizaciones pero no registran lo ejecutado, no será posible realizar correctamente la comparación Cotizado vs. Ejecutado ni acumular la información necesaria para construir el índice. E-01 ya manifestó que un registro adicional en terreno sería difícil porque él mismo realiza todas las actividades del trabajo.
    

## Tarjeta de prueba

- **Creemos que:** si el registro de ejecución requiere menos de 30 segundos y solicita solamente la información esencial, los contratistas estarán dispuestos a registrar datos reales de sus trabajos durante varios días.
    
- **Para verificarlo, haremos:** un experimento tipo **concierge / prototipo mínimo**, sin desarrollar todavía el software completo. Se entregará a al menos 5 contratistas un formulario móvil simplificado que simule el futuro registro de ÍndicePartidas. El formulario solicitará únicamente: trabajo o partida, cantidad ejecutada, costo/material utilizado, horas de trabajo y observación opcional.
    
- **Y mediremos:**
    
    - Número de participantes que efectivamente realizan registros.
        
    - Número de registros realizados por cada participante.
        
    - Tiempo requerido para completar cada registro.
        
    - Cantidad de participantes que vuelven a registrar información en un día diferente sin que el equipo complete los datos por ellos.
        
- **Tenemos razón si:** al menos **3 de 5 participantes realizan por cuenta propia un mínimo de 3 registros reales o reconstruidos de trabajos distintos durante el período del experimento**, y al menos **80 % de los registros se completan en 30 segundos o menos**, antes del 3 de noviembre de 2026.
    

## Plan

| Qué hay que hacer                                                                                  | Quién   | Para cuándo       |
| -------------------------------------------------------------------------------------------------- | ------- | ----------------- |
| Diseñar formulario mínimo de registro de ejecución                                                 | Grupo 1 | 30 de septiembre  |
| Probar internamente que el registro pueda realizarse rápidamente                                   | Grupo 1 | 30 de septiembre  |
| Seleccionar y contactar al menos 5 contratistas                                                    | Grupo 1 | 30 de septiembre  |
| Entregar instrucciones y comenzar el piloto                                                        | Grupo 1 | 30 de octubre     |
| Revisar diariamente únicamente si existen registros, sin completar información por el participante | Grupo 1 | 29 oct. al 2 nov. |
| Cerrar el experimento y contabilizar resultados                                                    | Grupo 1 | 3 de noviembre    |
| Completar tarjeta de aprendizaje                                                                   | Grupo 1 | 3 de noviembre    |

- **Fecha de inicio:** 30 de septiembre de 2026.
    
- **Fecha de cierre:** 3 de noviembre de 2026.
    
- **Costo estimado:** aproximadamente 8 a 12 horas de trabajo total del equipo. No se proyecta un gasto monetario significativo, ya que puede realizarse con herramientas gratuitas de formularios y hojas de cálculo.
    

> El formulario utilizado es únicamente un instrumento de prueba. No corresponde al desarrollo definitivo de ÍndicePartidas.

## Tarjeta de aprendizaje

> **Esta sección se completa una vez finalizado el experimento y no debe llenarse anticipadamente con resultados inventados.**

- **Creíamos que:** un registro de ejecución de menos de 30 segundos sería suficientemente simple para que los contratistas lo utilizaran de forma recurrente.
    
- **Observamos:** `[Completar el 3 de noviembre con los resultados reales: participantes, registros, tiempos y recurrencia]`.
    
- **De eso aprendimos que:** `[Completar después del experimento]`.
    
- **Por lo tanto, vamos a:** `[perseverar / redefinir el mecanismo de captura / concluir que el supuesto no se valida]`.