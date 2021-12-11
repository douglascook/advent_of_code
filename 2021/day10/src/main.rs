use std::collections::HashMap;
use std::fs;

fn main() {
    println!("Advent of Code Day 10 - Syntax Scoring");

    println!("\nTest Input:");
    check_syntax("test_input.txt");

    println!("\nReal Input:");
    check_syntax("input.txt");
}

fn check_syntax(filepath: &str) {
    let input: String = fs::read_to_string(filepath).expect("Failed to read input file.");

    let mut total_error_score: usize = 0;
    let mut completion_scores: Vec<usize> = Vec::new();
    for line in input.lines() {
        let (error_score, completion_score) = validate_line(line);
        total_error_score += error_score;
        if completion_score > 0 {
            completion_scores.push(completion_score);
        }
    }

    println!(
        "Total syntax error score for navigation system = {}",
        total_error_score
    );

    completion_scores.sort();
    println!(
        "Middle completion score = {}",
        completion_scores[(completion_scores.len() - 1) / 2]
    );
}

fn validate_line(line: &str) -> (usize, usize) {
    let mut char_stack: Vec<char> = Vec::new();
    println!("Processing line: {}", line);

    // TODO these can't be CONST, is there a nicer way to share them?
    let error_scores: HashMap<char, usize> = [(')', 3), (']', 57), ('}', 1197), ('>', 25137)]
        .into_iter()
        .collect();

    let close_to_open: HashMap<char, char> = [(')', '('), (']', '['), ('}', '{'), ('>', '<')]
        .into_iter()
        .collect();

    for c in line.chars() {
        // Opening tag, add to the stack
        if !close_to_open.contains_key(&c) {
            char_stack.push(c);
        // Check if closing tag matches previous tag left open and remove pair from stack
        } else if char_stack.last() == close_to_open.get(&c) {
            char_stack.pop();
        } else {
            let score = error_scores.get(&c).expect("Must be an error");
            println!("Line contains a syntax error with score {}", score);
            return (*score, 0);
        }
    }
    // Stack now contains all unclosed tags, iterate over them in reverse and
    // calculate score to close
    let completion_scores: HashMap<char, usize> = [('(', 1), ('[', 2), ('{', 3), ('<', 4)]
        .into_iter()
        .collect();

    let mut score = 0;
    char_stack.reverse();
    println!("Line is incomplete, remaining open tags = {:?}", char_stack);
    for c in char_stack {
        score *= 5;
        score += completion_scores.get(&c).expect("Must have a score");
    }
    (0, score)
}
