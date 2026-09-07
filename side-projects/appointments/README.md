# Appointments API

API de agendamentos em Go, construída como projeto de estudo com foco em **arquitetura**: DDD, SOLID e Repository Pattern. O objetivo aqui é entender como estruturar as camadas — performance e otimizações ficam para uma etapa posterior.

## Stack

- **Go** 1.27
- **[go-chi/chi](https://github.com/go-chi/chi)** v5 — roteamento e middlewares
- **[pgx/v5](https://github.com/jackc/pgx)** — driver/pool de conexão PostgreSQL
- **[tern](https://github.com/jackc/tern)** — migrações de banco (SQL versionado)
- **[sqlc](https://sqlc.dev)** — geração de código type-safe a partir do SQL (escolhido sobre squirrel; squirrel só entra se surgir query dinâmica)
- **[google/uuid](https://github.com/google/uuid)** — identificadores UUID
- **PostgreSQL** (via Docker) — Redis ainda planejado (stub)
- Módulo: `github.com/luizandrends/appointments`

## Arquitetura

O projeto é organizado por **domínio** (`modules/`) em vez de por camada técnica no topo. Cada módulo tem suas próprias camadas internas, e o que é transversal fica em `shared/`.

```
appointments/
├── shared/                     # infraestrutura compartilhada
│   ├── api/main.go             # entrypoint: sobe deps + servidor HTTP
│   ├── routes/routes.go        # roteador principal (monta os módulos)
│   ├── dependencies/           # inicialização de bancos
│   │   ├── postgres/
│   │   └── redis/
│   └── utils/response.go       # ApiResponse[T] genérico + SendJSON
│
└── modules/
    └── users/                  # módulo de usuários
        ├── handlers/           # camada HTTP (entrega)
        │   ├── router.go       # monta o chi.Router do módulo
        │   ├── in.go / out.go  # DTOs de entrada/saída
        │   ├── adapters.go     # mapeamento DTO <-> domínio
        │   └── create_user_handler.go
        ├── models/             # entidade de domínio (User)
        ├── services/           # regras de negócio (em construção)
        └── repositories/       # persistência (em construção)
```

As migrações do banco ficam em `shared/dependencies/postgres/migrations/` (gerenciadas pelo tern).

### Fluxo de uma request

```
HTTP body (JSON)
   │  json.Decode
   ▼
CreateUserIn (DTO da camada HTTP)
   │  in.requestToUser()      (método no DTO)
   ▼
User (entidade de domínio)
   │
   ▼
service  ──►  repository (postgres)
   │
   ▼
User (domínio)
   │  userToResponse(user)    (função pura)
   ▼
CreateUserOut (DTO de saída)
```

> Regra de bolso aplicada nos adapters: conversão que **usa** o dado de origem vira método com receiver (`requestToUser`); conversão que **constrói** um novo valor vira função pura (`userToResponse`).

### Princípios adotados

- **Módulo devolve `chi.Router`**; o roteador principal agrega com `r.Mount("/users", ...)`.
- **Direção de dependência única**: camada de entrega (handlers) depende do domínio (models), nunca o contrário.
- **Utilitários genéricos são pacotes-folha**: `ApiResponse[T]` e `SendJSON` vivem em `shared/utils` (não importam módulos) para evitar import cycles.
- **Injeção de dependência**: o dado do usuário chega por request (`r.Body`); as dependências (service, repo) são injetadas na inicialização.
- **Respostas padronizadas** via envelope genérico `ApiResponse[T]` (`{ "data": ... }` / `{ "error": ... }`).

## Como rodar

```bash
go run ./shared/api/main.go
```

O servidor sobe em `:8080` com timeouts de leitura/escrita configurados e os middlewares `Recoverer`, `RequestID` e `Logger`.

## Banco de dados & Migrações

Postgres sobe via Docker; o `postgres.go` conecta com `pgxpool`.

```bash
docker compose up -d
```

Migrações com tern (binário em `$(go env GOPATH)/bin/tern`):

```bash
tern migrate \
  --config shared/dependencies/postgres/migrations/tern.conf \
  --migrations shared/dependencies/postgres/migrations
```

Migrações aplicadas:

- `001_create_table_users.sql` — cria a tabela `users`
- `002_change_id_to_uuid.sql` — troca a PK de `serial` para `uuid` (`default gen_random_uuid()`)

## Endpoints

### `POST /users/create`

Cria um usuário (fluxo de camadas montado; persistência via sqlc em construção).

```bash
curl -i -X POST http://localhost:8080/users/create \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "cpf": "111.222.333-44",
    "email": "johndoe@test.com",
    "password": "my-secret-password"
  }'
```

> O `id` **não** é enviado pelo cliente — é um UUID gerado pelo banco (`gen_random_uuid()`) e retornado na resposta.

## Próximos passos

- [x] Adapters `CreateUserIn -> User` (`requestToUser`) e `User -> CreateUserOut` (`userToResponse`)
- [x] Corrigir `package` de `models/user.go` e remover a duplicação da entidade `User`
- [x] Conexão real com Postgres (`pgxpool`) + migrações com tern
- [x] `id` como UUID gerado pelo banco
- [ ] Migração `003` — remover coluna `name` duplicada e adicionar `unique` em `email`/`cpf`
- [ ] `sqlc generate` — gerar o código de acesso (`db.Queries`) a partir do `query.sql`
- [ ] Implementar `repositories` (mapeando `db.User` <-> `models.User`)
- [ ] Camada `services` com struct + construtor e hashing de senha (`golang.org/x/crypto/bcrypt`)
- [ ] Definir interfaces `UserService` (em handlers) e `UserRepository` (em services) no lado consumidor
- [ ] Injeção de dependência via composition root no `main`: `repo -> service -> handler -> router -> shared`
- [ ] Tratamento de erro no handler (remover `panic`, padronizar via `ApiResponse`)
- [ ] Redis (planejado)

## Notas

Projeto de estudo — o foco atual é arquitetura e o fluxo entre camadas. A persistência real (sqlc + repositories) e a injeção de dependência completa estão em construção; o handler ainda ecoa/monta a resposta sem gravar no banco.
