package usage

import "fmt"

// Innovator Usage ...
func Usage() {
	fmt.Println(`Usage: innovator -[hv]

     ------- < Commands Arguments > -------
optional:
  -h, help          Show this help message. 
  -v, version       Show the app version. 

For more help, you can use 'innovator help' for the detailed information
or you can check the docs: https://github.com/ka1i/innovator`)
}
