#include <stdio.h>
#include <stdlib.h>

int add(int a ,int b){
    return a+b;
}

int sub(int a,int b){
    return a-b;
}

char * echo(char * str){
    return str;
}

struct AddArgs {
    int a;
    int b;
};

int newadd(struct AddArgs arg){

    return arg.a + arg.b;
}