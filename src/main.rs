use adventofcode::problem::Problem;
use adventofcode::y2022;
use adventofcode::y2023;

use clap::Parser;

/// Simple program to run one day of AOC
#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    year: u16,
    day: u8,
}

fn day_to_problem(year: u16, day: u8) -> Option<Box<dyn Problem>> {
    match year {
        2022 => y2022::day_to_problem(day),
        2023 => y2023::day_to_problem(day),
        _ => None,
    }
}

fn main() {
    let args = Args::parse();
    let input = std::fs::read_to_string(format!("input/{}/{:02}/input.txt", args.year, args.day))
        .expect("Input file not found");

    let problem = day_to_problem(args.year, args.day).unwrap();

    println!("part1: {}", problem.part_one(&input));
    println!("part2: {}", problem.part_two(&input));
}
