package main

import (
	"cp/app/controller"
	"cp/app/server"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)







// Routes matches methods with endpoints
func Routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))

	r.GET("/", controller.StartPage)

	r.GET("/bookings", controller.SelectBookings)
	r.POST("/booking/insert", controller.InsertBooking)
	r.POST("/booking/delete", controller.DeleteBooking)
	
	r.GET("/rooms", controller.SelectRooms)
	r.POST("/rooms/update/category", controller.UpdateCategory)
	r.POST("/rooms/update/price", controller.UpdatePrice)
	
	r.GET("/clients", controller.SelectClients)

	r.GET("/staff", controller.SelectStaff)
	r.POST("/staff/insert", controller.InsertStaff)
	r.POST("/staff/delete", controller.DeleteStaff)

	
}

// InitConfig initialises configuration file
func InitConfig() error {
	viper.SetConfigFile("config.yml")
	return viper.ReadInConfig()
}

func main() {
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}
	configStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("db.host"), viper.GetString("db.port"),
		viper.GetString("db.user"), viper.GetString("db.dbname"),
		viper.GetString("db.password"), viper.GetString("db.sslmode"))

	if err := server.OpenDB(configStr); err != nil {
		log.Fatal(err)
		return
	}
	r := httprouter.New()
	Routes(r)
	if err := http.ListenAndServe(":"+viper.GetString("port"), r); err != nil {
		log.Fatal(err)
		return
	}
	if err := server.CloseDB(); err != nil {
		log.Fatal(err)
		return
	}
}
