# lang (No real name yet...)
A simple compiler written in go.

# Example

```lang


class lang {
  var a : Int = 100;
  var b : Int;

  func gnal (x: Int, y: Int = 6) -> (Int) {
    a = x;
    return y + x;
  }

}


```


# Language Definition

```ebnf

digit       =   "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" ;
uppercase   =   "A" | "B" | "C" | "D" | "E" | "F" | "G" | "H" | "I" | "J" | "K" | "L" | "M" | "N" | "O" | "P" | "Q" | "R" | "S" | "T" | "U"| "V" | "W" | "X" | "Y" | "Z"
lowercase   =   "a" | "b" | "c" | "d" | "e" | "f" | "g" | "h" | "i" | "j" | "k" | "l" | "m" | "n" | "o" | "p" | "q" | "r" | "s" | "t" | "u"| "v" | "w" | "x" | "y" | "z"

bool        =   "true" | "false"
type        =   "Int" | "Bool" | "Char" | "String"

operator    =   "+" | "-" | "*" | "/"
cmpop       =   "<" | ">" | "<=" | ">=" | "==" | "!="
boolop      =   "&&" | "||"

identifier  =   [uppercase | uppercase] {[uppercase | uppercase | digit]}
integer     =   ["-"] digit {digit}

declaration =   "var" identifier ":" type ["=" expression]
assignment  =   identifier equals expresison
expression  =   factor {operator factor}
factor      =   identifier | integer | "(" expression ")"

boolexp     =   boolelem | boolexp boolop boolexp | expression cmpop expression

boolelem    =   bool | "!" boolelem | "(" boolexp ")"

ifstatement =   "if" boolexp "{" statements "}"
                { "else if" boolexp "{" statements "}" }
                [("else" "{" statements "}")]

(* func someFunction (a: Int, b: Bool) -> (Bool) { a = 3; b = 4; return true; } *)
function    =   "func" identifier
                "(" {identifier ":" type ["=" integer] [","]} ")" "->"
                "(" {type [ "," ]} ")" "{" statements "}"

class       =   "class" identifier "{" {declaration} {function} "}"


statements  =   statement ";" {statements}
statement   =   declaration | assignment | ifstatement | function | class

```
