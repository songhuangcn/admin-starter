package locale

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	log "github.com/sirupsen/logrus"
	. "github.com/songhuangcn/admin-template/internal/common/core"
	"github.com/songhuangcn/admin-template/internal/common/enum"
	"github.com/vorlif/spreak"
	"github.com/vorlif/spreak/localize"
	"golang.org/x/text/language"
)

// 这里有非基本类型，无法使用常量存储，因此使用 var
var (
	LocaleBundle  = newLocaleBundle()
	DefaultDomain = "app"
	DefaultPath   = "./locales"
	DefaultLang   = language.Chinese
	AllLangs      = []language.Tag{
		// 默认语言需要排在第一位，匹配时未找到合适语言时用的第一个
		language.Chinese,
		language.English,
	}

	// go-playground/validator 的翻译
	uni = ut.New(zh.New(), en.New())
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerLang := ctx.GetHeader("Accept-Language")

		lang := getLang(headerLang)
		log.Debugf("lang: %#v", lang.String())
		locale := spreak.NewLocalizer(LocaleBundle, lang)
		ctx.Set(enum.Locale, locale)

		// go-playground/validator 的翻译
		var translang string
		switch headerLang {
		case "en", "en-US", "en_US":
			translang = "en"
		default:
			translang = "zh"
		}
		validatorTrans, _ := uni.GetTranslator(translang)
		v, _ := binding.Validator.Engine().(*validator.Validate)
		if translang == "en" {
			en_translations.RegisterDefaultTranslations(v, validatorTrans)
		} else {
			zh_translations.RegisterDefaultTranslations(v, validatorTrans)
		}
		ctx.Set(enum.ValidatorTrans, validatorTrans)
	}
}

func T(ctx context.Context, msgid localize.MsgID, args ...any) string {
	locale := ctx.Value(enum.Locale).(*spreak.Localizer)

	// TODO: 支持 args
	return locale.Get(msgid)
}

func newLocaleBundle() *spreak.Bundle {
	langs := Map(AllLangs, func(lang language.Tag) any { return any(lang) })
	bundle, err := spreak.NewBundle(
		// 不要指定源语言，因为指定后源语言就无法翻译
		// spreak.WithSourceLanguage(DefaultLang),
		// Set the path from which the translations should be loaded
		spreak.WithDomainPath(DefaultDomain, DefaultPath),
		// Specify the languages you want to load
		spreak.WithRequiredLanguage(langs...),
		// 指定默认 domain 之后，就不需要使用 NewKeyLocalizerForDomain，每次创建 localizer 时再指定
		spreak.WithDefaultDomain(DefaultDomain),
	)
	if err != nil {
		panic(err)
	}

	return bundle
}

func getLang(headerLang string) language.Tag {
	log.Debugf("header lang: %#v", headerLang)
	// 默认没匹配时，会取切片中第一个语言
	lang, _ := language.MatchStrings(
		language.NewMatcher(AllLangs),
		headerLang,
	)

	return lang
}
