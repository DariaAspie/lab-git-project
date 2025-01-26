package main

import (
    "fmt"
    "lab-git-project/src/utils"
)

func main() {
    result := utils.Add(3, 4)
    fmt.Println("The result of adding 3 and 4 is:", result)

    result = utils.Sub(3, 4)
    fmt.Println("The result of subtracting 3 and 4 is:", result)

    result = utils.Mul(3, 4)
    fmt.Println("The result of multiplying 3 and 4 is:", result)
}
