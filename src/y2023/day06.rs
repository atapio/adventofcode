use nom::bytes::complete::tag;
use nom::character::complete::newline;
use nom::sequence::{delimited, pair};
use nom::{character::complete::i32, character::complete::space1, multi::separated_list1, IResult};

use crate::problem::Problem;

pub struct Day06 {}

impl Problem for Day06 {
    fn part_one(&self, input: &str) -> String {
        let races = Day06::parse(input).expect("Failed to parse cards");
        //println!("cards {:?}", cards);
        let mut total = 1;
        for race in races {
            total *= race.ways_to_break_the_record();
        }
        format!("{}", total)
    }

    fn part_two(&self, input: &str) -> String {
        let races = Day06::parse(input).expect("Failed to parse cards");

        let mut records: Vec<i128> = vec![];
        let mut times: Vec<i128> = vec![];

        for race in races {
            records.push(race.record);
            times.push(race.duration);
        }

        let race = Race {
            record: Day06::concat_ints(records),
            duration: Day06::concat_ints(times),
        };

        let result = race.ways_to_break_the_record();

        format!("{}", result)
    }
}

impl Day06 {
    fn parse(input: &str) -> Result<Vec<Race>, nom::Err<nom::error::Error<&str>>> {
        let mut races: Vec<Race> = vec![];

        let (_, (times, distances)) = parse_input(input).expect("Failed to parse input");

        times.iter().zip(distances).for_each(|(t, d)| {
            races.push(Race {
                duration: *t as i128,
                record: d as i128,
            });
        });

        Ok(races)
    }

    fn concat_ints(ints: Vec<i128>) -> i128 {
        let mut s = String::new();
        ints.iter().for_each(|i| {
            s.push_str(&i.to_string());
        });
        s.parse::<i128>().unwrap()
    }
}

fn parse_input(input: &str) -> IResult<&str, (Vec<i32>, Vec<i32>)> {
    pair(
        delimited(pair(tag("Time:"), space1), numbers, newline),
        delimited(pair(tag("Distance:"), space1), numbers, newline),
    )(input)
}

fn numbers(input: &str) -> IResult<&str, Vec<i32>> {
    separated_list1(space1, i32)(input)
}

#[derive(Debug)]
struct Race {
    duration: i128,
    record: i128,
}

impl Race {
    fn ways_to_break_the_record(&self) -> i128 {
        let mut count = 0;

        for speed in 1..self.duration - 1 {
            let time_left = self.duration - speed;
            let distance = speed * time_left;

            if distance > self.record {
                count += 1;
            }
        }

        count
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "Time:      7  15   30
Distance:  9  40  200
";

    #[test]
    fn test_part1() {
        let p = Day06 {};
        assert_eq!(p.part_one(INPUT), "288");
    }

    #[test]
    fn test_part2() {
        let p = Day06 {};
        assert_eq!(p.part_two(INPUT), "71503");
    }
}
