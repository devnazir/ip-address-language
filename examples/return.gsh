func withReturn() {

  if 1 == 1 {
    echo "Yes"
  } else {
    var coba = "No"
    return coba
  }

  return 1
}

var coba = withReturn()

echo $coba