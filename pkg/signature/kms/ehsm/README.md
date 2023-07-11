# Start on the EHSM source
```
./run_with_single.sh 
```

# Cosign invokes the EHSM service process

After configuring the Go environment, add it to the import of the main.go file of cosign.
```
_ "github.com/syan10/sigstore/pkg/signature/kms/ehsm"
```
Update the GitHub repository and rebuild the cosign.
```
go mod tidy
go install ./cmd/cosign
```
# Start EHSM using cosign

Set environment variables EHSM_APPID、EHSM_APIKEY、EHSM_ADDR，for example

```
export EHSM_APPID=49b84175-1a82-419d-be35-fba6bc85ee69
export EHSM_APIKEY=B8FM3GVex6FAP45yeDuwSvuShjtn3gEJ
export EHSM_ADDR=https://10.112.241.110:9002
```
## Generate keys 
To generate keys using a KMS provider,Generate cosign.pub to store the public key, and some key files to store the keyid.
```
 cosign generate-key-pair --kms ehsm://<some key>
```
## Sign and Verify
For example
```
IMAGE DIGEST = lisiteng123/ubuntu@sha256:dfa84ec91daba0fecef93302a6f5754566a0249e90626dc7d289746ae412af75
```
Use keys for signing and authentication.
```
cosign sign --key ehsm://<some key> IMAGE DIGEST
cosign verify --key ehsm://<some key> IMAGE DIGEST
```

Reference https://docs.sigstore.dev/cosign/kms_support/