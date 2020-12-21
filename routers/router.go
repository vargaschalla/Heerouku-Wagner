package routers

import (
	"PROYECTintegrador/ProyectoGOI/app"
	"PROYECTintegrador/ProyectoGOI/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	conn, err := connectDBmysql()
	if err != nil {
		panic("failed to connect database: " + err.Error())
		//return
	}
	// Migrate the schema
	conn.AutoMigrate(
		&models.Persona{},
		&models.User{},
		&models.Curso{},
		&models.Periodo{},
		&models.PlanAcademico{},
		&models.Recurso{},
		&models.Rol{},
		&models.RolPersona{},
		&models.Seccion{},
		&models.Sesion{},
		&models.TipoRecurso{},
		&models.Trabajo{},
		&models.Unidad{},
		&models.Secuencia{},
	)

	r := gin.Default()

	//config := cors.DefaultConfig() https://github.com/rs/cors
	//config.AllowOrigins = []string{"http://localhost", "http://localhost:8086"}

	r.Use(CORSMiddleware())

	r.Use(dbMiddleware(*conn))

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", app.ItemsIndex)

		v1.GET("/personas", app.PersonsIndex)
		v1.POST("/personas", authMiddleWare(), app.PersonsCreate)
		v1.GET("/personas/:id", app.PersonsGet)
		v1.PUT("/personas/:id", app.PersonsUpdate)
		v1.DELETE("/personas/:id", app.PersonsDelete)

		v1.GET("/users", app.UsersIndex)
		v1.POST("/users", app.UsersCreate)
		v1.GET("/users/:id", app.UsersGet)
		v1.PUT("/users/:id", app.UsersUpdate)
		v1.DELETE("/users/:id", app.UsersDelete)
		v1.POST("/login", app.UsersLogin)
		v1.POST("/logout", app.UsersLogout)

		v1.GET("/curso", app.CursoIndex)
		v1.GET("/curso/:id", app.CursoGet)
		v1.POST("/curso", authMiddleWare(), app.CursoCreate)
		v1.PUT("/curso/:id", app.CursoUpdate)
		v1.DELETE("/curso/:id", app.CursoDelete)

		v1.GET("/periodo", app.PeriodoIndex)
		v1.POST("/periodo", authMiddleWare(), app.PeriodoCreate)
		v1.GET("/periodo/:id", app.PeriodoGet)
		v1.PUT("/periodo/:id", app.PeriodoUpdate)
		v1.DELETE("/periodo/:id", app.PeriodoDelete)

		v1.GET("/plan", app.PlanAcademicoIndex)
		v1.POST("/plan", authMiddleWare(), app.PlanAcademicoCreate)
		v1.GET("/plan/:id", app.PlanAcademicoGet)
		v1.PUT("/plan/:id", app.PlanAcademicoUpdate)
		v1.DELETE("/plan/:id", app.PlanAcademicoDelete)

		v1.GET("/recurso", app.RecursoIndex)
		v1.POST("/recurso", authMiddleWare(), app.RecursoCreate)
		v1.GET("/recurso/:id", app.RecursoGet)
		v1.PUT("/recurso/:id", app.RecursoUpdate)
		v1.DELETE("/recurso/:id", app.RecursoDelete)

		v1.GET("/rol", app.RolLista)
		v1.POST("/rol", authMiddleWare(), app.RolCreate)
		v1.GET("/rol/:id", app.RolGetID)
		v1.PUT("/rol/:id", app.RolUpdate)
		v1.DELETE("/rol/:id", app.RolDelete)

		v1.GET("/rolPersona", app.RolPersonaLista)
		v1.POST("/rolPersona", authMiddleWare(), app.RolPersonaCreate)
		v1.GET("/rolPersona/:id", app.RolPersonaGet)
		v1.PUT("/rolPersona/:id", app.RolPersonaUpdate)
		v1.DELETE("/rolPersona/:id", app.RolPersonaDelete)

		v1.GET("/seccion", app.SeccionIndex)
		v1.POST("/seccion", authMiddleWare(), app.SeccionCreate)
		v1.GET("/seccion/:id", app.SeccionGet)
		v1.PUT("/seccion/:id", app.SeccionUpdate)
		v1.DELETE("/seccion/:id", app.SeccionDelete)

		v1.GET("/sesion", app.SesionIndex)
		v1.POST("/sesion", authMiddleWare(), app.SesionCreate)
		v1.GET("/sesion/:id", app.SesionGet)
		v1.PUT("/sesion/:id", app.SesionUpdate)
		v1.DELETE("/sesion/:id", app.SesionDelete)

		v1.GET("/tipoRecurso", app.TipoRecursoIndex)
		v1.POST("/tipoRecurso", authMiddleWare(), app.TipoRecursoCreate)
		v1.GET("/tipoRecurso/:id", app.TipoRecursoGet)
		v1.PUT("/tipoRecurso/:id", app.TipoRecursoUpdate)
		v1.DELETE("/tipoRecurso/:id", app.TipoRecursoDelete)

		v1.GET("/trabajo", app.TrabajoIndex)
		v1.POST("/trabajo", authMiddleWare(), app.TrabajoCreate)
		v1.GET("/trabajo/:id", app.TrabajoGet)
		v1.PUT("/trabajo/:id", app.TrabajoUpdate)
		v1.DELETE("/trabajo/:id", app.TrabajoDelete)

		v1.GET("/unidad", app.UnidadIndex)
		v1.POST("/unidad", authMiddleWare(), app.UnidadCreate)
		v1.GET("/unidad/:id", app.UnidadGet)
		v1.PUT("/unidad/:id", app.UnidadUpdate)
		v1.DELETE("/unidad/:id", app.UnidadDelete)

		v1.GET("/secuencia", app.SecuenciaIndex)
		v1.POST("/secuencia", authMiddleWare(), app.SecuenciaCreate)
		v1.GET("/secuencia/:id", app.SecuenciaGet)
		v1.PUT("/secuencia/:id", app.SecuenciaUpdate)
		v1.DELETE("/secuencia/:id", app.SecuenciaDelete)

	}

	return r
}

func connectDBmysql() (c *gorm.DB, err error) {

	dsn := "root:aracelybriguit@tcp(localhost:3306)/wagner?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	return conn, err
}

func connectDB() (c *gorm.DB, err error) {
	////dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "user=wtsraexyfllzlj password=21bc9dcd59ddb8b14005d8bb0a535ef8d818293a55263f51e53e68cc831bb320 host=ec2-34-237-166-54.compute-1.amazonaws.com dbname=d7gjil23228635 port=5432 sslmode=require TimeZone=Asia/Shanghai"
	//dsn := "user=postgres password=postgres2 dbname=users_test host=localhost port=5435 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//c.Header("Access-Control-Allow-Origin", "http://localhost, http://localhost:8086,")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE ")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//https://dev.to/stevensunflash/a-working-solution-to-jwt-creation-and-invalidation-in-golang-4oe4

//https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
func authMiddleWare() gin.HandlerFunc { //ExtractToken
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated (IsTokenValid)."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
