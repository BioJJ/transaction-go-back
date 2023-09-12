# Transaction-go-back

Este é um projeto em Golang que implementa um CRUD de usuários com autenticação usando MongoDB.

## Rotas

### 1. Buscar usuário por ID

- Método: GET
- Rota: `/users/find/:userId`
- Middleware: `model.VerifyTokenMiddleware`
- Controlador: `userController.FindUserByID`

### 2. Buscar usuário por Email

- Método: GET
- Rota: `/users/:userEmail/email`
- Middleware: `model.VerifyTokenMiddleware`
- Controlador: `userController.FindUserByEmail`

### 3. Criar usuário

- Método: POST
- Rota: `/users`
- Controlador: `userController.CreateUser`

### 4. Atualizar usuário por ID

- Método: PUT
- Rota: `/users/:userId`
- Middleware: `model.VerifyTokenMiddleware`
- Controlador: `userController.UpdateUser`

### 5. Deletar usuário por ID

- Método: DELETE
- Rota: `/users/:userId`
- Middleware: `model.VerifyTokenMiddleware`
- Controlador: `userController.DeleteUser`

### 6. Autenticação de usuário

- Método: POST
- Rota: `/login`
- Controlador: `userController.LoginUserServices`

## Banco de Dados

Este projeto utiliza o MongoDB como banco de dados para armazenar informações de usuário e autenticação.

## Como Executar

Certifique-se de ter o Golang instalado e o MongoDB configurado corretamente antes de executar o projeto.

1. Clone o repositório:

- git clone https://github.com/seu-usuario/transaction-go-back.git

2. Certifique-se de criar um arquivo .env na raiz do seu projeto e adicionar essas variáveis de ambiente com seus valores correspondentes. Essas variáveis serão usadas para configurar a conexão com o MongoDB e a chave secreta JWT no seu aplicativo Go.

- MONGODB_URL=
- MONGODB_USER_DB=
- MONGODB_USER_COLLECTION=

JWT_SECRET_KEY=

3. Com Golang instalado:

- go run main.go

4. O servidor estará em execução em `http://localhost:8081`.

## Contribuição

Sinta-se à vontade para contribuir para este projeto. Se você tiver sugestões de melhorias, novos recursos ou correções de bugs, abra uma issue ou envie um pull request.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).
