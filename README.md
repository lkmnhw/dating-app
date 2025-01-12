# Dating App - REST API
A backend REST API built with Go and MongoDB for a dating application. This API provides features for user signup, login, profile creation, and swiping on other profiles.

## Features
User authentication with JSON Web Tokens (JWT).
Profile creation and preference settings.
Feed of potential matches based on preferences.
Swipe functionality (like or pass) with user match tracking.
Secure handling of user data using MongoDB.
## Getting Started
This guideline to help run locally

### Prerequisites
- Go (version >= 1.23.4)
- MongoDB
- gpg

### Steps
- Clone the repository & install dependencies:
```bash
git clone https://github.com/your-username/dating-app.git
cd dating-app
go mod tidy
```
- Set .env
| Variable | Descriptio|
|----------|----------|
| PORT | Specifies the port on which the application will run. |
| DATABASE_SOURCE | MongoDB connection string. |
| DATABASE_NAME | MongoDB database name. |
| COLLECTION_USERS | The name of the MongoDB collection used to store user data. |
| COLLECTION_PROFILES |The name of the MongoDB collection used to store user profile data. |
| COLLECTION_MATCHES | The name of the MongoDB collection used to store match data between users. |
| JWT_SECRET_KEY | A secret key used for signing and verifying JSON Web Tokens (JWT) for user authentication. |

- or genererate .env (key: dating-app)
```bash
gpg .env.gpg
```
- Run the app
```bash
go run main.go
```
- Run test
```bash
go test -race -covermode=atomic -v ./...
```
## Available API
- ```GET /ping```
Check if the server is running.
```bash
curl --location 'http://localhost:3000/ping'
```
- ```POST /signup```
Register a new user.
```bash
curl --location 'http://localhost:3000/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "your_email@example.com",
  "password": "your_password"
}'
```
- ```POST /login```
Log in and obtain a token.
```bash
curl --location 'http://localhost:3000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "your_email@example.com",
  "password": "your_password"
}'
```
- ```POST /profile```
Create or update the user profile. Requires authentication.
```bash
curl --location 'http://localhost:3000/profile' \
--header 'Authorization: Bearer your_jwt_token' \
--header 'Content-Type: application/json' \
--data '{
  "name": "FirstName LastName",
  "description": "A brief description",
  "gender": "male",
  "date_of_birth": "1998-10-19",
  "preference": {
    "gender": "female",
    "minimum_age": 24,
    "maximum_age": 30
  }
}'

```
- ```GET /feed```
Retrieve a feed of profiles based on preferences. Requires authentication.
```bash
curl --location 'http://localhost:3000/feed' \
--header 'Authorization: Bearer your_jwt_token'
```
- or import postman collection from postman folder
