package apis

import (
	"github.com/gin-gonic/gin"
	"go-member-api/controllers"
	"go-member-api/models"
	"log"
	"net/http"
	"strconv"
)

// GetMember godoc
// @Summary Retrieves member based on given client_member_id
// @Produce json
// @Param id path integer true "client_member_id"
// @Message:success Status:true {object} models.User
// @Router /client_member_id/{client_member_id} [get]
// @Security ApiKeyAuth
func GetMember(c) {
	//pass
}
