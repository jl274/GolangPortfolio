# Index

## Terrarium.go

### Topic
This application is a simulation of ants terrarium with a few leaves.
Ants move randomly, picking up leaves and caring them until they meet another leaf.

### String representation
Terrarium is defined by emojis:
* ant = 🐜
* leaf = 🍃
* ant caring leaf = 🐞
* empty field = 🔲

### Examples

#### 15x15 terrarium 7 ants 10 leaves
* 1 step
```text
🐜🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🔲🔲🔲🍃
🔲🍃🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🐜🔲🔲🐜🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲
🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🐜🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🍃🔲🔲
🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🍃🔲🔲🔲🔲🔲🔲🐜🔲🔲🐜🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲

---------------------

🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🍃
🔲🐞🔲🔲🔲🔲🔲🔲🍃🔲🔲🐜🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲
🔲🐞🔲🔲🐜🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🐞🔲🔲🔲🍃🔲🔲
🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🐜🔲🔲🐜🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
```

* 10 step
```text
🍃🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🍃🔲🐜🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🐜🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🐜🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🍃🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🐜🍃🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🍃🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🔲

---------------------

🍃🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🐜🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🐞🔲🔲🔲🍃🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲🔲🔲🔲🐞🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🐞
```

* 100 steps
```text
🐜🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🍃🔲🔲🔲🔲🐜🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🍃🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲🍃
🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🐜🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🍃
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🐜🔲🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🐜🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲

---------------------

🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🍃
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🍃
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐞🐞🔲🐞
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲🐞🐞🐞
```

* 1000 steps
```text
🐜🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🔲🔲
🔲🐜🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🐜🔲🔲🔲🔲
🔲🔲🍃🐜🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🍃🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🍃🔲🍃🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🐜🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲🍃🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🐜

---------------------

🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🐞🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🐞🐞🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🍃🔲🔲🐞🐞🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🐞🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🐞🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🍃🔲🔲🔲🔲🔲🔲🔲🔲
```

### 6x6 terrarium 1 ants 6 leaves - step by step

10 steps:
```text
🍃🔲🔲🍃🔲🔲
🔲🐜🔲🔲🔲🔲
🔲🔲🍃🔲🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲

---------------------

🍃🔲🔲🍃🔲🔲
🔲🐜🔲🔲🔲🔲
🔲🔲🍃🔲🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲

---------------------

🍃🔲🔲🍃🔲🔲
🔲🐜🔲🔲🔲🔲
🔲🔲🍃🔲🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲

---------------------

🐞🔲🔲🍃🔲🔲
🔲🔲🔲🔲🔲🔲
🔲🔲🍃🔲🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲

---------------------

🔲🔲🔲🍃🔲🔲
🔲🐞🔲🔲🔲🔲
🔲🔲🍃🔲🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲

---------------------

🔲🔲🔲🍃🔲🔲
🔲🍃🔲🔲🔲🔲
🔲🔲🐞🔲🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲

---------------------

🔲🔲🔲🍃🔲🔲
🔲🍃🔲🔲🔲🔲
🔲🔲🐞🔲🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲

---------------------

🔲🔲🔲🍃🔲🔲
🔲🐞🔲🔲🔲🔲
🔲🔲🍃🔲🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲

---------------------

🔲🔲🔲🍃🔲🔲
🔲🍃🔲🔲🔲🔲
🔲🔲🐞🔲🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲

---------------------

🔲🔲🔲🍃🔲🔲
🔲🍃🔲🔲🔲🔲
🔲🔲🐞🔲🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲

---------------------

🔲🔲🔲🍃🔲🔲
🔲🍃🔲🔲🔲🔲
🔲🔲🔲🐞🔲🔲
🍃🔲🔲🔲🔲🔲
🔲🔲🔲🔲🍃🔲
🔲🔲🔲🔲🔲🔲
```

#### 20x15 terrarium 8 ants 16 leaves - "animation"

50 steps:
![Animation](50x50x8antsx16leaves.gif)