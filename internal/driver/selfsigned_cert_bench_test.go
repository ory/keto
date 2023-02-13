// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/ory/x/tlsx"
)

func BenchmarkCertificateGeneration(b *testing.B) {
	cases := []struct {
		name  string
		curve elliptic.Curve
	}{
		{"P256", elliptic.P256()},
		{"P224", elliptic.P224()},
		{"P384", elliptic.P384()},
		{"P521", elliptic.P521()},
	}

	for _, tc := range cases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				key, err := ecdsa.GenerateKey(tc.curve, rand.Reader)
				if err != nil {
					b.Fatalf("could not create key: %v", err)
				}
				if _, err = tlsx.CreateSelfSignedTLSCertificate(key); err != nil {
					b.Fatalf("could not create TLS certificate: %v", err)
				}
			}
		})
	}
}
