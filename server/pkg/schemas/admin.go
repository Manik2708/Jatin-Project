package schemas

type Admin struct {
	BaseUser
	AcceptedAppointmentRequests   []Appointment `bson:"accepted_appointment_requests,omitempty" json:"accepted_appointment_requests,omitempty"`
	PendingAppointmentRequests  []Appointment `bson:"pending_appointment_requests,omitempty" json:"pending_appointment_requests,omitempty"`
	RejectedAppointmentRequests []Appointment `bson:"rejected_appointment_requests,omitempty" json:"rejected_appointment_requests,omitempty"`
	ClosedAppointmentRequests   []Appointment `bson:"closed_appointment_requests,omitempty" json:"closed_appointment_requests,omitempty"`
}

type AdminOrUser interface {
	*Customer | *Admin
}
