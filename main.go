package main
import (
  "fmt"
  "regexp"
  "strconv"
  "io/ioutil"
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
  dat, err := ioutil.ReadFile("test.lg")
  checkError(err)

  program := string(dat)
  var lastToken token
  var line int
  currentPos := 0
  line = 0
  program = program + " "
  for i:=1;i<=len(program);i++ {
    if currentToken := getToken(string(program[currentPos:i])); currentToken.token == IGNR {
      // If the token is important
      if lastToken.token != IGNR {
        // If the token is a new line.
        if lastToken.token == NL {
          line++
        // If the token is not a white space (We ignore white spaces)
        } else if lastToken.token != WS {
          fmt.Print(lastToken.token)
          fmt.Println(": " + lastToken.value)
        }
        i--
      } else {
        generateError("Invalid symbol", line, currentPos)
      }
      currentPos = i
      lastToken = currentToken
    } else {
      //fmt.Println(" != "+currentToken.value)
      lastToken = currentToken
    }

  }
}



func generateError(s string, line int, pos int) {
  panic("Error: " + s + " On Line " + strconv.Itoa(line) + "(" + strconv.Itoa(pos) + ")")
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
