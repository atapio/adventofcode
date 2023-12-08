use parse_display::{Display, FromStr};
use std::cmp::{max, min};
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
        let mut nodes = HashMap::new();
        let mut lines = input.lines();
        let instructions = lines.next().expect("no lines");
        let mut current_nodes = vec![];

        for line in lines {
            if line.is_empty() {
                continue;
            }
            let node = line.parse::<Node>().expect("bad line");
            nodes.insert(node.name.clone(), node.clone());
            if node.is_starting_node() {
                current_nodes.push(node);
            }
        }

        let mut steps = 0;
        let mut solutions = vec![];
        while !current_nodes.is_empty() {
            for i in instructions.chars() {
                let mut next_nodes = vec![];
                steps += 1;
                for current_node in current_nodes.iter() {
                    match i {
                        'L' => next_nodes
                            .push(nodes.get(&current_node.left).expect("bad left").clone()),
                        'R' => {
                            next_nodes
                                .push(nodes.get(&current_node.right).expect("bad right").clone());
                        }
                        _ => panic!("bad instruction"),
                    };
                }

                current_nodes = vec![];
                for node in next_nodes.iter() {
                    if node.is_ending_node() {
                        solutions.push(steps);
                        continue;
                    }
                    current_nodes.push(node.clone());
                }
            }

            // if steps > 10000000 {
            //     panic!("too many steps");
            // }
        }

        let mut solution = lcm(solutions[0], solutions[1]);
        for s in &solutions[2..] {
            solution = lcm(solution, *s);
        }

        println!("{:?}", solutions);

        format!("{}", solution)
    }
}

#[derive(Display, FromStr, PartialEq, Debug, Clone)]
#[display("{name} = ({left}, {right})")]
struct Node {
    name: String,
    left: String,
    right: String,
}

impl Node {
    fn is_starting_node(&self) -> bool {
        self.name.ends_with('A')
    }
    fn is_ending_node(&self) -> bool {
        self.name.ends_with('Z')
    }
}

fn gcd(a: usize, b: usize) -> usize {
    match ((a, b), (a & 1, b & 1)) {
        ((x, y), _) if x == y => y,
        ((0, x), _) | ((x, 0), _) => x,
        ((x, y), (0, 1)) | ((y, x), (1, 0)) => gcd(x >> 1, y),
        ((x, y), (0, 0)) => gcd(x >> 1, y >> 1) << 1,
        ((x, y), (1, 1)) => {
            let (x, y) = (min(x, y), max(x, y));
            gcd((y - x) >> 1, x)
        }
        _ => unreachable!(),
    }
}

fn lcm(a: usize, b: usize) -> usize {
    a * b / gcd(a, b)
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

    #[test]
    fn test_part2_example_2() {
        let input: &str = "LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
";
        let p = Day08 {};
        assert_eq!(p.part_two(input), "6");
    }
}
