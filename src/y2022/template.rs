use crate::problem::Problem;

pub struct DayXX {}

impl Problem for DayXX {
    fn part_one(&self, input: &str) -> String {
        format!("{}", "Part one not yet implemented.")
    }

    fn part_two(&self, input: &str) -> String {
        format!("{}", "Part two not yet implemented.")
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "";

    #[test]
    fn test_part1() {
        let p = DayXX {};
        assert_eq!(p.part_one(INPUT), "todo");
    }

    #[test]
    fn test_part2() {
        let p = DayXX {};
        assert_eq!(p.part_two(INPUT), "todo");
    }
}
