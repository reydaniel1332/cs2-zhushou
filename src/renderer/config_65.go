package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.4" )

func jVQyvLQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIVl3L2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q3HebFxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cz2WoeccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1PcnexdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vKWiPZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2MT1SC77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EaZRblvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZI51DJVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLkBOWoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A45zBgYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqjROwI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRsTAJe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzMKFDB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RSVKa1Y5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTvdMp2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPrKFu3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THHVX1p8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HmAlJhfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFJpVvRtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0ZTeI6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CfpreCx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhOXOVF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bs3MCux6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBP8yq5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jj0NBnrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8kurGeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDI0jrqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4e89uXvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDGXtoWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8dW0cZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ez6aJdXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QmsUNb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8aN6cpGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xE2VBpNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBnIuSAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func geZ16PzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFisP3P2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O839ANY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kR5pvMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GA5GwN2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNKBeD9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvfJ3TbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5dVPdCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q6YTIklSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSLk2swaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EHf5a04PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OO6hqmuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UeQdn7jzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNhG4xokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFDWrkKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AuPNWqrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umuxngRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXjIy5PQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7ITzxGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54N5W2H0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NjyiE0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBRbYzE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpHu7NXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmY8SE8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yw3aGR9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33gY73JBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEU2fqN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSpXT7hKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHFGePq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kpKBgxmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x417z64sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiHAY4LwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSxlMtk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLO4jViDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owEyi4RuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fPmGBoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymkxycFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDhJWbheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUDjVemRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2xt3bZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 97GRzg0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZsbwPwE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHmwyJa9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aya3A0lBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyv39wEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMign4OEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QaWABMbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tWLEK5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4eOykmbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKI7S3ZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4SCnCctOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func feD2ct59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCT3i0C9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXsqnbA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2GhfhD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etKWigKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQVJfOsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVBWhibIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjPo1FCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9ubQX4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EmGW68IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPuLD4tLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHPZHFsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acSkNuDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sPh6qKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6iwqhwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3PksdGECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CvoczZR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0WEwjoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AWXpwItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbvZcAjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTJsRciuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AEIB9F7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHzFm0dvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jB4I8bGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKncPJ22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdILm2JDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtTDewfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSkBDe7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTiHo97VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Qp5AigaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwMM3tPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9iPrqWcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ChoPKDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNxLZMVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHWF7YZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijNKHXnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHfcaNu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLZybMdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTynQM4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2h6C8MKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mo95gfILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B4D1NjYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pOVu0HRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YphME4woWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ce1KOAbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lr9GSVayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trOKqXBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MwqQVDpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTe97idXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ZnUq8cmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6joQHPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUlqpix2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZ2EX335Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VnBqhlfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1OpmVyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wPo8Ub7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pbe6fC77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2FttxJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYGK79JBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoYET2wrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0K9P5dJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFMiWkEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1eScO7gkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6DEcsMWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kbeUnNytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kto4VesuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVga8Ta0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mnhE3h9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtrLNnLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tloAGg26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWzPTH2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJu3WucOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMUotuRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vN54RYmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rabboTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErYQarjRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSX4e9YhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJraX63mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBVCWz6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKNTdg4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EQQ1b9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDM6ZwTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func quYAKkQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Im3MgzYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBIDM7FpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgGHwAV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5mfDHoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acu8THlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iu0DN4MjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0rl6aN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2VAhoyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S25USzoCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWFVWxCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xl915KkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wH4LatPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdziyzEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKwmFgUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwB4vL43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDz5ApydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRcJhm85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWo3c3gZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPR9hpSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQCEk7hdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4bzKkM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpSjY4qEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WhIHvwhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYsQnnmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqPFSGeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xynkZziiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3quiTxLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0vN8pYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6qx1mAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FzpJooWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRii489OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STNn4LMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5LFnYpbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWB3MVL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybkwMCXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnTuRWE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7JcEgdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHwxZwXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmDJsyLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5z1yjXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fo5rfaknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeVnjRU0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QnuXbSGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfkB6YRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLAhPHyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWVZi9vXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfqSoH1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jKK5OCChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G82uWZBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdjOMMZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C94InmSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APG9OcpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vT5ucscsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyf5RVATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykHkFWvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIXyVKwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKthtCDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyqCnbTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJ7QjdozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4PM9Kj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NR657aYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waU43lEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQVo5CG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMgX6Mn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJcpChRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0ZhIoUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGYNDS4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqRcnJojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nIMFKesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJunE2pGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5TXV3rybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxv2HH7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMxGxluSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KkAHAx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLIFj4dQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TqRgoodUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlo6xA2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0sVOrLc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TF6qU9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gudnoee4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyICnDRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WemjRKzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4352w8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x76c5BQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sQJF6MqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWzhAIS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcIA5fXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hmK8IgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8DN6DcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RqcUxDpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1pKLt2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func le7N9BkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FdWGjJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmoZZpKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mtbUx6ebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9py1NiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcfxBUIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVlxJsnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mH5TkQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHLK5PSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3iJaiFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8RoQBaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNRaG1pAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDUSqBKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3qj0LfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1l2klxhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wavV43msWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYYBw0yrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OR89jffUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1zcURgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kiTHl2AQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5rez6DfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tlpw97sZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcU5JY0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeT6F5KHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKDVgKanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func neIzhPVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jYaMiTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPvZDez0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMAZbSsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6CDnrsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lczmoJqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NCM4evcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KYyw3rcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ehoi9HwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouu32WycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gq16PDqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmYSAfPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFkqd4oRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xeZpyUuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMJTFVZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZe6GUapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIkbiCxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cu0yK4s9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ruLeF0twWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zd9hQe7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICRVJKTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztQVcDagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQ4vaN8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAP4TEaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwUQgbtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func las6RR2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJZxGFbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VeRKDWjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCmIKSb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mhr8vWJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wr4ZMxYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fr1uKlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JeLN15ajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETTS31HxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ooWtIjl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87nW7uqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dPnzDEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTxspKztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Myp1YkWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAkzCHosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7sPWRlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evCZQfsCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZ5l0djxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZyzjr6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvRu3FOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERICRDKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mBpUvv0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TalNGVUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxjWjejuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMAuOgmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBFYzmNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Az8KWFCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JlCLmAmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8IH6kFw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qN6ObqVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9PJQVD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJFJYO0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57DQWkoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjpM26xvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTtYd4BbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvq2Ye7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMS3lU2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eTOFTxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpDZP3qpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aCPEFvgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9zs4Pb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Emm458JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHhpXEjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8Smy1lMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JwpC6qyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxkGlNjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eEzaHJf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CgVgAkqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RrRqUVA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ppt9AsWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfpYszJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iet6AqwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJOnkHEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTlPbJ23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sQYWFXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PlHs1cvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcvI7h2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4l3In6qsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJznmun4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mx6X47OZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdFgFPffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPQbBpXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkBX5o3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyWTNBT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLth8wTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vITdKpPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DnkttbwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qErMeW9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKMY0BdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2j0sSalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y1ZxQEGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKYdHbhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEuZZPXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivxgF8uVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPmOSLXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoZz8jTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ePoAqCvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtGW0rvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func al8LhAiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9aatkSiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZjFesHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QF41YPPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4PIVXigSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stqYKRC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXtvvx4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgJc0ekqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jvl6ikE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcXSDDS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANTaE2HMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvrQBft6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWX198ziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Jz9xo4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nS7T2VdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SsHh0wxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNA4XHIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BkOwZNrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmBsozfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Dw4ZNYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2eXTqXJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdeg6bU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WL4ULZqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FklgLfxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmC4y81aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiV2UTYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nt0uxCyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HC66UFCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqAEkj0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fRqNr3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJbfcQgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prJURMk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wALFO18rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwqUP7qQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5xzc6qzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKlkf1ObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUQIupLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdLqh3KPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A63IX5x0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hr2ASrIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 37QwITeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZ6X0e7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dNUVZy53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdJ0eRmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5WBHg7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQ9GNTrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRMFvzK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnRaWh4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mPNxbgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTToZoxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZ35sCUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ve3uYNePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAJgg05bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWiOZjiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmXAoCuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9MWLfI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvn2oWIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcZXDcegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9jddjTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9N7REkstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reORNoIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmFuHXh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQ2l7xnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0v3D9kVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frHx1JQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S8ZC6uzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func paFdhnhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaLFNemvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2P08pLWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3lLSiRk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7MP8REfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdEbhwJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sQLOl0YMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GARjihekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SvB3fdHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func brh5bE2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzqpMyGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QkhxaTj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QnYYySZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWCs4713Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrQRjmuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gd2O6aFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxWRJvGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQunzJXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJvJBM3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgDkmXEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEoxfgV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbw0LY7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCP345U8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6jmYNygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ys2ENPVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qc7UXPLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVDNccDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUEJfNjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5h8lIriOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wR2RNKZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrMbWhH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ji3FQuYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnU1ij3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xiN6qTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RFyCnLVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ro3QUxcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDwUvlt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfdpH0moWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AuRJ8i2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wp85EJjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GERBYf52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7hJ4WaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTWlqKyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XR49FL7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjFjxwbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSEBPZeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHAP7oPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnfCThCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MjY8ySEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOqYbWpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrZzOY7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OQQ5XzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgynZ7rjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HyXcKai4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 254kHHbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36nZp6KDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bem2mLVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwAfM7jdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIUNDLKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XERDnFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4NiATuXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CC7D7ZVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQ9A0whGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zQ1lVBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmFDaHaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cou6v4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZe9c7KiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dC0bzxKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPLtMCEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRAPO11oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQYljiZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gt596yDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUps8pbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oII0fx3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifvPIYhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nrVUKWNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxor4CPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8xpXeZifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJizZqlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxqeJCMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6L1GpxMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayBehptzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQXVaizNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZcHVmORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12zBkKMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5ioDAjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FF1EnABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C0ro7Ji4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OpQDIcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func emeIqVORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VL14RopcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVD510b4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qoUsu8mvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nl3fXLh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4bYBbln7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztaMlJxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfuVtyYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6juJn2SwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKaUfSCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8Gegah8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func emw58VAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqQJ7qtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func excTsu2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhK2PhJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErNUztItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6UI565twWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ny40ehboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAFguYTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5kPGsS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C85LF0erWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Me1J6ikLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yg3Etu82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgSKglnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7JL2AGKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7npUkmj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVcrn3bsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jon8GSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwypEO18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmAgu90nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAHEeNvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MP3j0jK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBAz7kVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMvjhPVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nft0iYE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uksInBLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxQftSf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUrQOj83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SA0dJA09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTUO5BLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAC4T9QXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDoe2ZMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPPzeqQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfoo8OH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovgoco6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2l2UGcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONRPuyckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4cScOA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diAgeEyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQWfcVYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z1Q88ESUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEFYi0KqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbzaOmaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y592HwbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkpbQJRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfOruKZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKRUPeKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wnEicR06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nk7DXW55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYi5cCYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jX52iAwQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LybWKRftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05yw9BwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A9e9V0OfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBsBL681Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZL0eR9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlPPtdrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DF3lXw4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0JSCz5XAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIBdNA5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qyGAD88XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMk7SMamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7tPYv4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVFwf9NnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3AxcwVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcWiRutbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0uZKjeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0igJc4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXq1vEFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKJDGJiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3rMyeurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksTzEKYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tlESihZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JrwQP9CpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpkTnkd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57AcY4rTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZdC7f3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iy9nKPbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AarRf37sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gFCfsw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EgOLpQW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfPSpmIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSBDfyaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycYXjZDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezD4FWqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FKXSxY1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 095mXaIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c6vBCXxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tcPBUEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cHSkPTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kKrRlHJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 11k1g5QVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oq2UXpAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btxh4QbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4ZdYzXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2in6IS7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1fQC7M1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Db9t2HpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7sNEZNZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qT1j0SFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func As2MFch4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUQB4qw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLbWEk2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqtoMAdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9B7OC8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZValPJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4MPNTZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QEilkJb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8aBZgwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iU3jShT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6Rrsnz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OixLdnW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ioLkgpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhMo6GQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwalteUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwOli3ZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2T1p8SzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzX3IrVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2zAoJanhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q13PBebLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQYeSjsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2xQ8kqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RlX91UMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvrncWhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClXT8hjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezdQuCS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYfzvamgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofS8gVQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1B0Wi7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AcKsBzduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func phBjSh76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1SfVbq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPbTxLr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DsBqVyvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHLKgVqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func slCz8Q0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqoyR4NrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2C6ICd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SK2dKv6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9ojxOnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mBIiLn4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YGGX1GFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ttWAQnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IFE2QJBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bn4qAV0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMiUeYgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oqxz98XqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABOxD8BdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhUAvaA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1OlYyzCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXvYviS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMe3GpqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dttCzTQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGMP71m3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xeOjIH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Rc1xraXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pU9JCkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xulplfTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzdNVVLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jpquo34hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGjZXpr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9rMIxMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jK14gadQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqewXB8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 10pG12YRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUBayIvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFqjJaHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHMMLl1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xT7NvvvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NHR7k6OiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ux0ZnvDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvp6UtjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWi9v7fFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSCrHbZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPF3EaDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q6ChrspoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDOvQ8QAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWnmrETdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLCfuhsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JD8u0osGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Op8MjgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8c9ANq6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RAuB6hc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgncc6OWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEI5AGKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALcfqr56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQkx2kzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boOje9N8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PD5cRQZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfoFYeEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHD7mB6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKKWyTn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uu06EGoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

