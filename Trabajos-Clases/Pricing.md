# Pricing — Grupo 1 — Desafío: ÍndicePartidas

## Métrica de valor

- **Por qué unidad crece el valor para el cliente:** cantidad de **cotizaciones y trabajos gestionados al mes**.
    
    El valor de ÍndicePartidas aumenta a medida que el contratista cotiza y ejecuta más trabajos, porque puede reutilizar más información histórica, controlar una mayor cantidad de costos reales y detectar más desviaciones entre lo presupuestado y lo ejecutado. [E-01, E-02, E-04]
    
- **¿El cliente puede contarla solo, sin que se la expliquen?** Sí. El contratista conoce cuántas cotizaciones genera y cuántos trabajos ejecuta.
    
- **¿Crece cuando al cliente le va bien?** Sí. Si la empresa consigue más clientes y ejecuta más trabajos, aumenta la cantidad de cotizaciones, órdenes de trabajo y comparaciones de costos que administra mediante la plataforma.
    
- **Forma de cobro elegida:** **suscripción mensual por tramos**.
    
- **Por qué esa y no otra:** las entrevistas muestran clientes de tamaños y formas de trabajo diferentes. Un contratista independiente no debería pagar lo mismo que una empresa con varios proyectos y trabajadores. Además, E-03 indicó explícitamente preferir una **suscripción mensual** para acceder al índice y sus reportes. [E-03]
    
    No se propone cobrar por usuario, porque el valor principal no aumenta necesariamente por tener más personas con acceso, sino por la cantidad de cotizaciones y trabajos que la empresa administra. Tampoco se cobrará inicialmente por transacción o como porcentaje de cada obra, porque ÍndicePartidas no participa directamente en el pago del trabajo. [H]
    

---

## Cuánto vale el problema

Para estimar el valor mensual se utiliza como referencia un contratista pequeño que realiza aproximadamente **10 cotizaciones mensuales** [H].

E-01 indicó que cada cotización tarda entre **30 minutos y 2 horas**. Para la estimación se utiliza un promedio de **1,25 horas por cotización** [H a partir del rango E-01].

El valor de una hora del contratista se estima inicialmente en **$10.000 CLP por hora** [H]. Este valor deberá validarse posteriormente.

|Cuenta|Cálculo|Monto mensual|Fuente|
|---|---|---|---|
|Horas destinadas a cotizar|10 cotizaciones/mes [H] × 1,25 h [E-01/H] × $10.000/h [H]|**$125.000 [H]**|E-01 + [H]|
|Tiempo destinado a revisar o reconstruir precios anteriores|No existe información suficiente para monetizarlo actualmente|No calculado|E-02, E-04|
|Pérdidas por errores de cotización|E-01 reporta desviaciones de 20 % a 30 % en trabajos complejos, pero no conocemos el valor promedio de cada trabajo ni su frecuencia|No calculado|E-01|
|Pérdidas por desviaciones en construcción|E-05 reporta desviaciones de 2 % a 10 %, pero no disponemos de un monto promedio de obra para monetizarlas|No calculado|E-05|
|**Total monetizable actualmente**|Solo tiempo destinado a cotizar|**$125.000/mes [H]**|E-01 + [H]|

> Esta estimación es conservadora porque solo monetiza el tiempo destinado a cotizar. No incorpora todavía el valor económico de evitar errores de presupuesto, sobrecostos o pérdida de margen, aunque las entrevistas muestran que esos problemas existen. [E-01, E-02, E-04, E-05]

La clase propone utilizar el precio como una **fracción del valor creado**, aproximadamente entre un décimo y un tercio como referencia inicial. Con un problema estimado en $125.000 mensuales [H], un precio cercano a **$24.990 mensuales** representa aproximadamente un 20 % de ese valor. [H]

---

## Paquetes y precio

