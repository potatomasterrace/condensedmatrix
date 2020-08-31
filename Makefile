test:
	go test

lint:
	pylint --rcfile=pylint.rc matching_algorithms/
	pylint --rcfile=pylint.rc tests/

bench:
	go test -bench=. -benchtime 10000x

coverage:
	go test -coverprofile cover.out
	go tool cover -html=cover.out
	sleep 5 && rm cover.out 

regression_test:
	python dataset_test.py
