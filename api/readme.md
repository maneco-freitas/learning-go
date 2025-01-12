# Desenvolvendo API REST com Go

## 📚 Sobre este Projeto

Este projeto é parte do meu estudo da linguagem Go. Como desenvolvedor fullstack com mais de 10 anos de experiência em Java, Kotlin, Python e Node.js, decidi expandir meu conhecimento aprendendo Go através de uma abordagem prática, implementando uma API REST completa.

## 🛠 Tecnologias Utilizadas

- **Gin**: Framework web para Go
- **GORM**: ORM para Go
- **SQLite**: Banco de dados
- **Testify**: Framework de testes
- **GoDotEnv**: Gerenciamento de variáveis de ambiente

## 📁 Estrutura do Projeto

```
api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   ├── product.go
│   │   └── product_test.go
│   ├── model/
│   │   └── product.go
│   ├── repository/
│   │   ├── interface.go
│   │   └── product.go
│   │   └── product_test.go
│   └── database/
│       └── database.go
├── .env
├── go.mod
└── README.md
```

## 🚀 Como Executar

1. Clone o repositório
```bash
git clone [URL_DO_REPOSITORIO]
cd api
```

2. Configure o arquivo `.env`:
```env
PORT=8080
GIN_MODE=debug
DB_TYPE=sqlite
DB_PATH=products.db
```

3. Execute o projeto:
```bash
go run cmd/api/main.go
```

## 📝 Endpoints

### Produtos

- `GET /products`: Lista todos os produtos
  ```bash
  curl http://localhost:8080/products
  ```

- `GET /products/:id`: Busca um produto por ID
  ```bash
  curl http://localhost:8080/products/1
  ```

- `POST /products`: Cria um novo produto
  ```bash
  curl -X POST http://localhost:8080/products \
    -H "Content-Type: application/json" \
    -d '{"name":"Laptop","price":999.99}'
  ```

- `PUT /products/:id`: Atualiza um produto
  ```bash
  curl -X PUT http://localhost:8080/products/1 \
    -H "Content-Type: application/json" \
    -d '{"name":"Gaming Laptop","price":1299.99}'
  ```

- `DELETE /products/:id`: Remove um produto
  ```bash
  curl -X DELETE http://localhost:8080/products/1
  ```

## 🧪 Testes

O projeto inclui testes unitários usando `testify`. Para executar:

```bash
# Roda todos os testes
go test ./...

# Roda testes com cobertura
go test -cover ./...

# Roda testes em modo verboso
go test -v ./...
```

### Exemplos de Testes Implementados:
- Testes de Handler com mock do repository
- Testes de Repository usando SQLite em memória
- Testes de configuração

## 📚 Conceitos Implementados

1. **Estruturação de Projetos Go**
   - Organização de pacotes
   - Separação de responsabilidades
   - Clean Architecture

2. **Testes em Go**
   - Mocking com testify
   - Testes de integração
   - SQLite em memória para testes

3. **Configuração de Ambiente**
   - Variáveis de ambiente
   - Diferentes ambientes (dev, test, prod)
   - Gestão de configurações

4. **Padrões de Projeto**
   - Repository Pattern
   - Dependency Injection
   - Interface Segregation

## 🔄 Próximos Passos

- [ ] Implementar autenticação
- [ ] Adicionar logging estruturado
- [ ] Implementar cache
- [ ] Adicionar documentação com Swagger
- [ ] Containerização com Docker

## 📖 Recursos Úteis

- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [Testify Documentation](https://pkg.go.dev/github.com/stretchr/testify)
- [Go Project Layout](https://github.com/golang-standards/project-layout)

## 🤝 Contribuições

Contribuições são sempre bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---
Feito com ❤️ durante meus estudos de Go