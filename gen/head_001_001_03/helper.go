package head_001_001_03

import (
	"github.com/moov-io/fedwire20022/gen/xmldsig"
)

type SignatureEnvelope struct {
	Signature *xmldsig.Signature `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
}
