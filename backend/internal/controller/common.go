package controller

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/common/enum"
)

func setPagination(ctx *gin.Context, page, pageSize int, total int64) {
	ctx.Set(enum.Pagination, &Hash{"page": page, "per_page": pageSize, "total": total})
}

func renderSuccess(ctx *gin.Context, opts ...Hash) {
	opt := Hash{}
	if len(opts) > 0 {
		opt = opts[0]
	}
	status := HashGet[int](opt, "status", 200)
	data := HashGet[any](opt, "data", Hash{})
	meta := HashGet[Hash](opt, "meta", Hash{})

	pagination, ok := ctx.Value(enum.Pagination).(*Hash)
	if ok {
		meta["pagination"] = pagination
	}

	ctx.JSON(status, Hash{"data": data, "meta": meta})
}

func panicBind(ctx *gin.Context, obj any) {
	if err := ctx.ShouldBind(obj); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			PanicApiError(err.Error(), 422)
		}

		log.Debugf("errs: %#v\n", errs)
		validatorTrans := ctx.Value(enum.ValidatorTrans).(ut.Translator)
		msgs := errs.Translate(validatorTrans)

		var errorMessages []string
		for _, msg := range msgs {
			errorMessages = append(errorMessages, msg)
		}
		errorString := strings.Join(errorMessages, "; ")
		log.Debugf("errorString: %#v\n", errorString)

		PanicApiError(errorString, 422)
	}
}

func getParamID(ctx *gin.Context) uint {
	id, _ := strconv.Atoi(ctx.Param("id"))

	return uint(id)
}
