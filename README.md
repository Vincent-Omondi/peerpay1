# Peer Pay

Peer Pay is a platform that links young investors through a peer-to-peer method.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/Vincent-Omondi/peerpay.git
cd peerpay
```

markdown

# Project Title

## Environment Variables

Create a `.env` file in the root directory of the project and add the following environment variables:

```env
DB_USER=root
DB_PASSWORD=examplepassword
DB_HOST=db
DB_PORT=3306
DB_NAME=peerpay

Build and Run the Application

Use Docker Compose to build and run the application:

sh

docker-compose up --build

This will build the Docker images and start the application on http://localhost:8080.
API Endpoints
Authentication
Sign Up

    Endpoint: POST /auth/signup
    Request Body:

    json

{
  "name": "string",
  "email": "string",
  "password": "string"
}

Response:

json

    {
      "message": "User created successfully"
    }

Login

    Endpoint: POST /auth/login
    Request Body:

    json

{
  "email": "string",
  "password": "string"
}

Response:

json

    {
      "token": "jwt_token"
    }

Referral

    Endpoint: POST /auth/referral
    Request Body:

    json

{
  "email": "string"
}

Response:

json

    {
      "message": "Referral sent successfully"
    }

Profile Management
Create Profile

    Endpoint: POST /profile/create
    Request Body:

    json

{
  "user_id": "integer",
  "bio": "string"
}

Response:

json

    {
      "message": "Profile created successfully"
    }

Edit Profile

    Endpoint: PUT /profile/edit
    Request Body:

    json

{
  "user_id": "integer",
  "bio": "string"
}

Response:

json

    {
      "message": "Profile updated successfully"
    }

Get Profile

    Endpoint: GET /profile/
    Response:

    json

    {
      "profile": {
        "user_id": "integer",
        "bio": "string"
      }
    }

Share Management
Buy Shares

    Endpoint: POST /shares/buy
    Request Body:

    json

{
  "user_id": "integer",
  "amount": "float",
  "days": "integer"
}

Response:

json

    {
      "message": "Shares bought successfully"
    }

Sell Shares

    Endpoint: POST /shares/sell
    Request Body:

    json

{
  "user_id": "integer",
  "amount": "float",
  "days": "integer"
}

Response:

json

    {
      "message": "Shares sold successfully"
    }

Get Shares

    Endpoint: GET /shares/
    Response:

    json

    {
      "shares": [
        {
          "user_id": "integer",
          "amount": "float",
          "days": "integer"
        }
      ]
    }

Admin Panel
Get Users

    Endpoint: GET /admin/users
    Response:

    json

    {
      "users": [
        {
          "id": "integer",
          "name": "string",
          "email": "string"
        }
      ]
    }

Delete User

    Endpoint: DELETE /admin/user/
    Response:

    json

    {
      "message": "User deleted successfully"
    }

Get Dormant Accounts

    Endpoint: GET /admin/dormant-accounts
    Response:

    json

    {
      "dormant_users": [
        {
          "id": "integer",
          "name": "string",
          "email": "string"
        }
      ]
    }

License

This project is licensed under the MIT License - see the LICENSE file for details.
Final Steps
Build and run the Docker containers:

sh

docker-compose up --build

Check the application:

The application should be running on http://localhost:8080.
