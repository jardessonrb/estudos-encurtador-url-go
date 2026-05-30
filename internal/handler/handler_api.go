package handler

import (
	"encurtador-url-go/internal/domain"
	"encurtador-url-go/internal/dto"
	"encurtador-url-go/internal/repository/postgres"
	"encurtador-url-go/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

var geradorCodigo *usecase.GeradorCodigoService = usecase.NewGeradorCodigoService()
var repository *postgres.UrlRepository = postgres.NewUrlRepository()

func EncurtarUrlHandler(contexto *gin.Context) {
	var request dto.UrlEncurtadaForm

	if err := contexto.ShouldBindJSON(&request); err != nil {
		contexto.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "Erro no json enviado",
		})

		return
	}

	url := &domain.URL{
		OriginalUrl: request.URL,
		Codigo:      geradorCodigo.GerarCodigo(),
	}

	repository.Salvar(url)

	var response = dto.UrlEncurtadaCriadaResponse{
		ID:     url.ID,
		Codigo: url.Codigo,
	}

	contexto.JSON(http.StatusCreated, response)
}

func RetornarURLPorCodigoHandler(contexto *gin.Context) {
	codigo := contexto.Param("codigo")

	url, err := repository.BuscarPorCodigo(codigo)

	if err != nil {
		contexto.JSON(http.StatusNotFound, gin.H{
			"mensagem": "URL não encontrada para o código informado",
		})
	}

	response := &dto.UrlEncurtadaResponse{
		URL: url.OriginalUrl,
	}

	contexto.JSON(http.StatusPermanentRedirect, response)
}

func CheckHealth(contexto *gin.Context) {
	contexto.JSON(200, gin.H{
		"mensagem": "API On",
	})
}
