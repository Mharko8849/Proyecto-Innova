# Modelo de negocio — Grupo 1 — ÍndicePartidas

# 1. Tarjeta de aprendizaje

## Creíamos que

Si el registro de costos reales fuera rápido y sencillo, los contratistas estarían dispuestos a utilizarlo recurrentemente.

## Para verificarlo hicimos

Se diseñó un experimento de registro simplificado mediante un formulario móvil.

Las entrevistas posteriores también exploraron la aceptación comercial y nuevas funcionalidades.

## El criterio era

Al menos tres de cinco participantes debían realizar tres registros, con un mínimo del 80 % completado en 30 segundos o menos.

## Observamos

Las entrevistas muestran interés por herramientas sencillas y configurables.

Sin embargo, no se dispone de los resultados completos necesarios para verificar el cumplimiento del experimento original.

## De eso aprendimos que

La facilidad de uso es fundamental.

Además, los entrevistados valoran la posibilidad de configurar productos mediante componentes, mantener información privada y reutilizar datos.

## Lo que todavía no sabemos

Si los contratistas registrarán costos reales durante períodos prolongados y si contratarán el software después de una prueba gratuita.

---

# 2. Decisión

**Decisión propuesta: perseverar en la solución, manteniendo pendientes las validaciones técnicas y comerciales.**

Las entrevistas respaldan la relevancia del problema y muestran oportunidades para mejorar el producto.

No obstante, todavía no existe evidencia suficiente para considerar validado el modelo comercial completo.

---

# 3. Business Model Canvas

|Bloque|Propuesta y evidencia|
|---|---|
|Segmentos de clientes|Contratistas independientes y pymes de servicios técnicos/instalaciones que cotizan trabajos a medida; vidriería como primer segmento validado [E-01][E-02]. Futuramente, contratistas generales y tasadores que consultan referencias agregadas [E-03][H].|
|Propuesta de valor|Cotización ágil con componentes reutilizables, registro de costo real en terreno y comparación privada Cotizado vs. Ejecutado para conocer el margen real [E-01][E-02][H]. Futuro índice de mercado sectorial [E-03][H].|
|Canales|Contacto y venta directa, referidos del rubro y sitio web de autoservicio [H]. Asociaciones gremiales y directorio de especialistas en fases posteriores [E-03][H].|
|Relaciones con clientes|Autoservicio asistido, soporte personalizado para configurar la primera cotización y acompañamiento inicial [H].|
|Fuentes de ingreso|Suscripciones SaaS mensuales, semestrales y anuales ($9.990 Independiente, $24.990 Pyme y $59.990 Empresa) [H]. Aceptación de precios declarada en [E-01][E-02]; contratación y pago efectivo pendiente de validar en piloto [H]. Futura suscripción al índice [E-03][H].|
|Recursos clave|Plataforma web/móvil, catálogo de partidas y APU, base de datos de costos históricos privados y motor de componentes [H].|
|Actividades clave|Mantener la plataforma, facilitar registro rápido en terreno (< 30s), resguardar la confidencialidad de márgenes comerciales y soporte al usuario [H].|
|Socios clave|Proveedores de infraestructura cloud (Render/PostgreSQL), pasarelas de pago (Transbank/Flow/Mercado Pago) [H]. Futuras asociaciones de contratistas y distribuidores [H].|
|Estructura de costos|Infraestructura cloud prorrateada ($1.150 a $2.000/cuenta), comisiones de pasarela (~3,2 %), tiempo de soporte técnico y captación/onboarding ($120.000 CAC estimado) [H].|

---

# 4. Segmentos de clientes

## Cliente inicial

- Contratistas independientes.
    
- Pequeñas empresas.
    
- Pymes.
    
- Empresas de construcción.
    
- Empresas de electricidad.
    
- Empresas de vidriería.
    

## Cliente futuro

- Contratistas generales.
    
- Tasadores.
    
- Aseguradoras.
    
- Inmobiliarias.
    
- Mandantes.
    
- Profesionales que necesitan referencias de costos.
    

---

# 5. Propuesta de valor

ÍndicePartidas permitirá preparar cotizaciones, configurar productos mediante componentes y comparar los costos estimados con los reales.

La información histórica permitirá mejorar futuras cotizaciones.

