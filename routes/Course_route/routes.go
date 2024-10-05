package courseroute

import (
	coursecontroller "github.com/Akshat-Srivastava2004/educationportal/controllers/Course_controller"
	studentenrollcontroller "github.com/Akshat-Srivastava2004/educationportal/controllers/Studentenroll_controller"
	"github.com/gorilla/mux"
)

func CourseRouter(r *mux.Router) {
	r.HandleFunc("/coursestudent", coursecontroller.CreateCourse).Methods("POST")
	r.HandleFunc("/enrollstudent", studentenrollcontroller.Enrolledstudent).Methods("POST")

}
