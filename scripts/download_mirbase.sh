#!/bin/bash
set -euo pipefail

default_out_path="/mirbase"
out_path="${1:-$default_out_path}"
cd "$out_path"

fix() {
    sed -i 's/<br>/\n/g' $1
    sed -i 's/<\/p>//g' $1
    sed -i 's/<p>//g' $1
    sed -i '/^[[:space:]]*$/d' $1
    #sed -i 's/\&quot\;/"/g' $1
}

download() {
    printf "Downloading $1"
    wget -q -O $1 https://www.mirbase.org/download/CURRENT/database_files/$1
    fix $1
    size=$(du --block-size=K $1 | awk ' { print $1 } ')
    printf "  $size\n"
}

process() {

    for f in "$@"
    do
        download $f
    done
}

process confidence_score.txt \
 confidence.txt \
 dead_mirna.txt \
 literature_references.txt \
 mature_database_links.txt \
 mature_database_url.txt \
 mirna_2_prefam.txt \
 mirna_chromosome_build.txt \
 mirna_context.txt \
 mirna_database_links.txt \
 mirna_database_url.txt \
 mirna_literature_references.txt \
 mirna_pre_mature.txt \
 mirna_prefam.txt \
 mirna_mature.txt \
 mirna_species.txt \
 mirna.txt