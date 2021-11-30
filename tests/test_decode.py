from brownie import accounts, SolTest
import pytest


@pytest.fixture
def test():
    return accounts[0].deploy(SolTest)


def test_decode(test):
    msg = bytes.fromhex("1afc010a210a1f0a0552fdfc0721120582654f163f220fc79bc1a805bb899cc003818beab2030a1f0a1d0a055f0f9a62d81205681d0d86d1220db9969b43aca5a3a001c8fea5630a210a1f0a05e999eb9d181205a44784045d220fafb595c106a2ffb3db01f1bfe585030a210a1f0a0587f3c67cf212050badb37c58220f9aaa92e0038bede9a10295b991ac020a210a1f0a0521b6862163120525253fec73220f8f8e88d001da91c4f102e8a4afc80412270802121e8dd7a9e28bf9f5059875921e668a5bdf2c7fc4844592d2572bcd0668d2d61a030304051200120012001200120012001200120012001200120012001200120012001200120012001200")
    test.decodeMainMessag(msg)