Posteriormente, los datos agregados y anonimizados podrán alimentar un índice de costos y rendimientos.

---

# 6. Fuentes de ingresos

## Suscripciones SaaS

|Plan|Precio mensual|
|---|---|
|Independiente|$9.990 + IVA|
|Pyme|$24.990 + IVA|
|Empresa|$59.990 + IVA|

Se consideran pagos:

- Mensuales.
    
- Semestrales.
    
- Anuales.
    

También se propone una prueba gratuita de 30 días.

La aceptación comercial declarada por E-01 y E-02 respalda la propuesta, pero la contratación efectiva sigue pendiente de validación.

## Segunda línea de ingresos

Suscripción profesional para consultar información agregada del ÍndicePartidas.

Su precio y demanda todavía deben validarse.

---

# 7. Relaciones con clientes

Se considera un modelo de autoservicio asistido.

El usuario podrá:

1. Registrarse.
    
2. Configurar su empresa.
    
3. Seleccionar su rubro.
    
4. Cargar componentes.
    
5. Establecer precios.
    
6. Crear cotizaciones.
    
7. Registrar costos.
    
8. Comparar resultados.
    

Durante la configuración inicial podrá necesitar asistencia.

---

# 8. Recursos clave

- Backend Go.
    
- Frontend React + JavaScript.
    
- PostgreSQL.
    
- Docker.
    
- Catálogo de partidas.
    
- Motor de componentes y fórmulas.
    
- Historial de costos.
    
- Datos agregados anonimizados.
    
- Equipo de desarrollo.
    

---

# 9. Actividades clave

- Mantener el sistema.
    
- Facilitar cotizaciones.
    
- Configurar catálogos.
    
- Registrar costos reales.
    
- Comparar resultados.
    
- Actualizar información histórica.
    
- Validar datos.
    
- Proteger información comercial.
    
- Dar soporte.
    
- Captar clientes.
    

---

# 10. Directorio de prestadores

E-03 manifiesta interés por encontrar prestadores según especialidad y ubicación.

E-01 y E-02 muestran interés por publicar información profesional sin revelar precios individuales.

Por lo tanto, se propone un directorio voluntario con:

- Nombre comercial.
    
- Especialidad.
    
- Ubicación.
    
- Zonas de cobertura.
    
- Contacto.
    

La publicación será independiente del consentimiento para participar en el índice agregado.

---

# 11. ¿Cuenta una sola historia?

Sí, de manera preliminar.

El negocio comienza entregando valor directo al contratista mediante cotizaciones y control de costos.

La utilización continua de la plataforma permitiría acumular información histórica.

Posteriormente, los datos agregados podrían transformarse en un segundo producto de información.

La propuesta depende de que suficientes usuarios registren costos reales de manera recurrente.

---

# 12. Tres hipótesis pendientes

|Hipótesis|Cómo validarla|
|---|---|
|Los usuarios registrarán costos de manera recurrente|Experimento de registro y seguimiento|
|Los usuarios contratarán una suscripción|Piloto con precio declarado|
|Los canales permitirán captar clientes de manera sostenible|Medir conversiones y costos de adquisición|

Se incorporan además dos hipótesis técnicas:

- El motor configurable funcionará para otros rubros.
    
- Los mecanismos de validación permitirán generar un índice confiable.
    

---

# 13. Lo que cambió respecto del desafío inicial

Inicialmente el proyecto se centraba en un cotizador como herramienta para generar información de costos.

Las entrevistas mostraron que el valor principal también se encuentra en conocer el resultado real de los trabajos.

Posteriormente, E-01 y E-02 identificaron la necesidad de un configurador modular.

E-03 destacó la importancia de la confiabilidad del índice y la utilidad de un directorio de prestadores.

Estos hallazgos amplían la solución sin modificar su propósito central.

---

# 14. Conclusión

ÍndicePartidas mantiene un modelo SaaS dirigido inicialmente a independientes y pequeñas empresas.

El producto combinará cotizaciones configurables, registro de costos reales y comparación de resultados.

El ÍndicePartidas agregado y el directorio representan oportunidades de evolución.

Las entrevistas respaldan cualitativamente el modelo, pero todavía deben validarse la contratación efectiva, el uso continuo y la sostenibilidad económica.

