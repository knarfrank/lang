package lexer
import (
  "fmt"
  "regexp"
  "strconv"
  "lang/tokens"
)
func Lex(program string) []tokens.Token {

  tokenList := make([]tokens.Token, 0)
  var lastToken tokens.Token

  var lineNumber int
  var currentLine string
  var linePos int
  var columnNumber int

  var comment bool
  currentPos := 0
  lineNumber = 1
  columnNumber = 0
  comment = false
  program = program + " "

  for i:=1;i<=len(program);i++ {
    currentLine = string(program[linePos:i])
    if currentToken := getToken(string(program[currentPos:i]), lineNumber, columnNumber); currentToken.Token == tokens.IGNR {
      // If the last token was a start of a block comment
      if lastToken.Token == tokens.SCMNT {
        comment = true
      }

      // If the token is not recognised generate an error
      if lastToken.Token == tokens.IGNR {
        generateError("Invalid symbol", lineNumber, columnNumber, currentLine)
      } else {
        // If the token is a new line.
        if lastToken.Token == tokens.NL {
          // Reset line variables
          linePos = currentPos
          currentLine = ""
          columnNumber = 0
          // Increment Line number
          lineNumber++

        // If the token is not a white space (We ignore white spaces)
        } else if lastToken.Token != tokens.WS {
          // if we are not in a comment block, add to token list
          if comment == false {
            // Push to the back of the list (Only to the back to make it easy to read)
            tokenList = append(tokenList, lastToken)
          } else {
            // If end of block comment, get out of comment mode.
            if lastToken.Token == tokens.ECMNT {
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
      // If character is a new line.
      if program[i-1] == 10 {
        lineNumber++
        columnNumber = 0
        currentLine = ""
      }
    }
    columnNumber++
  }
  return tokenList
}


func getToken(s string, line int, column int) tokens.Token {
  tok := tokens.GetTokens()

  for _, t := range tok {
    // We return the first one that matches.
    if(checkToken(t, s)) {
      t.Value = s
      t.Line = line
      t.Col = column
      return t
    }
  }
  return tokens.Token{tokens.IGNR, "", "", line, column}
}

func checkToken(t tokens.Token, substr string) bool {

  match, err := regexp.MatchString(t.Regex, substr)
  checkError(err);
  if match {
    return true
  }
  return false
}

func generateError(s string, lineNumber int, pos int, line string) {
  fmt.Println(line)
  panic("Lex Error: " + s + " On Line " + strconv.Itoa(lineNumber) + " column " + strconv.Itoa(pos))
}


func checkError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
