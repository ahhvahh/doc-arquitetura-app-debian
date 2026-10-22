# OBJ-004 — Evidências de testes automatizados

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
A integração possui operações com impacto em credenciais, arquivos, branches e custos externos. Resultados de testes precisam ser reproduzíveis e auditáveis.

## Resultado esperado
Cada ação funcional possui testes unitários identificáveis, executados de modo automatizado com evidências consultáveis no GitHub Actions antes da publicação de nova versão.

## Critérios verificáveis
- Cada operação funcional definida em Scope possui casos positivos e negativos com identificadores estáveis.
- Execuções automatizadas disponibilizam resultado individual dos testes e relatório de cobertura.
- Falhas de teste impedem que o fluxo seja declarado concluído.
- Testes unitários usam simulações para OpenAI e Git; não consomem saldo da OpenAI nem alteram repositórios remotos.

## Relações
- [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md).
- [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

## Evidências obrigatórias para dupla avaliação — decisão 10 (2026-10-21)
O objetivo de qualidade aprovado prevê resultados de **testes estruturais**, **parecer de segundo modelo distinto do gerador** e **decisão de revisão humana**, todos rastreáveis ao mesmo conjunto de alterações e às baselines utilizadas.

**Critérios verificáveis de produto:** um erro de integridade reprova o resultado mesmo que ambos os modelos o considerem adequado; a ausência de uma revisão secundária exigida é evidenciada e impede declarar a proposta inteiramente validada; pareceres divergentes e intervenção humana devem permanecer auditáveis.

A homologação da semântica do conversor DSL, da cobertura normativa e dos testes de regressão ainda depende de [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md); esses testes não são considerados já existentes pela escolha desta política.

## Evidências de correção e reavaliação — decisão 12
Para cada divergência relevante, registrar: problema apontado, versão original, versão corrigida, **nova execução dos testes estruturais**, **nova avaliação por LLM distinto do gerador** e resultado de liberação ou bloqueio. Não considerar a correção validada com testes ou parecer de versão anterior.

Verificações negativas devem confirmar que falha estrutural, violação normativa ou divergência relevante não resolvida **impedem aceite e publicação de resultado validado**. Falta de orçamento ou de condições para novo parecer é motivo de bloqueio, não uma exceção. Critérios de relevância e de fidelidade continuam pendentes em [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md); a decisão 13 substitui o limite numérico de tentativas por ciclo condicionado aos bloqueios de segurança e custo.

## Verificabilidade do ciclo com contexto integral — decisão 13
As evidências previstas devem permitir constatar que **todas as tentativas de correção** conservam o conjunto original completo e a mesma baseline normativa/documental, modificando apenas o complemento da **última proposta e última observação de divergência**, sem omissão ou troca silenciosa de material de origem.

Cada rodada deve fornecer evidências próprias dos testes estruturais, da nova avaliação secundária, da divergência identificada, da correção apresentada e das condições financeiras autorizadas. Testar bloqueios quando houver erro normativo, ausência de proteção remota comprovada, orçamento insuficiente ou **capacidade de contexto inferior ao material integral**.

Não exigir completar todas as quatro correções quando uma resposta válida for alcançada antes: cada nova versão deve passar nos testes estruturais e na segunda avaliação, seguida da revisão humana. Esgotadas as quatro correções sem validade, a execução termina **bloqueada e não validada**. O procedimento ainda requer homologação técnica posterior. Relação: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Teste do limite de quatro iterações — decisão 14
Os cenários de evidência deverão distinguir **geração inicial** dos **quatro ciclos máximos de correção**, com nova validação estrutural e revisão por LLM em cada ciclo.

Verificar aceitação para encaminhamento humano quando a divergência for sanada na primeira, segunda, terceira ou quarta correção; e verificar bloqueio, sem quinta tentativa automática, quando a quarta não resolver a falha. Preservar também o bloqueio antecipado por orçamento, proteção remota, identidade, segurança, norma ou contexto integral incompatível, com registro auditável.

O critério é intenção de produto vinculada a [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md), não comprovação de testes implementados.

## Verificações obrigatórias na Pull Request — decisão 15
A funcionalidade futura deverá prever uma **Pull Request por sessão encerrada e aprovada pelo operador**, usando `feature/doc/{session-id}` como origem e `develop` do repositório de trabalho escolhido como destino, condicionada à resolução normativa em [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

**Verificações mínimas como objetivo de qualidade:** o resultado da PR deve tornar observável a conformidade da organização de arquivos e do escopo documental selecionado, vínculos/links, status e IDs, referência normativa correta e versionada, integridade de mudanças, ausência de segredos expostos e testes aplicáveis ao repositório. Evidências vinculadas à sessão e ao diff devem permitir identificar a causa de reprovação.

**Critério de proteção:** um check obrigatório falho, ausente ou indeterminado **impede a integração**. A proteção depende não só da execução de GitHub Actions, mas também da comprovação das configurações que **exigem** os checks para mesclar em `develop`, além da **autorização humana independente de merge**. Não presumir que checks, branch protection ou regras de revisão estejam implementados; sua definição técnica pertence a etapas futuras.

## Evidências de parada por insuficiência de contexto — decisão 16
Os critérios de qualidade deverão abranger um caso em que a documentação integral com seus complementos **não caiba na janela autorizada**, comprovando **ausência de chamada**, ausência de truncamento silencioso e bloqueio do aceite do resultado.

O operador deve poder ver a insuficiência e a distinção entre a **tarefa interrompida** e uma **nova tarefa explicitamente redefinida**, com origem, baselines e cobertura identificáveis. Testes negativos devem rejeitar reinicialização automática do mesmo ciclo para contornar as quatro correções, adoção de modelo alternativo sem aprovação, uso de orçamento anterior e promoção de conteúdo não validado.

Uma redefinição autorizada volta ao fluxo completo de validação documental e revisão humana, sem presumir homologação da equivalência semântica. Relação: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Verificação das configurações por inferência e da conversão DSL — decisão 17
A evidência de qualidade de cada inferência deverá associar **modelo/provedor/versão, prompt e origem, tamanho de contexto configurado e admissível, material e SHA de referência, autorização de custo, saída DSL, metadados de commit sugeridos, arquivos Markdown efetivamente produzidos e parecer da segunda avaliação**.

Os cenários de teste deverão verificar conversão sem perda de **títulos, seções, tabelas, links, badges, código, Mermaid, representações de gráficos e recursos externos referenciados**, preservando o conteúdo dos trechos não alterados. **DSL incompleta, Mermaid inválido, gráfico irrecuperável, caminho/escopo indevido, hash divergente ou violação normativa** deverão impedir a classificação de resultado válido e qualquer commit correspondente.

Confirmar que mudar modelo, contexto ou prompt entre inferências da mesma sessão **não altera a identidade da sessão nem apaga configurações anteriores**, não contorna autorização financeira nem reinicia automaticamente os quatro ciclos de correção da tarefa. O teste deve preservar como bloqueadas chamadas inviáveis por capacidade de contexto ou falta de proteção remota.

Os critérios de homologação da gramática e da fidelidade pertencem à [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md). Não declarar testes, conversor ou proteção Git existentes por esta decisão.

## Evidências da retomada após reprovação de PR — decisão 18
A futura validação deve demonstrar: falha de check impede merge; sessão retorna a **em correção** sem perder ID, histórico ou PR; novas inferências e commits são aceitos **somente após validação**; os checks obrigatórios são reexecutados sobre o novo SHA; aprovação anterior de encerramento não substitui a **nova aprovação humana**; autorização de merge é separada. Reprovação persistente mantém o bloqueio.

Testar também que retomada não apaga custos, não dispensa segregação de credenciais, revisão independente nem permite mais de quatro correções automáticas **na mesma inferência/tarefa**. Não presumir que workflows, proteção de branch ou a funcionalidade estejam implementados. Relação: [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

## Evidências de divergência material versus sugestão editorial — decisão 20
O produto deverá demonstrar, por casos auditáveis, que a revisão independente:
- identifica como **materiais** omissão de requisito, alteração semântica, contradição entre regras, perda na conversão DSL/Markdown, violação normativa e falhas estruturais, com bloqueio do aceite e exigência de correção/reavaliação;
- registra como **editoriais** somente sugestões de estilo, ortografia ou apresentação **sem efeito** no significado, na integridade e nas regras normativas, encaminhando-as à revisão humana sem impor correção automática;
- registra trecho afetado, baseline, justificativa e classificação, sem promover discrepâncias ambíguas ou não verificadas a «meramente editoriais»;
- conserva a geração inicial, as até **quatro correções** admissíveis, os testes estruturais repetidos, a revisão por outro LLM e os registros de custo/autorização, impedindo uma quinta rodada ou um aceite com falha material pendente.

Os resultados constituem **critérios de teste pretendidos**, não evidência de testes executados ou de critérios de relevância já homologados. Relação: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Casos de verificação da indisponibilidade do segundo LLM — decisão 27
**Evidências esperadas para aceitação futura do produto, não testes já executados:**
- Com **dois modelos distintos, homologados e disponíveis** e controles obrigatórios comprovados, a aplicação admite o início do fluxo validado, sujeitando cada chamada paga à sua própria autorização.
- Com gerador indisponível, revisor indisponível, identificação insuficiente do modelo subjacente no AgentCore, dois papéis apontando para o **mesmo LLM**, ou modelo desabilitado/não homologado, a aplicação **interrompe antes da primeira inferência documental validada**, informando o motivo verificável ao operador e registrando o bloqueio.
- É proibido tratar automaticamente a falha do revisor como autorização para uma **geração exploratória paga**, dispensar a segunda avaliação ou trocar silenciosamente o modelo indicado.
- Se a pré-verificação foi satisfeita, mas o revisor falhou **após** gerar conteúdo, o sistema preserva o histórico, registra custos/ocorrência, **não declara resultado validado e não produz commit de resultado validado**.
- A falha e a recuperação não mudam retroativamente os padrões fixados na sessão, não ignoram a herança controlada, não dispensam validações de integridade/DSL e **não reiniciam ou ampliam** as até quatro correções da mesma tarefa por artifício.

As condições financeiras remotas, credenciais, autorizações por chamada e restrições normativas/Git também precisam de evidência própria antes da operação e do aceite. Relações: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md) e [OBJ-007](OBJ-007-comparacao-controlada-de-modelos.md).

## Testes pretendidos para retomada da avaliação pendente — decisão 28
Os critérios de verificação futuros deverão demonstrar, **sem presumir testes já executados**:
- **Falha do revisor após geração íntegra:** a tarefa continua identificável, com proposta exata, baselines e documentos originais preservados, registro da interrupção e nenhum commit/aceite de resultado validado.
- **Retomada autorizada:** operador seleciona **revisor distinto, homologado e disponível**; a revisão pendente recebe **nova autorização financeira individual** e controles local/remoto comprovados; a aplicação avalia a **mesma proposta, sem gerar outra**, e vincula o parecer ao histórico original.
- **Falhas negativas:** proposta ou contexto alterado/corrompido, baseline diferente, revisor não distinto, desabilitado, sem capacidade de contexto ou sem controles financeiros suficientes **bloqueiam a retomada** e comunicam o motivo, sem modelo substituto automático.
- **Contabilidade de ciclos:** repetir apenas a revisão pendente **não consome, por si só, um ciclo de correção**, não adiciona nova geração e não reinicia o máximo de **quatro correções por tarefa**; eventual alteração corretiva real consome o ciclo correspondente e requer reavaliação.
- **Integridade e governança:** conservar hashes, uso/tokens/custos, versão de modelos/prompt, versões dos padrões carregados na sessão e escolhas por inferência; não permitir aprovação humana ou commit como substituto para avaliação por segundo LLM.

Esta decisão trata de **comportamento exigido para homologação futura**, não de implementação. Relações: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md) e [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md).

## Evidências de recuperação após reinício — decisão 29
Os testes **a serem definidos e executados posteriormente** deverão demonstrar:
- Uma tarefa com geração concluída e revisão pendente permanece **identificável pelo operador autorizado após fechar e reiniciar a aplicação**, recuperando **a mesma sessão, tarefa, proposta integral, material original, baseline, histórico, custos e contador de correções**.
- **Somente após** confirmação de integridade, identidade/permissão, revisor distinto e homologado, capacidade de contexto, limites locais e proteção remota efetivamente bloqueante, e **nova autorização financeira por chamada**, a revisão pendente pode continuar **sem gerar novamente o documento**.
- A recuperação cria **registro adicional da continuidade e da revisão**, mas não reescreve os eventos anteriores nem conta a retomada da avaliação como nova correção. Correção documental posterior continua limitada ao **saldo original de quatro ciclos**.
- Tentativas de recuperação com dados ausentes, hashes/baselines divergentes, contexto insuficiente, conteúdo excluído manualmente, credenciais indevidas, modelo inelegível ou falta de proteção financeira devem **bloquear com motivo**, sem fallback, geração automática ou aceite.
- Novo estado do perfil geral da aplicação ou mudanças posteriores de `develop` não substituem **silenciosamente** os padrões e a baseline da tarefa preservada. A eventual divergência demanda tratamento humano separado.
- Reiniciar o programa **não aciona automaticamente chamada paga nem aprova ou publica** a proposta; revisão independente e aprovação humana continuam necessárias.

Estes são **critérios de teste pretendidos**, não relatos de implementação. Relações: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md), [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md) e [OBJ-008](OBJ-008-historico-local-de-inferencias.md).

