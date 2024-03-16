use std::env::home_dir;

use walkdir::WalkDir;





pub fn get_all_files() -> Vec<String> {

    let mut git_projects = vec![];
    let home: String;

    match std::env::home_dir() {
        Some(d) => home = String::from(d.to_str().expect("Could not convert to string")),
        None => panic!("No home dir found")
    }

    for entry in WalkDir::new(home) {
        match entry {
            Ok(d) => {
                if &d.file_name().to_str().expect("Error in converting to string") == &".git" {
                    git_projects.insert(0, d.path().to_str().clone())                }
            },
            Err(err) => panic!("Error: {}", err)
        }
    };
    

    vec![]

}
