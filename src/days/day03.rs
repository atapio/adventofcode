use crate::problem::Problem;
use parse_display::{Display, FromStr};
use std::collections::HashSet;

pub struct Day03 {}

impl Problem for Day03 {
    fn part_one(&self, input: &str) -> String {
        let priority = self
            .parse(input)
            .iter()
            .fold(0, |total_priority, sack| total_priority + sack.priority());
        format!("{}", priority)
    }

    fn part_two(&self, input: &str) -> String {
        format!("{}", "Part two not yet implemented.")
    }
}

impl Day03 {
    fn parse(&self, input: &str) -> Vec<Rucksack> {
        input
            .lines()
            .map(|line| line.parse::<Rucksack>().unwrap())
            .collect()
    }
}

#[derive(Display, FromStr, PartialEq, Debug)]
#[display("{items}")]
struct Rucksack {
    #[from_str(regex = "[A-Za-z]+")]
    items: String,
}

impl Rucksack {
    fn priority(&self) -> u32 {
        let (first, second) = self.compartments();
        let a: HashSet<char> = first.chars().collect();
        let b: HashSet<char> = second.chars().collect();
        let intersection = a.intersection(&b);
        return intersection.fold(0, |sum, c| sum + priority(*c));
    }

    fn compartments(&self) -> (String, String) {
        let s = String::from(&self.items);
        let len = s.len();

        let first = &s[0..len / 2];
        let second = &s[len / 2..];
        return (first.to_string(), second.to_string());
    }
}

fn priority(c: char) -> u32 {
    if c.is_uppercase() {
        return 26 + c as u32 - 'A' as u32 + 1;
    }
    return c as u32 - 'a' as u32 + 1;
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";

    #[test]
    fn test_part1() {
        let p = Day03 {};
        assert_eq!(p.part_one(INPUT), "157");
    }

    #[test]
    fn test_part2() {
        let p = Day03 {};
        // assert_eq!(p.part_two(INPUT), "12");
    }

    #[test]
    fn test_priority() {
        assert_eq!(priority('a'), 1);
        assert_eq!(priority('z'), 26);
        assert_eq!(priority('A'), 27);
        assert_eq!(priority('Z'), 52);
    }
    #[test]
    fn test_compartments() {
        let r = Rucksack {
            items: String::from("abcdef"),
        };
        let (first, second) = r.compartments();
        assert_eq!(first, "abc");
        assert_eq!(second, "def");
    }
}
