package usecase

import "github.com/Higor-ViniciusDev/stress_cli/internal/entity"

type TesteUsecase struct {
}

func NewTesteUsecase() *TesteUsecase {
	return &TesteUsecase{}
}

func (t *TesteUsecase) ExecutarTeste(url string, request, concurrency int) {
	// Lógica para executar o teste de estresse
	/*
		Realizar requests HTTP para a URL especificada.
		Distribuir os requests de acordo com o nível de concorrência definido.
		Garantir que o número total de requests seja cumprido.
		Geração de Relatório:
	*/

	novoEntity := entity.NovoTeste(url, request, concurrency)

	for range novoEntity.Executar() {
		// Aqui você pode processar os resultados de cada request, se necessário
	}
}
