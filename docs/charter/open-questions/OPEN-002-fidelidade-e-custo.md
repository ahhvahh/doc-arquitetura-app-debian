# OPEN-002 — Fidelidade de DSL e custos

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Diretriz de validação aprovada, com critérios de comprovação de fidelidade e limites de custo ainda pendentes.
## Questão
**Diretriz de validação escolhida:** verificações estruturais obrigatórias, segunda avaliação por modelo de linguagem diferente do gerador e revisão humana final. Ainda é necessário comprovar a fidelidade semântica da DSL, preservar diagramas e referências e validar critérios mensuráveis de qualidade e tokens; o orçamento financeiro segue obrigatório em OBJ-005 e CON-004.
## Contexto
A compressão sintática existente ainda não possui prova de equivalência semântica. Diagramas, assets, critérios de aceite e exceções não podem desaparecer durante edição.
## Impacto
**Bloqueadora:** sim para declarar a inferência documental pronta para produção; não bloqueia planejamento e medição exploratória.
## Delimitação da decisão de armazenamento
A publicação manual da DSL em armazenamento consultável pode ser experimentada independentemente da inferência, conforme [OBJ-006](../objectives/OBJ-006-publicacao-manual-base-normativa.md). Upload e indexação bem-sucedidos não comprovam cobertura da norma, fidelidade semântica ou redução real de tokens. Essas dúvidas continuam bloqueadoras da inferência em produção.

## Próxima ação
Estabelecer evidências para integridade estrutural, avaliação por segundo LLM, **correção e reavaliação obrigatórias quando houver divergência relevante**, e posterior decisão humana; medir fidelidade, custos das tentativas e recuperação do original. Sem condições financeiras ou de segurança para reavaliar, bloquear o aceite. A concordância dos modelos não prova equivalência semântica.
## Destino previsto
Scope de extração, rastreabilidade e validação; Implementation de armazenamento e transporte.

## Pendências específicas do contrato DSL textual
A saída textual pretendida é `BIOMA-DSL-CHANGESET/2.0` com cabeçalho de tarefa e blocos `FILE` por arquivo, sem JSON. A decisão de formato **não** homologa a gramática interna. É necessário especificar caracteres reservados, delimitadores escapáveis, blocos multilinha, estrutura de metadados, badges, links, tabelas, código, Mermaid e regras de edição por seção, e demonstrar conversão DSL → Markdown que preserve todo conteúdo não tocado. O estado `BLOCKED` e sua representação textual também precisam ser formalizados.

Validar cenários de duplicidade de arquivos, hash divergente, alterações fora da etapa, symlinks, path traversal, tamanho excessivo e mudanças de baseline. Até lá, o conversor JSON/Markdown existente não representa implementação deste novo contrato.

## Proposta de regularização — Decisão 02 (2026-10-21)

**Origem:** decisão explícita do responsável — opção A, **DSL reversível com compactação seletiva**.
**Resultado da triagem:** aceita como diretriz para a regularização documental; a gramática DSL e os critérios de homologação permanecem pendentes.

- O material que puder resultar em criação ou alteração de arquivo deve preservar integralmente o conteúdo necessário à reconstrução sem perdas, inclusive Markdown original, cabeçalhos, badges, links, tabelas, blocos de código, diagramas Mermaid, metadados e separação por arquivo.
- Compactar seletivamente índices, catálogos e material de apoio à consulta quando for seguro, sem ocultar condições normativas obrigatórias, referências relevantes ou evidências necessárias à análise.
- Se a representação resumida não bastar para fundamentar uma análise ou alteração, consultar a fonte original íntegra antes de concluir; sem essa recuperação, bloquear a conclusão dependente.
- Manter por arquivo o caminho de origem, SHA da baseline e hash de integridade; identificar explicitamente conteúdo omitido, compactado e recuperado.
- A aplicação deve validar o ciclo Markdown → DSL → Markdown contra o documento original e impedir aplicação de mudanças que removam ou transformem conteúdo não solicitado.
- Medir tokens efetivamente consumidos e estimativa de custo, comparando o modo completo ao seletivo, sem considerar a redução de tokens comprovada antes de medição.

