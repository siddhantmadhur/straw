

mod file_explorer;

fn main() {

    let files = file_explorer::get_all_files(); 
    println!("Hello, world: {}", files.len());
}




