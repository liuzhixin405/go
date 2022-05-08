package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateResource godoc
// @Summary Creates a account
// @Description creates Resource directory
// @Tags Accounts
// @Param name path string true "name"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /account/create [post]
func CreateAccount(c *gin.Context) {
	fmt.Println("Creating account")
	name := c.Params.ByName("name")

	fmt.Println("Creating account called for %s", name)
	c.JSON(http.StatusOK, "Account created")

}

// UpdateAccount godoc
// @Summary updates account
// @Description creates Resource directory
// @Tags Accounts
// @Param name path string true "uuid"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /account/update [patch]
func UpdateAccount(c *gin.Context) {
	fmt.Println("Updating account")

	id := c.Params.ByName("id")

	fmt.Println("Creating account called for %s", id)
	c.JSON(http.StatusOK, "Account Updated")

}

// DeleteAccount godoc
// @Summary delete account
// @Description creates Resource directory
// @Tags Accounts
//@Param name path string true "uuid"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /account/delete [delete]
func DeleteAccount(c *gin.Context) {
	fmt.Println("Deleting account")
	id := c.Params.ByName("id")

	fmt.Println("Creating account called for %s", id)
	c.JSON(http.StatusOK, "Account Deleted")

}
