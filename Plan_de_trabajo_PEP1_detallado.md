# Plan de trabajo — 6 bloques de pendientes concretos para PEP 1

**Versión corregida: distribución de trabajo pendiente, no guía paso a paso.**

Este plan reemplaza la guía extensa de siete integrantes. Se basa en el último corte auditado (`480b627`), no en una nueva revisión de cambios posteriores. Cada bloque puede asumirlo una persona con apoyo puntual del equipo; no es necesario asignar un bloque a todos los integrantes.

| Responsable | Trabajo pendiente | Entrega principal |
|---|---|---|
| 1 | Corregir entrevistas y documentación faltante | Bitácora, resúmenes y guion completos. |
| 2 | Validar el prototipo y cerrar el experimento | Pruebas observadas, tiempos y aprendizaje. |
| 3 | Respaldar negocio y compromiso frente al precio | Supuestos justificados y evidencia comercial. |
| 4 | Consolidar decisiones, agentes y alcance | Bitácoras y documentos coherentes. |
| 5 | Completar el informe | PDF final con resultados y anexos. |
| 6 | Preparar prototipo, Canva y entrega | Prototipo consultable, slides PDF y entrega comprobada. |

---

## 1. Corregir entrevistas y completar documentación faltante

**Responsable 1**

- [ ] Completar el resumen de la entrevista de **climatización**: hay seis conversaciones en bitácora, pero cinco resúmenes.
- [ ] Corregir fechas inconsistentes de E-03/E-04 y precisar E-05 si existe respaldo.
- [ ] Aclarar el perfil de E-02 y fechar los seguimientos de entrevistas.
- [ ] Separar la **florería simulada** de las entrevistas reales y resolver el uso del código E-06.
- [ ] Confirmar consentimiento y publicación de los resúmenes en el banco del curso.
- [ ] Incorporar **`guion-entrevista.md`**, identificando el material existente que corresponde a ese entregable.
- [ ] Eliminar textos de plantilla y matizar afirmaciones exageradas como “plenamente validado” o “las constructoras no pagarían”.

**Archivos afectados:** `Trabajos-Clases/bitacora-contactos.md`, `resumen-entrevistas.md`, `Entrevista-Validar Hipotesis.md`, documentos que los citan y el guion faltante.

**Entrega:** bitácora, resúmenes y guion consistentes y completos; listado de datos que no pudieron recuperarse. No inventar fechas, consentimiento ni entrevistas.

## 2. Validar el prototipo y cerrar el estado del experimento

**Responsable 2, con apoyo de quien tenga los contactos.** Agrupa validación y cronometraje, pero distingue sus resultados.

- [ ] Probar el prototipo con **al menos dos personas del segmento**.
- [ ] Observar cómo cotizan, registran un costo y encuentran la comparación.
- [ ] Registrar **dónde se pierden, qué ignoran, qué ayuda necesitan y qué cambiar**.
- [ ] Cronometrar por separado el registro de cuenta, la elaboración de cotización y el **registro de costos**. Indicar uso de guía integrada o ayuda del entrevistador.
- [ ] Conseguir y documentar al menos un siguiente paso concreto: otra reunión, datos reales o presentación a otro contacto.
- [ ] Aclarar si el experimento original **se ejecutó, quedó parcial o no se realizó**. El informe y `experimento.md` deben coincidir.
- [ ] Completar `experimento.md` con resultados disponibles, comparación con el criterio y decisión.
- [ ] Documentar qué cambió en el prototipo por el feedback recibido.

**Entrega:** notas de pruebas, tiempos reales, siguientes pasos, cambios por feedback y tarjeta de aprendizaje.

**Importante:** el umbral de **30 segundos** corresponde al registro de costos, no a registrarse ni a preparar toda la cotización. El criterio original contempla **3 de 5 participantes**, tres registros y **80 % de registros en 30 segundos o menos**; conservar también sus condiciones de autonomía y recurrencia. Una sesión no demuestra recurrencia entre días, y una prueba nueva no cumple retroactivamente el plazo anterior al **3 de noviembre**.

**Coordinación:** el responsable 6 proporciona un recorrido estable; el 1 ayuda con contactos. Entregar resultados a 4, 5 y 6.

## 3. Respaldar el negocio y comprobar compromiso frente al precio

**Responsable 3, con apoyo para contactar clientes.** Agrupa sostenibilidad del negocio y disposición efectiva a pagar.

- [ ] Desglosar y justificar los supuestos existentes: costos de atención de **$2.500/$6.000**, **$120.000 de captación** y mezcla **70/30 de planes**.
- [ ] Añadir fuentes de infraestructura, comisiones y valorización del soporte.
- [ ] Comparar el precio con alternativas de `competencia.md`, distinguiendo compra única y suscripción.
- [ ] Mantener sensibilidad y revisar el efecto de descuentos sobre el ingreso.
- [ ] Identificar los **$125.000** como escenario de tiempo valorizado, no gasto observado ni ahorro demostrado.
- [ ] Aplicar el criterio comercial original: **dos compromisos verificables entre cinco potenciales clientes**, frente a la oferta y precio correspondientes.
- [ ] Registrar también rechazos y sus razones, no solo respuestas favorables.
- [ ] Sincronizar precios, Canvas y cifras entre Markdown, informe y slides.

**Archivos afectados:** `Plantilla_pep1/unit_economics.tex`, sección de negocio del informe, `Trabajos-Clases/Pricing.md` y `Modelo_Negocio.md`.

