use hexagon_service_client::hexagon_service_client::HexagonServiceClient;
use hexagon_service_client::Empty;

pub mod hexagon_service_client {
    tonic::include_proto!("_");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = HexagonServiceClient::connect("http://[::1]:8080").await?;

    let request = tonic::Request::new(Empty{});

    let response = client.get_status_server(request).await?;
    println!("RESPONSE={:?}", response);
    Ok(())
}