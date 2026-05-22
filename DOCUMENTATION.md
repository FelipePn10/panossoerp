# Módulos de Manufatura — Documentação

Cobre os módulos implementados neste ciclo:
**Roteiro de Fabricação · CRP · APS · Custo Padrão · Qualidade · Manutenção Preventiva · Previsão Estatística · Alertas MRP (e-mail + webhook)**

> Documentação fiscal em separado: **docs/FISCAL_FINANCEIRO.md**

---

## 1. Roteiro de Fabricação

### O que é

O roteiro descreve **como** um item é produzido: quais operações são executadas, em que sequência, em quais centros de trabalho, com quais tempos e quais dependências existem entre as etapas.

O roteiro é criado **manualmente** pelo PCP/engenharia de processo. O MRP, CRP e APS apenas o *leem* — nunca o criam nem o modificam.

### Estrutura de dados

```
operations                  ← biblioteca reutilizável de operações genéricas
  └─ manufacturing_routes   ← roteiro de um item específico
       └─ route_operations  ← instância de uma operação dentro do roteiro
            └─ route_operation_network  ← grafo de dependências entre operações
```

#### `operations` — biblioteca de operações

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `name` | string | Ex: "Solda TIG", "Pintura Eletrostática" |
| `origin` | enum | `INTERNA`, `EXTERNA`, `TERCEIROS` |
| `standard_time` | float64 | Tempo padrão em horas |

**O campo `origin` determina o tipo de ordem que o MRP gera:**

| Origin | Significado | Ordem gerada pelo MRP |
|--------|-------------|----------------------|
| `INTERNA` | Operação executada pelo próprio chão de fábrica | Ordem de Fabricação (OF) |
| `EXTERNA` | Operação enviada para fornecedor externo | Ordem de Serviço (OS) |
| `TERCEIROS` | Operação realizada por terceiros contratados | Ordem de Serviço (OS) |

Quando um item do tipo `FABRICACAO` possui operações com origin `EXTERNA` ou `TERCEIROS` no seu roteiro padrão, o MRP gera automaticamente **ordens de serviço adicionais** para cada uma dessas operações, além da ordem de fabricação principal.

#### `manufacturing_routes` — roteiro de um item

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `item_code` | int64 | Item ao qual o roteiro pertence |
| `is_standard` | bool | `TRUE` = roteiro usado pelo MRP/CRP; apenas um por item |
| `name` | string | Nome descritivo |

#### `route_operations` — operação dentro do roteiro

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `sequence` | int16 | Posição (ex: 10, 20, 30) |
| `operation_id` | int64 | FK para `operations` |
| `work_center_id` | int64? | Centro de trabalho (sobrescreve o padrão da operação) |
| `standard_time` | float64? | Tempo corrigido; se nulo, herda da operação |
| `setup_time` | float64? | Tempo de setup |
| `notes` | text? | Observações livres |

#### `route_operation_network` — grafo de dependências

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `predecessor_id` | int64 | Operação que deve terminar (ou estar suficientemente avançada) antes |
| `successor_id` | int64 | Operação que só pode iniciar após a predecessora |
| `overlap_pct` | float64 | % de sobreposição permitida (ver CPM abaixo) |

Operações sem sucessor simplesmente não aparecem como `predecessor_id` em nenhuma aresta. A última operação do roteiro não precisa de nenhum registro especial — ela contribui naturalmente para o cálculo de lead time.

---

### Como cadastrar um roteiro (passo a passo)

**Passo 1 — Criar operações genéricas (uma vez; ficam na biblioteca)**

```http
POST /api/routing/operations
{
  "name": "Corte a laser",
  "origin": "INTERNA",
  "standard_time": 0.5
}

POST /api/routing/operations
{
  "name": "Pintura eletrostática",
  "origin": "EXTERNA",
  "standard_time": 2.0
}
```

Operações criadas uma única vez e reutilizadas em múltiplos roteiros.

---

