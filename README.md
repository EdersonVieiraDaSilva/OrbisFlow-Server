# OrbisFlow Server

Aplicação backend escrita em Go para gerenciamento e distribuição de arquivos de mídia, oferecendo streaming eficiente e suporte a requisições HTTP Range.

### Funcionalidades
* Streaming e multimídia
	- Suporte Multi-formato: Reprodução fluida de arquivos de vídeo (.mp4, .mkv, .webm) e áudio (.mp3).
	- Seek & Play (Controle de Tempo): Implementação de HTTP Range Requests, permitindo avançar ou retroceder a mídia instantaneamente sem carregar o arquivo inteiro.
	- Streaming de Baixo Consumo: Utilização de buffers circulares para garantir que arquivos grandes (GBs) consumam apenas alguns MiB de RAM.
	
* Segurança e Proteção
	- Autenticação Basic Auth: Acesso restrito via usuário e senha, protegendo o conteúdo contra acessos não autorizados via rede externa.
	- Sanitização de Caminhos (Anti-Path Traversal): Proteção contra invasão de diretórios, garantindo que apenas arquivos na pasta designada sejam acessados.
	- Filtro de Extensões (Whitelist): Sistema de segurança que bloqueia a execução ou download de arquivos sensíveis do sistema (como .go, .log ou .exe).	
	
* Monitoramento e CLI
	- Terminal Interativo: Interface de linha de comando (CLI) integrada para gerenciamento do servidor em tempo real.
	- Show Data: Visualização detalhada do consumo de memória RAM e recursos do sistema.
	- Show Clients: Rastreamento de conexões ativas, exibindo IP do cliente, arquivo solicitado e horário do acesso.
	- Logs de Auditoria: Registro automático de todas as atividades em arquivos de log externos para análise posterior.	
	
### Fluxo de uma Requisição

1- O cliente solicita um arquivo via navegador/VLC.

2- O Middleware de Segurança valida as credenciais.

3- O Sanitizador limpa o nome do arquivo e verifica a extensão.

4- O Monitor registra o início da sessão e o IP.

5- O Streamer entrega o conteúdo em pedaços (chunks) via protocolo HTTP.	
	
### Arquitetura
* Estrutura do projeto
	- main.go: Ponto de entrada que gerencia o terminal interativo (CLI).
	- server.go: Configuração do servidor HTTP e roteamento protegido.
	- security.go: Camada de segurança (Auth, Sanitização e Filtros).
	- streamer.go: Motor de streaming e manipulação de buffers de mídia.
	- monitor.go: Gestão de métricas de sistema e logs de acesso.	
	
### Linguagem
- Go (Golang)

### Tecnologias
Este projeto foi desenvolvido utilizando a linguagem Go (Golang), aproveitando sua alta performance e suporte nativo à concorrência. Uma decisão de design importante foi o uso exclusivo de bibliotecas padrão, garantindo um binário leve e sem dependências externas.
	- net/http: Motor principal para o servidor de streaming e gerenciamento de protocolos.
	- time: Visa coletar dados de horário do relógio do Sistema Operacional, e o atribui a aos logs de auditoria e monitoramento de sessões.
	- os: Manipulação eficiente de arquivos no sistema operacional e fluxo de dados (streaming). 
	- path/filepath: Garantia de segurança no tratamento de nomes de arquivos e prevenção de ataques de diretório.
	- runtime: Coleta de métricas de hardware em tempo real (RAM e uso de CPU).
	
### Como Executar o Projeto
* Pré-requisitos: 
	- Certifique-se de ter o Go instalado (versão 1.18 ou superior).

* Preparação:
	- Crie uma pasta chamada Videos_OrbisFlow no diretório raiz do projeto.
	- Adicione seus arquivos de mídia (.mp4, .mp3, .mkv) dentro desta pasta.

* Execução:
	- go run .
* Acesso:
	- Abra o navegador ou VLC.
	- Acesse: http://localhost:8080/nome_do_arquivo.mp4
	- Credenciais: Use o usuário e senha definidos no módulo security.go.	

------------------------------------------------

