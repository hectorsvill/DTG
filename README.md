# Directory-Tree-Generator

This Go program generates a directory tree in either text or JSON format. 

## Command-Line Options

- `-d`: Specify the starting directory for the directory tree. The default value is `.` (the current directory).
- `-f`: Specify the output format for the directory tree. The available options are `txt` (text format) and `json` (JSON format). The default value is `txt`.
- `-p`: Print the directory tree to the console instead of writing it to a file. The default value is `false`.

## Usage

To run the program and generate a directory tree in text format, use the following command:

```
./dir_tree
```

This will create a file named `dir_tree.txt` in the current directory containing the directory tree.

Here's an example of what the directory tree might look like in text format:

```
.
├── file1.txt
├── file2.txt
└── dir1
    ├── file3.txt
    └── dir2
        └── file4.txt
```

To generate a directory tree in JSON format, use the `-f` option:

```
./dir_tree -f json
```

This will create a file named `dir_tree.json` in the current directory containing the directory tree.

Here's an example of what the directory tree might look like in JSON format:

```json
{
  "name": ".",
  "is_dir": true,
  "entries": [
    {
      "name": "file1.txt",
      "is_dir": false
    },
    {
      "name": "file2.txt",
      "is_dir": false
    },
    {
      "name": "dir1",
      "is_dir": true,
      "entries": [
        {
          "name": "file3.txt",
          "is_dir": false
        },
        {
          "name": "dir2",
          "is_dir": true,
          "entries": [
            {
              "name": "file4.txt",
              "is_dir": false
            }
          ]
        }
      ]
    }
  ]
}
```

To print the directory tree to the console instead of writing it to a file, use the `-p` option:

```
./dir_tree -p
```

This will print the directory tree to the console in text format.

To specify a different starting directory for the directory tree, use the `-d` option:

```
./dir_tree -d /path/to/directory
```

This will generate a directory tree starting from the specified directory.

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