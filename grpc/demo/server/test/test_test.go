package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gfgin "github.com/thoohv5/gf/gin"
	pbuser "github.com/thoohv5/gf/grpc/demo/sdk/user"
	"github.com/thoohv5/gf/grpc/server"
)

type DemoReq struct {
	Code string `form:"code" json:"code" binding:"required,test" label:"代码"`
}

func Test_testServer_Info(t *testing.T) {

	server.NewServer().Server(&server.Config{
		Network: "tcp",
		Address: "0.0.0.0:8028",
	}, func(gServer *grpc.Server) error {
		pbuser.RegisterUserServer(gServer, NewServer())
		return nil
	}, func(ctx context.Context, mux *runtime.ServeMux, addr string, dialOption []grpc.DialOption) error {
		if err := pbuser.RegisterUserHandlerFromEndpoint(ctx, mux, addr, dialOption); nil != err {
			return err
		}
		return nil
	}, func(mux *http.ServeMux) error {

		router, _ := gfgin.New().Handle(func(router *gin.Engine, trans ut.Translator) error {
			api := router.Group("/api")
			{
				api.GET("/demo", func(c *gin.Context) {
					req := new(DemoReq)

					if err := c.ShouldBind(req); nil != err {
						validatorErr, ok := err.(validator.ValidationErrors)
						if ok {
							err := validatorErr.Translate(trans)
							c.JSON(http.StatusBadRequest, gin.H{"error": err})
							return
						}

						c.JSON(http.StatusBadRequest, gin.H{"error": err})
						return
					}
					c.JSON(http.StatusOK, gin.H{"req": req})
				})
			}
			return nil
		})
		mux.Handle("/api/", router)
		return nil
	})
}
