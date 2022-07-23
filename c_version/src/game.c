//
// Created by NintenSAGA on 2020/12/11.
//

#include "include/Global.h"

int game (const int map_original[][12], int x, int y){


    int m, n, temp2;
    int restart_signal = 0;
    int exit_signal = 0;
    int map[y][12];
    int map_backup[y][12];
    int player_backup[2];
    int undo_signal = 0;
    for (n = 0;n < y;n++) for (m = 0;m < x;m++) map[n][m] = map_original[n][m];

    int aim[10][2];
    int i = 0, j;
    int test_success = 0;
    char operation;
    int move_x, move_y;
    int player_x, player_y;

    for (n = 0;n < y;n++) for (m = 0;m < x;m++) if (map[n][m]==1) {
                player_x = m;
                player_y = n;
                break;
            }

    for (n = 0;n < y;n++) for (m = 0;m < x;m++) if (map[n][m]==4) {
                aim[i][0] = n;
                aim[i][1] = m;
                i++;
            }

    while(1) {

        move_x = move_y = 0;
        print_map(map, x, y);
        while ((operation = getchar()) == '\n');

        if (operation != '\n') {


            switch (operation) {
                case 'w':
                case 'W':
                    move_y = -1;
                    break;
                case 's':
                case 'S':
                    move_y = 1;
                    break;
                case 'a':
                case 'A':
                    move_x = -1;
                    break;
                case 'd':
                case 'D':
                    move_x = 1;
                    break;
                case 'r':
                case 'R':
                    restart_signal = 1;
                    break;
                case '.':
                    exit_signal = 1;
                    break;
                case 'i':
                case 'I':
                    if(items>0) undo_signal = 1;
                    break;
            }



            if (exit_signal) {
                return 3;
                break;
            }

            if (restart_signal) {
                return 2;
                break;
            }

            if (undo_signal) {
                items -= 1;
                steps -= 1;
                for (n = 0;n < y;n++) for (m = 0;m < x;m++) map[n][m] = map_backup[n][m];
                player_x = player_backup[0];
                player_y = player_backup[1];
                undo_signal = 0;
                continue;
            }


            if (items > 0){
                for (n = 0;n < y;n++) for (m = 0;m < x;m++) map_backup[n][m] = map[n][m];
                player_backup[0] = player_x;
                player_backup[1] = player_y;
            }

            int temp = map[player_y + move_y][player_x + move_x];

            if (temp == 0) {
                map[player_y + move_y][player_x + move_x] = 1;
                if (map[player_y][player_x] == 5) map[player_y][player_x] = 4;
                else map[player_y][player_x] = 0;
                player_x += move_x;
                player_y += move_y;
                steps+=1;
            }

            if (temp == 4) {
                map[player_y + move_y][player_x + move_x] = 5;
                if (map[player_y][player_x] == 5) map[player_y][player_x] = 4;
                else map[player_y][player_x] = 0;
                player_x += move_x;
                player_y += move_y;
                steps+=1;
            }

            if (temp == 7) {
                map[player_y + move_y][player_x + move_x] = 1;
                if (map[player_y][player_x] == 5) map[player_y][player_x] = 4;
                else map[player_y][player_x] = 0;
                player_x += move_x;
                player_y += move_y;
                steps+=1;
                items += 1;
            }

            if (temp == 2) {
                temp2 = map[player_y + 2 * move_y][player_x + 2 * move_x];
                if (temp2 == 0 || temp2 == 7) {
                    map[player_y + 2 * move_y][player_x + 2 * move_x] = 2;
                    map[player_y + move_y][player_x + move_x] = 1;
                    if (map[player_y][player_x] == 5) map[player_y][player_x] = 4;
                    else map[player_y][player_x] = 0;
                    player_x += move_x;
                    player_y += move_y;
                    steps+=1;
                }
                if (temp2 == 4) {
                    map[player_y + 2 * move_y][player_x + 2 * move_x] = 6;
                    map[player_y + move_y][player_x + move_x] = 1;
                    if (map[player_y][player_x] == 5) map[player_y][player_x] = 4;
                    else map[player_y][player_x] = 0;
                    player_x += move_x;
                    player_y += move_y;
                    steps+=1;
                }
            }

            if (temp == 6) {
                temp2 = map[player_y + 2 * move_y][player_x + 2 * move_x];
                if (temp2 == 0 || temp2 == 7) {
                    map[player_y + 2 * move_y][player_x + 2 * move_x] = 2;
                    map[player_y + move_y][player_x + move_x] = 5;
                    if (map[player_y][player_x] == 5) map[player_y][player_x] = 4;
                    else map[player_y][player_x] = 0;
                    player_x += move_x;
                    player_y += move_y;
                    steps+=1;
                }
                if (temp2 == 4) {
                    map[player_y + 2 * move_y][player_x + 2 * move_x] = 6;
                    map[player_y + move_y][player_x + move_x] = 5;
                    if (map[player_y][player_x] == 5) map[player_y][player_x] = 4;
                    else map[player_y][player_x] = 0;
                    player_x += move_x;
                    player_y += move_y;
                    steps+=1;
                }
            }


            for (j = 0; j < i; j++) {
                test_success = 1;
                if (map[aim[j][0]][aim[j][1]] != 6) {
                    test_success = 0;
                    break;
                }
            }

            if (test_success) {
                return 1;
                break;
            }


        }
    }
}