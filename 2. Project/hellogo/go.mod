module hellogo

go 1.16


require (
	hellogo.com/constants v0.0.0
	hellogo.com/logics v0.0.0
)

replace (
	hellogo.com/constants v0.0.0 => ./constants
	hellogo.com/logics v0.0.0 => ./logics
)
