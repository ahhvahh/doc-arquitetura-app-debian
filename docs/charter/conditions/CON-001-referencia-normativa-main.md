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

## Restrição vigente na primeira entrega — referência fixa, sem parametrização (2026-10-22)

**Origem:** o responsável escolheu a **opção A — manter a referência normativa fixa na primeira entrega**, preservando a decisão histórica de parametrização controlada como intenção para evolução futura.

- **Fonte única da norma vigente:** `https://github.com/ahhvahh/doc-arquitetura`, branch **`main`**, raiz **`ProjectCharter/1.0`**, exclusivamente **somente leitura**. A aplicação não poderá selecionar outro repositório, branch ou caminho como fonte normativa nesta primeira entrega.
- **Rastreabilidade e integridade:** cada consulta/compilação/sincronização normativa relevante deverá identificar a origem fixa e o **SHA da `main` efetivamente consultado**. A cópia local ou base consultável é derivada, e não substitui a fonte normativa. Atualizações somente por solicitação explícita, não durante inferência em andamento.
- **Parametrização adiada:** a proposta anterior de usar outras URLs previamente autorizadas continua registrada na seção **Decisão 01**, mas não está operacionalmente habilitada; sua adoção futura **depende de alteração normativa externa formalmente autorizada**, verificação de confiança das novas fontes e refinamento de contratos posteriormente.
- **Falha fechada:** se a fonte fixa ou seu SHA/integridade não puderem ser confirmados, bloquear a operação normativa dependente. Nenhum parâmetro administrativo, arquivo local ou resultado de modelo pode relaxar essas regras.
- **Governança Git distinta:** fixar a fonte normativa não amplia as permissões Git. A [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md) registra o adiamento de branches/PRs; gravações por agentes documentais permanecem limitadas à branch `develop` do repositório-alvo devidamente autorizado.

**Critérios de verificação pendentes:** confirmar origem/SHA da norma na operação real; detectar URL/branch divergente, atualização silenciosa ou falta de integridade; demonstrar negativa a tentativas de contornar a fonte fixa. A decisão é `DOCUMENTED`, não prova de implementação ou homologação.

**Status:** `refinement` / `v1`; preservar restrição anterior e histórico, inclusive a intenção futura de referência configurável. Relações: [OBJ-006](../objectives/OBJ-006-publicacao-manual-base-normativa.md), [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).
