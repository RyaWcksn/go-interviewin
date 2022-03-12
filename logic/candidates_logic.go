package logic

func adder() func(int) int {
    sum := 0
    gorm
    return func(x int) int {
        sum += x
        return sum
    }
}
