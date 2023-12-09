use itertools::Itertools;

use crate::problem::Problem;

pub struct Day09 {}

impl Problem for Day09 {
    fn part_one(&self, input: &str) -> String {
        let mut histories = vec![];

        for line in input.lines() {
            if line.is_empty() {
                continue;
            }
            let mut history = vec![];
            for v in line.split(' ') {
                history.push(v.parse::<i32>().expect("bad number"));
            }
            histories.push(history);
        }
        //println!("{:?}", histories);

        let mut sum = 0;
        for history in histories {
            sum += next_value(&history);
        }

        format!("{}", sum)
    }

    fn part_two(&self, input: &str) -> String {
        let mut histories = vec![];

        for line in input.lines() {
            if line.is_empty() {
                continue;
            }
            let mut history = vec![];
            for v in line.split(' ') {
                history.push(v.parse::<i32>().expect("bad number"));
            }
            histories.push(history);
        }
        //println!("{:?}", histories);

        let mut sum = 0;
        for history in histories {
            let mut rev_history: Vec<i32> = history.clone();
            rev_history.reverse();
            sum += next_value(&rev_history);
        }

        format!("{}", sum)
    }
}

fn next_value(history: &[i32]) -> i32 {
    let mut rows = vec![];
    let mut row: Vec<i32> = Vec::from(history);
    rows.push(row.clone());
    'outer: loop {
        let mut diffs = vec![];
        for (a, b) in row.iter().tuple_windows() {
            diffs.push(*b - *a);
        }

        rows.push(diffs.clone());
        row = diffs.clone();

        for d in diffs {
            if d != 0 {
                continue 'outer;
            }
        }
        // all zeros
        break;
    }

    //println!("rows: {:?}", rows);

    for j in 1..rows.len() {
        let i = rows.len() - j;
        let v = rows[i].last().expect("no last") + rows[i - 1].last().expect("no last");
        //println!("i: {} v: {}", i, v);
        rows[i - 1].push(v);
    }

    let last = *rows.first().expect("no first").last().expect("no last");
    //println!("history: {:?} last: {}", history, last);
    last
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
";

    #[test]
    fn test_part1() {
        let p = Day09 {};
        assert_eq!(p.part_one(INPUT), "114");
    }

    #[test]
    fn test_part2() {
        let p = Day09 {};
        assert_eq!(p.part_two(INPUT), "2");
    }
}
