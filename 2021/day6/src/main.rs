use std::fs;

fn main() {
    println!("Advent of Code Day 6 - Lantern Fish Explosion!");

    println!("\nTest Input:");
    count_fish("test_input.txt", 80);
    count_fish("test_input.txt", 256);

    println!("\nReal Inputs:");
    count_fish("input.txt", 80);
    count_fish("input.txt", 256);
}

fn count_fish(filepath: &str, generations: usize) {
    let input: String = fs::read_to_string(filepath).expect("Failed to read input file.");

    let fish: Vec<usize> = input
        .trim()
        .split(',')
        .map(|n| n.parse().expect("Must be number"))
        .collect();

    let mut counters: Vec<usize> = vec![0; 9];
    for f in fish {
        counters[f] += 1;
    }

    for _generation in 0..generations {
        counters = update_population(counters);
        // println!("Counters = {:?}", counters);
        // let total_fish: usize = counters.iter().sum();
        // println!("{} fish left after {} generations", total_fish, generations);
    }
    let total_fish: usize = counters.iter().sum();
    println!("{} fish left after {} generations", total_fish, generations);
}

fn update_population(counters: Vec<usize>) -> Vec<usize> {
    let current_generation = counters;
    let mut counters: Vec<usize> = vec![0; 8];

    // shift values to left by 1 position to decrement counts
    counters.copy_from_slice(&current_generation[1..9]);
    // new fish are born
    counters.push(current_generation[0]);
    // fish that just have timer reset to 6
    counters[6] += current_generation[0];

    counters
}
