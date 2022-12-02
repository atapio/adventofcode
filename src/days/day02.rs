use crate::problem::Problem;
use parse_display::{Display, FromStr};

pub struct Day02 {}

impl Problem for Day02 {
    fn part_one(&self, input: &str) -> String {
        let score = self
            .parse(input)
            .iter()
            .fold(0, |total_score, rps| total_score + rps.score());
        format!("{}", score)
    }

    fn part_two(&self, input: &str) -> String {
        let score = self.parse(input).iter().fold(0, |total_score, rps| {
            total_score + rps.score_with_strategy()
        });
        format!("{}", score)
    }
}

impl Day02 {
    fn parse(&self, input: &str) -> Vec<RPS> {
        input
            .lines()
            .map(|line| line.parse::<RPS>().unwrap())
            .collect()
    }
}

#[derive(Display, FromStr, PartialEq, Debug)]
#[display("{first} {second}")]
struct RPS {
    #[from_str(regex = "[A-C]")]
    first: String,
    #[from_str(regex = "[X-Z]")]
    second: String,
}

impl RPS {
    fn score(&self) -> u32 {
        match self.first.as_str() {
            "A" => match self.second.as_str() {
                "X" => 1 + 3,
                "Y" => 2 + 6,
                "Z" => 3 + 0,
                _ => panic!("Invalid second value"),
            },
            "B" => match self.second.as_str() {
                "X" => 1 + 0,
                "Y" => 2 + 3,
                "Z" => 3 + 6,
                _ => panic!("Invalid second value"),
            },
            "C" => match self.second.as_str() {
                "X" => 1 + 6,
                "Y" => 2 + 0,
                "Z" => 3 + 3,
                _ => panic!("Invalid second value"),
            },
            _ => panic!("Invalid first value"),
        }
    }
    fn score_with_strategy(&self) -> u32 {
        match self.first.as_str() {
            "A" => match self.second.as_str() {
                // Rock
                "X" => 0 + 3,
                "Y" => 3 + 1,
                "Z" => 6 + 2,
                _ => panic!("Invalid second value"),
            },
            "B" => match self.second.as_str() {
                // Paper
                "X" => 0 + 1,
                "Y" => 3 + 2,
                "Z" => 6 + 3,
                _ => panic!("Invalid second value"),
            },
            "C" => match self.second.as_str() {
                // Scissors
                "X" => 0 + 2,
                "Y" => 3 + 3,
                "Z" => 6 + 1,
                _ => panic!("Invalid second value"),
            },
            _ => panic!("Invalid first value"),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let p = Day02 {};
        let input = "A Y
B X
C Z";
        assert_eq!(p.part_one(&input), "15");
    }

    #[test]
    fn test_part2() {
        let p = Day02 {};
        let input = "A Y
B X
C Z";
        assert_eq!(p.part_two(&input), "12");
    }
}
