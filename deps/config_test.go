package deps

import (
	"os"
	"testing"

	"github.com/PayloadPro/svc.pro.payload.bins/configs"
	. "github.com/franela/goblin"
)

func TestCockroachSetup(t *testing.T) {

	g := Goblin(t)

	proto := "postgresql"
	host := "localhost"
	port := "5432"
	user := "dbuser"
	pass := "secret"
	database := "payloadpro"

	os.Setenv("DB_PROTO", proto)
	os.Setenv("DB_HOST", host)
	os.Setenv("DB_PORT", port)
	os.Setenv("DB_USER", user)
	os.Setenv("DB_PASS", pass)
	os.Setenv("DB_DATABASE", database)

	c := &Config{
		DB: &configs.DBConfig{},
	}
	c.Setup()

	g.Describe("Sets up App and DB configs correctly", func() {

		g.It("Contains the correct proto", func() {
			g.Assert(c.DB.Proto).Equal(proto)
		})

		g.It("Contains the correct host", func() {
			g.Assert(c.DB.Host).Equal(host)
		})

		g.It("Contains the correct port", func() {
			g.Assert(c.DB.Port).Equal(port)
		})

		g.It("Contains the correct user", func() {
			g.Assert(c.DB.User).Equal(user)
		})

		g.It("Contains the correct pass", func() {
			g.Assert(c.DB.Pass).Equal(pass)
		})

		g.It("Contains the correct database", func() {
			g.Assert(c.DB.Database).Equal(database)
		})

	})
}
