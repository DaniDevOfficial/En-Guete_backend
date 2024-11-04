package meal

import (
	"database/sql"
	"enguete/modules/group"
	"enguete/util/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateNewMeal(c *gin.Context, db *sql.DB) {
	var newMeal RequestNewMeal
	err := c.ShouldBindJSON(&newMeal)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, MealError{Error: "Invalid request body"})
		return
	}

	jwtPayload, err := auth.GetJWTPayloadFromHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, MealError{Error: "Unauthorized"})
		return
	}
	err = group.CheckIfUserIsAdminOrOwnerOfGroupInDB(newMeal.GroupId, jwtPayload.UserId, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, MealError{Error: "Unauthorized"})
		return
	}

	mealId, err := CreateNewMealInDB(newMeal, jwtPayload.UserId, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MealError{Error: "Internal server error"})
		return
	}
	log.Println("New meal created with id:", mealId)
	c.JSON(http.StatusCreated, ResponseNewMeal{MealId: mealId})
}

func AddCookToMeal(c *gin.Context, db *sql.DB) {
	var addCookToMealData RequestAddCookToMeal

	if err := c.ShouldBindJSON(&addCookToMealData); err != nil {
		c.JSON(http.StatusBadRequest, MealError{Error: "Invalid request body"})
		return
	}
	jwtPayload, err := auth.GetJWTPayloadFromHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, MealError{Error: "Unauthorized"})
		return
	}
	err = group.CheckIfUserIsAdminOrOwnerOfGroupInDB(newMeal.GroupId, jwtPayload.UserId, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, MealError{Error: "Unauthorized"})
		return
	}
}
