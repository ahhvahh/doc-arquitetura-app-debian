# CON-001 — Referência normativa versionada

![Tipo: CON](../../assets/documentation-style/badges/badge-type-con.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Restrição
O sistema deve utilizar exclusivamente `ahhvahh/doc-arquitetura`, branch `main`, caminho `ProjectCharter/1.0`, como origem normativa; essa referência será somente leitura.
## Origem verificável
Decisão explícita do responsável e regra descrita no [AGENT-PROMPT normativo](https://github.com/ahhvahh/doc-arquitetura/blob/main/ProjectCharter/1.0/AGENT-PROMPT.md).
## Motivo e impacto
A referência precisa ser reproduzível; sem identificação da versão, análises diferentes não seriam comparáveis.
## Verificação
Uma execução deve registrar o SHA da referência normativa usada, manter uma cópia local verificável sem alterar o repositório de referência e vincular o SHA normativo ao cadastro de cada agente. A norma só será reenviada quando o operador solicitar atualização do agente.
## Aplicação à base consultável
O Vector Store é cópia derivada e pesquisável, não uma nova fonte normativa. Cada publicação identifica SHA da main e integridade do compilado. Sincronização somente por comando explícito, nunca silenciosamente durante inferência.

## Relações
- [OBJ-001](../objectives/OBJ-001-agentes-documentais-especializados.md).
- [OBJ-006](../objectives/OBJ-006-publicacao-manual-base-normativa.md).
## Condição de revisão
Alteração explícita da referência normativa pelo responsável.

## Proposta de regularização — Decisão 01 (2026-10-21)

**Origem:** decisão explícita do responsável pelo projeto — opção A, referência parametrizável com controle.
**Resultado da triagem:** proposta aceita para regularização documental; **não** constitui alteração já vigente da referência normativa externa.

- Permitir configurar a URL do repositório de referência entre repositórios explicitamente autorizados, mantendo `https://github.com/ahhvahh/doc-arquitetura` como padrão.
- Exigir leitura da branch `main` do repositório autorizado, sem consultar branches alternativas para formar a referência normativa.
- Identificar e registrar o SHA exato da `main` e a URL canônica utilizados em cada atualização e análise, com rastreabilidade da DSL derivada.
- Atualizar o material de referência apenas por solicitação manual explícita do operador, não durante inferências comuns.
- Bloquear a operação quando a URL não estiver autorizada, a branch não for `main`, o SHA não puder ser confirmado ou a integridade da cópia não puder ser verificada.

**Conflito e encaminhamento:** a restrição vigente deste CON e as instruções normativas de `ahhvahh/doc-arquitetura/main/ProjectCharter/1.0` ainda fixam essa origem como exclusiva. Até que a governança normativa seja formalmente compatibilizada e seus contratos dependentes sejam refinados, o aplicativo **não deve considerar a nova parametrização liberada para operação normativa**. A fonte externa permanece somente leitura; esta decisão registra somente o encaminhamento aprovado no Charter do aplicativo.

**Verificação pendente:** formalizar o contrato de configuração e autorização da lista de repositórios em Scope/Implementation após o gate, sem presumir o mecanismo técnico de armazenamento ou autenticação.
