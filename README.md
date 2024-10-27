# mkauth

`mkauth` is an hobby  authentication server in Go, built to be compatible with OpenID Connect (OIDC) and OAuth 2.0 to begin with. Idea is to stick to RFC and implementing as much of the specifications as possible to serve as a robust auth server. This is a hobby project to creat my own Auth server. 

## Stack
- **Go** – for backend logic and API implementation
- **HTMX** – for building simple, dynamic frontend components

## Goals
- **Create a client**: Build a client to interact with and test the authentication server.
- **RFC-Compliant API**: Implement an end-to-end API that follows the OAuth 2.0 and OIDC specifications closely.
- **Basic Authentication and Consent Pages**: Use HTMX to create simple, interactive pages for login and user consent.

## Work In Progress (WIP)
- **Base Setup**:
  - Logging, validation, error handling, and routing
  - In-memory database setup
- **OIDC Auth Code Flow Support**:

Auth code flow 
- /authrozie endpoint
- validation
- create redirect to login page 
- create clients and secret in inmemDB 
- create user table
- prompt for login test
- generate and store the code and with proper expiry and checks
- Send back JWT. 
- 

    
## Future Ideas
- Add support for additional OAuth 2.0 and OIDC flows (e.g., Implicit, Hybrid).
- Expand client testing and documentation for easier integration.
- Implement consent management and additional user data storage options.

## Getting Started
1. **Clone the repository**:
   ```bash
   install https://github.com/air-verse/air
   git clone https://github.com/smukk9/mkauth.git
   cd mkauth
   air
   
