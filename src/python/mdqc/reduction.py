from __future__ import annotations

from dataclasses import dataclass
from typing import Literal

import numpy as np

Algorithm = Literal["pca", "umap"]


@dataclass(frozen=True)
class ReductionConfig:
    target_dims: int = 4
    algorithm: Algorithm = "pca"
    seed: int = 42


def reduce_dimensions(data: np.ndarray, config: ReductionConfig) -> np.ndarray:
    if data.ndim != 2:
        raise ValueError("data must be a 2D array")
    if config.target_dims <= 0:
        raise ValueError("target_dims must be > 0")
    if config.target_dims > data.shape[1]:
        raise ValueError("target_dims cannot exceed input dimensions")

    if config.algorithm == "pca":
        return _pca_reduce(data, config.target_dims)
    if config.algorithm == "umap":
        return _umap_reduce(data, config.target_dims, config.seed)

    raise ValueError(f"unknown algorithm: {config.algorithm}")


def _pca_reduce(data: np.ndarray, target_dims: int) -> np.ndarray:
    centered = data - np.mean(data, axis=0, keepdims=True)
    _, _, vh = np.linalg.svd(centered, full_matrices=False)
    components = vh[:target_dims].T
    return centered @ components


def _umap_reduce(data: np.ndarray, target_dims: int, seed: int) -> np.ndarray:
    try:
        import umap
    except ImportError as exc:  # pragma: no cover - optional dependency
        raise ImportError("umap-learn is required for algorithm='umap'") from exc

    reducer = umap.UMAP(n_components=target_dims, random_state=seed)
    return reducer.fit_transform(data)
