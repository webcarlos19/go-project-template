# Changelog

Todas as mudanças notáveis neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
e este projeto adere ao [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-08-08

### Added
- Estrutura inicial do projeto seguindo padrões Go
- Configuração via Viper com suporte a YAML e variáveis de ambiente
- Logger estruturado com Zap
- API REST com Gorilla Mux
- Conexão com PostgreSQL usando pgx
- Dockerfile multi-stage para otimização
- Docker Compose para desenvolvimento
- Makefile com comandos de automação
- Scripts de setup e inicialização
- Testes unitários com testify
- CI/CD pipeline com GitHub Actions
- Endpoints de exemplo:
  - GET /health - Health check
  - GET /api/v1/users - Listar usuários (mock)
  - POST /api/v1/users - Criar usuário
  - GET /api/v1/users/{id} - Buscar usuário por ID
  - PUT /api/v1/users/{id} - Atualizar usuário
  - DELETE /api/v1/users/{id} - Deletar usuário
- Middleware de logging e CORS
- Graceful shutdown
- Configuração de linting com golangci-lint
- Documentação completa no README
- Exemplos de uso da API
- Licença MIT

### Documentation
- README.md com instruções detalhadas
- API_EXAMPLES.md com exemplos de uso
- Comentários em código seguindo padrões Go
- Documentação inline em todas as funções públicas

### Infrastructure
- GitHub Actions para CI/CD
- Docker configuration para produção
- Configuração de development environment
- Scripts de automação
