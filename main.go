package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func printUsageAndExit() {
  fmt.Println("Usage: file-transfer-over-powershell FILE")
  os.Exit(1)
}

func must[T any](val T, err error) T {
  if err != nil {
    panic(err)
  }
  return val
}

func main() {
  flag.Parse()

  if flag.NArg() != 1 {
    printUsageAndExit()
  }

  filePath := flag.Arg(0)
  fileName := filepath.Base(filePath)
  fileBytes := must(os.ReadFile(filePath))
  fileBase64 := base64.StdEncoding.EncodeToString(fileBytes)

  fmt.Printf(`$b64 = '%s'
$filename = "$env:TEMP\%s"
$bytes = [Convert]::FromBase64String($b64)
[IO.File]::WriteAllBytes($filename, $bytes)
explorer.exe "$env:TEMP"
`, fileBase64, fileName)
}
