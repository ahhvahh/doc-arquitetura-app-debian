# CON-004 — Orçamentos obrigatórios em duas camadas

![Tipo: CON](../../assets/documentation-style/badges/badge-type-con.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Restrição
A aplicação deve negar inferências pagas na ausência de limite local válido por execução e por agente, cotação e previsão de consumo. Deve utilizar também as proteções disponibilizadas pelo provedor, incluindo limite de sessão, quando suportado, e configuração administrativa de limite mensal do projeto OpenAI.

## Origem verificável
Decisão explícita do responsável sobre limite na aplicação **e** na conta OpenAI. Orientações documentadas em [`src/docs/OPENAI-SPEND-LIMITS.md`](../../../src/docs/OPENAI-SPEND-LIMITS.md), arquivo existente na baseline da aplicação.

## Motivo e impacto
Limita riscos financeiros e obriga autorização consciente antes de usar serviços cobrados por token. Limites locais controlam chamadas da aplicação; limites do projeto OpenAI protegem a conta também contra chamadas por outros consumidores autorizados.

## Forma de verificação
- Teste negativo demonstra que não existe requisição de rede quando orçamento, preço ou câmbio estão ausentes ou vencidos.
- Testes concorrentes validam reservas de saldo por transação.
- Todas as sessões suportadas possuem limite positivo configurado; a ausência do recurso exige política explícita de bloqueio ou modo compatível aprovado.
- A documentação fornece o passo a passo para cadastrar hard limit e alertas no projeto OpenAI, sem alegar que a CLI já os configura remotamente.

## Itens afetados
- [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md)

## Condição de revisão
Somente após revisão de segurança, viabilidade de cobrança da API e autorização do responsável.
