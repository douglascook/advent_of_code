use std::fs;

// Working with binary values
const BASE: i32 = 2;
const SET: char = '1';
const UNSET: char = '0';

fn main() {
    println!("Advent of Code Day 3 - Binary Diagnostic");

    println!("\nTest Results:");
    run_diagnostics("test_input.txt");

    println!("\nReal Results:");
    run_diagnostics("input.txt");
}

fn run_diagnostics(filepath: &str) {
    let contents = fs::read_to_string(filepath).expect("Failed to read input file.");
    let (first, _) = contents.split_once('\n').expect("At least two lines.");
    let line_length: usize = first.len();
    let num_rows = contents.lines().count();

    let bits = parse_bit_strings(&contents);
    calculate_epsilon_and_gamma(&bits, line_length, num_rows);

    let oxygen = calculate_rating(&bits, &chars_equal);
    let co2 = calculate_rating(&bits, &chars_not_equal);
    println!(
        "Oxygen generator rating = {}, CO2 scrubber rating = {}, Product = {}",
        oxygen,
        co2,
        oxygen * co2
    );
}

// TODO can these be passed in as anonymous functions instead?
fn chars_equal(a: char, b: char) -> bool {
    return a == b;
}

fn chars_not_equal(a: char, b: char) -> bool {
    return a != b;
}

fn parse_bit_strings(contents: &str) -> Vec<Vec<char>> {
    let mut bits: Vec<Vec<char>> = Vec::new();
    for line in contents.lines() {
        let line_bits: Vec<char> = line.chars().collect();
        bits.push(line_bits);
    }
    return bits;
}

fn calculate_epsilon_and_gamma(bits: &Vec<Vec<char>>, line_length: usize, num_rows: usize) {
    let counters = count_set_bits_for_all_columns(&bits, line_length);

    let mut gamma = 0;
    let mut epsilon = 0;
    // Build up binary values for epsilon and gamma
    for (i, count) in counters.iter().rev().enumerate() {
        let power = BASE.pow(i as u32);
        // If bit in this position is set more often than not
        if count > &(num_rows / 2) {
            gamma += power;
        } else {
            epsilon += power;
        }
    }
    println!(
        "Gamma = {}, Epsilon = {}, Product = {}",
        gamma,
        epsilon,
        gamma * epsilon
    );
}

fn count_set_bits_for_all_columns(bits: &Vec<Vec<char>>, line_length: usize) -> Vec<usize> {
    let mut counters = vec![0 as usize; line_length];
    for line in bits {
        for (i, &bit) in line.iter().enumerate() {
            if bit == SET {
                counters[i] += 1;
            }
        }
    }
    return counters;
}

fn calculate_rating(bits: &Vec<Vec<char>>, filter_fn: &dyn Fn(char, char) -> bool) -> i32 {
    let mut oxygen = bits.clone();
    let mut index = 0;
    while oxygen.len() > 1 {
        let mut most_common = UNSET;
        let oxygen_sum: usize = oxygen.iter().filter(|b| b[index] == SET).count();
        // If number of set bits is greater or equal to number of unset bits
        if oxygen_sum as f32 >= oxygen.len() as f32 / 2.0 {
            most_common = SET;
        }
        oxygen = oxygen
            .iter()
            .filter(|b| filter_fn(b[index], most_common))
            .cloned()
            .collect();
        index += 1;
    }
    return convert_to_decimal(&oxygen[0]);
}

// TODO what is the deal with unsigned vs signed ints here???
fn convert_to_decimal(bits: &Vec<char>) -> i32 {
    let mut total: i32 = 0;
    for (i, bit) in bits.iter().rev().enumerate() {
        if bit == &SET {
            total += BASE.pow(i as u32);
        }
    }
    return total;
}
