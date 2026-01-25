#[derive(Debug)]
pub enum ComputeError {
    EmptyVector,
    MismatchedLength,
    ZeroNorm,
}

pub fn normalize(vec: &[f64]) -> Result<Vec<f64>, ComputeError> {
    if vec.is_empty() {
        return Err(ComputeError::EmptyVector);
    }
    let norm = vec.iter().map(|v| v * v).sum::<f64>().sqrt();
    if norm == 0.0 {
        return Err(ComputeError::ZeroNorm);
    }
    Ok(vec.iter().map(|v| v / norm).collect())
}

pub fn cosine_similarity(a: &[f64], b: &[f64]) -> Result<f64, ComputeError> {
    if a.is_empty() || b.is_empty() {
        return Err(ComputeError::EmptyVector);
    }
    if a.len() != b.len() {
        return Err(ComputeError::MismatchedLength);
    }

    let dot = a.iter().zip(b.iter()).map(|(x, y)| x * y).sum::<f64>();
    let norm_a = a.iter().map(|v| v * v).sum::<f64>().sqrt();
    let norm_b = b.iter().map(|v| v * v).sum::<f64>().sqrt();
    if norm_a == 0.0 || norm_b == 0.0 {
        return Err(ComputeError::ZeroNorm);
    }
    Ok(dot / (norm_a * norm_b))
}
