package piscine
func Swap(a *int, b *int) {
    p := *a
    *a = *b
    *b = p
}