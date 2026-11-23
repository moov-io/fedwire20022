package fedwire

import (
	"cmp"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/moov-io/base"
)

func AddError(errs *base.ErrorList, fieldName string, err error) {
	if err != nil {
		if el, ok := err.(base.ErrorList); ok {
			for indx := range el {
				errs.Add(el[indx])
			}
		} else {
			errs.Add(fmt.Errorf("%v: %v", fieldName, err))
		}
	}
}

var (
	compiledRegexeLock sync.RWMutex
	compiledRegexes    = make(map[string]*regexp.Regexp)
)

// ValidatePattern checks that the whole value matches an XSD pattern.
// XSD patterns are implicitly anchored at both ends, so the regex is
// compiled as ^(?:regex)$.
func ValidatePattern(value string, regex string) error {
	compiledRegexeLock.RLock()
	rr, exists := compiledRegexes[regex]
	compiledRegexeLock.RUnlock()
	if !exists {
		r, err := regexp.Compile("^(?:" + regex + ")$")
		if err != nil {
			return err
		}

		compiledRegexeLock.Lock()
		compiledRegexes[regex] = r
		compiledRegexeLock.Unlock()

		rr = r
	}
	if !rr.MatchString(value) {
		return fmt.Errorf("%s fails validation with pattern %s", value, regex)
	}
	return nil
}

func ValidateEnumeration(value string, enumeration ...string) error {
	for indx := range enumeration {
		if enumeration[indx] == value {
			return nil
		}
	}
	return fmt.Errorf("%s fails enumeration validation", value)
}

func ValidateLength(value string, length int) error {
	var strLength = utf8.RuneCountInString(value)
	if strLength != length {
		return fmt.Errorf("%s fails validation with length %v, length %v required", value, strLength, length)
	}
	return nil
}

func ValidateMinLength(value string, minLength int) error {
	var strLength = utf8.RuneCountInString(value)
	if strLength < minLength {
		return fmt.Errorf("%s fails validation with length %v >= required minLength %v", value, strLength, minLength)
	}
	return nil
}

func ValidateMaxLength(value string, maxLength int) error {
	var strLength = utf8.RuneCountInString(value)
	if strLength > maxLength {
		return fmt.Errorf("%s fails validation with length %v <= required maxLength %v", value, strLength, maxLength)
	}
	return nil
}

func ValidateMinInclusive[N cmp.Ordered](value N, minValue N) error {
	if value < minValue {
		return fmt.Errorf("%v fails validation as it is >= minInclusive %v", value, minValue)
	}
	return nil
}

func ValidateMaxInclusive[N cmp.Ordered](value N, maxValue N) error {
	if value > maxValue {
		return fmt.Errorf("%v fails validation as it is <= maxInclusive %v", value, maxValue)
	}
	return nil
}

func ValidateMinExclusive[N cmp.Ordered](value N, minValue N) error {
	if value <= minValue {
		return fmt.Errorf("%v fails validation as it is > minExclusive %v", value, minValue)
	}
	return nil
}

func ValidateMaxExclusive[N cmp.Ordered](value N, maxValue N) error {
	if value >= maxValue {
		return fmt.Errorf("%v fails validation as it is < maxExclusive %v", value, maxValue)
	}
	return nil
}

func ValidateFractionDigits(value string, maxValue int) error {
	splt := strings.Split(value, ".")
	if len(splt) > 1 {
		var strLength = utf8.RuneCountInString(splt[1])
		if strLength > maxValue {
			return fmt.Errorf("%s fails validation with length %v <= required fractionDigits %v", value, strLength, maxValue)
		}
	}
	return nil
}

// ValidateFinite returns an error when value is NaN or infinite.
// XSD decimal values must be finite numbers.
func ValidateFinite(value float64) error {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return fmt.Errorf("%v fails validation as it is not a finite number", value)
	}
	return nil
}

// FormatDecimal writes value as a plain decimal number with no exponent.
// Use it with ValidateTotalDigits and ValidateFractionDigits.
func FormatDecimal(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func ValidateTotalDigits(value string, maxValue int) error {
	// The sign and leading zeros are not digits in the XSD totalDigits facet.
	str := strings.TrimLeft(value, "+-")
	str = strings.TrimLeft(str, "0")
	str = strings.ReplaceAll(str, ".", "")
	var strLength = utf8.RuneCountInString(str)
	if strLength > maxValue {
		return fmt.Errorf("%s fails validation with length %v <= required totalDigits %v", value, strLength, maxValue)
	}
	return nil
}
