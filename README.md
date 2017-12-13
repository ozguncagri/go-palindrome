# go-palindrome

After watching [this](https://www.youtube.com/watch?v=bN8PE3eljdA) video I've decided to make my own reverse addition palindromic number calculator.

# What is palindromic number?

From Wikipedia;

A **palindromic number** or **numeral palindrome** is a number that remains the same when its digits are reversed. Like 16461, for example, it is "symmetrical". The term palindromic is derived from palindrome, which refers to a word (such as rotor or racecar) whose spelling is unchanged when its letters are reversed.

# How calculation works?

So if we enter number 395 as an input, calculation works like;

```plain
395 + 593 = 988
988 + 889 = 1877
1877 + 7781 = 9658
9658 + 8569 = 18227
18227 + 72281 = 90508
90508 + 80509 = 171017
171017 + 710171 = 881188
```

until result of addition is palindromic number.

# What is max supported number?

Because of working with big integers there is no explicit limit for input or calculations.

# Known biggest palindrome

If you enter `1186060307891929990` as an input. It will going to find palindromic number in 261 calculation cycles.

# What about 196?

It will probably run forever.

# Usage

Usage is very simple and there is only 3 flags exists.

-l (uint) : Cycle limit for calculating palindromic number (default 0 means no limit)

-n (string) : Sets test number for calculating palindromic number (accepst only base 10 numbers)

-nolog : Truncates reversal add operation logs

# Example

For calculating palindrome of 395

```sh
go-palindrome -n 395
```

For calculating palindrome of 196 but limiting calculation cycle on 1000

```sh
go-palindrome -n 196 -l 1000
```

For calculating palindrome of 196 but limiting calculation cycle on 1000 and not print calculation operation logs

```sh
go-palindrome -n 196 -l 1000 -nolog
```