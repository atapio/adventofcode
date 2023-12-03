use crate::problem::Problem;
use itertools::Itertools;

pub struct Day06 {}

impl Problem for Day06 {
    fn part_one(&self, input: &str) -> String {
        let marker = find_marker(input, 4);
        format!("{}", marker)
    }

    fn part_two(&self, input: &str) -> String {
        let marker = find_marker(input, 14);
        format!("{}", marker)
    }
}

fn find_marker(input: &str, wsize: usize) -> usize {
    let v = input.chars().collect::<Vec<_>>();

    let idx = v
        .windows(wsize)
        .enumerate()
        .find(|w| w.1.iter().unique().count() == w.1.len());
    idx.unwrap().0 + wsize
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let p = Day06 {};
        assert_eq!(p.part_one("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), "7");
        assert_eq!(p.part_one("bvwbjplbgvbhsrlpgdmjqwftvncz"), "5");
        assert_eq!(p.part_one("nppdvjthqldpwncqszvftbrmjlhg"), "6");
        assert_eq!(p.part_one("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), "10");
        assert_eq!(p.part_one("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), "11");
    }

    #[test]
    fn test_part2() {
        let p = Day06 {};
        assert_eq!(p.part_two("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), "19");
        assert_eq!(p.part_two("bvwbjplbgvbhsrlpgdmjqwftvncz"), "23");
        assert_eq!(p.part_two("nppdvjthqldpwncqszvftbrmjlhg"), "23");
        assert_eq!(p.part_two("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), "29");
        assert_eq!(p.part_two("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), "26");
    }
}