**Limites da decisão:** esta proposta não define delimitadores, escape, formato de campos, recuperação remota, limite numérico de tokens ou algoritmo de compactação. O contrato `BIOMA-DSL-CHANGESET/2.0` continua sujeito às validações descritas nesta OPEN. Essas definições deverão ser formalizadas em Scope/Implementation, seguindo os gates documentais.

**Impacto no gate:** a diretriz de arquitetura foi escolhida, mas a `OPEN-002` permanece em `refinement` e continua bloqueadora para declarar inferência documental pronta para produção até comprovar fidelidade e segurança dos casos de teste.

## Extensão aprovada — comparação objetiva entre modelos (2026-10-21)
Uma comparação válida deverá usar a mesma tarefa, prompt efetivo, documentação e SHAs de referência, ou identificar explicitamente variações. Avaliar integridade do Markdown, links, diagramas, validade do changeset, aderência à norma, correções humanas, tempo, tokens e custos previstos e realizados, quando mensuráveis.

Pesos, escala de qualidade e limites numéricos de aprovação ainda exigem definição e validação. Registrar dados experimentais não comprova automaticamente superioridade, fidelidade ou equivalência semântica.

Relações: [OBJ-007](../objectives/OBJ-007-comparacao-controlada-de-modelos.md), [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md), [OPEN-005](OPEN-005-estrategia-de-provedores-e-modelos.md).

## Decisão 10 — Dupla avaliação de qualidade com revisão humana (2026-10-21)
**Origem:** o responsável selecionou explicitamente a opção B — **validação automática + segunda análise por LLM + revisão humana** — e aprovou sua documentação exclusivamente na `develop`.

**Resultado da triagem:** modalidade de controle de qualidade **aceita como objetivo funcional**, sem afirmar implementação, confiabilidade homologada ou disponibilidade financeira dos modelos.

- **Primeira barreira:** a aplicação realiza verificações automatizadas obrigatórias da integridade das alterações propostas, abrangendo escopo e hashes de origem, conteúdo não solicitado, links, badges, tabelas, blocos de código, referências e diagramas pertinentes. Falhas estruturais impedem aceitar e aplicar o resultado; uma segunda análise por LLM não dispensa nem substitui esses testes.
- **Segunda avaliação:** após a validação estrutural, **outro modelo de linguagem**, distinto do que gerou a alteração, analisa criticamente o resultado frente à tarefa, às evidências e à versão normativa. Seu parecer deverá apontar qualidade, omissões, contradições, riscos e trechos sem sustentação, com identificação do provedor/modelo e contexto avaliado.
- **Decisão humana:** apresentar as alterações, os testes e o parecer secundário ao operador para revisão humana antes do aceite final e de qualquer integração Git que venha a ser normativamente permitida. Parecer favorável do revisor LLM não constitui aprovação automática.
- **Rastreabilidade:** relacionar inferência geradora, avaliação secundária, parecer, métricas de uso/custo, eventuais divergências, ajustes e decisão humana ao projeto, sessão e baselines. Não tratar concordância entre LLMs como prova da equivalência semântica da DSL.
- **Orçamento:** a revisão secundária também é uma operação de inferência e obedece a [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md). Sem proteção remota comprovada ou autorização financeira exigida para chamada paga, **não executar a revisão paga nem considerar o fluxo de validação concluído**.
- **Governança Git:** a aprovação humana final de qualidade não substitui a autorização explícita para integração à `develop`; commits em branch de sessão continuam condicionados à [OPEN-003](OPEN-003-politica-de-branches.md). Nenhuma gravação fora de `develop` é autorizada por esta decisão.

