# Orçamentos e limites de gastos — OpenAI API (proposta de integração)

> **Status: refinement — documentação de configuração e contrato funcional.**
> Os comandos `openai` descritos como planejados **ainda não estão implementados no binário atual**. Esta página não altera limites da conta OpenAI.
> Fonte normativa de arquitetura: [ProjectCharter/1.0 — main](https://github.com/ahhvahh/doc-arquitetura/tree/main/ProjectCharter/1.0), somente leitura; não duplicar modelos normativos neste repositório.

## Objetivo e escopos de proteção

Aplicar limites em **camadas independentes** para os agentes Charter, Scope, Implementation e Source, com bloqueio local antes de enviar requisições pagas e proteção adicional no lado da OpenAI.

| Camada | Onde fica o limite | O que é aplicado | Observação |
|---|---|---|---|
| Aplicação Go | Banco local privado | Valor máximo por operação, orçamento diário e mensal de cada agente e orçamento global | **Bloqueio local obrigatório** antes da inferência; valores em BRL |
| Managed Agent / sessão | `spend_control.limit` da sessão | Limite de gasto da **sessão**, em centavos de USD | Controle de consumo **best-effort**; não equivale a limite mensal permanente do agente |
| Projeto de API OpenAI | Plataforma → projeto → Limits → Spend | Limite mensal com opção **Enforce a hard limit** | Bloqueia novas requisições após o consumo contabilizado atingir o teto; não é instantâneo |
| Organização OpenAI | Organization limits → Spend | Limite mensal opcional para todos os projetos | Contenção adicional de gastos entre projetos |

**Não confundir:** alertas de gastos só notificam e não interrompem chamadas. Um limite mensal do **projeto OpenAI** não é automaticamente um limite por agente quando vários agentes compartilham o mesmo projeto. Uma sessão de Managed Agents sem `spend_control` não possui teto de sessão.

Referências oficiais (consultadas em 20/10/2026):
- [Spend limits — configurar alertas e hard limits](https://developers.openai.com/api/docs/guides/spend-limits)
- [Managing projects — projetos, membros, chaves e limites](https://help.openai.com/en/articles/9186755-managing-your-work-in-the-api-platform-with-projects)
- [Create an agent session — parâmetro spend_control](https://developers.openai.com/api/reference/ruby/resources/beta/subresources/agents/subresources/sessions/methods/create)
- [Update an agent session — alterar limite](https://developers.openai.com/api/reference/resources/beta/subresources/agents/subresources/sessions/methods/update)
- [Spend limits e APIs administrativas](https://developers.openai.com/api/docs/guides/admin-apis)

## 1. Cadastrar o limite mensal dentro da conta OpenAI (recomendado)

1. Acesse [OpenAI API Platform](https://platform.openai.com/) com uma conta que tenha permissão para administrar o projeto.
2. Se necessário, crie um projeto de API exclusivo para o aplicativo BIOMA. Para um **limite mensal OpenAI separado por agente**, crie quatro projetos: `bioma-charter`, `bioma-scope`, `bioma-implementation` e `bioma-source`.
3. Abra **Project settings → Limits → Spend → Edit spend limit** no projeto correto.
4. Defina **Monthly spend limit**, em dólares (USD), e **ative Enforce a hard limit**. Salve.
5. Configure **Spend alerts** em patamares antecipados, por exemplo 50%, 80% e 100%; alertas não substituem o hard limit.
6. Repita em cada projeto se desejar quatro orçamentos independentes. Opcionalmente configure também um hard limit da organização em **Organization limits → Spend**.
7. Em **Project settings → API keys**, crie uma chave **do projeto correto**, com apenas as permissões necessárias. Associe a chave ao agente correspondente no cadastro local; nunca a versione no Git.

Se usar **um único projeto OpenAI** para os quatro agentes, ele terá um teto mensal agregado do lado da plataforma; as cotas individuais serão responsabilidade do banco local da aplicação.

**Importante:** a cobrança da API é separada do plano ChatGPT. Os hard limits da OpenAI são mensais e contabilizados em USD; a propagação não é instantânea e pode haver pequeno excedente após atingir o valor. Verifique limites e permissões da sua conta na plataforma.

## 2. Limite por inferência dentro do Managed Agent

Na criação de uma **nova sessão para cada inferência**, incluir o objeto abaixo na chamada de criação de sessão, além dos campos normais requeridos pela API:

```json
{
  "spend_control": {
    "limit": 100
  }
}
```

Neste exemplo, `100` representa **US$ 1,00 (100 centavos de dólar)** de orçamento da sessão; não representa R$ 100,00. O campo é opcional e, se omitido, a sessão fica sem limite. A contabilização de consumo da sessão é aproximada (`best-effort`). Não afirmar que isso constitui garantia de cobrança máxima exata.

Esse controle é **da sessão**, e não um atributo permanente de teto mensal definido no cadastro do agente. A aplicação deverá abrir uma sessão isolada por requisição, passar um limite explícito e jamais reutilizar sessões entre inferências. Não atualizar automaticamente as instruções normativas persistidas no agente; atualização da DSL normativa exige comando separado.

### Conversão do orçamento de BRL para USD

O valor em reais é o parâmetro de aprovação do operador. Imediatamente antes da chamada, consultar uma cotação USD/BRL recente, registrar fonte, data e valor e converter para centavos de dólar **arredondando para baixo** para não ultrapassar o limite autorizado em reais. Como o limite de sessão precisa ser inteiro positivo (mínimo 1 centavo de USD), **bloquear** a chamada se não couber sequer US$ 0,01 no saldo autorizado. Não presumir que a cotação efetiva do emissor do cartão ou os tributos serão idênticos à taxa de referência.

## 3. Limites na aplicação Go (contrato planejado)

A aplicação manterá registros locais por `account_id`, `agent_id`, `api_project_id` e `operation_id`. Os valores monetários devem ser inteiros em centavos (evitar `float` como autoridade contábil). Políticas:

- **Falhar fechado:** sem limite válido, modelo/tarifa identificável ou cotação recente, **não executar** inferência paga.
- Antes da inferência, estimar tokens de entrada (instruções persistentes do agente + DSL de `/docs` + metadados + envelope da chamada) e custo de saída, incluindo reserva para raciocínio, além de ferramentas e outros itens cobrados, quando aplicáveis.
- Mostrar custo esperado e cenário máximo; exigir confirmação explícita; reservar o saldo em transação atômica para evitar concorrência entre inferências.
- Bloquear quando o pior cenário previsto superar qualquer limite: por inferência, por agente/dia, por agente/mês, global/dia ou global/mês.
- Limitar a saída com `max_output_tokens`, quando suportado, sem tratar isso como garantia de custo total.
- Após a execução, conciliar métricas de consumo/custo retornadas pela API; registrar diferenças, erros e mudanças cambiais. Dados ausentes ou inconsistentes mantêm a reserva em aberto até reconciliação, sem liberar orçamento prematuramente.
- Tratar explicitamente os erros `project_spend_limit_exceeded`, `organization_spend_limit_exceeded`, `credit_balance_exhausted`; **não** repetir automaticamente requisições pagas ao recebê-los.
- A cotação não é armazenada como cache reutilizável entre operações; é um dado auditável associado à execução. A DSL do projeto e o histórico da inferência não são reaproveitados entre sessões.

### Interface CLI proposta (ainda não implementada)

```bash
# Cotas locais em reais, por agente
doc-arquitetura openai agents budget charter --por-execucao 3.00 --diario 20.00 --mensal 100.00
doc-arquitetura openai agents budget scope --por-execucao 3.00 --diario 20.00 --mensal 100.00

# Consultar configurações e consumo local
doc-arquitetura openai agents budget status

# Inferência: exibir custo e pedir aprovação antes do envio à API
doc-arquitetura infer --repo https://github.com/OWNER/REPO --agente charter
```

Os valores acima são **exemplos**, não limites já cadastrados. A implementação futura também deverá permitir configurar limite global, associar uma chave/projeto OpenAI a cada agente e registrar o SHA da DSL normativa persistida.

A CLI comum não deverá criar hard limits remotos silenciosamente. Alterar o orçamento da OpenAI pelo painel é operação administrativa independente. Automação futura de limites remotos, se implementada, deve usar credencial administrativa separada, permissões adequadas e aprovação explícita; uma API key comum não é uma chave Admin.

## 4. Exemplo de validação anterior à execução

```text
Agente: Charter
Projeto OpenAI: bioma-charter
Limite local por execução: R$ 3,00
Estimativa de entrada: 40.000 tokens
Estimativa de saída: 8.000 tokens
Custo esperado: R$ 1,20
Custo máximo reservado: R$ 2,80
Limite de sessão OpenAI: US$ 0,55 [exemplo após conversão]
Cotação: USD/BRL, origem/data registradas
Limite remoto mensal: verificar configuração do projeto na OpenAI

Confirmar execução paga? [s/N]
```

**Exemplo ilustrativo**: números não representam uma medição de custo real. A sessão pode ter cobrança aproximada e o limite mensal remoto deve ser verificado na conta.

## 5. Critérios de validação e testes unitários obrigatórios

Cada ação precisa de casos de sucesso, erro e concorrência documentados para evidência no GitHub Actions:
- Cadastro, edição, remoção e consulta dos limites locais por agente e globais; nenhum agente pode executar sem política válida.
- Cotação ausente, inválida, vencida ou inacessível; arredondamento BRL → USD cents e limite mínimo de 1 centavo USD.
- Identificação de projeto/chave corretos e rejeição de credenciais cruzadas.
- Inclusão de `spend_control.limit` em **todas** as sessões; rejeição de sessão sem limite; sessão nova por inferência.
- Estimativa de entrada/saída, cenário máximo, insuficiência de contexto ou orçamento e recusa na confirmação do operador.
- Reserva concorrente e conciliação de gastos efetivos; proteção contra execução duplicada na repetição de uma requisição.
- Erros 429 de orçamento da OpenAI, créditos esgotados, timeout, conexão interrompida e custo ainda não reconciliado.
- Nenhum segredo em stdout, logs, arquivos do projeto, DSL, diffs, relatórios Actions ou Git.
- Evidências de `go test`, relatório de cobertura e publicação dos artefatos CI após implementação dos testes.

**Estado desta documentação:** especificação e instruções de configuração manual. O código de orçamento, sessões OpenAI, banco local e workflow específico **não foi implementado** neste commit.
