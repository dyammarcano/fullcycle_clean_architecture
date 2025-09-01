# Arquitectura y Separación de Capas

Este proyecto sigue los principios de Clean Architecture / Hexagonal (Ports and Adapters), promoviendo:
- Independencia del framework (HTTP/GRPC/DB reemplazables)
- Independencia de UI/Transporte (JSON/Protobuf/GraphQL)
- Independencia de base de datos (Memoria/Postgres)
- Testabilidad y bajo acoplamiento

## Capas y Dependencias (de afuera hacia adentro)

- cmd
  - Puntos de entrada (CLI/cmd). Hace composición de dependencias: crea repositorio concreto, crea use case, crea servidor (HTTP/GRPC). No contiene lógica de negocio.
- internal/adapter
  - http, grpc: Adaptadores de entrada/salida. Transforman I/O (JSON/Protobuf/GraphQL) en tipos de dominio y llaman a casos de uso. No conocen detalles de persistencia.
- internal/usecase
  - Orquesta reglas de aplicación. Depende solo de internal/domain (puerto OrderRepository). No conoce transporte ni base de datos.
- internal/domain
  - Núcleo: entidades y puertos (interfaces). Cero dependencias con el exterior.
- internal/repository
  - Implementaciones de puertos (adapters de infraestructura): memoria y Postgres. Pueden usar config, migraciones, drivers. No exponen detalles de infraestructura hacia adentro.
- pkg
  - Utilidades transversales (config, logger, util, grpc/pb). Usadas por capas externas. No deben ser requeridas por domain/usecase.

## Reglas de dependencia
- domain no importa nada.
- usecase importa solo domain.
- adapters (http/grpc) importan usecase (y util/logger) y traducen datos.
- repositories implementan interfaces de domain y pueden depender de infra (sql/migraciones/config).
- cmd ensambla todo y decide implementaciones concretas.

## Puntos revisados
- HTTP adapter deserializa JSON a domain.Order y llama al usecase. Correcto: el transporte no contamina el usecase.
- gRPC adapter convierte domain.Order a pb.Order. Se corrigió un detalle en la creación del slice para evitar entradas nulas (ver cambios). ✓
- Usecase solo usa el puerto OrderRepository y el tipo Order del dominio. ✓
- Repositorios (memoria y Postgres) implementan el puerto del dominio. Migraciones y SQL quedan encapsulados. ✓
- Config y logger están fuera del núcleo y se usan en composición/adapters/infra. ✓

## Posibles mejoras futuras (no disruptivas)
- DTOs de transporte: crear structs específicos para requests/responses en HTTP/GraphQL para no acoplar JSON directamente a la entidad de dominio (actualmente aceptable, pero mejora aislamiento).
- Normalizar errores en internal/repository/errors.go (hay mensajes de "user" que no aplican a orders). Podría moverse a un paquete de errores por contexto o mapear en adapters.
- Validación de dominio: agregar métodos/constructores en domain para invariantes (e.g., Amount > 0).

## Cambios realizados en esta revisión
- Corrección en internal/adapter/grpc/server.go: se ajustó la inicialización del slice para mapear orders sin producir entradas nulas:
  - Antes: make(len), luego append -> duplicaba longitud con nulos.
  - Ahora: make(0, len), luego append -> lista correcta.

Con estos ajustes, la separación de arquitectura se mantiene clara y consistente con Clean Architecture.