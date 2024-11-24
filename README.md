# Cryptographic algorithms

In this repository, you can find the source code for cryptographic algorithms implemented in Go.

## List of algorithms

### Stream ciphers

* Vernam cipher

### Block ciphers

* GOST 34.12-2015 (64-bit block - Magma)

### Hash functions

* SHA-1

## How to use

Makefile help message:

```sh
$ make help

Available commands:
        make help               - print this help
        make build              - build crypto-stream and crypto-block
        make clean              - clean build directory
        make test               - run tests
```

After executing `make build`, you can find the compiled binaries in the `build` directory.

```sh
$ make build
...
$ tree build
build
├── crypto-block
├── crypto-hash
└── crypto-stream
```

More information about this CLI application can be found in the help message:

```sh
$ ./build/crypto-block --help
Block cipher to encrypt/decrypt a file

Usage:
  crypto-block [flags]
  crypto-block [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  gost341264  GOST 34.12-2015 64-bit to encrypt/decrypt a file
  help        Help about any command

Flags:
  -h, --help            help for crypto-block
  -o, --output string   Output file path

Use "crypto-block [command] --help" for more information about a command.
```
```sh
$ ./build/crypto-hash --help
Hash a string

Usage:
  crypto-hash [flags]
  crypto-hash [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  sha1        SHA1 hash a string

Flags:
  -h, --help   help for crypto-hash

Use "crypto-hash [command] --help" for more information about a command.
```
```sh
$ ./build/crypto-stream --help
Stream cipher to encrypt/decrypt a string

Usage:
  crypto-stream [flags]
  crypto-stream [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  vernam      Vernam cipher to encrypt/decrypt a string

Flags:
  -h, --help   help for crypto-stream

Use "crypto-stream [command] --help" for more information about a command.
```

## Testing

```sh
$ echo "Hello, World!" > plain.txt 
$ ./build/crypto-block gost341264 encrypt ./plain.txt -o encrypted.bin
$ ./build/crypto-block gost341264 decrypt encrypted.bin -o decrypted.txt
$ cat decrypted.txt
Hello, World!
```
```sh
$ ./build/crypto-hash sha1 Hello, World! 
hash: 0a0a9f2a6772942557ab5355d76af442f8f65e01
```
```sh
$ ./build/crypto-stream vernam encrypt Hello, World!
key (hex): 7dc94c1628bab50dee784247e9
encrypted text (hex): 35ac207a4796955a810a2e23c8
$ ./build/crypto-stream vernam decrypt 7dc94c1628bab50dee784247e9 35ac207a4796955a810a2e23c8
decrypted text: Hello, World!
```

## Avalanche effect

```sh
$ ./build/crypto-hash sha1 Hello, World!                                                    
hash: 0a0a9f2a6772942557ab5355d76af442f8f65e01                                                                                                                           ptrvsrg@MacBook-Pro: ~/projects/nsu/crypto kube:(devrc/deploy-registry) git:(main) ✗  
$ ./build/crypto-hash sha1 Hello, World 
hash: 907d14fb3af2b0d4f18c2d46abe8aedce17367bd   
```

## License

This project is released under the [MIT License](https://github.com/ptrvsrg/crypto/blob/master/LICENSE).
