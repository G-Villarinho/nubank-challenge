# 🟪 Desafio Backend - Nubank

Este projeto faz parte de um processo seletivo para a vaga de Desenvolvedor Backend no Nubank. A proposta era construir uma API REST para gerenciamento de **clientes e seus contatos**, com foco em boas práticas, relacionamento entre entidades e persistência de dados.

> 💡 Apesar da sugestão ser em Spring Boot + PostgreSQL, este desafio foi implementado utilizando **Go (Golang)**, explorando toda sua performance e estrutura limpa de APIs.

---

## ⚙️ Tecnologias utilizadas

- Go 1.24+
- PostgreSQL
- GORM (ORM)
- Echo Framework
- Docker & Docker Compose
- Swagger (documentação automática)
- Mockery + Testify (testes)
- Makefile (scripts automatizados)

---

## 📌 Requisitos Atendidos

- ✅ Cadastro de Cliente: `POST /clients`
- ✅ Cadastro de Contato (vinculado a um cliente): `POST /contacts`
- ✅ Listagem de todos os clientes com seus contatos: `GET /clients`
- ✅ Listagem dos contatos de um cliente específico: `GET /clients/{id}/contacts`

---

## 🚀 Como rodar o projeto

1. **Clone o repositório**

```bash
git clone https://github.com/G-Villarinho/nubank-challenge.git
cd nubank-challenge
```
2. **Instale as dependências de desenvolvimento**

```bash
make setup
```


3. **Configure o arquivo .env.local**
```bash
ENV=DEV

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=nu_user
POSTGRES_PASSWORD=nu_@123
POSTGRES_NAME=NUBANK_DEV
POSTGRES_SSL_MODE=disable
POSTGRES_MAX_CONN=10
POSTGRES_MAX_IDLE=5
POSTGRES_MAX_LIFE_TIME=1800
POSTGRES_TIMEOUT=3
```

4. **Suba o banco com Docker**
```bash
$ make docker-run
```

5. **Rode as migrations**
```bash
$ make migrations
```

6. **Inicie a aplicação**
```bash
$ make run
```

7. **Acesse a documentação Swagger (copie e cole no seu navegado)**
```bash
http://localhost:8080/swagger/index.html
```

## ✅ Testes
```bash
make test
```
**Isso gera:**
- Execução completa dos testes com cobertura
- Relatório HTML interativo (coverage.html)

# 📁 Estrutura do Projeto
```bash
.
├── handlers        # Controllers / rotas
├── models          # Entidades + Payloads
├── services        # Lógica de negócio
├── repositories    # Repositórios (GORM)
├── configs         # Carregamento de .env
├── mocks           # Mocks gerados com mockery
├── docs            # Swagger
├── pkgs            # Container de dependências helpers (injeção de dependência)
├── migrations      # Scripts de migração
├── storages        # Conexões com banco
├── Makefile        # Scripts de automação
└── main.go
```

# ✨  Diferenciais

- Utilização de Swagger para documentação automática

- Estrutura pensada para testabilidade e escalabilidade

- Cobertura de testes com cenários reais e mocks

- Separação clara de camadas com injeção de dependência

