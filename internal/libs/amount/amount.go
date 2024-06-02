package amount

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidAmount = errors.New("le montant invalide")

func ConvertToCents(amount float64) (int, error) {
	value := strconv.FormatFloat(amount, 'f', -1, 64)
	if r, err := regexp.MatchString(`^\d+(\.\d{0,2}[0]*)?$`, value); err != nil || !r {
		return 0, ErrInvalidAmount
	}

	parts := strings.Split(fmt.Sprintf("%.2f", amount*100), ".")
	v, err := strconv.Atoi(parts[0])

	if err != nil {
		return 0, err
	}

	return v, nil
}
