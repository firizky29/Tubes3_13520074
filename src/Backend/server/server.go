package server

import (
	"backend/disease"
	"backend/service"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-contrib/static"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type Server struct {
	Router            *gin.Engine
	Database          *gorm.DB
	DiseaseService    service.DiseaseService
	PredictionService service.PredictionService
}

func NewServer(config DatabaseConfig) *Server {
	s := &Server{}
	s.initService(config)
	s.initRouter()
	return s
}

func (s *Server) initService(config DatabaseConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Database)
	sqlDB, err := sql.Open("mysql", dsn)

	s.Database, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database...")
	}

	fmt.Printf("Creating Migration to Database...")
	err = s.Database.AutoMigrate(&disease.Disease{})
	if err != nil {
		log.Fatal(err)
	}
	err = s.Database.AutoMigrate(&disease.DiseasePrediction{})
	if err != nil {
		log.Fatal(err)
	}

	DiseaseList := disease.NewDisease(s.Database)
	PredictionLog := disease.NewPredictionLogs(s.Database)

	s.DiseaseService = service.NewDiseaseService(DiseaseList)
	s.PredictionService = service.NewPredictionService(PredictionLog)
}

func (s *Server) initRouter() {
	s.Router = gin.Default()
	s.Router.POST("/inputpenyakit", s.PostDisease)
	s.Router.Use(static.Serve("/", static.LocalFile("../Frontend/rucikawavin/dist", true)))
	s.Router.Use(static.Serve("/inputpenyakit", static.LocalFile("../Frontend/rucikawavin/dist", true)))
	s.Router.Use(static.Serve("/tesdna", static.LocalFile("../Frontend/rucikawavin/dist", true)))
	s.Router.Use(static.Serve("/hasilprediksi", static.LocalFile("../Frontend/rucikawavin/dist", true)))

	s.Router.LoadHTMLGlob("../Frontend/rucikawavin/dist/index.html")
	s.Router.Run()
}

func (s *Server) Run(port string) {
	err := s.Router.Run(port)
	if err != nil {
		panic("Failed to listen to port")
	}
}
