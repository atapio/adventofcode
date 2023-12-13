use crate::problem::Problem;

mod day01;
mod day02;
mod day04;
mod day06;
mod day07;
mod day08;
mod day09;
mod day13;

pub fn day_to_problem(day: u8) -> Option<Box<dyn Problem>> {
    match day {
        1 => Some(Box::new(day01::Day01 {})),
        2 => Some(Box::new(day02::Day02 {})),
        4 => Some(Box::new(day04::Day04 {})),
        6 => Some(Box::new(day06::Day06 {})),
        //7 => Some(Box::new(day07::Day07 {})),
        8 => Some(Box::new(day08::Day08 {})),
        9 => Some(Box::new(day09::Day09 {})),
        13 => Some(Box::new(day13::Day13 {})),
        // ...
        _ => None,
    }
}
