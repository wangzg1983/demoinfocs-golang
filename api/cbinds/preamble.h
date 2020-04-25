#ifndef PREAMBLE_H
#define PREAMBLE_H

#include <stdlib.h>
#include <stdbool.h>


typedef unsigned int ParserHandle;

typedef void* EventPtr;

typedef void EventHandler(ParserHandle p, EventPtr e);

extern void CallEventHandler(ParserHandle h, EventPtr e, EventHandler *cfn);

/*
 * demoinfocs package
 */

typedef struct gameState {
    int TotalRoundsPlayed;
} GameState;

/*
 * events package
 */

typedef enum event {
    KILL_EVENT,
    ROUND_END_EVENT,
    ROUND_START_EVENT
} Event;

typedef struct killEvent {
    struct equipment *Weapon;
    struct player *Victim;
    struct player *Killer;
    int PenetratedObjects;
    bool IsHeadshot;
} KillEvent;

typedef struct roundEnd {
    const char *Message;
} RoundEndEvent;

typedef struct roundStart {
    const char *Objective;
} RoundStartEvent;

/*
 * common package
 */

typedef struct demoHeader {
    const char *Filestamp;
    const char *ServerName;
    const char *ClientName;
    const char *MapName;
    const char *GameDirectory;
    long long PlaybackTime;
    int	Protocol;
    int	NetworkProtocol;
    int	PlaybackTicks;
    int	PlaybackFrames;
    int	SignonLength;
} DemoHeader;

typedef struct equipment {
    const char *Type;
} Equipment;

typedef struct player {
    const char *Name;
} Player;

#endif
