# Exemplo de Concorrência em Go

## Sobre o Projeto
Este projeto demonstra um exemplo prático de concorrência em Go, implementando um sistema de processamento de períodos em paralelo com controle de concorrência, retry e timeout.

## Funcionalidades Implementadas

### Estruturas Principais
- `Period`: Estrutura para representar um período de tempo
- `Result`: Estrutura para armazenar resultados e erros das operações
- Sistema de processamento concorrente com número limitado de goroutines

### Controles de Concorrência
- **Semáforo**: Limita o número máximo de goroutines concorrentes
- **WaitGroup**: Sincroniza a conclusão de todas as goroutines
- **Context**: Gerencia timeouts e cancelamento de operações
- **Channels**: Comunica resultados entre goroutines

### Tratamento de Erros
- Sistema de retry com backoff exponencial
- Timeout global para todas as operações
- Tratamento de erros individuais por período

## Como Usar

```go
func main() {
    // Configurações
    maxConcurrent := 5    // Número máximo de goroutines
    maxRetries := 4       // Tentativas máximas por request
    
    // Inicialização
    sem := make(chan struct{}, maxConcurrent)
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    // Processamento dos períodos...
}
```

## Detalhes de Implementação

### Semáforo
```go
sem := make(chan struct{}, maxConcurrent)
// Antes de iniciar uma goroutine
sem <- struct{}{}
go func() {
    defer func() { <-sem }()
    // Processamento...
}()
```

### Retry com Backoff
```go
for attempt := 0; attempt < maxRetries; attempt++ {
    // Tenta operação
    if err == nil {
        return data, nil
    }
    // Espera com backoff exponencial
    time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
}
```

## Próximos Passos

### Melhorias Propostas
1. **Métricas**
   - Adicionar tracking de tempo de execução por request
   - Monitorar taxa de sucesso/falha

2. **Ordenação de Resultados**
   - Implementar ordenação dos resultados por período
   - Adicionar identificadores únicos para tracking

3. **Circuit Breaker**
   - Implementar mecanismo de circuit breaker
   - Adicionar monitoramento de falhas consecutivas

4. **Melhorias de Performance**
   - Otimizar backoff strategy
   - Implementar cache para resultados comuns

### Considerações de Produção
- Adicionar logs estruturados
- Implementar métricas para monitoramento
- Adicionar testes de carga
- Documentar padrões de erro

## Como Contribuir
1. Fork o repositório
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Crie um Pull Request