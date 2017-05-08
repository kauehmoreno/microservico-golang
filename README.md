# Aplicação microservico-golang - Materia
##Projeto de Monografia - API feita em Golang
>Este é o projeto final de faculdade
>-Está aplicação servirá como base para o testes de performance entre aplicações monolíticas vs microserviço
>Aplicação Microserviço-golang feita com go==1.8.1 e mongo

`Cenários de testes entre arquiteturas`

| Testes | Ferramentas | arquitetura |
| -------| ----------- | ----------- |
| Números de requisições | apachebanchmark, pagespeedtest | Microserviço |
| Requisições paralelas | apachebanchmark | Microserviço |
| Tempo de resposta backend | usando modulo python, speedcurve | Microserviço |
| Matérias com fotos | apachebanchmark,wrk, PagesppedTest | Microserviço |

`Rotas da Aplicação`

| Função  |  parametros |
| -------- | ----------- |
| Todas as matérias | Filtro per_page, Filtro order_by(default será publicacao), uuid retorná uma única matéria |
| Rota para post | Rota que irá salvar as matérias criadas, aceitará uma lista quanto uma única matéria para procesar |

` Autor - Kauêh Moreno`


