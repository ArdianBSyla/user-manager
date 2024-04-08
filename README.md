# user-manager
This is a Fullstack project for managing users built with Golang and Vue.js. 
The backend is a RESTful API that allows users to create, read, update and delete users and search.
The frontend is a single page application that allows users to interact with the backend.

## Backend
For starting the backend server, you need to have Golang installed on your machine.
However, you need to have a MySQL database running on your machine, the steps to create the database are as followings:

1. Please create database schema and insert users by executing the commands in files `V1_create_database.sql`, `V2_create_company_table.sql`, `V3_create_user_table.sql`, in the folder `cmd` in root directory of the project. 

Now you have the database schema and some users in the database.

1. Please create a `.env` file in the root directory of the project and add the following environment variables found in `.env-template` file.

2. Install the dependencies by running the following command:
```bash
go install
```

4. Run the following command to start the backend server:
```bash
go run main.go
```

The backend server will start on port `3000`.

Here are the endpoints of the backend server:
1. `GET /api/v1/users?query={query}&page={page}&size={size}` - Get all users
2. `GET /api/v1/users/{id}` - Get a user by id
3. `POST /api/v1/users` - Create a user with the request body:
```json
{
    "company_id": 1,
    "first_name": "one",
    "last_name": "two",
    "email": "email@email.com"
}
```
4. `PUT /api/v1/users/{id}` - Update a user by id with the request body
```json
{
    "company_id": 1,
    "first_name": "one",
    "last_name": "two",
    "email": "email@email.com"
}
```
5. `DELETE /api/v1/users/{id}` - Delete a user by id

## Frontend
Frontend application is inside `client/` directory in the root directory of the project.

For starting the frontend application here are the steps:

1. Install the dependencies by running the following command:
```bash
npm install
```

2. Run the following command to start the frontend application:
```bash
npm run dev
```

The frontend application will start on port `5173`.