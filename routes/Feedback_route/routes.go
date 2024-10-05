package feedbackroute

import (
	feedbackcontroller "github.com/Akshat-Srivastava2004/educationportal/controllers/Feedback_controller"
	"github.com/gorilla/mux"
)

func FeedbackRouter(r *mux.Router) {

	r.HandleFunc("/feedback", feedbackcontroller.Feedback).Methods("POST")
}
