package controllers

import (
	"net/http"
	"backend/api/utils"
	"backend/api/models"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	logs, err := models.PaginateLogs(utils.PageRequest(r, 5))
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.ToJson(w, logs, http.StatusOK)
}