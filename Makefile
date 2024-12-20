# Diretórios e variáveis
APP_NAME = myapp
CMD_DIR = ./cmd/$(APP_NAME)
BIN_DIR = ./bin

# Targets
.PHONY: all build run docker-build docker-run docker-compose clean

# Construir o binário Go
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

# Executar a aplicação Go localmente
run: build
	$(BIN_DIR)/$(APP_NAME)

# Construir imagem Docker
docker-build:
	docker build -t $(APP_NAME):latest .

# Rodar container Docker
docker-run: docker-build
	docker run --rm -p 8080:8080 $(APP_NAME):latest

# Usar Docker Compose para orquestração
docker-compose:
	docker-compose up -d
# Limpar binários e imagens temporárias
clean:
	rm -rf $(BIN_DIR)/*
	docker-compose down --rmi all --volumes --remove-orphans
