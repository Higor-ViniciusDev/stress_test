# Stress CLI

Ferramenta de teste de carga para HTTP. Realiza stress tests configurando requisições totais e concorrência simultânea.

## Como Usar

```bash
docker run higorsouzadev/stess-cli:latest -url <URL> -requests <N> -concurrency <N>
```

**Flags:**
- `-url`: URL a testar
- `-requests`: Total de requisições
- `-concurrency`: Requisições simultâneas

**Exemplo:**
```bash
docker run higorsouzadev/stess-cli:latest -url https://api.example.com -requests 1000 -concurrency 100
```

Retorna relatório com métricas de desempenho (tempo total, sucessos/falhas, latência).