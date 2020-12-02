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
        let n = count_letter(c, pwd);
        if n >= r1 && n <= r2 {
            num_ok += 1;
        }
    }
    println!("Number of accepted passwords: {}", num_ok);
}

fn count_letter(ch:char, s:&str) -> i32 {
    let mut n = 0;
    let v:Vec<char> = s.chars().collect();
    for c in v {
        if c == ch {
            n += 1;
        }
    }
    return n;
}

