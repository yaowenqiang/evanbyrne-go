package controller

import (
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"net/http"
)

func GetAdminIndex() (int, string) {
	var db = new(util.Database)
	defer db.Close()
	posts := service.GetPostListing(db)
	return util.RespondTemplate(http.StatusOK, "template/admin/index.html", posts)
}