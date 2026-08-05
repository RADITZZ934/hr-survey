package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StoreLocation struct {
	AplikasiID int    `json:"aplikasi_id"`
	ID         int    `json:"id"`
	Kode       string `json:"kode"`
	Name       string `json:"name"`
	Koordinat  string `json:"koordinat"`
}

// GetStores handles GET /api/LIT/store.
// Mengembalikan daftar toko / DC Laskar Buah untuk integrasi survey external.
func GetStores(c *gin.Context) {
	stores := []StoreLocation{
		{AplikasiID: 1, ID: 1, Kode: "DC_MAGETAN_1", Name: "DC Magetan 1", Koordinat: "-7.6315,111.3283"},
		{AplikasiID: 1, ID: 2, Kode: "DC_GRESIK", Name: "DC Gresik", Koordinat: "-7.1566,112.6555"},
		{AplikasiID: 1, ID: 3, Kode: "DC_REMBANG_1", Name: "DC Rembang 1", Koordinat: "-6.7118,111.3444"},
		{AplikasiID: 1, ID: 4, Kode: "DC_JOMBANG_1", Name: "DC Jombang 1", Koordinat: "-7.5461,112.2331"},
		{AplikasiID: 1, ID: 5, Kode: "DC_PURWODADI", Name: "DC Purwodadi", Koordinat: "-7.0864,110.9169"},
		{AplikasiID: 1, ID: 6, Kode: "DC_KLATEN_1", Name: "DC Klaten 1", Koordinat: "-7.7025,110.6031"},
		{AplikasiID: 1, ID: 7, Kode: "DC_LAMONGAN", Name: "DC Lamongan", Koordinat: "-7.1198,112.4154"},
		{AplikasiID: 1, ID: 8, Kode: "LASKAR_BUAH_MADIUN", Name: "Laskar Buah Madiun", Koordinat: "-7.6298,111.5298"},
		{AplikasiID: 1, ID: 9, Kode: "LASKAR_BUAH_SOLO", Name: "Laskar Buah Solo", Koordinat: "-7.5755,110.8243"},
		{AplikasiID: 1, ID: 10, Kode: "LASKAR_BUAH_KEDIRI", Name: "Laskar Buah Kediri", Koordinat: "-7.8172,112.0119"},
	}

	c.JSON(http.StatusOK, stores)
}
