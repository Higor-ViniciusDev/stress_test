package main

import (
	"flag"
	"fmt"

	"github.com/Higor-ViniciusDev/stress_cli/internal"
)

func init() {
	internal.FlagURL = flag.String("url", "", "flag para informar a URL do serviço a ser testado")
	internal.FlagReq = flag.Int("requests", 0, "flag para informar o número total de requests")
	internal.FlagConc = flag.Int("concurrency", 0, "flag para informar o número de chamadas simultâneas")

	flag.Parse()
}

func main() {
	fmt.Println("valor da flag URL:", *internal.FlagURL)
	fmt.Println("valor da flag requests:", *internal.FlagReq)
	fmt.Println("valor da flag concurrency:", *internal.FlagConc)
}
