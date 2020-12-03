use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();

    let mut board: Vec<bool> = [].to_vec();
    let lines = stdin.lock().lines();
    let mut rows = 0;
    for line in lines {
        let chars:Vec<char> = line.unwrap()[..].chars().collect();
        for c in chars {
            board.push(c == '#');
        }
        rows += 1;
    }
    let cols = (board.len() as i32) / rows;
    println!("rows: {}, cols: {}", rows, cols);

    let slopes = [(1,1),(3,1),(5,1),(7,1),(1,2)];
    let mut prod = 1;
    for slope in slopes.iter() {
        let num = count_trees(&board, rows, cols, slope.0, slope.1);
        prod *= num;
        println!("slope {},{} number of trees: {}", slope.0,slope.1,num);
    }

    println!("Product: {}", prod);

}

fn count_trees(board: &Vec<bool>, rows: i32, cols: i32, cstep: i32, rstep: i32) -> i32 {
    let mut n = 0;
    let mut r = 0;
    let mut c = 0;
    loop {
        let cm = c % cols;
        let ix = (r*cols + cm) as usize;
        //println!("r: {}, c: {}, tree: {}", r, c, board[ix]);
        if board[ix] {
            n += 1;
        }
        r += rstep;
        c += cstep;

        if r >= rows {
            break;
        }
    }
    return n;
}

