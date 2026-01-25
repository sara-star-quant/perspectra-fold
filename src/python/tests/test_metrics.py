import math

from mdqc.metrics import bits_per_photon, qber_tolerance


def test_bits_per_photon():
    assert bits_per_photon(8) == math.log2(8)


def test_qber_tolerance_known_dimension():
    assert qber_tolerance(4) == 0.18
