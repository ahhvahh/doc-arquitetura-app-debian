# OPEN-006 — Retenção e proteção do histórico de inferências

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Decisão funcional de retenção e proteção resolvida; execução dos controles ainda sujeita a validação.

## Questão
**Decidida:** manter prompts, respostas, avaliações, custos, medições e métricas JSON até exclusão manual. Pela decisão complementar de 2026-10-22, os **membros autenticados do grupo Linux autorizado** podem consultar o histórico compartilhado **exclusivamente por intermédio da aplicação/serviço**, sem acesso direto ao banco e sem alterar a proteção de segredos.

## Contexto e origem
O histórico local solicitado inclui conteúdo enviado e recebido que pode carregar informações confidenciais. A decisão de manter um banco local não resolve controle de acesso, descarte ou tratamento de segredos.

## Impacto
**Bloqueadora para decisão de Charter:** não; o responsável definiu retenção, acesso e descarte. Para produção, a implementação e validação efetivas de proteção, identidade e autorização continuam pendentes, inclusive [OPEN-001](OPEN-001-politica-credenciais.md).

## Próxima ação verificável
Verificar em Scope/Implementation os controles de acesso, a proteção de conteúdo, a extensão do descarte a backups/cópias, as autorizações de exportação e a auditoria da exclusão. Não declarar eliminação de cópias que a aplicação não controla.

## Relações
- [OBJ-008 — Histórico](../objectives/OBJ-008-historico-local-de-inferencias.md)
- [CON-005 — Persistência](../conditions/CON-005-persistencia-local-de-inferencias.md)
- [CON-002 — Credenciais](../conditions/CON-002-protecao-credenciais.md)
- [OPEN-001 — Segurança](OPEN-001-politica-credenciais.md)

## Resultado da análise — decisão C (2026-10-21)
**Triagem: RESOLVIDA.** O responsável selecionou a opção C e autorizou o registro documental nesta `develop`.

- **Retenção:** conservar prompts, respostas, avaliações e métricas JSON locais **sem expiração automática**, até solicitação manual de exclusão pelo operador autorizado.
- **Consulta:** permitir acesso ao operador e às aplicações consumidoras expressamente autorizadas; a autorização para consultar não implica permissão para excluir.
- **Proteção:** credenciais e segredos não devem integrar registros, exportações, logs nem trilhas de descarte. Trechos omitidos ou mascarados por segurança devem ser identificados sem expor seus valores.
- **Exclusão e rastreabilidade:** exigir comando ou pedido manual autorizado; conservar somente evidência mínima de quem solicitou, data, identificação dos dados afetados, alcance e resultado, sem manter o conteúdo descartado na auditoria.
- **Cópias:** indicar o alcance real da exclusão em backups e exportações; não presumir descarte de cópias sob controle externo.
- **Interoperabilidade:** as métricas JSON obedecem à mesma regra de retenção, acesso e exclusão manual dos registros de inferência.

**Destinação canônica:** [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md) e [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md). Mecanismos técnicos para Scope e Implementation.

**Situação do documento:** decisão de produto concluída, mas o badge permanece em `refinement`, pois falta o asset `badge-status-refined.svg` em `docs/assets/` e sua criação está fora do escopo desta alteração de Charter. O encerramento da análise não é comprovação de implementação; pendências de segurança continuam em [OPEN-001](OPEN-001-politica-credenciais.md).

## Complemento vigente — leitura compartilhada por grupo Linux via aplicação (2026-10-22)

**Origem:** aprovação expressa da simplificação da autorização pelo grupo Linux do servidor administrado, sem acesso direto ao banco por usuários comuns.

- A decisão original **C de retenção até exclusão manual** permanece integralmente vigente, assim como o descarte manual autorizado e a auditoria.
- **Consulta:** todos os usuários Linux autenticados e associados ao grupo autorizado podem consultar pela aplicação tarefas e resultados compartilhados, mesmo quando solicitados por outro integrante. A inclusão ou exclusão do grupo cabe à administração Linux; a aplicação não cria concessões individuais de leitura.
- **Meio obrigatório:** o usuário comum e o guardião opcional recebem dados somente pela aplicação/serviço, que utiliza a identidade técnica autorizada para consultar o banco local. A participação no grupo não dá acesso direto aos arquivos, credenciais ou consultas do banco.
- **Limites:** consulta não concede exclusão de registros, leitura de segredos, autorização financeira, controle administrativo ou publicação Git. Preservar autoria e trilha de auditoria.
- **Histórico:** os termos anteriores “operador” e “consumidores expressamente autorizados” permanecem como registro da decisão original, mas a **forma vigente de autorizar a consulta compartilhada** é a associação ao grupo Linux, sem permissões individuais por tarefa.

**Pendência:** verificar tecnicamente no ambiente controlado autenticação/grupo, restrições ao acesso direto, exposição da aplicação e rastreabilidade. `DOCUMENTED`, não `IMPLEMENTED`; badge permanece em `refinement`.

**Relações:** [OPEN-001](OPEN-001-politica-credenciais.md), [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md) e [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md).

## Complemento aprovado — backup e restauração manuais do histórico (2026-10-22)

**Origem:** Backup/restauração **1A/2A/3A**, complementando Continuidade **3A** e a anterior decisão **C de retenção até exclusão manual**. Não reabrir decisão de retenção ou autorização de consulta.

- **Backup:** somente **manual por administrador autorizado**, em **diretório local protegido e separado dos dados ativos**. A cópia do banco deverá ser consistente, verificável, recuperável e inacessível a usuários comuns. Sem rotina automática ou armazenamento externo obrigatório na primeira entrega.
- **Restauração:** **manual supervisionada**, com autorização administrativa, verificação de origem/integridade antes de substituir dados e registro do resultado. Não executar tarefas, abrir sessões ou autorizar despesa por efeito da restauração.
- **Corrupção:** bloquear operações dependentes sem reparar/restaurar automaticamente nem remover evidências. Reativação requer regularização demonstrável e liberação autorizada.
- **Sigilo e descarte:** cópias não conterão segredos administrativos, permanecendo sujeitas à proteção e à retenção **até exclusão manual autorizada**. Cada descarte registrará alcance real sobre backups e outras cópias controladas pela aplicação; não presumir eliminação de cópias externas nem substituir a decisão original.
- **Limitação aceita:** cópias no mesmo servidor não oferecem garantia contra perda total do equipamento ou comprometimento simultâneo da base e da cópia. Backups do histórico não são snapshots documentais em `/docs/backup/`.

**Verificação pendente:** consistência e restauração das cópias, negativa de acesso a clientes, integridade de origem, falha segura, autoria da exclusão e extensão real do descarte. Tecnologia/método físico pertencem a Scope/Implementation.

**Estado:** decisão `DOCUMENTED`, não `IMPLEMENTED`. Badge `refinement` preservado porque os critérios ainda não foram comprovados e o asset `badge-status-refined.svg` não existe na baseline observada. Relações: [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md), [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md), [OPEN-003](OPEN-003-politica-de-branches.md).