**Passo 2 — Criar o roteiro do item**

```http
POST /api/routing/routes
{
  "item_code": 1001,
  "name": "Roteiro Padrão – Produto X",
  "is_standard": true
}
→ { "id": 7, "item_code": 1001, "is_standard": true, ... }
```

---

**Passo 3 — Adicionar operações ao roteiro**

```http
POST /api/routing/route-operations/7
{ "operation_id": 1, "sequence": 10, "work_center_id": 2, "standard_time": 0.5 }

POST /api/routing/route-operations/7
{ "operation_id": 3, "sequence": 20, "work_center_id": 4, "standard_time": 1.5, "setup_time": 0.25 }

POST /api/routing/route-operations/7
{ "operation_id": 5, "sequence": 30, "work_center_id": 2, "standard_time": 0.5 }
```

---

**Passo 4 — Definir dependências entre operações**

```http
POST /api/routing/routes/7/edges
{
  "predecessor_id": 10,
  "successor_id": 20,
  "overlap_pct": 0.0
}

POST /api/routing/routes/7/edges
{
  "predecessor_id": 20,
  "successor_id": 30,
  "overlap_pct": 0.20
}
```

A operação 30 pode iniciar quando a 20 estiver 80% concluída (`overlap_pct = 0.20`).
A operação 30 não tem sucessor — ela é a última do roteiro e não precisa de nenhum registro adicional.

---

### Lead Time via CPM (Critical Path Method)

**CPM** é uma técnica de engenharia para calcular o **tempo mínimo** de um conjunto de atividades com dependências. Independente de quantas operações existam, o tempo total é determinado pelo **caminho mais longo** (caminho crítico). Operações em paralelo não somam; as em série somam.

#### Conceitos

- **`early_start[op]`** — o mais cedo que a operação pode iniciar
- **`early_finish[op]`** — o mais cedo que ela pode terminar
- **`overlap_pct`** — fração do predecessor que o sucessor pode "pular" ao iniciar

#### O que é overlap_pct?

`overlap_pct = 0.20` significa: o sucessor pode iniciar quando o predecessor está 80% concluído (ou seja, quando já passou `duração × 0.80` do tempo do predecessor). Isso modela situações de fluxo parcial de lote ou processos que podem se sobrepor.

`overlap_pct = 0.0` (padrão) significa: o sucessor só inicia após o predecessor terminar completamente.

#### Algoritmo

```
Para operação SEM predecessora (primeira(s) do roteiro):
    early_start[op] = 0

Para operação COM predecessora(s):
    early_start[op] = max sobre todos os predecessores {
        early_start[pred] + duração[pred] × (1 − overlap_pct[aresta pred→op])
    }

Para qualquer operação:
    early_finish[op] = early_start[op] + duração[op]

Lead Time (horas) = max { early_finish[op] }   ← máximo entre TODAS as operações
Lead Time (dias)  = Lead Time (horas) ÷ horas_disponíveis_por_dia
```

#### Por que `max` e não apenas a última operação?

Porque o roteiro pode ter **ramificações paralelas**: dois grupos de operações que correm ao mesmo tempo e convergem depois. A que terminar mais tarde dita o lead time — independente de ser ou não a última da sequência. A última operação serial normalmente coincide com o máximo, mas não é garantido em roteiros mais complexos.

#### Exemplo prático

Roteiro com 3 operações em série, overlap de 20% entre op2 e op3:

```
op1: duração = 2h,  early_start = 0h,   early_finish = 2h
op2: duração = 3h,  early_start = 2h,   early_finish = 5h   (após op1, sem overlap)
op3: duração = 1h,  early_start = 4.4h, early_finish = 5.4h

early_start[op3] = early_start[op2] + duração[op2] × (1 − 0.20)
                 = 2 + 3 × 0.80 = 4.4h

Lead Time = max(2h, 5h, 5.4h) = 5.4h
Lead Time em dias (turno de 8h) = 5.4 ÷ 8 = 0.675 dias ≈ 1 dia útil
```

