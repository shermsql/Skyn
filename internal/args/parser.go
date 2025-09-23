package args

import "flag"

type Config struct {
	Target string
	Help   bool
}

func Parse() Config {
	target := flag.String("t", "", "Target Host Or IP To Scan.")
	help := flag.Bool("help", false, "Show Help Message.")

	flag.Parse()

	return Config{
		Target: *target,
		Help:   *help,
	}
}