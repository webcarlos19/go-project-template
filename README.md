# Go Project Template

Um template de projeto Go seguindo as melhores práticas adotadas por grandes empresas como Google, Uber e HashiCorp, com estrutura limpa e escalável.

## 🚀 Características

- **Estrutura de projeto limpa e escalável** seguindo padrões da comunidade Go
- **Configuração via Viper** com suporte a arquivos YAML e variáveis de ambiente
- **Logging estruturado** com Zap para alta performance
- **Conexão com PostgreSQL** usando pgx (driver mais performático)
- **API REST** com roteamento via Gorilla Mux
- **Testes unitários** com testify e mocks
- **Docker e Docker Compose** para desenvolvimento e produção
- **Makefile** com comandos úteis para desenvolvimento
- **Scripts de automação** para setup e deploy
- **Graceful shutdown** para encerramento seguro da aplicação

## 📁 Estrutura do Projeto

```
go-project-template/
├── cmd/                          # Pontos de entrada da aplicação
│   └── go-project-template/
│       └── main.go              # Arquivo principal
├── internal/                     # Código privado da aplicação
│   ├── config/                  # Configurações da aplicação
│   ├── handlers/                # Handlers HTTP
│   ├── models/                  # Estruturas de dados
│   ├── repository/              # Camada de acesso a dados
│   └── service/                 # Lógica de negócio
├── pkg/                         # Código reutilizável
│   ├── database/               # Configuração do banco de dados
│   └── logger/                 # Configuração do logger
├── api/                         # Definições de API
│   └── routes/                 # Definição de rotas
├── configs/                     # Arquivos de configuração
│   ├── config.yaml             # Configuração de desenvolvimento
│   ├── config.test.yaml        # Configuração de teste
│   └── config.prod.yaml        # Configuração de produção
├── build/                       # Scripts de build e Dockerfile
├── scripts/                     # Scripts de automação
├── test/                        # Testes de integração
├── docker-compose.yml          # Configuração do Docker Compose
├── Makefile                    # Comandos de automação
└── README.md                   # Este arquivo
```

### Explicação dos Diretórios

- **`cmd/`**: Contém os pontos de entrada da aplicação. Cada subdiretório representa um executável diferente.
- **`internal/`**: Código privado da aplicação que não deve ser importado por outros projetos.
- **`pkg/`**: Código que pode ser reutilizado por outros projetos ou aplicações.
- **`api/`**: Definições relacionadas à API (rotas, middlewares, documentação).
- **`configs/`**: Arquivos de configuração para diferentes ambientes.
- **`build/`**: Scripts de build, Dockerfiles e configurações de CI/CD.
- **`scripts/`**: Scripts auxiliares para desenvolvimento, teste e deploy.
- **`test/`**: Testes de integração e arquivos de teste que não ficam junto ao código.

## 🛠 Tecnologias Utilizadas

