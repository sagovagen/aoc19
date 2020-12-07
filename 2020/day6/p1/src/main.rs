use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut num = 0;
    let mut yes = vec![false; 26];    
    for line in stdin.lock().lines() {
        let l = &line.unwrap()[..];
        if l.len() == 0 {
            num += check_yeses(&mut yes);
            // println!("{}", num);
        } else {
            for c in l.chars() {
                let ix = (c as usize) - 97;
                //println!("{} {}", c, ix);
                yes[ix]=true;
            }
        }
    }
    num += check_yeses(&mut yes);
    println!("{}", num);
}

fn check_yeses(yes: &mut Vec<bool>) -> i32 {
    let mut num = 0;
    for i in 0..26 {
        let u = (i+97) as u8;
        if yes[i] {
            num += 1;
        }
    }
    *yes = vec![false; 26];
    return num;
}