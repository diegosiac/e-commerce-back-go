# E-commerce Backend

![Go](https://img.shields.io/badge/Go-1.22-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-blue)
![GraphQL](https://img.shields.io/badge/GraphQL-1.0.0-pink)
![Docker](https://img.shields.io/badge/Docker-26.1.3-blue)
![Stripe](https://img.shields.io/badge/Stripe-API-lightgrey)

Este proyecto es un backend para una aplicación de e-commerce, desarrollado en Go y utilizando microservicios. Emplea GraphQL para la API, PostgreSQL como base de datos, Docker para la contenedorización y Stripe para los pagos.

## Características

- **Microservicios**: Arquitectura basada en microservicios para escalabilidad y mantenibilidad.
- **GraphQL**: API desarrollada con GraphQL para consultas eficientes y flexibles.
- **PostgreSQL**: Base de datos relacional potente y fiable.
- **Docker**: Contenedorización para un despliegue sencillo y consistente.
- **Stripe**: Integración con Stripe para la gestión de pagos.

## Requisitos

- [Go 1.16+](https://golang.org/dl/)
- [Docker 20.10+](https://www.docker.com/products/docker-desktop)
- [PostgreSQL 16+](https://www.postgresql.org/download/)
- [Stripe API Key](https://stripe.com/docs/keys)

## Instalación

1. **Clonar el repositorio**

    ```bash
    git clone https://github.com/diegosiac/e-commerce-back-go
    cd e-commerce-back-go
    ```

2. **Configurar variables de entorno**

    Crear un archivo `.env` en el directorio raíz del proyecto.

3. **Iniciar servicios con Docker**

    ```bash
    docker compose up --build -d
    ```

4. **Ejecutar migraciones de base de datos**

    ```bash
    go run cmd/migrations/main.go
    ```

5. **Iniciar el servidor**

    ```bash
    go run cmd/server/main.go
    ```

## Uso

El servidor estará disponible en `http://localhost:8080`. Puedes acceder a la interfaz de GraphQL en `http://localhost:8080/graphql`.

:)
