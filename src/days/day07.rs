use std::collections::BTreeMap;

use nom::{
    branch::alt,
    bytes::complete::tag,
    character::complete::{alphanumeric1, multispace1, newline, u64},
    combinator::recognize,
    multi::{many1, separated_list0},
    sequence::{delimited, preceded, tuple},
    Err as NomError, IResult,
};
use parse_display::IntoResult;

use crate::problem::Problem;

pub struct Day07 {}

impl Problem for Day07 {
    fn part_one(&self, input: &str) -> String {
        let cwd = &mut Cwd::init();

        Self::parse(cwd, input);
        let total_size = cwd.get_root().dirs().iter().fold(0, |total, d| {
            if d.size() < 100000 {
                total + d.size()
            } else {
                total
            }
        });

        format!("{}", total_size)
    }

    fn part_two(&self, input: &str) -> String {
        format!("{}", "Part two not yet implemented.")
    }
}

impl Day07 {
    fn parse(cwd: &mut Cwd, input: &str) {
        let mut remaining_input = input;

        while remaining_input.len() > 0 {
            match cmd_cd(remaining_input) {
                Ok((next_input, dir)) => {
                    match dir {
                        "/" => {
                            cwd.root();
                        }
                        ".." => {
                            cwd.parent();
                        }
                        _ => {
                            cwd.cd(dir);
                        }
                    }
                    remaining_input = next_input;
                    println!(
                        "cwd: {} {:?}",
                        cwd.current().unwrap().name,
                        cwd.current().unwrap().children,
                    );
                }
                Err(e1) => match cmd_ls(remaining_input) {
                    Ok((next_input, entries)) => {
                        for entry in entries {
                            println!("Entry: {}", entry.name);
                            cwd.current_mut().unwrap().add(entry.clone());
                        }
                        remaining_input = next_input;
                        println!(
                            "dir: {} {:?}",
                            cwd.current().unwrap().name,
                            cwd.current().unwrap().children,
                        );
                    }
                    Err(e2) => {
                        panic!(
                            "Error:\ncd: {:?}\nls: {:?}\nremaining: {}",
                            e1, e2, remaining_input
                        );
                    }
                },
            }
        }
    }
}

fn cmd_cd(input: &str) -> IResult<&str, &str> {
    delimited(
        tag("$ cd "),
        recognize(many1(alt((tag(".."), alphanumeric1, tag("/"))))),
        newline,
    )(input)
}

fn cmd_ls(input: &str) -> IResult<&str, Vec<DirEntry>> {
    tuple((
        tag("$ ls"),
        newline,
        separated_list0(multispace1, direntry),
        newline,
    ))(input)
    .map(|(next_input, res)| {
        let (_, _, direntry, _) = res;
        (next_input, direntry)
    })
}
fn direntry<'a>(input: &'a str) -> IResult<&str, DirEntry> {
    alt((direntry_dir, direntry_file))(input)
}

//fn direntry_dir<'a>(input: &str) -> IResult<&str, &DirEntry<'a>> {
fn direntry_dir(input: &str) -> IResult<&str, DirEntry> {
    preceded(
        tag("dir "),
        recognize(many1(alt((alphanumeric1, tag("/"))))),
    )(input)
    .map(|(next_input, res)| {
        let name = res;
        (next_input, DirEntry::dir(name))
    })
}
fn direntry_file(input: &str) -> IResult<&str, DirEntry> {
    tuple((
        u64,
        tag(" "),
        recognize(many1(alt((alphanumeric1, tag("_"), tag("."))))),
    ))(input)
    .map(|(next_input, res)| {
        let (size, _, name) = res;
        (next_input, DirEntry::file(name, size))
    })
}

#[derive(Debug)]
struct Cwd {
    pub path: Vec<DirEntry>,
}

impl Cwd {
    fn init<'a>() -> Cwd {
        Cwd {
            path: vec![DirEntry::dir("/")],
        }
    }

    fn root(&mut self) {
        self.path.truncate(1);
    }

    fn get_root(&self) -> &DirEntry {
        &self.path[0]
    }

    fn parent(&mut self) {
        let cur = self.path.pop().unwrap();
        self.current_mut().unwrap().update(cur);
    }

    fn cd(&mut self, dir: &str) {
        //let last = &self.path[self.path.len() - 1];
        match self.path.last().unwrap().cd(dir) {
            //match last.cd(dir) {
            Ok(newcwd) => self.path.push(newcwd.clone()),
            Err(e) => panic!("Error: {} {}", e, dir),
        };
    }
    fn current_mut(&mut self) -> Option<&mut DirEntry> {
        self.path.last_mut()
    }
    fn current(&self) -> Option<&DirEntry> {
        self.path.last()
    }

    // fn ls(&self) {
    //     for e in self.path.last().iter() {
    //         if e.is_dir {
    //             println!("dir {}", e.name);
    //         } else {
    //             println!("{} {}", e.size, e.name);
    //         }
    //     }
    // }
}

#[derive(Debug, Clone, PartialEq)]
struct DirEntry {
    pub is_dir: bool,
    pub name: String,
    pub size: u64,
    pub children: BTreeMap<String, usize>,
    pub child_entries: Vec<DirEntry>,
}

impl DirEntry {
    fn dir(name: &str) -> DirEntry {
        DirEntry {
            is_dir: true,
            name: name.to_string(),
            size: 0,
            children: BTreeMap::new(),
            child_entries: vec![],
        }
    }
    fn file<'a>(name: &str, size: u64) -> DirEntry {
        DirEntry {
            is_dir: false,
            name: name.to_string(),
            size,
            children: BTreeMap::new(),
            child_entries: vec![],
        }
    }

    fn size(&self) -> u64 {
        match self.is_dir {
            true => self.child_entries.iter().map(|c| c.size()).sum(),
            false => self.size,
        }
    }
    //fn cd(&mut self, dir: &str) -> Result<&mut &DirEntry, &str> {
    fn cd(&self, dir: &str) -> Result<&DirEntry, &str> {
        match self.is_dir {
            true => {
                let idx = self.children.get(dir).unwrap();
                return self.child_entries.get(*idx).ok_or("dir not found");
            }

            false => Err("not a directory"),
        }
    }

    fn add(&mut self, entry: DirEntry) {
        self.child_entries.push(entry.clone());
        self.children
            .insert(entry.name.clone(), self.child_entries.len() - 1);
    }
    fn update(&mut self, entry: DirEntry) {
        let idx = self.children.get(&entry.name).unwrap();
        self.child_entries[*idx] = entry.clone();
    }

    // return recursive list of directories
    fn dirs(&self) -> Vec<&DirEntry> {
        let mut dirs = vec![];
        if self.is_dir {
            dirs.push(self);
            for child in &self.child_entries {
                println!("child: {:?}", child.name);
                dirs.extend(child.dirs().iter());
            }
            println!("{:?}: {:?}", self.name, dirs);
        }
        dirs
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
";

    #[test]
    fn test_part1() {
        let p = Day07 {};
        assert_eq!(p.part_one(INPUT), "95437");
    }

    #[test]
    fn test_part2() {
        let p = Day07 {};
        assert_eq!(p.part_two(INPUT), "24933642");
    }

    #[test]
    fn test_direntry() {
        assert_eq!(direntry("dir a"), Ok(("", DirEntry::dir("a"))));
        assert_eq!(
            direntry("14848514 b.txt"),
            Ok(("", DirEntry::file("b.txt", 14848514)))
        );
    }
}
