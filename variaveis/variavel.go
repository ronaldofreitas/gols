package variaveis

var endereco string
var telefone = "9999-9999"
var id int64
var ano int16
var ativo bool
var valor float64
var palavras rune // tipo rune?
var ddd string

// Reverse inverte ordem das letras
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Telefone adiciona o DDD
func Telefone(t string) string {
	ddd = "(071) "
	return string(ddd + t)
}
