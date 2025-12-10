package usecase

import "github.com/Higor-ViniciusDev/stress_cli/internal/entity"

type TesteUsecase struct {
}

func NewTesteUsecase() *TesteUsecase {
	return &TesteUsecase{}
}

type RelatorioDTO struct {
}

func (t *TesteUsecase) ExecutarTeste(url string, request, concurrency int) *RelatorioDTO {
	novoEntity := entity.NovoTeste(url, request, concurrency)

	//Aqui processa retorno relatorio e preencher dto
	novoEntity.Executar()

	return &RelatorioDTO{}
}
