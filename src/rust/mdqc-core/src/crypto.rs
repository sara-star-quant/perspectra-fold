use hkdf::Hkdf;
use sha2::Sha256;

#[derive(Debug)]
pub enum CryptoError {
    InvalidLength,
}

pub fn derive_session_key(
    qkd: &[u8],
    pqc: &[u8],
    classical: &[u8],
    info: &[u8],
    out_len: usize,
) -> Result<Vec<u8>, CryptoError> {
    let mut ikm = Vec::with_capacity(qkd.len() + pqc.len() + classical.len());
    ikm.extend_from_slice(qkd);
    ikm.extend_from_slice(pqc);
    ikm.extend_from_slice(classical);

    let hk = Hkdf::<Sha256>::new(None, &ikm);
    let mut okm = vec![0u8; out_len];
    hk.expand(info, &mut okm)
        .map_err(|_| CryptoError::InvalidLength)?;
    Ok(okm)
}
