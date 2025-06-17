# 📄 Solicitud de Préstamo API

Este proyecto expone una API RESTful para gestionar el flujo de solicitudes de préstamo, implementado en Go y listo para ejecutarse en un entorno Dockerizado. Utiliza Postman para pruebas e incluye integración con Sentry para el seguimiento de errores.

---

## ⚙️ Requisitos

Asegúrate de tener instalados:

- [Docker](https://www.docker.com/)
- [Postman](https://www.postman.com/) u otra herramienta compatible con el formato de colecciones de Postman

---

## 🔐 Variables de Entorno

Debes crear un archivo `.env` con las siguientes variables:

```env
SENTRY=https://aad80be270d9aa0a5b3c9b40a46bff31@o4509387464179712.ingest.us.sentry.io/4509387467653120
DB_PATH=data/solicitudesPrestamos.db


 Ubicación del archivo .env

Este archivo debe estar ubicado en:

${HOME}/environments/solicitud-prestamo-env/.env

1. Construcción y despliegue del contenedor

Desde el directorio raíz del proyecto, ejecuta:

./scripts/run.sh


2. Verificación del estado del contenedor

Puedes verificar si el contenedor se levantó correctamente con:

docker logs -f solicitud-prestamo-container


🧪 Pruebas con Postman
3. Importar archivos en Postman

Para realizar pruebas, importa los siguientes archivos en Postman:

    🧪 Archivo de entorno (SOLICITUD-PRESTAMO-LOCAL.postman_environment.json)

    📬 Colección de pruebas (API-SOLICITUD-PRESTAMO.postman_collection.json)

    Asegúrate de seleccionar el entorno correcto al momento de ejecutar cada request.


4. Ejecutar los endpoints

A continuación, se listan los endpoints disponibles:

| Método | Endpoint                      | Descripción                                           |
| ------ | ----------------------------- | ----------------------------------------------------- |
| GET    | `/mutant`                     | Verifica la salud del contenedor (endpoint de prueba) |
| POST   | `/iniciar-solicitud`          | Inicia una nueva solicitud de préstamo                |
| POST   | `/consultar-score`            | Consulta el score crediticio del solicitante          |
| POST   | `/verificar-identidad`        | Verifica la identidad del solicitante                 |
| POST   | `/consultar-estado-solicitud` | Consulta el estado actual de una solicitud            |




🔄 Flujo de Uso

    Iniciar solicitud:
    Ejecuta /iniciar-solicitud y obtén el uuid o documento de identidad generado.

    Consultar score:
    Usa el uuid recibido para llamar a /consultar-score.

    Verificar identidad:
    Utiliza los mismos datos para llamar a /verificar-identidad.

    Consultar estado de la solicitud:
    Consulta el estado actual con /consultar-estado-solicitud, usando nuevamente el uuid o el documento.

    ⚠️ Nota: Los endpoints dependen de los datos generados en pasos anteriores. Asegúrate de mantener el flujo correcto entre ellos.




