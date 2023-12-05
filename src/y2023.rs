use crate::problem::Problem;

mod day01;
mod day04;

pub fn day_to_problem(day: u8) -> Option<Box<dyn Problem>> {
    match day {
        1 => Some(Box::new(day01::Day01 {})),
        4 => Some(Box::new(day04::Day04 {})),
        // ...
        _ => None,
    }
}
