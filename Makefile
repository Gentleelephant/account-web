.PHONY: build check help clean doc

help:
	@echo "use [make check]: code review"
	@echo "use [make build]: build project"
build:
	@go build -o account-web main.go
check:
	@golangci-lint run
clean:
	@rm -f ./account-web
	@rm -rf ./cache
	@rm -rf ./logs
doc:
	@swag init --parseDependency --parseInternal

update:
	@git pull