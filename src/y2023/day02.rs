use nom::{
    branch::alt,
    bytes::complete::tag,
    character::complete::space1,
    character::complete::{i32, newline},
    multi::{separated_list0, separated_list1},
    sequence::{delimited, pair, separated_pair, tuple},
    IResult,
};

use crate::problem::Problem;

pub struct Day02 {}

impl Problem for Day02 {
    fn part_one(&self, input: &str) -> String {
        let games = Day02::parse(input).expect("Failed to parse games");
        let mut sum = 0;
        for game in games {
            if game.possible() {
                sum += game.id;
            }
        }
        format!("{}", sum)
    }

    fn part_two(&self, input: &str) -> String {
        let games = Day02::parse(input).expect("Failed to parse games");
        let mut sum = 0;
        for game in games {
            let subset = game.min_subset();
            sum += subset.red * subset.green * subset.blue;
        }
        format!("{}", sum)
    }
}

impl Day02 {
    fn parse(input: &str) -> Result<Vec<Game>, nom::Err<nom::error::Error<&str>>> {
        let mut games: Vec<Game> = vec![];
        for line in input.lines() {
            if line.is_empty() {
                continue;
            }
            let (_, game) = game(line).expect("Failed to parse game");
            games.push(game);
        }
        Ok(games)
    }
}

fn game(input: &str) -> IResult<&str, Game> {
    tuple((
        delimited(pair(tag("Game"), space1), i32, tag(": ")),
        separated_list0(tag("; "), subset),
    ))(input)
    .map(|(next_input, (id, subsets))| {
        (
            next_input,
            Game {
                id: id as u32,
                subsets,
            },
        )
    })
}

fn subset(input: &str) -> IResult<&str, Subset> {
    separated_list1(tag(", "), color)(input).map(|(next_input, colors)| {
        let mut subset = Subset {
            red: 0,
            green: 0,
            blue: 0,
        };
        for (count, color) in colors {
            match color {
                "red" => subset.red = count as u32,
                "green" => subset.green = count as u32,
                "blue" => subset.blue = count as u32,
                _ => panic!("invalid color"),
            };
        }
        (next_input, subset)
    })
}

fn color(input: &str) -> IResult<&str, (i32, &str)> {
    separated_pair(i32, space1, alt((tag("red"), tag("green"), tag("blue"))))(input)
}

#[derive(Debug, Clone)]
struct Game {
    id: u32,
    subsets: Vec<Subset>,
}

impl Game {
    fn possible(&self) -> bool {
        for subset in self.subsets.iter() {
            if !subset.possible() {
                return false;
            }
        }
        true
    }

    fn min_subset(&self) -> Subset {
        let mut min_subset = Subset {
            red: 0,
            green: 0,
            blue: 0,
        };
        for subset in self.subsets.iter() {
            if subset.red > min_subset.red {
                min_subset.red = subset.red;
            }
            if subset.green > min_subset.green {
                min_subset.green = subset.green;
            }
            if subset.blue > min_subset.blue {
                min_subset.blue = subset.blue;
            }
        }
        min_subset
    }
}

#[derive(Debug, Clone)]
struct Subset {
    red: u32,
    green: u32,
    blue: u32,
}

impl Subset {
    fn possible(&self) -> bool {
        self.red <= 12 && self.green <= 13 && self.blue <= 14
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
";

    #[test]
    fn test_part1() {
        let p = Day02 {};
        assert_eq!(p.part_one(INPUT), "8");
    }

    #[test]
    fn test_part2() {
        let p = Day02 {};
        assert_eq!(p.part_two(INPUT), "2286");
    }
}
