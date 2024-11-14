source (
  "variable.gsh"
  "math.gsh" as m
)

var fromExample = Example
var fromMath = m.Z

echo $fromExample;
echo $fromMath;

echo $m.Z $Example;