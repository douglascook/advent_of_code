use std::fs;

fn main() {
    println!("Advent of Code Day 13 - Transparent Origami");

    println!("\nTest Input:");
    fold_it_up("test_input.txt");

    println!("\nReal Input:");
    fold_it_up("input.txt");
}

fn fold_it_up(filepath: &str) {
    let input: String = fs::read_to_string(filepath).expect("Failed to read input file.");
    let (coords, folding_instructions) = input.split_once("\n\n").expect("Must have coords");

    let mut paper = Paper::from_string(coords);
    paper.print();

    for instruction in folding_instructions.lines() {
        paper = paper.fold(instruction);
        paper.print();
    }
}

struct Paper {
    height: usize,
    width: usize,
    grid: Vec<Vec<char>>,
}

impl Paper {
    fn from_string(string: &str) -> Paper {
        let coords: Vec<(usize, usize)> = string.lines().map(parse_string_to_coords).collect();
        let height = coords.iter().map(|(_x, y)| *y).max().expect("Has a max") + 1;
        let width = coords.iter().map(|(x, _y)| *x).max().expect("Has a max") + 1;

        let mut grid = instantiate_grid(height, width);
        for (x, y) in coords {
            grid[y][x] = '#';
        }

        Paper {
            height,
            width,
            grid,
        }
    }

    fn print(&self) {
        println!("\nPaper dimensions are {} x {}", self.width, self.height);
        let mut total_dots = 0;
        for row in &self.grid {
            println!("{}", row.iter().collect::<String>());
            total_dots += row.iter().filter(|d| **d == '#').count();
        }
        println!("Contains {} dots in total", total_dots);
    }

    fn fold(&self, instruction: &str) -> Paper {
        let (axis, line) = instruction
            .split_once('=')
            .expect("Must have axis and value");
        let axis = axis.chars().last().expect("Must have an axis");
        let line = line.parse().expect("Must have a value");

        let new_height: usize;
        let new_width: usize;
        let mut new_grid: Vec<Vec<char>>;
        if axis == 'y' {
            new_height = line;
            new_width = self.width;

            // Next grid starts as the top half of the existing paper
            new_grid = instantiate_grid(new_height, new_width);
            new_grid.clone_from_slice(&self.grid[0..line]);

            // Go through each row in bottom half in reverse, updating the new grid
            let bottom = &self.grid[line..self.height];
            for (j, row) in bottom.iter().rev().enumerate() {
                for (i, dot) in row.iter().enumerate() {
                    if *dot == '#' {
                        new_grid[j][i] = '#';
                    }
                }
            }
        } else {
            new_height = self.height;
            new_width = line;
            new_grid = Vec::new();
            for row in &self.grid {
                // Each row starts as the left half of the existing row
                let mut new_row = vec!['.'; new_width];
                new_row.clone_from_slice(&row[0..line]);
                // Go through each element on RHS in reverse, updating the new row
                for (i, dot) in row[line..self.width].iter().rev().enumerate() {
                    if *dot == '#' {
                        new_row[i] = '#';
                    }
                }
                new_grid.push(new_row);
            }
        }
        Paper {
            height: new_height,
            width: new_width,
            grid: new_grid,
        }
    }
}

fn parse_string_to_coords(line: &str) -> (usize, usize) {
    let (x, y) = line.split_once(',').expect("Must have x and y");
    (
        x.parse().expect("Must be an int"),
        y.parse().expect("Must be an int"),
    )
}

fn instantiate_grid(height: usize, width: usize) -> Vec<Vec<char>> {
    let mut grid: Vec<Vec<char>> = Vec::new();
    for _row in 0..height {
        grid.push(vec!['.'; width]);
    }
    grid
}
