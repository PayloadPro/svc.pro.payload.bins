package rpc

import (
	"net/http"
	"reflect"

	"golang.org/x/net/context"

	"github.com/PayloadPro/svc.pro.payload.bins/deps"
	"github.com/PayloadPro/svc.pro.payload.bins/models"
	"github.com/gorilla/mux"
)

// GetBin is a func which takes the incoming request for the bins and returns bins
type GetBin func(context.Context, *http.Request) (*models.Bin, int, error)

// NewGetBin is the concrete func for GetBins
func NewGetBin(services *deps.Services, config *deps.Config) GetBin {
	return func(ctx context.Context, r *http.Request) (*models.Bin, int, error) {

		var bin = &models.Bin{}

		vars := mux.Vars(r)
		bin, err := services.Bin.GetByID(vars["id"])

		if err != nil {
			if reflect.DeepEqual(err, models.ErrBinNotFound) {
				return nil, http.StatusNotFound, err
			}
			return nil, http.StatusInternalServerError, err
		}

		return bin, http.StatusOK, nil
	}
}
