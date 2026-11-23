[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![Go Reference](https://pkg.go.dev/badge/github.com/moov-io/fedwire20022.svg)](https://pkg.go.dev/github.com/moov-io/fedwire20022)
[![Build Status](https://github.com/moov-io/fedwire20022/workflows/Go/badge.svg)](https://github.com/moov-io/fedwire20022/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/fedwire20022)](https://goreportcard.com/report/github.com/moov-io/fedwire20022)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/fedwire20022?label=project%20size)](https://github.com/moov-io/fedwire20022)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/fedwire20022/master/LICENSE)
[![Slack Channel](https://img.shields.io/badge/slack-moov--io-e01563?logo=slack)](https://slack.moov.io/)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/fedwire20022)](https://github.com/moov-io/fedwire20022)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)


## moov-io/fedwire20022

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

This repository contains a subset of ISO 20022 messages for FedWire payments. ISO 20022 is a standard for electronic data interchange between financial institutions. It describes a metadata repository containing descriptions of messages and business processes, and a maintenance process for the repository content. The standard covers financial information transferred between financial institutions that includes payment transactions, securities trading and settlement information, credit and debit card transactions, and other financial information.

## Project Status

Go code is generated with [`moov-io/xsd2go`](https://github.com/moov-io/xsd2go) inside the `gen/` folder from the XSDs in `xsd/`. Please star the project if you are interested in its progress. Feedback on this early version of ISO 20022 is appreciated and vital to its success. Please let us know if you encounter any bugs/unclear documentation or have feature suggestions by opening up an issue. Thanks!

This project is a Go library only. There is no CLI, HTTP server, or Docker image. API documentation is on [pkg.go.dev](https://pkg.go.dev/github.com/moov-io/fedwire20022).

## Naming Conventions

The XSDs define some naming that can be confusing.

| Go Package                  | Description                                                        |
|-----------------------------|--------------------------------------------------------------------|
| `gen/fedwirefunds_incoming` | Messages sent INTO the Fed Wire service - from a bank or vendor.   |
| `gen/fedwirefunds_outgoing` | Messages coming INTO a bank or vendor - from the Fed Wire service. |

## Supported Messages

Each FedWire message is wrapped in an application header (`head.001.001.03`) and a `Document`. The wrapper types are in `gen/fedwirefunds_incoming` and `gen/fedwirefunds_outgoing`.

| Message | Go Package | Incoming | Outgoing |
|---------|------------|----------|----------|
| `admi.002.001.01` Message Reject | `gen/admi_002_001_01` | Yes | Yes |
| `admi.004.001.02` Connection Check / Broadcast | `gen/admi_004_001_02` | Yes | Yes |
| `admi.006.001.01` Retrieval Request | `gen/admi_006_001_01` | Yes | |
| `admi.007.001.01` Acknowledgement | `gen/admi_007_001_01` | | Yes |
| `admi.011.001.01` System Response | `gen/admi_011_001_01` | | Yes |
| `camt.029.001.09` Return Request Response | `gen/camt_029_001_09` | Yes | Yes |
| `camt.052.001.08` Account Balance, Activity and Endpoint Reports | `gen/camt_052_001_08` | | Yes |
| `camt.056.001.08` Return Request | `gen/camt_056_001_08` | Yes | Yes |
| `camt.060.001.05` Account Reporting Request | `gen/camt_060_001_05` | Yes | |
| `camt.110.001.01` Investigation Request | `gen/camt_110_001_01` | Yes | Yes |
| `camt.111.001.01` Investigation Response | `gen/camt_111_001_01` | Yes | Yes |
| `head.001.001.03` Business Application Header | `gen/head_001_001_03` | Yes | Yes |
| `pacs.002.001.10` Payment Status | `gen/pacs_002_001_10` | | Yes |
| `pacs.004.001.10` Payment Return | `gen/pacs_004_001_10` | Yes | Yes |
| `pacs.008.001.08` Customer Credit Transfer | `gen/pacs_008_001_08` | Yes | Yes |
| `pacs.009.001.08` Financial Institution Credit Transfer | `gen/pacs_009_001_08` | Yes | Yes |
| `pacs.028.001.03` Payment Status Request | `gen/pacs_028_001_03` | Yes | Yes |
| `pain.013.001.07` Drawdown Request | `gen/pain_013_001_07` | Yes | Yes |
| `pain.014.001.07` Drawdown Response | `gen/pain_014_001_07` | Yes | Yes |

## Usage

### Install

Requires Go 1.26 or newer.

```
go get github.com/moov-io/fedwire20022@latest
go mod tidy
```

### Go Library

This project offers Go structures which can read and write XML encoded ISO 20022 messages. Every generated type has a `Validate() error` method which checks values against the restrictions (patterns, lengths, enumerations and numeric ranges) in the XSD. `Validate()` does not check FedWire business rules.

Read a message and validate it:

```go
package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/moov-io/fedwire20022/gen/fedwirefunds_incoming"
)

func main() {
	bs, err := os.ReadFile("testdata/PACS008_CustomerCreditTransfer.txt")
	if err != nil {
		log.Fatal(err)
	}

	var msg fedwirefunds_incoming.FedwireFundsIncoming
	if err := xml.Unmarshal(bs, &msg); err != nil {
		log.Fatal(err)
	}
	if err := msg.Validate(); err != nil {
		log.Fatal(err)
	}

	cct := msg.FedwireFundsIncomingMessage.FedwireFundsCustomerCreditTransfer
	fmt.Println(cct.AppHdr.MsgDefIdr)                        // pacs.008.001.08
	fmt.Println(cct.Document.FIToFICstmrCdtTrf.GrpHdr.MsgId) // 20240228B1QDRCQR000001
}
```

To write a message, build the wrapper struct, call `Validate()`, then call `xml.Marshal`. See [`example_test.go`](example_test.go) for a complete `pacs.009` example. Sample messages are in [`testdata/`](testdata/).

The `pkg/fedwire` package has helpers for FedWire values:

- `fedwire.InputMessageAccountabilityData(cycleDate, endpointID, sequenceNumber)` creates an IMAD, and `fedwire.ValidIMAD` checks one.
- `fedwire.BusinessMessageID(prefix)` creates a Business Message Identifier, and `fedwire.ValidBusinessMessageID` checks one.
- `fedwire.ISODate`, `fedwire.ISODateTime` and `fedwire.ISOTime` read and write ISO 20022 dates and times.
- `fedwire.Amount` writes amounts with two decimal places. A value such as `1.234` is written as `1.23`.

### XML Namespaces

Decoding accepts messages with or without namespace prefixes (for example `<urn:FedwireFundsIncoming xmlns:urn="...">`), because Go matches elements by their local name. Encoding writes a default namespace on the wrapper, `AppHdr` and `Document` elements with no prefixes, for example `<FedwireFundsIncoming xmlns="urn:fedwirefunds:incoming:v001">`. Go has several [known limitations around XML namespace prefixes](https://github.com/golang/go/issues/13400).

### Generating Code

The code in `gen/` is generated from the XSDs in `xsd/` and the templates in `internal/templates/`. Do not edit the files in `gen/` by hand. To generate the code again:

```
make install-generator   # installs moovio_xsd2go and goimports into $(go env GOPATH)/bin
export PATH=$PATH:$(go env GOPATH)/bin
make generate
```

## Getting Help

 channel | info
 ------- | -------
Twitter [@moov](https://twitter.com/moov) | You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/fedwire20022/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#fedwire`) to have an interactive discussion about the development of the project.

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows
- Go 1.26 or newer

Tests run on Linux and macOS. The Windows CI job skips `make check`.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
