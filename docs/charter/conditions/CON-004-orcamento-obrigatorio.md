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

Para **OpenAI, Anthropic e Google Gemini**, operações **pagas** exigem duas condições cumulativas: (1) orçamento local válido e autorizado e (2) **proteção remota verificada e suficiente para interromper consumo remunerado acima do limite aplicável**. A falta de qualquer delas determina o bloqueio antes da requisição.

- Não admitir limite apenas local, alerta sem corte de chamadas ou declaração não comprovada de hard limit como substituto da proteção remota.
- Verificar a efetividade do mecanismo do provedor e o contexto de abrangência/identidade administrativa; aplicar em **cada um dos quatro projetos OpenAI independentes** que sejam habilitados.
- Se a validade da proteção não puder ser demonstrada, negar a chamada paga sem presumir capacidades da API externa.
- Testes simulados ou **comprovadamente sem cobrança** não equivalem a permissão para produzir custos.
- Evidências da análise de viabilidade, garantias comerciais reais e meios técnicos de comprovação devem ser aprofundados no Discovery pertinente e depois no Scope/Implementation, sem antecipar mecanismo nesta restrição.

**Impacto:** a seleção nominal dos três provedores continua registrada, mas nenhum é declarado apto a operar inferência paga por esta decisão. Relações: [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) e [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md).
