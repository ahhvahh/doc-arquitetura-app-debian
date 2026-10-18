# Doc Arquitetura — Aplicação Debian

Aplicação local em Go para compilar a referência normativa ProjectCharter/1.0 em um pacote DSL e auditar a documentação de repositórios GitHub nas etapas Charter, Scope e Implementation.

## Organização das branches

- `main`: base inicial do repositório, sem aplicação publicada.
- `develop`: desenvolvimento, documentação do projeto, código-fonte e testes.

## Referência normativa

A fonte normativa é [ahhvahh/doc-arquitetura](https://github.com/ahhvahh/doc-arquitetura/tree/main/ProjectCharter/1.0), sempre lida na branch `main`.

## Estado

Repositório inicializado. A especificação, migração do protótipo Go, integração com modelos locais e scripts Debian deverão ser desenvolvidos em `develop`, após validação dos contratos do projeto.