- **[Go 1.21+](https://golang.org/)** - Linguagem de programação
- **[Gorilla Mux](https://github.com/gorilla/mux)** - Roteador HTTP
- **[Viper](https://github.com/spf13/viper)** - Gerenciamento de configuração
- **[Zap](https://github.com/uber-go/zap)** - Logger estruturado de alta performance
- **[pgx](https://github.com/jackc/pgx)** - Driver PostgreSQL nativo
- **[Testify](https://github.com/stretchr/testify)** - Framework de testes
- **[Docker](https://www.docker.com/)** - Containerização
- **[PostgreSQL](https://www.postgresql.org/)** - Banco de dados relacional
- **[Redis](https://redis.io/)** - Cache (opcional)

## 🚀 Início Rápido

### Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose (opcional)
- Make (opcional, mas recomendado)

### Instalação

1. **Clone o repositório:**
   ```bash
   git clone <repository-url>
   cd go-project-template
   ```

2. **Execute o script de setup:**
   ```bash
   chmod +x scripts/setup.sh
   ./scripts/setup.sh
   ```

3. **Configure as variáveis de ambiente:**
   ```bash
   cp .env.example .env
   # Edite o arquivo .env com suas configurações
   ```

4. **Instale as dependências:**
   ```bash
   make deps
   ```

### Execução

#### Desenvolvimento Local

```bash
# Executar a aplicação
make run

# Executar com live reload (se tiver o Air instalado)
air

# Executar em modo de teste
make run-test

# Executar em modo de produção
make run-prod
```

#### Usando Docker

```bash
# Build da imagem
make docker-build

# Executar container
make docker-run

# Usar Docker Compose (inclui PostgreSQL e Redis)
make docker-compose-up
```

### Endpoints Disponíveis

- **GET /health** - Health check da aplicação
- **GET /** - Endpoint raiz com informações da API
- **GET /api/v1/users** - Lista todos os usuários (mock)
- **POST /api/v1/users** - Cria um novo usuário
- **GET /api/v1/users/{id}** - Busca usuário por ID
- **PUT /api/v1/users/{id}** - Atualiza usuário
- **DELETE /api/v1/users/{id}** - Remove usuário

### Exemplo de Requisição

```bash
# Health check
curl http://localhost:8080/health

# Listar usuários
curl http://localhost:8080/api/v1/users

# Criar usuário
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Carlos dos Santos", "email": "carlos@example.com"}'
```

## 🧪 Testes

```bash
# Executar todos os testes
make test

# Executar testes com coverage
make test-coverage

# Executar testes com race detection
make test-race

# Executar benchmarks
make benchmark
```

## 🏗 Build e Deploy

### Build Local

```bash
# Build para Linux
make build

# Build para Windows
make build-windows

# Build para macOS
make build-macos
```

### Docker

```bash
# Build da imagem Docker
make docker-build

# Deploy com Docker Compose
make docker-compose-up
```

### Produção

O projeto está configurado para deploy em ambiente de produção com:

- **Dockerfile multi-stage** para imagens otimizadas
- **Health checks** configurados
- **Usuário não-root** para segurança
- **Graceful shutdown** para encerramento seguro
- **Configuração via variáveis de ambiente**

## 📊 Qualidade de Código

```bash
# Linting
make lint

# Formatação
make format

# Verificação com go vet
make vet

# Análise de segurança
make security-scan
```

## 🔧 Configuração

### Variáveis de Ambiente

| Variável | Descrição | Padrão |
|----------|-----------|---------|
| `APP_ENV` | Ambiente da aplicação | `development` |
| `SERVER_PORT` | Porta do servidor | `8080` |
| `DB_HOST` | Host do banco de dados | `localhost` |
| `DB_PORT` | Porta do banco de dados | `5432` |
| `DB_USER` | Usuário do banco | `postgres` |
| `DB_PASSWORD` | Senha do banco | `postgres` |
| `DB_NAME` | Nome do banco | `go_project_template` |
| `LOG_LEVEL` | Nível de log | `info` |

### Arquivos de Configuração

- `configs/config.yaml` - Desenvolvimento
- `configs/config.test.yaml` - Testes
- `configs/config.prod.yaml` - Produção

## 📝 Padrões de Desenvolvimento

### Commits

Utilizamos o padrão [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: adiciona nova funcionalidade
fix: corrige bug
docs: atualiza documentação
style: formatação de código
refactor: refatoração sem mudança de funcionalidade
test: adiciona ou modifica testes
chore: tarefas de manutenção
```

### Branches

- `main` - Branch principal (produção)
- `develop` - Branch de desenvolvimento
- `feature/nome-da-feature` - Novas funcionalidades
- `fix/nome-do-fix` - Correções
- `hotfix/nome-do-hotfix` - Correções urgentes

### Code Review

1. Crie um branch para sua feature
2. Faça commits seguindo o padrão
3. Execute os testes: `make test`
4. Execute o linting: `make lint`
5. Abra um Pull Request para `develop`

## 🤝 Como Contribuir

1. **Fork** o projeto
2. Crie uma **branch** para sua feature (`git checkout -b feature/nova-feature`)
3. **Commit** suas mudanças (`git commit -m 'feat: adiciona nova feature'`)
4. **Push** para a branch (`git push origin feature/nova-feature`)
5. Abra um **Pull Request**

### Diretrizes

- Siga os padrões de código Go
- Escreva testes para novas funcionalidades
- Mantenha a documentação atualizada
- Use commits semânticos

## 📚 Recursos Adicionais

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👥 Autores

- **Carlos Eduardo Fernandes dos Santos** - [GitHub](https://github.com/webcarlos19)

## 🐛 Reportar Bugs

Encontrou um bug? Por favor, abra uma [issue](issues) com:

- Descrição detalhada do problema
- Passos para reproduzir
- Comportamento esperado vs atual
- Versão do Go e sistema operacional

## 💡 Solicitação de Features

Tem uma ideia? Abra uma [issue](issues) com:

- Descrição da feature
- Caso de uso
- Benefícios esperados

---

**Happy Coding! 🚀**
