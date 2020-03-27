package servicestest

import (
	"testing"

	"github.com/glaubergoncalves/analise/api/models"
	"github.com/glaubergoncalves/analise/api/services"
)

func TestAnaliseFraudeService(t *testing.T) {
	resultadoAnalise := services.AnaliseFraudeService(models.Venda{})

	if resultadoAnalise.ResultadAnalise != true {
		t.Errorf("era para retornar true na analise")
	}
}
