//
// Created by NintenSAGA on 2020/12/11.
//
#include "Global.h"

void make_blank(int num){
    int i;
    for (i = 1;i <=num;i++) printf(" ");
}

void print_map(int map[][12], int x, int y){

    int blank = (50-x)/2 + 1;
    int line = (50-y-6+1)/2;
    int i, m, n, temp;

    for(i = 0;i < line;i++) printf("\n");

    printf("┌");
    for(i = 0;i < 48;i++) printf("─");
    printf("┐\n");

    printf("│");
    printf("\tLevel:%d\t\tSteps:%d \tItems:%d",level,steps,items);
    make_blank(2);
    printf("│\n");


    printf("│");
    make_blank(blank-1);
    printf("┌");
    for(i = 0;i < x;i++) printf("─");
    printf("┐");
    make_blank(47-blank-x);
    printf("│\n");

    for (n = 0;n < y;n++){
        printf("│");
        make_blank(blank-1);
        printf("│");
        for (m = 0;m < x;m++){
            temp = map[n][m];
            if (temp == 0) printf(" ");
            if (temp == 1) printf("○");
            if (temp == 2) printf("◆");
            if (temp == 3) printf("■");
            if (temp == 4) printf("◇");
            if (temp == 5) printf("○");
            if (temp == 6) printf("◆");
            if (temp == 7) printf("△");
        }
        printf("│");
        make_blank(47-blank-x);
        printf("│\n");
    }
    printf("│");
    make_blank(blank-1);
    printf("└");
    for(i = 0;i < x;i++) printf("─");
    printf("┘");
    make_blank(47-blank-x);
    printf("│\n");

    printf("│   ◇Move:W,S,A,D  ◇Restart:R  ◇Exit:.  ◇Item:I  │\n");

    printf("└");
    for(i = 0;i < 48;i++) printf("─");
    printf("┘\n");

    for(i = 0;i < ((50-y-6+1) - line);i++) printf("\n");

}