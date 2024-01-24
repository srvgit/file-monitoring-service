package internal

type EventProcessor interface {
	
	updateStatus() error
	processEvent() error
}

type Config struct {
	SrcDirectory    string `json:"sourceDirectory"`
	TargetDirectory string `json:"ReportDirectory"`
	MaxGoRoutines   int    `json:"maxGoRoutines"`
}

func init() {
	// TODO: LOAD CONFIG FILE
}

func (c *Config) loadConfig() error {
	//TODO: Implement
	return nil
}

func (c *Config) processEvent() error {
	//TODO: Implement
	return nil
}

func (c *Config) updateStatus() error {
	//TODO: Implement
	return nil
}
