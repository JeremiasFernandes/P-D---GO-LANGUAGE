## Pré-requisitos

Antes de começar, certifique-se de ter os seguintes pré-requisitos instalados no seu sistema:


- [Go](https://golang.org/dl/):

## Instalação:

Para instalar o go no Ubuntu: 
    sudo apt-get update && sudo apt-get -y install golang-go

## Executando o Aplicativo

    go run main.go

O aplicativo estará acessível em 
    http://localhost:8080.

## Testando o Aplicativo

Se preferir, após executar o projeto, visite: 
    http://localhost:8080/swagger/index.html# 
para ver e testar todos os endpoints.

O projeto oferece endpoints REST para criar, listar, atualizar e excluir usuários. 


## Modelos de Dados:

### request.UserLogin:

Estrutura contendo os campos necessários para o login do usuário.

- `email` (string, obrigatório): O email do usuário (deve ser um endereço de email válido).
- `password`  (string, obrigatório): A senha do usuário (deve ter pelo menos 6 caracteres e conter pelo menos um dos caracteres: !@#$%*).

### request.UserRequest:

Estrutura contendo os campos necessários para criar um novo usuário.

- `age` (inteiro, obrigatório): A idade do usuário (deve estar entre 1 e 140).
- `email`  (string, obrigatório): O email do usuário (deve ser um endereço de email válido).
- `name` (string, obrigatório): O nome do usuário (deve ter pelo menos 4 caracteres e no máximo 100 caracteres).
- `password` (string, obrigatório): A senha do usuário (deve ter pelo menos 6 caracteres e conter pelo menos um dos caracteres: !@#$%*).

### request.UserUpdateRequest:

Estrutura contendo campos para atualizar as informações do usuário.

- `age` (inteiro, obrigatório): A idade do usuário (deve estar entre 1 e 140).
- `name`  (string, obrigatório): O nome do usuário (deve ter pelo menos 4 caracteres e no máximo 100 caracteres).

### response.UserResponse:

Estrutura de resposta contendo informações do usuário.

- `age` (inteiro): A idade do usuário.
- `email`  (string): O email do usuário.
- `id`  (string): O ID único do usuário.
- `name` (string): O nome do usuário.

### rest_err.Causes:

Estrutura representando as causas de um erro.

- `field`  (string): O campo associado à causa do erro.
- `message` (string): Mensagem de erro descrevendo a causa.

### rest_err.RestErr:

Estrutura descrevendo por que ocorreu um erro.

- `causes` (array de rest_err.Causes): Causas do erro.
- `code` (inteiro): Código do erro.
- `error` (string): Descrição do erro.
- `message` (string): Mensagem de erro.

## Endpoints:

- Para autenticação, você deve incluir o token de acesso no cabeçalho Authorization como "Bearer <Inserir token de acesso aqui>" para os endpoints que necessitam do token.

A API oferece os seguintes endpoints:

1. **POST /createUser**
    - Descrição: Cria um novo usuário com as informações fornecidas.
    - Parâmetros:
        - `userRequest` (corpo, obrigatório): Informações do usuário para registro.
    - Respostas:
        - 200: OK (Usuário criado com sucesso)
        - 400: Requisição Inválida (Erro na requisição)
        - 500: Erro Interno do Servidor (Erro interno do servidor)
2. **DELETE /deleteUser/{userId}**
    - Descrição: Exclui um usuário com base no ID fornecido como parâmetro.
    - Parâmetros:
        - `userId` (caminho, obrigatório): ID do usuário a ser excluído.
    - Respostas:
        - 200: OK (Usuário excluído com sucesso)
        - 400: Requisição Inválida (Erro na requisição)
        - 500: Erro Interno do Servidor

3. **GET /getUserByEmail/{userEmail}**
    - Descrição: Retorna os dados do usuário baseado no email fornecido.
    - Parâmetros:
        - `userEmail` (caminho, obrigatório): Email do usuário a ser buscado
    - Responses:
        - 200: Ok (usuário encontrado e retornado com sucesso)
        - 400: Erro: Email inválido
        - 404: Usuário não encontrado
4. **GET /getUserById/{userId}**
    - Descrição: Retorna os dados do usuário baseado no ID.
    - Parâmetros:
        - `userId` (path, required): O ID do usuário a ser buscado.
    - Resposta:
        - 200: Ok (usuário encontrado e retornado com sucesso)
        - 400: Erro: ID inválido
        - 404: Usuário não encontrado

5. **POST /login**
    - Descrição: Login do usuário, retorna um token de sessão.
    - Parâmetros:
        - `userLogin` (body, required): Credênciais do usuário.
    - Resposta:
        - 200: Login feito com sucesso
        - 403: Erro: Credenciais inválidas

6. **PUT /updateUser/{userId}**
    - Descrição: Atualização do usuário a partir do ID.
    - Parâmetros:
      - `userId` (path, required): ID do usuário a ser atualizado.
      - `userRequest` (body, required): Informações a serem atualizadas do usuário.
    - Respostas:
      - 200: OK (Usuário atualizado com sucesso)
      - 400: Request inválido 
      - 500: Erro Interno do Servidor