use std::net::SocketAddr;

use mdqc_core::compute::{cosine_similarity, ComputeError};
use tonic::{transport::Server, Request, Response, Status};

pub mod corepb {
    tonic::include_proto!("mdqc.core.v1");
}

use corepb::core_compute_server::{CoreCompute, CoreComputeServer};
use corepb::{CosineSimilarityRequest, CosineSimilarityResponse};

#[derive(Default)]
struct CoreComputeService;

#[tonic::async_trait]
impl CoreCompute for CoreComputeService {
    async fn cosine_similarity(
        &self,
        request: Request<CosineSimilarityRequest>,
    ) -> Result<Response<CosineSimilarityResponse>, Status> {
        let payload = request.into_inner();
        if payload.a.is_empty() || payload.b.is_empty() {
            return Err(Status::invalid_argument("vectors must be non-empty"));
        }
        if payload.a.len() != payload.b.len() {
            return Err(Status::invalid_argument("vectors must have equal length"));
        }

        let value = cosine_similarity(&payload.a, &payload.b).map_err(|err| match err {
            ComputeError::EmptyVector => Status::invalid_argument("vectors must be non-empty"),
            ComputeError::MismatchedLength => {
                Status::invalid_argument("vectors must have equal length")
            }
            ComputeError::ZeroNorm => Status::failed_precondition("zero-norm vector"),
        })?;

        Ok(Response::new(CosineSimilarityResponse { value }))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr: SocketAddr = std::env::var("MDQC_CORE_ADDR")
        .unwrap_or_else(|_| "0.0.0.0:50051".to_string())
        .parse()?;

    let service = CoreComputeService;

    println!("mdqc-core compute service listening on {}", addr);

    Server::builder()
        .add_service(CoreComputeServer::new(service))
        .serve(addr)
        .await?;

    Ok(())
}
