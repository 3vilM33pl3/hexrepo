fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
         .build_server(false)
         .build_client(true)
         .out_dir("./src/")
         .compile(
             &["E:/Hex/hexrepo/api/hexagon.proto"],
             &["E:/"],
         )?;
    Ok(())
 }