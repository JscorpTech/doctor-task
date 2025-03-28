package domain

type DoctorResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Specialty string `json:"specialty"`
	WorkStart string `json:"work_start"`
	WorkEnd   string `json:"work_end"`
}

type DoctorUsecase interface {
	List() ([]*DoctorResponse, error)
}
