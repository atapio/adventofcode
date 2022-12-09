use parse_display::{Display, FromStr};
use std::{collections::HashMap, default};

use crate::problem::Problem;

pub struct Day09 {}

impl Problem for Day09 {
    fn part_one(&self, input: &str) -> String {
        let mut grid = SparseGrid::<u8>::new();
        let mut snake = Snake::new();
        grid.set(snake.tail_pos(), 1);
        for m in self.parse(input) {
            let mut left = m.steps;
            while left > 0 {
                snake.step(&m.direction);
                println!("{:?} {:?}", snake.head, snake.tail_pos());
                grid.set(snake.tail_pos(), 1);
                left -= 1;
            }
        }

        let total = grid.iter().fold(0, |sum, _| sum + 1);

        format!("{}", total)
    }

    fn part_two(&self, input: &str) -> String {
        format!("{}", "Part two not yet implemented.")
    }
}

impl Day09 {
    fn parse(&self, input: &str) -> Vec<Move> {
        input
            .lines()
            .map(|line| line.parse::<Move>().unwrap())
            .collect()
    }
}

#[derive(Debug)]
struct Snake {
    // absolute coordinate
    head: Coordinate,
    // relative to head
    tail: Coordinate,
}
impl Snake {
    fn new() -> Snake {
        Snake {
            head: Coordinate { x: 0, y: 0 },
            tail: Coordinate { x: 0, y: 0 },
        }
    }
    fn tail_pos(&self) -> Coordinate {
        Coordinate {
            x: self.head.x + self.tail.x,
            y: self.head.y + self.tail.y,
        }
    }

    fn step(&mut self, dir: &Direction) {
        match dir {
            Direction::U => {
                self.head.y += 1;
                self.tail.y -= 1;
            }
            Direction::D => {
                self.head.y -= 1;
                self.tail.y += 1;
            }
            Direction::R => {
                self.head.x += 1;
                self.tail.x -= 1;
            }
            Direction::L => {
                self.head.x -= 1;
                self.tail.x += 1;
            }
        }

        if !(self.tail.x >= -1 && self.tail.x <= 1 && self.tail.y >= -1 && self.tail.y <= 1) {
            self.tail = match self.tail {
                Coordinate { x: 2, y: 0 } => Coordinate { x: 1, y: 0 },
                Coordinate { x: -2, y: 0 } => Coordinate { x: -1, y: 0 },
                Coordinate { x: 0, y: 2 } => Coordinate { x: 0, y: 1 },
                Coordinate { x: 0, y: -2 } => Coordinate { x: 0, y: -1 },
                Coordinate { x: 2, y: 1 } => Coordinate { x: 1, y: 0 },
                Coordinate { x: -2, y: 1 } => Coordinate { x: -1, y: 0 },
                Coordinate { x: 2, y: -1 } => Coordinate { x: 1, y: 0 },
                Coordinate { x: -2, y: -1 } => Coordinate { x: -1, y: 0 },
                Coordinate { x: 1, y: 2 } => Coordinate { x: 0, y: 1 },
                Coordinate { x: 1, y: -2 } => Coordinate { x: 0, y: -1 },
                Coordinate { x: -1, y: 2 } => Coordinate { x: 0, y: 1 },
                Coordinate { x: -1, y: -2 } => Coordinate { x: 0, y: -1 },
                _ => panic!("unsupported tail position: {:?}", self.tail),
            };
        }
    }
}

#[derive(Display, FromStr, PartialEq, Debug, Clone)]
#[display(style = "UPPERCASE")]
enum Direction {
    R,
    U,
    D,
    L,
}

#[derive(Display, FromStr, PartialEq, Debug)]
#[display("{direction} {steps}")]
struct Move {
    direction: Direction,
    steps: u8,
}

#[derive(Debug, PartialEq, Eq, Hash, Clone)]
struct Coordinate {
    x: i32,
    y: i32,
}

struct SparseGrid<T> {
    cells: HashMap<Coordinate, T>,
}

impl<T> SparseGrid<T> {
    fn new() -> SparseGrid<T> {
        SparseGrid {
            cells: HashMap::new(),
        }
    }

    fn set(&mut self, coord: Coordinate, value: T) {
        self.cells.insert(coord, value);
    }
    fn get(&self, x: i32, y: i32) -> Option<&T> {
        self.cells.get(&Coordinate { x, y })
    }
    fn iter(&self) -> impl Iterator<Item = (&Coordinate, &T)> {
        self.cells.iter()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2";

    #[test]
    fn test_part1() {
        let p = Day09 {};
        assert_eq!(p.part_one(INPUT), "13");
    }

    #[test]
    fn test_part2() {
        let p = Day09 {};
        assert_eq!(p.part_two(INPUT), "todo");
    }
}
