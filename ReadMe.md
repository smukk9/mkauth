# MkAuth

A experimental and educational SSO Server. Think of KeyCloak or PingFederate but smaller scale.

## Goal of the project
- Use AI as project manager, planner and code reviewer
- No code generation - I will still write and make the decision
- Learn along the way
- TDD approach (Test-Driven Development)

## Implementation Progress
All this is planned by Claude as per my goal and we keep changing based on pattern or things i find as we go along, First step is to get a
Client credentials grant type.


| # | Task | Test | Impl | Status | Date | Notes |  
|---|------|------|------|--------|------|-------|
| 1 | Project Setup | ✅ | ✅ | ✅ | | go mod init, dependencies installed |
| 2A | Config with Viper | ✅ | ✅ | ✅ | | Using MKAUTH_FILE env + YAML |
| 2B | Expand Config for OAuth | ⬜ | ⬜ | ⏭️ | | DEFERRED - Add as needed |
| 2C | Health Handler | ✅ | ✅ | ✅ | | Handler pattern established |
| 2D | Wire in Main | ✅ | ✅ | ✅ | | Router + dependencies working |
| 3 | PocketBase Initialization | ⬜ | ⬜ | 🔄 | | NEXT: Database wrapper |
| 4 | Database Schema Migration | ⬜ | ⬜ | ⬜ | | oauth_clients, oauth_tokens |
| 4B | Seed Test OAuth Client | ⬜ | ⬜ | ⬜ | | test-client-123 with bcrypt secret |
| 5 | Test Infrastructure Setup | ⬜ | ⬜ | ⬜ | | Test helpers and utilities |
| 6A | OAuthClient Model | ⬜ | ⬜ | ⬜ | | Client struct definition |
| 6B | Client.SupportsGrantType | ⬜ | ⬜ | ⬜ | | Grant type validation |
| 7A | OAuthToken Model | ⬜ | ⬜ | ⬜ | | Token struct definition |
| 7B | Token.IsExpired | ⬜ | ⬜ | ⬜ | | Time-based expiration |
| 7C | Token.IsValid | ⬜ | ⬜ | ⬜ | | Complete validity check |
| 8A | CryptoService.HashPassword | ⬜ | ⬜ | ⬜ | | Bcrypt password hashing |
| 8B | CryptoService.VerifyPassword | ⬜ | ⬜ | ⬜ | | Bcrypt verification |
| 8C | CryptoService.GenerateRandomToken | ⬜ | ⬜ | ⬜ | | Secure token generation |
| 9A | JWTService Struct + Constructor | ⬜ | ⬜ | ⬜ | | JWT service initialization |
| 9B | JWTService.GenerateToken | ⬜ | ⬜ | ⬜ | | Create signed JWT |
| 9C | JWTService.ValidateToken | ⬜ | ⬜ | ⬜ | | Verify JWT signature |
| 10A | OAuth Error Responses | ⬜ | ⬜ | ⬜ | | RFC 6749 error format |
| 11A | Logger Setup | ⬜ | ⬜ | ⬜ | | Structured logging |
| 12A | ClientService.ValidateClient | ⬜ | ⬜ | ⬜ | | Authenticate client credentials |
| 12B | ClientService.CreateClient | ⬜ | ⬜ | ⬜ | | Register new OAuth client |
| 12C | ClientService.GetClientByID | ⬜ | ⬜ | ⬜ | | Retrieve client details |
| 12D | ClientService.ListClients | ⬜ | ⬜ | ⬜ | | Get all clients |
| 12E | ClientService.DeleteClient | ⬜ | ⬜ | ⬜ | | Remove client |
| 13A | TokenService - Issue JWT | ⬜ | ⬜ | ⬜ | | Generate JWT access token |
| 13B | TokenService - Issue Reference | ⬜ | ⬜ | ⬜ | | Generate opaque token |
| 13C | TokenService - Save to DB | ⬜ | ⬜ | ⬜ | | Persist token record |
| 13D | TokenService - Introspect Token | ⬜ | ⬜ | ⬜ | | Validate & return metadata |
| 13E | TokenService - Validate JWT Sig | ⬜ | ⬜ | ⬜ | | Verify JWT signature |
| 13F | TokenService - Revoke Token | ⬜ | ⬜ | ⬜ | | Invalidate token |
| 14A | Admin API - Create Client | ⬜ | ⬜ | ⬜ | | POST /admin/clients |
| 14B | Admin API - List Clients | ⬜ | ⬜ | ⬜ | | GET /admin/clients |
| 14C | Admin API - Get Client | ⬜ | ⬜ | ⬜ | | GET /admin/clients/:id |
| 14D | Admin API - Delete Client | ⬜ | ⬜ | ⬜ | | DELETE /admin/clients/:id |
| 15A | OAuth - Extract Client Creds | ⬜ | ⬜ | ⬜ | | Parse Authorization header |
| 15B | OAuth - Validate Client | ⬜ | ⬜ | ⬜ | | Authenticate client |
| 15C | OAuth - Issue Token | ⬜ | ⬜ | ⬜ | | Generate access token |
| 15D | OAuth - Return Token Response | ⬜ | ⬜ | ⬜ | | POST /oauth/token endpoint |
| 16A | Introspect - Parse Token | ⬜ | ⬜ | ⬜ | | Extract token from request |
| 16B | Introspect - Validate & Respond | ⬜ | ⬜ | ⬜ | | POST /oauth/introspect endpoint |
| 17A | Discovery - Metadata Response | ⬜ | ⬜ | ⬜ | | GET /.well-known/oauth-authorization-server |
| 18 | Wire All Components in Main | ⬜ | ⬜ | ⬜ | | Dependency injection |
| 19 | E2E Test - Full OAuth Flow | ⬜ | ⬜ | ⬜ | | Complete client credentials flow |
| 20 | Manual Testing Script | N/A | ⬜ | ⬜ | | curl commands for testing |

**Legend:**
- ✅ Complete
- 🔄 In Progress
- ⬜ Not Started
- ⏭️ Deferred

**Current Focus:** Task 3 - PocketBase Initialization & Wrapper
