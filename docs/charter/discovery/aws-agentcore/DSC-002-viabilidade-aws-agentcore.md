# DSC-002 — Viabilidade do Amazon Bedrock AgentCore e proteção financeira

![Tipo: DSC](../../assets/documentation-style/badges/badge-type-dsc.svg) ![Status: refinement](../../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../../assets/documentation-style/badges/badge-version-v1.svg)

## Pergunta e motivação

O Amazon Bedrock AgentCore pode executar agentes para a primeira entrega com limites preventivos e proteção financeira remota suficiente, cobrindo **todos** os serviços utilizados por cada tarefa?

O risco prioritário é o processamento encadeado de várias tarefas, cuja cobrança não corresponde apenas à inferência do modelo subjacente.

## Tecnologia, escopo e baseline

- **Objeto:** AgentCore; Runtime/harness, modelo efetivo, memória, ferramentas, observabilidade e serviços auxiliares quando utilizados.
- **Ambiente:** conta, região, papéis IAM e projeto AWS reais: UNKNOWN; disponibilidade e elegibilidade ainda não comprovadas.
- **Data de pesquisa:** 2026-10-22, por documentação oficial, sem login à conta do operador e sem chamadas remuneradas.
- **Framework, modelos e desenho de execução:** PENDING — escolha técnica pertence às etapas seguintes, não decorre automaticamente do estudo de viabilidade.

## Método e evidências documentais

