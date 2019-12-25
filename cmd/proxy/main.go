package main

import "github.com/chubaka358/intern/pkg/proxy"

func main() {
	dao := proxy.DAOProxy{}
	dao.Add("trololo")
	dao.Delete(1)
	dao.Update(1, "tralala")
	dao.Get(1)
}
