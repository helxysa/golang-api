package controller

import (
	model "go-api-catalog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type solucaoController struct {
	db *gorm.DB
}

func NewSolucaoController(db *gorm.DB) *solucaoController {
	return &solucaoController{db: db}
}

func (p *solucaoController) GetSolucoes(ctx *gin.Context) {
	var solucoes []model.Solucao

	result := p.db.Find(&solucoes)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, solucoes)
}

func (p *solucaoController) CreateSolucao(ctx *gin.Context) {
	var solucao model.Solucao
	if err := ctx.ShouldBindJSON(&solucao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON INVALIDO" + err.Error()})
		return
	}

	if err := p.db.Create(&solucao).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, solucao)
}


func (p *solucaoController) UpdateSolucao(ctx *gin.Context) {	
	id := ctx.Param("id")
	var solucao model.Solucao
	if err := p.db.First(&solucao, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Solucao não encontrada"})
		return	
	}


	var atualizacao map[string]interface{}
	if err := ctx.ShouldBindJSON(&atualizacao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON INVALIDO:" + err.Error()})
		return
	}

	if err := p.db.Model(&solucao).Updates(atualizacao).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, solucao)
}


func (p *solucaoController) DeleteSolucao(ctx *gin.Context) {
	id := ctx.Param("id")
	var solucao model.Solucao
	if err := p.db.First(&solucao, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Solucao não encontrada"})
		return
	}

	if err := p.db.Model(&solucao).Delete(&solucao).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Erro ao deletar solucao"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Solucao deletada com sucesso"})


}






