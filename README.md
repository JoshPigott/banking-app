# Banking app
- The project recreates a banking app like ASB

## Stack 
- go 
- SSR html with htmx

## Feature 
- Login & Sign up
- Different types of bank accounts: everyday, saver, kiwisaver

**To come**
- transfers between accounts
- payments 

## Screenshots (To come)

## File structure
```text
│   go.mod
│   go.sum
│   improvents.md
│   plan.md
│   README.md
│
├───.vscode
│       launch.json
│
├───cmd
│       main.go
│
├───data
│       database.db
│
├───internal
│   │
│   ├───database
│   │       bank-accounts.go
│   │       schema.go
│   │       session.go
│   │       users.go
│   │
│   ├───domain
│   │       bank-account.go
│   │       password.go
│   │       session.go
│   │       user.go
│   │
│   ├───handlers
│   │       auth.go
│   │       bank-accounts.go
│   │
│   ├───middleware
│   │       auth.go
│   │
│   ├───server
│   │       routes.go
│   │
│   ├───helpers
│   │       auth.go
│   │       bank-accounts.go
│   │       sessions.go
│   │       user-account.go
│   │
│   └───utils
│           utils.go
│
└───web
    └───templates
            balance.html
```

## Key logic flow (To come)

## Problems issues know constraints
A lot of the current code lacks the go idiomatic style
