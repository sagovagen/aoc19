use std::io::{self, BufRead};

fn main() {
    let mut num_ok = 0;
    let stdin = io::stdin();
    for line in stdin.lock().lines() {
        let s = &line.unwrap()[..];
        let v:Vec<&str> = s.split("-").collect();
        let r1 = v[0].parse::<i32>().unwrap();
        let v:Vec<&str> = v[1].split(" ").collect();
        let r2 = v[0].parse::<i32>().unwrap();
        let pwd = v[2].trim();
        let v:Vec<&str> = v[1].split(":").collect();
        let c = v[0].parse::<char>().unwrap();
        if check_pwd(r1-1,r2-1,c,pwd) {
            num_ok += 1;
        }
    }
    println!("Number of accepted passwords: {}", num_ok);
}

fn check_pwd(p1:i32,p2:i32,ch:char,pwd:&str) -> bool {
    let p1:usize = p1 as usize;
    let p2:usize = p2 as usize;
    let v:Vec<char> = pwd.chars().collect();
    if p1<v.len() && p2<v.len() && v[p1] != v[p2] && (v[p1] == ch || v[p2] == ch) {
        return true;
    }   
    return false;
}

