# Instruções para execução
Caso tenha o Go e o docker instalado em sua máquina apenas execute o arquivo: ./cria_env.sh

# Introdução
## Arquitetura
![Netflix Case](https://app.lucidchart.com/publicSegments/view/dd021a71-81fa-4ce7-aec9-669c3b5d43d1/image.png)

# Url para os servicos
## Serviço de Filmes
1. Procurar por genero:
GET localhost:8088/movie/genre?genre=comedy
2. Buscar descrição:
GET localhost:8088/movie/desc/12
3. Palavra chave:
GET localhost:8088/movie/search?key=porno
4. Top por categoria:
GET localhost:8088/movie/top?genre=romance

## Serviço de Usuário
1. Votar em um filme:
PUT localhost:8089/user/vote '{"userId": 1, "moveId": 2}'
2. Adicionar para assistir depois:
PUT localhost:8089/user/addlist '{"userId": 2, "movieId": 8}'
3. Verificar os filmes assistidos por um usuário:
GET localhost:8089/user/watched/1

## Serviço de Chamados
1. Mostrar chamado por id:
GET localhost:8090/issue/1
2. Inserir um novo chamado:
POST localhost:8090/issue -d '{"username": "xxxxx", "desc": "Error ao abrir aplicativo"}'
