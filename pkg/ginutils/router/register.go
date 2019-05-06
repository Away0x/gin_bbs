package router

import (
	"github.com/gin-gonic/gin"
)

// MyRoute -
type MyRoute struct {
	Router gin.IRouter
}

// Middleware 注册中间件
func (r *MyRoute) Middleware(middlewares ...gin.HandlerFunc) *MyRoute {
	return &MyRoute{
		Router: r.Router.Group("/", middlewares...),
	}
}

// Group 注册路由组
func (r *MyRoute) Group(relativePath string, handlers ...gin.HandlerFunc) *MyRoute {
	return &MyRoute{
		Router: r.Router.Group(relativePath, handlers...),
	}
}

// Register 注册路由
func (r *MyRoute) Register(method string, name string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	Name(r.Router, name, method, relativePath)
	return r.Router.Handle(method, relativePath, handlers...)
}
