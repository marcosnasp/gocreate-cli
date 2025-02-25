# gocreate

Um CLI em Go para criar projetos Go com uma estrutura de diretórios padrão.

## Instalação

```bash
go install github.com/marcosnasp/gocreate-cli@latest
```

# Estrutura do Projeto após Execução do `gocreate library-api`

# library-api

Um projeto Go com estrutura organizada

## 📦 Estrutura do Projeto

```
/
├── cmd/
│   └── library-api/
│       └── main.go
├── internal/
│   └── app/
├── pkg/
│   └── utils/
├── .gitignore
└── README.md
```

## ⚙️ Instalação
```bash
go mod download
```

## 🚀 Execução
```bash
go run cmd/library-api/main.go
```