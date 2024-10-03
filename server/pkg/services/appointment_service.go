package services

import "jatin/pkg/schemas"

type AppointmentServiceTemplate interface {
	CreateAppointment() (*schemas.Appointment, error)
	UpdateAppointmentStatus() (*schemas.Appointment, error) 
	GetAllAppointmentsForAdmin() (*[]schemas.Appointment, error)
	GetAppointmentsByStatus() (*[]schemas.Appointment, error)
	GetAppointmentsForCustomer() (*[]schemas.Appointment, error)
}

func (as *GlobalService) CreateAppointment(){
	
}
