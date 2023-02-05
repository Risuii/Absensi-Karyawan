package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"

	"github.com/Risuii/config"
	"github.com/Risuii/config/bcrypt"
	"github.com/Risuii/helpers/constant"
	"github.com/Risuii/internal/absensi"
	"github.com/Risuii/internal/activity"
	"github.com/Risuii/internal/user"
)

func main() {
	cfg := config.New()

	db, err := sql.Open("mysql", cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}

	validator := validator.New()
	router := mux.NewRouter()
	bcrypt := bcrypt.NewBcrypt(cfg.Bcrypt.HashCost)

	userRepo := user.NewUserRepository(db, constant.TableEmployee)
	userUseCase := user.NewUserUseCase(userRepo, bcrypt)

	activityRepo := activity.NewActivityRepositoryImpl(db, constant.TableActivity)
	activityUseCase := activity.NewActivityUseCaseImpl(activityRepo)

	absensiRepo := absensi.NewAbsensiRepositoryImpl(db, constant.TableAbsensi)
	absensiUseCase := absensi.NewAbsensiUseCase(absensiRepo)

	user.NewUserHandler(router, validator, userUseCase)
	activity.NewActivityHandler(router, validator, activityUseCase)
	absensi.NewAbsensiHandler(router, validator, absensiUseCase)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.App.Port),
		Handler: router,
	}

	port := os.Getenv("PORT")

	fmt.Println("SERVER ON")
	fmt.Println("PORT :", port)
	log.Fatal(server.ListenAndServe())
}
