package internal

var (
	Version = "v1.0.0"
	Build   = "0001"

	/*
			--url: URL do serviço a ser testado.
		--requests: Número total de requests.
		--concurrency: Número de chamadas simultâneas.

	*/
	//Flags via CLI
	FlagURL  *string
	FlagReq  *int
	FlagConc *int
)
