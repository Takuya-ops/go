package sample

import (
	"net/http"
	"regexp"
)

// フレームワークを使わずにルーティング実装
func (rt *router) GET(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodGet,
		matcher: regexp.MustCompile(prefix),
		fnc:     fnc,
	})
}