| Fonte oficial | Evidência documentada | Limitações |
|---|---|---|
| [AWS Settings — Spend limits](https://docs.aws.amazon.com/accounts/latest/reference/create-spend-limit.html) | Limites mensais por projeto, com pausa de recursos; experiência ainda em lançamento para número limitado de clientes; exige Paid Plan; mínimo de US$ 20 ou projeção conservadora maior | Não prova disponibilidade na conta, cobertura de AgentCore nem proteção de todas as cobranças |
| [AWS Settings — efeitos da pausa](https://docs.aws.amazon.com/accounts/latest/reference/create-spend-limit.html) | Controles antecipados opcionais incluem parar recursos que elevam custos, inclusive Bedrock entre serviços elegíveis; limite por projeto é pré-impostos | Pausa pode interromper aplicações; atenção ao risco de exclusão de dados do projeto após 90 dias parado sem ação |
| [AWS Budgets — gestão de custos](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-managing-costs.html) | Informações atualizadas até três vezes ao dia, tipicamente a cada 8–12 horas; notificações podem atrasar | Orçamento tradicional e alertas não demonstram corte imediato de cobrança |
| [AWS Budgets — ações](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-controls.html) | IAM/SCP e ações sobre recursos podem ser disparadas após limiar | Não constituem, isoladamente, proteção remota comprovada para todo AgentCore |
| [AgentCore — observabilidade e custos](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/harness-operations.html) | Uma execução de harness pode envolver muitas inferências; `maxIterations`, `timeoutSeconds`, `maxTokens`, tempo ocioso e lifetime são limites de execução; traz itens de custo distintos | Controles de execução não equivalem a limite monetário remoto da conta |
| [AgentCore — preços](https://aws.amazon.com/bedrock/agentcore/pricing/) | Tabelas oficiais por capacidade, com mudanças de preço/medição por recurso | Tarifas, região, moedas e descontos exigem consulta na homologação; não fixar cotação desatualizada |

## Fatos, inferências e pendências

### DOCUMENTED — capacidades do fornecedor

1. A oferta de gasto controlado por projeto da AWS Settings está disponível apenas a clientes elegíveis durante a liberação gradual, com mínimo mensal de **US$ 20** ou estimativa maior. Isso não contradiz o teto local de **US$ 10 por sessão**: são unidades de controle distintas.
2. O AWS Budgets convencional sofre atraso de dados, logo alerta/ação tardia não é sinônimo de limite rígido instantâneo.
3. AgentCore tem custos separados por Runtime (CPU/memória), modelo efetivo, ferramentas, Memory, Gateway, observabilidade e outros recursos vinculados, conforme o uso realizado.
4. A AWS documenta controles adicionais de duração, iterações e tokens em harness; eles limitam processos, mas não demonstram integralmente o limite mensal de cobrança.
5. A suspensão por spend limit pode afetar disponibilidade e exigir reativação, e uma pausa superior a 90 dias sem ação pode ocasionar exclusão de dados do projeto.

### INFERRED — viabilidade condicionada

A combinação de isolamento por projeto AWS elegível, limite remoto aplicável a **todos os recursos cobrados**, controle de tempo/iterações, orçamento local por chamada e autorização individual pode ser suficiente para o objetivo. A abrangência de projetos AWS Settings sobre recursos do AgentCore **não foi demonstrada** nesta análise. Não homologar por analogia com o serviço Bedrock de modelos.

### UNKNOWN — comprovações ainda necessárias

- Elegibilidade da conta para AWS Settings spend limit e ativação por projeto; organização, faturamento, titular e região.
- Quais recursos específicos do AgentCore estão vinculados ao projeto com spend limit, inclusive modelos, ECR/S3/CloudWatch/rede, e como são bloqueados quando o limite é atingido.
- Identificação do LLM efetivo por agente, versão, preços por região, limites de tokens, métricas de uso e modelo independente de revisão.
- Custos máximos por duração, memória, execução concorrente e ferramenta; isolamento de solicitações distintas.
- Compatibilidade da política exigida de bloqueio financeiro com custos residuais, cobrança diferida e impostos.

## Critérios de homologação administrativa

- Consultar a conta real e comprovar disponibilidade de AWS Settings spend limit, Paid Plan, valor mínimo aplicável e abrangência dos serviços cobrados.
- Inventariar todos os recursos/serviços usados por agente e projetar o pior caso, especialmente chamadas múltiplas de modelos e ferramentas.
- Verificar separadamente os custos por região e evidências de limites financeiros remotos; se Budgets for o único controle, manter a inferência remunerada bloqueada.
- Validar identidade/permissões, segregação, modelo efetivo e controle por autorização financeira individual, reservas locais e limites de sessão.
- Exigir definição da política de pausa, recuperação e retenção de dados antes de usar projeto sujeito a suspensão.
- Validar no ambiente, mediante autorização e orçamento separados, comportamento seguro sem gastar além do permitido; revalidar ao modificar recursos.

## Conclusão de viabilidade

**INCONCLUSIVO para habilitação paga do AgentCore nesta conta.** Existem ferramentas oficiais de orçamento e de limitação de execução, mas cobertura, elegibilidade e eficácia remota sobre **todo** o gasto não estão demonstradas. Manter este DSC em `refinement`; não habilitar inferência remunerada.

## Impacto e rastreabilidade

- [OPEN-004 — proteção financeira](../../open-questions/OPEN-004-projeto-openai-e-hard-limit.md)
- [OPEN-005 — oferta AWS AgentCore aprovada como intenção](../../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md)
- [OBJ-005 — consumo previsível](../../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [CON-004 — exigências cumulativas](../../conditions/CON-004-orcamento-obrigatorio.md)
- **PENDING — contratos/fluxos de integração em Scope:** depende do gate do Charter.

## Sequenciamento da homologação — D-001, alternativa B (2026-10-23)

- **Priorização:** a equipe verificará primeiro a proteção financeira dos quatro projetos OpenAI previstos em [DSC-001](../openai-api/DSC-001-viabilidade-openai-api.md). Esta decisão não cancela a integração AWS AgentCore planejada em [OPEN-005](../../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md).
- **Estado AWS:** `UNKNOWN` para elegibilidade administrativa, bloqueio remoto e cobertura integral dos custos do AgentCore e recursos associados. Toda chamada remunerada via AWS permanece bloqueada; nenhum resultado obtido na OpenAI é evidência transferível.
- **Retomada independente:** após ou em paralelo à conclusão documental da OpenAI, verificar conta/região/projeto AWS, serviços faturados, condições de suspensão, custos residuais e evidência de proteção remota suficiente segundo [CON-004](../../conditions/CON-004-orcamento-obrigatorio.md). Pesquisa sem cobrança pode continuar, mas não comprova eficácia por si só.
- **Resultado:** permanecer em `refinement`, com D-001/D-002 abertas e sem autorização para chamadas pagas ou testes de cobrança.
