package init

func ZshInit() string {
	return `
function pulsarship_prompt() {
  "/usr/bin/pulsarship prompt"
}

PROMPT='$(pulsarship_prompt)'
`
}
