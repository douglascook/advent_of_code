use std::fs;

fn main() {
    println!("Advent of Code Day 5 - Hydrothermal Venture!");

    println!("\nTest Input:");
    map_vents("test_input.txt");

    println!("\nReal Inputs:");
    map_vents("input.txt");
}

fn map_vents(filepath: &str) {
    let contents: String = fs::read_to_string(filepath).expect("Failed to read input file.");

    let vents: Vec<Line> = contents.lines().map(Line::from_string).collect();
    let max_x = vents
        .iter()
        .map(|l| Ord::max(l.a.x, l.b.x))
        .max()
        .expect("Must have values");
    let max_y = vents
        .iter()
        .map(|l| Ord::max(l.a.y, l.b.y))
        .max()
        .expect("Must have values");
    println!("Map dimensions = {} x {}", max_x, max_y);

    // values are zero indexed so map needs to have size max + 1
    let mut map = Map::inititalise(max_x + 1, max_y + 1);
    for vent in vents {
        map.plot_vent(&vent);
    }
    println!(
        "Number of coordinates with at least two vents = {}",
        map.count_at_least_n_vents(2)
    );
}

struct Map {
    vent_counts: Vec<Vec<usize>>,
}

impl Map {
    fn inititalise(width: usize, height: usize) -> Map {
        Map {
            vent_counts: vec![vec![0; width]; height],
        }
    }

    fn plot_vent(&mut self, vent: &Line) {
        for point in vent.get_coords() {
            self.vent_counts[point.x][point.y] += 1;
        }
    }

    fn count_at_least_n_vents(&self, vent_count: usize) -> usize {
        let mut total = 0;
        for row in &self.vent_counts {
            for count in row {
                if count >= &vent_count {
                    total += 1;
                }
            }
        }
        total
    }
}

struct Line {
    a: Point,
    b: Point,
}

impl Line {
    fn from_string(string: &str) -> Line {
        let (a, b) = string.split_once(" -> ").expect("Must have a and b");
        Line {
            a: Point::from_string(a),
            b: Point::from_string(b),
        }
    }

    fn get_coords(&self) -> Vec<Point> {
        let min_x = Ord::min(self.a.x, self.b.x);
        let max_x = Ord::max(self.a.x, self.b.x);
        let min_y = Ord::min(self.a.y, self.b.y);
        let max_y = Ord::max(self.a.y, self.b.y);

        let left;
        let right;
        if self.a.x == min_x {
            left = &self.a;
            right = &self.b;
        } else {
            left = &self.b;
            right = &self.a;
        }
        let delta = right.x - left.x;

        // horizontal
        if self.a.y == self.b.y {
            return (min_x..=max_x)
                .map(|i| Point { x: i, y: self.a.y })
                .collect();
        // vertical
        } else if self.a.x == self.b.x {
            return (min_y..=max_y)
                .map(|j| Point { x: self.a.x, y: j })
                .collect();
        // diagonal bottom left to top right
        } else if left.y < right.y {
            return (0..=delta)
                .map(|d| Point {
                    x: left.x + d,
                    y: left.y + d,
                })
                .collect();
        // diagonal top left to bottom right
        } else {
            return (0..=delta)
                .map(|d| Point {
                    x: left.x + d,
                    y: left.y - d,
                })
                .collect();
        }
    }
}

struct Point {
    x: usize,
    y: usize,
}

impl Point {
    fn from_string(string: &str) -> Point {
        let (x, y) = string.split_once(',').expect("Must have x and y coords");
        Point {
            x: x.parse().expect("Must be integer"),
            y: y.parse().expect("Must be integer"),
        }
    }
}
