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