# SPECS

Specifications for the GOSH language.

### Keywords

```
source | if | else | func | return | var | for
while | break | continue | true | false | nil | args
```

### Shell Keywords

```
cd | mkdir | ls | rm | cp | mv | touch | cat | echo | pwd | clear | exit
git | npm | yarn
```

### Wrapper Shell Keywords

Used when there is no shell keyword that currently hasn't been implemented.

```
$[command]
```

#### Example

```
$echo "Hello, World!"
```

### Imports

#### Single import

```
source "path/to/file.gsh"
```

#### Multiple imports

```
source (
  "path/to/file1.gsh"
  "path/to/file2.gsh"
)
```

```
source "path/to/file1.gsh"
source "path/to/file2.gsh"
```

### Running Scripts

```
gsh path/to/file.gsh
```

#### With Flags

```
gsh path/to/file.gsh --flag1 [value1] --flag2 [value2]
```