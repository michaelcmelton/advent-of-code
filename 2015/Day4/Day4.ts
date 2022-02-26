/**
 * 
 * --- Day 4: The Ideal Stocking Stuffer ---
Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts
for all the economically forward-thinking little girls and boys.

To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least
five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below)
followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number
(no leading zeroes: 1, 2, 3, ...) that produces such a hash.

For example:

If your secret key is abcdef, the answer is 609043, because the MD5 hash of 
bcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.

If your secret key is pqrstuv, the lowest number it combines with to make an
MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....

Your answer was 282749.

--- Part Two ---
Now find one that starts with six zeroes.

Your answer was 9962624.
* 
 */

(() => {
  let fs = require("fs");
  let crypto = require("crypto");

  const part1 = (input: string): number => {
    let num = 0;
    while (true) {
      let hash = crypto
        .createHash("md5")
        .update(`${input}${num}`)
        .digest("hex");
      if (/^00000.*/.test(hash)) {
        break;
      } else {
        num++;
      }
    }
    return num;
  };

  const part2 = (input: string): number => {
    let num = 0;
    while (true) {
      let hash = crypto
        .createHash("md5")
        .update(`${input}${num}`)
        .digest("hex");
      if (/^000000.*/.test(hash)) {
        break;
      } else {
        num++;
      }
    }
    return num;
  };

  try {
    let input = fs.readFileSync("./input.txt", "utf-8");
    console.log(part1(input));
    console.log(part2(input));
  } catch (error) {}
})();
