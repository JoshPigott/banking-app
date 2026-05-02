# Banking app
- The project recreates a banking app like ASB

## Stack 
- go 
- SSR html with htmx

## Feature 
- Login & Sign up & logout
- Different types of bank accounts: everyday, saver, kiwisaver
- transfers between accounts
- payments 

## Screenshots 
login page
![login page screenshot desktop](/screenshots/login-page-desktop.png)
![login page screenshot phone](/screenshots/login-page-phone.png)

online banking page
![online banking page screenshot desktop](/screenshots/online-banking-page-desktop.png)
![online banking page screenshot phone](/screenshots/online-banking-page-phone.png)

transfer page
![transfer page screenshot desktop](/screenshots/transfer-page-desktop.png)
![transfer page screenshot phone](/screenshots/transfer-page-phone.png)

payments page
![payments page screenshot desktop](/screenshots/payments-page-desktop.png)
![payments page screenshot phone](/screenshots/payments-page-phone.png)


## File structure
```text
│   .air.toml
│   .gitignore
│   docker-compose.yml
│   Dockerfile
│   go.mod
│   go.sum
│   improvements.md
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
│       .gitkeep
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
│   │       html-pages.go
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
├───web
│   ├───assets
│   │       default-account.jpg
│   │       dollar-sign-icon.png
│   │       dropdown.js
│   │       everyday-account.jpg
│   │       favicon.png
│   │       kiwisaver-account.jpg
│   │       password-icon.png
│   │       saver-account.jpg
│   │       style.css
│   │       user-icon.png
│   │
│   ├───templates
│   │       account.html
│   │       balance.html
│   │
│   └───views
│           login.html
│           online-banking.html
│           payment.html
│           transfer.html
│
├───screenshots
│       login-page-desktop.png
│       login-page-phone.png
│       online-banking-page-desktop.png
│       online-banking-page-phone.png
│       payments-page-desktop.png
│       payments-page-phone.png
│       transfer-page-desktop.png
│       transfer-page-phone.png
│
└───tmp
        main
```

## Key flow and logic

### Auth (signup + login)
- `handlers.SignUp` → validate → `helpers.CreateUserAccount`
- `CreateUserAccount` → hash password → `database.CreateUserAccount`
- Then `createAllAccounts` → creates 3 account rows
- Session created via `helpers.CreateSession`
- Cookie set in `setSessionCookie`

- `handlers.LoginAuth` → `helpers.ValidLoginCredentials`
- If valid → new session → cookie

---

### Auth protection
- `middleware.RequireAuth`
  - Reads `session_id` cookie
  - Calls `database.GetSession`
  - Checks expiry
  - Rejects or continues

---

### Account data (HTMX)
- Page loads → HTMX calls:
  - `/get-welcome-message` → `handlers.GetWelcomeMessage`
  - `/account-balance` → `handlers.GetAccount`

- Flow:
  - sessionID → `database.GetUserID`
  - account type → `BankAccountType.GetTableName`
  - balance → `database.GetAccountBalance`

---

### Transfers
- `handlers.TransferMoney`
  - parse → `getTransferData`
  - validate → `helpers.CanTransfer`
  - execute → `helpers.MakeTransfer`

- Core flow:
  - session → userID (`GetUserID`)
  - map enums → table names
  - call `database.MakeTransfer`

- DB layer:
  - wrapped in `withTx`
  - `transfer`:
    - `WithDraw`
    - `Deposit`

---

### Payments (to another user)
- `handlers.Payment`
  - parse → `getPaymentData`
  - validate → `helpers.IsValidPayment`
  - execute → `helpers.MakePayment`

- Key differences:
  - `receiver()` → gets receiver userID
  - prevents sending to self
  - deposits into default account

---

### Transactions (important)
- `withTx`
  - begin tx
  - run logic
  - rollback on error
  - commit if success

- Ensures:
  - no partial updates
  - withdraw + deposit always linked

---

### Validation rules
- `isValidAmount`
  - amount > 0
  - enough balance

- `BankAccountType`
  - `IsValid`
  - `CanWithdraw`

- Password:
  - `Password.IsStrong`

---

### Database design
- Tables:
  - `users`
  - `sessions`
  - `everydayAccount`, `saverAccount`, `kiwiSaverAccount`

- Each account:
  - `userID`
  - `balanceCents`

---


## Problems issues know constraints
A lot of the current code lacks the go idiomatic style
The CSS on a latop screen size looks bad (smaller and bigger ok)
The transfer and payment are not beign clear after transaction