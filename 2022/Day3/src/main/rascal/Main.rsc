module Main

import IO;
import List;
import String;

tuple[str,str] splitString(str S) {
  return <S[..(size(S)/2)], S[(size(S)/2)..]>;
}

// This probably could and should be done easier
str matchLetters(str S1, str S2){
    str retVal = "";
    for(/<e:[a-zA-Z]>/ := S1){
        if (/<e>/ := S2) {
            retVal = e;
        }
    }
    return retVal;
}

int getDecValue(str letter) {
    int asciiCode = head(chars(letter));
    if(asciiCode < 97){
        return asciiCode - 38;
    }

    return asciiCode - 96;
}

void main() {
    loc file = |project://Day3/input.txt|;
    list[str] fileLines = readFileLines(file);
    list[tuple[str,str]] rucksackList = [splitString(ruckSack) | ruckSack <- fileLines];
    list[str] matches = [ matchLetters(Fst,Snd)| <Fst, Snd> <- rucksackList];
    println(sum([getDecValue(match) | match <- matches ])); //8176
}
