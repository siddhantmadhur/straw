use std::{env::home_dir, fs, os, process::exit};



pub struct Config {
    directories: Vec<String>
}


pub fn new() -> Config {
    let mut cfg_content: String = String::from("");
    let cfg_file = fs::read_to_string(String::from(home_dir().expect("Error getting home dir").to_str().expect("Error converting to string")) + "/.straw");
    match cfg_file {
        Ok(d)=> cfg_content = d,
        Err(_)=> {
            println!("Config does not exist. Consider adding paths for it to scan.");
        }
    }

    let mut directories: Vec<String> = vec![];

    for line in cfg_content.lines() {
        directories.push(String::from(line));
    }

    Config{
        directories 
    }
}
