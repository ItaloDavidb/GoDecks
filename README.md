# GoDecks

## Descrição

Este projeto é uma aplicação backend desenvolvida em Go para gerenciar cartas de Pokémon, usuários e decks de cartas. Utiliza PostgreSQL como banco de dados e o GORM para interação com o banco. A aplicação suporta a criação, leitura e exclusão de cartas e decks, além de associar cartas aos usuários e verificar a disponibilidade de cartas em decks.

## Estrutura do Banco de Dados

As seguintes tabelas são criadas e gerenciadas:

- **users**: Armazena informações dos usuários.
- **cards**: Armazena informações das cartas de Pokémon.
- **user_cards**: Relaciona cartas aos usuários e mantém a quantidade de cada carta.
- **decks**: Armazena os decks criados pelos usuários.
- **deck_cards**: Relaciona cartas aos decks e mantém a quantidade de cada carta em um deck.

## Estrutura do Projeto

- **controllers**: Contém funções para manipulação de dados e resposta a requisições HTTP.
- **database**: Configura a conexão com o banco de dados.
- **models**: Define os modelos de dados e suas relações.
- **routes**: Configura as rotas e suas respectivas funções.

## Rotas Disponíveis

- `POST /api/Cards`: Cria uma nova carta ou várias cartas.
- `GET /api/Cards`: Obtém todas as cartas ou pesquisa cartas pelo nome, set_code ou number.
- `DELETE /api/Cards`: Deleta uma ou várias cartas com base na combinação de set_code e number.

## Exemplo de Payload

### Criar Múltiplas Cartas

```json
[
    {
        "set_code": "PAR",
        "number": "18",
        "name": "Wo-Chien",
        "type": "Pokémon",
        "json_data": "{}"
    },
    {
        "set_code": "BRS",
        "number": "121",
        "name": "Bibarel",
        "type": "Pokémon",
        "json_data": "{}"
    }
]
```
### Excluir Múltiplas Cartas

```json
{
    "cards": [
        "PAR18",
        "BRS121"
    ]
}
```
# Inicialização do Projeto

Clone o repositório:

```bash
git clone https://github.com/ItaloDavidb/GoDecks.git
cd GoDecks
```

### Inicialize o banco de dados:

```
docker-compose up
```
### Instale as dependências do Go:
```
go mod tidy
```
### Inicie o Servidor 
```
go run main.go
```
### Acesse a aplicação na URL:
```
http://localhost:8080