**Pendências de comprovação:** critérios de equivalência Markdown → DSL → Markdown, parâmetros de qualidade e escala, classificação de divergências relevantes, limites de tokens e validação dos modelos elegíveis e testes de custos/segurança. A política de insuficiência de janela de contexto está definida pela decisão 16, ainda sem comprovação operacional. **O limite vigente é de até quatro ciclos de correção e reavaliação, após a análise inicial**, conforme a decisão 14, que substitui o aspecto sem teto numérico da decisão 13. O tratamento de divergências relevantes está definido na decisão 12 abaixo. Definições técnicas pertencem a Scope/Implementation após os gates. Manter a OPEN em `refinement` enquanto tais evidências faltarem.

**Relações canônicas:** [OBJ-003 — Alterações verificáveis](../objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md), [OBJ-004 — Evidências de testes](../objectives/OBJ-004-evidencias-de-testes-automatizados.md), [OBJ-007 — Avaliação de modelos](../objectives/OBJ-007-comparacao-controlada-de-modelos.md).

## Decisão 12 — Correção obrigatória e reavaliação antes do aceite (2026-10-21)
**Origem:** opção A aprovada explicitamente pelo responsável, para documentação somente no Charter de `develop`.

**Resultado da triagem:** regra de qualidade aceita como intenção de produto, **sem implementação nem fidelidade semântica homologada**.

- Toda divergência **relevante** identificada pela segunda análise por LLM exige **corrigir a proposta** antes de encaminhá-la à revisão humana de aceite.
- Após a correção, **reexecutar a validação estrutural e a análise crítica por outro modelo de linguagem, distinto do gerador**, sobre a **nova versão** da proposta. A avaliação anterior não vale para conteúdo diferente.
- Enquanto persistir divergência relevante, omissão ou contradição sem solução, **não declarar a alteração validada, não gerar commit de resultado validado e não permitir sua integração Git**. A revisão humana final só ocorre depois de tratar e reavaliar a divergência.
- **Falhas estruturais, alterações fora do escopo, perda de informação, links quebrados ou violações normativas** permanecem bloqueadoras independentemente de quaisquer pareceres de LLM ou autorização humana.
- Se não for possível corrigir, comprovar adequação, obter nova avaliação, ou satisfazer as condições financeiras e de segurança, **registrar o bloqueio e preservar a rastreabilidade** sem contornar a verificação exigida.
- Cada nova geração ou avaliação potencialmente paga requer **autorização financeira e orçamento aplicáveis**, além da **proteção remota comprovada** exigida em [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md); a decisão de corrigir não autoriza gastos adicionais automaticamente.
- Registrar por tentativa as versões do conteúdo, a origem da divergência, o reparo, o segundo parecer, as evidências estruturais, as métricas e os custos previstos e medidos, sem presumir que consenso entre modelos significa fidelidade semântica.
- A revisão e a decisão humanas finais continuam obrigatórias; **aprovar integração à `develop` exige autorização humana distinta**. Branch e commits automáticos de sessão só serão permitidos após resolver [OPEN-003](OPEN-003-politica-de-branches.md) na norma externa.

**Pendências:** critérios para definir relevância, medidas quantitativas de fidelidade/qualidade, demonstração de cabimento do contexto integral e homologação do conversor DSL ainda carecem de decisão e evidências. A ausência de limite numérico adotada na decisão 13 foi **substituída pelo máximo de quatro correções** da decisão 14, sem alterar os bloqueios financeiros e normativos. Contratos técnicos pertencem a Scope/Implementation. Manter a `OPEN-002` em `refinement`.

## Decisão 13 — Ciclo iterativo com contexto integral e feedback recente (2026-10-21)

**Atualização de governança:** a regra originalmente sem teto numérico, registrada nesta decisão, foi **substituída** pelo limite de quatro ciclos de correção da **decisão 14** abaixo. Permanecem válidos o contexto integral, o feedback recente e os bloqueios.
**Origem:** aprovação explícita do responsável para substituir a definição de limite fixo de tentativas por uma repetição de correção e reavaliação **sem limite numérico pré-definido**, exclusivamente no Charter da `develop`.

