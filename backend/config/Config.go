package config

// Number of decimal points for prices
var DecimalPoints int = 4

// Time between coinGecko price fetchings in seconds
var UpdateInterval int = 60

// Symbols to monitor
// - needs to be the same as id on coinGecko
var Symbols = []string{
	"ethereum",
	"bitcoin",
}