| Bueno — **Independiente**               | Mejor — **Pyme**                                                                     | Óptimo — **Empresa**                                                                                     |                                                                                                                                                                             |
| --------------------------------------- | ------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Para quién**                          | Trabajador independiente o microempresa con pocos trabajos mensuales                 | Pyme contratista que cotiza y ejecuta trabajos regularmente                                              | Empresa con varios proyectos, responsables y necesidad de mayor control                                                                                                     |
| **Cotizaciones / trabajos gestionados** | Hasta 10 al mes [H]                                                                  | Hasta 40 al mes [H]                                                                                      | Sin límite definido para el piloto [H]                                                                                                                                      |
| **Usuarios**                            | 1 usuario [H]                                                                        | Hasta 5 usuarios [H]                                                                                     | Hasta 15 usuarios [H]                                                                                                                                                       |
| **Qué incluye**                         | Cotizaciones, catálogo propio, partidas, costos y comparación Cotizado vs. Ejecutado | Todo lo anterior + historial completo, órdenes de trabajo, dashboard, roles y rubro personalizado “Otro” | Todo lo anterior + múltiples responsables, reportes avanzados, prioridad de soporte y acceso ampliado a información del ÍndicePartidas cuando exista suficiente información |
| **Qué lo separa del anterior**          | Funciones esenciales para un contratista individual                                  | Más trabajos, usuarios, historial y control de empresa                                                   | Mayor capacidad, administración y analítica                                                                                                                                 |
| **Precio mensual (más IVA)**            | **$9.990 [H]**                                                                       | **$24.990 [H]**                                                                                          | **$59.990 [H]**                                                                                                                                                             |
| **Precio anual (con descuento)**        | **$99.900 + IVA [H]**                                                                | **$249.900 + IVA [H]**                                                                                   | **$599.900 + IVA [H]**                                                                                                                                                      |

### Tramos de tamaño

Los paquetes se diferencian principalmente por el volumen de trabajos gestionados y por la estructura de la empresa.

Como referencia complementaria, la clase indica que en Chile pueden utilizarse los tramos de la Ley 20.416:

- Microempresa: hasta 2.400 UF anuales.
    
- Pequeña empresa: hasta 25.000 UF anuales.
    
- Mediana empresa: hasta 100.000 UF anuales.
    

Sin embargo, para ÍndicePartidas consideramos más comprensible para el cliente utilizar como principal métrica la **cantidad de cotizaciones y trabajos gestionados al mes**, porque corresponde directamente al uso y valor obtenido de la plataforma. [H]

### Prueba o piloto

Sí.

Se propone realizar un **piloto con fecha de término y precio declarado**, evitando un modelo freemium permanente.

- Duración propuesta: **30 días [H]**.
    
- Precio posterior informado desde el inicio: **$24.990 + IVA mensuales para el paquete Pyme [H]**.
    
- Durante el piloto se observará si el usuario registra cotizaciones y costos reales y si manifiesta un compromiso concreto frente al precio.
    

El piloto puede ser gratuito únicamente como experimento y a cambio de permitir medir el uso, pero el usuario deberá conocer desde el comienzo cuál será el precio posterior. [H]

---

## Justificación del precio del paquete del medio

El paquete **Pyme de $24.990 + IVA mensuales [H]** es el paquete que se espera que concentre la mayor parte de los clientes.

|Fuente|Dato|De dónde sale|
|---|---|---|
|Lo que el cliente gasta hoy en el problema|Aproximadamente **$125.000 mensuales [H]** solamente considerando tiempo de cotización para el escenario utilizado|E-01 + supuestos declarados [H]|
|Lo que cobra la alternativa principal|Excel puede adquirirse por aproximadamente **$179.990 como licencia individual**, mientras APU Chile ofrece herramientas desde aproximadamente **$54.990**, según el levantamiento realizado por el equipo|`competencia.md`|
|Barrera observada en alternativas|E-02 investigó software especializado y desistió debido a su alto precio|E-02|
|Fracción del valor del problema que captura el precio|$24.990 / $125.000 ≈ **20 % [H]**|cálculo|
|Preferencia de forma de pago|Un entrevistado indicó preferir una **suscripción mensual con reportes**|E-03|
|Lo que alguien estuvo dispuesto a hacer frente a un precio|**Todavía no validado con un precio concreto**|pendiente de experimento|

### Por qué $24.990 y no otro valor

El precio no se obtiene del costo de desarrollar el software, porque el costo marginal de atender a un cliente adicional en un SaaS es relativamente bajo.

Tampoco se copia directamente el valor de la competencia.

Se propone inicialmente **$24.990 + IVA mensuales [H]** porque:

1. representa aproximadamente un **20 % del valor mensual estimado del problema** en el escenario utilizado [H];
    
