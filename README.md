# 🎄 Advent of Code 2024 in Go

Solutions for [Advent of Code](https://adventofcode.com/) in Golang

<!--- advent_readme_stars table --->
## 2024 Results

| Day | Part 1 | Part 2 |
| :---: | :---: | :---: |
| [Day 1](https://adventofcode.com/2024/day/1) | ⭐ | ⭐ |
| [Day 2](https://adventofcode.com/2024/day/2) | ⭐ | ⭐ |
| [Day 3](https://adventofcode.com/2024/day/3) | ⭐ | ⭐ |
| [Day 4](https://adventofcode.com/2024/day/4) | ⭐ | ⭐ |
| [Day 5](https://adventofcode.com/2024/day/5) | ⭐ | ⭐ |
| [Day 6](https://adventofcode.com/2024/day/6) | ⭐ | ⭐ |
| [Day 7](https://adventofcode.com/2024/day/7) | ⭐ | ⭐ |
| [Day 8](https://adventofcode.com/2024/day/8) | ⭐ | ⭐ |
| [Day 9](https://adventofcode.com/2024/day/9) | ⭐ | ⭐ |
| [Day 10](https://adventofcode.com/2024/day/10) | ⭐ | ⭐ |
| [Day 11](https://adventofcode.com/2024/day/11) | ⭐ | ⭐ |
| [Day 12](https://adventofcode.com/2024/day/12) | ⭐ | ⭐ |
| [Day 13](https://adventofcode.com/2024/day/13) | ⭐ | ⭐ |
| [Day 14](https://adventofcode.com/2024/day/14) | ⭐ | ⭐ |
<!--- advent_readme_stars table --->

<!--- benchmarking table --->
## Benchmarks

| Day | Part 1 | Part 2 |
| :---: | :---: | :---:  |
| [Day 1](./src/2024/days/01/code.go) | `231.90µs` | `237.18µs` |
| [Day 2](./src/2024/days/02/code.go) | `104.74µs` | `267.70µs` |
| [Day 3](./src/2024/days/03/code.go) | `310.17µs` | `119.57µs` |
| [Day 4](./src/2024/days/04/code.go) | `109.64µs` | `522.25µs` |
| [Day 5](./src/2024/days/05/code.go) | `313.29µs` | `1.68ms` |
| [Day 6](./src/2024/days/06/code.go) | `647.16µs` | `759.54ms` |
| [Day 7](./src/2024/days/07/code.go) | `1.27ms` | `655.29ms` |
| [Day 8](./src/2024/days/08/code.go) | `60.10µs` | `114.06µs` |
| [Day 9](./src/2024/days/09/code.go) | `1.49ms` | `15.98ms` |
| [Day 10](./src/2024/days/10/code.go) | `279.66µs` | `281.03µs` |
| [Day 11](./src/2024/days/11/code.go) | `265.70ns` | `254.80ns` |
| [Day 12](./src/2024/days/12/code.go) | `6.31ms` | `8.34ms` |
| [Day 13](./src/2024/days/13/code.go) | `198.93µs` | `198.14µs` |

**Total: 1453.89ms**
<!--- benchmarking table --->

---

<details>
A handy template repository to hold your [Advent of Code](https://adventofcode.com) solutions in Go (golang).

Advent of Code (<https://adventofcode.com>) is a yearly series of programming questions based on the [Advent Calendar](https://en.wikipedia.org/wiki/Advent_calendar). For each day leading up to christmas, there is one question released, and from the second it is released, there is a timer running and a leaderboard showing who solved it first.

---

### Features

* A directory per question `<year>/<day>`
* Auto-download questions into `<year>/<day>/README.md`
* Auto-download example input into `<year>/<day>/input-example.txt`
* With env variable `AOC_SESSION` set:
  * Auto-download part 2 of questions into `<year>/<day>/README.md`
  * Auto-download user input into `<year>/<day>/input-user.md`
  * Only runs part 2 once part 1 is completed
* When you save `code.go`, it will execute your `run` function 4 times:
  * Input `input-example.txt` and `part2=false`
  * Input `input-example(2).txt` and `part2=true`
  * Input `input-user.txt` and `part2=false`
  * Input `input-user(2).txt` and `part2=true`
  * Each run will display the return value and timing.
  * Part 2 will use the `<file>2.txt` if it exists.
* Control execution with `PART= INPUT= ./run.sh <year> <day>`, where
  * `PART` can be `1` or `2`, and
  * `INPUT` can be `example` or `user`

---

### Usage

1. Click "**Use this template**" above to fork it into your account
1. Setup repo, either locally or in codespaces
   * Locally
      * Install Go from <https://go.dev/dl/> or from brew, etc
      * Git clone your fork
      * Open in VS Code, and install the Go extension
   * Codespaces
      * Click "Open in Codespaces"
1. Open a terminal and `./run.sh <year> <day>` like this:

   ```sh
   $ ./run.sh 2023 1
   Created directory ./2023/01
   Created file code.go
   Created file README.md
   Created file input-example.txt
   run(part1, input-example) returned in 616µs => 42
   ```

1. Implement your solution in `./2023/01/code.go` inside the `run` function
   * I have provided solutions for year `2022`, days `2`,`4`,`7` – however you can delete them and do them yourself if you'd like
1. Changes will re-run the code
   * For example, update `code.go` to `return 43` instead you should see:

   ```sh
   file changed code.go
   run(part1, input-example) returned in 34µs => 43
   ```

1. The question is downloaded to `./2023/01/README.md`
1. Login to <https://adventofcode.com>
1. Find your question (e.g. <https://adventofcode.com/2023/day/1>) and **[get your puzzle input](https://adventofcode.com/2023/day/1/input)** and save it to `./2023/01/input-user.txt`
   * See **Session** below to automate this step
1. Iterate on `code.go` until you get the answer
1. Submit it to <https://adventofcode.com/2023/day/1>

---

#### Session

**Optionally**, you can `export AOC_SESSION=<session>` from your adventofcode.com `session` cookie. That is:

* Login with your browser
* Open developer tools > Application/Storage > Cookies
* Retrieve the contents of `session`
* Export it as `AOC_SESSION`

With your session set, running `code.go` will download your user-specifc `input-user.txt` and also update `README.md` with part 2 of the question once you've completed part 1.

Currently, your session is NOT used to submit your answer. You still need to login to <https://adventofcode.com> to submit.
</details>
