package models

import (
	"net/http"
	"strings"
	"testing"

	. "github.com/franela/goblin"
)

func TestBinModel(t *testing.T) {

	g := Goblin(t)

	g.Describe("Bins are handled correctly", func() {

		g.It("NewBin() returns a correct bin", func() {

			e := Bin{
				Name: "I'm a bin",
			}

			data := []byte(`{"data":{"attributes":{"name":"I'm a bin"}}}`)

			r, _ := http.NewRequest("GET", "", strings.NewReader(string(data)))

			bin, _ := NewBin(r)
			g.Assert(bin.Name).Equal(e.Name)

		})

		g.It("Bin not found returns correct error", func() {
			nf := ErrBinNotFound
			g.Assert(nf.Error()).Equal("Bin could not be found")
		})

	})

}
