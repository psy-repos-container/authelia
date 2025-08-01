---
title: "authelia crypto certificate rsa generate"
description: "Reference for the authelia crypto certificate rsa generate command."
lead: ""
date: 2024-03-14T06:00:14+11:00
draft: false
images: []
weight: 905
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---

## authelia crypto certificate rsa generate

Generate an RSA private key and certificate

### Synopsis

Generate an RSA private key and certificate.

This subcommand allows generating an RSA private key and certificate.

```
authelia crypto certificate rsa generate [flags]
```

### Examples

```
authelia crypto certificate rsa generate --help
```

### Options

```
  -b, --bits int                        number of RSA bits for the certificate (default 2048)
      --bundles strings                 enables generating bundles options are 'chain' and 'privkey-chain'
      --ca                              create the certificate as a certificate authority certificate
  -n, --common-name string              certificate common name
      --country strings                 certificate country
  -d, --directory string                directory where the generated keys, certificates, etc will be stored
      --duration string                 duration of time the certificate is valid for (default "1y")
      --extended-usage strings          specify the extended usage types of the certificate
      --file.bundle.chain string        name of the file to export the certificate chain PEM bundle to when the --bundles flag includes 'chain' (default "public.chain.pem")
      --file.bundle.priv-chain string   name of the file to export the certificate chain and private key PEM bundle to when the --bundles flag includes 'priv-chain' (default "private.chain.pem")
      --file.ca-certificate string      certificate authority certificate to use when signing this certificate (default "ca.public.crt")
      --file.ca-private-key string      certificate authority private key to use to signing this certificate (default "ca.private.pem")
      --file.certificate string         name of the file to export the certificate data to (default "public.crt")
      --file.extension.legacy string    string to include before the actual extension as a sub-extension on the PKCS#1 and SECG1 legacy formats (default "legacy")
      --file.private-key string         name of the file to export the private key data to (default "private.pem")
  -h, --help                            help for generate
      --legacy                          enables the output of the legacy PKCS#1 and SECG1 formats when enabled
  -l, --locality strings                certificate locality
      --not-after string                latest date and time the certificate is considered valid in various formats
      --not-before string               earliest date and time the certificate is considered valid in various formats (default is now)
  -o, --organization strings            certificate organization (default [Authelia])
      --organizational-unit strings     certificate organizational unit
      --path.ca string                  source directory of the certificate authority files, if not provided the certificate will be self-signed
  -p, --postcode strings                certificate postcode
      --province strings                certificate province
      --sans strings                    subject alternative names
      --signature string                signature algorithm for the certificate (default "SHA256")
  -s, --street-address strings          certificate street address
```

### Options inherited from parent commands

```
  -c, --config strings                        configuration files or directories to load, for more information run 'authelia -h authelia config' (default [configuration.yml])
      --config.experimental.filters strings   list of filters to apply to all configuration files, for more information run 'authelia -h authelia filters'
```

### SEE ALSO

* [authelia crypto certificate rsa](authelia_crypto_certificate_rsa.md)	 - Perform RSA certificate cryptographic operations

