package middlewares

import "github.com/gin-gonic/gin"

type MetaHandler gin.HandlerFunc

func NewMetaHandler() MetaHandler {
	return func(context *gin.Context) {
	}
}

func (m MetaHandler) SetTitle(title string) MetaHandler {
	return func(context *gin.Context) {
		m(context)
		context.Set("title", title)
	}
}

func (m MetaHandler) ToHFunc() gin.HandlerFunc {
	return gin.HandlerFunc(m)
}

func (m MetaHandler) SetType(typeS string) MetaHandler {
	return func(context *gin.Context) {
		m(context)
		context.Set("type", typeS)
	}
}

func (m MetaHandler) Set(key, value string) MetaHandler {
	return func(context *gin.Context) {
		m(context)
		context.Set(key, value)
	}
}
