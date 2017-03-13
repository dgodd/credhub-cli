package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Secret", func() {
	Describe("Terminal", func() {
		It("renders string secrets", func() {
			stringSecret := Secret{
				SecretBody: SecretBody{
					Name:             "stringSecret",
					SecretType:       "value",
					Value:            "my-value",
					VersionCreatedAt: "2016-01-01T12:00:00Z",
				},
			}

			Expect(stringSecret.Terminal()).To(Equal("" +
				"Type:          value\n" +
				"Name:          stringSecret\n" +
				"Value:         my-value\n" +
				"Updated:       2016-01-01T12:00:00Z"))
		})

		It("renders ssh secrets", func() {
			ssh := RsaSsh{PublicKey: "my-pub", PrivateKey: "my-priv"}
			sshSecret := Secret{
				SecretBody: SecretBody{
					Name:             "sshSecret",
					SecretType:       "ssh",
					Value:            ssh,
					VersionCreatedAt: "2016-01-01T12:00:00Z",
				},
			}

			Expect(sshSecret.Terminal()).To(Equal("" +
				"Type:          ssh\n" +
				"Name:          sshSecret\n" +
				"Public Key:    my-pub\n" +
				"Private Key:   my-priv\n" +
				"Updated:       2016-01-01T12:00:00Z"))
		})

		It("renders rsa secrets", func() {
			rsa := RsaSsh{PublicKey: "my-pub", PrivateKey: "my-priv"}
			sshSecret := Secret{
				SecretBody: SecretBody{
					Name:             "rsaSecret",
					SecretType:       "rsa",
					Value:            rsa,
					VersionCreatedAt: "2016-01-01T12:00:00Z",
				},
			}

			Expect(sshSecret.Terminal()).To(Equal("" +
				"Type:          rsa\n" +
				"Name:          rsaSecret\n" +
				"Public Key:    my-pub\n" +
				"Private Key:   my-priv\n" +
				"Updated:       2016-01-01T12:00:00Z"))
		})

		Describe("renders certificate secrets", func() {
			It("when fields have non-nil values", func() {
				certificate := Certificate{Ca: "my-ca", Certificate: "my-cert", PrivateKey: "my-priv"}
				certificateSecret := Secret{
					SecretBody: SecretBody{
						Name:             "nonNulledSecret",
						SecretType:       "certificate",
						Value:            certificate,
						VersionCreatedAt: "2016-01-01T12:00:00Z",
					},
				}

				Expect(certificateSecret.Terminal()).To(Equal("" +
					"Type:          certificate\n" +
					"Name:          nonNulledSecret\n" +
					"Ca:            my-ca\n" +
					"Certificate:   my-cert\n" +
					"Private Key:   my-priv\n" +
					"Updated:       2016-01-01T12:00:00Z"))
			})

			It("when some fields have nil values", func() {
				certificate := Certificate{Ca: "my-ca", Certificate: "", PrivateKey: "my-priv"}
				certificateSecret := Secret{
					SecretBody: SecretBody{
						Name:             "nonNulledSecret",
						SecretType:       "certificate",
						Value:            certificate,
						VersionCreatedAt: "2016-01-01T12:00:00Z",
					},
				}

				Expect(certificateSecret.Terminal()).To(Equal("" +
					"Type:          certificate\n" +
					"Name:          nonNulledSecret\n" +
					"Ca:            my-ca\n" +
					"Private Key:   my-priv\n" +
					"Updated:       2016-01-01T12:00:00Z"))
			})

			It("when fields all have nil values", func() {
				certificateSecret := Secret{
					SecretBody: SecretBody{
						Name:             "nulledSecret",
						SecretType:       "certificate",
						Value:            Certificate{},
						VersionCreatedAt: "2016-01-01T12:00:00Z",
					},
				}

				Expect(certificateSecret.Terminal()).To(Equal("" +
					"Type:          certificate\n" +
					"Name:          nulledSecret\n" +
					"Updated:       2016-01-01T12:00:00Z"))
			})
		})
	})

	Describe("JSON", func() {
		It("renders string secrets", func() {
			stringSecret := Secret{
				SecretBody: SecretBody{
					Name:             "stringSecret",
					SecretType:       "value",
					Value:            "my-value",
					VersionCreatedAt: "2016-01-01T12:00:00Z",
				},
			}

			Expect(stringSecret.Json()).To(MatchJSON(`{
				"type": "value",
				"name": "stringSecret",
				"value": "my-value",
				"version_created_at": "2016-01-01T12:00:00Z"
			}`))
		})

		It("renders ssh secrets", func() {
			ssh := RsaSsh{PublicKey: "my-pub", PrivateKey: "my-priv"}
			sshSecret := Secret{
				SecretBody: SecretBody{
					Name:             "sshSecret",
					SecretType:       "ssh",
					Value:            ssh,
					VersionCreatedAt: "2016-01-01T12:00:00Z",
				},
			}

			Expect(sshSecret.Json()).To(MatchJSON(`{
				"type": "ssh",
				"name": "sshSecret",
				"version_created_at": "2016-01-01T12:00:00Z",
				"value": {
					"public_key": "my-pub",
					"private_key": "my-priv"
				}
			}`))
		})

		It("renders rsa secrets", func() {
			rsa := RsaSsh{PublicKey: "my-pub", PrivateKey: "my-priv"}
			sshSecret := Secret{
				SecretBody: SecretBody{
					Name:             "rsaSecret",
					SecretType:       "rsa",
					Value:            rsa,
					VersionCreatedAt: "2016-01-01T12:00:00Z",
				},
			}

			Expect(sshSecret.Json()).To(MatchJSON(`{
				"type": "rsa",
				"name": "rsaSecret",
				"version_created_at": "2016-01-01T12:00:00Z",
				"value": {
					"public_key": "my-pub",
					"private_key": "my-priv"
				}
			}`))
		})

		Describe("renders certificate secrets", func() {
			It("when fields have non-nil values", func() {
				certificate := Certificate{Ca: "my-ca", Certificate: "my-cert", PrivateKey: "my-priv"}
				certificateSecret := Secret{
					SecretBody: SecretBody{
						Name:             "nonNulledSecret",
						SecretType:       "certificate",
						Value:            certificate,
						VersionCreatedAt: "2016-01-01T12:00:00Z",
					},
				}

				Expect(certificateSecret.Json()).To(MatchJSON(`{
					"type": "certificate",
					"name": "nonNulledSecret",
					"version_created_at": "2016-01-01T12:00:00Z",
					"value": {
						"ca": "my-ca",
						"certificate": "my-cert",
						"private_key": "my-priv"
					}
				}`))
			})

			It("when some fields have nil values", func() {
				certificate := Certificate{Ca: "my-ca", Certificate: "", PrivateKey: "my-priv"}
				certificateSecret := Secret{
					SecretBody: SecretBody{
						Name:             "nonNulledSecret",
						SecretType:       "certificate",
						Value:            certificate,
						VersionCreatedAt: "2016-01-01T12:00:00Z",
					},
				}

				Expect(certificateSecret.Json()).To(MatchJSON(`{
					"type": "certificate",
					"name": "nonNulledSecret",
					"version_created_at": "2016-01-01T12:00:00Z",
					"value": {
						"ca": "my-ca",
						"private_key": "my-priv"
					}
				}`))
			})

			It("when fields all have nil values", func() {
				certificateSecret := Secret{
					SecretBody: SecretBody{
						Name:             "nulledSecret",
						SecretType:       "certificate",
						Value:            Certificate{},
						VersionCreatedAt: "2016-01-01T12:00:00Z",
					},
				}

				Expect(certificateSecret.Json()).To(MatchJSON(`{
					"type": "certificate",
					"name": "nulledSecret",
					"value": {},
					"version_created_at": "2016-01-01T12:00:00Z"
				}`))
			})
		})
	})
})
