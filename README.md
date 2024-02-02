# Desafio 1 - Client Server API (Pós Graduação GoExpert)

### DESCRIÇÃO DO DESAFIO
 
Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.
 
Você precisará nos entregar dois sistemas em Go:
- client.go
- server.go
 
Os requisitos para cumprir este desafio são:
 
O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
 
O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
 
Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
 
O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.
 
Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.
 
O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

### PRÉ-REQUISITOS

#### 1. Instalar o GO no sistema operacional.

É possível encontrar todas as instruções de como baixar e instalar o GO nos sistemas operacionais Windows, Mac ou Linux [aqui](https://go.dev/doc/install).

#### 2. Clonar o repositório:

```
git clone git@github.com:raphapaulino/pos-graduacao-goexpert-desafio-1-client-server-api.git
```

#### 3. Criar o banco de dados Sqlite para uso nesse projeto

À partir da raiz do projeto, criar o arquivo do banco utilizando o seguinte comando via terminal:

```
touch posgosqlite-2.db
```

### EXECUTANDO O PROJETO

1. Estando na raiz do projeto, baixar as dependências:

```
go mod tidy
```

2. Na sequência, acessar o diretório **server**, gerar o programa servidor e executá-lo:

```
cd server
```

```
go build -o server
```

```
./server
```

3. Feitos os passos anteriores, em uma outra aba do terminal, acessar o diretório **client** (à partir da raiz do projeto), gerar o programa client e executá-lo da seguinte forma:

```
cd client
```

```
go build -o client
```

```
./client
```


### PLUS: COMO VERIFICAR AS COTAÇÕES GRAVADAS NO BANCO?

À partir da raiz do projeto, acessar o diretório **server**, executar a leitura ao banco de dados, na command line do sqlite3 visualizar as tabelas e depois visualizar os registros da tabela **quotations**:

```
cd server
```

```
sqlite3 posgosqlite-2.db
```

```
sqlite> .tables
```

```
sqlite> SELECT * FROM quotations;
```

That's all folks! : )


## Contacts

[LinkedIn](https://www.linkedin.com/in/raphaelalvespaulino/)

[GitHub](https://github.com/raphapaulino/)

[My Portfolio](https://www.raphaelpaulino.com.br/)
