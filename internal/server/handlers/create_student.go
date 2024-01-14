package server

import (
	"net/http"

	"github.com/RomanTykhyi/students-api/internal/data"
	"github.com/RomanTykhyi/students-api/internal/di"
	"github.com/RomanTykhyi/students-api/internal/models"
	utils "github.com/RomanTykhyi/students-api/internal/server/utils"
	"github.com/google/uuid"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	repo, err := di.GetAppContainer().Resolve("students-store")
	if err != nil {
		panic("Cannot get students repository")
	}
	studentsRepo := repo.(data.StudentsStore)

	if err := r.ParseForm(); err != nil {
		utils.WriteMessageResponse(w, "Failed to parse form...", http.StatusBadRequest)
		return
	}

	fullName := r.PostForm["FullName"][0]

	student := models.Student{
		PartitionId: "students",
		Id:          uuid.New().String(),
		FullName:    fullName,
	}

	err = studentsRepo.PutStudent(&student)
	if err != nil {
		utils.WriteMessageResponse(w, "Error putting student.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.WriteString(w, student.Id)
}
