package main

type config struct {
	logLevel string
	logPath string

	chanSize int
	kafkaAddr string
	collectConf []tailf.CollectConf
}

func loadConf(filename string) (err error) {

}