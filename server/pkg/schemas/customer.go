package schemas

type Customer struct{
	BaseUser
	Address 					Address 	  `bson:"address,omitempty" json:"address,omitempty"`
	Cars 						[]Car 		  `bson:"cars,omitempty" json:"cars,omitempty"`
	TriageAppointmentRequests   []Appointment `bson:"triage_appointment_requests,omitempty" json:"triage_appointment_requests,omitempty"`
	PendingAppointmentRequests  []Appointment `bson:"pending_appointment_requests,omitempty" json:"pending_appointment_requests,omitempty"`
	RejectedAppointmentRequests []Appointment `bson:"rejected_appointment_requests,omitempty" json:"rejected_appointment_requests,omitempty"`
	ClosedAppointmentRequests   []Appointment `bson:"closed_appointment_requests,omitempty" json:"closed_appointment_requests,omitempty"`
}