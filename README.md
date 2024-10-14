# luizalabs-technical-test

Olá, pessoal! Realizei este teste técnico com base em todo o meu conhecimento, com o objetivo de apresentar o meu melhor com o uso da liguagem GO. Espero que gostem!

Se houver alguma observação a ser feita, peço que me enviem para que eu possa realizar as correções necessárias. Agradeço desde já.

## 1 - Explicação técnica:

### Descrição

Este projeto tem como objetivo criar um serviço para a busca rápida de CEPs, utilizando as principais APIs de CEP disponíveis no Brasil.

A linguagem Go foi escolhida para realizar buscas concorrentes devido ao uso de goroutines e channels, além do seu escalonamento de tarefas de forma preemptiva. Sua agilidade, proporcionada por um runtime compacto e suas dependências reduzidas, juntamente com o uso de multi-stage builds no Docker, potencializa ainda mais a ferramenta. Fora isso, Go suporta HTTP/2, que torna a comunicação mais rápida, e, ao não delegar as threads do sistema operacional, gera lightweight threads, permitindo um processo de resposta ágil para esta aplicação.

Se um CEP válido não retornar um endereço, o sistema substituirá, um a um, os dígitos da direita para a esquerda por zero até encontrar um endereço correspondente, aumentando a ordem de grandeza. Por exemplo, ao fornecer o CEP 14570006 e não obter resultado, o sistema tentará variações como 14570000, 14500000 e assim por diante, até ter sucesso, retornando as informações da cidade de Buritizal neste caso.

### Como rodar o projeto

### Arquitetura e documentação

Para a arquitetura do projeto, utilizei a convenção da comunidade proposta no repositório https://github.com/golang-standards/project-layout. Essa mesma abordagem é utilizada em grandes projetos (como Kubernetes) para manter um padrão, além de empregar técnicas já validadas por pesquisas de benchmark.

## 2 - Como funcionam as requisições HTTP entre o Client e o Server:

(usei o exemplo do própio site da netshoes, conforme o enunciado)

- Digitar a URL: Você digita "http://www.netshoes.com.br" no navegador.
- Resolução de DNS: O navegador transforma "www.netshoes.com.br" em um endereço IP.
- Conexão: O navegador se conecta ao servidor da Netshoes.
- Envio da solicitação HTTP: O navegador manda uma solicitação HTTP para o servidor, pedindo a página inicial (ex.: GET).
- Processamento: O servidor da Netshoes recebe e processa a solicitação.
- Recebimento da resposta: O navegador recebe a resposta com a página da Netshoes.
- Renderização: O navegador exibe a página HTML recebida.
- Exibição: A página do site da Netshoes aparece para você.
