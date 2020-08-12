package main

import (
    "os"
    "fmt"
    "flag"
    "bufio"
    "io"
    "strings"
    "errors"
    "time"
    "github.com/xlzd/gotp"
)

const (
        InfoColor    = "\033[1;34m%s\033[0m"
        NoticeColor  = "\033[1;36m%s\033[0m"
        WarningColor = "\033[1;33m%s\033[0m"
        ErrorColor   = "\033[1;31m%s\033[0m"
        DebugColor   = "\033[0;36m%s\033[0m"
)

type Row struct {
    name string
    code string
}

func main() {
    inputFile := flag.String("f", "", "Input file");
    watchFlag := flag.Bool("w", false, "Refresh data")
    flag.Parse()

    var reader *bufio.Reader
    if *inputFile != "" {
        file, err := os.Open(*inputFile)
        if err != nil {
            panic(err)
        }
        reader = bufio.NewReader(file)
    } else {
        info, err := os.Stdin.Stat()
        if err != nil {
            panic(err)
        }

        if info.Mode() & os.ModeNamedPipe == 0 {
            panic("Nothing to read")
        }

        reader = bufio.NewReader(os.Stdin)
    }

    rows, err := parseInput(reader)
    if err == nil {
        for {
            if *watchFlag {
                clearTerminal()
            }
            for _, row := range rows {
                totp := gotp.NewDefaultTOTP(row.code)
                code := totp.Now()
                fmt.Printf(InfoColor, row.name)
                fmt.Printf("\t")
                fmt.Printf(DebugColor, code)
                fmt.Print(" (", int(totpChangeAfter(code, totp).Seconds()), ")")
                fmt.Println()
            }

            if !*watchFlag {
                break
            }
            time.Sleep(1 * time.Second)
        }
    }
}

func parseInput(reader *bufio.Reader) ([]Row, error) {
    if reader == nil {
        return nil, errors.New("Null reader")
    }

    var output []Row
    for {
        input, err := reader.ReadString('\n')
        if err != nil && err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
        input = strings.TrimSuffix(input, "\n")
        s := strings.Split(input, ":")
        var row Row
        if len(s) > 1 {
            row = Row{name: strings.TrimSpace(s[0]), code: strings.TrimSpace(s[1])}
        } else {
            row = Row{name: "", code: strings.TrimSpace(s[0])}
        }
        output = append(output, row)
    }

    return output, nil
}

func clearTerminal() {
    print("\033c")
}

func totpChangeAfter(oldCode string, totp *gotp.TOTP) time.Duration {
    start := time.Now()
    current := start
    for {
        current = current.Add(1 * time.Second)
        code := totp.At(int(current.Unix()))
        if oldCode != code {
            return current.Sub(start)
        }
    }
}
