# Onde Estão as Minhas Classes?

**Introdução:** O choque cultural (O 'preconceito' inicial com a simplicidade do Go)

**Capítulo 1:** O mistério da sintaxe (O `:=`, a ordem das variáveis e a tipagem que me deu dores de cabeça)

**Capítulo 2:** O deserto dos Loops (Como sobrevivi apenas com o `for`)

**Capítulo 3:** O momento 'Aha!' (Como as Goroutines e Channels mudaram a minha forma de pensar)

**Capítulo 4:** Fazer as pazes com os Erros (Por que o `if err != nil` é, na verdade, meu amigo)

**Conclusão:** O que eu diria ao 'eu' do passado antes de ter começado

---

## 📂 Estrutura do Projeto

```
cafe-com-go/
├── README.md                        # Artigo/saga principal
├── cheatsheet.go                    # Código consolidado para consulta rápida
├── go.mod                           # Arquivo de módulo Go
├── .gitignore                       # Arquivos ignorados pelo Git
└── capitulos/                       # Exemplos práticos por capítulo
    ├── 01-sintaxe/                  # Capítulo 1: O mistério da sintaxe
    │   └── sintaxe.go
    ├── 02-loops/                    # Capítulo 2: O deserto dos Loops
    │   └── loops.go
    ├── 03-concorrencia/             # Capítulo 3: O momento 'Aha!'
    │   └── concorrencia.go
    └── 04-erros/                    # Capítulo 4: Fazer as pazes com os Erros
        └── erros.go
```

## 🚀 Como Usar

```bash

# Exemplos por capítulo
go run capitulos/01-sintaxe/sintaxe.go
go run capitulos/02-loops/loops.go
go run capitulos/03-concorrencia/concorrencia.go
go run capitulos/04-erros/erros.go
```
