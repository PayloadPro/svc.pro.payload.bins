package configs

import (
	"os"
	"testing"

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

	db := &DBConfig{}
	db.Setup()

	g.Describe("Returns OS env vars as values", func() {

		g.It("Contains the correct proto", func() {
			g.Assert(db.Proto).Equal(proto)
		})

		g.It("Contains the correct host", func() {
			g.Assert(db.Host).Equal(host)
		})

		g.It("Contains the correct port", func() {
			g.Assert(db.Port).Equal(port)
		})

		g.It("Contains the correct user", func() {
			g.Assert(db.User).Equal(user)
		})

		g.It("Contains the correct pass", func() {
			g.Assert(db.Pass).Equal(pass)
		})

		g.It("Contains the correct database", func() {
			g.Assert(db.Database).Equal(database)
		})

	})

}
