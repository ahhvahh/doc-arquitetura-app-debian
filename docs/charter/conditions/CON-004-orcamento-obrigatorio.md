# CON-004 — Orçamentos obrigatórios em duas camadas

![Tipo: CON](../../assets/documentation-style/badges/badge-type-con.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Restrição
A aplicação deve negar inferências pagas quando não houver limite local válido por execução e especialização, previsão de consumo, tarifa/cotação, autorização financeira **e proteção remota suficiente, efetiva e comprovada no provedor/projeto aplicável**. Alertas e orçamentos administrativos sem bloqueio demonstrado não atendem a esta restrição. Para OpenAI, manter a previsão de quatro projetos independentes, um por especialização, cada um sujeito à verificação remota.

## Origem verificável
Decisão explícita do responsável sobre limite na aplicação **e** na conta OpenAI. Orientações documentadas em [`src/docs/OPENAI-SPEND-LIMITS.md`](../../../src/docs/OPENAI-SPEND-LIMITS.md), arquivo existente na baseline da aplicação.

## Motivo e impacto
Limita riscos financeiros e obriga autorização consciente antes de usar serviços cobrados por token. Limites locais controlam chamadas da aplicação; limites do projeto OpenAI protegem a conta também contra chamadas por outros consumidores autorizados.

## Forma de verificação
- Teste negativo demonstra que não existe requisição de rede quando orçamento, preço ou câmbio estão ausentes ou vencidos.
- Testes concorrentes validam reservas de saldo por transação.
- Em cada provedor/projeto que realizará inferência paga, exigir evidência válida de proteção remota efetivamente bloqueante para o escopo do gasto autorizado. Ausência, incerteza, recurso meramente informativo ou falha na verificação exigem **bloqueio**, sem exceção por aceite de risco.
- Manter os quatro projetos OpenAI segregados, com verificações independentes de credenciais, consumo e eficácia dos mecanismos de proteção; não chamar alertas de gasto nem valores de orçamento de hard limits garantidos sem evidência.

## Itens afetados
- [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md)

## Condição de revisão
Somente após revisão de segurança, viabilidade de cobrança da API e autorização do responsável.

## Aplicabilidade da decisão 06 — provedores remotos
Os controles locais de orçamento são obrigatórios também para provedores remotos adicionais. Proteções do próprio provedor devem ser usadas **quando comprovadamente oferecidas**, mas não podem ser apresentadas como hard limits garantidos sem evidência. Serviços com insuficiência de dados de preço/uso ou sem mecanismo de controle compatível com o risco autorizado não serão habilitados para inferência paga até decisão verificável.

A aprovação humana do merge à `develop` não autoriza, por si só, a realização de novas despesas. Relações: [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) e [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md).

## Política de falha fechada — decisão 09 (2026-10-21)
**Origem:** opção A escolhida pelo responsável em [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md).

Para **as integrações vigentes OpenAI e agentes Amazon Bedrock (AWS)** — e qualquer provedor adicionado no futuro —, operações **pagas** exigem duas condições cumulativas: (1) orçamento local válido e autorizado e (2) **proteção remota verificada e suficiente para interromper consumo remunerado acima do limite aplicável**. A falta de qualquer delas determina o bloqueio antes da requisição.

- Não admitir limite apenas local, alerta sem corte de chamadas ou declaração não comprovada de hard limit como substituto da proteção remota.
- Verificar a efetividade do mecanismo do provedor e o contexto de abrangência/identidade administrativa; aplicar em **cada um dos quatro projetos OpenAI independentes** que sejam habilitados.
- Se a validade da proteção não puder ser demonstrada, negar a chamada paga sem presumir capacidades da API externa.
- Testes simulados ou **comprovadamente sem cobrança** não equivalem a permissão para produzir custos.
- Evidências da análise de viabilidade, garantias comerciais reais e meios técnicos de comprovação devem ser aprofundados no Discovery pertinente e depois no Scope/Implementation, sem antecipar mecanismo nesta restrição.

**Impacto:** a lista histórica de três provedores foi substituída pela decisão 22 de [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md); nenhuma das duas integrações atualmente previstas (OpenAI/AWS) é declarada apta a operar chamadas pagas por esta decisão. Relações: [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) e [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md).

## Limites locais configuráveis por inferência e sessão — decisão de 2026-10-22

**Restrição adicional:** o sistema deve respeitar limites padrão **configuráveis** de **1.000.000 de tokens por inferência** (entrada e saída, observando todos os componentes faturáveis informados) e **US$ 10,00 de orçamento monetário acumulado por sessão**. Os tokens gastos pelas inferências são acumulados como métrica da sessão, mas um milhão **não é** o teto de tokens total da sessão.

**Verificação obrigatória:**
- Conferir capacidade e limites reais do modelo antes da inferência, sem pressupor suporte a um milhão de tokens em uma chamada nem permitir truncamento do contexto exigido.
- Bloquear preventivamente qualquer nova chamada remunerada cujo custo conservador reservado, somado ao consumo e às demais reservas da sessão, ultrapasse o teto monetário da sessão; contabilizar os custos de planejamento, geração, segunda avaliação e correções no mesmo acumulador.
- Persistir consumo e reservas, inclusive em interrupção, retomada e concorrência, e impedir reinicialização artificial do orçamento por nova instrução.
- Negar inferências pagas sem **autorização financeira individual específica**, preço/câmbio válidos, demais orçamentos aplicáveis e **proteção remota efetivamente bloqueante e comprovada**. O teto monetário local não substitui a proteção remota nem assegura teto absoluto na fatura externa.

**Origem:** decisão expressa do responsável em 2026-10-22, vinculada a [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) e [OBJ-009](../objectives/OBJ-009-execucao-automatica-de-tarefas-elegiveis.md). Situação `DOCUMENTED`, sem implementação ou homologação demonstradas; o documento permanece em `refinement`.

## Evidências de Discovery para homologação por provedor — D-001

- [DSC-001 — OpenAI API](../discovery/openai-api/DSC-001-viabilidade-openai-api.md): hard spend limits mensais documentados, mas configuração e eficácia não verificadas na organização e nos quatro projetos.
- [DSC-002 — AWS AgentCore](../discovery/aws-agentcore/DSC-002-viabilidade-aws-agentcore.md): orçamento controlado na AWS Settings está em oferta limitada; cobertura dos recursos e custos adicionais não comprovada; Budgets tradicional insuficiente sozinho.

A viabilidade documental não constitui aceite de risco ou autorização de cobrança. Toda chamada paga permanece bloqueada sem proteção remota suficiente e comprovada, autorização individual, preços válidos, reservas preventivas e limites locais por inferência e por sessão.
