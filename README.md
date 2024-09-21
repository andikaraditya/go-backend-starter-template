# Go Backend Template

This templates includes:
- Go Fiber 
- pgx
- JWT authentication
- Validator
- Inital user migration

Requirement:
- Go 
- Golang Migrate
- Postgresql
- Make

How to use this template
1. Clone this template
2. Create a new repository
3. Run `Go mod init "your_module_name"`
4. Copy all these files and directory into your module root folder
5. Add `.env` and add the value based on the `.env.example`
6. Run `make update-imports module="your_module_name"`
7. Run `make run-migration`
   1. Optional: update the initial migration to your needs
8. Run `make run-api` 