**Entrega:** números respaldados, pricing coherente y resultados comerciales, separando aceptación verbal, compromiso y pago.

**Corrección importante:** los unit economics **ya existen y sus cálculos son coherentes**; falta fortalecer sus fuentes. No necesitan medir meses de permanencia para esta entrega: retención y LTV pueden seguir como hipótesis explícitas. El umbral 2/5 es del equipo, no una condición literal universal de la rúbrica. Una carta con precio no equivale a pago; conservar firmas y datos privados fuera del repo público.

## 4. Consolidar decisiones, agentes y alcance

**Responsable 4**

- [ ] Crear una bitácora con **fecha, evidencia, decisión y cambio de alcance**, incluyendo qué pensaban antes y qué encontraron.
- [ ] Registrar qué se pidió a los agentes, qué produjeron y qué verificó/corrigió el equipo. Incluir asistencia documental y Figma Make/Stitch cuando efectivamente se usaron.
- [ ] Completar y registrar la revisión humana del informe generado.
- [ ] Unificar el **segmento inicial**, las **tres capacidades principales del MVP**, qué queda para después y qué queda fuera.
- [ ] Marcar el **corte de primera entrega en el mapa de historias**.
- [ ] Corregir contradicciones residuales Vue/React y la referencia a “Spring Security”.

**Archivos afectados:** nuevas bitácoras, `problema-solucion.md`, `propuesta-valor.md`, informe y referencias técnicas inconsistentes.

**Entrega:** bitácoras verificables y documentos que cuenten la misma historia. Si se reconstruye una decisión ahora, identificarlo; no simular un registro escrito en octubre.

## 5. Completar el informe existente

**Responsable 5**

- [ ] Continuar el borrador actual; **no empezar desde cero**.
- [ ] Incorporar resultados de entrevistas, pruebas, experimento y negocio entregados por 1–4.
- [ ] Actualizar la descripción del prototipo y sus limitaciones con el responsable 6.
- [ ] Completar anexos con evidencia consultable.
- [ ] Revisar citas `[E-NN]`, hipótesis `[H]`, cifras, anonimización y coherencia con el repositorio.
- [ ] Detallar qué no se validó, cómo se intentó y qué faltó; no limitarse a prometer pruebas futuras.
- [ ] Eliminar placeholders, tablas cortadas y problemas de maquetación del PDF.
- [ ] Incorporar feedback de la defensa antes del envío del jueves.

**Archivos afectados:** `Plantilla_pep1/informe_indicepartidas.tex`, archivos incluidos y PDF final.

**Entrega:** informe de máximo **15 páginas más anexos**, revisado por otro integrante. No inventar resultados para rellenar pendientes de otros bloques.

## 6. Preparar prototipo, Canva y entrega digital

**Responsable 6, con revisión y apoyo del equipo**

- [ ] Corregir solo fallos que afecten el recorrido: campos de ejecución que no se guardan; condiciones comerciales con texto fijo; cálculos inconsistentes de cantidades/precios; resultado “final” con costos incompletos.
- [ ] Si una limitación no se corrige, recortar el flujo y declararla. No presentar una función defectuosa o simulada como validada.
- [ ] Crear **`prototipo/`** con el recorrido consultable, versión identificada y notas de pruebas del responsable 2.
- [ ] Preparar Canva con resultados y fuentes de los demás; exportar slides a PDF.
- [ ] Ensayar **10 minutos de pitch y 8 de preguntas** con todos.
- [ ] Tener respaldo del recorrido si falla Figma o internet.
- [ ] Comprobar etiqueta **`pep1`**, contenido, acceso y correo de entrega por un integrante autorizado.

**Entrega:** prototipo entregable, slides PDF y entrega comprobada.

**No ampliar el producto ahora:** no hace falta implementar todo el listado de API, ML, directorio o un SaaS completo para esta PEP. Una primera entrega de una partida puede presentarse si se declara y se recorre de principio a fin.

Si el plazo pasó o la etiqueta fue entregada, consultar al profesor antes de modificar la versión evaluable. El informe posterior no cambia automáticamente el contenido de `pep1`.

---

## Orden de prioridad

1. **Contactar usuarios y realizar pruebas/compromisos:** 2 y 3, con apoyo de quienes tienen los contactos. Depende de disponibilidad externa.
2. **Completar evidencia, números y decisiones en paralelo:** 1, 3 y 4. El 6 estabiliza el recorrido para las sesiones.
3. **Integrar informe, slides y entrega:** 5 y 6 reciben resultados de los demás; pueden preparar estructura desde el inicio.

**No agregar más desarrollo salvo lo necesario para que la prueba y el recorrido sean honestos y consistentes.**

### Para todos

- Poder explicar problema, evidencia, experimento, precio y alcance, aunque no sean su bloque.
- Revisar el trabajo de otro integrante y ensayar la defensa.
- Completar coevaluación y autoevaluación en plazo.
- Mantener fechas reales, fuentes e hipótesis identificadas.

### Fechas originales de la guía

- **Lunes 9 de noviembre, 23:59:** repositorio con `pep1`, slides PDF, prototipo en `prototipo/` y enlace enviado al correo del curso.
- **Martes 10:** defensa presencial; todos presentes.
- **Jueves 12, 23:59:** informe PDF; formularios individuales en el plazo indicado por el curso.

Si las fechas ya pasaron, confirmar qué admite el profesor fuera del corte, sin asumir autorización ni alterar silenciosamente lo entregado.
