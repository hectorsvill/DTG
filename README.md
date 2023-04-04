# Directory-Tree-Generator
This program generates a directory tree for a specified directory and outputs it in either text or JSON format. It can also print the directory tree to the console.

## Usage

```
Usage of dir_tree:
  -d string
        starting directory (default ".")
  -f string
        output format (txt or json) (default "txt")
  -p    print to console
```

By default, the program will generate a directory tree for the current directory and output it in text format to a file named `dir_tree.txt`. You can use the `-d` flag to specify a different starting directory and the `-f` flag to specify a different output format (`txt` or `json`). If you specify `json` as the output format, the program will output the directory tree to a file named `dir_tree.json`.

You can use the `-p` flag to print the directory tree to the console instead of writing it to a file. If no flags are given when running the program, it will output the current directory tree to the console.

## Examples

### Example 1: Generate a directory tree for the current directory in text format

To generate a directory tree for the current directory and output it in text format, you would run the following command:

```
./dir_tree
```

This would create a file named `dir_tree.txt` in the current directory containing the directory tree in text format.

### Example 2: Generate a directory tree for a specified directory in text format

To generate a directory tree for the `/home/user/documents` directory and output it in text format, you would run the following command:

```
./dir_tree -d /home/user/documents
```

This would create a file named `dir_tree.txt` in the current directory containing the directory tree for the `/home/user/documents` directory in text format.

### Example 3: Generate a directory tree for a specified directory in JSON format

To generate a directory tree for the `/home/user/documents` directory and output it in JSON format, you would run the following command:

```
./dir_tree -d /home/user/documents -f json
```

This would create a file named `dir_tree.json` in the current directory containing the directory tree for the `/home/user/documents` directory in JSON format.

### Example 4: Print a directory tree for the current directory to the console

To print a directory tree for the current directory to the console, you would run the following command:

```
./dir_tree -p
```

This would print the current directory tree to the console.

### Example 5: Print a directory tree for a specified directory to the console

To print a directory tree for the `/home/user/documents` directory to the console, you would run the following command:

```
./dir_tree -d /home/user/documents -p
```

This would print the `/home/user/documents` directory tree to the console.