use std::fs;

fn main() {
    println!("Advent of Code Day 2 - Dive!");
    let contents = fs::read_to_string("input.txt").expect("Failed to read input file.");

    println!("\nCalculating position using basic directions...");
    basic_directions(&contents);

    println!("\nCalculating position using correct aiming directions...");
    aiming_directions(&contents);
}

fn basic_directions(contents: &str) {
    let mut depth = 0;
    let mut horizontal_position = 0;
    let mut delta: u32;
    for line in contents.lines() {
        let (direction, distance) = line
            .split_once(" ")
            .expect("Line must have direction and distance");

        delta = distance.parse().expect("Should be integer");
        if direction == "forward" {
            horizontal_position += delta;
        } else if direction == "down" {
            depth += delta;
        } else {
            depth -= delta;
        }
    }
    println!(
        "Depth = {}, Horizontal Position = {}, Product = {}",
        depth,
        horizontal_position,
        depth * horizontal_position,
    );
}

fn aiming_directions(contents: &str) {
    let mut depth = 0;
    let mut horizontal_position = 0;
    let mut aim = 0;
    let mut delta: u32;
    for line in contents.lines() {
        let (direction, distance) = line
            .split_once(" ")
            .expect("Line must have direction and distance");

        delta = distance.parse().expect("Should be integer");
        if direction == "forward" {
            horizontal_position += delta;
            depth += aim * delta;
        } else if direction == "down" {
            aim += delta;
        } else {
            aim -= delta;
        }
    }
    println!(
        "Depth = {}, Horizontal Position = {}, Product = {}",
        depth,
        horizontal_position,
        depth * horizontal_position
    );
}
