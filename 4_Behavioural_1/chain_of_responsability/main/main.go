// My Code
package main

import cor "GDP/chain_of_responsability"

func main() {
	//t-chaining set of tasks which will be executed in one shot.
	writerLogger := cor.WriterLogger{}
	second := cor.SecondLogger{NextChain: &writerLogger}
	chain := cor.FirstLogger{NextChain: &second}

	chain.Next("tinkaal")
}
