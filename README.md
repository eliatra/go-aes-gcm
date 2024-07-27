# Go AES GCM

A small go library and command line tool (CLI) to encrypt/decrypt data and files with symmetric encryption algorithm AES in GCM mode.

## Eliatra

- https://github.com/eliatra/go-aes-gcm
- https://eliatra.com
- https://www.linkedin.com/company/7116719

Contact: sales@eliatra.com

## Use CLI

The CLI reads the secret AES key from an environment variable named `AES_SECRET_KEY`.
This needs to be a string with 16 or 32 bytes (AES-128 or AES-256).
 
Make sure you have a strong key which is truly random. 

If you only have a human style password set `AES_SECRET_PASSWORD` environment variable.

Never use a password directly as a key. Instead, use Password-Based Keys (PBK), where the key is generated from the password via a key derivation function (KDF) with key stretching, such as PBKDF2. This is what we do with the content of `AES_SECRET_PASSWORD`.

The salt for PBKDF2 is a 16-byte hardcoded value, and we perform 210,000 iterations using HMAC-SHA512 as the PRF. If a hardcoded salt is not applicable, create a truly random key and use the `AES_SECRET_KEY` environment variable.

Encrypt/decrypt files

```
./go-aes-gcm encrypt|decrypt <source-file> <target-file>

./go-aes-gcm encrypt <plain-text-file> <cipher-text-file>
./go-aes-gcm decrypt <cipher-text-file> <plain-text-file>
```

The `<target-file>` is created (with perm 0600) if it does not exist.
If it does exist, it is truncated first (i.e., overwritten).

Encrypt/decrypt stdin

```
echo "input" | ./go-aes-gcm encrypt|decrypt [aad]

echo "plaintext" | ./go-aes-gcm encrypt
echo "base64-ciphertext" | ./go-aes-gcm decrypt
echo "plaintext" | ./go-aes-gcm encrypt | ./go-aes-gcm decrypt
```

Plaintext will be encrypted and printed to stdout as base64-encoded bytes.
Base64-encoded ciphertext will be decrypted and printed to stdout as a string.

[AAD (Additional Authenticated Data)](https://crypto.stackexchange.com/questions/35727/does-aad-make-gcm-encryption-more-secure) is optional.

Additional Authenticated Data (AAD) is a portion of the input that is authenticated but not encrypted, ensuring the integrity of the data while it remains in plaintext. This allows the recipient to verify that the data has not been altered without the need to decrypt it.

## Use as Library

See `main.go`

## Note

AES GCM has some limitations:

- A nonce must never ever be reused with the same key. In this library, we use random nonces. In some use cases, atomic counter nonces may be a better choice. 
- Usage of more that 2^32 random nonces with the same key is insecure.
- Do not encrypt more than 350 GB of total input data with the same key.

## License

```
Copyright 2024 Eliatra

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```