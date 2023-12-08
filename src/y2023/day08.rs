use parse_display::{Display, FromStr};
use std::collections::HashMap;

use crate::problem::Problem;

pub struct Day08 {}

impl Problem for Day08 {
    fn part_one(&self, input: &str) -> String {
        let mut nodes = HashMap::new();

        let mut lines = input.lines();

        let instructions = lines.next().expect("no lines");

        for line in lines {
            if line.is_empty() {
                continue;
            }
            let node = line.parse::<Node>().expect("bad line");

            nodes.insert(node.name.clone(), node.clone());
        }

        //println!("{:?}", nodes);

        let mut steps = 0;

        let mut current_node = nodes.get("AAA").expect("no AAA");

        loop {
            for i in instructions.chars() {
                match i {
                    'L' => current_node = nodes.get(&current_node.left).expect("bad left"),
                    'R' => {
                        current_node = nodes.get(&current_node.right).expect("bad right");
                    }
                    _ => panic!("bad instruction"),
                }
                steps += 1;

                if current_node.name == "ZZZ" {
                    return format!("{}", steps);
                }
            }

            if steps > 100000 {
                panic!("too many steps");
            }
        }

        //format!("{}", "Part one not yet implemented.")
    }

    fn part_two(&self, input: &str) -> String {
        format!("{}", "Part two not yet implemented.")
    }
}

#[derive(Display, FromStr, PartialEq, Debug, Clone)]
#[display("{name} = ({left}, {right})")]
struct Node {
    name: String,
    left: String,
    right: String,
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
";

    #[test]
    fn test_part1() {
        let p = Day08 {};
        assert_eq!(p.part_one(INPUT), "2");
    }

    #[test]
    fn test_part1_example_2() {
        let input: &str = "LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
";
        let p = Day08 {};
        assert_eq!(p.part_one(input), "6");
    }

    #[test]
    fn test_part2() {
        let p = Day08 {};
        assert_eq!(p.part_two(INPUT), "todo");
    }
}
