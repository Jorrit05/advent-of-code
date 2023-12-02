module Main

import IO;
import List;
import Set;
import String;

void main() {
    loc file = |project://Day13/test.txt|;

    pairList = [];
    list[str] pair = [];
    for (line <- readFileLines(file)) {
        pair += line;

        if(line := "") {
            pairList += <pair[0], pair[1]>;
            pair = [];
        }
    }
    // println(readFileLines(file));
}
