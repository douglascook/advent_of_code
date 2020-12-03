.PHONY: day1 day2 day3

day1: lib.o
	gcc day1/expense_report.c lib.o -g -o run.o
	cp day1/input.txt input.txt

day2: lib.o
	gcc day2/password_policy.c lib.o -g -o run.o
	cp day2/input.txt input.txt

day3: lib.o
	gcc day3/toboggan.c lib.o -g -o run.o
	cp day3/input.txt input.txt

run:
	@./run.o < input.txt

debug:
	lldb run.o

lib.o:
	gcc -c lib/input.c -o lib.o

clean:
	rm -f input.txt
	rm -f *.o
	rm -rf run.o.dSYM
