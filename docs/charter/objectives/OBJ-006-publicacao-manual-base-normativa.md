# OBJ-006 — Publicação manual da base normativa consultável

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
Reenviar integralmente a norma documental em cada inferência aumenta os tokens consumidos e mistura instruções permanentes do agente com o acervo consultável.

## Resultado esperado
A aplicação permitirá ao operador compilar uma representação textual compacta e rastreável da norma doc-arquitetura/main/ProjectCharter/1.0 e publicar, sob comando explícito, um conjunto pesquisável separado para Charter, Scope, Implementation e Development (src).

O armazenamento remoto não altera a fonte normativa nem substitui os documentos originais. A associação posterior dos agentes à ferramenta de consulta é responsabilidade independente da publicação.

## Critérios verificáveis
- O operador escolhe uma especialidade ou solicita atualização de todas elas; a origem é sempre main normativa verificada por SHA.
- O comando padrão compila a DSL e exibe prévia sem upload. Somente opção explícita de aplicação publica os arquivos.
- A compilação identifica arquivos de origem e hashes; exclui backups e não inclui tokens ou material do projeto-cliente na base normativa.
- A representação enviada utiliza formato aceito pela pesquisa remota, preservando conteúdo suficiente para rastreabilidade.
- A aplicação compara hashes e não reenvia conteúdo idêntico.
- A indexação remota deve ser confirmada antes de considerar a nova versão ativa; uma falha não invalida automaticamente a versão anterior.
- IDs remotos, commit da referência e hashes são registrados localmente sem credenciais.
- A consulta pelo agente só será considerada disponível após vínculo e teste próprios da ferramenta de pesquisa.
- Testes automáticos cobrem prévia, atualização, idempotência, falha parcial, integridade e ausência de segredo nos registros.
- O operador será informado sobre custos e limites do serviço antes de habilitar operações pagas.

## Fora deste objetivo
A publicação isolada não realiza inferência, não altera o projeto-alvo e não prova equivalência semântica entre DSL e Markdown.

## Relações
- [OBJ-001 — Agentes especializados](OBJ-001-agentes-documentais-especializados.md)
- [OBJ-002 — Contexto compacto e rastreável](OBJ-002-contexto-compacto-rastreavel.md)
- [CON-001 — Origem normativa](../conditions/CON-001-referencia-normativa-main.md)
- [CON-002 — Proteção de credenciais](../conditions/CON-002-protecao-credenciais.md)
- [OPEN-002 — Fidelidade e custo](../open-questions/OPEN-002-fidelidade-e-custo.md)

## Situação
Intenção aprovada pelo responsável; documento em refinement. A integração externa, formatos específicos, contratos de consulta e liberação de produção ainda exigem Discovery, Scope e Implementation.

## Provisionamento integrado por especialidade — decisão do responsável
O operador poderá realizar, em uma **única ação explícita**, o provisionamento de cada agente: baixar `doc-arquitetura/main` por SHA, compilar os documentos normativos da especialidade em DSL pesquisável, sincronizar o acervo no Vector Store, aguardar a indexação, criar ou atualizar o agente e registrar a associação de agente, Vector Store e versão normativa. O padrão continuará sendo **prévia sem custos nem gravação remota**; `--apply` autorizará o provisionamento, sujeito às verificações de credenciais e orçamento.

O sistema separará a **DSL normativa** (somente a referência `main`) da **DSL do projeto-alvo** (clone de `develop` por tarefa). Não publicará dados do projeto-alvo no acervo normativo. A vinculação poderá ocorrer na requisição de execução pelo aplicativo se não houver suporte remoto a `file_search` no cadastro persistente de agentes.

A sincronização do acervo e o cadastro do agente são processos verificáveis separadamente, com hashes, IDs e falhas visíveis. A existência do acervo não comprova que o modelo consultou a norma: será exigido teste de recuperação.
