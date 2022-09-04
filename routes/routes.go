package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "final-project/controller"
    "final-project/middlewares"

    swaggerFiles "github.com/swaggo/files"     // swagger embed files
    ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // set db to gin context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

    r.POST("/register", controller.Register)
    r.POST("/login", controller.Login)
    r.POST("/ganti-password", controller.GantiPassword)

    r.GET("/pasien", controller.GetAllPasien)
    r.GET("/pasien/:id", controller.GetPasienById)

    pasienMiddlewareRoute := r.Group("/pasien")
    pasienMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    pasienMiddlewareRoute.POST("/", controller.CreatePasien)
    pasienMiddlewareRoute.PATCH("/:id", controller.UpdatePasien)
    pasienMiddlewareRoute.DELETE("/:id", controller.DeletePasien)

	r.GET("/dokter", controller.GetAllDokter)
    r.GET("/dokter/:id", controller.GetDokterById)

    dokterMiddlewareRoute := r.Group("/dokter")
    dokterMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    dokterMiddlewareRoute.POST("/", controller.CreateDokter)
    dokterMiddlewareRoute.PATCH("/:id", controller.UpdateDokter)
    dokterMiddlewareRoute.DELETE("/:id", controller.DeleteDokter)

    r.GET("/perawat", controller.GetAllPerawat)
    r.GET("/perawat/:id", controller.GetPerawatById)

    perawatMiddlewareRoute := r.Group("/perawat")
    perawatMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    perawatMiddlewareRoute.POST("/", controller.CreatePerawat)
    perawatMiddlewareRoute.PATCH("/:id", controller.UpdatePerawat)
    perawatMiddlewareRoute.DELETE("/:id", controller.DeletePerawat)

	r.GET("/kamar", controller.GetAllKamar)
    r.GET("/kamar/:id", controller.GetKamarById)

    kamarMiddlewareRoute := r.Group("/kamar")
    kamarMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    kamarMiddlewareRoute.POST("/", controller.CreateKamar)
    kamarMiddlewareRoute.PATCH("/:id", controller.UpdateKamar)
    kamarMiddlewareRoute.DELETE("/:id", controller.DeleteKamar)

    r.GET("/obat", controller.GetAllObat)
    r.GET("/obat/:id", controller.GetObatById)

    obatMiddlewareRoute := r.Group("/obat")
    obatMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    obatMiddlewareRoute.POST("/", controller.CreateObat)
    obatMiddlewareRoute.PATCH("/:id", controller.UpdateObat)
    obatMiddlewareRoute.DELETE("/:id", controller.DeleteObat)

    r.GET("/resep", controller.GetAllResep)
    r.GET("/resep/:id", controller.GetResepById)

    resepMiddlewareRoute := r.Group("/resep")
    resepMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    resepMiddlewareRoute.POST("/", controller.CreateResep)
    resepMiddlewareRoute.PATCH("/:id", controller.UpdateResep)
    resepMiddlewareRoute.DELETE("/:id", controller.DeleteResep)

    r.GET("/rawat-inap", controller.GetAllRawatInap)
    r.GET("/rawat-inap/:id", controller.GetRawatInapById)

    rawatInapMiddlewareRoute := r.Group("/rawatInap")
    rawatInapMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    rawatInapMiddlewareRoute.POST("/", controller.CreateRawatInap)
    rawatInapMiddlewareRoute.PATCH("/:id", controller.UpdateRawatInap)
    rawatInapMiddlewareRoute.DELETE("/:id", controller.DeleteRawatInap)

    r.GET("/appointment", controller.GetAllAppointment)
    r.GET("/appointment/:id", controller.GetAppointmentById)

    appointmentMiddlewareRoute := r.Group("/appointment")
    appointmentMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    appointmentMiddlewareRoute.POST("/", controller.CreateAppointment)
    appointmentMiddlewareRoute.PATCH("/:id", controller.UpdateAppointment)
    appointmentMiddlewareRoute.DELETE("/:id", controller.DeleteAppointment)

    r.GET("/diagnosa", controller.GetAllDiagnosa)
    r.GET("/diagnosa/:id", controller.GetDiagnosaById)

    diagnosaMiddlewareRoute := r.Group("/diagnosa")
    diagnosaMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    diagnosaMiddlewareRoute.POST("/", controller.CreateDiagnosa)
    diagnosaMiddlewareRoute.PATCH("/:id", controller.UpdateDiagnosa)
    diagnosaMiddlewareRoute.DELETE("/:id", controller.DeleteDiagnosa)

    r.GET("/transaksi", controller.GetAllTransaksi)
    r.GET("/transaksi/:id", controller.GetTransaksiById)

    transaksiMiddlewareRoute := r.Group("/transaksi")
    transaksiMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    transaksiMiddlewareRoute.POST("/", controller.CreateTransaksi)
    transaksiMiddlewareRoute.PATCH("/:id", controller.UpdateTransaksi)
    transaksiMiddlewareRoute.DELETE("/:id", controller.DeleteTransaksi)

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}