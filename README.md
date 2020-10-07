# suduku

## Usage

``` bash
suduku -i examples/in.txt -o examples/out.txt 
```

## Example

examples/in.txt
```
5 1 . | . . . | 4 . .
. . . | . 5 3 | . 6 .
. . 6 | 4 . . | 2 . .
- - - + - - - + - - -
6 2 . | . 4 . | . . .
. . . | . . 5 | 7 . .
. . . | . . 2 | . 1 .
- - - + - - - + - - -
. . 5 | 3 . . | . . .
. 3 . | 9 6 . | . . 7
. . . | . . . | 8 4 .
```

examples/out.txt
```
5 1 9 | 2 7 6 | 4 3 8
7 4 2 | 8 5 3 | 9 6 1
3 8 6 | 4 1 9 | 2 7 5
- - - + - - - + - - -
6 2 7 | 1 4 8 | 3 5 9
1 9 4 | 6 3 5 | 7 8 2
8 5 3 | 7 9 2 | 6 1 4
- - - + - - - + - - -
2 7 5 | 3 8 4 | 1 9 6
4 3 8 | 9 6 1 | 5 2 7
9 6 1 | 5 2 7 | 8 4 3
```

## License

Licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/suduku/blob/master/LICENSE) for the full license text.
