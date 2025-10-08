# Elden Ring Stat Optimizer

A CLI tool that allows users to enter character build info such as:
- Rune level  
- Weapon
- Base Class 

 which then returns an optimized attribute distribution for achieving the maximum physical damage from the weapon.

### Building and Running
to build and run the program, clone the repository, `cd` into the repository, and run
```
go build main.go
./main
```
### Important Info
This is a very barebones optimizer. There are various other factors that tie into damage output for a specific weapon that are not implemented (but hope to provide support for in the future) such as:
- Weapon scaling factors improving due to weapon level upgrading.
- Weapon scaling factors improving due to affinity application.
- Type of enemy weapon is used against (damage negation).

and since this is optimizing for physical damage, split damage is not taken into consideration (which sounds very interesting to implement). 

### Sources and thanks to
For gathering all Elden Ring data within the program:
- [eldenring-api](https://github.com/deliton/eldenring-api)

This project relies on weapon scaling soft cap data and damage calculation from various external sources (since Elden Ring is not open source), which have been derived from reverse-engineering/experimentation within Elden Ring:

- [Scaling Soft Cap Data from u/getcheddarttv](https://www.reddit.com/r/Eldenring/comments/ta9kmj/put_this_soft_cap_cheat_sheet_together_credit_to/)   
- [Damage Calculation Info](https://eldenring.wiki.fextralife.com/Calculating+Damage)

