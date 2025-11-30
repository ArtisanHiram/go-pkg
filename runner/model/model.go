package __obf_db8bd7ea4a5f8615

type Payload struct {
	Lang    string          `json:"lang"`
	Files   []*InMemoryFile `json:"files"`
	Stdin   string          `json:"stdin"`
	Command string          `json:"command"`
}

type InMemoryFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}
