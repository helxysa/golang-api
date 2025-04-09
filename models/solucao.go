package model

type Solucao struct {
	ID          int    `json:"id_solucoes" gorm:"column:id;primaryKey"`
	DemandaID   int    `json:"demanda_id" gorm:"column:demanda_id"`
	Nome        string `json:"nome" gorm:"column:nome"`
	Sigla       string `json:"sigla" gorm:"column:sigla"`
	Versao      string `json:"versao" gorm:"column:versao"`
	Link        string `json:"link" gorm:"column:link"`
	Andamento   string `json:"andamento" gorm:"column:andamento"`
	Repositorio string `json:"repositorio" gorm:"column:repositorio"`
	Criticidade string `json:"criticidade" gorm:"column:criticidade"`
	// tipo_id int
	// linguaguem_id string
	// desenvolvedor_id int
	// categoria_id int
	// responsavel_id int
	// status_id int
	// proprietario_id int
	DataStatus string `json:"data_status" gorm:"column:data_status"`
}

// TableName specifies the table name for the Solucao model
func (Solucao) TableName() string {
	return "solucoes"
}
