use std::fs;

const BOARD_SIZE: usize = 5;

#[derive(Debug)]
struct Number {
    value: usize,
    marked: bool,
}

#[derive(Debug)]
struct BingoBoard {
    numbers: Vec<Vec<Number>>,
}

impl BingoBoard {
    fn from_string(string: &str) -> BingoBoard {
        let mut numbers: Vec<Vec<Number>> = Vec::new();
        for row in string.trim().split('\n') {
            numbers.push(
                row.split_whitespace()
                    .map(|n| Number {
                        value: n.parse().expect("Must be a number"),
                        marked: false,
                    })
                    .collect(),
            );
        }
        BingoBoard { numbers }
    }

    /// Update the board, marking any numbers matching the called number
    fn update(&mut self, called_number: usize) {
        let mut updated = false;
        for i in 0..BOARD_SIZE {
            if updated {
                return;
            }
            for j in 0..BOARD_SIZE {
                if self.numbers[i][j].value == called_number {
                    self.numbers[i][j].marked = true;
                    updated = true;
                    break;
                }
            }
        }
    }

    /// A board is a winner if any column or row is completely marked
    fn is_winner(&self) -> bool {
        for i in 0..BOARD_SIZE {
            if self.numbers[i].iter().all(|n| n.marked) {
                // println!("Row {} is a winner!", i);
                return true;
            }
        }
        for j in 0..BOARD_SIZE {
            if self.numbers.iter().all(|r| r[j].marked) {
                // println!("Column {} is a winner!", j);
                return true;
            }
        }
        false
    }

    /// Calculate score for the board = sum of all unmarked numbers
    fn get_score(&self) -> usize {
        let mut total: usize = 0;
        for i in 0..BOARD_SIZE {
            for j in 0..BOARD_SIZE {
                if !self.numbers[i][j].marked {
                    total += self.numbers[i][j].value;
                }
            }
        }
        total
    }
}

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

    let mut boards: Vec<BingoBoard> = boards_string
        .split("\n\n")
        .map(BingoBoard::from_string)
        .collect();

    let scores = calculate_scores(&mut boards, draw);
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
fn calculate_scores(boards: &mut Vec<BingoBoard>, draw: Vec<usize>) -> Vec<(usize, usize, usize)> {
    let mut scores: Vec<(usize, usize, usize)> = Vec::new();
    let mut completed: Vec<usize> = Vec::new();
    let board_count = boards.len();

    for called_number in draw {
        for board_index in 0..board_count {
            // skip any boards that have already won
            if !completed.contains(&board_index) {
                let board = &mut boards[board_index];
                board.update(called_number);
                if board.is_winner() {
                    completed.push(board_index);
                    let score = board.get_score();
                    scores.push((board_index, score, score * called_number));
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
