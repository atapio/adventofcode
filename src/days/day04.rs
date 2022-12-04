use crate::problem::Problem;
use parse_display::{Display, FromStr};

pub struct Day04 {}

impl Problem for Day04 {
    fn part_one(&self, input: &str) -> String {
        let count = self
            .parse(input)
            .iter()
            .fold(0, |total, pair| total + pair.fully_contained() as u32);
        format!("{}", count)
    }

    fn part_two(&self, input: &str) -> String {
        let count = self.parse(input).iter().fold(0, |total, pair| {
            // println!("{:?} {}", pair, pair.has_overlap());
            total + pair.has_overlap() as u32
        });
        format!("{}", count)
    }
}

impl Day04 {
    fn parse(&self, input: &str) -> Vec<Pair> {
        input
            .lines()
            .map(|line| line.parse::<Pair>().unwrap())
            .collect()
    }
}

#[derive(Display, FromStr, PartialEq, Debug)]
#[display("{s1}-{e1},{s2}-{e2}")]
struct Pair {
    s1: u32,
    e1: u32,
    s2: u32,
    e2: u32,
}

impl Pair {
    fn fully_contained(&self) -> bool {
        (self.s1 <= self.s2 && self.e1 >= self.e2) || (self.s2 <= self.s1 && self.e2 >= self.e1)
    }
    fn has_overlap(&self) -> bool {
        (self.e1 >= self.s2 && self.s1 <= self.s2) || (self.e2 >= self.s1 && self.s2 <= self.s1)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";

    #[test]
    fn test_part1() {
        let p = Day04 {};
        assert_eq!(p.part_one(INPUT), "2");
    }

    #[test]
    fn test_part2() {
        let p = Day04 {};
        assert_eq!(p.part_two(INPUT), "4");
    }
}
