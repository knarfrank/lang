package main
import (
  "fmt"
  "regexp"
)




func main() {
  program := "43 ^ 34 & 34 | ~345"

  var lastToken token
  currentPos := 0

  program = program + " "
  for i:=1;i<=len(program);i++ {
    if currentToken := getToken(string(program[currentPos:i])); currentToken.token == IGNR {
      if lastToken.token != IGNR {
        fmt.Print(lastToken.token)
        fmt.Println(": " + lastToken.value)
        i--
      }
      currentPos = i
      lastToken = currentToken
    } else {
      //fmt.Println(" != "+currentToken.value)
      lastToken = currentToken
    }

  }

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
