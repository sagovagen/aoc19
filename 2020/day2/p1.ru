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
        let n = count_letter(c, pwd);
        if n >= r1 && n <= r2 {
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

