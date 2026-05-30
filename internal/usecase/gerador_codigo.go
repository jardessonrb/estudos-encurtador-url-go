package usecase

import (
	"encurtador-url-go/internal/repository"
	"math/rand"
)

type GeradorCodigoService struct {
	Charset  string
	Contador repository.Contador
}

func NewGeradorCodigoService(contador repository.Contador) *GeradorCodigoService {
	return &GeradorCodigoService{
		Charset:  "w8oVxhifqWcFQRtXvS209alLOM67KDCjpAb14NzUrPmg5ndkuGY3HyJZBsIeET",
		Contador: contador,
	}
}

func (service *GeradorCodigoService) CalcularValorBase62(valor int) string {
	if valor < 62 {
		return string(service.Charset[valor%len(service.Charset)])
	}

	var hash string = ""
	var restante = valor

	for restante >= 62 {
		resto := restante % 62

		hash += string(service.Charset[resto%len(service.Charset)])
		// fmt.Printf("resto: %d - %s\n", resto, hash)
		restante = restante / 62

		if restante < 62 {
			hash += string(service.Charset[restante])
			// fmt.Printf("restante: %d - %s\n", restante, hash)
			return hash
		}
	}

	return ""
}

func (service *GeradorCodigoService) ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (service *GeradorCodigoService) GerarNumeroAleatorio() int {
	return rand.Intn(1000000000-100000) + 100000
}

func (service *GeradorCodigoService) GerarCodigo() string {
	// numero := service.GerarNumeroAleatorio()
	numero, err := service.Contador.Next()

	if err != nil {
		panic(err)
	}

	codigo := service.CalcularValorBase62(int(numero))

	return service.ReverseString(codigo)
}
