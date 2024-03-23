package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        // Obtenez le répertoire de travail actuel
        currentDir, err := os.Getwd()
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
        fmt.Printf("%s $ ", currentDir)
        cmdString, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
        err = runCommand(cmdString)
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}

func runCommand(commandStr string) error {
    commandStr = strings.TrimSuffix(commandStr, "\n")
    arrCommandStr := strings.Fields(commandStr)
    switch arrCommandStr[0] {
    case "exit":
        os.Exit(0)
    case "cd":
        if len(arrCommandStr) != 2 {
            return fmt.Errorf("Usage: cd <directory>")
        }
        err := os.Chdir(arrCommandStr[1])
        if err != nil {
            return err
        }
    default:
        cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
        cmd.Stderr = os.Stderr
        cmd.Stdout = os.Stdout
        return cmd.Run()
    }
    return nil
} 
