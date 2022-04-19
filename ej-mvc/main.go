package main
import (
	//"fmt"
	"net/http"
	"github.con/gin-gonic/gin"
)

func main(){
	engine := gin.New()

	engine.GET("/ping", Ping)
	engine.POST("/Test", Test)
	engine.Run(":3000")
}

func Ping(ctx *gin.Context){
	ctx.String(200, "pong")
}

type Body struct{
	Name string `json:"name"`
}
func Test(ctx *gin.Context){
	var body BodyConst

	err := ctx.BindJSON(&body)
	if err != nil{
		ctx.AbortWithError(400,err)
		return
	}
	ctx.JSON(http.StatusOK, body)
}
