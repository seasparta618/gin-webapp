# JSON Processor Web API Server
This is a web API server designed to work with the Go app json-processor. It provides a simple interface for processing JSON data and interacting with the JSON Processor application.

## Features
- Authentication: Provides an endpoint for user authentication and token generation.
- Token Refresh: Supports refreshing expired tokens using a refresh token.
- Enquiry Endpoint: Offers an endpoint for saving enquiries, which can be used to test the JSON processing capabilities of the json-processor app.
## Endpoints
- POST /auth/login: Authenticate with `username:"admin"` and `password:"password"` to receive a JWT token and a refresh token. In order to test, you can add another field `expired:true` in the request json body to get an expired token in order to see the 401 result from json-processor
- POST /auth/refresh: Use a refresh token to obtain a new JWT token and a new refresh token.
- POST /api/enquiry/save: Accepts JSON data for processing. Requires a valid JWT token for authorization.
## Usage
- Start the Server: Run the server using the command go run main.go. The server will start listening for requests.

- Authenticate: Send a POST request to /auth/login with a JSON body containing username and password fields to receive authentication tokens.

- Save Enquiry: Send a POST request to /api/enquiry/save with a JSON body containing enquiry data. Include the JWT token in the Authorization header for authentication.

- Refresh Token: If the JWT token expires, send a POST request to /auth/refresh with a JSON body containing the refresh_token field to obtain new tokens.

## Configuration
The server relies on environment variables for configuration. Ensure the following variables are set in your .env file:

1. API_HOST: The host address for the API server.
2. SIGNATURE: The secret signature for generating JWT tokens.

## Dependencies
This server is built using the Gin web framework and uses the jwt-go library for JWT token handling.