# Path finding

## Maps

Maps will include a first line with a "[X]x[Y]" where X is the width and Y is the hight of the map
The next lines will contain the rows ( X number of rows).
Each row will contain Y number of tuples.
Each tuple has X coordinate, Y coordinate, and 1|0 to indicate if it is walkable or not.
Example:

```
4x4
0-0-1 1-0-1 2-0-1 3-0-1
0-1-1 1-1-1 2-1-0 3-1-1
0-2-1 1-2-0 2-2-1 3-2-1
0-3-1 1-3-1 2-3-1 3-3-1
```

The coordinates start with 0,0 at the top left corner
