use crate::problem::Problem;

pub struct Day01 {}

pub struct Elves {
    elves: Vec<u32>,
}

impl Elves {
    fn parse(&mut self, input: &str) {
        let input = input.split("\n");
        let mut elf = 0;
        for line in input {
            if line == "" {
                self.elves.push(elf);
                elf = 0;
                continue;
            }
            elf += line.parse::<u32>().unwrap();
        }
        self.elves.push(elf);
    }

    fn with_max_calories(&self) -> (usize, u32) {
        let mut idx = 0;
        let mut max = 0;
        for (i, cal) in self.elves.iter().enumerate() {
            if cal > &max {
                max = *cal;
                idx = i;
            }
        }

        return (idx, max);
    }

    fn max_calories(&self) -> u32 {
        return self.with_max_calories().1;
    }
}

impl Problem for Day01 {
    fn part_one(&self, input: &str) -> String {
        let mut elves = Elves { elves: vec![] };
        elves.parse(input);
        return elves.max_calories().to_string();
    }

    fn part_two(&self, input: &str) -> String {
        let mut elves = Elves { elves: vec![] };
        elves.parse(input);
        let mut max = 0;

        for _ in 1..4 {
            let (i, cal) = elves.with_max_calories();
            max += cal;
            elves.elves.remove(i);
        }

        return max.to_string();
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let p = Day01 {};
        let input = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";
        assert_eq!(p.part_one(&input), "24000");
    }

    #[test]
    fn test_part2() {
        let p = Day01 {};
        let input = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";
        assert_eq!(p.part_two(&input), "45000");
    }
}
