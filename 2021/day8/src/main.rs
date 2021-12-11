use std::collections::HashMap;
use std::collections::HashSet;
use std::fs;

fn main() {
    println!("Advent of Code Day 8 - Seven Segment Search!");

    println!("\nSingle Input:");
    process_displays("single_input.txt");

    println!("\nTest Input:");
    process_displays("test_input.txt");

    println!("\nReal Input:");
    process_displays("input.txt");
}

fn process_displays(filepath: &str) {
    let input: String = fs::read_to_string(filepath).expect("Failed to read input file.");
    let displays = input.lines().map(Display::from_line);

    let mut total_output = 0;
    for mut display in displays {
        println!(
            "Digits = {:?}, Output = {:?}",
            display.digits, display.output
        );
        display.map_digits();
        total_output += display.get_output_value();
    }
    println!("Sum of outputs for all displays = {}", total_output);
}

// Lifetime of a struct must match/exceed lifetime of any references it contains
struct Display {
    // each digit is made up of a set of segments
    digits: Vec<HashSet<char>>,
    output: Vec<HashSet<char>>,
    // mapping from digit value to the set of segments that it uses
    mapping: HashMap<usize, HashSet<char>>,
}

// TODO how do lifetimes work???
impl Display {
    fn from_line(line: &str) -> Display {
        let (digits, output) = line.split_once(" | ").expect("Must have digits and output");

        Display {
            digits: digits
                .split(' ')
                .map(|d| HashSet::from_iter(d.chars()))
                .collect(),
            output: output
                .split(' ')
                .map(|d| HashSet::from_iter(d.chars()))
                .collect(),
            mapping: HashMap::new(),
        }
    }

    fn map_digits(&mut self) {
        // Insert digits that can be uniquely identified by their number
        // of segments: 1, 4, 7, 8
        for (value, segment_count) in &[(1, 2), (4, 4), (7, 3), (8, 7)] {
            let digit = self
                .digits
                .iter()
                .find(|d| d.len() == *segment_count)
                .expect("Digit exists");
            // TODO how does to_owned differ from clone?
            println!("{} = {:?}", value, digit);
            self.mapping.insert(*value, digit.to_owned());
        }

        // Next look at digits made up of 6 segments: 0, 6, 9
        for digit in self.digits.iter().filter(|d| d.len() == 6) {
            // Contains 4 as a subset -> 9
            if self
                .mapping
                .get(&4)
                .expect("Digit has been mapped.")
                .is_subset(digit)
            {
                println!("9 = {:?}", digit);
                self.mapping.insert(9, digit.to_owned());
            // contains 1 as a subset -> 0
            } else if self
                .mapping
                .get(&1)
                .expect("Digit has been_mapped.")
                .is_subset(digit)
            {
                println!("0 = {:?}", digit);
                self.mapping.insert(0, digit.to_owned());
            // contains neither -> 6
            } else {
                println!("6 = {:?}", digit);
                self.mapping.insert(6, digit.to_owned());
            }
        }
        //
        // Look at remaining digits, which are made up of 5 segments: 2, 3, 5
        for digit in self.digits.iter().filter(|d| d.len() == 5) {
            // Contains 1 as a subset -> 3
            if self
                .mapping
                .get(&1)
                .expect("Digit has been_mapped.")
                .is_subset(digit)
            {
                println!("3 = {:?}", digit);
                self.mapping.insert(3, digit.to_owned());
            // Left with 2 and 5, check size of intersection with 4 to tell
            // which is which
            } else {
                let four = self.mapping.get(&4).expect("Digit has been_mapped.");
                if four.intersection(digit).count() == 2 {
                    println!("2 = {:?}", digit);
                    self.mapping.insert(2, digit.to_owned());
                } else {
                    println!("5 = {:?}", digit);
                    self.mapping.insert(5, digit.to_owned());
                }
            }
        }
    }

    fn get_output_value(&self) -> usize {
        let mut total: usize = 0;
        let base: usize = 10;
        for (i, digit) in self.output.iter().rev().enumerate() {
            let (value, _digit) = self
                .mapping
                .iter()
                .find(|(_k, v)| *v == digit)
                .expect("please work");
            total += value * base.pow(i as u32);
        }
        println!("Output value = {}", total);
        total
    }
}
