#include<stdio.h>

int main(){
    char a = 'x';
    int b= 5;
    float c = 3.1;
    double d = 10.1;

    printf("문자열의 크기 %d\n", sizeof(char));
    printf("정수형의 크기 %d\n", sizeof(int));
    printf("실수의 크기 %f\n", sizeof(float));
    printf("실수의 크기 %f\n", sizeof(double));

}