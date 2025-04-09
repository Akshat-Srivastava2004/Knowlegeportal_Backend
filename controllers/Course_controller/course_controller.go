package coursecontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Akshat-Srivastava2004/educationportal/database"
	model "github.com/Akshat-Srivastava2004/educationportal/models/Course_model"
)

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://blue-meadow-0b28d241e.6.azurestaticapps.net/")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	courseduaration := r.FormValue("Courseduration")
	courseprice := r.FormValue("Courseprice")

	var course model.CourseStudent

	course.Coursename = r.FormValue("Coursename")
	course.Description = r.FormValue("Description")
	course.Category = r.FormValue("Category")
	course.Instructor = r.FormValue("Instructor")
	course.Prerequist = r.FormValue("Prerequist")
	course.Rating = r.FormValue("Rating")
	course.Language = r.FormValue("Language")

	Coursedur, err := strconv.ParseInt(courseduaration, 10, 64)
	if err != nil {
		http.Error(w, "Invalid courseduaration", http.StatusBadRequest)
	}
	Coursepric, err := strconv.ParseInt(courseprice, 10, 64)
	if err != nil {
		http.Error(w, "Invalid courseduaration", http.StatusBadRequest)
	}
	course.Courseduration = Coursedur
	course.Courseprice = Coursepric

	insertedID := database.Insertcourse(course)
	course.ID = insertedID

	fmt.Println("the inserted idis ", insertedID)
	fmt.Println("the coursed added successfully")
	http.Redirect(w, r, "https://blue-meadow-0b28d241e.6.azurestaticapps.net/form.html", http.StatusSeeOther)
}
