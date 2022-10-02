package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paramount_school/internal/helpers"
	"paramount_school/internal/model"
)

// CreateHighSchool godoc
// @Summary      user creates high school
// @Description  Collects student's high school details during registration and uses it to create a new high school in the database.
// @Tags         HighSchool
// @Accept       json
// @Produce      json
// @Param HighSchool body model.HighSchool true "school, country, city, address, postal"
// @Success      201  {string} string "created successfully"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /create_high_schools [post]
func (u *HTTPHandler) CreateHighSchool(c *gin.Context) {
	var highSchool *model.HighSchool
	err := c.ShouldBindJSON(&highSchool)
	if err != nil {
		helpers.Response(c, "bad request", 400, nil, []string{"validation error"})
		return
	}
	err = u.Repository.CreateHighSchool(highSchool)
	if err != nil {
		helpers.Response(c, "Internal Server Error", http.StatusInternalServerError, nil, []string{"Could not create high schools"})
		return
	}

	helpers.Response(c, "created successfully", http.StatusCreated, nil, nil)
}

// GetHighSchools godoc
// @Summary      Lists all the high schools in the database
// @Description  This displays all the available high schools in the database for students to choose from during registration.
// @Tags         HighSchool
// @Accept       json
// @Produce      json
// @Success      200  {string} string "found successfully"
// @Failure      500  {string}  string "internal server error"
// @Router       /get_high_schools [get]
func (u *HTTPHandler) GetHighSchools(c *gin.Context) {
	highSchools, err := u.Repository.FindHighSchools()
	if err != nil {
		helpers.Response(c, "Internal Server Error", http.StatusInternalServerError, nil, []string{"Could not get high schools"})
		return
	}

	helpers.Response(c, "found successfully", 200, highSchools, nil)
}
