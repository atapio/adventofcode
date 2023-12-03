use crate::problem::Problem;

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

    fn part_two(&self, input: &str) -> String {
        format!("{}", 0)
    }
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

    //     #[test]
    //     fn test_part2() {
    //         let p = Day01 {};
    //         let input = "1000
    // 2000
    // 3000
    //
    // 4000
    //
    // 5000
    // 6000
    //
    // 7000
    // 8000
    // 9000
    //
    // 10000";
    //         assert_eq!(p.part_two(&input), "45000");
    //     }
}
