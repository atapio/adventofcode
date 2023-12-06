use parse_display::Display;

use nom::bytes::complete::tag;
use nom::character::complete::space0;
use nom::sequence::{delimited, pair, tuple};
use nom::{character::complete::i32, character::complete::space1, multi::separated_list1, IResult};

use crate::problem::Problem;

pub struct Day04 {}

impl Problem for Day04 {
    fn part_one(&self, input: &str) -> String {
        let cards = Day04::parse(input).expect("Failed to parse cards");
        //println!("cards {:?}", cards);
        let mut total = 0;
        for card in cards {
            total += card.points();
        }
        format!("{}", total)
    }

    fn part_two(&self, input: &str) -> String {
        let mut cards = Day04::parse(input).expect("Failed to parse cards");
        //println!("cards {:?}", cards);
        for i in 0..cards.len() {
            let card = &cards[i];
            let matches = card.matches();
            let copies = card.copies;

            for j in 0..(matches as usize) {
                if i + j + 1 >= cards.len() {
                    break;
                }
                cards[i + j + 1].copies += copies;
            }
        }
        //println!("{:?}", cards);
        let mut total = 0;
        cards.iter().for_each(|c| {
            total += c.copies;
        });
        format!("{}", total)
    }
}

impl Day04 {
    fn parse(input: &str) -> Result<Vec<Card>, nom::Err<nom::error::Error<&str>>> {
        let mut cards: Vec<Card> = vec![];
        for line in input.lines() {
            let card = card(line).expect("Failed to parse card");
            cards.push(card);
        }

        Ok(cards)
    }
}

fn card(input: &str) -> Result<Card, nom::Err<nom::error::Error<&str>>> {
    tuple((
        card_number,
        space0,
        numbers,
        space0,
        tag("|"),
        space0,
        numbers,
    ))(input)
    .map(|(_, res)| {
        let (card_number, _, winning_numbers, _, _, _, numbers) = res;

        Card {
            id: card_number,
            copies: 1,
            winning_numbers,
            numbers,
        }
    })
}

fn card_number(input: &str) -> IResult<&str, i32> {
    delimited(pair(tag("Card"), space1), i32, tag(":"))(input)
}

fn numbers(input: &str) -> IResult<&str, Vec<i32>> {
    separated_list1(space1, i32)(input)
}

#[derive(Debug)]
struct Card {
    id: i32,
    copies: u32,
    winning_numbers: Vec<i32>,
    numbers: Vec<i32>,
}

impl Card {
    fn points(&self) -> u32 {
        match self.matches() {
            0 => 0,
            i => 2u32.pow(i - 1),
        }
    }
    fn matches(&self) -> u32 {
        self.winning_numbers
            .iter()
            .filter(|n| self.numbers.contains(n))
            .count() as u32
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
";

    #[test]
    fn test_part1() {
        let p = Day04 {};
        assert_eq!(p.part_one(INPUT), "13");
    }

    #[test]
    fn test_part2() {
        let p = Day04 {};
        assert_eq!(p.part_two(INPUT), "30");
    }
}
