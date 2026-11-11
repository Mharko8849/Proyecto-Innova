# Guía para abrir el prototipo

## Acceso

El archivo [index.html](index.html) incluye la aplicación, React, estilos y tipografía. No necesita instalación, cuenta de Figma ni conexión a internet.

1. Descargar la carpeta del prototipo o clonar el repositorio que la contiene.
2. Abrir `index.html` en Chrome, navegador en el que se comprobó la apertura como archivo local.
3. Elegir **«Vidriería»** en **«Explorar por rubro»** para entrar con datos ficticios. También hay accesos demostrativos de Electricidad y Construcción.

GitHub muestra el código del HTML, pero no ejecuta la aplicación en su visor. Para utilizarla hay que descargarla y abrirla en el navegador. El [recorrido con capturas](README.md) sí puede consultarse directamente en GitHub.

## Datos de prueba y reinicio

Los cambios se guardan en el almacenamiento local del navegador: no modifican el HTML ni se suben a Git. No introducir contraseñas reales ni datos privados de clientes; la sesión de la maqueta no es autenticación de producción.

Para borrar los datos del prototipo y recuperar su estado inicial, abrirlo en Chrome, pulsar **F12**, entrar en **Consola** y ejecutar:

```js
localStorage.removeItem("indicepartidas-v3");
location.reload();
```

Esto elimina las cotizaciones, órdenes y demás cambios locales del prototipo en ese acceso. Recargar normalmente o cerrar sesión no los borra.

## Alternativa: servidor local

Si el navegador restringe el almacenamiento de archivos locales, ejecutar desde la raíz del repositorio que contenga esta carpeta:

```bash
python3 -m http.server 8000 --bind 127.0.0.1 --directory prototipo
```

Después, abrir `http://127.0.0.1:8000/`. Los datos guardados mediante este acceso son independientes de los utilizados al abrir el archivo con doble clic.

## Versión de la exportación

**Versión entregable: PEP1-HTML-01.** SHA-256 de `index.html`:

```text
7d0af18e683e645a0ca3db82f427f8d40fc7ccfcd6b31aa40d321bff55caf3c4
```

El HTML incluye también los hashes de sus archivos fuente. Este identificador corresponde a la exportación entregada, no acredita cuál fue la versión mostrada a P-01 y P-02.

La exportación conserva el comportamiento de la copia de `ÍndicePartidas/` utilizada para generarla; no se sincroniza automáticamente con cambios posteriores de Figma ni corrige las limitaciones descritas en el README.

## Comprobación técnica

En Chrome, como archivo local y sin conexión, se comprobó acceso demostrativo, creación de una cotización de dos partidas, conservación de condiciones comerciales, documento comercial sin desglose de costos internos, aceptación manual, creación de una OT, registro de un monto de costo y persistencia al recargar. Ese recorrido no produjo errores JavaScript ni solicitudes HTTP externas.

Esta comprobación no es una entrevista con usuarios ni certifica todas las funciones o cálculos del prototipo. Los resultados de validación del equipo están en [pruebas-solucion.md](pruebas-solucion.md).
