use crate::problem::Problem;
use std::cmp;

pub struct Day15 {}

impl Problem for Day15 {
    fn part_one(&self, input: &str) -> String {
        let mut sum = 0;
        for step in input.strip_suffix('\n').expect("no suffix").split(',') {
            println!("{}", step);
            sum += hash(step);
        }

        format!("{}", sum)
    }

    fn part_two(&self, input: &str) -> String {
        format!("{}", "Part two not yet implemented.")
    }
}

fn hash(s: &str) -> u64 {
    let mut h: u64 = 0;
    for c in s.chars() {
        h += c as u64;
        h *= 17;
        h %= 256;
    }
    h
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
";

    #[test]
    fn test_sample() {
        let p = Day15 {};
        assert_eq!(p.part_one("HASH"), "52");
    }

    #[test]
    fn test_part1() {
        let p = Day15 {};
        assert_eq!(p.part_one(INPUT), "1320");
    }

    #[test]
    fn test_part2() {
        let p = Day15 {};
        assert_eq!(p.part_two(INPUT), "todo");
    }
}
