# Testes por grupo

O workflow apresenta 14 grupos independentes, além de inventário, vet/build e validação Charter. Cada job usa um runner Ubuntu novo; o executor isola diretórios XDG e temporários e remove credenciais do ambiente de testes. OpenAI usa HTTP simulado. Não é homologação com API real nem isolamento de rede.

## Executar

Requisitos: Linux, Go conforme go.mod, Bash, Git, jq e compilador C para race detector.

```bash
cd src
bash test.sh --list
bash test.sh --check
bash test.sh --group dsl-charter
bash test.sh --group openai-agents
bash test.sh --all
```

No GitHub Actions, abra Functional test groups. Cada Tests / grupo tem resultado próprio. Em uma falha, use Re-run failed jobs para repetir os jobs que falharam. Novas execuções ocorrem somente em push para develop e em abertura, reabertura ou atualização de commits de PR de develop para main no mesmo repositório. Não há workflow_dispatch, agendamento ou gatilho de push direto na main. Todos os grupos são executados nesses eventos; a seleção individual permanece disponível localmente.

## Grupos e resultados esperados

| Grupo | Resultado esperado |
|---|---|
| dsl-charter | Converter cinco documentos da etapa, preservar conteúdo/status/hash e gerar DSL determinística sem badges. |
| dsl-scope | Converter cinco documentos da etapa, preservar conteúdo/status/hash e gerar DSL determinística sem badges. |
| dsl-implementation | Converter cinco documentos da etapa, preservar conteúdo/status/hash e gerar DSL determinística sem badges. |
| dsl-src | Converter cinco documentos da etapa, preservar conteúdo/status/hash e gerar DSL determinística sem badges. |
| dsl-core | Converter o projeto completo, detectar alterações, excluir backups e validar cache/integração normativa. |
| prompts | Selecionar somente o prompt correto, conferir as 34 rotas e rejeitar catálogo vazio ou acesso a outra etapa. |
| openai-agents | Simular criação, consulta, atualização idempotente e exclusão; bloquear chave ausente e separar norma de prompts. |
| changeset | Decodificar retorno, preservar seções, conferir SHA e rejeitar caminhos, operações e metadados inválidos. |
| auth | Validar configuração GitHub, permissões e rejeição de symlinks; nenhuma credencial real. |
| config | Validar caminhos, configuração de contexto e valores inválidos. |
| models | Validar catálogo GGUF, integridade, associação e seleção persistente do modelo. |
| runtime | Validar runtime com servidor e socket simulados, sem baixar modelos. |
| terminal | Validar URLs, branches, arquivos permitidos, histórico e encerramento de sessões. |
| backup | Validar origem/destino, snapshot Git local e integridade dos backups. |

## Evidências e critérios

Cada grupo gera test-results/<grupo>/ com tests.json, stderr.log, expected.txt, passed.txt, coverage.out, coverage.txt e summary.md. No Actions, os artefatos tests-<grupo> ficam disponíveis por 30 dias, inclusive em falhas. O resumo mostra o resultado por teste e subteste.

Falha, skip ou ausência de um teste esperado reprovam o grupo. Inventário incompleto reprova antes de executar a matriz. Ao adicionar um teste, registre seu nome exato em [groups.json](groups.json); o teste golden tem uma seleção por etapa. As asserções originais continuam sendo os oráculos; a organização não comprova cenários ainda não implementados.

A cobertura é medida por grupo sobre o pacote inteiro. Não somar percentuais nem apresentar cobertura parcial como cobertura global.

A separação preserva os arquivos _test.go no pacote Go existente, pois compartilham helpers e acessam funções internas. [Fixtures Markdown](markdown/README.md) continuam com cinco documentos por etapa. Esta mudança organiza execução e evidências, sem alterar produção ou promover gates documentais.

Referência consultada: doc-arquitetura/main 652e1035a6ebb7ade0f21a6f2b9d08156ba35b37. A reorganização de CI e testes foi solicitada explicitamente pelo usuário.
