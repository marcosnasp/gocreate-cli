package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Uso: gocreate <nome-do-projeto>\n")
		fmt.Fprintf(os.Stderr, "Cria um novo projeto Go com uma estrutura de diretórios padrão.\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	projectName := flag.Args()[0]

	err := createProject(projectName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar o projeto: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Projeto %s criado com sucesso!\n", projectName)
}

func createProject(projectName string) error {
	// Lista de diretórios a serem criados
	dirs := []string{
		filepath.Join(projectName, "cmd", projectName),
		filepath.Join(projectName, "internal", "app"),
		filepath.Join(projectName, "pkg", "utils"),
	}

	// Criar os diretórios
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("falha ao criar diretório %s: %v", dir, err)
		}
	}

	// Mapa de arquivos com seus conteúdos iniciais
	files := map[string]string{
		filepath.Join(projectName, "cmd", projectName, "main.go"): `package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}`,
		filepath.Join(projectName, "internal", "app", "app.go"): `package app

import "fmt"

func SayHello() {
	fmt.Println("Hello from internal/app!")
}`,
		filepath.Join(projectName, "pkg", "utils", "utils.go"): `package utils

import "fmt"

func PrintMessage(msg string) {
	fmt.Println(msg)
}`,
		filepath.Join(projectName, ".gitignore"): `# Binários
*.exe
*.exe~
*.dll
*.so
*.dylib

# Arquivos de build
*.o
*.obj

# Diretórios de saída
/bin/
/pkg/
/dist/`,

		filepath.Join(projectName, "README.md"): `# ` + projectName + `

Um projeto Go com estrutura organizada

## 📦 Estrutura do Projeto

` + "```" + `
/
├── cmd/
│   └── ` + projectName + `/
│       └── main.go
├── internal/
│   └── app/
├── pkg/
│   └── utils/
├── .gitignore
└── README.md
` + "```" + `

## ⚙️ Instalação
` + "```" + `bash
go mod download
` + "```" + `

## 🚀 Execução
` + "```" + `bash
go run cmd/` + projectName + `/main.go
` + "```" + ``,
	}

	// Criar os arquivos
	for filePath, content := range files {
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			return fmt.Errorf("falha ao criar arquivo %s: %v", filePath, err)
		}
	}

	// Inicializar o módulo Go
	err := os.Chdir(projectName)
	if err != nil {
		return fmt.Errorf("falha ao entrar no diretório %s: %v", projectName, err)
	}

	err = initGoModule(projectName)
	if err != nil {
		return fmt.Errorf("falha ao inicializar o módulo Go: %v", err)
	}

	return nil
}

func initGoModule(moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
