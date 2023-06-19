package response

type ProductResponse struct {
	ID                  int64                     `json:"id"`
	Code                string                    `json:"code"`
	Name                string                    `json:"name"`
	UnitMassAcronym     string                    `json:"unit_mass_acronym"`
	UnitMassDescription string                    `json:"unit_mass_description"`
	CreatedAt           string                    `json:"created_at"`
	UpdatedAt           string                    `json:"updated_at"`
	ProductQualities    []*ProductQualityResponse `json:"product_qualities,omitempty"`
}
