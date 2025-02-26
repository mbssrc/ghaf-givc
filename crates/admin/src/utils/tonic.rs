use anyhow;
use std::future::Future;
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
        Err(any_err) => {
            // Convert root cause and stack to strings
            let stack: Vec<_> = any_err.chain().skip(1).map(ToString::to_string).collect();
            let cause = any_err.root_cause().to_string();

            // ...then dump them...
            error!("Local error cause is {cause}");
            stack.iter().for_each(|e| error!("Local reasons is {e}"));

            // ...then pack to ErrorDetails
            let err_details = ErrorDetails::with_debug_info(stack, cause);
            // Generate error status
            let status = Status::with_error_details(
                Code::InvalidArgument,
                "request contains invalid arguments",
                err_details,
            );
            error!("error handling GRPC request: {any_err}");

            Err(status)
        }
    }
}
