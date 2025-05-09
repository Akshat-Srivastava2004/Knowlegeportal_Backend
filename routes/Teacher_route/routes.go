package teacherroute

import (
	"net/http"

	teachercontroller "github.com/Akshat-Srivastava2004/educationportal/controllers/Teacher_controller"
	"github.com/Akshat-Srivastava2004/educationportal/middleware"
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	
	r.HandleFunc("/teacherregister", teachercontroller.CreateUser).Methods("POST")
	r.HandleFunc("/teacherlogin", teachercontroller.Checkuser).Methods("POST")
    r.HandleFunc("/teacherlogintoken",teachercontroller.Checkuserserver).Methods("POST")
	r.Handle("/teacherresume", middleware.AuthMiddleware(http.HandlerFunc(teachercontroller.UploadResumeHandler))).Methods("POST")
	r.HandleFunc("/teacherdetails", teachercontroller.TeacherDashboard).Methods("GET")
	r.HandleFunc("/teachertest", teachercontroller.TeacherMCq).Methods("GET")
	r.HandleFunc("/upload", teachercontroller.UploadFile).Methods("POST")
	r.Handle("/quiz", middleware.AuthMiddleware(http.HandlerFunc(teachercontroller.FetchQuestionsHandler))).Methods("GET")
	r.Handle("/evaluate",middleware.AuthMiddleware(http.HandlerFunc(teachercontroller.EvaluateAnswersHandler))).Methods("POST")
	// r.HandleFunc("/submit-mcq", teachercontroller.SubmitMCq).Methods("POST")

	// r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./template"))))

}
