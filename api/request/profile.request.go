package request

type Profile struct {
	Name           string            `json:"name" validate:"required"`
	Description    string            `json:"description" validate:"required,min=10"`
	Gender         string            `json:"gender" validate:"required,oneof=male female other"`
	DateOfBirth    string            `json:"date_of_birth" validate:"required"`
	Preference     ProfilePreference `json:"preference"`
	PremiumPackage PremiumPackage    `json:"premium_package"`
}

type ProfilePreference struct {
	Gender     string `json:"gender" validate:"required,oneof=male female other"`
	MinimumAge int    `json:"minimum_age" validate:"required,min=18"`
	MaximumAge int    `json:"maximum_age" validate:"required,max=100"`
}

type PremiumPackage struct {
	PurchaseDate string `json:"purchase_date"`
	ExpireDate   string `json:"expire_date"`
}
