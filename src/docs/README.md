# Documentação técnica do protótipo

- [Arquitetura](ARCHITECTURE.md)
- [Plano de implementação](IMPLEMENTATION-PLAN.md)

Esta documentação descreve o aplicativo migrado, não substitui a referência normativa consultada em tempo de execução.

## Evolução para execução em serviço

- [Plano do serviço systemd isolado e autorização Unix](SERVICE-ISOLATION-PLAN.md) — **refinement, não implementado**.
- [Contrato de branches por operação e eventos de progresso](TASK-BRANCH-LIFECYCLE.md) — **refinement, não implementado**.

- [Administração root e cadastro de usuários, grupos e credenciais GitHub](ADMIN-SETUP.md) — **refinement, não implementado**. Define `sudo doc-arquitetura admin` e `/etc/doc-arquitetura-app/config.yaml`.

- [Conexão GitHub por arquivo local no aplicativo](GITHUB-CONFIG.md) — suporte à CLI legado, não usa configuração global do Git.
