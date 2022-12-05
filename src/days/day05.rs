use std::fmt::{Display, Error, Formatter};

use crate::problem::Problem;
use nom::{
    branch::alt,
    bytes::complete::tag,
    character::complete::{anychar, i8, newline, space1},
    character::complete::{char, space0},
    error::Error as NomError,
    multi::separated_list1,
    sequence::{delimited, preceded, tuple},
    Finish, IResult,
};

pub struct Day05 {}

impl Problem for Day05 {
    fn part_one(&self, input: &str) -> String {
        match self.parse(input) {
            Ok((_, (mut stacks, instructions))) => {
                for i in instructions {
                    stacks.execute(&i);
                }
                format!("{}", stacks.top())
            }
            Err(e) => format!("Error: {:?}", e),
        }
    }

    fn part_two(&self, input: &str) -> String {
        format!("{}", "Part two not yet implemented.")
    }
}

impl Day05 {
    fn parse<'a>(
        &'a self,
        input: &'a str,
    ) -> Result<(&str, (Stacks, Vec<Instruction>)), NomError<&str>> {
        tuple((
            stack_rows,
            newline,
            stack_numbers,
            space0,
            newline,
            newline,
            instructions,
            newline,
        ))(input)
        .map(|(next_input, res)| {
            let (stack_rows, _, _stack_numbers, _, _, _, instructions, _) = res;
            let mut stacks = Stacks { stacks: vec![] };

            for row in stack_rows {
                stacks.insert_row(row.to_vec());
            }

            //stack_rows.iter().map(|row| stacks.insert_row(row.to_vec()));
            (next_input, (stacks, instructions))
        })
        .finish()
    }
}

fn stack_row(input: &str) -> IResult<&str, Vec<char>> {
    separated_list1(tag(" "), empty_or_crate)(input)
}

fn stack_rows(input: &str) -> IResult<&str, Vec<Vec<char>>> {
    separated_list1(newline, stack_row)(input)
}

fn stack_numbers(input: &str) -> IResult<&str, Vec<i8>> {
    preceded(char(' '), separated_list1(space1, i8))(input)
}

fn instruction(input: &str) -> IResult<&str, Instruction> {
    tuple((tag("move "), i8, tag(" from "), i8, tag(" to "), i8))(input).map(|(next_input, res)| {
        let (_, count, _, from, _, to) = res;
        (
            next_input,
            Instruction {
                count: count.try_into().unwrap(),
                from: from.try_into().unwrap(),
                to: to.try_into().unwrap(),
            },
        )
    })
}

fn instructions(input: &str) -> IResult<&str, Vec<Instruction>> {
    separated_list1(newline, instruction)(input)
}

fn empty_or_crate(input: &str) -> IResult<&str, char> {
    alt((three_spaces, one_crate))(input)
}

fn three_spaces(input: &str) -> IResult<&str, char> {
    delimited(char(' '), char(' '), char(' '))(input)
}
fn one_crate(input: &str) -> IResult<&str, char> {
    delimited(char('['), anychar, char(']'))(input)
}

#[derive(Debug)]
struct Stacks {
    stacks: Vec<Stack>,
}

impl Stacks {
    fn insert_row(&mut self, row: Vec<char>) {
        while self.stacks.len() < row.len() {
            self.stacks.push(Stack::new());
        }

        for (i, c) in row.iter().enumerate() {
            self.stacks[i].insert(*c);
        }
    }

    fn execute(&mut self, instruction: &Instruction) {
        let mut crates: Vec<char> = vec![];
        {
            let from = &mut self.stacks[instruction.from as usize - 1];

            for _ in 0..instruction.count {
                crates.push(from.items.pop().expect("empty stack"));
            }
        }

        let to = &mut self.stacks[instruction.to as usize - 1];

        for c in crates {
            to.items.push(c);
        }
    }

    fn top(&self) -> String {
        self.stacks.iter().fold("".to_string(), |crates, c| {
            format!("{}{}", crates, c.items.last().expect("empty"))
        })
    }
}

impl Display for Stacks {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result<(), Error> {
        for stack in self.stacks.iter() {
            writeln!(f, "{}", stack)?;
        }
        Ok(())
    }
}

#[derive(Debug, PartialEq)]
struct Instruction {
    from: u8,
    to: u8,
    count: u8,
}

#[derive(Debug)]
struct Stack {
    items: Vec<char>,
}

impl Stack {
    fn new() -> Stack {
        return Stack { items: vec![] };
    }

    fn insert(&mut self, item: char) {
        // insert to bottom of stack
        if item != ' ' {
            self.items.insert(0, item);
        }
    }
}

impl Display for Stack {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result<(), Error> {
        for i in self.items.iter() {
            writeln!(f, "{}", i)?;
        }
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
";

    #[test]
    fn test_part1() {
        let p = Day05 {};
        assert_eq!(p.part_one(INPUT), "CMZ");
    }

    #[test]
    fn test_part2() {
        let p = Day05 {};
        assert_eq!(p.part_two(INPUT), "todo");
    }

    #[test]
    fn test_parse_stack_row() {
        assert_eq!(
            stack_row("    [D]     [A]"),
            Ok(("", vec![' ', 'D', ' ', 'A']))
        );
    }

    #[test]
    fn test_parse_stack_numbers() {
        assert_eq!(stack_numbers(" 1   2   3"), Ok(("", vec![1, 2, 3])));
    }

    #[test]
    fn test_parse_instruction() {
        assert_eq!(
            instruction("move 1 from 2 to 1"),
            Ok((
                "",
                Instruction {
                    count: 1,
                    from: 2,
                    to: 1
                }
            ))
        );
    }
}
