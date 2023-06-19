package request

type CreateProductRequest struct {
	Name                string                         `json:"name" validate:"required,max=100"`
	UnitMassAcronym     string                         `json:"unit_mass_acronym" validate:"required,max=20"`
	UnitMassDescription string                         `json:"unit_mass_description" validate:"required,max=50"`
	ProductQualities    []*CreateProductQualityRequest `json:"product_qualities" validate:"required"`
}

type UpdateProductRequest struct {
	ID                  int64
	Name                string                         `json:"name" validate:"required,max=100"`
	UnitMassAcronym     string                         `json:"unit_mass_acronym" validate:"required,max=20"`
	UnitMassDescription string                         `json:"unit_mass_description" validate:"required,max=50"`
	ProductQualities    []*UpdateProductQualityRequest `json:"product_qualities" validate:"required"`
}