---

### Endpoints do módulo de roteiro

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | `/api/routing/operations` | Criar operação genérica |
| GET | `/api/routing/operations` | Listar operações |
| GET | `/api/routing/operations/{id}` | Buscar operação |
| POST | `/api/routing/routes` | Criar roteiro |
| GET | `/api/routing/routes` | Listar roteiros |
| GET | `/api/routing/routes/{id}` | Buscar roteiro |
| POST | `/api/routing/route-operations/{routeId}` | Adicionar operação ao roteiro |
| POST | `/api/routing/routes/{id}/edges` | Definir dependência predecessor→sucessor |
| GET | `/api/routing/routes/{id}/lead-time` | Calcular lead time via CPM |

---

## 2. CRP — Capacity Requirements Planning

### O que é

O CRP calcula **quanto** de cada centro de trabalho será necessário para executar as ordens de produção, em quais datas, comparando com a capacidade disponível e sinalizando sobrecargas.

### Posição no fluxo de produção

```
Pedido de Venda
    ↓
MRP (lê BOM + roteiro + estoque + parâmetros)
    ↓ gera
Sugestões de Ordens Planejadas
    ↓ PCP analisa e aprova
Ordens Aprovadas / Planejadas confirmadas
    ↓ PCP aciona manualmente o CRP
CRP (lê ordens + roteiros + capacidade dos centros)
    ↓ aponta sobrecargas
PCP ajusta datas / redistribui carga / autoriza hora extra
    ↓
PCP libera as ordens para o chão de fábrica
```

**Por que o CRP é manual e não automático?**
Porque a decisão de verificar capacidade, ajustar ordens e liberar ao chão de fábrica é humana e contextual. O PCP decide quando rodar o CRP:
- Após aprovar um lote de ordens (verificar viabilidade antes de liberar)
- Como análise prévia sobre sugestões do MRP ("e se eu aprovar tudo isso?")
- Após alterar datas de ordens manualmente

**O que o CRP lê:** ordens com qualquer status exceto `CANCELLED` — isso permite usá-lo tanto sobre ordens aprovadas quanto como análise preventiva sobre sugestões.

### Algoritmo

```
Para cada ordem com roteiro:
    Para cada operação do roteiro:
        required_hours[centro_trabalho, data] += tempo_efetivo × quantidade

Para cada (centro_trabalho, data):
    avail = horas_nominais_do_centro
    avail -= horas_bloqueadas_por_manutenção  ← integrado com Manutenção Preventiva
    if avail < 0: avail = 0
    load_pct = required_hours / avail × 100
    if load_pct > 100: sinalizar sobrecarga
```

### Integração com Manutenção Preventiva

Quando existem ordens de manutenção com status `PLANNED` ou `IN_PROGRESS` para um centro de trabalho em uma data específica, o CRP **subtrai** essas horas da capacidade disponível. Exemplo: centro com 8h nominais e 2h de manutenção agendada tem apenas 6h disponíveis para produção.

### Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | `/api/crp/calculate` | Calcular CRP para um plano MRP |
| GET | `/api/crp/plans/{planCode}` | Listar todos os registros de capacidade do plano |
| GET | `/api/crp/plans/{planCode}/overload` | Listar apenas centros sobrecarregados |
| GET | `/api/crp/work-centers/{id}?from=&to=` | Capacidade de um centro em um período |

**`POST /api/crp/calculate`:**
```json
{ "plan_code": 42 }
```
**Resposta:**
```json
{ "plan_code": 42, "total_entries": 18, "overload_count": 2 }
```

**`GET /api/crp/plans/42/overload`:**
```json
[
  {
    "work_center_id": 3,
    "req_date": "2026-06-10",
    "required_hours": 12.5,
    "available_hours": 8.0,
    "load_pct": 156.25,
    "is_overloaded": true
  }
]
```

---

