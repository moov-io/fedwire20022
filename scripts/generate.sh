#!/bin/bash
set -e

moovio_xsd2go convert \
   xsd/xmldsig-core-schema.xsd \
   github.com/moov-io/fedwire20022 \
   gen \
   --template-name=internal/templates/fedwire20022/xmldsig.tgo \
   --xmlns-override="http://www.w3.org/2000/09/xmldsig#=xmldsig"


files=($(find ./xsd -name "*iso15.xsd" | sort -u))
for file in "${files[@]}"
do
    # file = "./xsd/AccountActivityDetailsReport_camt_052_001_08_20241122_1718_iso15.xsd"
    ns_under=$(echo "$file" | grep -oE '[a-z]{4}_[0-9]{3}_[0-9]{3}_[0-9]{2}')
    ns_dots=$(echo "$ns_under" | tr '_' '.')
    msg=$(basename "$file" .xsd | grep -oE '[a-zA-Z]{1,}_[a-z]{4}_[0-9]{3}_[0-9]{3}_[0-9]{2}')

    moovio_xsd2go convert \
                  "$file" github.com/moov-io/fedwire20022 gen \
                  --template-name=internal/templates/fedwire20022/model.tgo \
                  --template-name=internal/templates/fedwire20022/write.tgo \
                  --template-name=internal/templates/fedwire20022/validate.tgo \
                  --xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.02=head_001_001_02" \
                  --xmlns-override="urn:iso:std:iso:20022:tech:xsd:${ns_dots}=${msg}" \
                  --xmlns-override="urn:fedwire=Message_${msg}"
done

# run go fmt and goimports for every generated file
files=($(find ./gen -name '*.go'))
for file in "${files[@]}"
do
    gofmt -w $file
    goimports -w $file
done
