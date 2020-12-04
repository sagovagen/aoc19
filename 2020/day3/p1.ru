use std::io::{self, BufRead};

type Slope = (usize,usize);
type TreeMap = Vec<Vec<bool>>;

fn main() {
    let trees = read_map();
    println!("rows: {}, cols: {}", trees.len(), trees[0].len());

    let slopes:Vec<Slope> = [(1,1),(3,1),(5,1),(7,1),(1,2)].to_vec();
    let mut prod = 1;
    for slope in slopes.iter() {
        let num = count_trees(&trees, slope);
        prod *= num;
        println!("slope {},{} number of trees: {}", slope.0,slope.1,num);
    }

    println!("Product: {}", prod);

}

// Read tree map from stdin (lines of '..#.#....##.#' where . means no tree)
fn read_map() -> TreeMap {
    let stdin = io::stdin();
    let mut trees: TreeMap = [].to_vec();
    let lines = stdin.lock().lines();
    for line in lines {
        let mut row = [].to_vec();
        let chars:Vec<char> = line.unwrap()[..].chars().collect();
        for c in chars {
            row.push(c == '#');
        }
        trees.push(row);
    }
    return trees;
}

fn count_trees(trees: &TreeMap, slope: &Slope) -> i32 {
    let mut n = 0;
    let mut r:usize = 0;
    let mut c:usize = 0;
    let cols = trees[0].len();
    loop {
        let cm = c % cols;
        if trees[r][cm] {
            n += 1;
        }
        c += slope.0;
        r += slope.1;

        if r >= trees.len() {
            break;
        }
    }
    return n;
}

