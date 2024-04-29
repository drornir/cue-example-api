source ".bingo/variables.env"

function gen_api() {
	(cd api
		${CUE} get go --local
	)
}

gen_api
