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
