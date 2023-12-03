use core::fmt;

use crate::problem::Problem;

pub struct Day10 {}

impl Problem for Day10 {
    fn part_one(&self, input: &str) -> String {
        let mut c = Computer::new();
        for line in input.lines() {
            let parts = line.split(' ').collect::<Vec<&str>>();
            match parts[0] {
                "addx" => c.add(parts[1].parse::<i32>().unwrap()),
                "noop" => c.noop(),
                _ => {
                    panic!("Unknown instruction: {}", parts[0])
                }
            }
        }

        let signal_strength = vec![20, 60, 100, 140, 180, 220]
            .iter()
            .fold(0, |sum, cycle| sum + c.signal_strength(*cycle));

        format!("{}", signal_strength)
    }

    fn part_two(&self, input: &str) -> String {
        let mut c = Computer::new();
        for line in input.lines() {
            let parts = line.split(' ').collect::<Vec<&str>>();
            match parts[0] {
                "addx" => c.add(parts[1].parse::<i32>().unwrap()),
                "noop" => c.noop(),
                _ => {
                    panic!("Unknown instruction: {}", parts[0])
                }
            }
        }
        format!("\n{}", c)
    }
}

#[derive(Debug)]
struct Computer {
    register_x: Vec<i32>,
}

impl Computer {
    fn new() -> Self {
        Self {
            register_x: vec![1],
        }
    }

    fn signal_strength(&self, cycle: usize) -> i32 {
        self.register_x.get(cycle - 1).unwrap().to_owned() * cycle as i32
    }

    fn noop(&mut self) {
        self.register_x
            .push(self.register_x.last().unwrap().clone());
    }

    fn add(&mut self, x: i32) {
        let last = self.register_x.last().unwrap().clone();
        self.register_x.push(last);
        self.register_x.push(last + x);
    }
}

impl fmt::Display for Computer {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        for i in 0..self.register_x.len() - 1 {
            let pos = i % 40;
            let signal = self.register_x[i];

            if signal == pos as i32 || signal == (pos as i32 - 1) || (signal == pos as i32 + 1) {
                write!(f, "#")?;
            } else {
                write!(f, ".")?;
            }

            if i > 0 && pos % 40 == 39 {
                write!(f, "\n")?;
            }
        }
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
";

    #[test]
    fn test_part1() {
        let p = Day10 {};
        assert_eq!(p.part_one(INPUT), "13140");
    }

    #[test]
    fn test_part2() {
        let p = Day10 {};
        assert_eq!(
            p.part_two(INPUT),
            "##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
"
        );
    }
}
