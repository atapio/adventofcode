use crate::problem::Problem;
use std::{collections::HashMap, ops};

pub struct Day13 {}

impl Problem for Day13 {
    fn part_one(&self, input: &str) -> String {
        let grid = Self::parse(input);
        let count = grid.iter().count();
        format!("{}", count)
    }

    fn part_two(&self, input: &str) -> String {
        format!("{}", "Part two not yet implemented.")
    }
}

impl Day13 {
    fn parse(input: &str) -> SparseGrid<u8> {
        SparseGrid::<u8>::new()
    }
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

    fn set(&mut self, p: Point, value: T) {
        self.cells.insert(coord, value);
    }

    fn get(&self, p: &Point) -> Option<&T> {
        self.cells.get(p)
    }
    fn iter(&self) -> impl Iterator<Item = (&Point, &T)> {
        self.cells.iter()
    }

    fn draw_line(&mut self, start: Point, end: Point, value: T) {
        let mut p = start;
        while p != end {
            self.set(p, value);
            if p.x < end.x {
                p.x += 1;
            } else if p.x > end.x {
                p.x -= 1;
            } else if p.y < end.y {
                p.y += 1;
            } else if p.y > end.y {
                p.y -= 1;
            }
        }
        self.set(end, value);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9";

    #[test]
    fn test_part1() {
        let p = Day13 {};
        assert_eq!(p.part_one(INPUT), "todo");
    }

    #[test]
    fn test_part2() {
        let p = Day13 {};
        assert_eq!(p.part_two(INPUT), "todo");
    }
}
