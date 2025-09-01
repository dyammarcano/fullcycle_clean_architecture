# Arquitetura e Separação de Camadas

Este projeto segue os princípios de Clean Architecture / Hexagonal (Ports and Adapters), promovendo:
- Independência de framework (HTTP/GRPC/BD substituíveis)
- Independência de UI/Transporte (JSON/Protobuf/GraphQL)
- Independência de banco de dados (Memória/Postgres)
- Testabilidade e baixo acoplamento

## Camadas e Dependências (de fora para dentro)

- cmd
  - Pontos de entrada (CLI/cmd). Faz a composição de dependências: cria o repositório concreto, cria o caso de uso e cria o servidor (HTTP/GRPC). Não contém lógica de negócio.
- internal/adapter
  - http, grpc: Adaptadores de entrada/saída. Transformam I/O (JSON/Protobuf/GraphQL) em tipos de domínio e chamam os casos de uso. Não conhecem detalhes de persistência.
- internal/usecase
  - Orquestra as regras de aplicação. Depende apenas de internal/domain (porta OrderRepository). Não conhece transporte nem banco de dados.
- internal/domain
  - Núcleo: entidades e portas (interfaces). Zero dependências com o exterior.
- internal/repository
  - Implementações das portas (adapters de infraestrutura): memória e Postgres. Podem usar config, migrações, drivers. Não expõem detalhes de infraestrutura para dentro.
- pkg
  - Utilidades transversais (config, logger, util, grpc/pb). Usadas pelas camadas externas. Não devem ser requeridas por domain/usecase.

## Regras de dependência
- domain não importa nada.
- usecase importa apenas domain.
- adapters (http/grpc) importam usecase (e util/logger) e fazem a tradução de dados.
- repositories implementam interfaces de domain e podem depender de infra (sql/migrações/config).
- cmd faz o assembly de tudo e decide implementações concretas.

## Pontos revisados
- O adapter HTTP desserializa JSON para domain.Order e chama o usecase. Correto: o transporte não contamina o usecase.
- O adapter gRPC converte domain.Order para pb.Order. Foi corrigido um detalhe na criação do slice para evitar entradas nulas (ver mudanças). ✓
- O usecase usa apenas a porta OrderRepository e o tipo Order do domínio. ✓
- Repositórios (memória e Postgres) implementam a porta do domínio. Migrações e SQL ficam encapsulados. ✓
- Config e logger estão fora do núcleo e são usados na composição/adapters/infra. ✓

## Possíveis melhorias futuras (não disruptivas)
- DTOs de transporte: criar structs específicos para requests/responses em HTTP/GraphQL para não acoplar o JSON diretamente à entidade de domínio (atualmente aceitável, mas melhora o isolamento).
- Normalizar erros em internal/repository/errors.go (há mensagens de "user" que não se aplicam a orders). Poderia ser movido para um pacote de erros por contexto ou mapeado nos adapters.
- Validação de domínio: adicionar métodos/construtores em domain para invariantes (ex.: Amount > 0).

## Mudanças realizadas nesta revisão
- Correção em internal/adapter/grpc/server.go: ajustada a inicialização do slice para mapear orders sem produzir entradas nulas:
  - Antes: make(len), depois append -> duplicava o comprimento com nulos.
  - Agora: make(0, len), depois append -> lista correta.

Com esses ajustes, a separação de arquitetura se mantém clara e consistente com Clean Architecture.