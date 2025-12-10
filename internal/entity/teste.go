package entity

import (
	"context"
	"sync"
	"time"

	"github.com/Higor-ViniciusDev/stress_cli/internal/service"
)

type TesteRequest struct {
	Url             string
	Requests        int
	Conc            int
	InputChan       chan MessageRequest
	TimeInicial     time.Time
	TimeFinal       time.Time
	mu              sync.Mutex
	ResponseResults []*Response
	wg              sync.WaitGroup
	ServiceRequest  *service.ServiceRequest
}

type MessageRequest struct {
	Reply chan ResponseRequest
}

type ResponseRequest struct {
	CodeResponse int
	Status       bool
	Error        error
}

type Response struct {
	Code  int
	Error error
}

func NovoTeste(url string, intRequests int, intConc int) *TesteRequest {
	teste := &TesteRequest{
		Url:            url,
		Requests:       intRequests,
		Conc:           intConc,
		InputChan:      make(chan MessageRequest, intRequests),
		ServiceRequest: service.NewServiceRequest(),
	}

	for range teste.Conc {
		teste.wg.Add(1)
		go teste.Processar()
	}

	return teste
}

func (t *TesteRequest) Executar() []*Response {
	t.TimeInicial = time.Now()
	reply := make(chan ResponseRequest, t.Requests)

	for range t.Requests {
		msg := MessageRequest{
			Reply: reply,
		}
		t.InputChan <- msg
	}
	// send to rate limiter
	for range t.Requests {
		retornoRequest := <-reply
		t.mu.Lock()
		t.ResponseResults = append(t.ResponseResults, &Response{
			Code:  retornoRequest.CodeResponse,
			Error: retornoRequest.Error,
		})
		t.mu.Unlock()
	}

	t.Stop()
	t.TimeFinal = time.Now()
	return t.ResponseResults
}

func (t *TesteRequest) Processar() {
	defer t.wg.Done() // Espera concluir o processamento da request

	for msg := range t.InputChan {
		ctx := context.Background()

		retorno := t.ServiceRequest.FazerRequest(ctx, t.Url)

		requestRetorno := ResponseRequest{
			CodeResponse: retorno.Code,
			Error:        retorno.Error,
		}

		msg.Reply <- requestRetorno
	}
}

func (t *TesteRequest) Stop() {
	close(t.InputChan)
	t.wg.Wait()
}
