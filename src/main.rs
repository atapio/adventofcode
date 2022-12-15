use adventofcode::days::day01::Day01;
use adventofcode::days::day02::Day02;
use adventofcode::days::day03::Day03;
use adventofcode::days::day04::Day04;
use adventofcode::days::day05::Day05;
use adventofcode::days::day06::Day06;
use adventofcode::days::day07::Day07;
use adventofcode::days::day08::Day08;
use adventofcode::days::day09::Day09;
use adventofcode::days::day10::Day10;
use adventofcode::days::day13::Day13;
use adventofcode::days::day14::Day14;
use adventofcode::problem::Problem;

use clap::Parser;

/// Simple program to run one day of AOC
#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    year: u16,
    day: u8,
}

fn day_to_problem(day: u8) -> Option<Box<dyn Problem>> {
    match day {
        1 => Some(Box::new(Day01 {})),
        2 => Some(Box::new(Day02 {})),
        3 => Some(Box::new(Day03 {})),
        4 => Some(Box::new(Day04 {})),
        5 => Some(Box::new(Day05 {})),
        6 => Some(Box::new(Day06 {})),
        7 => Some(Box::new(Day07 {})),
        8 => Some(Box::new(Day08 {})),
        9 => Some(Box::new(Day09 {})),
        10 => Some(Box::new(Day10 {})),
        13 => Some(Box::new(Day13 {})),
        14 => Some(Box::new(Day14 {})),
        // ...
        _ => None,
    }
}

fn main() {
    let args = Args::parse();
    let input = std::fs::read_to_string(format!("input/{}/{:02}/input.txt", args.year, args.day))
        .expect("Input file not found");

    let problem = day_to_problem(args.day).unwrap();

    println!("part1: {}", problem.part_one(&input));
    println!("part2: {}", problem.part_two(&input));
}
