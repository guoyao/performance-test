#include <stdio.h>

int fibonacci(n) {
  return n < 2 ? n : fibonacci(n - 2) + fibonacci(n - 1);
}

int main() {
  printf("%d", fibonacci(45));
  //fibonacci(40);
  return 0;
}
