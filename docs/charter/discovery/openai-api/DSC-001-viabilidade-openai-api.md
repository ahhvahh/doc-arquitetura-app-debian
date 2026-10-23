# DSC-001 — Viabilidade da OpenAI API e controle de consumo

![Tipo: DSC](../../assets/documentation-style/badges/badge-type-dsc.svg) ![Status: refinement](../../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../../assets/documentation-style/badges/badge-version-v1.svg)

## Pergunta e motivação

A integração direta à OpenAI pode atender às quatro especializações previstas e impedir que tarefas e inferências executem sem controle financeiro suficiente, conforme os limites aprovados no Charter?

A motivação é evitar faturamento excessivo provocado por múltiplas tarefas, geração, revisão por segundo modelo, correções e reprocessamentos.

## Tecnologia, escopo e baseline

- **Objeto:** API OpenAI e controles administrativos de organização e projeto; não se confunde com assinatura ChatGPT.
- **Especializações planejadas:** `charter`, `scope`, `implementation`, `source`, cada uma com projeto e credencial independentes.
- **Data de pesquisa:** 2026-10-22; documentação oficial online, sem acesso administrativo à organização do usuário.
- **Modelos/região/versões concretas:** PENDING — selecionar e verificar por especialização antes da homologação; não pressupor janela de contexto de 1.000.000 de tokens.

## Método e fontes verificáveis

Pesquisa documental em fontes oficiais, sem criar projetos, configurar chaves, gastar créditos ou executar testes remunerados.

| Fonte oficial | Evidência documental | Limite da evidência |
|---|---|---|
| [Spend limits](https://developers.openai.com/api/docs/guides/spend-limits) | Hard spend limits mensais por organização e projeto; opção explícita *Enforce a hard limit*; erros HTTP 429 na ultrapassagem do gasto rastreado | Aplicação não instantânea; possível pequeno excedente; não prova configuração da conta |
| [Error codes](https://developers.openai.com/api/docs/guides/error-codes) | `organization_spend_limit_exceeded` e `project_spend_limit_exceeded` são códigos distintos de outros erros 429 | Não prova que uma política específica foi ativada |
| [Admin APIs](https://developers.openai.com/api/docs/guides/admin-apis) e [API de limite por projeto](https://developers.openai.com/api/reference/cli/resources/admin/subresources/organization/subresources/projects/subresources/spend_limit) | Administração e leitura dos limites por API são documentadas | Requer credenciais administrativas e permissões não fornecidas neste estudo |
| [Controle por execução](https://developers.openai.com/cookbook/articles/per_run_spending_controller_responses_api) | Reserva preventiva e contabilização por chamada são padrões recomendados para orçar execuções com múltiplas chamadas | Exemplo não demonstra implantação deste aplicativo |

## Fatos, inferências e pendências

### DOCUMENTED — capacidades documentadas pelo provedor

1. Limites **mensais** de organização e de projeto podem interromper novas requisições da API com HTTP 429. Alertas de gasto isolados não bloqueiam tráfego.
2. Limites por projeto e de organização podem atuar em conjunto; o valor aprovado de uso/tier é um conceito diferente.
3. O bloqueio é assíncrono e pode permitir leve excedente; não é garantia matemática de despesa total igual ao teto.
4. Projetos distintos são compatíveis com a organização planejada do Charter, mas seu cadastro, permissões e limites precisam ser comprovados individualmente.

### INFERRED — adequação condicionada

A combinação de limites mensais na plataforma com controles locais por chamada, reserva de custo, autorização específica, concorrência e teto de sessão é uma solução **potencialmente adequada**, desde que a tolerância ao excedente residual seja avaliada contra a política obrigatória de proteção financeira. Essa inferência não homologa o provedor.

### UNKNOWN — evidências não obtidas

- Identificadores reais e isolamento dos quatro projetos; permissões, credenciais e titulares.
- Valores de limites mensais **efetivamente aplicados** nos projetos e na organização; mecanismo *Enforce a hard limit* habilitado em cada escopo.
- Demonstração administrativa de bloqueio e acompanhamento em cada projeto, sem testes não autorizados.
- Modelos/versões disponíveis, limites reais de entrada/saída, preços atuais por tipo de token, custos de ferramentas e rastreabilidade financeira do segundo modelo.
- Tolerância administrativa ao pequeno excedente da plataforma, considerando que a regra do projeto exige proteção remota suficiente.

## Custos e riscos

O limite de **1.000.000 tokens por inferência** e o teto local de **US$ 10 por sessão** são parâmetros aprovados do aplicativo, não permissões para exceder capacidade do modelo nem equivalentes aos limites **mensais** OpenAI. O fluxo pode incluir várias chamadas por tarefa; controles da plataforma não substituem a reserva de custo por chamada e a contabilização de tarefas concorrentes. Erros 429 de faturamento não devem ser repetidos automaticamente. Não usar segredos em arquivos Git, logs ou DSL.

## Critérios de homologação administrativa (sem nova execução neste estudo)

- Identificar organização e quatro projetos e validar credenciais segregadas, sem registrar segredos.
- Verificar em cada projeto e no limite de organização: teto mensal, ativação do modo bloqueante, permissões, abrangência e evidência datada.
- Validar o comportamento previsto de erros de limite a partir de evidências seguras ou ensaio explicitamente autorizado, com restrição financeira própria.
- Validar orçamento local com reserva preventiva, limite por inferência/sessão, limites efetivos do modelo e autorização **individual** para cada geração, revisão e correção.
- Registrar aprovação administrativa, responsável, data de verificação, validade e revalidação quando a configuração mudar.

## Conclusão de viabilidade

**VIÁVEL EM TESE para segregação e bloqueio mensal pela API; INCONCLUSIVO para habilitação remunerada no projeto.** Faltam evidências da conta, da eficácia aplicável e da cobertura de custos reais. Manter documento em `refinement` e impedir chamadas pagas até a homologação. Não converter esta pesquisa em mudança de política do Charter.

## Impacto e rastreabilidade

- [OPEN-004 — bloqueio remoto financeiro](../../open-questions/OPEN-004-projeto-openai-e-hard-limit.md)
- [OPEN-005 — escolha de provedores](../../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md)
- [OBJ-005 — limites e previsibilidade](../../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [CON-004 — restrição financeira cumulativa](../../conditions/CON-004-orcamento-obrigatorio.md)
- **PENDING — INT/contratos de integração e orçamento em Scope**: somente após gate; sem antecipação técnica.
