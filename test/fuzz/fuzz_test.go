// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzz

import (
	"encoding/xml"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/fedwire20022/gen/fedwirefunds_incoming"
	"github.com/moov-io/fedwire20022/gen/fedwirefunds_outgoing"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/fedwire20022/pkg/fedwiretest"
)

func FuzzFedWireXML(f *testing.F) {
	populateCorpus(f)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}
		data := []byte(contents)

		var incoming fedwirefunds_incoming.FedwireFundsIncoming
		if err := xml.Unmarshal(data, &incoming); err == nil {
			_, _ = xml.Marshal(&incoming)
		}

		var outgoing fedwirefunds_outgoing.FedwireFundsOutgoing
		if err := xml.Unmarshal(data, &outgoing); err == nil {
			_, _ = xml.Marshal(&outgoing)
		}

		if out, err := fedwiretest.FlipMessageDirection(data); err == nil && len(out) > 0 {
			_, _ = fedwiretest.FlipMessageDirection(out)
		}
	})
}

func FuzzFedWireTypes(f *testing.F) {
	f.Add("2019-03-21")
	f.Add("2019-03-21T10:36:19-04:00")
	f.Add("10:36:19")
	f.Add("0")
	f.Add("1234.56")
	f.Add("")
	f.Add("not-a-date")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > 256 {
			t.Skip()
		}
		var d fedwire.ISODate
		_ = d.UnmarshalText([]byte(s))
		_, _ = d.MarshalText()
		_ = d.Validate()

		var dt fedwire.ISODateTime
		// May or may not exist — use ISOTime as well
		_ = dt.UnmarshalText([]byte(s))

		var tm fedwire.ISOTime
		_ = tm.UnmarshalText([]byte(s))
		_, _ = tm.MarshalText()
		_ = tm.Validate()

		_ = fedwire.ValidBusinessMessageID(s)
		_ = fedwire.ValidatePattern(s, `^[A-Za-z0-9]+$`)
		_ = fedwire.ValidateLength(s, 35)
		_ = fedwire.ValidateMinLength(s, 1)
		_ = fedwire.ValidateMaxLength(s, 35)
	})
}

func populateCorpus(f *testing.F) {
	f.Helper()

	f.Add("")
	f.Add("<FedwireFundsIncoming></FedwireFundsIncoming>")

	roots := []string{
		filepath.Join("..", "..", "testdata"),
		filepath.Join("..", "..", "data"),
	}
	for _, root := range roots {
		_ = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil || info == nil || info.IsDir() {
				return nil
			}
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".xml" || ext == ".txt" {
				bs, err := os.ReadFile(path)
				if err != nil || len(bs) > 512*1024 {
					return nil
				}
				f.Add(string(bs))
			}
			return nil
		})
	}
}
