package exchange

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/shopspring/decimal"
	"strings"
)

func printPriceData(price decimal.Decimal, exchange, base, quote string) {
	fmt.Print("\n")
	color.HiGreen("===== %s/%s =====\n", strings.ToUpper(base), strings.ToUpper(quote))
	color.Cyan("exchange:\t%s\n", exchange)
	color.Cyan("price:\t%s\n", price.String())
}
