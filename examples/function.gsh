var name = "John Doe"
echo "Hello, $name!"

func sayHello() {
  var name = "Ulum"
  echo "Hello, $name!"

  func sayHello2() {
    var name = "Doe"
    var age = 20
    echo "Hello, $name!"
    
    func sayHello3() {
      var name = "John"
      echo "Hello, $name!"
      echo "Age: $age"
    }

    sayHello3()
  }

  sayHello2()
}

sayHello()

func withRest(name, age, ...rest) {
  echo "Hello, $name!"
  echo $rest[0]
}

withRest("Nazir", "Ulum", "Doe")

var fn = func() {
  echo "Hello, From Anonymous Function"
}

fn()

var fn2 = func(name, ...rest) {
  echo "Hello, $name From Anonymous Function"
  echo "Rest:" $rest[1]
}

fn2("Ulum", "Doe", "Nazir")