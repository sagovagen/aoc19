use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut numbers = Vec::<u32>::new();
    for line in stdin.lock().lines() {
        let s = &line.unwrap()[..];
        let num:u32 = s.parse::<u32>().unwrap();
        find_2020(num, &numbers);
        numbers.push(num);
    }
}

fn find_2020(num:u32, numbers:&Vec::<u32>) {
    if numbers.len() > 0 {
        for i in 0..numbers.len()-1 {
            if sum_is_2020(num, numbers[i]) {
                println!("{} + {} : {}", num, numbers[i], num*numbers[i]);
            }
        }
    }
}

fn sum_is_2020(first:u32, second:u32) -> bool {
    if first + second == 2020 {
        return true;
    }
    return false;
}

