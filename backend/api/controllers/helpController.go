package controllers

import (
	"net/http"
	"backend/api/utils"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	utils.ToJson(w, utils.INFO, http.StatusOK)
}

func GetHelp(w http.ResponseWriter, r *http.Request) {
	utils.ToJson(w, utils.HELPS, http.StatusOK)
}