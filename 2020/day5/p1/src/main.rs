extern crate regex;
use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut max_id = 0;
    let mut min_id = 1000;
    let mut id_list: Vec<i32> = [].to_vec();
    for line in stdin.lock().lines() {
        let l = &line.unwrap()[..];
        let binary = l.replace("F", "0").replace("B", "1").replace("L", "0").replace("R", "1");
        let row = &binary[..7];
        let rownum = i64::from_str_radix(&row, 2).unwrap() as i32;
        let col = &binary[7..];
        let colnum = i64::from_str_radix(&col, 2).unwrap() as i32;
        let id = rownum * 8 + colnum;
        
        //println!("row: {}, column: {}, ID: {}", rownum, colnum, id);
        if id > max_id { max_id = id; }
        if id < min_id { min_id = id; }
        id_list.push(id);
    }
    println!("Min ID: {}", min_id);
    println!("Max ID: {}", max_id);

    for i in min_id .. max_id {
        let mut found = false;
        for n in id_list.iter() {
            if i == *n { found = true; break; }
        }
        if ! found { println!("ID {} could not be found", i); }
    }

}

