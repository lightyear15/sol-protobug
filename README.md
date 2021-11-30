# sol-protobug

repository with the only purpose of reporting a bug in the sol-protobuf compiler

The bug occurs when trying to decode a specific instance of the ``MainMessage`` as defined in ``test.proto``.
The messagge is generated by the Go implementation of the protobuf encoder, and is then passed to the protobuf ``decode`` function using a dummy contract defined in ``contracts/contract.sol``.
We use [brownie](https://eth-brownie.readthedocs.io/en/stable/toctree.html) as the EVM framework to compile and run tests.
## steps to reproduce the error

the repository already contains the compiled files for both Go and Solidity protobuf messages. In case you want to re-compile them, run:
- for Go protobuf
```bash
$ protoc --version
libprotoc 3.14.0

$ protoc-gen-go --version
protoc-gen-go v1.25.0-devel

$ protoc -I=. --go_out=. test.proto
```
For solidity protobuf
```bash
$ git show --summary --oneline 
be80f5f (HEAD -> master, origin/master, origin/HEAD) Merge pull request #22 from siburu/resupport-solidity-types

$ /run.sh --input ~/go/src/github.com/lightyear15/sol-protobug/test.proto  --output ~/go/src/github.com/lightyear15/sol-protobug/contracts
```

1. generate the Go-encoded protobug ``MainMessage``
```
$ go run ./...
1afc010a210a1f0a0552fdfc0721120582654f163f220fc79bc1a805bb899cc003818beab2030a1f0a1d0a055f0f9a62d81205681d0d86d1220db9969b43aca5a3a001c8fea5630a210a1f0a05e999eb9d181205a44784045d220fafb595c106a2ffb3db01f1bfe585030a210a1f0a0587f3c67cf212050badb37c58220f9aaa92e0038bede9a10295b991ac020a210a1f0a0521b6862163120525253fec73220f8f8e88d001da91c4f102e8a4afc80412270802121e8dd7a9e28bf9f5059875921e668a5bdf2c7fc4844592d2572bcd0668d2d61a030304051200120012001200120012001200120012001200120012001200120012001200120012001200
```
2. copy-paste the hex-encoded string into the ``tests/test_decode.py`` at line 10
```Python
    msg = bytes.fromhex("1afc010a210a1f0a0552fdfc0721120582654f163f220fc79bc1a805bb899cc003818beab2030a1f0a1d0a055f0f9a62d81205681d0d86d1220db9969b43aca5a3a001c8fea5630a210a1f0a05e999eb9d181205a44784045d220fafb595c106a2ffb3db01f1bfe585030a210a1f0a0587f3c67cf212050badb37c58220f9aaa92e0038bede9a10295b991ac020a210a1f0a0521b6862163120525253fec73220f8f8e88d001da91c4f102e8a4afc80412270802121e8dd7a9e28bf9f5059875921e668a5bdf2c7fc4844592d2572bcd0668d2d61a030304051200120012001200120012001200120012001200120012001200120012001200120012001200")
```
run the test using the brownie framework
```bash
$ brownie test
Brownie v1.17.1 - Python development framework for Ethereum

================================================================ test session starts =================================================================
platform linux -- Python 3.8.10, pytest-6.2.5, py-1.10.0, pluggy-1.0.0
rootdir: /home/giulio/go/src/github.com/lightyear15/sol-protobug
plugins: eth-brownie-1.17.1, web3-5.24.0, forked-1.3.0, hypothesis-6.24.0, xdist-1.34.0
collected 1 item                                                                                                                                     

Launching 'ganache-cli --port 8545 --gasLimit 12000000 --accounts 10 --hardfork istanbul --mnemonic brownie'...

tests/test_decode.py F                                                                                                                         [100%]

====================================================================== FAILURES ======================================================================
____________________________________________________________________ test_decode _____________________________________________________________________

test = <SolTest Contract '0x3194cBDC3dbcd3E11a07892e7bA5c3394048Cc87'>

    def test_decode(test):
        msg = bytes.fromhex("1afc010a210a1f0a0552fdfc0721120582654f163f220fc79bc1a805bb899cc003818beab2030a1f0a1d0a055f0f9a62d81205681d0d86d1220db9969b43aca5a3a001c8fea5630a210a1f0a05e999eb9d181205a44784045d220fafb595c106a2ffb3db01f1bfe585030a210a1f0a0587f3c67cf212050badb37c58220f9aaa92e0038bede9a10295b991ac020a210a1f0a0521b6862163120525253fec73220f8f8e88d001da91c4f102e8a4afc80412270802121e8dd7a9e28bf9f5059875921e668a5bdf2c7fc4844592d2572bcd0668d2d61a030304051200120012001200120012001200120012001200120012001200120012001200120012001200")
>       test.decodeMainMessag(msg)
E       brownie.exceptions.VirtualMachineError: revert

tests/test_decode.py:12: VirtualMachineError
============================================================== short test summary info ===============================================================
FAILED tests/test_decode.py::test_decode - brownie.exceptions.VirtualMachineError: revert
================================================================= 1 failed in 2.46s ==================================================================
```
