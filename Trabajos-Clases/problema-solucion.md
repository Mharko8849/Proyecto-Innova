# Problema y solución — Grupo 1 — Desafío: ÍndicePartidas

## El problema

**Enunciado:**

> Los contratistas independientes y pequeñas empresas de especialidades necesitan preparar cotizaciones que cubran correctamente sus costos y les permitan obtener una utilidad, pero actualmente muchos trabajan con cuadernos, planillas o presupuestos anteriores y no comparan sistemáticamente lo cotizado con lo realmente ejecutado. Esta situación aparece cada vez que preparan y posteriormente ejecutan un trabajo, pudiendo requerir desde 30 minutos hasta varias horas para cotizar y generando desviaciones informadas de entre 2 % y 30 % según el tipo de proyecto. Como resultado, algunos contratistas descubren los errores cuando ya compraron materiales o terminaron el trabajo, sin conocer con certeza su margen real. [E-01, E-02, E-04, E-05]

## Quién tiene el problema

- Contratistas independientes.
    
- Pequeñas empresas.
    
- Pymes.
    
- Prestadores de servicios especializados.
    
- Administradores de proyectos.
    

Los rubros iniciales considerados son construcción, electricidad y vidriería.

## Qué quieren lograr

Preparar cotizaciones precisas, conocer sus costos reales y utilizar la información histórica para mejorar futuras decisiones.

## Cuándo ocurre

Durante la elaboración de presupuestos y posteriormente durante la ejecución de los trabajos.

## Cuánto cuesta

E-01 reporta desviaciones de entre un 20 % y un 30 % en trabajos complejos.

E-05 menciona desviaciones de entre un 2 % y un 10 % en determinados proyectos de construcción.

Estos porcentajes representan experiencias declaradas, no promedios generales del mercado.

---

# 2. Reencuadre

Inicialmente se consideró que el problema principal era el tiempo necesario para elaborar una cotización.

Sin embargo, las entrevistas mostraron que la falta de control posterior de costos puede producir consecuencias económicas más importantes.

Por esta razón, el problema se amplió hacia la trazabilidad entre los costos presupuestados y los realmente ejecutados.

Las entrevistas posteriores también identificaron la necesidad de permitir cotizaciones personalizadas mediante componentes.

---

# 3. Alternativas actuales

Los usuarios resuelven el problema mediante:

1. Cuadernos.
    
2. Excel.
    
3. Presupuestos anteriores.
    
4. Experiencia personal.
    
5. Solicitudes de precios a proveedores.
    
6. Software especializado.
    

Estas alternativas permiten elaborar presupuestos, pero no siempre integran la cotización con el control posterior de costos.

---

# 4. Hipótesis de solución

Se consideraron tres alternativas.

## Alternativa 1: Cotizador digital

Permite crear presupuestos mediante materiales, cantidades, mano de obra y márgenes.

## Alternativa 2: Control de costos reales

Permite registrar gastos y cantidades utilizadas durante la ejecución.

## Alternativa 3: Plataforma integrada

Conecta cotizaciones, ejecución, comparación y aprendizaje histórico.

### Alternativa seleccionada

Se seleccionó la tercera alternativa porque permite abordar el problema completo.

**Cotizar → Ejecutar → Registrar → Comparar → Aprender.**

---

# 5. Hipótesis principal

Creemos que una plataforma sencilla que conecte la cotización con el registro de costos reales permitirá a los contratistas mejorar la precisión de sus presupuestos y conocer la rentabilidad de sus trabajos.

Lo verificaremos mediante pruebas de registro, cotización y comparación con usuarios reales.

---

# 6. Alcance del MVP

|Problema|Funcionalidad|Prioridad|
|---|---|---|
|Información dispersa|Crear cotizaciones|Imprescindible|
|Errores de cálculo|Costos y márgenes automáticos|Imprescindible|
|Falta de registro|Registrar ejecución real|Imprescindible|
|Desconocimiento del resultado|Comparación Cotizado vs. Ejecutado|Imprescindible|
|Productos personalizados|Configurador mínimo de componentes|Prioritario, sujeto a prueba|
|Documentos poco claros|Cotización comercial resumida|Prioritario|

El núcleo del MVP seguirá siendo la comparación entre presupuesto y ejecución.

---

# 7. Cotizador modular

Las entrevistas E-01 y E-02 identificaron la necesidad de configurar productos mediante componentes.

