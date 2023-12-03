use crate::problem::Problem;
use std::{collections::HashMap, ops};

use nom::{
    bytes::complete::tag, character::complete::i32, multi::separated_list1, sequence::tuple,
    IResult,
};

pub struct Day14 {}

impl Problem for Day14 {
    fn part_one(&self, input: &str) -> String {
        let mut grid = Self::parse(input);
        let max_y = grid.iter().map(|(x, _y)| x.y).max().unwrap();
        let mut finished = false;
        let potentials = vec![
            Point { x: 0, y: 1 },
            Point { x: -1, y: 1 },
            Point { x: 1, y: 1 },
        ];
        let mut count = 0;
        //while !finished && count < 100 {
        while !finished {
            //println!("count: {}", count);
            let mut p = Point { x: 500, y: 0 };
            let mut stopped = false;
            while p.y <= max_y {
                //println!("sand dropping: {} {:?}", inner_count, p);
                for potential in potentials.iter() {
                    stopped = true;
                    let next = p + potential;
                    match grid.get(&next) {
                        Some(_) => {}
                        None => {
                            stopped = false;
                            p = next;
                            break;
                        }
                    }
                }
                if stopped {
                    //println!("sand dropped: {} {:?}", count, p);
                    grid.set(p, 2);
                    break;
                }
            }
            //println!("loop finished: {} {:?} {}", count, p, max_y);
            if p.y <= max_y {
                count += 1;
                continue;
            }
            finished = true;
        }

        format!("{}", count)
    }

    fn part_two(&self, input: &str) -> String {
        let mut grid = Self::parse(input);
        let max_y = grid.iter().map(|(x, _y)| x.y).max().unwrap() + 2;
        let potentials = vec![
            Point { x: 0, y: 1 },
            Point { x: -1, y: 1 },
            Point { x: 1, y: 1 },
        ];
        let mut count = 0;
        //while !finished && count < 100 {
        loop {
            count += 1;
            //println!("count: {}", count);
            let mut p = Point { x: 500, y: 0 };
            let mut stopped = false;
            while p.y < max_y - 1 {
                //println!("sand dropping: {} {:?}", inner_count, p);
                for potential in potentials.iter() {
                    stopped = true;
                    let next = p + potential;
                    match grid.get(&next) {
                        Some(_) => {}
                        None => {
                            stopped = false;
                            p = next;
                            break;
                        }
                    }
                }
                if stopped {
                    //println!("sand dropped: {} {:?}", count, p);
                    grid.set(p, 2);
                    break;
                }
                if p.y == max_y - 1 {
                    grid.set(p, 2);
                    break;
                }
            }
            //println!("loop finished: {} {:?} {}", count, p, max_y);
            if (p == Point { x: 500, y: 0 }) {
                break;
            }
        }

        format!("{}", count)
    }
}

impl Day14 {
    fn parse(input: &str) -> SparseGrid<u32> {
        let mut grid = SparseGrid::<u32>::new();
        for line in input.lines() {
            match points(line) {
                Ok((_, points)) => {
                    points.windows(2).for_each(|window| {
                        let (start, end) = (window[0], window[1]);
                        grid.draw_line(start, end, 1)
                    });
                }
                Err(e) => {
                    println!("Error: {:?}", e);
                }
            }
        }
        grid
    }
}

fn points(input: &str) -> IResult<&str, Vec<Point>> {
    separated_list1(tag(" -> "), point)(input)
}

fn point(input: &str) -> IResult<&str, Point> {
    tuple((i32, tag(","), i32))(input).map(|(next_input, res)| {
        let (x, _, y) = res;
        (next_input, Point { x, y })
    })
}

#[derive(Debug, PartialEq, Eq, Hash, Clone, Copy)]
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

impl<T: Clone> SparseGrid<T> {
    fn new() -> SparseGrid<T> {
        SparseGrid {
            cells: HashMap::new(),
        }
    }

    fn set(&mut self, p: Point, value: T) {
        self.cells.insert(p, value);
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
            self.set(p, value.clone());
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
        let p = Day14 {};
        assert_eq!(p.part_one(INPUT), "24");
    }

    #[test]
    fn test_part2() {
        let p = Day14 {};
        assert_eq!(p.part_two(INPUT), "93");
    }
}
