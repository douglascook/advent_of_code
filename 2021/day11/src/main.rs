use std::fs;

fn main() {
    println!("Advent of Code Day 11 - Flashing Octopi!");

    println!("\nSample Input:");
    update_octopi("sample_input.txt", 2);

    println!("\nTest Input:");
    update_octopi("test_input.txt", 100);

    println!("\nTest Input - searching for synchronisation step:");
    update_octopi("test_input.txt", 1000);

    println!("\nReal Input:");
    update_octopi("input.txt", 100);

    println!("\nReal Input - searching for synchronisation step:");
    update_octopi("input.txt", 1000);
}

fn update_octopi(filepath: &str, steps: usize) {
    let input: String = fs::read_to_string(filepath).expect("Failed to read input file.");
    let mut cave = Cave::from_string(&input);
    for step in 1..=steps {
        println!("\nRunning step {}", step);
        let step_flash_count = cave.update();
        if step_flash_count == cave.height * cave.width {
            println!("All octopi synchronised at step {}!", step);
            return;
        }
    }
}

struct Cave {
    octopi: Vec<Vec<u32>>,
    height: usize,
    width: usize,
    flash_count: usize,
}

impl Cave {
    fn from_string(string: &str) -> Cave {
        let mut octopi: Vec<Vec<u32>> = Vec::new();
        for line in string.lines() {
            let row: Vec<u32> = line
                .chars()
                .map(|c| c.to_digit(10).expect("Must be digit"))
                .collect();
            octopi.push(row);
        }
        Cave {
            octopi: octopi.to_owned(),
            height: octopi.len(),
            width: octopi.first().expect("Must contain a row").len(),
            flash_count: 0,
        }
    }

    /// Update octopi energy levels for one step
    fn update(&mut self) -> usize {
        for j in 0..self.height {
            for i in 0..self.width {
                self.increment_energy(j, i);
            }
        }
        // reset any flashed octopi to energy level 0
        let mut step_flash_count = 0;
        for j in 0..self.height {
            for i in 0..self.width {
                if self.octopi[j][i] >= 10 {
                    self.octopi[j][i] = 0;
                    step_flash_count += 1;
                }
            }
        }
        println!("Energy levels =");
        for row in &self.octopi {
            println!("\t{:?}", row);
        }
        println!("Total flash count = {}", self.flash_count);

        step_flash_count
    }

    /// Flash the octopus at given coordinates, updating energy levels of its
    /// neighbours and flashing them in turn if they exceed the threshold
    fn increment_energy(&mut self, y: usize, x: usize) {
        self.octopi[y][x] += 1;
        // octopus has already flashed, cannot flash more than once per step
        if self.octopi[y][x] != 10 {
            return;
        }
        self.flash_count += 1;
        let mut min_y = 0;
        if y >= 1 {
            min_y = y - 1;
        }
        let mut min_x = 0;
        if x >= 1 {
            min_x = x - 1;
        }
        for j in min_y..=(y + 1).min(self.height - 1) {
            for i in min_x..=(x + 1).min(self.width - 1) {
                self.increment_energy(j, i);
            }
        }
    }
}
