#!/usr/bin/env bash

fileExt="mp3"
ASSET_PATH=./assets/soundscapes

TYPES=(
"WHITE"
"RAIN"
"STORM"
"BINAURAL"
"OFFICE"
)
VARIANTS=(
"0a"
"1a"
"1b"
"2a"
"2b"
"3a"
"3b"
"4a"
"4b"
"5a"
"5b"
"6a"
"6b"
"7a"
"7b"
"8a"
"8b"
"9a"
"9b"
)

baseUrl="https://mynoise.world/Data"

for audiotype in "${TYPES[@]}"; do
    for variant in "${VARIANTS[@]}"; do
        wget "${baseUrl}/${audioType}/${variant}.${fileExt}" -O "${ASSET_PATH}/${audiotype}/${variant}.${fileExt}"
    done
done
