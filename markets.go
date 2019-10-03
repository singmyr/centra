package centra

// Market contains the market data.
type Market struct {
	Market    string   `json:"market"`
	Name      string   `json:"name"`
	Countries []string `json:"countries"`
	Products  []string `json:"products"`
}

// GetMarkets gets all the markets.
func (c *Client) GetMarkets() (m map[string]Market, e error) {
	_, e = c.get("/markets", &m)
	return
}

// GetMarket gets a specific market.
func (c *Client) GetMarket(market string) (m Market, e error) {
	_, e = c.get("/markets/"+market, &m)
	return
}
