package main
import (
  "fmt"
  "regexp"
  "strconv"
  "io/ioutil"
  "container/list"
)





func main() {
  defer func() {
    if r := recover(); r != nil {
        fmt.Println("Compilation Failed.", r)
    }
  }()
  lex()
}

func lex() {
  // Read in file (should move to function...)
  dat, err := ioutil.ReadFile("test.lg")
  checkError(err)
  program := string(dat)

  tokenList := list.New()
  var lastToken token

  var lineNumber int
  var currentLine string
  var linePos int
  var columnNumber int

  var comment bool
  currentPos := 0
  lineNumber = 0
  columnNumber = 0
  comment = false
  program = program + " "
  for i:=1;i<=len(program);i++ {
    //fmt.Println( strconv.Itoa(columnNumber) + ": " + string(program[currentPos:i]))
    currentLine = string(program[linePos:i])
    if currentToken := getToken(string(program[currentPos:i])); currentToken.token == IGNR {

      // If the last token was a start of a block comment
      if lastToken.token == SCMNT {
        comment = true
      }

      // If the token is not recognised generate an error
      if lastToken.token == IGNR {
        generateError("Invalid symbol", lineNumber, columnNumber, currentLine)
      } else {
        // If the token is a new line.
        if lastToken.token == NL {
          // Reset line variables
          linePos = currentPos
          currentLine = ""
          columnNumber = 0
          // Increment Line number
          lineNumber++

        // If the token is not a white space (We ignore white spaces)
        } else if lastToken.token != WS {
          // if we are not in a comment block, add to token list
          if comment == false {
            // Push to the back of the list (Only to the back to make it easy to read)
            tokenList.PushBack(lastToken)
          } else {
            // If end of block comment, get out of comment mode.
            if lastToken.token == ECMNT {
              comment = false
            }
          }
        }
        // Decrement i, so we start checking again at the correct position
        i--
        columnNumber--
      }

      currentPos = i
      lastToken = currentToken
    } else {
      lastToken = currentToken
    }
    columnNumber++
  }
  for e := tokenList.Front(); e != nil; e = e.Next() {
		fmt.Println(getTokenName((e.Value).(token).token))
	}
}



func generateError(s string, lineNumber int, pos int, line string) {
  fmt.Println(line)
  panic("Error: " + s + " On Line " + strconv.Itoa(lineNumber) + " column " + strconv.Itoa(pos))
}

func getToken(s string) token {
  tokens := getTokens()
  for _, t := range tokens {
    // We return the first one that matches.
    if(checkToken(t, s)) {
      t.value = s
      return t
    }
  }
  return token{IGNR, "", ""}
}

func checkToken(t token, substr string) bool {

  match, err := regexp.MatchString(t.regex, substr)
  checkError(err);
  if match {
    return true
  }
  return false
}

func checkError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
