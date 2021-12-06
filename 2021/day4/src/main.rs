use std::collections::HashMap;
use std::fs;

const BOARD_SIZE: usize = 5;

fn main() {
    println!("Advent of Code Day 4 - Squid (Bingo) Game!");

    println!("\nTest Input:");
    play_game("test_input.txt");

    println!("\nReal Inputs:");
    play_game("input.txt");
}

fn play_game(filepath: &str) {
    let contents: String = fs::read_to_string(filepath).expect("Failed to read input file.");
    let (draw_string, boards_string) = contents
        .split_once("\n\n")
        .expect("Must be at least one board");

    let draw: Vec<usize> = draw_string
        .trim()
        .split(',')
        .map(|n| n.parse().expect("Must be number"))
        .collect();

    // Keep track of which numbers have been called
    let mut called_numbers: HashMap<usize, bool> = draw.iter().map(|&n| (n, false)).collect();

    let mut boards: Vec<BingoBoard> = boards_string
        .split("\n\n")
        .map(BingoBoard::from_string)
        .collect();

    let scores = calculate_scores(&mut boards, draw, &mut called_numbers);
    let winner = scores.first().expect("Must be a winner");
    println!(
        "Winning board is board {}, Unmarked sum = {}, Final score = {}",
        winner.0 + 1,
        winner.1,
        winner.2
    );
    let loser = scores.last().expect("Must be a loser");
    println!(
        "Losing board is board {}, Unmarked sum = {}, Final score = {}",
        loser.0 + 1,
        loser.1,
        loser.2
    );
}

/// Update boards as numbers are called until all boards are complete and return score for each board
fn calculate_scores(
    boards: &mut Vec<BingoBoard>,
    draw: Vec<usize>,
    called_numbers: &mut HashMap<usize, bool>,
) -> Vec<(usize, usize, usize)> {
    let mut scores: Vec<(usize, usize, usize)> = Vec::new();
    let mut completed: Vec<usize> = Vec::new();
    let board_count = boards.len();

    for number in draw {
        called_numbers.insert(number, true);

        for (board_index, board) in boards.iter_mut().enumerate().take(board_count) {
            // skip any boards that have already won
            if !completed.contains(&board_index) {
                if board.is_winner(called_numbers) {
                    completed.push(board_index);
                    let score = board.get_score(called_numbers);
                    scores.push((board_index, score, score * number));
                }
            }
        }
        if completed.len() == board_count {
            println!("All boards done!");
            break;
        }
    }
    scores
}

#[derive(Debug)]
struct BingoBoard {
    numbers: Vec<Vec<usize>>,
}

impl BingoBoard {
    fn from_string(string: &str) -> BingoBoard {
        let mut numbers: Vec<Vec<usize>> = Vec::new();
        for row in string.trim().split('\n') {
            numbers.push(
                row.split_whitespace()
                    .map(|n| n.parse().expect("Must be a number"))
                    .collect(),
            );
        }
        BingoBoard { numbers }
    }

    /// A board is a winner if any column or row is completely marked
    fn is_winner(&self, called_numbers: &HashMap<usize, bool>) -> bool {
        for i in 0..BOARD_SIZE {
            if self.numbers[i]
                .iter()
                .all(|n| *called_numbers.get(n).expect("Key exists"))
            {
                // println!("Row {} is a winner!", i);
                return true;
            }
        }
        for j in 0..BOARD_SIZE {
            if self
                .numbers
                .iter()
                .all(|row| *called_numbers.get(&row[j]).expect("Key exists"))
            {
                // println!("Column {} is a winner!", j);
                return true;
            }
        }
        false
    }

    /// Calculate score for the board = sum of all unmarked numbers
    fn get_score(&self, called_numbers: &HashMap<usize, bool>) -> usize {
        let mut total: usize = 0;
        for i in 0..BOARD_SIZE {
            for j in 0..BOARD_SIZE {
                if !called_numbers.get(&self.numbers[i][j]).expect("Key exists") {
                    total += self.numbers[i][j];
                }
            }
        }
        total
    }
}
