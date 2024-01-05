package hello

import "rsc.io/quote/v3"

func Hello() string {
	return quote.Concurrency()
}

func Proverb() string {
	return quote.HelloV3()
}
