# Appointments API

API de agendamentos em Go, construída como projeto de estudo com foco em **arquitetura**: DDD, SOLID e Repository Pattern. O objetivo aqui é entender como estruturar as camadas — performance e otimizações ficam para uma etapa posterior.

## Stack

- **Go** 1.27
- **[go-chi/chi](https://github.com/go-chi/chi)** v5 — roteamento e middlewares
- Persistência planejada: **PostgreSQL** e **Redis** (ainda em stubs)
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
        │   └── create_user_handler.go
        ├── adapters/           # mapeamento DTO <-> domínio
        ├── models/             # entidade de domínio (User)
        ├── services/           # regras de negócio (a implementar)
        └── repositories/       # persistência (a implementar)
```

### Fluxo de uma request

```
HTTP body (JSON)
   │  json.Decode
   ▼
CreateUserIn (DTO da camada HTTP)
   │  adapter (ToUser)
   ▼
User (entidade de domínio)
   │
   ▼
service  ──►  repository (postgres/redis)
```

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

## Endpoints

### `POST /users/create`

Cria um usuário (atualmente ecoa o input; persistência a implementar).

```bash
curl -i -X POST http://localhost:8080/users/create \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "id": "1",
    "cpf": "111.222.333-44",
    "email": "johndoe@test.com",
    "password": "my-secret-password"
  }'
```

> `id` trafega como string por causa da tag `json:"id,string"`.

## Próximos passos

- [ ] Finalizar o adapter `CreateUserIn -> User` (método `ToUser()` no DTO)
- [ ] Corrigir `package` de `models/user.go` e remover a duplicação da entidade `User`
- [ ] Implementar a camada `services` com hashing de senha (`golang.org/x/crypto/bcrypt`)
- [ ] Definir a interface `UserService` no lado consumidor (handlers)
- [ ] Implementar `repositories` com Postgres/Redis reais
- [ ] Injeção de dependência: `service -> handler -> router -> shared`
- [ ] Configurar `docker-compose.yaml` (Postgres + Redis)

## Notas

Projeto de estudo — algumas partes ainda são stubs (`RunDB` apenas loga; handlers ecoam o input). A prioridade atual é a arquitetura e o fluxo entre camadas.
