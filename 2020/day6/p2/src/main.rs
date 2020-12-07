use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut num = 0;
    let mut yes = vec![true; 26];
    for line in stdin.lock().lines() {
        let l = &line.unwrap()[..];
        if l.len() == 0 {
            num += check_yeses(&mut yes);
            //println!("{}", num);
        } else {
            let line_str = String::from(l);
            for i in 0..26 {
                let c = ((i + 97) as u8) as char;
                if ! line_str.contains(c) {
                    yes[i]=false;
                }
            }
        }
    }
    num += check_yeses(&mut yes);
    println!("Total number of yes: {}", num);
}

fn check_yeses(yes: &mut Vec<bool>) -> i32 {
    let mut num = 0;
    for i in 0..26 {
        if yes[i] {
            num += 1;
        }
    }
    *yes = vec![true; 26];
    return num;
}