package _779_k_th_symbol_in_grammar

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	return (^K & 1) ^ kthGrammar(N-1, (K+1)/2)
}
