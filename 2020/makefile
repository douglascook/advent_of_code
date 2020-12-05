.PHONY: day1 day2 day3 day4 day5

# TODO work out how to template these
day1: lib.o
	gcc day1/*.c lib.o -g -o run.o
	cp day1/input.txt input.txt

day2: lib.o
	gcc day2/*.c lib.o -g -o run.o
	cp day2/input.txt input.txt

day3: lib.o
	gcc day3/*.c lib.o -g -o run.o
	cp day3/input.txt input.txt

day4: lib.o
	gcc day4/*.c lib.o -g -o run.o
	cp day4/input.txt input.txt

day5: lib.o
	gcc day5/*.c lib.o -g -o run.o
	cp day5/input.txt input.txt

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