## Evidências de suspensão sem publicação — decisão 30
**Critérios de testes futuros, não evidências de execução:**
- Uma sessão com tarefa **em avaliação pendente** permite ao operador **suspender o trabalho ativo**, mas o estado de conclusão para publicação permanece **pendente/bloqueado**, sem disparar PR, commit de proposta inválida ou merge.
- Após encerrar e reiniciar a aplicação, a sessão suspensa pode ser **reaberta por operador autorizado** com a mesma identidade, tarefas e estados, baselines/hashes, propostas, custos, autorização histórica, perfis iniciais e contadores de correção.
- Reabertura **não dispara inferência paga automaticamente**. Retomada de revisão exige integridade e segundo modelo distinto/homologado, com **nova autorização financeira** por chamada e proteção remota comprovadamente bloqueante, preservando o saldo de quatro correções.
- **Negativos:** uma tarefa pendente impede concluir a sessão para publicação; cancelamento implícito por suspensão ou exclusão automática dos registros é proibido; tentativas de publicar sem resolução rastreável e aprovação humana são bloqueadas e auditadas.
- Após resolver as pendências válidas ou cancelar/retirar expressamente tarefas que não integrarão o conjunto final, exigir **nova decisão humana de conclusão**, mantendo separada a autorização de integração Git e registrando o conjunto realmente elegível.

A implementação e os contratos técnicos desses comportamentos pertencem às etapas posteriores, e nenhum fluxo de PR/branch é autorizado enquanto [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md) permanecer bloqueadora. Relações: [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md) e [OBJ-008](OBJ-008-historico-local-de-inferencias.md).
