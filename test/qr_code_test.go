package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)


func TestGenerateCode(t *testing.T) {
	characterByte := strconv.FormatInt(time.Now().Unix(), 36)
	char := string(characterByte)
	
	fmt.Println(time.Now().Unix())
	fmt.Println("Character:", char)	
}