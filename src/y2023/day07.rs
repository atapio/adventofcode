use parse_display::{Display, FromStr};

use crate::problem::Problem;

pub struct Day07 {}

impl Problem for Day07 {
    fn part_one(&self, input: &str) -> String {
        let mut hands = vec![];
        for line in input.lines() {
            if line.is_empty() {
                continue;
            }
            hands.push(line.parse::<Hand>().expect("bad line"));
        }

        format!("{:?}", hands);

        format!("{}", "Part one not yet implemented.")
    }

    fn part_two(&self, input: &str) -> String {
        format!("{}", "Part two not yet implemented.")
    }
}

#[derive(Display, FromStr, PartialEq, Debug)]
#[display("{cards} {bid}")]
struct Hand {
    cards: String,
    bid: u32,
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
";

    #[test]
    fn test_part1() {
        let p = Day07 {};
        assert_eq!(p.part_one(INPUT), "6440");
    }

    #[test]
    fn test_part2() {
        let p = Day07 {};
        assert_eq!(p.part_two(INPUT), "todo");
    }
}
