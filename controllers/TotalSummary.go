package controllers

import (
	"net/http"

	"github.com/ilham13/Covid-2020/config"
)

// TotalSummaryController struct
type TotalSummaryController struct{}

// TotalSummaryResponse response
type TotalSummaryResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SummaryResponse response
type SummaryResponse struct {
	Meninggal   float64 `json:"meninggal"`
	Sembuh      float64 `json:"sembuh"`
	Perawatan   float64 `json:"perawatan"`
	JumlahKasus float64 `json:"jumlah_kasus"`
}

// GetList get list all
func (a *TotalSummaryController) GetList(w http.ResponseWriter, r *http.Request) {

	resp, err := GetService(config.CovidIndonesiaURL)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	// perProvinsi := resp["perProvinsi"].(map[string]interface{})

	summary := SummaryResponse{}
	summary.Meninggal = resp["meninggal"].(float64)
	summary.Sembuh = resp["sembuh"].(float64)
	summary.Perawatan = resp["perawatan"].(float64)
	summary.JumlahKasus = resp["jumlahKasus"].(float64)

	data := TotalSummaryResponse{true, "success", summary}
	respondWithJSON(w, http.StatusOK, data)
}
