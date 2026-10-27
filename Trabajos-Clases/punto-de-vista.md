# Punto de vista — Grupo 1 — Desafío: ÍndicePartidas

## Con quién hablamos

| Código | Rol y sector                                                                          | Fecha        | ¿Propia o del banco? |
| ------ | ------------------------------------------------------------------------------------- | ------------ | -------------------- |
| E-01   | Contratista independiente de instalación de vidrios y ventanas                        | 02-10-2026   | Propia               |
| E-02   | Maestro contratista independiente de obras menores, terminaciones y remodelación      | 02-10-2026   | Propia               |
| E-03   | Contratista general (Construccion y electricidad)                                     | 03-10-2026   | Propia               |
| E-04   | Profesional de finanzas e ingeniería con experiencia en pyme y servicios industriales | 03-10-2026   | Propia               |
| E-05   | Constructor Civil con experiencia en presupuestos, cubicaciones y APU                 | Octubre 2026 | Propia               |

## Mapas de empatía

### E-01 — Contratista independiente de vidriería

- **Dice:** “Poco, ya que soy yo el único que hace todo.”
    
- **Hace:** realiza personalmente cotizaciones y trabajos. Cotiza manualmente, no realiza una comparación sistemática al finalizar y detecta los sobrecostos cuando debe comprar materiales adicionales o aparece un imprevisto.
    
- **Piensa:** _inferencia / hipótesis:_ cualquier herramienta que le quite demasiado tiempo en terreno probablemente será abandonada.
    
- **Siente:** _inferencia / hipótesis:_ preocupación por perder dinero debido a imprevistos y temor a que la información compartida permita identificarlo o perjudicarlo comercialmente.
    

### E-02 — Maestro contratista de obras menores

- **Dice:** “No, no lo comparo, no sé realmente cuál es mi margen de ganancia.”
    
- **Hace:** prepara presupuestos en cuaderno, dedica varias horas a cotizar y descubre errores de cálculo cuando el dinero disponible queda demasiado ajustado frente a los materiales necesarios.
    
- **Piensa:** _inferencia / hipótesis:_ conocer el margen real sería útil, siempre que obtener esa información no implique aumentar mucho su trabajo administrativo.
    
- **Siente:** _inferencia / hipótesis:_ frustración frente a errores de cubicación y rechazo a pagar por software cuyo costo considera alto.
    

### E-03 — Contratista general / coordinador de obras

- **Dice:** “Preferiría una suscripción mensual con reportes, pero una de las desventajas que considero es que se depende de la buena fe de las empresas.”
    
- **Hace:** compara presupuestos de distintos maestros y utiliza experiencia previa para decidir si un precio es razonable.
    
- **Piensa:** _inferencia / hipótesis:_ un índice podría ser útil si los datos son confiables y además facilita encontrar especialistas disponibles.
    
- **Siente:** _inferencia / hipótesis:_ desconfianza respecto de información de mercado que no tenga mecanismos de validación.
    

### E-04 — Profesional de finanzas e ingeniería

- **Dice:** “Muchas empresas quiebran por no hacer los presupuestos como corresponde.”
    
- **Hace:** utiliza información histórica y criterios financieros para estructurar presupuestos y destaca la importancia de mantener precios actualizados y controlar el flujo de caja.
    
- **Piensa:** _inferencia / hipótesis:_ la rapidez de cotizar tiene menos valor que disponer de información histórica confiable que evite vender bajo costo.
    
- **Siente:** _inferencia / hipótesis:_ preocupación por el impacto financiero que puede producir una mala estimación en una pyme.
    

### E-05 — Constructor Civil

- **Dice:** “En la construcción, el presupuesto es una apuesta, una apuesta que puede ser buena o mala.”
    
- **Hace:** trabaja con APU, cubicaciones, precios históricos y presupuestos a suma alzada. Identifica al Administrador de Obra como responsable adecuado para registrar costos y cantidades.
    
- **Piensa:** _inferencia / hipótesis:_ una herramienta nueva solo tendrá posibilidades de adopción si es extremadamente intuitiva y protege los márgenes de la empresa.
    
- **Siente:** _inferencia / hipótesis:_ preocupación por entregar información que posteriormente pueda ser utilizada por competidores para reducir precios.
    

## Patrones

### Se repite

- Las cotizaciones todavía dependen fuertemente de procesos manuales, planillas, cuadernos o históricos propios.
    
- La información de trabajos anteriores no siempre se utiliza de forma sistemática para mejorar el siguiente presupuesto.
    
- Existen desviaciones entre lo presupuestado y lo realmente ejecutado.
    
- Los contratistas pequeños no siempre saben cuál fue su margen real.
    
- Registrar información adicional en terreno genera resistencia si el proceso requiere demasiado tiempo.
    
- La facilidad de uso es fundamental para lograr adopción.
    
- Existe preocupación por la confidencialidad de costos y márgenes.
    
- Los datos históricos son considerados más valiosos cuando permiten tomar mejores decisiones futuras.
    

### Nos sorprendió

- Inicialmente pensábamos que cualquier trabajador podría registrar costos en terreno, pero E-05 indicó que esa función debería recaer principalmente en el Administrador de Obra o Jefe de Terreno.
    
- Suponíamos que los contratistas pequeños conocían aproximadamente su rentabilidad después de cada trabajo, pero E-02 declaró directamente no saber cuál es su margen real.
    
- Creíamos que el principal valor sería cotizar más rápido, pero E-04 mostró que evitar errores y utilizar costos históricos confiables puede tener mayor impacto.
    
- Pensábamos que las constructoras serían compradores naturales del índice, pero E-05 mostró poco interés en datos de competidores, mientras E-03 sí encontró valor y agregó la necesidad de un directorio de prestadores.
    

## Punto de vista

**Un contratista pequeño o trabajador independiente que cotiza y ejecuta trabajos** necesita **conocer de forma simple cuánto le costó realmente cada trabajo y utilizar esa experiencia para preparar mejores presupuestos futuros**, porque **actualmente dedica tiempo a cotizar, pero muchas veces no realiza un análisis posterior y descubre las pérdidas o errores cuando el trabajo ya está ejecutado**.

- **Basado en:** E-01, E-02, E-04 y E-05.
    
- **¿Qué parte es hipótesis?:** que un proceso digital suficientemente simple logrará que el contratista registre esta información de manera recurrente. Este supuesto debe validarse mediante el experimento.
    

## ¿Cómo podríamos...?

1. ¿Cómo podríamos permitir que un contratista registre el costo real de un trabajo sin transformarlo en una nueva carga administrativa?
    
2. ¿Cómo podríamos utilizar los trabajos anteriores para advertir al usuario cuando una nueva cotización presenta riesgo de estar subestimada?
    
3. ¿Cómo podríamos construir un índice útil de costos y rendimientos sin revelar la identidad, clientes ni márgenes comerciales de quienes aportan los datos?
    

## Qué preguntar en la próxima conversación

- ¿Cuál es la cantidad mínima de información que estaría dispuesto a registrar después de ejecutar un trabajo?
    
- ¿En qué momento del día o del proceso sería más natural realizar ese registro?
    
- ¿Qué función concreta haría que utilizara ÍndicePartidas nuevamente después de preparar su primera cotización?
    
- ¿Qué información debería ocultarse obligatoriamente antes de permitir que sus datos alimenten un índice agregado?
    
- ¿Qué referencia del índice utilizaría realmente para cambiar el precio de una próxima cotización?