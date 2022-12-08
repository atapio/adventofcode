use crate::problem::Problem;

pub struct Day08 {}

impl Problem for Day08 {
    fn part_one(&self, input: &str) -> String {
        let grid = Self::parse(input);
        format!("{}", Self::count_visible(&grid))
    }

    fn part_two(&self, input: &str) -> String {
        let grid = Self::parse(input);
        format!("{}", Self::max_scenic_score(&grid))
    }
}

impl Day08 {
    fn parse(input: &str) -> Grid {
        let cells: Vec<Vec<u8>> = input
            .lines()
            .map(|line| {
                line.chars()
                    .map(|d| d.to_digit(10).unwrap() as u8)
                    .collect()
            })
            .collect();

        let mut grid = Grid::new(cells.len(), cells.get(0).unwrap().len());
        for (y, line) in cells.iter().enumerate() {
            for (x, cell) in line.iter().enumerate() {
                grid.set(x, y, *cell);
            }
        }

        grid
    }

    fn count_visible(grid: &Grid) -> u32 {
        let edges: u32 = u32::try_from(grid.height * 2 + grid.width * 2 - 4).unwrap();

        let mut visible = Grid::new(grid.width, grid.height);

        // left to right
        for y in 1..grid.height - 1 {
            let mut max_height = grid.get(0, y);
            for x in 1..grid.width - 1 {
                let height = grid.get(x, y);
                if height > max_height {
                    // println!("{} {} {} {}", y, x, height, max_height);
                    visible.set(x, y, height);
                    max_height = height;
                }
            }
        }
        // right to left
        for y in (1..grid.height - 1).rev() {
            let mut max_height = grid.get(grid.width - 1, y);
            for x in (1..grid.width - 1).rev() {
                let height = grid.get(x, y);
                if height > max_height {
                    // println!("{} {} {} {}", y, x, height, max_height);
                    visible.set(x, y, height);
                    max_height = height;
                }
            }
        }
        // up to down
        for x in 1..grid.width - 1 {
            let mut max_height = grid.get(x, 0);
            for y in 1..grid.height - 1 {
                let height = grid.get(x, y);
                if height > max_height {
                    // println!("{} {} {} {}", y, x, height, max_height);
                    visible.set(x, y, height);
                    max_height = height;
                }
            }
        }
        // bottop to up
        for x in (1..grid.width - 1).rev() {
            let mut max_height = grid.get(x, grid.height - 1);
            for y in (1..grid.height - 1).rev() {
                let height = grid.get(x, y);
                if height > max_height {
                    // println!("{} {} {} {}", y, x, height, max_height);
                    visible.set(x, y, height);
                    max_height = height;
                }
            }
        }

        return edges + visible.count(1);
    }

    fn max_scenic_score(grid: &Grid) -> u32 {
        let mut max_score = 0;
        for y in 1..grid.height - 1 {
            for x in 1..grid.width - 1 {
                let score = Self::scenic_score(grid, x, y);
                if score > max_score {
                    max_score = score;
                }
            }
        }

        max_score
    }

    fn scenic_score(grid: &Grid, tree_x: usize, tree_y: usize) -> u32 {
        let mut left = 0;
        let mut right = 0;
        let mut up = 0;
        let mut down = 0;

        let tree_height = grid.get(tree_x, tree_y);
        // println!("tree: {} {} {}", tree_x, tree_y, tree_height);

        // looking right
        for x in tree_x + 1..grid.width {
            right += 1;
            // println!(
            //     "right: x:{} y:{} count: {} height: {} tree: {}",
            //     x,
            //     tree_y,
            //     left,
            //     grid.get(x, tree_y),
            //     tree_height,
            // );
            if grid.get(x, tree_y) >= tree_height {
                break;
            }
        }
        // looking left
        for x in (0..tree_x).rev() {
            left += 1;
            // println!(
            //     "left: x:{} y:{} count: {} height: {} tree: {}",
            //     x,
            //     tree_y,
            //     left,
            //     grid.get(x, tree_y),
            //     tree_height,
            // );
            if grid.get(x, tree_y) >= tree_height {
                break;
            }
        }
        // looking down
        for y in tree_y + 1..grid.height {
            down += 1;
            if grid.get(tree_x, y) >= tree_height {
                break;
            }
        }
        // looking up
        for y in (0..tree_y).rev() {
            up += 1;
            if grid.get(tree_x, y) >= tree_height {
                break;
            }
        }

        // println!(
        //     "x {} y {}: {} left {} right {} up {} down {}",
        //     tree_x,
        //     tree_y,
        //     left * right * up * down,
        //     left,
        //     right,
        //     up,
        //     down
        // );

        return left * right * up * down;
    }
}

#[derive(Debug)]
struct Grid {
    pub width: usize,
    pub height: usize,
    cells: Vec<u8>,
}

impl Grid {
    fn new(width: usize, height: usize) -> Self {
        Self {
            width,
            height,
            cells: vec![0; width * height],
        }
    }

    fn get(&self, x: usize, y: usize) -> u8 {
        self.cells[y * self.width + x]
    }

    fn set(&mut self, x: usize, y: usize, value: u8) {
        self.cells[y * self.width + x] = value;
    }

    fn count(&self, min_value: u8) -> u32 {
        self.cells.iter().fold(
            0,
            |total, &c| if c >= min_value { total + 1 } else { total },
        )
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    const INPUT: &str = "30373
25512
65332
33549
35390";

    #[test]
    fn test_parse() {
        let grid = Day08::parse(INPUT);
        assert_eq!(grid.height, 5);
        assert_eq!(grid.width, 5);
    }

    #[test]
    fn test_part1() {
        let p = Day08 {};
        assert_eq!(p.part_one(INPUT), "21");
    }

    #[test]
    fn test_part2() {
        let p = Day08 {};
        assert_eq!(p.part_two(INPUT), "8");
    }
}
