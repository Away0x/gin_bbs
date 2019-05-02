package routes

import (
	"gin_bbs/pkg/ginutils/router"

	"gin_bbs/app/controllers/page"
)

func registerWeb(r *router.MyRoute) {
	r.Register("GET", "root", "/", page.Root)
}
