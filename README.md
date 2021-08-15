# Go Todo API Documentation

| Method | Route           | Description                                  |
| :----- | :-------------- | :------------------------------------------- |
| POST   | /users/register | Create new user account to access Go Todo |
| POST   | /users/login    | User login to access Go Todo              |
| POST   | /todos          | Add new todo to Go Todo                   |
| GET    | /todos          | Show all todos in Go Todo                 |
| GET    | /todos/:id      | Show todo in Go Todo by ID                |
| PUT    | /todos/:id      | Update all todo field in Go Todo          |
| PATCH  | /todos/:id      | Update only todo status field in Go Todo  |
| DELETE | /todos/:id      | Delete todo from Go Todo                  |

---

## Register

Create new user account to access Go Todo

- **URL**

  `/users/register`

- **Method:**

  `POST`

- **URL Params**: none

- **Data Params**: none

- **Request Body**

  ````json
    {
      "email": "<user email>",
      "password": "<user password>"
    }
    ```

  ````

- **Success Response:**

  - **Code:** 201 <br />
    **Content:**
    ```json
    {
      "success": true,
      "user": {
        "id": "<user ID>",
        "email": "<user email>"
      }
    },
    ```

- **Error Response:**

  - **Code:** 400 BAD REQUEST <br />
    **Content:** `"message": "Email <user email> is already registered"`

    OR

  - **Code:** 400 BAD REQUEST <br />
    **Content:** `"message": "Email cannot be empty/null"`

    OR

  - **Code:** 400 BAD REQUEST <br />
    **Content:** `"message": "Email format is wrong"`

    OR

  - **Code:** 400 BAD REQUEST <br />
    **Content:** `"message": "Password cannot be empty/null"`

    OR

  - **Code:** 400 BAD REQUEST <br />
    **Content:** `"message": "Password minimum character is 6"`

    OR

  - **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `"message": "Internal server error`

- **Sample Call:**

  Request body:

  ```json
  {
    "email": "user1@mail.com",
    "password": "password1"
  }
  ```

  Response:

  ```json
  {
    "success": true,
    "user": {
      "id": 14,
      "email": "user1"
    }
  }
  ```

- **Notes:** none

---

## Login

Login to access Go Todo

- **URL**

  `/users/login`

- **Method:**

  `POST`

- **URL Params**: none

- **Data Params**: none

- **Request Body**

  ````json
    {
      "email": "<user email>",
      "password": "<user password>"
    }
    ```

  ````

- **Success Response:**

  - **Code:** 201 <br />
    **Content:**
    ```json
    {
      "success": true,
      "access_token": "<access token>"
    },
    ```

- **Error Response:**

  - **Code:** 403 FORBIDDEN <br />
    **Content:** `"message": "Email is not registered"`

  OR

  - **Code:** 403 FORBIDDEN <br />
    **Content:** `"message": "Wrong password"`

    OR

  - **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `"message": "Internal server error`

- **Sample Call:**

  Request body:

  ```json
  {
    "email": "user1",
    "password": "password1"
  }
  ```

  Response:

  ```json
  {
    "success": true,
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWF0IjoxNjIyMjEyMzk5fQ.aWfNArS1JYnNYkDxrIyqFBqWuwxEKfEmFYs65t0bcjs"
  }
  ```

- **Notes:** none

---

## Add Todo

Add new todo to Go Todo

- **URL**

  `/todos`

- **Method:**

  `POST`

- **URL Params**: none

- **Data Params**: none

- **Request Body**

  ````json
    {
      "title": "<todo title>",
      "description": "<todo description>",
      "status": "<todo status: done/undone>",
      "due_date": "<todo due date, format: YYYY-MM-DD>"
    }
    ```

  ````

- **Success Response:**

  - **Code:** 201 <br />
    **Content:**
    ```json
    {
      "message": created
      "data":
        {
          "id": "<id number>",
          "title": "<todo title>",
          "description": "<todo description>",
          "status": "<todo status: done/undone>",
          "due_date": "<todo due date>",
          "createdAt": "2021-05-24T15:01:21.735Z",
          "updatedAt": "2021-05-24T15:01:21.735Z"
        }
    },
    ```

- **Error Response:**

  - **Code:** 400 BAD REQUEST <br />
    **Content:** `"message": "Validation error: Due date cannot be the day before today"`

    OR

  - **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `"message": "Internal server error`

- **Sample Call:**

  Request body:

  ```json
  {
    "title": "REST API",
    "description": "learning about REST API",
    "status": "undone",
    "due_date": "2021-05-26"
  }
  ```

  Response:

  ```json
  {
    "id": 1,
    "title": "REST API",
    "description": "learning about REST API",
    "status": "undone",
    "due_date": "2021-05-26T00:00:00.000Z",
    "createdAt": "2021-05-24T15:01:21.735Z",
    "updatedAt": "2021-05-24T15:01:21.735Z"
  }
  ```

