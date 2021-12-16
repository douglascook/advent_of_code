use std::collections::HashMap;
use std::fs;

fn main() {
    println!("Advent of Code Day 11 - Polymer explosion!");

    println!("\nTest Input:");
    polymerise("test_input.txt", 12);

    // println!("\nReal Input:");
    // polymerise("input.txt", 40);
}

fn polymerise(filepath: &str, generations: usize) {
    let input: String = fs::read_to_string(filepath).expect("Failed to read input file.");
    let (mut polymer, rules) = parse_input(&input);
    println!("Initial polymer = {}", polymer);

    for _generation in 0..generations {
        polymer = process_polymer(polymer, &rules);
        println!("Next generation = {}", polymer);
    }
    println!(
        "Result after {} generations is {}",
        generations,
        get_result(polymer)
    );
}

fn parse_input(input: &str) -> (String, HashMap<&str, &str>) {
    let (polymer, rule_list) = input.split_once("\n\n").expect("Has polymer; and rules");
    let mut rules: HashMap<&str, &str> = HashMap::new();
    for line in rule_list.trim().lines() {
        let (pair, new_element) = line
            .split_once(" -> ")
            .expect("Must have pair and new_element");
        rules.insert(pair, new_element);
    }
    (polymer.trim().to_string(), rules)
}

fn process_polymer(polymer: String, rules: &HashMap<&str, &str>) -> String {
    let mut next_polymer = polymer.to_string();
    // iterate in reverse so that we don't need to worry about shifting indices
    for i in (1..polymer.len()).rev() {
        let pair = &polymer[i - 1..=i];
        let new_element = rules.get(pair);
        if new_element.is_some() {
            println!(
                "Found a rule for {}, adding {}",
                pair,
                new_element.expect("asdasd")
            );
            next_polymer = vec![
                &next_polymer[0..i],
                new_element.expect("We know it's a string"),
                &next_polymer[i..next_polymer.len()],
            ]
            .join("");
        }
    }
    next_polymer
}

fn get_result(polymer: String) -> usize {
    let mut counts: HashMap<char, usize> = HashMap::new();

    for c in polymer.chars() {
        let count = counts.entry(c).or_insert(0);
        *count += 1;
    }
    let mut max: usize = 0;
    let mut min: usize = usize::max_value();
    for (_char, count) in counts.drain() {
        if count > max {
            max = count;
        }
        if count < min {
            min = count;
        }
    }
    max - min
}
