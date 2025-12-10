package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/Higor-ViniciusDev/stress_cli/internal/entity"
)

type TesteUsecase struct {
}

func NewTesteUsecase() *TesteUsecase {
	return &TesteUsecase{}
}

type RelatorioDTO struct {
	TempoTotalGasto    time.Duration
	QuantityTotalGasto int
	ResponsesCode      map[int]int
	Code200Sucess      int
}

func (t *TesteUsecase) ExecutarTeste(url string, request, concurrency int) *RelatorioDTO {
	novoEntity := entity.NovoTeste(url, request, concurrency)

	retornoResponse := novoEntity.Executar()
	totalRequests := len(retornoResponse)
	successCount := 0 // Code 200
	errorCount := 0
	statusCodes := make(map[int]int)

	for _, result := range retornoResponse {
		if result.Error != nil {
			errorCount++
			continue
		}

		statusCodes[result.Code]++
		if result.Code == 200 {
			successCount++
		}
	}

	return &RelatorioDTO{
		TempoTotalGasto:    novoEntity.TimeFinal.Sub(novoEntity.TimeInicial),
		QuantityTotalGasto: totalRequests,
		ResponsesCode:      statusCodes,
		Code200Sucess:      successCount,
	}
}

func (t *TesteUsecase) PrintRelatorioStdout(rel *RelatorioDTO) {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("RESULTADO DO TESTE")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Printf("Tempo total:          %s\n", t.formatDuration(rel.TempoTotalGasto))
	fmt.Printf("Total de requisições: %d\n", rel.QuantityTotalGasto)
	fmt.Printf("Sucessos (200):       %d\n", rel.Code200Sucess)
	fmt.Printf("Erros:                %d\n", rel.QuantityTotalGasto-rel.Code200Sucess)

	fmt.Println("\nCódigos HTTP recebidos:")
	for code, count := range rel.ResponsesCode {
		fmt.Printf("   • %d -> %d\n", code, count)
	}

	fmt.Println(strings.Repeat("=", 40))
}

func (t *TesteUsecase) formatDuration(tempo time.Duration) string {
	return fmt.Sprintf("%dms", tempo.Milliseconds())
}