2. queda considerablemente por debajo de algunas alternativas especializadas identificadas;
    
3. E-02 mostró que el precio del software es una barrera real, por lo que el paquete principal debe mantenerse accesible;
    
4. la suscripción mensual coincide con la modalidad preferida por E-03;
    
5. permite que el cliente conserve la mayor parte del valor económico generado.
    

El precio todavía debe considerarse una **hipótesis**, porque ningún entrevistado ha enfrentado aún una decisión real de contratación a $24.990 mensuales.

---

## Segunda línea de negocio: ÍndicePartidas

La tesis del desafío plantea que el software de cotización permite capturar información, pero el activo de mayor valor a largo plazo corresponde a los **datos agregados de costos y rendimientos reales**.

Por ello se considera una segunda línea de negocio dirigida a:

- aseguradoras;
    
- tasadores;
    
- inmobiliarias;
    
- mandantes;
    
- contratistas generales que necesitan comparar precios de especialidades.
    

E-03 mostró interés en pagar mediante suscripción por información agregada y reportes, especialmente si se complementa con un directorio de contratistas. [E-03]

Sin embargo:

- E-05 mostró que una empresa constructora puede valorar más sus propios datos que los datos de competidores. [E-05]
    
- todavía no existe suficiente evidencia para establecer un precio concreto para el acceso profesional al índice.
    

Por lo tanto, el precio de esta segunda línea se mantiene **[H]** y deberá validarse después de demostrar que existe suficiente volumen y confiabilidad de datos.

**Modelo futuro propuesto [H]:**

- Suscripción profesional mensual para acceso a reportes del índice.
    
- Segmentación de datos por rubro, partida, zona y período.
    
- Exportación de reportes.
    
- Directorio voluntario de prestadores.
    
- Precio diferenciado respecto del SaaS utilizado por los contratistas.
    

---

## Experimento de precio

- **Creemos que** un pequeño contratista o dueño de pyme estaría dispuesto a contratar el paquete **Pyme por $24.990 + IVA mensuales [H]** si la plataforma le permite crear cotizaciones, registrar costos reales y conocer automáticamente sus desviaciones y rentabilidad.
    
- **Para verificarlo:** presentaremos una propuesta simple con las funciones incluidas y el precio mensual claramente visible. Se solicitará una **firma de intención** para participar en un piloto posterior al desarrollo o, si es posible, aceptar un piloto pagado.
    
- **Mediremos:** cuántas personas realizan una acción concreta frente al precio:
    
    - firman una carta/intención de participación;
        
    - solicitan iniciar el piloto;
        
    - entregan datos de contacto para contratar cuando esté disponible;
        
    - o aceptan realizar un pago/piloto pagado.
        
- **Tenemos razón si:** al menos **2 de 5 potenciales clientes [H]** realizan una acción de compromiso concreta frente al precio de $24.990 + IVA mensuales antes del **9 de noviembre de 2026 [H]**.
    
- **Cuándo y con quién:** antes del 9 de noviembre de 2026, idealmente con contratistas independientes y pequeñas empresas similares a los perfiles E-01 y E-02.
    

> Preguntar simplemente “¿pagarías $24.990?” no se considerará evidencia suficiente. Se medirá una acción observable frente al precio.

---

## Lo que todavía es [H]

- La cantidad promedio de **10 cotizaciones mensuales** utilizada para estimar el escenario.
    
- El valor de **$10.000 por hora** utilizado para monetizar el tiempo del contratista.
    
- El valor mensual estimado del problema de **$125.000**.
    
- Los límites de 10 y 40 trabajos mensuales de los paquetes.
    
- Los límites de usuarios por paquete.
    
- Los precios de **$9.990, $24.990 y $59.990**.
    
- El descuento anual propuesto.
    
- Que $24.990 sea un precio suficientemente bajo para un contratista pequeño y suficientemente alto para sostener el negocio.
    
- Que la cantidad de cotizaciones/trabajos gestionados sea la mejor métrica de valor.
    
- El precio y modelo comercial del acceso profesional al ÍndicePartidas.
    
- La disposición real a contratar después de enfrentar un precio concreto.
    
- Que al menos 2 de 5 usuarios realicen un compromiso concreto en el experimento de precio.