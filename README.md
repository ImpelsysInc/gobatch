# gobatch
GoBatch is a Golang/Go  based batch processing framework like Spring Batch in Java.

## Architecture

## Folder Structure
```sh
├── app
│   ├── app.go  # App Intitialization
│   ├── database
│   │   ├── db
│   │   │   └── db.go
│   │   └── gorm
│   │       └── gorm.go
│   ├── dbs.go # Database Intitialization
│   └── locale
│       └── locale.go
├── cmd  # Starting point for any application
│   ├── api
│   │   └── main.go # Main application start
│   └── migrate
│       └── main.go # Migration start
├── config
│   └── config.go # App configuration. Any environment specific will get from .env
├── go.mod
├── go.sum
├── locale  # Multi language support for the API
│   ├── el-GR
│   │   └── example.yml
│   ├── en-US
│   │   └── example.yml
│   └── zh-CN
│       └── example.yml
├── log  # Log folder
│   └── micro.log
├── Makefile  # Common command 
├── migration # Migration files
│   ├── 20190805170000_create_user_table.sql
│   └── 20210130204915_create_tenant_table.sql
├── module  # Module contain actual business logic
│   ├── module.go
│   └── tenant  # Tenant Module
│   │   ├── handler.go
│   │   ├── handler_test.go
│   │   ├── inject.go
│   │   ├── mocks
│   │   │   ├── ITenantRepo.go
│   │   │   └── ITenantService.go
│   │   ├── model
│   │   │   └── tenant.go
│   │   ├── repo
│   │   │   └── tenant.go
│   │   ├── routes.go
│   │   └── service
│   │       ├── tenant_service.go
│   │       └── tenant_service_test.go
│   └── user # Tenant Module
│       ├── handler.go
│       ├── handler_test.go
│       ├── inject.go
│       ├── mocks
│       │   ├── IUserRepo.go
│       │   └── IUserService.go
│       ├── model
│       │   └── user.go
│       ├── repo
│       │   └── user.go
│       ├── routes.go
│       └── service
│           ├── user_service.go
│           └── user_service_test.go
└── README.md