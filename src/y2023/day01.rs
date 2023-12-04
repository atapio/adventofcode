use crate::problem::Problem;
use nom::{
    branch::alt,
    bytes::complete::{tag, take},
    character::complete::anychar,
    combinator::{value, verify},
    IResult,
};

pub struct Day01 {}

impl Problem for Day01 {
    fn part_one(&self, input: &str) -> String {
        let mut calibration_values: Vec<u32> = vec![];
        for line in input.lines() {
            let mut digits: Vec<u8> = vec![];
            for c in line.chars() {
                if let Some(d) = c.to_digit(10) {
                    //println!("{}: {} {}", line, c, d);
                    digits.push(d as u8);
                }
            }

            calibration_values.push((10 * digits.first().unwrap() + digits.last().unwrap()) as u32)
        }

        format!("{}", calibration_values.iter().sum::<u32>())
    }

    // 55330 too low
    fn part_two(&self, input: &str) -> String {
        let mut calibration_values: Vec<u32> = vec![];
        for line in input.lines() {
            let first = line_to_first_number(line).expect("first number failed");
            let last = line_to_last_number(line).expect("last number failed");
            println!("{} {} {}", line, first, last);

            let val = 10 * first + last;

            calibration_values.push(val as u32);
        }

        format!("{}", calibration_values.iter().sum::<u32>())
    }
}

//fn line_to_numbers(input: &str) -> IResult<&str, Vec<u32>> {
fn line_to_first_number(input: &str) -> Result<u32, &str> {
    let mut remaining_input = input;
    while !remaining_input.is_empty() {
        match alt((number_string, value("x", anychar)))(remaining_input) {
            Ok((next_input, digits)) => {
                if let Some(d) = digits.chars().next().expect("empty").to_digit(10) {
                    //println!("{}: {} {}", line, c, d);
                    return Ok(d);
                }

                remaining_input = next_input;
            }
            Err(_) => todo!(), //Err() => { }
        }
    }
    Err("nothing found")
}

fn line_to_last_number(input: &str) -> Result<u32, &str> {
    let mut count = 0;
    let split_pos = input.char_indices().nth_back(count).unwrap().0;
    let mut current_slice = &input[split_pos..];
    while input.len() >= current_slice.len() {
        match number_string(current_slice) {
            Ok((_, digits)) => {
                if let Some(d) = digits.chars().next().expect("empty").to_digit(10) {
                    //println!("{}: {} {}", line, c, d);
                    return Ok(d);
                }

                count += 1;
                let split_pos = input.char_indices().nth_back(count).unwrap().0;
                current_slice = &input[split_pos..];
            }
            Err(_) => {
                count += 1;
                let split_pos = input.char_indices().nth_back(count).unwrap().0;
                current_slice = &input[split_pos..];
            }
        }
    }
    Err("nothing found")
}

fn number_string(input: &str) -> IResult<&str, &str> {
    alt((
        value("1", tag("one")),
        value("2", tag("two")),
        value("3", tag("three")),
        value("4", tag("four")),
        value("5", tag("five")),
        value("6", tag("six")),
        value("7", tag("seven")),
        value("8", tag("eight")),
        value("9", tag("nine")),
        verify(take(1usize), |s: &str| s.parse::<u8>().is_ok()),
    ))(input)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let p = Day01 {};
        let input = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";
        assert_eq!(p.part_one(input), "142");
    }

    #[test]
    fn test_part2() {
        let p = Day01 {};
        let input = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";
        assert_eq!(p.part_two(input), "281");
    }
}
