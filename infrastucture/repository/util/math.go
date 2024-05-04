package util

import (
	"math"

	"github.com/shopspring/decimal"
)

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RoundAndConvertDecimalToFloat64(decimalNumber decimal.Decimal, precision int) float64 {
	decimalNumberRound := decimal.Decimal.Round(decimalNumber, int32(precision))
	floatNumber, _ := decimalNumberRound.Float64()
	return floatNumber
}

func ConvertDecimalToFloat64(decimalNumber decimal.Decimal) float64 {
	floatNumber, _ := decimalNumber.Float64()
	return floatNumber
}

func CheckDocumentWaitingForApproval(numbers ...int) bool {
    for _, num := range numbers {
        if num > 0 {
            return true
        }
    }
    return false
}