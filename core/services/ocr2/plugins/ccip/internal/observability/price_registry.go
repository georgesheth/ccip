package observability

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

type ObservedPriceRegistryReader struct {
	ccipdata.PriceRegistryReader
	metric metricDetails
}

func NewPriceRegistryReader(origin ccipdata.PriceRegistryReader, chainID int64, pluginName string) *ObservedPriceRegistryReader {
	return &ObservedPriceRegistryReader{
		PriceRegistryReader: origin,
		metric: metricDetails{
			histogram:  priceRegistryHistogram,
			pluginName: pluginName,
			chainId:    chainID,
		},
	}
}

func (o *ObservedPriceRegistryReader) GetTokenPriceUpdatesCreatedAfter(ctx context.Context, ts time.Time, confs int) ([]ccipdata.Event[ccipdata.TokenPriceUpdate], error) {
	return withObservedContract(o.metric, "GetTokenPriceUpdatesCreatedAfter", func() ([]ccipdata.Event[ccipdata.TokenPriceUpdate], error) {
		return o.PriceRegistryReader.GetTokenPriceUpdatesCreatedAfter(ctx, ts, confs)
	})
}

func (o *ObservedPriceRegistryReader) GetGasPriceUpdatesCreatedAfter(ctx context.Context, chainSelector uint64, ts time.Time, confs int) ([]ccipdata.Event[ccipdata.GasPriceUpdate], error) {
	return withObservedContract(o.metric, "GetGasPriceUpdatesCreatedAfter", func() ([]ccipdata.Event[ccipdata.GasPriceUpdate], error) {
		return o.PriceRegistryReader.GetGasPriceUpdatesCreatedAfter(ctx, chainSelector, ts, confs)
	})
}

func (o *ObservedPriceRegistryReader) GetFeeTokens(ctx context.Context) ([]common.Address, error) {
	return withObservedContract(o.metric, "GetFeeTokens", func() ([]common.Address, error) {
		return o.PriceRegistryReader.GetFeeTokens(ctx)
	})
}

func (o *ObservedPriceRegistryReader) GetTokenPrices(ctx context.Context, wantedTokens []common.Address) ([]ccipdata.TokenPriceUpdate, error) {
	return withObservedContract(o.metric, "GetTokenPrices", func() ([]ccipdata.TokenPriceUpdate, error) {
		return o.PriceRegistryReader.GetTokenPrices(ctx, wantedTokens)
	})
}
