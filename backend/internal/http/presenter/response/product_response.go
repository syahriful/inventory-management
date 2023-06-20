package response

type ProductResponse struct {
	ID                  int64                     `json:"id,omitempty"`
	Code                string                    `json:"code,omitempty"`
	Name                string                    `json:"name,omitempty"`
	UnitMassAcronym     string                    `json:"unit_mass_acronym,omitempty"`
	UnitMassDescription string                    `json:"unit_mass_description,omitempty"`
	CreatedAt           string                    `json:"created_at,omitempty"`
	UpdatedAt           string                    `json:"updated_at,omitempty"`
	ProductQualities    []*ProductQualityResponse `json:"product_qualities,omitempty"`
}
