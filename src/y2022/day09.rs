use parse_display::{Display, FromStr};
use std::{
    collections::{HashMap, VecDeque},
    ops,
};

use crate::problem::Problem;

pub struct Day09 {}

impl Problem for Day09 {
    fn part_one(&self, input: &str) -> String {
        let mut grid = SparseGrid::<u8>::new();
        let mut snake = Snake::new(2);
        grid.set(snake.tail_pos(), 1);
        for m in self.parse(input) {
            let mut left = m.steps;
            //println!("move: {:?}", m);
            while left > 0 {
                snake.step(&m.direction);
                grid.set(snake.tail_pos(), 1);
                left -= 1;
            }
        }

        let total = grid.iter().fold(0, |sum, _| sum + 1);

        format!("{}", total)
    }

    fn part_two(&self, input: &str) -> String {
        let mut grid = SparseGrid::<u8>::new();
        let mut snake = Snake::new(10);
        grid.set(snake.tail_pos(), 1);
        for m in self.parse(input) {
            let mut left = m.steps;
            //println!("move: {:?}", m);
            while left > 0 {
                snake.step(&m.direction);
                grid.set(snake.tail_pos(), 1);
                left -= 1;
            }
        }

        let total = grid.iter().fold(0, |sum, _| sum + 1);

        format!("{}", total)
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
    knots: VecDeque<Point>,
}
impl Snake {
    fn new(len: usize) -> Snake {
        let mut s = Snake {
            knots: VecDeque::with_capacity(len),
        };
        for _ in 0..len {
            s.knots.push_back(Point { x: 0, y: 0 });
        }
        s
    }
    fn tail_pos(&self) -> Point {
        self.knots.back().unwrap().clone()
    }

    fn step(&mut self, dir: &Direction) {
        self.step_head(dir);
        self.step_knots();
    }

    fn step_head(&mut self, dir: &Direction) {
        let head = self.knots.front_mut().unwrap();

        match dir {
            Direction::U => {
                head.y += 1;
            }
            Direction::D => {
                head.y -= 1;
            }
            Direction::R => {
                head.x += 1;
            }
            Direction::L => {
                head.x -= 1;
            }
        }
    }

    fn step_knots(&mut self) {
        for i in 1..self.knots.len() {
            let head = self.knots.get(i - 1).unwrap().clone();
            let knot = self.knots.get(i).unwrap().clone();

            let diff = knot.clone() - &head;

            if diff.x >= -1 && diff.x <= 1 && diff.y >= -1 && diff.y <= 1 {
                //let knot_mut = self.knots.get_mut(i).unwrap();
                //*knot_mut = diff.clone() + &head;
                //println!("i: {} diff: {:?} knot: {:?}", i, diff, *knot_mut);
                //println!("i: {} diff: {:?} knot: {:?}", i, diff, knot);
            } else {
                let delta = match diff {
                    Point { x: 2, y: 0 } => Point { x: -1, y: 0 },
                    Point { x: -2, y: 0 } => Point { x: 1, y: 0 },
                    Point { x: 0, y: 2 } => Point { x: 0, y: -1 },
                    Point { x: 0, y: -2 } => Point { x: 0, y: 1 },
                    Point { x: 2, y: 1 } => Point { x: -1, y: -1 },
                    Point { x: -2, y: 1 } => Point { x: 1, y: -1 },
                    Point { x: 2, y: -1 } => Point { x: -1, y: 1 },
                    Point { x: -2, y: -1 } => Point { x: 1, y: 1 },
                    Point { x: 1, y: 2 } => Point { x: -1, y: -1 },
                    Point { x: 1, y: -2 } => Point { x: -1, y: 1 },
                    Point { x: -1, y: 2 } => Point { x: 1, y: -1 },
                    Point { x: -1, y: -2 } => Point { x: 1, y: 1 },
                    Point { x: 2, y: 2 } => Point { x: -1, y: -1 },
                    Point { x: -2, y: -2 } => Point { x: 1, y: 1 },
                    Point { x: 2, y: -2 } => Point { x: -1, y: 1 },
                    Point { x: -2, y: 2 } => Point { x: 1, y: -1 },
                    _ => panic!("unsupported tail position: {:?}", diff),
                };
                let knot_mut = self.knots.get_mut(i).unwrap();
                *knot_mut = delta + &knot;
                //println!("i: {} diff: {:?} knot: {:?}", i, diff, *knot_mut);
            }
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
struct Point {
    x: i32,
    y: i32,
}

impl ops::Add<&Point> for Point {
    type Output = Point;

    fn add(self, rhs: &Point) -> Point {
        Point {
            x: self.x + rhs.x,
            y: self.y + rhs.y,
        }
    }
}

impl ops::Sub<&Point> for Point {
    type Output = Point;

    fn sub(self, rhs: &Point) -> Point {
        Point {
            x: self.x - rhs.x,
            y: self.y - rhs.y,
        }
    }
}

struct SparseGrid<T> {
    cells: HashMap<Point, T>,
}

impl<T> SparseGrid<T> {
    fn new() -> SparseGrid<T> {
        SparseGrid {
            cells: HashMap::new(),
        }
    }

    fn set(&mut self, coord: Point, value: T) {
        self.cells.insert(coord, value);
    }
    fn get(&self, x: i32, y: i32) -> Option<&T> {
        self.cells.get(&Point { x, y })
    }
    fn iter(&self) -> impl Iterator<Item = (&Point, &T)> {
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
        assert_eq!(p.part_two(INPUT), "1");
        assert_eq!(
            p.part_two(
                "R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20"
            ),
            "36"
        );
    }
}
