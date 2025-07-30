package init

func FishInit() string {
	return `
set -Ux PULSARSHIP_SHELL "fish"
function fish_prompt
    /usr/bin/pulsarship prompt
end

function fish_right_prompt
	/usr/bin/pulsarship prompt --right
end
`
}
