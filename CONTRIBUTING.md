# Contribuindo para o Go Project Template

Obrigado por seu interesse em contribuir! Este documento fornece diretrizes e informações sobre como contribuir para este projeto.

## 📋 Código de Conduta

Ao participar deste projeto, você concorda em manter um ambiente respeitoso e inclusivo. Seja gentil e respeitoso com outros contribuidores.

## 🚀 Como Contribuir

### Reportando Bugs

Antes de reportar um bug, verifique se ele já foi reportado nas [issues existentes](../../issues).

Ao reportar um bug, inclua:

- **Descrição clara e concisa** do bug
- **Passos para reproduzir** o comportamento
- **Comportamento esperado** vs comportamento atual
- **Screenshots** (se aplicável)
- **Informações do ambiente**:
  - Versão do Go
  - Sistema operacional
  - Versão do projeto

### Sugerindo Melhorias

Para sugerir uma nova funcionalidade:

1. Verifique se já existe uma issue para a funcionalidade
2. Abra uma nova issue com:
   - Descrição detalhada da funcionalidade
   - Justificativa para a adição
   - Exemplos de uso
   - Possível implementação (opcional)

### Processo de Desenvolvimento

1. **Fork** o repositório
2. **Clone** seu fork localmente
3. **Crie uma branch** para sua feature/fix
4. **Configure o ambiente** de desenvolvimento
5. **Faça suas alterações**
6. **Teste** suas alterações
7. **Commit** seguindo as convenções
8. **Push** para seu fork
9. **Abra um Pull Request**

### Configurando o Ambiente de Desenvolvimento

```bash
# Clone seu fork
git clone https://github.com/webcarlos19/go-project-template.git
cd go-project-template

# Execute o script de setup
./scripts/setup.sh

# Instale ferramentas de desenvolvimento
make dev-setup

# Execute os testes para verificar se tudo está funcionando
make test
```

### Padrões de Código

#### Estilo de Código

- Siga as convenções padrão do Go
- Use `gofmt` para formatação
- Execute `golangci-lint` antes de fazer commit
- Mantenha funções pequenas e focadas
- Use nomes descritivos para variáveis e funções

#### Comentários

- Documente todas as funções e tipos públicos
- Use comentários para explicar lógica complexa
- Mantenha comentários atualizados com o código

#### Testes

- Escreva testes para todas as novas funcionalidades
- Mantenha cobertura de testes alta (>80%)
- Use nomes descritivos para testes
- Agrupe testes relacionados em subtestes

### Convenções de Commit

Usamos [Conventional Commits](https://www.conventionalcommits.org/):

```
<tipo>[escopo opcional]: <descrição>

[corpo opcional]

[footer opcional]
```

#### Tipos

- `feat`: Nova funcionalidade
- `fix`: Correção de bug
- `docs`: Documentação
- `style`: Formatação de código
- `refactor`: Refatoração sem mudança de funcionalidade
- `test`: Testes
- `chore`: Tarefas de manutenção
- `ci`: Configuração de CI/CD
- `perf`: Melhoria de performance
- `build`: Mudanças no sistema de build

#### Exemplos

```bash
feat(api): adiciona endpoint para criação de usuários
fix(database): corrige connection pool leak
docs(readme): atualiza instruções de instalação
test(handlers): adiciona testes para user handlers
```

### Padrões de Branch

- `main` - Branch principal (produção)
- `develop` - Branch de desenvolvimento
- `feature/nome-da-feature` - Novas funcionalidades
- `fix/nome-do-fix` - Correções
- `hotfix/nome-do-hotfix` - Correções urgentes
- `docs/nome-da-doc` - Documentação

### Pull Request Guidelines

#### Antes de Abrir um PR

- [ ] Execute `make test` e certifique-se de que todos os testes passam
- [ ] Execute `make lint` e corrija todos os problemas
- [ ] Execute `make format` para formatar o código
- [ ] Atualize a documentação se necessário
- [ ] Adicione testes para novas funcionalidades

#### Template do Pull Request

```markdown
## Descrição

Breve descrição das mudanças realizadas.

## Tipo de Mudança

- [ ] Bug fix (mudança que corrige um problema)
- [ ] Nova funcionalidade (mudança que adiciona funcionalidade)
- [ ] Breaking change (correção ou funcionalidade que quebra compatibilidade)
- [ ] Documentação

## Como Foi Testado

Descreva como você testou suas mudanças.

## Checklist

- [ ] Meu código segue o estilo do projeto
- [ ] Revisei meu próprio código
- [ ] Comentei partes complexas do código
- [ ] Fiz mudanças correspondentes na documentação
- [ ] Minhas mudanças não geram novos warnings
- [ ] Adicionei testes que provam que a correção é efetiva ou que a funcionalidade funciona
- [ ] Testes novos e existentes passam localmente
```

### Revisão de Código

Todo Pull Request passa por revisão. Esperamos:

- **Feedback construtivo** nos comentários
- **Aprovação** de pelo menos um mantenedor
- **Todos os checks** do CI passando
- **Resolução** de todos os comentários

### Documentação

#### Documentando APIs

```go
// CreateUser creates a new user in the system.
// It validates the input data and returns the created user or an error.
//
// Parameters:
//   - ctx: Request context for cancellation and timeouts
//   - req: User creation request containing name and email
//
// Returns:
//   - *UserResponse: The created user data
//   - error: Any error that occurred during creation
func (s *userService) CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
    // Implementation...
}
```

#### README

- Mantenha o README atualizado
- Inclua exemplos de uso para novas funcionalidades
- Documente novas variáveis de ambiente
- Atualize a seção de instalação se necessário

### Releases

Releases são criados automaticamente quando:

1. Um PR é mergeado na branch `main`
2. A versão no `go.mod` é atualizada
3. Um tag é criado seguindo semantic versioning

### Ferramentas Recomendadas

#### Editor/IDE

- **VS Code** com extensão Go
- **GoLand** (JetBrains)
- **Vim/Neovim** com plugins Go

#### Ferramentas de Linha de Comando

```bash
# Ferramentas essenciais
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest

# Ferramentas opcionais para desenvolvimento
go install github.com/cosmtrek/air@latest  # Live reload
go install github.com/swaggo/swag/cmd/swag@latest  # Swagger docs
```

### Recursos Úteis

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

### Dúvidas?

Se tiver dúvidas sobre como contribuir:

1. Verifique a [documentação](README.md)
2. Procure em [issues existentes](../../issues)
3. Abra uma nova issue com a tag `question`
4. Entre em contato com os mantenedores

---

**Obrigado por contribuir! 🚀**
