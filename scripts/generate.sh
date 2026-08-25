#!/bin/bash
set -e

moovio_xsd2go convert \
   xsd/xmldsig-core-schema.xsd \
   github.com/moov-io/fedwire20022 \
   gen \
   --template-name=internal/templates/fedwire20022/xmldsig.tgo \
   --xmlns-override="http://www.w3.org/2000/09/xmldsig#=xmldsig"

# Generate Go models
files=($(find ./xsd -name "fedwirefunds-*.xsd" | sort -u))
for file in "${files[@]}"
do
    moovio_xsd2go convert \
                  "$file" github.com/moov-io/fedwire20022 gen \
                  --template-name=internal/templates/fedwire20022/model.tgo \
                  --template-name=internal/templates/fedwire20022/write.tgo \
                  --template-name=internal/templates/fedwire20022/validate.tgo
done

# run go fmt and goimports for every generated file
files=($(find ./gen -name '*.go'))
for file in "${files[@]}"
do
    gofmt -w $file
    goimports -w $file
done

# xsd2go camel-cases element names into XML tags. Fed's outgoing XSD keeps
# underscores on the account-balance report wrappers, and those names appear
# on the wire. Restore them after generate so unmarshal matches Fed.
perl -pi -e 's/xml:"FedwireFundsAccountBalanceReportMaster/xml:"FedwireFundsAccountBalanceReport_Master/g; s/xml:"FedwireFundsAccountBalanceReportSelf/xml:"FedwireFundsAccountBalanceReport_Self/g' \
	gen/fedwirefunds_outgoing/model.go
