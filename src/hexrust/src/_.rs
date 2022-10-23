#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Empty {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct HexLocation {
    #[prost(int64, tag="1")]
    pub x: i64,
    #[prost(int64, tag="2")]
    pub y: i64,
    #[prost(int64, tag="3")]
    pub z: i64,
    #[prost(string, tag="5")]
    pub hex_id: ::prost::alloc::string::String,
    #[prost(map="string, string", tag="6")]
    pub local_data: ::std::collections::HashMap<::prost::alloc::string::String, ::prost::alloc::string::String>,
    #[prost(map="string, string", tag="7")]
    pub global_data: ::std::collections::HashMap<::prost::alloc::string::String, ::prost::alloc::string::String>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct HexInfo {
    #[prost(string, tag="1")]
    pub id: ::prost::alloc::string::String,
    #[prost(map="string, string", tag="2")]
    pub data: ::std::collections::HashMap<::prost::alloc::string::String, ::prost::alloc::string::String>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct HexInfoList {
    #[prost(message, repeated, tag="1")]
    pub hex_info: ::prost::alloc::vec::Vec<HexInfo>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct HexLocationList {
    #[prost(message, repeated, tag="1")]
    pub hex_loc: ::prost::alloc::vec::Vec<HexLocation>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct HexagonGetRequest {
    #[prost(message, optional, tag="1")]
    pub hex_loc: ::core::option::Option<HexLocation>,
    #[prost(int64, tag="2")]
    pub radius: i64,
    #[prost(bool, tag="3")]
    pub fill: bool,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct HexIdList {
    #[prost(string, repeated, tag="1")]
    pub hex_id: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct HexIdData {
    #[prost(string, tag="1")]
    pub hex_id: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub data_key: ::prost::alloc::string::String,
    #[prost(string, tag="3")]
    pub value: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct HexLocData {
    #[prost(int64, tag="1")]
    pub x: i64,
    #[prost(int64, tag="2")]
    pub y: i64,
    #[prost(int64, tag="3")]
    pub z: i64,
    #[prost(string, tag="4")]
    pub data_key: ::prost::alloc::string::String,
    #[prost(string, tag="5")]
    pub value: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct HexLocDataList {
    #[prost(message, repeated, tag="1")]
    pub hex_loc_data: ::prost::alloc::vec::Vec<HexLocData>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Status {
    #[prost(string, tag="1")]
    pub msg: ::prost::alloc::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Result {
    #[prost(bool, tag="1")]
    pub success: bool,
}
/// Generated client implementations.
pub mod hexagon_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct HexagonServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl HexagonServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: std::convert::TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> HexagonServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> HexagonServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            HexagonServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        pub async fn repo_add_hexagon_info(
            &mut self,
            request: impl tonic::IntoRequest<super::HexInfoList>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/RepoAddHexagonInfo",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn repo_get_hexagon_info(
            &mut self,
            request: impl tonic::IntoRequest<super::HexIdList>,
        ) -> Result<tonic::Response<super::HexInfoList>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/RepoGetHexagonInfo",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn repo_get_hexagon_info_data(
            &mut self,
            request: impl tonic::IntoRequest<super::HexIdData>,
        ) -> Result<tonic::Response<super::HexIdData>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/RepoGetHexagonInfoData",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn repo_get_all_hexagon_info(
            &mut self,
            request: impl tonic::IntoRequest<super::Empty>,
        ) -> Result<tonic::Response<super::HexInfoList>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/RepoGetAllHexagonInfo",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn repo_del_hexagon_info(
            &mut self,
            request: impl tonic::IntoRequest<super::HexIdList>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/RepoDelHexagonInfo",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn repo_del_hexagon_info_data(
            &mut self,
            request: impl tonic::IntoRequest<super::HexIdData>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/RepoDelHexagonInfoData",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn map_add(
            &mut self,
            request: impl tonic::IntoRequest<super::HexLocationList>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/HexagonService/MapAdd");
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn map_add_data(
            &mut self,
            request: impl tonic::IntoRequest<super::HexLocDataList>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/MapAddData",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn map_get(
            &mut self,
            request: impl tonic::IntoRequest<super::HexagonGetRequest>,
        ) -> Result<tonic::Response<super::HexLocationList>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/HexagonService/MapGet");
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn map_update(
            &mut self,
            request: impl tonic::IntoRequest<super::HexLocation>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/HexagonService/MapUpdate");
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn map_update_data(
            &mut self,
            request: impl tonic::IntoRequest<super::HexLocation>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/MapUpdateData",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn map_remove(
            &mut self,
            request: impl tonic::IntoRequest<super::HexLocationList>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/HexagonService/MapRemove");
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn map_remove_data(
            &mut self,
            request: impl tonic::IntoRequest<super::HexLocation>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/MapRemoveData",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn get_status_server(
            &mut self,
            request: impl tonic::IntoRequest<super::Empty>,
        ) -> Result<tonic::Response<super::Status>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/GetStatusServer",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn get_status_storage(
            &mut self,
            request: impl tonic::IntoRequest<super::Empty>,
        ) -> Result<tonic::Response<super::Status>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/GetStatusStorage",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn get_status_clients(
            &mut self,
            request: impl tonic::IntoRequest<super::Empty>,
        ) -> Result<tonic::Response<super::Status>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/GetStatusClients",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn repo_del_all_hexagon_info(
            &mut self,
            request: impl tonic::IntoRequest<super::Empty>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/RepoDelAllHexagonInfo",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn map_remove_all(
            &mut self,
            request: impl tonic::IntoRequest<super::Empty>,
        ) -> Result<tonic::Response<super::Result>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/HexagonService/MapRemoveAll",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
    }
}
