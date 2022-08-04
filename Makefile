.PHONY: build check help

help:
	@echo "use [make check]: code review"
	@echo "use [make build]: build project"
build:
	@go build -o account-web.exe main.go

check:
	@golangci-lint run

clean:
	@rm -f ./account-web
	@rm -rf ./cache
	@rm -rf ./logs