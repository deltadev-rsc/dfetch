# Binary file name
BINARY = bin/fetch

# source code file
SRC = main.go 

# GoLang compile and linking flags
GO_FLAGS = -ldflags="-s -w"

build:
	@echo "Сборка кода фетча... "
	@go build -o $(BINARY) $(GO_FLAGS) $(SRC) 

test:
	@echo "Запуск тестов... "
	@go test ./...

clean:
	@echo "Очистка бинарей... "
	@rm -f $(BINARY) 

run: build 
	@./$(BINARY)

help:
	@echo "Доступные команды:"
	@echo "  make build   - Собрать бинарный файл"
	@echo "  make test    - Запустить тесты"
	@echo "  make clean   - Очистить артефакты"
	@echo "  make run     - Собрать и запустить"
	@echo "  make help    - Показать это сообщение"
