package feedbackroute

import (
	"net/http"

	feedbackcontroller "github.com/Akshat-Srivastava2004/educationportal/controllers/Feedback_controller"
	"github.com/Akshat-Srivastava2004/educationportal/middleware"
	"github.com/gorilla/mux"
)

func FeedbackRouter(r *mux.Router) {

	r.Handle("/feedback", middleware.AuthMiddleware(http.HandlerFunc(feedbackcontroller.Feedback))).Methods("POST")
}
