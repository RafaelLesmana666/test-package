package main

import (
	"unit-test/config"

	jamethelper "github.com/RafaelLesmana666/jamet-helper"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Open()

	met := config.Jamet

	router := gin.Default()
	router.GET("/service/pkb", func(c *gin.Context) {
		query := met.GetData("pkbs", "")

		jamethelper.CreateData(c, query, []string{"id"})
	})

	router.POST("/service/pkb", func(c *gin.Context) {

		request := jamethelper.Converter(c)

		//validation
		check := map[string]map[string]string{
			"field": {
				"pkb_type":             "required",
				"kilometer":            "required",
				"kode_alasan":          "required",
				"est_waktu_pengerjaan": "required",
				"est_waktu_tunggu":     "required",
				"signature":            "required",
				"sa_signature":         "required",
				"hscs_part_not_ok":     "required",
				"hscs_catatan_teknisi": "required",
				"hscs_signature":       "required",
			},
			"alias": {
				"est_waktu_pengerjaan": "Waktu Pengerjaan",
				"est_waktu_tunggu":     "Waktu Tunggu",
				"hscs_part_not_ok":     "Part Not Ok",
				"hscs_catatan_teknisi": "Catatan Teknisi Hscs",
				"sa_signature":         "Tanda Tangan Customer",
				"signature":            "Tanda Tangan SA",
				"hscs_signature":       "Tanda Tangan HSCS",
			},
			"message": {
				"pkb_type":    "Jenis Layanan belum dipilih",
				"kode_alasan": "Alasan Ke Ahass belum dipilih",
			},
		}

		err := jamethelper.Validation(request, check)
		if err != "" {
			jamethelper.PrintJSON(c, jamethelper.Response{
				Status:  false,
				Message: err,
			})
			return
		}

		connection := config.Jamet.Connection("default")

		data := map[string]interface{}{
			"name":  "Alice",
			"age":   30,
			"email": "alice@example.com",
		}

		if err := jamethelper.InsertData(c, connection, "barangs", data); err != nil {
			jamethelper.EPrintJSON(c, jamethelper.Response{
				Status:  false,
				Message: err,
			})
			return
		}

		connection.Commit()
		jamethelper.PrintJSON(c, jamethelper.Response{
			Status:  false,
			Message: "Data Berhasil Tersimpan!",
		})
	})

	router.Run()

}
