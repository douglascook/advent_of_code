use std::fs;

fn main() {
    println!("Advent of Code Day 1 - Sonar Sweep!");
    let contents = fs::read_to_string("input.txt").expect("Failed to read input file.");
    let lines: Vec<i32> = contents.lines().map(parse_to_int).collect();

    println!(
        "The total number of single line increases is {}",
        count_single_line_increases(&lines)
    );
    println!(
        "The total number of three line window increases is {}",
        count_sliding_window_increases(&lines, 3)
    );
}

fn parse_to_int(line: &str) -> i32 {
    return line.trim().parse().expect("Expecting an integer!");
}

fn count_single_line_increases(lines: &[i32]) -> i32 {
    let mut increase_count = 0;
    let mut index = 1;
    while index < lines.len() {
        if lines[index] > lines[index - 1] {
            increase_count += 1;
        }
        index += 1;
    }
    return increase_count;
}

fn count_sliding_window_increases(lines: &[i32], window_size: usize) -> i32 {
    let mut increase_count = 0;
    let mut index = 1;
    let mut previous_sum: i32 = lines[0..window_size].iter().sum();
    while index < lines.len() - 2 {
        let current_sum: i32 = lines[index..index + window_size].iter().sum();
        if current_sum > previous_sum {
            increase_count += 1;
        }
        previous_sum = current_sum;
        index += 1;
    }
    return increase_count;
}
