### Original

```mermaid
flowchart LR
AA -- 1 --- DD
AA -- 1 --- II
AA -- 1 --- BB
BB -- 1 --- CC
CC -- 1 --- DD
DD -- 1 --- EE
EE -- 1 --- FF
FF -- 1 --- GG
GG -- 1 --- HH
II -- 1 --- JJ

```

### New

```mermaid
flowchart LR
AA -- 1 --- BB
AA -- 2 --- CC
AA -- 1 --- DD
AA -- 2 --- EE
AA -- 5 --- HH
AA -- 2 --- JJ
BB -- 1 --- CC
BB -- 2 --- DD
BB -- 3 --- EE
BB -- 6 --- HH
BB -- 3 --- JJ
CC -- 1 --- DD
CC -- 2 --- EE
CC -- 5 --- HH
CC -- 4 --- JJ
DD -- 1 --- EE
DD -- 4 --- HH
DD -- 3 --- JJ
EE -- 3 --- HH
EE -- 4 --- JJ
HH -- 7 --- JJ

```