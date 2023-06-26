package keeper

import api "core/x/oracle/api"

func (k Keeper) GetCoinPrice(coin string) (uint64, error) {
	// TODO: each api uses diferent symbols, create special cases for each one
	symbol := coin + "USDT"

	fromBinance, err := api.GetFromBinance(symbol)
	if err == nil {
		return uint64(fromBinance * 1_000_000_000), nil
	}

	fromGecko, err := api.GetFromGecko(coin)
	if err == nil {
		return uint64(fromGecko * 1_000_000_000), nil
	}

	//TODO: Agregate from more apis and control errors choosing the most reliable

	return 0, err
}
