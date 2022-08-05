.PHONY: build check help clean doc update image

help:
	@echo "use [make check]: code review"
	@echo "use [make build]: build project"
	@echo "use [make clean]: clean project"
	@echo "use [make update]: git pull project"
	@echo "use [make doc]: generate doc"
	@echo "use [make image]: build docker image. do not use it"
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

image:
	@bash ./build.sh