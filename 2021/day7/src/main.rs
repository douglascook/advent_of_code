use std::collections::HashMap;
use std::fs;

fn main() {
    println!("Advent of Code Day 7 - Whale Attack!");

    println!("\nTest Input:");
    crab_shuffle("test_input.txt");

    println!("\nReal Input:");
    crab_shuffle("input.txt");
}

fn crab_shuffle(filepath: &str) {
    let input: String = fs::read_to_string(filepath).expect("Failed to read input file.");

    let crabs: Vec<isize> = input
        .trim()
        .split(',')
        .map(|n| n.parse().expect("Must be number"))
        .collect();

    let mut positions: HashMap<isize, isize> = HashMap::new();
    for c in &crabs {
        let count = positions.entry(*c).or_insert(0);
        *count += 1;
    }

    let min_pos: &isize = crabs.iter().min().expect("Must have minimum");
    let max_pos: &isize = crabs.iter().max().expect("Must have maximum");

    let min_fuel = (*min_pos..=*max_pos)
        .map(|p| calculate_fuel_cost(&positions, p))
        .min()
        .expect("Must be a number");
    println!(
        "Predicted minimum fuel required to align crabs = {}",
        min_fuel
    );

    let min_fuel = (*min_pos..=*max_pos)
        .map(|p| calculate_real_fuel_cost(&positions, p))
        .min()
        .expect("Must be a number");
    println!("Real minimum fuel required to align crabs = {}", min_fuel);
}

/// Move 1 position = 1 fuel
fn calculate_fuel_cost(positions: &HashMap<isize, isize>, target_position: isize) -> isize {
    return positions
        .iter()
        .map(|(pos, count)| (pos - target_position).abs() * count)
        .sum();
}

/// Move 1 position = n where n is the number of moves that crab has made
fn calculate_real_fuel_cost(positions: &HashMap<isize, isize>, target_position: isize) -> isize {
    // println!("Calculating target position = {}", target_position);
    let cost: isize = positions
        .iter()
        .map(|(pos, count)| fuel_cost(*pos, target_position) * count)
        .sum();
    // println!("Cost = {}", cost);
    cost
}

fn fuel_cost(position: isize, target_position: isize) -> isize {
    let diff = (position - target_position).abs();
    if diff == 0 {
        return 0;
    }
    (diff * (diff + 1)) / 2
}
