# Index

## NicknameGenerator.go

Programme generating random nickname based on
name, surname and day of birth.
All using maps.

Example:
> go run NicknameGenerator.go -day 8 -name Jakub -surname Lord

Result:
```text
Lord Wielki Rozczłonkowany
```

If flag is no set, programme will ask for it.

## QuizGame.go

Micro terminal game that loads questions from csv file (../data/QuizData.csv)

Example playtrough:
```text
Welcome to the quiz app!
You will hear 10 random questions. Each one cant grant you 1 point.
Enter anything to continue!

1: have poland ever had more than 1 000 000 m2 area?
yes
2: 8+6?
14
3: is gopher golang logo right now?
no
4: 5+5?
10
5: is Serbia member of UE?
yes
6: was docker founded before 2015?
no
7: 7+3?
10
8: 8+3?
11
9: does penguin have knees?
yes
10: is messi taller than isignie?
yes

The end!
***************
*Score :  8/10*
***************
```
