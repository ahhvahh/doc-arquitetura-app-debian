# Arquitetura atual

O executável Go sincroniza referências documentais, converte Markdown em DSL, administra agentes OpenAI e prepara tarefas sobre repositórios GitHub.

## Fluxos disponíveis

1. **Referência/DSL:** clona a branch normativa `main`, guarda o SHA e compila os documentos Markdown para DSL.
2. **Preparação de tarefa:** clona `develop` do alvo, cria uma branch local, seleciona um prompt normativo e gera `task.json` e `target.pcdsl`.
3. **Agentes OpenAI:** com chave e confirmação explícita, cria, consulta, atualiza ou exclui agentes persistentes. Ainda não processa tarefas por inferência.

A aplicação não baixa nem executa modelos locais. A inferência OpenAI de análise continua pendente dos controles de gasto e da implementação do contrato de retorno.

## Empacotamento

O pacote Debian instala o executável Go e a DSL normativa. A compilação usa `CGO_ENABLED=0`; não compila llama.cpp e não instala binários ou bibliotecas C++ para inferência.

## Execução legada

O comando `terminal` permanece para navegação e operações documentais locais. Mensagens livres são informadas como não processáveis; o fluxo não tenta iniciar LLM local nem chamar inferência remota. Utilize `openai analyze prepare` para validar o clone, o prompt e a geração da DSL.

Os comandos `check`, `analyze` e `routes` operam sobre regras DSL; `analyze` é determinístico e não chama modelo.