**Resultado da triagem:** diretriz funcional aprovada; não comprova existência de mecanismo de execução, qualidade homologada ou capacidade de contexto de provedores.

- **Contexto original inalterado:** a cada divergência relevante, reenviar **integralmente o mesmo material de origem da primeira inferência**, com tarefa, prompt original, documentos, normas e referências, conservando identificadores, versões, hashes e SHAs da baseline. A nova tentativa não pode substituir silenciosamente essa baseline nem omitir material original necessário.
- **Complemento da tentativa:** acrescentar ao material original **a proposta mais recente**, a **observação do avaliador da última inferência** e a instrução de corrigir os problemas apontados. Não tratar uma reescrita ou um resumo do contexto original como reenvio integral.
- **Ciclo de qualidade:** produzir a correção; repetir as validações estruturais obrigatórias; solicitar **nova avaliação por LLM diferente do gerador** sobre a versão corrigida. Se persistir divergência relevante, retomar o ciclo com o **mesmo material original** e o novo resultado/parecer mais recente.
- **Regra histórica de iterações (substituída pela decisão 14):** originalmente, a continuidade foi prevista sem teto numérico, até alcançar resposta válida. A regra vigente limita o processo a quatro ciclos de correção após a análise inicial, sem autorizar execução ou gastos ilimitados.
- **Histórico e rastreabilidade:** preservar localmente todas as tentativas e pareceres, versões propostas, baselines, divergências, correções, testes, modelos/versões, tempo, tokens, custos e decisões; o material reenviado ao modelo corresponde à origem integral e ao **feedback mais recente**, sem necessidade de concatenar todo o histórico no prompt de cada rodada.
- **Bloqueios invioláveis:** falhas estruturais, violações normativas, perda de integridade, escopo inadequado e divergências relevantes pendentes impedem aceite e publicação de resultado validado, mesmo com consenso de modelos ou desejo de prosseguir.
- **Autorização financeira a cada chamada paga:** cada nova geração ou avaliação paga requer previsão de custo, reserva/limite local e autorização explícita, além de **proteção remota efetiva comprovada**, de acordo com [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md). Sem qualquer condição, **interromper sem nova requisição** e registrar o resultado como não validado.
- **Limite de contexto e disponibilidade:** se o material integral com os complementos não couber na capacidade comprovada do modelo ou não puder ser transmitido com segurança, **bloquear a tentativa**, sem truncamento silencioso. A decisão posterior passou a ser a **redefinição explícita de tarefa e baseline pelo operador**, conforme decisão 16; nenhuma mudança é implícita.
- **Revisão humana:** somente após obter versão válida, apresentar evidências e proposta para revisão humana final; a **integração Git à `develop` exige aprovação humana separada** e permanece condicionada às restrições normativas em [OPEN-003](OPEN-003-politica-de-branches.md).

**Pendências técnicas e de homologação:** demonstrar preservação integral do contexto, fidelidade da DSL, capacidade de contexto, critérios objetivos de divergência relevante e testes do ciclo de correção. O Charter não determina algoritmos, contratos ou implementação. Documento segue em `refinement`.

## Decisão 14 — Máximo de quatro ciclos de correção (2026-10-21)
**Origem:** o responsável determinou posteriormente **um máximo de quatro iterações de correção**, interrompendo o processo quando não houver resposta válida após a quarta.

**Resultado da triagem:** ajuste aprovado do objetivo funcional; substitui **somente** a ausência de teto numérico da decisão 13, mantendo as exigências já documentadas.

