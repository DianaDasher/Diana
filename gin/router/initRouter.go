package router

import (
	"github.com/gin-gonic/gin"
)

type RouteConfigurer func(*gin.Engine)

// 注册中间件
func ApplyMiddlewares(r *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	for _, middleware := range middlewares {
		r.Use(middleware)
	}
	return r
}

// 注册路由
func RegisterRouteConfigs(r *gin.Engine, configs ...RouteConfigurer) *gin.Engine {
	for _, config := range configs {
		config(r)
	}
	return r
}
