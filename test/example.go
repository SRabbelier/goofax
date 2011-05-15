package foo

import "fmt";

// Send the sequence 2, 3, 4, ... to channel 'ch'.

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter() {
	for i := range src {	// Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i;  // Send 'i' to channel 'dst'.
		};
	};
};

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make();  // Create a new channel.
	go generate(ch);       // Start generate() as a subprocess.
	for {
		prime := <-ch;
		Print(prime);
		ch1 := make();
		go filter(ch, ch1, prime);
		ch = ch1;
	};
};

func main() {
	sieve();
};