## 3. APS — Advanced Planning and Scheduling

### O que é

O APS faz o sequenciamento finito de ordens: aloca cada operação de cada ordem nos centros de trabalho respeitando a capacidade real, produzindo um Gantt de produção.

### Diferença entre CRP e APS

| | CRP | APS |
|-|-----|-----|
| Objetivo | Medir carga por período | Sequenciar operação por operação |
| Granularidade | Centro × Dia | Operação × Hora exata |
| Capacidade | Infinita (aponta sobrecargas) | Finita (resolve sobrecargas deslocando datas) |
| Resultado | Relatório de carga % | Gantt com data/hora de início e fim de cada op |

### Algoritmo (capacidade finita)

```
Para cada ordem (ordenadas por prioridade/data de necessidade):
    Para cada operação do roteiro (ordem topológica):
        startCandidate = max(wcNextAvailable[centro], earliestStart[op])
        startCandidate = avançar para próximo dia útil se necessário
        endTime = startCandidate + duração
        wcNextAvailable[centro] = endTime
        registrar no Gantt: (ordem, operação, centro, startCandidate, endTime)
```

### Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | `/api/aps/schedule` | Gerar sequenciamento para um plano |
| GET | `/api/aps/plans/{planCode}` | Obter Gantt do plano |
| GET | `/api/aps/work-centers/{id}?from=&to=` | Gantt de um centro em um período |

**Gantt (trecho):**
```json
[
  {
    "order_id": 101,
    "operation_name": "Corte a laser",
    "work_center_id": 2,
    "start_time": "2026-06-09T07:00:00Z",
    "end_time": "2026-06-09T07:30:00Z"
  },
  {
    "order_id": 101,
    "operation_name": "Pintura eletrostática",
    "work_center_id": 5,
    "start_time": "2026-06-10T07:00:00Z",
    "end_time": "2026-06-10T09:00:00Z"
  }
]
```

---

## 4. Custo Padrão

### O que é

Calcula o custo de fabricação de um item considerando materiais (BOM) e mão de obra/máquina (roteiro). Suporta rollup multinível.

### Fórmula

```
custo_material  = Σ (qtd_componente × custo_padrão_componente)
custo_operação  = Σ (tempo_operação × taxa_centro_trabalho)
custo_overhead  = custo_operação × taxa_overhead (%)
custo_total     = custo_material + custo_operação + custo_overhead
```

### Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | `/api/standard-cost/calculate/{itemCode}` | Calcular custo padrão |
| GET | `/api/standard-cost/{itemCode}` | Buscar custo padrão salvo |
| GET | `/api/standard-cost/` | Listar todos os custos padrão |

---

## 5. Qualidade

### O que é

Registra pontos de inspeção ao longo do processo produtivo com laudo (aprovado/reprovado/condicional), quantidades e observações.

### Tipos de ponto de inspeção

| Tipo | Momento |
|------|---------|
| `RECEIVING` | Inspeção de recebimento (matéria-prima) |
| `IN_PROCESS` | Durante a fabricação, após uma operação |
| `FINAL` | Produto acabado, antes do estoque |

### Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | `/api/quality/inspection-points` | Criar ponto de inspeção |
| GET | `/api/quality/inspection-points` | Listar pontos |
| POST | `/api/quality/inspection-points/{id}/results` | Registrar laudo |
| GET | `/api/quality/inspection-points/{id}/results` | Buscar resultado |

---

## 6. Manutenção Preventiva

### O que é

Gerencia planos de manutenção periódica de máquinas e centros de trabalho. Gera ordens automaticamente conforme a frequência definida. As horas agendadas são descontadas da capacidade pelo CRP.

### Entidades

**Plano de Manutenção (`maintenance_plans`):**