## Ejemplo: ventana de aluminio

El usuario podrá seleccionar:

1. Tipo de vidrio.
    
2. Espesor.
    
3. Dimensiones.
    
4. Material de estructura.
    
5. Línea de perfiles.
    
6. Tipo de manilla.
    
7. Cantidad de rodamientos.
    
8. Accesorios.
    

El software calculará internamente los costos y generará un documento comercial.

## Resultado esperado

Ventana de aluminio línea 5000.

Medidas: ancho × alto.

Precio: valor calculado.

El detalle de materiales y costos no aparecerá en la propuesta comercial.

---

# 8. Después del MVP

Se consideran las siguientes funcionalidades:

- Dashboard avanzado.
    
- Historial completo.
    
- Control de versiones.
    
- Partidas extraordinarias.
    
- Registro fotográfico.
    
- Actualización de precios.
    
- Índice agregado.
    
- Detección de datos anómalos.
    
- Directorio de prestadores.
    
- Automatización avanzada.
    

---

# 9. Fuera del MVP

- Marketplace completo.
    
- Contabilidad empresarial.
    
- Facturación electrónica.
    
- Remuneraciones.
    
- ERP completo.
    
- Lectura automática de planos.
    
- Inteligencia artificial predictiva avanzada.
    
- Aplicación móvil nativa.
    

---

# 10. Requisitos técnicos

El sistema deberá:

1. Registrar cotizaciones y partidas.
    
2. Almacenar costos y cantidades.
    
3. Relacionar cotizaciones con trabajos ejecutados.
    
4. Calcular diferencias monetarias y porcentuales.
    
5. Permitir configurar componentes.
    
6. Reutilizar estructuras de cotización.
    
7. Proteger los datos privados de cada empresa.
    
8. Permitir posteriormente generar información agregada y anonimizada.
    

## Tecnologías

|Componente|Tecnología|
|---|---|
|Frontend|React + JavaScript|
|Backend|Go|
|Base de datos|PostgreSQL|
|Contenedores|Docker|

---

# 11. Mapa de historias

|Preparar cotización|Ejecutar trabajo|Registrar costos|Analizar resultados|
|---|---|---|---|
|Crear cliente|Aceptar presupuesto|Registrar materiales|Comparar costos|
|Crear proyecto|Crear trabajo|Registrar cantidades|Calcular desviaciones|
|Seleccionar partidas|Iniciar ejecución|Registrar mano de obra|Mostrar rentabilidad|
|Configurar componentes|Registrar cambios|Agregar costos|Consultar historial|
|Calcular precio|Cerrar trabajo|Guardar información|Aprender para futuras cotizaciones|

---

# 12. Prototipo de baja fidelidad

Se consideran las siguientes pantallas:

1. Inicio.
    
2. Selección de empresa o proyecto.
    
3. Creación de cotización.
    
4. Configuración de partidas.
    
5. Configuración de componentes.
    
6. Vista previa de cotización.
    
7. Cotización aceptada.
    
8. Registro de costos.
    
9. Comparación Cotizado vs. Ejecutado.
    

---

# 13. Entrevistas de solución

|Entrevistado|Objetivo|
|---|---|
|E-01|Probar cotización modular y facilidad de uso|
|E-02|Validar configuración de componentes y documento final|
|E-03|Evaluar utilidad del índice y directorio|
|Usuario de otro rubro|Validar opción «Otro»|

---

# 14. Después de las entrevistas

## Qué cambió en el alcance

Se identificó la necesidad de incorporar un configurador modular de productos.

También surgieron requerimientos para mejorar la confiabilidad del índice y considerar un directorio voluntario.

## Qué cambió en el prototipo

Se propone incorporar la selección de componentes y la generación de un documento comercial resumido.

## Qué sigue siendo hipótesis

- Cotización en menos de diez minutos.
    
- Uso recurrente.
    
- Registro de costos en menos de 30 segundos.
    
- Adaptabilidad a otros rubros.
    
- Suscripción efectiva.
    
- Confiabilidad del índice.
    

---

# 15. Conclusión

Las entrevistas respaldan continuar con una plataforma integrada de cotización y control de costos.

El configurador modular puede mejorar la utilidad del producto, especialmente en vidriería.

Sin embargo, debe desarrollarse sin desplazar el objetivo principal de ÍndicePartidas: registrar costos reales, compararlos con lo presupuestado y utilizar esa información para mejorar futuras cotizaciones.