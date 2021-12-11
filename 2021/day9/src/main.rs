use std::collections::HashSet;
use std::fs;

fn main() {
    println!("Advent of Code Day 9 - Lava Tubes!");

    println!("\nTest Input:");
    navigate_tubes("test_input.txt");

    println!("\nReal Input:");
    navigate_tubes("input.txt");
}

fn navigate_tubes(filepath: &str) {
    let input: String = fs::read_to_string(filepath).expect("Failed to read input file.");
    let tubes = parse_input(&input);

    let low_points = find_low_points(&tubes);
    let risk_level: u32 = low_points.iter().map(|(j, i)| tubes[*j][*i] + 1).sum();
    println!("Total risk level = {}", risk_level);

    let mut basin_sizes: Vec<usize> = low_points
        .iter()
        .map(|(y, x)| calculate_basin_size(&tubes, *y, *x))
        .collect();
    println!("Calculated basin sizes {:?}", basin_sizes);

    basin_sizes.sort();
    basin_sizes.reverse();
    let product: usize = basin_sizes[0..3].iter().product();
    println!("Product of 3 largest basin sizes = {}", product);
}

fn parse_input(input: &str) -> Vec<Vec<u32>> {
    let mut tubes: Vec<Vec<u32>> = Vec::new();

    for line in input.lines() {
        let numbers: Vec<u32> = line
            .chars()
            .map(|c| c.to_digit(10).expect("Must be integer"))
            .collect();
        tubes.push(numbers);
    }
    tubes
}

/// Find all low points on the map, that is tubes whose height is lower than
/// all their neighbours
fn find_low_points(tubes: &[Vec<u32>]) -> Vec<(usize, usize)> {
    let height = tubes.len();
    let width = tubes.first().expect("Must be at least one row").len();

    let mut low_points = Vec::new();
    for j in 0..height {
        for i in 0..width {
            let tube_height = tubes[j][i];
            let mut low_point = true;
            for (y, x) in get_neighbours(j, i, width, height) {
                if tube_height >= tubes[y][x] {
                    low_point = false;
                    break;
                }
            }
            if low_point {
                low_points.push((j, i));
            }
        }
    }
    low_points
}

/// Return the size of basin around given low point
fn calculate_basin_size(tubes: &[Vec<u32>], y: usize, x: usize) -> usize {
    let height = tubes.len();
    let width = tubes.first().expect("Must be at least one row").len();

    let mut to_visit: HashSet<(usize, usize)> = HashSet::from_iter(vec![(y, x)]);
    let mut visited: HashSet<(usize, usize)> = HashSet::new();
    // loop over all remaining things to visit
    while !to_visit.is_empty() {
        let mut next_to_visit: HashSet<(usize, usize)> = HashSet::new();

        for (y, x) in to_visit.iter() {
            visited.insert((*y, *x));
            for (j, i) in get_neighbours(*y, *x, width, height) {
                if !visited.contains(&(j, i)) & (tubes[j][i] != 9) {
                    next_to_visit.insert((j, i));
                }
            }
        }
        to_visit = next_to_visit;
    }
    visited.len()
}

/// Return list of coordinates for all neighbours. Diagonals are not considered
/// as neighbours.
fn get_neighbours(y: usize, x: usize, width: usize, height: usize) -> Vec<(usize, usize)> {
    let mut neighbours: Vec<(usize, usize)> = Vec::new();
    if y > 0 {
        neighbours.push((y - 1, x));
    }
    if y < height - 1 {
        neighbours.push((y + 1, x));
    }
    if x > 0 {
        neighbours.push((y, x - 1));
    }
    if x < width - 1 {
        neighbours.push((y, x + 1));
    }
    neighbours
}
