package response

type ProductResponse struct {
	ID                  int64                     `json:"id"`
	Code                string                    `json:"code"`
	Name                string                    `json:"name"`
	UnitMassAcronym     string                    `json:"unit_mass_acronym,omitempty"`
	UnitMassDescription string                    `json:"unit_mass_description,omitempty"`
	CreatedAt           string                    `json:"created_at,omitempty"`
	UpdatedAt           string                    `json:"updated_at,omitempty"`
	ProductQualities    []*ProductQualityResponse `json:"product_qualities,omitempty"`
}
