/* --- Day 3: Perfectly Spherical Houses in a Vacuum ---
Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at
the North Pole calls him via radio and tells him where to move next. Moves are always exactly one
house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the
house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off,
and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

> delivers presents to 2 houses: one at the starting location, and one to the east.
^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

Answer is 2592

--- Part Two ---
The next year, to speed up the process, Santa creates a robot version of himself,
Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the same starting
house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from
the same script as the previous year.

This year, how many houses receive at least one present?

For example:

^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.



*/
(() => {
  let fs = require("fs");

  const calculateMove = (move: string): [string, number] => {
    switch (move) {
      case "v":
        return ["y", -1];
      case "^":
        return ["y", 1];
      case "<":
        return ["x", -1];
      case ">":
        return ["x", 1];
    }
  };

  const makeMove = (axis: string, number: number, coord: number[]) => {
    if (axis === "x") {
      return [coord[0] + number, coord[1]];
    } else {
      return [coord[0], coord[1] + number];
    }
  };

  const part1 = (inputs: string): number => {
    let coord = [0, 0];
    let arr: number[][] = [];
    arr.push(coord);
    inputs.split("").forEach((input) => {
      const [axis, amount] = calculateMove(input);
      coord = makeMove(axis, amount, coord);
      arr.push(coord);
    });

    return [...new Set(arr.map((i) => JSON.stringify(i)))].length;
  };

  const part2 = (inputs: string): number => {
    let coord1 = [0, 0];
    let coord2 = [0, 0];
    let arr: number[][] = [];
    arr.push(coord1);
    inputs.split("").forEach((input, index) => {
      const [axis, amount] = calculateMove(input);
      if (index % 2 === 0) {
        coord1 = makeMove(axis, amount, index % 2 === 0 ? coord1 : coord2);
      } else {
        coord2 = makeMove(axis, amount, index % 2 === 0 ? coord1 : coord2);
      }
      arr.push(index % 2 === 0 ? coord1 : coord2);
    });

    return [...new Set(arr.map((i) => JSON.stringify(i)))].length;
  };

  try {
    let input = fs.readFileSync("./input.txt", "utf-8");
    console.log(part1(input));
    console.log(part2(input));
  } catch (error) {}
})();
