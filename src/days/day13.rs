use std::cmp::Ordering;

use crate::problem::Problem;
use itertools::Itertools;
use serde_json::Value;

pub struct Day13 {}

impl Problem for Day13 {
    fn part_one(&self, input: &str) -> String {
        let pairs = Self::parse(input);

        let count = pairs.iter().fold(0, |total, pair| {
            //println!("{:?}", pair);
            match right_order(&pair.left, &pair.right) {
                Some(true) => total + pair.n,

                _ => total,
            }
        });

        format!("{}", count)
    }

    fn part_two(&self, input: &str) -> String {
        let mut pairs = Self::parse(input);
        let dividers = Self::parse(
            "[[2]]
[[6]]",
        );
        let divider_1 = Packet {
            value: dividers[0].left.clone(),
        };
        let divider_2 = Packet {
            value: dividers[0].right.clone(),
        };

        pairs.extend(dividers);
        let mut packets: Vec<Packet> = pairs
            .iter()
            .map(|p| {
                vec![
                    Packet {
                        value: p.left.clone(),
                    },
                    Packet {
                        value: p.right.clone(),
                    },
                ]
            })
            .flatten()
            .collect();

        packets.sort_unstable_by(|a, b| a.partial_cmp(b).unwrap_or(Ordering::Equal));

        let decoder = packets
            .iter()
            .enumerate()
            .filter(|(_, p)| **p == divider_1 || **p == divider_2)
            .fold(1, |total, (i, _)| {
                return total * (i + 1);
            });

        format!("{}", decoder)
    }
}

impl Day13 {
    fn parse(input: &str) -> Vec<Pair> {
        let pairs = input
            .lines()
            .chunks(3)
            .into_iter()
            .enumerate()
            .map(|(i, mut chunks)| {
                //let parsed_left = json!(chunks.next().unwrap());
                let parsed_left = serde_json::from_str(chunks.next().unwrap()).unwrap();
                let parsed_right = serde_json::from_str(chunks.next().unwrap()).unwrap();

                Pair {
                    n: i + 1,
                    left: parsed_left,
                    right: parsed_right,
                }
            })
            .collect();

        pairs
    }
}

#[derive(Debug)]
struct Pair {
    n: usize,
    left: serde_json::Value,
    right: serde_json::Value,
}

#[derive(Debug, PartialEq)]
struct Packet {
    value: serde_json::Value,
}

impl PartialOrd for Packet {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        match right_order(&self.value, &other.value) {
            Some(true) => Some(Ordering::Less),
            Some(false) => Some(Ordering::Greater),
            None => None,
        }
    }
}

fn right_order(left: &serde_json::Value, right: &serde_json::Value) -> Option<bool> {
    //println!("{:?} {:?}", left, right);
    // both values integers
    if let (Value::Number(left), Value::Number(right)) = (left, right) {
        return match left.as_u64().unwrap().cmp(&right.as_u64().unwrap()) {
            Ordering::Less => Some(true),
            Ordering::Greater => Some(false),
            Ordering::Equal => None,
        };
    }

    // both values lists
    if let (Value::Array(left), Value::Array(right)) = (left, right) {
        let mut left_iter = left.iter();
        let mut right_iter = right.iter();

        loop {
            let left = left_iter.next();
            let right = right_iter.next();

            if left.is_none() && right.is_none() {
                return None;
            }

            if left.is_none() {
                return Some(true);
            }

            if right.is_none() {
                return Some(false);
            }

            let left = left.unwrap();
            let right = right.unwrap();

            if let Some(order) = right_order(left, right) {
                return Some(order);
            }
        }
    }

    // exactly one value is an integer
    if let (Value::Number(left), _) = (left, right) {
        let l = serde_json::json!([left.as_u64().unwrap()]);
        return right_order(&l, right);
    }
    if let (_, Value::Number(right)) = (left, right) {
        let r = serde_json::json!([right.as_u64().unwrap()]);
        return right_order(left, &r);
    }

    panic!("Unexpected values: {:?} {:?}", left, right);
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]";

    #[test]
    fn test_part1() {
        let p = Day13 {};
        assert_eq!(p.part_one(INPUT), "13");
    }

    #[test]
    fn test_part2() {
        let p = Day13 {};
        assert_eq!(p.part_two(INPUT), "140");

    }
}
