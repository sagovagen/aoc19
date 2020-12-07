use std::io::{self, BufRead};

enum Attribute {  Nul = 0, Ecl, Hcl, Byr, Iyr, Eyr, Pid, Hgt }

fn main() {
    let stdin = io::stdin();
    let mut num_valid = 0;
    let mut p:Vec<bool> = [false, false, false, false, false, false, false, false].to_vec();
    for line in stdin.lock().lines() {
        let l = &line.unwrap()[..];
        if l.len() > 2 {
            let av:Vec<&str> = l.split(" ").collect();
            for a in av {
                let kv:Vec<&str> = a.split(":").collect();
                let key = kv[0];
                let atr:Attribute = match key {
                    "ecl" => Attribute::Ecl,
                    "hcl" => Attribute::Hcl,
                    "pid" => Attribute::Pid,
                    "eyr" => Attribute::Eyr,
                    "byr" => Attribute::Byr,
                    "iyr" => Attribute::Iyr,
                    "hgt" => Attribute::Hgt,
                    _ => Attribute::Nul
                };
                let ix = atr as usize;
                p[ix] = true;
            } 
        } else { // Empty line means end of record
            if check_valid(&mut p) {
                num_valid += 1;
            }
        }
    }
    if check_valid(&mut p) {
        num_valid += 1;
    }

    println!("Number of valid passports: {}", num_valid);
}

fn check_valid(p: &mut [bool]) -> bool {
    let mut valid = true;
    for i in 1..8 {
        if ! p[i] { valid = false; }
        p[i] = false;
    }
    return valid;
}