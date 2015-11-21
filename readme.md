This is a basic tool to help with [mini](http://www.nytimes.com/crosswords/game/mini) crossword construction. You give it a dictionary, a partially filled grid like this:

    SLATE
    .E...
    .V...
    .E...
    .E...

and it will give you possible completions.

# Installation

    $ go get https://github.com/thanthese/crossword

# Usage

    $ crossword -dict modifiedMacDict.txt -grid grid.txt | head -n 20

         1:  SLATE SHARD HEMAD AVAHI RELIC DEALT LEVEE AMALA TAHIL EDICT
    SLATE
    HEMAD
    AVAHI
    RELIC
    DEALT

         2:  SLATE SHARD HEMAD AVAHI REPIC DEALT LEVEE AMAPA TAHIL EDICT
    SLATE
    HEMAD
    AVAHI
    REPIC
    DEALT

         3:  SLATE SLART LEGAL AVAHI REDIA TEENS LEVEE AGADE TAHIN ELIAS
    SLATE
    LEGAL
    AVAHI
    REDIA
    TEENS

or, more compactly,

    $ crossword -dict modifiedMacDict.txt -grid grid.txt -abbr | head -n 20

         1:  SLATE SHARD HEMAD AVAHI RELIC DEALT LEVEE AMALA TAHIL EDICT
         2:  SLATE SHARD HEMAD AVAHI REPIC DEALT LEVEE AMAPA TAHIL EDICT
         3:  SLATE SLART LEGAL AVAHI REDIA TEENS LEVEE AGADE TAHIN ELIAS
         4:  SLATE SLART LEGAL AVAHI RETIA TEENS LEVEE AGATE TAHIN ELIAS
         5:  SLATE SMART MEDAL AVAHI REGIA TEENS LEVEE ADAGE TAHIN ELIAS
         6:  SLATE SMALT MESAL AVAHI LELIA TEENS LEVEE ASALE TAHIN ELIAS
         7:  SLATE SPART PEDAL AVAHI REGIA TEENS LEVEE ADAGE TAHIN ELIAS
         8:  SLATE SHAPS HEMAL AVAHI PETIT SEINE LEVEE AMATI TAHIN ELITE
         9:  SLATE SNACK NEWAR AVAHI CERIC KEENA LEVEE AWARE TAHIN ERICA
        10:  SLATE STACK TELAR AVAHI CETIC KEENA LEVEE ALATE TAHIN ERICA
        11:  SLATE SMASH MELOE AVAIL SENSE HEDER LEVEE ALAND TOISE EELER
        12:  SLATE STARN TEMSE AVAIL RENNE NEGER LEVEE AMANG TSINE EELER
        13:  SLATE SCART CENSE AVAIL RENNE TEAER LEVEE ANANA TSINE EELER
        14:  SLATE SMART MENSE AVAIL RENNE TEAER LEVEE ANANA TSINE EELER
        15:  SLATE START TENSE AVAIL RENNE TEAER LEVEE ANANA TSINE EELER
        16:  SLATE SMART MESSE AVAIL RENNE TEAER LEVEE ASANA TSINE EELER
        17:  SLATE SMART MERSE AVAIL REUNE TEAER LEVEE ARAUA TSINE EELER
        18:  SLATE SPART PERSE AVAIL REUNE TEAER LEVEE ARAUA TSINE EELER
        19:  SLATE START TERSE AVAIL REUNE TEAER LEVEE ARAUA TSINE EELER
        20:  SLATE SHALT HEMAL AVAIL LEDGE TEIAN LEVEE AMADI TAIGA ELLEN

# Advanced grid

You can put `#`s in the grid to indicate black squares.

    #LATE
    .E...
    .V...
    .E...
    .E..#

Anything after the first blank line in grid is considered a comment. This lets you keep multiple versions of grids you're working on.

# Dictionaries

There is a tension with the dictionary: too many words and the grid becomes too obscure. Too few words and there aren't any solutions. If you find a good dictionary file to use, let me know!

# License

MIT
