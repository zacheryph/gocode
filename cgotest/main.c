#include <dlfcn.h>
#include <stdio.h>

typedef char* (*Hello)(char*);

int main(int argc, char *argv[]) {
	void *hand = dlopen("./dynlib.so", RTLD_LAZY);
  if (!hand) {
    printf("Failed dlopen: %s\n", dlerror());
    return 1;
  }

  Hello helloWorld = (Hello) dlsym(hand, "helloWorld");
  if (!helloWorld) {
    printf("Failed dlsym: %s\n", dlerror());
  }

	printf("Calling helloWorld()\n");
	char *val = helloWorld("Jane");
	printf("Value: %s\n", val);

	return 0;
}
