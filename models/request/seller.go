package request

type SellerRegistrationRequest struct {
	FullName         string `json:"fullName"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	BusinessName     string `json:"businessName"`
	BusinessCategory string `json:"businessCategory"`
	Username         string `json:"businessUsername"`
	Password         string `json:"password"`
}
