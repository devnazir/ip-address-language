var count int
var count2 int
var str = "Hello World!"

str = "Hello Again" 

var Example = "Hello World!"

var str2 = 'Hello World!'
var template = `Hello World! $str`

var booleanTrue = true
var booleanFalse = false

echo $booleanTrue
echo $booleanFalse

var array = [1, 2, 3, 4, 5]
echo $array[1]

var multiArray = [
  [1, 2, 3], 
  [4, 5, 6], 
  [7, 8, 9]
]

echo $multiArray[1][1]

var object = {
  name: "John",
  age: 30,
  array: [1, 2, 3, 4, 5],
  bool: false,
  nested: {
    name: "John",
    age: 30,
    array: [1, 2, 3, 4, 5],
    bool: false
  }
} 

withoutVar := 2 + 2 * 2
echo $withoutVar