- **Contagem inequívoca:** a geração inicial e sua avaliação constituem a tentativa de origem; em caso de divergência relevante, podem ser realizadas **no máximo quatro novas iterações de correção** (`1/4`, `2/4`, `3/4` e `4/4`). Cada iteração compreende proposta corrigida, nova validação estrutural e nova análise por LLM distinto do gerador.
- **Material reenviado:** em cada correção, conservar **integralmente o contexto e o material de origem da primeira inferência**, acrescentando **a proposta e a observação mais recentes do avaliador**. O limite numérico não autoriza compactar ou omitir esse contexto.
- **Condição de conclusão:** se qualquer uma das quatro iterações produzir uma resposta válida, com todos os testes estruturais e avaliação secundária sem divergência relevante pendente, encerrar o ciclo de correção e encaminhar à **revisão humana final**.
- **Condição de parada:** se a quarta iteração ainda apresentar problema relevante, **interromper sem quinta correção automática**, registrar `BLOCKED` para o trabalho e preservar as evidências. O limite não representa aceite automático de resposta incompleta, nem autoriza contornar uma divergência por ter esgotado tentativas.
- **Parada antecipada obrigatória:** falha estrutural ou violação normativa impede aceitar/aplicar a proposta; ausência de autorização financeira, saldo local, proteção remota comprovada, credenciais, segurança ou capacidade suficiente de contexto impede **novas chamadas** mesmo antes da quarta iteração.
- **Orçamento de cada operação:** geração e avaliação secundária remuneradas exigem suas próprias previsões e autorizações, de acordo com [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md).
- **Rastreabilidade:** registrar ordinal da iteração, versão, baseline, última divergência, correção, testes, modelos, tokens, custo, resultado e motivo de encerramento, inclusive quando o processo terminar por bloqueio antecipado.
- **Governança Git:** não gerar commit de proposta inválida; a integração à `develop` continua dependente de aprovação humana separada e da política normativa indicada em [OPEN-003](OPEN-003-politica-de-branches.md).

**Pendente:** evidenciar os critérios de relevância, homologação da DSL, verificações financeiras e testes de parada nas quatro iterações. Não presumir que a convergência será alcançada nem que o recurso já está implementado. Manter a `OPEN-002` em `refinement`.

## Decisão 16 — Contexto insuficiente exige redefinição explícita pelo operador (2026-10-21)
**Origem:** escolha explícita do responsável — **opção C, permitir redefinição da tarefa** — autorizada para documentação somente no Charter da branch `develop`.

**Resultado da triagem:** diretriz funcional aprovada, sem alegar implementação ou capacidade efetiva de contexto dos provedores.

- **Condição de disparo:** antes de solicitar geração, correção ou segunda avaliação, se o **material original integral** mais o resultado e a observação recentes não couber na capacidade de contexto **comprovada e disponível** do modelo autorizado, **interromper a tentativa sem realizar a chamada**. Não truncar, resumir, omitir ou substituir material de modo silencioso.
- **Comunicação ao operador:** informar a insuficiência de contexto, qual etapa ficou bloqueada e que a **tarefa atual não obteve validação**. Preservar a entrada original, a baseline, a última proposta, o parecer, as evidências e os custos efetivamente incorridos; não interpretar a limitação como aprovação do conteúdo.
- **Redefinição consciente:** solicitar ao operador **nova definição explícita do escopo e da baseline**, com ciência das partes mantidas e das partes excluídas. Não escolher um modelo alternativo, alterar o escopo ou substituir a baseline por conta própria. Se o operador não redefinir, **manter o trabalho bloqueado**.
- **Novo ciclo identificado:** somente depois da autorização explícita registrar a nova tarefa e a **nova baseline com seus próprios identificadores e versões**, distinta da origem anterior e com ligação rastreável à tentativa interrompida. Reavaliar a viabilidade do **contexto integral da nova tarefa** antes de executar, sem reclassificar a tentativa antiga como válida e sem reabrir automaticamente um ciclo esgotado sob a mesma tarefa.
- **Limite de correções preservado:** para cada tarefa efetivamente autorizada, a análise inicial pode ser seguida por **no máximo quatro correções e reavaliações** segundo a decisão 14. A mudança deliberada para **outra tarefa e baseline** não permite encadear indefinidamente reinícios automáticos para contornar esse limite.
- **Validação preservada:** aplicar integralmente os controles estruturais e normativos, segunda avaliação por LLM distinto do gerador e posterior revisão humana. O escopo reduzido autorizado não equivale à comprovação de fidelidade sobre o material excluído; apontar claramente a limitação da cobertura.
- **Financeiro e segurança:** qualquer nova inferência paga sob a nova tarefa requer estimativa, **autorização financeira própria**, proteção remota comprovadamente bloqueante, segregação por usuário e credenciais válidas conforme [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md).
- **Governança Git:** redefinir a tarefa **não autoriza commits inválidos, criação de branches proibidas, PR ou merge**. A integração à `develop` depende de aprovação humana final específica e do alinhamento normativo pendente em [OPEN-003](OPEN-003-politica-de-branches.md). O vínculo da nova tarefa à sessão permanece rastreável sem criar refs novas neste registro.

