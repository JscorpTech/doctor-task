package domain

type DoctorResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Specialty string `json:"specialty"`
	WorkStart string `json:"work_start"`
	WorkEnd   string `json:"work_end"`
}

type DoctorDetailResponse struct {
	DoctorResponse
	Appointments []*AppointmentResponse `json:"appointments"`
}

type DoctorUsecase interface {
	List(search string) []*DoctorResponse
	FindByID(uint) (*DoctorDetailResponse, error)
}
