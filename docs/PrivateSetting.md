# Private Key Management
> We use the SM4 encryption algorithm to manage the private keys of the wallet. The encrypted private key is stored locally and can only be decrypted for transaction operations after the user enters the key.

## 1. Private Key Encryption
### Prerequisites
#### Install Golang
```shell
wget https://golang.org/dl/go1.17.2.linux-amd64.tar.gz
sudo tar -xvf go1.17.2.linux-amd64.tar.gz
sudo mv go /usr/local
echo "export GOROOT=/usr/local/go" >> ~/.profile
echo "export GOPATH=$HOME/go" >> ~/.profile
echo "export PATH=$GOPATH/bin:$GOROOT/bin:$PATH" >> ~/.profile
source ~/.profile

# verify that Go has been installed
go version

```
#### Install Project Dependencies
```shell
go mod tidy

```

### Usage
```shell
cd verifier/pkg/sm4

# Generate private key
KEY=1234567890abcdef IV=123456789asdfghb PRIVATE_DATA=[Your Private Key] go test -v -run TestSm4Decrypt

```
> After running the code, you will get the following information
```shell
Your SM4 Encode Private Data Result: xxxxxxx , Please update configs/config.yaml
 Your Key：1234567890abcdef , Please update configs/config.yaml
 Your IV：123456789asdfghb , Please update configs/config.yaml

```
> Update the generated encrypted private key, Key, and IV to configs/config.yaml
```