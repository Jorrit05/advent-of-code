module Main

import IO;
import List;
import Set;
import String;

tuple[set[int],set[int]] getSet(str line) {
    list[str] splitLine = split(",", line);
    list[str] Fst = split("-", splitLine[0]);
    list[str] Snd = split("-", splitLine[1]);
    return <toSet([toInt(Fst[0])..toInt(Fst[1]) +1 ]), toSet([toInt(Snd[0])..toInt(Snd[1]) +1 ])>;
}

int puzzle1(list[str] fileLines){
    int counter = 0;
    for (line <- fileLines) {
        <range1, range2> = getSet(line);

        // Check for subset
        if(range1 <= range2 || range2 <= range1) {
            counter += 1;
        }
    }

    return counter; //448
}

int puzzle2(list[str] fileLines){
    int counter = 0;
    for (line <- fileLines) {
        <range1, range2> = getSet(line);

        // Check for size after a Set difference
        if(size(range1) != size(range1 - range2) ||
           size(range2) != size(range2 - range1)) {
            counter += 1;
        }
    }

    return counter; //794
}

void main() {
    loc file = |project://Day4/input.txt|;
    list[str] fileLines = readFileLines(file);
    println(puzzle1(fileLines));
    println(puzzle2(fileLines));
}