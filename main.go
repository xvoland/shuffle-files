/*
    Copyright © 2023, Vitalii Tereshchuk | DOTOCA.NET All rights reserved.
    Homepage: https://dotoca.net/shuffle-files

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package main

import (
    "fmt"
    "io/ioutil"
    "math/rand"
    "os"
    "path/filepath"
    "time"
    "encoding/base64"
)

func main() {
    printCopyrights()

    // Check for the correct number of command-line arguments
    if len(os.Args) != 2 {
        printUsage()
        return
    }

    // Get the directory path from the command-line argument
    dir := os.Args[1]

    // List all files in the specified directory
    fileList, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Seed the random number generator
    rand.Seed(time.Now().Unix())

    // Shuffle the file names
    shuffledFileNames := shuffleFileNames(fileList)

    // Rename the files with shuffled names
    for i, file := range fileList {
        oldPath := filepath.Join(dir, file.Name())
        newName := shuffledFileNames[i]
        newPath := filepath.Join(dir, newName)

        if err := os.Rename(oldPath, newPath + ".tmp"); err != nil {
            fmt.Printf("Error renaming %s to %s: %v\n", oldPath, newPath, err)
        }
    }

    for i := range shuffledFileNames {
        newName := shuffledFileNames[i]
        newPath := filepath.Join(dir, newName)

        if err := os.Rename(newPath + ".tmp", newPath); err != nil {
            fmt.Printf("Error renaming %s to %s: %v\n", newPath + ".tmp", newPath, err)
        }
    }

    fmt.Printf("Shuffled %d file(s)\n", len(shuffledFileNames))
}



func shuffleFileNames(fileList []os.FileInfo) []string {
    names := make([]string, len(fileList))
    for i, file := range fileList {
        names[i] = file.Name()
    }

    // Shuffle the file names using the Fisher-Yates algorithm
    for i := len(names) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        names[i], names[j] = names[j], names[i]
    }
    return names
}



func printUsage() {
    fmt.Println("Shuffle the contents of files, keeping their names in the same order.\n")
    fmt.Println("Usage: shuffle-files <directory>")
    fmt.Println("Options:")
    fmt.Println("  <directory string> - Path to the directory containing files to shuffle")
}



func printCopyrights() {
    // Copyright © 2023, Vitalii Tereshchuk | DOTOCA.NET All rights reserved. Homepage: https://dotoca.net/shuffle-files
    str := "Q29weXJpZ2h0IMKpIDIwMjMsIFZpdGFsaWkgVGVyZXNoY2h1ayB8IERPVE9DQS5ORVQgQWxsIHJpZ2h0cyByZXNlcnZlZC4gSG9tZXBhZ2U6IGh0dHBzOi8vZG90b2NhLm5ldC9zaHVmZmxlLWZpbGVz"

    decoded, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(decoded))
}

