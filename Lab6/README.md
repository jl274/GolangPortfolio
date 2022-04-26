# Index

## pongTable.go v1

Concurrency based ping pong game

* Main logic:
    * Ball is passed between pre-declared opponents. After passing it, player/racket waits for
    ball that should come from personal channel.
      ```go
        func (thisRacket *Racket) passTheBall() {
          for {
              fmt.Printf("%s: %s\n", thisRacket.name, thisRacket.motto)
              time.Sleep(1000 * time.Millisecond)
              *thisRacket.opponent.receive <- Ball("")
              <-*thisRacket.receive
          }
        }
      ```
      
* Details:
  
    * After initialization of rackets, opponents must be set for each:
    ```go
    func main() {
        fmt.Println("Test")
        racket1 := NewRacket("Player1", "Ping!")
        racket2 := NewRacket("Player2", "Pong!")
        racket1.setOpponent(racket2)
        racket2.setOpponent(racket1)
        //...
    }
    ```
  
    * Then we have **getReady** and **serve** methods. 
      * getReady starts loop but sets first action as "waiting for the ball"
      * serve also starts loop but first action is passing (serving) the ball (someone has to be first)
    ```go
    func (thisRacket *Racket) getReady() {
        go func() {
            <-*thisRacket.receive
            go thisRacket.passTheBall()
        }()
    }
  
    func (thisRacket *Racket) serve() {
        go thisRacket.passTheBall()
    }
    ```
  * To start the game, all rackets except "first player" needs to be set ready and then "first player"
    may start the game. All go routines are killed as soon as main function ends 
    that's why some time is put to prevent it.
    ```go
    func main() {
        // ...
        racket2.getReady()
        racket1.serve()
        time.Sleep(20000 * time.Millisecond)
    }
    ```
    
### Ping Pong in action:

![GIF](./pongv1.gif)


## pongTable.go v2

Changes:
* The ball was modified to count hits:
  ```go
  type Ball struct {
      timesHit int
  }
  ```
* "Waiting for the ball" loop was modified to increment Ball counter and send precise Ball pointer
  ```go
  func (thisRacket *Racket) passTheBall() {
      var ball Ball
      for {
          ball = <-*thisRacket.receive
          ball.timesHit++
          fmt.Printf("(Hit %d) %s: %s\n", ball.timesHit, thisRacket.name, thisRacket.motto)
          time.Sleep(1000 * time.Millisecond)
          *thisRacket.opponent.receive <- ball
      }
  
  }
  ```
* Because of that change, some refactoring was also needed in getReady and serve:
  ```go
  func (thisRacket *Racket) getReady() {
      go thisRacket.passTheBall()
  }
  
  func (thisRacket *Racket) serve() {
      go thisRacket.passTheBall()
      *thisRacket.receive <- Ball{0}
  }
  ```
  
## wordCounter.go

Program that Concurrently counts words in file and displays that result.

### Usage:
Flags:
* **-file** selects file where words should be counted (text.txt by default)
* **-w** selects number of threads ("gorutines") to handle the task (5 by default)

### Example

>  go run .\wordCounter.go -file text.txt -w 3

```text
the: 15 times
his: 12 times
a: 11 times
was: 8 times
The: 6 times
of: 6 times
and: 6 times
in: 5 times
it: 5 times
coat: 4 times
man: 4 times
to: 4 times
He: 4 times
carried: 3 times
stranger: 3 times
not: 3 times
he: 3 times
but: 3 times
shoulders: 2 times
innkeeper: 2 times
said: 2 times
Old: 2 times
on: 2 times
were: 2 times
horse: 2 times
came: 2 times
at: 2 times
empty: 2 times
stood: 2 times
As: 2 times
street: 2 times
from: 2 times
It: 2 times
front: 2 times
sword: 2 times
almost: 2 times
Narakort: 2 times
one: 2 times
stopped: 1 times
best: 1 times
hour: 1 times
quiver: 1 times
listened: 1 times
wiped: 1 times
drew: 1 times
further: 1 times
afternoon: 1 times
had: 1 times
unpleasant: 1 times
tanners: 1 times
attention: 1 times
hands: 1 times
Not: 1 times
laced: 1 times
pickled: 1 times
hair: 1 times
Beneath: 1 times
saddlers: 1 times
another: 1 times
chipped: 1 times
stalls: 1 times
Gate: 1 times
worn: 1 times
black: 1 times
measured: 1 times
Inn: 1 times
closed: 1 times
Beer: 1 times
old: 1 times
above: 1 times
did: 1 times
Wyzim: 1 times
back: 1 times
entirely: 1 times
or: 1 times
Ropers: 1 times
smaller: 1 times
neck: 1 times
leading: 1 times
hot: 1 times
will: 1 times
off: 1 times
every: 1 times
strapped: 1 times
ropers: 1 times
noticed: 1 times
usual: 1 times
those: 1 times
pulled: 1 times
for: 1 times
foot: 1 times
around: 1 times
if: 1 times
laden: 1 times
His: 1 times
up: 1 times
unusual: 1 times
thrown: 1 times
silent: 1 times
white: 1 times
something: 1 times
Later: 1 times
by: 1 times
hubbub: 1 times
itself: 1 times
still: 1 times
full: 1 times
took: 1 times
voices: 1 times
jerkin: 1 times
this: 1 times
raised: 1 times
no: 1 times
north: 1 times
late: 1 times
people: 1 times
enter: 1 times
earthenware: 1 times
over: 1 times
gaze: 1 times
stiffly: 1 times
What: 1 times
down: 1 times
be: 1 times
wore: 1 times
nearly: 1 times
there: 1 times
already: 1 times
leather: 1 times
that: 1 times
reputations: 1 times
counter: 1 times
him: 1 times
bow: 1 times
weapon: 1 times
himself: 1 times
filled: 1 times
canvas: 1 times
motionless: 1 times
outsider: 1 times
enjoying: 1 times
called: 1 times
cucumbers: 1 times
with: 1 times
apron: 1 times
barrel: 1 times
head: 1 times
bridle: 1 times
tavern: 1 times
tankard: 1 times
Fox: 1 times
as: 1 times
voice: 1 times
moment: 1 times
```