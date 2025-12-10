package entity

import (
	"context"
	"sync"

	"github.com/Higor-ViniciusDev/stress_cli/internal/service"
)

type TesteRequest struct {
	Url       string
	Requests  int
	Conc      int
	InputChan chan any
}

type FilaRequest struct {
	Cont int
	mu   sync.Mutex
}

func NovoTeste(url string, intRequests int, intConc int) *TesteRequest {
	teste := &TesteRequest{
		Url:       url,
		Requests:  intRequests,
		Conc:      intConc,
		InputChan: make(chan any, intRequests),
	}

	return teste
}

func (t *TesteRequest) Executar() chan any {
	novoFilaRequest := &FilaRequest{
		Cont: t.Requests,
	}

	for {
		if novoFilaRequest.Cont == 0 {
			break
		}

		timeStarted := make(chan string)

		for range t.Conc {
			go t.Processar(t.Url, timeStarted, novoFilaRequest)
		}

		close(timeStarted)
	}

	return t.InputChan
}

func (t *TesteRequest) Processar(url string, timeStarted chan (string), fila *FilaRequest) {
	novoService := &service.ServiceRequest{}
	ctx := context.Background()

	if fila.Cont == 0 {
		return
	}

	fila.mu.Lock()
	fila.Cont--
	fila.mu.Unlock()

	<-timeStarted
	retorno := novoService.FazerRequest(ctx, url)

	t.InputChan <- retorno
}
