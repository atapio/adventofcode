use adventofcode::days::day01::Day01;
use adventofcode::days::day02::Day02;
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
