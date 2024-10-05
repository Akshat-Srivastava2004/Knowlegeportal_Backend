package studentroute

import (
	contactcontroller "github.com/Akshat-Srivastava2004/educationportal/controllers/Contact_controller"
	studentcontroller "github.com/Akshat-Srivastava2004/educationportal/controllers/Student_controller"
	"github.com/gorilla/mux"
)

func StudentRouter(r *mux.Router) {

	r.HandleFunc("/studentregister", studentcontroller.CreateUserstudent).Methods("POST")
	r.HandleFunc("/studentlogin", studentcontroller.Checkuserstudent).Methods("POST")
	r.HandleFunc("/studentdetails", studentcontroller.StudentDashboard).Methods("GET")

	r.HandleFunc("/contact", contactcontroller.ContactPageHandler)
	r.HandleFunc("/studentenroll", contactcontroller.StudentEnrollmentPageHandler)
}
