fn main() {
    println!("cargo:rerun-if-changed=../../../proto/core_compute.proto");
    println!("cargo:rerun-if-changed=../../../proto");

    tonic_build::configure()
        .build_server(true)
        .build_client(false)
        .compile(
            &["../../../proto/core_compute.proto"],
            &["../../../proto"],
        )
        .expect("failed to compile proto");
}
