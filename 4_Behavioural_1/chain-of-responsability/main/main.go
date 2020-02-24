// My Code
package main

import chain_of_responsability "GDP/4_Behavioural_1/chain-of-responsability"

func main() {
	//t-chaining set of tasks which will be executed in one shot.
	writerLogger := chain_of_responsability.WriterLogger{}
	second := chain_of_responsability.SecondLogger{NextChain: &writerLogger}
	chain := chain_of_responsability.FirstLogger{NextChain: &second}

	chain.Next("tinkaal")
}
