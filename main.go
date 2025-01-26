package main

// const nbrPizzas = 10

// var pizzasMade, pizzasFailed, total int

// producer // factory
type Producer struct {
	data chan PizzaOrder
}

// consumer // client // order pizza
type PizzaOrder struct {
	id         int
	message    string
	statusCode int
}

func main() {
	// time.Now().UnixNano()
	// create producer
	// run the producer in the background
	// create and run the consumer
}
