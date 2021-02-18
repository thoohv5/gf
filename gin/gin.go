package gin

import (
	"net/http"

	gonicgin "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	"github.com/thoohv5/gf/gin/middleware"
	libvalidator "github.com/thoohv5/gf/validator"
)

type gin struct {
}

type RegisterRouter func(engine *gonicgin.Engine, trans ut.Translator) error

func New() *gin {
	return &gin{}
}

func (g *gin) Handle(registerRouter RegisterRouter) (http.Handler, error) {
	router := gonicgin.New()

	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := libvalidator.NewValidator().Register(v, trans); nil != err {
			return nil, errors.Wrap(err, "validator Register err")
		}
	}

	gonicgin.SetMode(gonicgin.DebugMode)

	router.Use(middleware.Recovery(), middleware.Logger(), middleware.Cors(), middleware.Opentracing())

	if err := registerRouter(router, trans); nil != err {
		return router, err
	}

	return router, nil
}
