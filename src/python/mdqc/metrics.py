import math

_QBER_TOLERANCE = {
    2: 0.11,
    4: 0.18,
    8: 0.24,
}


def qber_tolerance(dimension: int) -> float:
    if dimension in _QBER_TOLERANCE:
        return _QBER_TOLERANCE[dimension]
    raise ValueError("unsupported dimension for QBER tolerance")


def bits_per_photon(dimension: int) -> float:
    if dimension <= 0:
        raise ValueError("dimension must be > 0")
    return math.log2(dimension)
