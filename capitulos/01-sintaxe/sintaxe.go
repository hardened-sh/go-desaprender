package main // 1️. pacote

import ( // 2️. importações
	"fmt"
)

func main() {          // 3️. ponto de entrada
    // variáveis
    var nome, sobrenome string = "Luan", "Rodrigues"
    var idade int = 30

    // uso de :=
    contador := 0

    // loop
    for i := 0; i < 5; i++ {
        fmt.Printf("%d – %s %s (%d anos)\n",
            i+1, nome, sobrenome, idade)
        contador++
    }

    fmt.Println("Total de iterações:", contador)
}