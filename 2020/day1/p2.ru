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
    if numbers.len() > 1 {
        for i in 0..numbers.len()-1 {
            for j in i+1..numbers.len()-1 {
                if sum_is_2020(num, numbers[i], numbers[j]) {
                    println!("{} + {} + {}: {}", num, numbers[i], numbers[j], num*numbers[i]*numbers[j]);
                }
            }
        }
    }
}

fn sum_is_2020(x:u32, y:u32, z:u32) -> bool {
    return (x+y+z) == 2020;
}

