# API-Service

API para estudo, temos end-points para cadastro de usuários, e geração de token, essas duas são publicas.

Para produtos temos um CRUD, onde todas as rotas são protegidas pelo token JWT.

## Tests Unitários

Para executar os testes:

```
go test --count=1 ./...
```
![image](https://github.com/user-attachments/assets/39d4b14b-44fb-4446-be4e-c304d9ea82f8)


## Documentação

Para visualizar a documentação temos dentro da pasta API os arquivos da documentação gerados pelo swegger, mas também temos a opção de visualizar no navegador por meio da rota "http://localhost:8000/docs/index.html"
![image](https://github.com/user-attachments/assets/4d3f8c88-293d-4813-8094-f7f637251baf)

## Banco de dados

Para utilização do banco de dados não é necessário a configuração, pois esse projeto utiliza um banco sqlite. A menos que você não tenha sqlite instalado.
