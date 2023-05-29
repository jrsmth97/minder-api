package seed

import (
	"log"
	"minder/src/server/model"

	"gorm.io/gorm"
)

func SeedLocation(db *gorm.DB) {
	for _, location := range locations {
		var locationExist []model.Location
		db.Where("location_name=?", location).Find(&locationExist)
		if len(locationExist) > 0 {
			continue
		}

		locationModel := model.Location{
			LocationName: location,
		}

		err := db.Debug().Model(&model.Location{}).Create(&locationModel).Error
		if err != nil {
			log.Fatalf("cannot seed location table: %v", err)
		}
	}
}

var locations = []string{
	"Kepulauan Seribu",
	"Jakarta Selatan",
	"Jakarta Timur",
	"Jakarta Pusat",
	"Jakarta Barat",
	"Jakarta Utara",
	"Bogor",
	"Sukabumi",
	"Cianjur",
	"Bandung",
	"Garut",
	"Tasikmalaya",
	"Ciamis",
	"Kuningan",
	"Cirebon",
	"Majalengka",
	"Sumedang",
	"Indramayu",
	"Subang",
	"Purwakarta",
	"Karawang",
	"Bekasi",
	"Bandung Barat",
	"Pangandaran",
	"Bogor",
	"Sukabumi",
	"Bandung",
	"Cirebon",
	"Bekasi",
	"Depok",
	"Cimahi",
	"Tasikmalaya",
	"Banjar",
	"Cilacap",
	"Banyumas",
	"Purbalingga",
	"Banjarnegara",
	"Kebumen",
	"Purworejo",
	"Wonosobo",
	"Magelang",
	"Boyolali",
	"Klaten",
	"Sukoharjo",
	"Wonogiri",
	"Karanganyar",
	"Sragen",
	"Grobogan",
	"Blora",
	"Rembang",
	"Pati",
	"Kudus",
	"Jepara",
	"Demak",
	"Semarang",
	"Temanggung",
	"Kendal",
	"Batang",
	"Pekalongan",
	"Pemalang",
	"Tegal",
	"Brebes",
	"Magelang",
	"Surakarta",
	"Salatiga",
	"Semarang",
	"Pekalongan",
	"Tegal",
	"Pacitan",
	"Ponorogo",
	"Trenggalek",
	"Tulungagung",
	"Blitar",
	"Kediri",
	"Malang",
	"Lumajang",
	"Jember",
	"Banyuwangi",
	"Bondowoso",
	"Situbondo",
	"Probolinggo",
	"Pasuruan",
	"Sidoarjo",
	"Mojokerto",
	"Jombang",
	"Nganjuk",
	"Madiun",
	"Magetan",
	"Ngawi",
	"Bojonegoro",
	"Tuban",
	"Lamongan",
	"Gresik",
	"Bangkalan",
	"Sampang",
	"Pamekasan",
	"Sumenep",
	"Kediri",
	"Blitar",
	"Malang",
	"Probolinggo",
	"Pasuruan",
	"Mojokerto",
	"Madiun",
	"Surabaya",
	"Batu",
}
