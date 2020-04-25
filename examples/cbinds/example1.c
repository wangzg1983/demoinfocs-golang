#include <stdio.h>
#include <string.h>
#include "libdemoinfocs.h"

void printKills(ParserHandle parser, KillEvent *data) {
    char headshot[6] = "";
    if (data->IsHeadshot) {
        strcpy(headshot, " (HS)");
    }
    char wallbang[6] = "";
    if (data->PenetratedObjects > 0) {
        strcpy(wallbang, " (WB)");
    }
    printf("%s <%s%s%s> %s\n", data->Killer->Name, 
            data->Weapon->Type, headshot, wallbang, 
            data->Victim->Name);
    DestroyKillEvent(data);
}

void printStartMessage(ParserHandle parser, RoundStartEvent *data) {
    GameState *gameState = GetGameState(parser);
    printf("\nRound #%d has started...\n", gameState->TotalRoundsPlayed + 1);
    DestroyGameState(gameState);
}

int main() {
    ParserHandle parser = GetNewParser("demo.dem");
    DemoHeader *demoHeader = ParseHeader(parser);

    printf("\nMap:\t%s\n\n", demoHeader->MapName);

    RegisterEventHandler(parser, KILL_EVENT, (EventHandler *)printKills);
    RegisterEventHandler(parser, ROUND_START_EVENT, (EventHandler *)printStartMessage);

    ParseToEnd(parser);

    DestroyDemoHeader(demoHeader);
    return 0;
}
