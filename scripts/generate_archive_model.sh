#!/bin/bash
set -e

moovio_xsd2go convert \
   xsd/xmldsig-core-schema.xsd \
   github.com/moov-io/fedwire20022 \
   gen \
   --template-name=internal/templates/fedwire20022/xmldsig.tgo \
   --xmlns-override="http://www.w3.org/2000/09/xmldsig#=xmldsig"

# Ensure ./xsd exists
if [ ! -d "./xsd" ]; then
  echo "Directory ./xsd does not exist."
  exit 1
fi

mkdir -p ./gen

find ./xsd -mindepth 2 -type f -name "*.xsd" | while read -r filepath; do
  relative_path="${filepath#./xsd/}"
  folder_path=$(dirname "$relative_path")
  versionname=$(basename "$filepath" .xsd)
  ns_dots=${versionname//./_}
  mkdir -p "./gen/$folder_path"
  mkdir -p "./gen/$folder_path/$ns_dots"

  msg="${folder_path}_${ns_dots}"
  genFolder="gen/$folder_path"
  moovio_xsd2go convert \
                  "$filepath" github.com/moov-io/fedwire20022 ${genFolder} \
                  --template-name=internal/templates/fedwire20022/model.tgo \
                  --template-name=internal/templates/fedwire20022/write.tgo \
                  --template-name=internal/templates/fedwire20022/validate.tgo \
                  --xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.02=head_001_001_02" \
                  --xmlns-override="urn:iso:std:iso:20022:tech:xsd:${ns_dots}=${msg}" \
                  --xmlns-override="urn:fedwire=Message_${msg}"
done

# run go fmt and goimports for every generated file
find ./gen -type f -name "*.go" | while read -r file; do
    gofmt -w "$file"
    goimports -w "$file"
    echo "Formatted: $file"
done