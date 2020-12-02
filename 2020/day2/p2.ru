use std::io::{self, BufRead};

fn main() {
    let mut num_ok = 0;
    let stdin = io::stdin();
    for line in stdin.lock().lines() {
        let s = &line.unwrap()[..];
        let v = split_line(s);
        let r1 = v[0].parse::<i32>().unwrap();
        let r2 = v[1].parse::<i32>().unwrap();
        let c = v[2].parse::<char>().unwrap();
        let pwd=v[3];
        if check_pwd(r1-1,r2-1,c,pwd) {
            num_ok += 1;
        }
    }
    println!("Number of accepted passwords: {}", num_ok);
}

fn split_line(line:&str) -> Vec<&str> {
    let mut value:Vec<&str> = [].to_vec();
    let v:Vec<&str> = line.split(" ").collect();
    let v1:Vec<&str> = v[0].split("-").collect();
    let v2:Vec<&str> = v[1].split(":").collect();
    value.push(v1[0]);
    value.push(v1[1]);
    value.push(v2[0]);
    value.push(v[2]);
    return value;
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

