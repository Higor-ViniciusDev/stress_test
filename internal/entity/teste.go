package entity

type TesteRequest struct {
	Url       string
	Requests  int
	Conc      int
	InputChan chan string
}

type FilaRequest struct {
	Request chan int
	Cont    int
	Url     string
}

func NovoTeste(url string, intRequests int, intConc int) *TesteRequest {
	teste := &TesteRequest{
		Url:      url,
		Requests: intRequests,
		Conc:     intConc,
	}

	return teste
}

func (t *TesteRequest) Executar() {
	novoFilaRequest := &FilaRequest{
		// Request: make(int, t.re)
	}

	for range t.Requests {
		t.InputChan <- t.Url
	}
}

func (t *TesteRequest) Processar(url string, timeStarted chan (string)) {

	<-timeStarted
}