| Campo | Descrição |
|-------|-----------|
| `machine_id` | Máquina que receberá manutenção |
| `work_center_id` | Centro de trabalho (afeta capacidade no CRP) |
| `frequency` | `DAILY`, `WEEKLY`, `MONTHLY`, `CUSTOM_DAYS` |
| `frequency_days` | Intervalo em dias |
| `estimated_hours` | Horas estimadas de parada |
| `next_scheduled_at` | Calculado automaticamente |

**Ordem de Manutenção (`maintenance_orders`):**

| Campo | Descrição |
|-------|-----------|
| `plan_id` | Plano de origem |
| `machine_id` | Máquina (copiado do plano) |
| `scheduled_date` | Data programada |
| `status` | `PLANNED` → `IN_PROGRESS` → `DONE` / `CANCELLED` |
| `actual_hours` | Preenchido ao concluir |
| `started_at` / `completed_at` | Timestamps automáticos na mudança de status |

### Ciclo de vida

```
Plano criado
    ↓ GenerateOrders (disparo manual ou periódico)
Ordem PLANNED  (idempotente: não cria duplicata para mesmo plano+data)
    ↓ AdvanceOrder {status: "IN_PROGRESS"}
Ordem IN_PROGRESS  (registra started_at)
    ↓ AdvanceOrder {status: "DONE", actual_hours: 1.5}
Ordem DONE  (registra completed_at)
```

### Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | `/api/maintenance/plans` | Criar plano |
| GET | `/api/maintenance/plans` | Listar planos (`?active=true`) |
| GET | `/api/maintenance/plans/{id}` | Buscar plano |
| GET | `/api/maintenance/machines/{machineId}/plans` | Planos de uma máquina |
| DELETE | `/api/maintenance/plans/{id}` | Desativar plano |
| POST | `/api/maintenance/orders` | Criar ordem manual |
| PUT | `/api/maintenance/orders/{id}/advance` | Avançar status / registrar horas reais |
| GET | `/api/maintenance/plans/{planId}/orders` | Ordens de um plano |
| GET | `/api/maintenance/work-centers/{wcId}/orders?from=&to=` | Ordens por período |
| POST | `/api/maintenance/orders/generate` | Gerar ordens automáticas (`{ "horizon_days": 30 }`) |

---

## 7. Previsão Estatística

### O que é

Calcula previsões de demanda futura aplicando modelos estatísticos a uma série histórica. Retorna o modelo de melhor ajuste (menor MAPE).

### Modelos disponíveis

| Modelo | Quando é melhor |
|--------|----------------|
| Holt-Winters (aditivo) | Séries com tendência + sazonalidade |
| Suavização Exponencial | Séries com tendência sem sazonalidade |
| Média Móvel (k=3) | Séries estáveis / sem padrão claro |
| Média Móvel (k=6) | Séries estáveis com mais histórico |

O sistema calcula o MAPE de cada modelo e retorna o de menor erro. O campo `model_used` indica qual foi selecionado.

### Endpoint

**`POST /api/forecast/statistical`**

```json
{
  "item_code": 1001,
  "history": [
    { "period": "2026-01", "quantity": 120.0 },
    { "period": "2026-02", "quantity": 135.0 },
    { "period": "2026-03", "quantity": 118.0 },
    { "period": "2026-04", "quantity": 142.0 },
    { "period": "2026-05", "quantity": 130.0 }
  ],
  "periods_ahead": 3
}
```

**Resposta:**
```json
{
  "item_code": 1001,
  "model_used": "exponential_smoothing",
  "mape": 4.82,
  "forecasts": [
    { "period": "2026-06", "quantity": 133.2 },
    { "period": "2026-07", "quantity": 135.8 },
    { "period": "2026-08", "quantity": 134.5 }
  ]
}
```

> A previsão é calculada em tempo real e não é persistida automaticamente. Para armazenar, use `POST /api/sales-forecast/blocks`.

---

## 8. Alertas de Exceções MRP

### O que é

Após o MRP rodar, exceções são geradas para situações que exigem atenção do PCP. Este módulo consolida e envia os alertas via **webhook** e/ou **e-mail**.

