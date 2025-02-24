# Makefile para o projeto gocreate

# Variáveis
BINARY_NAME = gocreate
BUILD_DIR = build
VERSION = v0.1.0

# Plataformas alvo
PLATFORMS = linux windows darwin
ARCH = amd64

# Comando padrão
.PHONY: all
all: build

# Compilar para todas as plataformas
.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		if [ "$$platform" = "windows" ]; then \
			ext=".exe"; \
		else \
			ext=""; \
		fi; \
		GOOS=$$platform GOARCH=$(ARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME)-$$platform-$(ARCH)$$ext ./cmd/gocreate-cli/main.go; \
		echo "Gerado: $(BUILD_DIR)/$(BINARY_NAME)-$$platform-$(ARCH)$$ext"; \
	done

# Limpar arquivos gerados
.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)
	@echo "Diretório $(BUILD_DIR) removido."

# Compactar os binários para release
.PHONY: package
package: build
	@for platform in $(PLATFORMS); do \
		if [ "$$platform" = "windows" ]; then \
			cd $(BUILD_DIR) && tar -a -c -f $(BINARY_NAME)-$(VERSION)-$$platform-$(ARCH).zip $(BINARY_NAME)-$$platform-$(ARCH).exe && cd ..; \
			echo "Gerado: $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-$$platform-$(ARCH).zip"; \
		else \
			tar -czvf $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-$$platform-$(ARCH).tar.gz -C "$(BUILD_DIR)" "$(BINARY_NAME)-$$platform-$(ARCH)"; \
			echo "Gerado: $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-$$platform-$(ARCH).tar.gz"; \
		fi; \
	done

# Testar o binário localmente
.PHONY: test
test: build
	@$(BUILD_DIR)/$(BINARY_NAME)-$(shell go env GOOS)-$(ARCH) meu-projeto
	@echo "Teste concluído. Verifique o diretório 'meu-projeto'."