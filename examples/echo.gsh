var name = "nazir"
var age = 100

echo "my name is $name and my age is $age"

age = 200;
name = "ulum"

// echo with color
echo -e "\nmyname is\n$name and my age is $age"

echo -e "\e[31mHello World in Red Color\e[0m"
echo -e "Hello World in Green Color"
echo -e "Hello World in Yellow Color"
echo -e "Hello World in Blue Color";
echo -e "Hello World in Magenta Color"

// with background white and text color black
echo -e "\e[47;30mHello World in White Background and Black Text Color\e[0m"

echo $(ls > ./output/ls.txt)

/**
 * Capture the output of a command
 */

var capturedLs = $(ls)
echo $capturedLs

$(echo "Added" >> ./output/ls.txt)

echo $(echo $(echo 1 2 3) 4 5)

var capturedLs2 = "List of files: \n" + $(ls) + "\nEnd of list"
echo -e $capturedLs2

var str = "Hello $name"
echo $str;

echo Hi $str