### Tipos de exceção

| Tipo | Significado |
|------|-------------|
| `LATE_ORDER` | Ordem planejada com data de necessidade no passado |
| `OVERDUE_PURCHASE` | Sugestão de compra com prazo vencido |
| `EXCESS_STOCK` | Estoque projetado acima do máximo definido |
| `OPEN_ORDER_NO_DEMAND` | Ordem aberta sem demanda correspondente |
| `CAPACITY_OVERLOAD` | Centro de trabalho sobrecarregado no período |

### Canais de notificação

| Canal | Campo no body | Requisito |
|-------|--------------|-----------|
| Webhook HTTP | `webhook_url` | URL do sistema destino |
| E-mail | `email_to` | SMTP configurado via `.env` |

Ambos os canais funcionam simultaneamente. Se o SMTP não estiver configurado, o e-mail é silenciosamente ignorado sem afetar o webhook.

### Configuração SMTP (`.env`)

```dotenv
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=seu@email.com
SMTP_PASSWORD=sua_senha_app
SMTP_FROM=erp@suaempresa.com
```

### Endpoint

**`POST /api/mrp-calculation/exceptions`**

```json
{
  "plan_code": 42,
  "webhook_url": "https://chat.empresa.com/mrp-alerts",
  "email_to": ["pcp@empresa.com", "gerencia@empresa.com"]
}
```

**Resposta:**
```json
{
  "plan_code": 42,
  "generated_at": "2026-05-22T10:00:00Z",
  "total": 3,
  "by_type": { "LATE_ORDER": 2, "EXCESS_STOCK": 1 },
  "exceptions": [
    {
      "item_code": 1001,
      "message_type": "LATE_ORDER",
      "description": "Ordem planejada para 2026-05-18, já vencida"
    }
  ]
}
```

**Corpo do e-mail gerado:**
```
Relatório de Exceções MRP — Plano 42
Gerado em: 22/05/2026 10:00
Total de exceções: 3

Por tipo:
  LATE_ORDER                     2
  EXCESS_STOCK                   1

Detalhes:
  Item 1001   [LATE_ORDER              ] Ordem planejada para 2026-05-18, já vencida
  Item 1002   [LATE_ORDER              ] ...
  Item 1008   [EXCESS_STOCK            ] Estoque projetado acima do máximo
```

---

## 9. Restrições e Configurador

### O que é

Permite definir regras de negócio que controlam quais combinações de atributos de um item são válidas. Útil em configuradores de produto ou validações de cadastro.

### Operadores suportados

`==`, `!=`, `>`, `<`, `>=`, `<=`, `IN`, `NOT_IN`

### Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| POST | `/api/restrictions` | Criar restrição |
| GET | `/api/restrictions` | Listar restrições |
| GET | `/api/restrictions/{id}` | Buscar restrição |
| POST | `/api/restrictions/{id}/evaluate` | Avaliar restrição com um contexto |
| DELETE | `/api/restrictions/{id}` | Remover restrição |

---

## Relação entre módulos

```
Pedido de Venda
      │
      ▼
    MRP ──────── BOM (estrutura do produto)
      │    └──── Roteiro (lead time via CPM, tipo de ordem por origin)
      │    └──── Estoque (saldo disponível)
      │    └──── Parâmetros (lote mínimo, estoque de segurança)
      │
      ├── Sugestões de Compra  → Pedido de Compra
      │
      └── Sugestões de Fabricação
                  │  origin INTERNA  → Ordem de Fabricação (OF)
                  │  origin EXTERNA/TERCEIROS → Ordem de Serviço (OS)
                  │
                  ▼ (PCP analisa e aprova)
            Ordens Aprovadas
                  │
          ┌───────┴────────┐
          ▼                ▼
        CRP              APS
   (carga % por       (sequenciamento
    centro/dia)        finito / Gantt)
          │
          └── Manutenção Preventiva
              (desconta horas de parada da capacidade disponível)
```