- **Notes:** none

---

## Show All Todos

Show all todos in Go Todo

- **URL**

  `/todos`

- **Method:**

  `GET`

- **URL Params**: none

- **Data Params**: none

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**
    ```json
    [
      {
        "id": "<id number>",
        "title": "<todo title>",
        "description": "<todo description>",
        "status": "<todo status: done/undone>",
        "due_date": "<todo due date>",
        "createdAt": "2021-05-24T15:01:21.735Z",
        "updatedAt": "2021-05-24T15:01:21.735Z"
      },
      {
        "id": "<id number>",
        "title": "<todo title>",
        "description": "<todo description>",
        "status": "<todo status: done/undone>",
        "due_date": "<todo due date>",
        "createdAt": "2021-05-24T15:01:21.735Z",
        "updatedAt": "2021-05-24T15:01:21.735Z"
      }
    ]
    ```

- **Error Response:**

  - **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `"message": "Internal server error`

- **Notes:** none

---

## Show Todo by ID

Show todo in Go Todo by ID

- **URL**

  `/todos/:id`

- **Method:**

  `GET`

- **URL Params**

  `/:id`

  **Required:**

  `id=[integer]`

- **Data Params**: none

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**
    ```json
    {
      "id": "<id number>",
      "title": "<todo title>",
      "description": "<todo description>",
      "status": "<todo status: done/undone>",
      "due_date": "<todo due date>",
      "createdAt": "2021-05-24T15:01:21.735Z",
      "updatedAt": "2021-05-24T15:01:21.735Z"
    }
    ```

- **Error Response:**

  - **Code:** 404 NOT FOUND <br />
    **Content:** `"message": "Todo not found"`

- **Notes:** none

---

## Update Todo

Update all todo field in Go Todo

- **URL**

  `/todos/:id`

- **Method:**

  `PUT`

- **URL Params**

  `/:id`

  **Required:**

  `id=[integer]`

- **Data Params**: none

- **Request Body**

  ````json
    {
      "title": "<todo title>",
      "description": "<todo description>",
      "status": "<todo status: done/undone>",
      "due_date": "<todo due date, format: YYYY-MM-DD>"
    }
    ```

  ````

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**
    ```json
    {
      "id": "<id number>",
      "title": "<todo title>",
      "description": "<todo description>",
      "status": "<todo status: done/undone>",
      "due_date": "<todo due date>",
      "createdAt": "2021-05-24T15:01:21.735Z",
      "updatedAt": "2021-05-24T15:01:21.735Z"
    }
    ```

- **Error Response:**

  - **Code:** 400 BAD REQUEST <br />
    **Content:** `"message": "Validation error: Due date cannot be the day before today"`

    OR

  - **Code:** 404 NOT FOUND <br />
    **Content:** `"message": "Todo not found"`

    OR

  - **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `"message": "Internal server error"`

- **Notes:** none

---

## Update Todo Status

Update only todo status field in Go Todo

- **URL**

  `/todos/:id`

- **Method:**

  `PATCH`

- **URL Params**

  `/:id`

  **Required:**

  `id=[integer]`

- **Data Params**: none

- **Request Body**

  ````json
    {
      "status": "<todo status: done/undone>"
    }
    ```

  ````

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**
    ```json
    {
      "id": "<id number>",
      "title": "<todo title>",
      "description": "<todo description>",
      "status": "<todo status: done/undone>",
      "due_date": "<todo due date>",
      "createdAt": "2021-05-24T15:01:21.735Z",
      "updatedAt": "2021-05-24T15:01:21.735Z"
    }
    ```

- **Error Response:**

  - **Code:** 400 BAD REQUEST <br />
    **Content:** `"message": "Validation error: Due date cannot be the day before today"`

    OR

  - **Code:** 404 NOT FOUND <br />
    **Content:** `"message": "Todo not found"`

    OR

  - **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `"message": "Internal server error"`

- **Notes:** none

---

## Delete Todo

Delete todo from Go Todo

- **URL**

  `/todos/:id`

- **Method:**

  `DELETE`

- **URL Params**

  `/:id`

  **Required:**

  `id=[integer]`

- **Data Params**: none

- **Success Response:**

  - **Code:** 200 <br />
    **Content:**
    ```json
    {
      "message": "todo deleted",
      "deletedData": {
        "id": "<id number>",
        "title": "<todo title>",
        "description": "<todo description>",
        "status": "<todo status: done/undone>",
        "due_date": "<todo due date>",
        "createdAt": "2021-05-24T15:01:21.735Z",
        "updatedAt": "2021-05-24T15:01:21.735Z"
      }
    }
    ```

- **Error Response:**

- **Code:** 404 NOT FOUND <br />
  **Content:** `"message": "Todo not found"`

  OR

- **Code:** 500 INTERNAL SERVER ERROR <br />
  **Content:** `"message": "Internal server error"`

- **Notes:** none
