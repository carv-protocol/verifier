# Verifier

Verifier nodes are lightweight nodes managed by the community. They play a crucial role in maintaining the protocol's integrity, thereby enhancing its trustworthiness and security. Currently, these nodes primarily check TEE attestations on-chain, verifying that the results are reliable and that the process maintains the privacy of user data.

For more information about verifier please read [verifier-node-explained](https://docs.carv.io/carv-protocol/verifier-node-explained)

## Building the source

Building verifier requires a Go (version 1.21 or later). You can install it using your favourite package manager.

To download dependencies during compilation, run

```shell
make init
```

These dependencies will be downloaded and compiled to `$GOPATH/bin`. This path needs to be added to the environment variable by running `export PATH=$PATH:$GOPATH/bin`. If it has already been added, ignore it.

Then, build the project by running

```shell
make all
```

## Running verifier

### Hardware Requirements

Minimum:

- CPU with 1+ cores
- 2GB RAM
- 4 MBit/sec download Internet service

Recommended:

- Fast CPU with 2+ cores
- 4GB+ RAM
- 8+ MBit/sec download Internet service

### Run via binary program

If this is your first time running verifier, you need to specify a private key. The private key will sign the verification transaction, and the corresponding address will receive carv rewards.

The private key can be passed to the verifier through startup parameters, or written into the configuration file.

#### through startup parameters

If you want to pass the private key through startup parameters, you need to set `wallet.mode` in the configuration file to `0`.

```shell
# Pass the private key in clear text
./verifier -conf ../configs/config.yaml -private-key <Your Private Key>

# By specifying keystore
./verifier -conf ../../configs/config.yaml -keystore-path <Path to keystore file> -keystore-password <keystore's password>
```

In order to facilitate user operation, verifier provides a tool to generate a new keystore, run
```shell
./verifier -generate-keystore -keystore-path <path to generate your keystore file>
```

#### through configuration file

- Configure the plain text private key: Set `wallet.mode` in the configuration file to `1`, and write the plain text private key into `wallet.private_key`

- Configure the path and password of the keystore: Set `wallet.mode` in the configuration file to `2`, and write the path and password of the keystore file into `wallet.keystore_path` and `wallet.keystore_password`.

### Run via docker

One of the quickest ways to run verifier is by using Docker:

```shell
docker run -d --name verifier -v /<Path To This Repository>/verifier/configs:/data/conf carv/verifier
```

Note that you need to configure your private key in the configuration file.Similarly, you can configure both plaintext and keystore methods.

## Contribution
We welcome contributions of all forms! If you have any questions, suggestions, or have identified any bugs, please submit them through Issues. If you would like to contribute code, please submit it via Pull Requests.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.