**Pendências de verificação:** critérios de previsão de tamanho de contexto e reserva para saída/raciocínio, comprovação da capacidade dos provedores, rastreabilidade da redefinição e testes negativos de truncamento/reinício. Regras técnicas detalhadas pertencem a Scope/Implementation. `OPEN-002` permanece em `refinement`.

## Decisão 17 — Configuração por inferência e conversão DSL integral para Markdown (2026-10-21)
**Origem:** aprovação explícita do responsável para refinar o conceito de sessão, seleção de contexto/modelo/prompt **a cada inferência** e entrega de alterações em DSL convertidas pela aplicação, exclusivamente no Charter de `develop`.

**Classificação:** diretriz funcional de produto **DOCUMENTED**; não comprova gramática homologada, implementação de conversão, modelos elegíveis ou fidelidade sem perdas.

- **Sessão versus inferência:** repositório-alvo, nível documental, identificador, histórico de ações e branch **prevista** pertencem à sessão. A **cada solicitação de inferência**, o operador escolhe explicitamente **provedor/modelo autorizado**, **tamanho de contexto** dentro da capacidade efetiva e **prompt** (do catálogo ou informado/editado pelo operador); valores usados em uma solicitação não são assumidos como política imutável de toda a sessão.
- **Contexto verificável:** identificar por inferência a versão e a origem do prompt, configuração/tamanho de contexto solicitado e admissível, identidade e versão do modelo, documentos e SHAs utilizados. A seleção de uma janela maior não aumenta o limite real de contexto do provedor nem permite omitir documentos de forma silenciosa.
- **Resposta pretendida:** o LLM apresenta **título do commit, descrição do commit, lista identificável de arquivos adicionados/modificados/excluídos e conteúdo documental correspondente exclusivamente em DSL textual**. O título e a descrição são sugestões sujeitas a validação, não autorização para comandos Git fornecidos pelo LLM.
- **Conversão sob responsabilidade da aplicação:** a aplicação interpreta e valida a DSL e gera os **arquivos Markdown completos afetados**, conservando o conteúdo preexistente fora das mudanças autorizadas, incluindo **seções, texto, tabelas, links e referências, badges, blocos de código, diagramas Mermaid e representações de gráficos**. Recursos visuais ou arquivos externos indispensáveis deverão continuar íntegros e resolúveis; não simular capacidade de reconstrução quando não houver dados e suporte comprovados.
- **Falha fechada e nenhuma perda silenciosa:** não aplicar mudanças, gerar commit de resultado declarado validado ou aceitar publicação se houver DSL inválida, arquivo faltante, perda de semântica, referência quebrada, conteúdo gráfico não representável, Mermaid inválido, violação normativa ou incapacidade de converter com fidelidade verificável. Preservar a saída DSL original, o motivo e as evidências para inspeção.
- **Validação e revisão:** validar escopo, caminhos, hashes, integridade da estrutura Markdown e conteúdos especiais **antes de qualquer commit**; o resultado também deverá passar pela **segunda avaliação independente de LLM** e revisão humana nos pontos anteriormente definidos. A mesma inferência dispõe de **até quatro ciclos de correção** após a primeira tentativa, sem quinta correção automática; a correção preserva a baseline e os insumos originalmente autorizados e incorpora a divergência mais recente.
- **Alteração de configuração entre inferências:** trocar modelo, contexto ou prompt para uma **nova inferência** da mesma sessão não apaga evidências, não substitui a baseline anterior sem autorização nem habilita automaticamente chamadas remuneradas. Mudanças substanciais do escopo de uma tarefa bloqueada exigem a nova definição e a baseline explícitas descritas na **decisão 16**.
- **Proteções financeiras e Git:** toda chamada paga (geradora ou revisora) permanece submetida à autorização financeira, orçamento local, controle remoto bloqueante comprovado e segregação por usuário de [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md). A branch `feature/doc/{session-id}` e a PR previstas continuam **não autorizadas para execução por agentes** enquanto a [OPEN-003](OPEN-003-politica-de-branches.md) não resolver o conflito com a norma atual.

