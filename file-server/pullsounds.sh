#!/usr/bin/env bash

file_ext="mp3"
ASSET_PATH=./assets/tracks
banner_remote_name="fb"
banner_local_name="banner"
banner_ext="jpg"


TYPES=(
"WHITE"
"RAIN"
"RAIN3"
"STORM"
"BINAURAL"
"OFFICE"
"GREY"
"VIBRAPHONE2"
"PIANO"
"JAPANGARDEN"
"FOREST"
"IRISHCOAST"
"NOTWHITE"
"WINDCHIMES"
"STALACTITES"
"THROAT"
"BOWLS"
"AUTUMNWALK"
"FINWHALE"
"DEEPCHANT"
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
imgBaseUrl="https://mynoise.net/Data"

for audio_type in "${TYPES[@]}"; do
    mkdir -p "${ASSET_PATH}/${audio_type}"
    for variant in "${VARIANTS[@]}"; do
        wget "${baseUrl}/${audio_type}/${variant}.${file_ext}" -O "${ASSET_PATH}/${audio_type}/${variant}.${file_ext}"
    done
    wget "${imgBaseUrl}/${audio_type}/${banner_remote_name}.${banner_ext}" -O "${ASSET_PATH}/${audio_type}/${banner_local_name}.${banner_ext}"
done
