#include "Global.h"
#include "map_data.h"
int level = 0;
int items = 0;
int steps = 0;


int main(){
    int return_value;
    int item_backup;
    int i = 1;


    title();
    introduction();
    manual();

    item_backup = items;
    while(i){
        level = 1;
        steps = 0;
        items = item_backup;
        return_value = game(map1,9,6);
        if (return_value == 1) break;
        if (return_value == 2) continue;
        if (return_value == 3) i = 0;
    }

    if(i){
        Level_CLear();
    }

    item_backup = items;
    while(i){
        level = 2;
        steps = 0;
        items = item_backup;
        return_value = game(map2, 12, 8);
        if (return_value == 1) break;
        if (return_value == 2) continue;
        if (return_value == 3) i = 0;
    }

    if(i){
        Level_CLear();
    }

    item_backup = items;
    while(i){
        level = 3;
        steps = 0;
        items = item_backup;
        return_value = game(map3, 11, 11);
        if (return_value == 1) break;
        if (return_value == 2) continue;
        if (return_value == 3) i = 0;
    }

    if(i){
        Level_CLear();
        All_CLear();
    }
    Credit();
}



