import numpy as np

from mdqc.reduction import ReductionConfig, reduce_dimensions


def test_pca_reduction_shape():
    data = np.random.default_rng(0).normal(size=(10, 8))
    config = ReductionConfig(target_dims=4, algorithm="pca")
    reduced = reduce_dimensions(data, config)
    assert reduced.shape == (10, 4)
