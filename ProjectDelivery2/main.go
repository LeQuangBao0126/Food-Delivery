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
		//nên đặt biến môi trường để devops biết service thứ mấy
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



// 1 Đơn hàng được tạo thong qua rest API rồi thì server sẽ có
//nhiệm vụ thong báo cho những socket lien quan tới rest api đó
//chủ cửa hàng , shipper ,hoặc chủ nhà hàng , hoặc chính user đó chẳng hạn
/*Hệ thống danh sách tài xế
1/Loại tài xế onl
2/Loại tài xế ko onl
khi khách booking 1 món của hàng a
Mình có vị trí khách hàng vs nhà hàng
Rồi tìm tài xế online ưu tiên gần nhất (khong có đơn hàng), Nếu có ng duyệt thì đơn hàng đc kích hoạt
Sau khi giao tới khách hàng thành công thì khách hàng xác nhận thành công
=>>>>>>>>>
Từ user => lấy list tài xế is online
Có thể lưu lat long của user tài xe trong db => query nearby sql tìm lat long gần giống
muốn chuyên hơn dùng db chuyên về location như mongodb , geofencing ,elastic search
*/


