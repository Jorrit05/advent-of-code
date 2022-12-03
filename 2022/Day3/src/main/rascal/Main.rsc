module Main

import IO;
import List;
import String;

tuple[str,str] splitString(str S) {
  return <S[..(size(S)/2)], S[(size(S)/2)..]>;
}

// This probably could and should be done easier
str matchLetters(str S1, str S2){
    for(/<e:[a-zA-Z]>/ := S1){
        if (/<e>/ := S2) return e;
    }
    return "";
}

// This probably could and should be done easier
str matchLetters(str S1, str S2, str S3){
    for(/<e:[a-zA-Z]>/ := S1){
        if (/<e>/ := S2 && /<e>/ := S3) return e;
    }
    return "";
}

int getDecValue(str letter) {
    int asciiCode = head(chars(letter));

    if(asciiCode < 97) return asciiCode - 38;

    return asciiCode - 96;
}

list[tuple[str,str ,str]] getGroups([]) {return [<"","","">];}

list[tuple[str,str ,str]] getGroups(list[str] groupsList ){
    list[str] three = take(3, groupsList);
    tuple[str,str ,str] tupleOfThree =  <three[0], three[1], three[2]>;
    list[tuple[str,str ,str]] snd = getGroups(drop(3,groupsList));

    return  [tupleOfThree] + snd;
}

void main() {
    loc file = |project://Day3/input.txt|;
    list[str] fileLines = readFileLines(file);
    list[tuple[str,str]] rucksackList = [splitString(ruckSack) | ruckSack <- fileLines];
    list[str] matches = [ matchLetters(Fst,Snd)| <Fst, Snd> <- rucksackList];
    println(sum([getDecValue(match) | match <- matches ])); //8176

    // puzzle 2
    list[str] rucksackListTwo = [n | n <- fileLines];
    list[tuple[str,str ,str]]  groupsList = getGroups(rucksackListTwo);
    list[str] matchesTwo = [ matchLetters(Fst,Snd,Thrd)| <Fst, Snd,Thrd> <- delete(groupsList, size(groupsList) -1)];
    println(sum([getDecValue(match) | match <- matchesTwo ])); //2689
}
