use std::time::Instant;



mod file_explorer;
mod config;

fn main() {
    let cfg = config::new();    


    let start = Instant::now();
    let files = file_explorer::get_all_files(); 
    println!("Hello, world: {}", files.len());

    println!("Time elapsed: {:?}", start.elapsed());
}




