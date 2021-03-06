package main

import (
	"ProjectDelivery2/component"
	"ProjectDelivery2/middleware"
	"ProjectDelivery2/modules/restaurant/restauranttransport/restaurantgin"
	"ProjectDelivery2/modules/restaurantlike/transport/ginrestaurantlike"
	"ProjectDelivery2/modules/user/usertransport/ginuser"
	"ProjectDelivery2/pubsub/pblocal"
	"ProjectDelivery2/skio"
	"ProjectDelivery2/subscriber"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	fmt.Println("Loi khi load file env")
	//}
	USERNAME := os.Getenv("username")
	PASSWORD := os.Getenv("PASSWORD")
	DatabaseHost := os.Getenv("DATABASE_HOST")
	DatabasePort := os.Getenv("DATABASE_PORT")
	DatabaseName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, DatabaseHost, DatabasePort, DatabaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("loi khi ket noi database")
	}
	runService(db)
}
func runService(db *gorm.DB) error {
	db = db.Debug()
	appContext := component.NewAppContext(db, os.Getenv("SECRET"), pblocal.NewPubSub())
	r := gin.Default()

	//subscriber.SetupSubscribers(appContext)
	rtEngine := skio.NewEngine()

	if err := subscriber.NewEngine(appContext,rtEngine).Start(); err != nil {
		log.Fatalln(err)
	}
	rtEngine.Run(appContext,r)

	r.StaticFile("/demo/", "./demo.html")

	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	v1 := r.Group("/v1")
	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/login", ginuser.Login(appContext))
	v1.GET("/profile", middleware.RequireAuth(appContext), ginuser.GetProfile(appContext))
	restaurants := v1.Group("/restaurants")
	{
		restaurants.POST("", middleware.RequireAuth(appContext), restaurantgin.CreateRestaurant(appContext))
		restaurants.GET("", restaurantgin.ListRestaurant(appContext))
		restaurants.GET("/:id", restaurantgin.GetRestaurant(appContext))
		restaurants.PUT("/:id", restaurantgin.UpdateRestaurant(appContext))
		restaurants.GET("/:id/liked_users", ginrestaurantlike.ListUser(appContext))
		restaurants.POST(":id/un-like", middleware.RequireAuth(appContext), ginrestaurantlike.UnLikeRestaurant(appContext))
		restaurants.POST(":id/like", middleware.RequireAuth(appContext), ginrestaurantlike.LikeRestaurant(appContext))
	}
	////exporter
	agentEndpointURI := "localhost:6831"


	exporter, _ := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint:          agentEndpointURI,
		Process:             jaeger.Process{ServiceName: "Deliveryfood"},
		//n??n ?????t bi???n m??i tr?????ng ????? devops bi???t service th??? m???y
	})
	//if err != nil {
	//	log.Fatalf("Failed to create the Jaeger exporter: %v", err)
	//}
	// And now finally register it as a Trace Exporter

	 trace.RegisterExporter(exporter)
	 trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(1)})
	 return http.ListenAndServe("localhost:8080",&ochttp.Handler{
			Handler: r,
	 })

	// return r.Run()
}



// 1 ????n h??ng ???????c t???o thong qua rest API r???i th?? server s??? c??
//nhi???m v??? thong b??o cho nh???ng socket lien quan t???i rest api ????
//ch??? c???a h??ng , shipper ,ho???c ch??? nh?? h??ng , ho???c ch??nh user ???? ch???ng h???n
/*H??? th???ng danh s??ch t??i x???
1/Lo???i t??i x??? onl
2/Lo???i t??i x??? ko onl
khi kh??ch booking 1 m??n c???a h??ng a
M??nh c?? v??? tr?? kh??ch h??ng vs nh?? h??ng
R???i t??m t??i x??? online ??u ti??n g???n nh???t (khong c?? ????n h??ng), N???u c?? ng duy???t th?? ????n h??ng ??c k??ch ho???t
Sau khi giao t???i kh??ch h??ng th??nh c??ng th?? kh??ch h??ng x??c nh???n th??nh c??ng
=>>>>>>>>>
T??? user => l???y list t??i x??? is online
C?? th??? l??u lat long c???a user t??i xe trong db => query nearby sql t??m lat long g???n gi???ng
mu???n chuy??n h??n d??ng db chuy??n v??? location nh?? mongodb , geofencing ,elastic search
*/


