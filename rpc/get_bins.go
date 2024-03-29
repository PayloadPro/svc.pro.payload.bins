package rpc

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/PayloadPro/svc.pro.payload.bins/deps"
	"github.com/PayloadPro/svc.pro.payload.bins/models"
)

// GetBins is a func which takes the incoming request for the bins and returns bins
type GetBins func(context.Context, *http.Request) ([]*models.Bin, int, error)

// NewGetBins is the concrete func for GetBins
func NewGetBins(services *deps.Services, config *deps.Config) GetBins {
	return func(ctx context.Context, r *http.Request) ([]*models.Bin, int, error) {

		var bins = make([]*models.Bin, 0)
		var err error

		if bins, err = services.Bin.GetBins(); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return bins, http.StatusOK, nil
	}
}
