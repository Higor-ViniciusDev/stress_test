package main

import (
	"flag"
	"fmt"

	"github.com/Higor-ViniciusDev/stress_cli/internal"
	"github.com/Higor-ViniciusDev/stress_cli/internal/usecase"
)

func init() {
	internal.FlagURL = flag.String("url", "", "flag para informar a URL do serviço a ser testado")
	internal.FlagReq = flag.Int("requests", 0, "flag para informar o número total de requests")
	internal.FlagConc = flag.Int("concurrency", 0, "flag para informar o número de chamadas simultâneas")

	flag.Parse()
}

func main() {
	if *internal.FlagURL == "" {
		fmt.Println("URL invalida, por favor informe uma URL válida usando a flag -url")
		return
	}

	if *internal.FlagReq <= 0 {
		fmt.Println("Número de requests inválido, por favor informe um número maior que zero usando a flag -requests")
		return
	}

	if *internal.FlagConc <= 0 {
		fmt.Println("Número de concorrência inválido, por favor informe um número maior que zero usando a flag -concurrency")
		return
	}

	novoTeste := usecase.NewTesteUsecase()
	novoTeste.ExecutarTeste(*internal.FlagURL, *internal.FlagReq, *internal.FlagConc)
}
