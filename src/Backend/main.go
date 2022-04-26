package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	// "regexp"
)

/*
Informasi terkait database bisa jadi tidak dibuat disini (dibaca dari file .env/.txt), untuk sementara simpan disini dulu
*/
const (
	host     = "rucikawavin.mysql.database.azure.com"
	database = "rucikawavin"
	user     = "rucikawavin@rucikawavin"
	password = "rucika.wavin123"
	port     = "3306"
)

func initDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	sqlDB, err := sql.Open("mysql", dsn)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return gormDB, err
}

func readVirus(path string) string {
	data, err := os.ReadFile(path)
	check(err)
	return string(data)
}

func readHuman(path string) string {
	data, err := os.ReadFile(path)
	check(err)
	return string(data)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World!")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	//fs := http.FileServer(http.Dir("../Frontend/rucikawavin/dist"))
	//http.Handle("/", fs)
	//
	//fmt.Println("Server listening on port 3000")
	//log.Panic(
	//	http.ListenAndServe(":3000", nil),
	//)
	//Router := gin.Default()
	//
	//Router.Static("/", "../Frontend/rucikawavin/dist")
	//Router.StaticFS("/more_static", http.Dir("my_file_system"))

	//Router.POST("/inputpenyakit", func(context *gin.Context) {
	//	var diseaseResponse DiseaseResponse
	//	err := context.ShouldBindJSON(&diseaseResponse)
	//	if err != nil {
	//		fmt.Printf("Error jsonnya")
	//	}
	//
	//	context.JSON(http.StatusOK, gin.H{
	//		"DiseaseName": diseaseResponse.DiseaseName,
	//		"DNA":         diseaseResponse.DNA,
	//	})
	//
	//	fmt.Printf(diseaseResponse.DiseaseName)
	//	fmt.Printf(diseaseResponse.DNA)
	//},
	//)

	//err := Router.Run()
	//if err != nil {
	//	return
	//}
	/* COBA-COBA */
	//var err error
	//var db *gorm.DB
	//db, err = initDB()
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = db.AutoMigrate(&disease.Disease{})
	//if err != nil {
	//	return
	//}
	//err = db.AutoMigrate(&disease.DiseasePrediction{})
	//if err != nil {
	//	return
	//}
	//
	//var repo = disease.NewPredictionLogs(db)
	//var prediction = disease.DiseasePrediction{
	//	UserName:          "Firizky Ardiansyah",
	//	DNA:               "ATCGCGCATGAATATCGATCGATGCATCGCGGCGCGCTAGTACGATCGATGCATGTACATGCATCGTAGCATCGATCGATCGATCGCATCG",
	//	DiseasePrediction: "virus",
	//	SimilarityLevel:   decimal.NewFromFloat32(9.29),
	//	Status:            false,
	//}
	//repo.Create(prediction)

	/*
		TODO: yang ini masih bingung, pengennya bisa parse generic date tp gagal wae
	*/
	var input = "02-02-2002 00:00:00 WIB"
	var inputToTime, err2 = time.Parse("2006-01-02 15:02:01 WIB", input)
	if err2 != nil {
		fmt.Printf("Input salah format")
		panic(err2)
	}
	fmt.Printf(inputToTime.String())
	//var logs []disease.DiseasePrediction
	//logs, _ = repo.FindByDateAndDiseaseName("virus", time.Now())
	//
	//for _, log := range logs {
	//	fmt.Printf(log.UserName)
	//}

}
