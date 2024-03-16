use std::{env::home_dir, str::FromStr};

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
                if (&d).clone().file_name().to_str().expect("Error in converting to string") == ".git" {
                    git_projects.insert(0, String::from_str(d.path().to_str().expect("Not str")).expect("Could not convert"));
                }
            },
            Err(err) => () 
        }
    };


    for prj in &git_projects {
        println!("Prj: {}", prj);
    }

    git_projects
}
