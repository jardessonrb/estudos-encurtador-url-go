package usecase

import "testing"

var service = NewGeradorCodigoService(nil)

func TestReverseString(t *testing.T) {
	tests := []struct {
		entrada  string
		esperado string
	}{
		{"abc", "cba"},
		{"123abc", "cba321"},
	}

	for _, test := range tests {
		resultado := service.ReverseString(test.entrada)

		if resultado != test.esperado {
			t.Errorf("Esperado %s, obtido %s", test.esperado, resultado)
		}
	}
}

func TestCalcularValorBase62(t *testing.T) {
	tests := []struct {
		entrada  int
		esperado string
	}{
		{61, "T"},
		{62, "w8"},
		{125, "8o"},
	}

	for _, test := range tests {
		resultado := service.CalcularValorBase62(test.entrada)

		if resultado != test.esperado {
			t.Errorf("Esperado %s, obtido %s", test.esperado, resultado)
		}
	}
}

func TestNewGeradorCodigoService(t *testing.T) {
	service = NewGeradorCodigoService(nil)

	if service == nil {
		t.Fatal("Esperado service, obtido nil")
	}
}

func TestNewGeradorCodigoServiceTamanhoCharset(t *testing.T) {
	service = NewGeradorCodigoService(nil)
	var tamanhoEsperado int = 62

	if len(service.Charset) != tamanhoEsperado {
		t.Errorf("Esperado %d, obtido %d", tamanhoEsperado, len(service.Charset))
	}
}
