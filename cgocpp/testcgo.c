#include <stdio.h>
#include "echo.h"

int main(){
    GoString st;
    st.p = "hello";
    st.n = 5;
    printf(":%s\n",Echo(st).p);
    return 0;
}
//gcc -o testecho testcgo.c echo.so
//gcc -o testecho testcgo.c echo.a