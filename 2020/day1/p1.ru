use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let lines = stdin.lock().lines();
    let report: Vec<i32> = lines.map(|x| x.parse::i32().unwrap()).collect();
    for r in report {
        println!("{}", r);
    }
}

fn sum_is_2020(first: i32, second: i32) -> bool {
    if first + second == 2020 {
        return true;
    }
    return false;
}

