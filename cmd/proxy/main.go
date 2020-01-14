package main

import "github.com/jayhrat/intern/pkg/proxy"

func main() {
	dao := proxy.NewDAOProxy()
	dao.Add("trololo")
	dao.Delete(1)
	dao.Update(1, "tralala")
	dao.Get(1)
}
