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

        self.elves.sort_unstable();
    }

    fn sum_of_max_calories(self, n: usize) -> u32 {
        return self.elves[self.elves.len() - n..].iter().sum();
    }
}

impl Problem for Day01 {
    fn part_one(&self, input: &str) -> String {
        let mut elves = Elves { elves: vec![] };
        elves.parse(input);
        return elves.sum_of_max_calories(1).to_string();
    }

    fn part_two(&self, input: &str) -> String {
        let mut elves = Elves { elves: vec![] };
        elves.parse(input);
        return elves.sum_of_max_calories(3).to_string();
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
