package validator

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Config() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	v.RegisterValidation("permission_names", permissionNames)
}

var permissionNames validator.Func = func(fl validator.FieldLevel) bool {
	names, ok := fl.Field().Interface().([]string)
	if !ok {
		return false
	}

	// 合法的权限名称应该是：controller#action
	pattern := regexp.MustCompile(`^\w+#\w+$`)

	for _, name := range names {
		if !pattern.MatchString(name) {
			return false
		}
	}

	return true
}
