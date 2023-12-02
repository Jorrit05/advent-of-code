module lib::Debug

import IO;
import lang::json::IO;

public void dumpToJson(str target, value values) {
    loc targetFile = |project://Day7/src/main/rascal|;
    targetFile += target;
    touch(targetFile);
    writeJSON(targetFile, values);
}