**Pendências de homologação:** definir e demonstrar em Scope/Implementation a gramática `BIOMA-DSL-CHANGESET/2.0`, a conversão fiel de conteúdos e recursos externos, critérios objetivos de equivalência e casos de teste para Markdown, Mermaid e gráficos. Não criar contratos técnicos nesta decisão. Manter `OPEN-002` em `refinement`.

## Decisão 20 — Corrigir somente divergências materiais e registrar sugestões editoriais (2026-10-21)
**Origem da escolha:** o responsável aprovou expressamente a **opção B: corrigir apenas divergências materiais**, preservando revisão humana e bloqueios normativos/estruturais. A decisão qualifica o termo «divergência relevante» empregado nas decisões 10, 12, 13 e 14; não substitui o processo de dupla avaliação.

**Regra funcional:**
- **Divergência material:** qualquer falha fundamentada que **elimine, omita, altere ou deturpe requisitos**; **mude o significado** de uma regra ou decisão; produza **contradição** entre documentos ou com a baseline aprovada; cause **falha de integridade/fidelidade** (incluindo informações perdidas na conversão DSL → Markdown, links, tabelas, Mermaid ou gráficos); **viole a referência normativa**, os limites da etapa, restrições de segurança ou outras condições obrigatórias de aceite. Tais divergências **exigem correção e nova avaliação**, antes de declarar a proposta válida.
- **Sugestão apenas editorial:** recomendação de ortografia, pontuação, estilo, clareza ou formatação que **não altere conteúdo, intenção, estrutura obrigatória, significado, links normativos ou rastreabilidade**. Registrar a sugestão e sua justificativa juntamente com o parecer do LLM revisor; apresentar ao responsável na **revisão humana**, sem exigir por si só uma nova inferência de correção.
- **Evidência da classificação:** o parecer secundário deverá identificar o trecho/documento afetado, a baseline, a divergência ou sugestão, sua justificativa, a classificação «material» ou «somente editorial» e o encaminhamento. Uma mera concordância entre modelos não supre provas de integridade nem a validação estrutural.
- **Sem descaracterizar bloqueios:** erro estrutural, omissão de conteúdo obrigatório, violação normativa ou dúvida não resolvida sobre possível perda de requisitos **não podem ser classificados como meramente editoriais para liberar a proposta**. Quando não houver evidência suficiente para demonstrar que a diferença é exclusivamente editorial, o resultado **permanece não validado** até esclarecimento e tratamento seguro; não há aceitação automática por incerteza.
- **Ciclo de correção:** após a geração e a avaliação iniciais, divergências materiais podem iniciar **no máximo quatro iterações de correção e reavaliação** da mesma tarefa; cada iteração preserva o material original integral, a baseline e a proposta e observação mais recentes, conforme decisões 13 e 14. Não é necessário gastar rodadas automáticas exclusivamente com sugestões editoriais.
- **Proteção financeira:** cada chamada paga de geração ou de segunda avaliação e cada nova correção continuam exigindo autorização própria, orçamento e **proteção remota bloqueante comprovada**, com segregação de credenciais conforme [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md). Se faltar alguma condição, bloquear sem inferir aceite.
- **Governança humana e Git:** propostas estrutural/normativamente inválidas ou com divergência material pendente **não geram commit de resultado validado**. A revisão humana final permanece obrigatória nos momentos definidos, e a aprovação para integração em `develop` continua **distinta** do aceite de encerramento/PR, condicionada à [OPEN-003](OPEN-003-politica-de-branches.md).

