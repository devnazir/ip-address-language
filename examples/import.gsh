source (
  "variable.gsh"
  "math.gsh" as M
)

var fromExample int = Example

var fromMath string = M.Z

echo $fromExample;

echo $fromMath

echo $M.Z $Example;