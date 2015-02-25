# lang (No real name yet... will there ever be?)
A simple compiler written in go.


# Language Definition

digit       =   "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" ;

uppercase   =   "A" | "B" | "C" | "D" | "E" | "F" | "G" | "H" | "I" | "J" | "K" | "L" | "M" | "N" | "O" | "P" | "Q" | "R" | "S" | "T" | "U"| "V" | "W" | "X" | "Y" | "Z"

lowercase   =   "a" | "b" | "c" | "d" | "e" | "f" | "g" | "h" | "i" | "j" | "k" | "l" | "m" | "n" | "o" | "p" | "q" | "r" | "s" | "t" | "u"| "v" | "w" | "x" | "y" | "z"

bool        =   "true" | "false"
operator    =   "+" | "-" | "*" | "/"

boolop      =   "<" | ">" | "<=" | ">=" | "==" | "!="

identifier  =   [uppercase | uppercase] {[uppercase | uppercase | digit]}

integer     =   ["-"] digit {digit}


assignment  =   identifier equals expresison

expression  =   factor {operator factor}

factor      =   identifier | integer | "(" expression ")"

boolexp     =    bool | (expression boolop expression)

ifstatement =   "if" boolexp "{" statements "}"
