extern crate regex;
use std::io::{self, BufRead};
use regex::Regex;


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
                let value = kv[1];
                let mut atr: Attribute = Attribute::Nul; 
                match key {
                    "ecl" => if check_ecl(value) { atr = Attribute::Ecl; }
                    "hcl" => if check_hcl(value) { atr = Attribute::Hcl; }
                    "pid" => if check_pid(value) { atr = Attribute::Pid; }
                    "eyr" => if check_eyr(value) { atr = Attribute::Eyr; }
                    "byr" => if check_byr(value) { atr = Attribute::Byr; }
                    "iyr" => if check_iyr(value) { atr = Attribute::Iyr; }
                    "hgt" => if check_hgt(value) { atr = Attribute::Hgt; }
                    _ => atr = Attribute::Nul
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

fn check_ecl(value: &str) -> bool { 
    let colors = [ "amb", "blu", "brn", "gry", "grn", "hzl", "oth" ];
    for s in colors.iter() {
        if s == &value { return true; }
    }
    return false;
}

fn check_hcl(value: &str) -> bool { 
    let re = Regex::new(r"^#[0-9a-f]{6}$").unwrap();
    return re.is_match(value);
 }

fn check_pid(value: &str) -> bool { 
    let re = Regex::new(r"^[0-9]{9}$").unwrap();
    return re.is_match(value);
}

fn check_eyr(value: &str) -> bool {
    let num:u32 = value.parse::<u32>().unwrap();
    return num >= 2020 && num <= 2030;
}

fn check_byr(value: &str) -> bool { 
    let num:u32 = value.parse::<u32>().unwrap();
    return num >= 1920 && num <= 2002;
}

fn check_iyr(value: &str) -> bool { 
    let num:u32 = value.parse::<u32>().unwrap();
    return num >= 2010 && num <= 2020;
}

fn check_hgt(value: &str) -> bool { 
    let re = Regex::new(r"^(\d*)([incm]{2})$").unwrap();
    for cap in re.captures_iter(value) {
        let num:u32 = cap[1].parse::<u32>().unwrap();
        if &cap[2] == "in" {
            return num >= 59 && num <= 76;
        } else if &cap[2] == "cm" {
            return num >= 150 && num <= 193;    
        }
    }
    return false;
}

fn check_valid(p: &mut [bool]) -> bool {
    let mut valid = true;
    for i in 1..8 {
        if ! p[i] { valid = false; }
        p[i] = false;
    }
    return valid;
}