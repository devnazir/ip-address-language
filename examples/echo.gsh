var name = "nazir"
var age = 100

echo my name is $name and my age is $age

age = 200;
name = "ulum"

echo -e my name is\n$name and my age is $age;

// echo with color
echo -e \e[31mHello World in Red Color\e[0m
echo -e \e[32mHello World in Green Color\e[0m
echo -e \e[33mHello World in Yellow Color\e[0m
echo -e \e[34mHello World in Blue Color\e[0m
echo -e \e[35mHello World in Magenta Color\e[0m

// with background white and text color black
echo -e \e[47;30mHello World in White Background and Black Text Color\e[0m

echo $(ls > ./output/ls.txt)

var capturedLs = $(ls)
echo $capturedLs

$(echo "Added" >> ./output/ls.txt)

echo $(echo $(echo 1 2 3) 4 5)

var capturedLs2 = "List of files: \n" + $(ls) + "\nEnd of list"
echo -e $capturedLs2

var str = "Hello $name"
echo $str;

echo Hi $str
