#include <stdio.h>

#define EXIT_SUCCESS 0
#define EXIT_ERROR 1

int write(char ch);
int main(int argc, char *argv[]) {

  char ch;
  FILE *fptr;

  fptr = fopen("old.txt", "r");

  if (fptr == NULL) {
    printf("Error: File cant be opend");
    return EXIT_ERROR;
  }

  while (ch != EOF) {
    ch = fgetc(fptr);
    if (ch != '*' && ch != '#' && ch != '-')
      write(ch);
  }

  fclose(fptr);

  return EXIT_SUCCESS;
};

int write(char ch) {
  FILE *fptr;

  fptr = fopen("new.txt", "a");

  if (fptr == NULL) {
    printf("Error: File cant be opend");
    return EXIT_ERROR;
  }

  fprintf(fptr, "%c", ch);
  fclose(fptr);

  return EXIT_SUCCESS;
};
