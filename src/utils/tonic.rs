use anyhow;
use std::future::Future;
use std::result::Result;
use tonic::{Code, Response, Status};
use tonic_types::{ErrorDetails, StatusExt};
use tracing::error;

pub async fn escalate<T, R, F, FA>(
    req: tonic::Request<T>,
    fun: F,
) -> Result<tonic::Response<R>, tonic::Status>
where
    F: FnOnce(T) -> FA,
    FA: Future<Output = anyhow::Result<R>>,
{
    let result = fun(req.into_inner()).await;
    match result {
        Ok(res) => Ok(Response::new(res)),
        Err(any) => {
            let err_details = ErrorDetails::new();
            // Generate error status
            let status = Status::with_error_details(
                Code::InvalidArgument,
                "request contains invalid arguments",
                err_details,
            );
            error!("error handling GRPC request: {}", any);

            Err(status)
        }
    }
}
