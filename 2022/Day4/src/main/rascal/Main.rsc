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

tuple[int,int] puzzles(list[str] fileLines){
    int puzzle1 = 0;
    int puzzle2 = 0;

    for (line <- fileLines) {
        <range1, range2> = getSet(line);

        // Check for subset
        if(range1 <= range2 || range2 <= range1) {
            puzzle1 += 1;
        }

        // Check for size after a Set difference
        if(size(range1) != size(range1 - range2) ||
           size(range2) != size(range2 - range1)) {
            puzzle2 += 1;
        }
    }

    return <puzzle1, puzzle2>; //448, 794
}

void main() {
    loc file = |project://Day4/input.txt|;
    list[str] fileLines = readFileLines(file);
    println(puzzles(fileLines));
}
