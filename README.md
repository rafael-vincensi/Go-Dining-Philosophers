# Jantar dos Filósofos

> Projeto desenvolvido em Go para demonstrar problemas clássicos de concorrência e sincronização em sistemas operacionais e sistemas distribuídos. Baseado no Problema do Jantar dos Filósofos, a aplicação simula cinco filósofos disputando recursos compartilhados (garfos), reproduzindo um cenário de deadlock e implementando uma solução baseada em semáforo para preveni-lo.

## Tecnologias Utilizadas

<div align="center">

[![Go](https://img.shields.io/badge/Go-1.27-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Goroutines](https://img.shields.io/badge/Goroutines-Concorrência-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/tour/concurrency/1)
[![Channels](https://img.shields.io/badge/Channels-Sincronização-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/tour/concurrency/2)
[![Go Mutex](https://img.shields.io/badge/Sync-Mutex-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/sync)

</div>

## Funcionalidades

- Simulação de cinco filósofos executando concorrentemente por meio de Goroutines.
- Gerenciamento de recursos compartilhados através de garfos protegidos por Mutex.
- Reprodução controlada do cenário de deadlock do Problema do Jantar dos Filósofos.
- Prevenção de deadlock utilizando semáforo implementado com canais (Channels).
- Sincronização entre rotinas concorrentes.
- Simulação dos estados dos filósofos: pensando, aguardando e comendo.
- Dois cenários independentes de execução para comparação.

## Arquitetura

A aplicação está organizada em três pacotes principais:

- `philosopher`: Define a estrutura dos filósofos, seus estados e comportamentos, incluindo pensar, adquirir garfos, comer e liberar recursos.
- `fork`: Representa os garfos compartilhados, utilizando `sync.Mutex` para garantir a exclusão mútua durante o acesso aos recursos.
- `simulation`: Responsável por inicializar os filósofos e garfos, coordenar as Goroutines, reproduzir o deadlock e implementar a estratégia de prevenção baseada em semáforo.

O arquivo `main.go` atua como ponto de entrada da aplicação, executando os dois cenários sequencialmente.

## Concorrência e Sincronização

O projeto demonstra duas abordagens para o Problema do Jantar dos Filósofos.

### 1. Cenário de Deadlock

O primeiro cenário reproduz intencionalmente uma situação de deadlock, coordenando os cinco filósofos para que todos adquiram seus respectivos garfos esquerdos antes de tentar adquirir os garfos direitos.

Quando todos os filósofos estão segurando um garfo, cada um passa a aguardar indefinidamente pelo recurso que está em posse de seu vizinho, estabelecendo uma condição de espera circular.

Um canal de sincronização garante que todos os filósofos adquiram o primeiro garfo antes de serem liberados para tentar adquirir o segundo. Um mecanismo de timeout é utilizado para identificar e informar o deadlock.

Esse cenário demonstra as quatro condições necessárias para a ocorrência de deadlock:

- **Exclusão Mútua:** Cada garfo pode ser utilizado por apenas um filósofo por vez.
- **Posse e Espera:** Os filósofos mantêm um garfo enquanto aguardam a aquisição de outro.
- **Não Preempção:** Os garfos não podem ser retirados à força de um filósofo.
- **Espera Circular:** Cada filósofo aguarda por um garfo que está sendo utilizado pelo próximo filósofo na disposição circular.

### 2. Prevenção de Deadlock

O segundo cenário implementa um semáforo utilizando um canal Go com buffer e capacidade para quatro elementos.

Ao permitir que, no máximo, quatro filósofos tentem adquirir os garfos simultaneamente, a implementação impede que os cinco filósofos mantenham um garfo enquanto aguardam pelo segundo recurso.

Essa estratégia quebra a condição de espera circular e impede a ocorrência de deadlock.

Cada filósofo realiza dois ciclos de alimentação, enquanto um `sync.WaitGroup` coordena a conclusão de todas as Goroutines.

## Estrutura do Projeto

```text
dining-philosophers/
│
├── fork/
│   └── fork.go
│
├── philosopher/
│   └── philosopher.go
│
├── simulation/
│   └── simulation.go
│
├── main.go
├── go.mod
└── .gitignore
```

## Como Executar

### Pré-requisitos

- Go 1.27 ou versão compatível.
- Git (opcional, para clonar o repositório).

### Clonar o Repositório

```bash
git clone https://github.com/rafael-vincensi/dining-philosophers.git
cd dining-philosophers
```

### Executar a Aplicação

Execute o comando abaixo na raiz do projeto:

```bash
go run .
```

O programa executará os dois cenários sequencialmente.

### Compilar a Aplicação

Para verificar a compilação de todos os pacotes:

```bash
go build ./...
```

## Resultados Esperados

A aplicação apresenta dois cenários principais:

**Cenário 1 — Demonstração de Deadlock**

Os cinco filósofos adquirem seus respectivos garfos esquerdos e, em seguida, tentam adquirir os garfos direitos. Como cada recurso necessário está em posse de um filósofo vizinho, todos acabam bloqueados.

Saída esperada:

```text
All philosophers acquired their left forks.

Allowing philosophers to acquire their right forks...

DEADLOCK DETECTED!
All philosophers are waiting for their right forks.
```

**Cenário 2 — Prevenção de Deadlock**

O semáforo limita a quantidade de filósofos que podem tentar adquirir os garfos simultaneamente a quatro, permitindo que a simulação avance sem ocorrência de deadlock.

Saída esperada:

```text
All philosophers finished eating!
Deadlock successfully prevented.
```

## Objetivo Acadêmico

A aplicação demonstra como rotinas concorrentes disputam recursos compartilhados e como mecanismos de sincronização podem ser utilizados para evitar situações de deadlock.

---

<div align="center">

Desenvolvido por [rafael-vincensi](https://github.com/rafael-vincensi) e Marcos Felipe Schmid.

</div>
