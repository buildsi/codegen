all:
	clang++ -fPIC -shared -O3 -g -o libfoo.so foo.c
	clang++ -O3 -o v main.c -L. -lfoo -Wl,-rpath,`pwd`
