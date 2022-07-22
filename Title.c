//
// Created by NintenSAGA on 2020/12/11.
//

#include "Global.h"

void title(){
    int i;
    for(i = 0;i < 19;i++) printf("\n");

    printf(" *┌");
    for(i = 0;i < 47;i++) printf("─");
    printf("┐*\n");

    printf(
           " *│   _____      _   _               _            │*\n"
           " *│  / ____|    | | | |             (_)           │*\n"
           " *│ | |     __ _| |_| |__   ___ _ __ _ _ __   ___ │*\n"
           " *│ | |    / _` | __| '_ \\ / _ \\ '__| | '_ \\ / _ \\│*\n"
           " *│ | |___| (_| | |_| | | |  __/ |  | | | | |  __/│*\n"
           " *│  \\_____\\__,_|\\__|_| |_|\\___|_|  |_|_| |_|\\___|│*\n"
           " *│                                               │*\n"
           " *│            ©2002-2020 NINTENSAGA              │*\n");

    printf(" *└");
    for(i = 0;i < 47;i++) printf("─");
    printf("┘*\n");

    make_blank(15);
    printf("Press Enter to Start");
    make_blank(15);

    for(i = 0;i < 20;i++) printf("\n");

    getchar();


}


void introduction(){
    int i;
    for(i = 0;i < 19;i++) printf("\n");

    printf(" *┌");
    for(i = 0;i < 47;i++) printf("─");
    printf("┐*\n");

    printf(
            " *│                                               │*\n"
            " *│Your name is Vincent.                          │*\n"
            " *│Because of your getting too close with a girl  │*\n"
            " *│called Catherine, you're cursed by your fiancee│*\n"
            " *│Katherine.                                     │*\n"
            " *│During a nightmare, you're stuck in a dungeon. │*\n"
            " *│To avoid being killed, you have to push all the│*\n"
            " *│blocks to the correct places.                  │*\n"
            " *│                                               │*\n");

    printf(" *└");
    for(i = 0;i < 47;i++) printf("─");
    printf("┘*\n");

    make_blank(15);
    printf("Press Enter to Continue");
    make_blank(15);

    for(i = 0;i < 19;i++) printf("\n");

    getchar();

}


void manual(){
    int i;
    for(i = 0;i < 19;i++) printf("\n");

    printf(" *┌");
    for(i = 0;i < 47;i++) printf("─");
    printf("┐*\n");

    printf(
            " *│                Operation Manual               │*\n"
            " *│                                               │*\n"
            " *│     ○ You               ■ Unmovable Blocks    │*\n"
            " *│     ◆ Object Blocks     ◇ Object Places       │*\n"
            " *│     △ Magical Sand Clocks                     │*\n"
            " *│                                               │*\n"
            " *│  *You can turn back time by entering 'I' when │*\n"
            " *│   holding a Magical Sand Clock*               │*\n"
            " *│                                               │*\n"
);

    printf(" *└");
    for(i = 0;i < 47;i++) printf("─");
    printf("┘*\n");

    make_blank(15);
    printf("Press Enter to Continue");
    make_blank(15);

    for(i = 0;i < 19;i++) printf("\n");

    getchar();

}

void Level_CLear(){
    while(getchar() != '\n');
    int i;
    for(i = 0;i < 19;i++) printf("\n");

    printf(" *┌");
    for(i = 0;i < 47;i++) printf("─");
    printf("┐*\n");

    printf(
            " *│                                               │*\n"
            " *│                                               │*\n"
            " *│                                               │*\n"
            " *│            ╔═╗╦  ╔═╗╔═╗╔╦╗╦═╗╔╦╗              │*\n"
            " *│            ║  ║  ║╣ ╠═╣ ║ ╠╦╝ ║║              │*\n"
            " *│            ╚═╝╩═╝╚═╝╩ ╩ ╩ ╩╚══╩╝              │*\n"
            " *│                                               │*\n"
            " *│                                               │*\n"
            " *│                                               │*\n"
    );

    printf(" *└");
    for(i = 0;i < 47;i++) printf("─");
    printf("┘*\n");

    make_blank(15);
    printf("Press Enter to Continue");
    make_blank(15);

    for(i = 0;i < 19;i++) printf("\n");

    while(getchar() != '\n');

}

void All_CLear(){

    while(getchar() != '\n');

    int i;
    for(i = 0;i < 19;i++) printf("\n");

    printf(" *┌");
    for(i = 0;i < 47;i++) printf("─");
    printf("┐*\n");

    printf(
            " *│               Congratulations!                │*\n"
            " *│                                               │*\n"
            " *│     You've successfully escaped from          │*\n"
            " *│  Katherine's dungeon.                         │*\n"
            " *│                                               │*\n"
            " *│     It's time for you to take a breath!       │*\n"
            " *│                                               │*\n"
            " *│  (But seemingly she won't just let you get    │*\n"
            " *│   away with it so easily......)               │*\n"
    );

    printf(" *└");
    for(i = 0;i < 47;i++) printf("─");
    printf("┘*\n");

    make_blank(15);
    printf("Press Enter to Continue");
    make_blank(15);

    for(i = 0;i < 19;i++) printf("\n");

    while(getchar() != '\n');

}

void Credit(){

    while(getchar() != '\n');

    int i;
    for(i = 0;i < 12;i++) printf("\n");


    printf(
            "       _______ _                 _        \n"
            "      |__   __| |               | |       \n"
            "         | |  | |__   __ _ _ __ | | _____ \n"
            "         | |  | '_ \\ / _` | '_ \\| |/ / __|\n"
            "         | |  | | | | (_| | | | |   <\\__ \\\n"
            "         |_|  |_| |_|\\__,_|_| |_|_|\\_\\___/\n"
            "       ______                             \n"
            "      |  ____|                            \n"
            "      | |__ ___  _ __                     \n"
            "      |  __/ _ \\| '__|                    \n"
            "      | | | (_) | |                       \n"
            "      |_|  \\___/|_|                       \n"
            "      __     __                           \n"
            "      \\ \\   / /                           \n"
            "       \\ \\_/ /__  _   _ _ __              \n"
            "        \\   / _ \\| | | | '__|             \n"
            "         | | (_) | |_| | |                \n"
            "         |_|\\___/ \\__,_|_|                \n"
            "       _____  _             _             \n"
            "      |  __ \\| |           (_)            \n"
            "      | |__) | | __ _ _   _ _ _ __   __ _ \n"
            "      |  ___/| |/ _` | | | | | '_ \\ / _` |\n"
            "      | |    | | (_| | |_| | | | | | (_| |\n"
            "      |_|    |_|\\__,_|\\__, |_|_| |_|\\__, |\n"
            "                       __/ |         __/ |\n"
            "                      |___/         |___/ \n"
    );


    make_blank(11);
    printf("Press Enter to Exit");
    make_blank(11);

    for(i = 0;i < 12;i++) printf("\n");

    while(getchar() != '\n');

}