**Pendência:** validar a eficácia da classificação material/editorial com exemplos representativos, critérios verificáveis e testes de falsos negativos; a fidelidade do conversor DSL ainda requer homologação. Este registro é **DOCUMENTED**, não `IMPLEMENTED`. A `OPEN-002` permanece em `refinement`.

## Decisão 24 — Pré-configuração por papel e bloqueio do aceite sem revisor distinto (2026-10-21)
**Origem:** decisão expressamente aprovada pelo responsável para configuração prévia, persistente e reutilizável das inferências, antes de o operador abrir uma sessão documental.

- **Configuração inicial:** a aplicação oferecerá configuração prévia por **especialização** (Charter, Scope, Implementation, Source), **operação/tipo de análise** e **papel**. Serão mantidos e apresentados os valores padrão de **motor (OpenAI ou Amazon Bedrock AgentCore), modelo elegível e prompt de tarefa** para **gerador** e **revisor independente**. Os padrões deverão poder ser definidos antes da sessão e recuperados após reinício, sem guardar credenciais no arquivo.
- **Abertura da sessão:** o operador escolhe o repositório e o nível documental, carrega a configuração correspondente e pode inspecionar a elegibilidade do gerador e do revisor. O arquivo/tela de configuração é um **requisito funcional de planejamento**, sem especificação de formato, caminho, implementação ou tecnologia no Charter.
- **Configuração por inferência preservada:** os valores padrão **não se sobrepõem à decisão 17**. Durante a sessão, o usuário pode selecionar **outros modelos/motores elegíveis, tamanho de contexto admissível e prompt** em cada chamada, com registro completo da origem/versão do perfil, dos valores efetivos, da baseline e das alterações autorizadas. Alterações posteriores da configuração não reescrevem o histórico.
- **Independência obrigatória:** o segundo parecer continua exigindo **outro modelo de linguagem distinto e homologado**, não apenas outro nome de provedor ou de agente. Para AgentCore, deve ser verificável o modelo subjacente usado. Se o revisor não estiver elegível, disponível ou for equivalente ao gerador, o sistema **não pode declarar o resultado validado, aceitá-lo ou produzir commit de resultado validado**. Não dispensar a revisão por substituição automática pelo próprio gerador ou somente pela revisão humana.
- **Falha fechada e segurança:** configuração persistente não habilita gasto automaticamente. Cada geração, revisão e correção remunerada requer autorização financeira própria, segregação de credenciais, limites locais e **proteção remota bloqueante comprovada**. Falha normativa, estrutural, de contexto integral ou de validação continua bloqueadora; mantêm-se a revisão humana posterior e **até quatro ciclos de correção** após a tentativa inicial.

**Pendências de verificação:** homologação de modelos e motores, capacidade efetiva do AgentCore de identificar os modelos, disponibilidade de revisor independente, testes de persistência e troca de valores, segurança de configuração e custos. Esses aspectos não estão IMPLEMENTED por este registro. Relações: [OBJ-001](../objectives/OBJ-001-agentes-documentais-especializados.md), [OBJ-002](../objectives/OBJ-002-contexto-compacto-rastreavel.md), [OBJ-007](../objectives/OBJ-007-comparacao-controlada-de-modelos.md) e [OPEN-005](OPEN-005-estrategia-de-provedores-e-modelos.md). `OPEN-002` permanece em `refinement`.
