use crate::problem::Problem;
use std::cmp;

pub struct Day13 {}

impl Problem for Day13 {
    fn part_one(&self, input: &str) -> String {
        let mut patterns = vec![];
        let mut pattern = vec![];
        for line in input.lines() {
            if line.is_empty() {
                patterns.push(pattern);
                pattern = vec![];
                continue;
            }
            pattern.push(line.to_owned())
        }
        patterns.push(pattern);

        let mut sum = 0;
        for pattern in &patterns {
            println!("{:?}", pattern);
            match reflection(pattern) {
                Some(i) => {
                    println!("Found horizontal reflection at {}", i);
                    sum += 100 * i;
                }
                None => println!("No horizontal reflection found"),
            }
            let transposed = transpose(pattern);
            println!("{:?}", transposed);
            match reflection(&transposed) {
                Some(i) => {
                    println!("Found vertical reflection at {}", i);
                    sum += i;
                }
                None => println!("No vertical reflection found"),
            }
        }

        format!("{}", sum)
    }

    fn part_two(&self, input: &str) -> String {
        format!("{}", "Part two not yet implemented.")
    }
}

fn reflection(pattern: &Vec<String>) -> Option<usize> {
    'outer: for i in 0..pattern.len() - 1 {
        println!("i {} min {} ", i, cmp::min(i, pattern.len() - i - 1));
        for j in 0..cmp::min(i + 1, pattern.len() - i - 1) {
            println!(
                "rows {} {}\n{}\n{}",
                i - j,
                i + j + 1,
                pattern[i - j],
                pattern[i + j + 1]
            );
            if pattern[i - j] != pattern[i + j + 1] {
                continue 'outer;
            }
        }
        return Some(i + 1);
    }
    None
}

fn transpose(pattern: &[String]) -> Vec<String> {
    let mut lines = vec![];
    for x in 0..pattern.first().expect("no rows").len() {
        let mut row = String::new();
        for line in pattern.iter().rev() {
            row.push(line.chars().nth(x).expect("no char"));
        }
        lines.push(row)
    }
    lines
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
";

    #[test]
    fn test_part1() {
        let p = Day13 {};
        assert_eq!(p.part_one(INPUT), "405");
    }

    #[test]
    fn test_part2() {
        let p = Day13 {};
        assert_eq!(p.part_two(INPUT), "todo");
    }
}
