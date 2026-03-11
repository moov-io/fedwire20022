package fedwire

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

// BusinessMessageID creates a FedWire Business Message Identifier from a base64-encoded UUID v4 and a prefix (up to 3 characters)
func BusinessMessageID(prefix string) string {
	random := strings.ReplaceAll(uuid.New().String(), "-", "")
	return strings.TrimSpace(fmt.Sprintf("%3.3s", prefix) + random)
}

func ValidBusinessMessageID(id string) error {
	n := utf8.RuneCountInString(id)
	if n < 1 || n > 35 {
		return fmt.Errorf("business message id must be 1-35 characters long - was %d", n)
	}

	// The CPMI permissible character set is composed of lower-case characters a-z, upper-case
	// characters A-Z, numeric characters 0-9, complemented with the following additional
	// limited set of special characters: / - ? : ( ) . , ' + ! # & % * = ^ _ ` { | } ~ " ; @ [ \ ] $ > < and space
	for _, r := range id {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			continue
		}
		switch r {
		case '/', '-', '?', ':', '(', ')', '.', ',', '\'', '+', '!', '#', '&', '%', '*', '=',
			'^', '_', '`', '{', '|', '}', '~', '"', ';', '@', '[', '\\', ']', '$', '>', '<', ' ':
			continue
		default:
			return fmt.Errorf("unexpected %v in business message id", string(r))
		}
	}
	return nil
}
