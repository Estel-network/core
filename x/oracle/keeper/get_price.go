package keeper

import api "core/x/oracle/api"

func (k Keeper) GetCoinPrice(coin string) (float64, error) {
	// TODO: each api uses diferent symbols, create special cases for each one

	fromBinance, err := api.GetFromBinance(coin)
	if err == nil {
		return fromBinance, nil
	}

	fromGecko, err := api.GetFromGecko(coin)
	if err == nil {
		return fromGecko, nil
	}

	//TODO: Agregate from more apis and control errors choosing the most reliable

	return 0, err
}
